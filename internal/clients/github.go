package clients

import (
	"context"
	"encoding/json"

	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	"github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/crossplane/upjet/pkg/terraform"
	"github.com/k0kubun/pp/v3"
	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"

	hook "github.com/lacroi-m-insta/provider-github/apis/hook/v1alpha1"
)

const (
	// error messages
	errNoProviderConfig              = "no providerConfigRef provided"
	errGetProviderConfig             = "cannot get referenced ProviderConfig"
	errTrackUsage                    = "cannot track ProviderConfig usage"
	errExtractCredentials            = "cannot extract credentials"
	errUnmarshalCredentials          = "cannot unmarshal github credentials as JSON"
	errProviderConfigurationBuilder  = "cannot build configuration for terraform provider block"
	errTerraformProviderMissingOwner = "github provider app_auth needs owner key to be set"
	keyOwner                         = "owner"
	keyToken                         = "token"
)

type githubConfig struct {
	Owner *string `json:"owner,omitempty"`
	Token *string `json:"token,omitempty"`
}

func terraformProviderConfigurationBuilder(creds githubConfig) (terraform.ProviderConfiguration, error) {
	cnf := terraform.ProviderConfiguration{}

	if creds.Owner != nil {
		cnf[keyOwner] = *creds.Owner
	}

	if creds.Token != nil {
		cnf[keyToken] = *creds.Token
	}
	return cnf, nil
}

// TerraformSetupBuilder builds Terraform a terraform.SetupFn function which
// returns Terraform provider setup configuration
func TerraformSetupBuilder(version, providerSource, providerVersion string) terraform.SetupFn {
	return func(ctx context.Context, client client.Client, mg resource.Managed) (terraform.Setup, error) {
		ps := terraform.Setup{
			Version: version,
			Requirement: terraform.ProviderRequirement{
				Source:  providerSource,
				Version: providerVersion,
			},
		}

		configRef := mg.GetProviderConfigReference()
		if configRef == nil {
			return ps, errors.New(errNoProviderConfig)
		}

		pc := &hook.Webhook{}
		if err := client.Get(ctx, types.NamespacedName{Name: configRef.Name}, pc); err != nil {
			return ps, errors.Wrap(err, errGetProviderConfig)
		}

		// For secretSecretRef
		selectorCreds := xpv1.CommonCredentialSelectors{
			SecretRef: &xpv1.SecretKeySelector{
				SecretReference: xpv1.SecretReference{
					Name:      pc.Spec.ForProvider.Configuration[0].SecretSecretRef.Name,
					Namespace: pc.Spec.ForProvider.Configuration[0].SecretSecretRef.Namespace,
				},
				Key: pc.Spec.ForProvider.Configuration[0].SecretSecretRef.Key,
			},
		}

		// For urlSecretRef
		/*
			selectorUrl := xpv1.CommonCredentialSelectors{
				SecretRef: &xpv1.SecretKeySelector{
					SecretReference: xpv1.SecretReference{
						Name:      pc.Spec.ForProvider.Configuration[0].URLSecretRef.Name,
						Namespace: pc.Spec.ForProvider.Configuration[0].URLSecretRef.Namespace,
					},
					Key: pc.Spec.ForProvider.Configuration[0].URLSecretRef.Key,
				},
			}
		*/

		data, err := resource.CommonCredentialExtractor(ctx, xpv1.CredentialsSourceSecret, client, selectorCreds)
		if err != nil {
			return ps, errors.Wrap(err, errExtractCredentials)
		}

		pp.Println("===============================================")
		pp.Println(data)
		pp.Println("===============================================")

		creds := githubConfig{}
		if err := json.Unmarshal(data, &creds); err != nil {
			return ps, errors.Wrap(err, errUnmarshalCredentials)
		}

		ps.Configuration, err = terraformProviderConfigurationBuilder(creds)
		if err != nil {
			return ps, errors.Wrap(err, errProviderConfigurationBuilder)
		}
		return ps, nil
	}
}
