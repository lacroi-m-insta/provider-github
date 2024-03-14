// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type ConfigurationInitParameters struct {

	// The content type for the payload. Valid values are either form or json.
	// The content type for the payload. Valid values are either 'form' or 'json'.
	ContentType *string `json:"contentType,omitempty" tf:"content_type,omitempty"`

	// Url for the callback webhook.
	// Url for the callback webhook.
	Url *string `json:"url,omitempty" tf:"url,omitempty"`

	// Insecure SSL boolean toggle. Defaults to false.
	// Insecure SSL boolean toggle. Defaults to 'false'.
	InsecureSSL *bool `json:"insecureSsl,omitempty" tf:"insecure_ssl,omitempty"`
}

type ConfigurationObservation struct {

	// The content type for the payload. Valid values are either form or json.
	// The content type for the payload. Valid values are either 'form' or 'json'.
	ContentType *string `json:"contentType,omitempty" tf:"content_type,omitempty"`

	// Url for the callback webhook.
	// Url for the callback webhook. 
	Url *string `json:"url,omitempty" tf:"url,omitempty"`

	// Insecure SSL boolean toggle. Defaults to false.
	// Insecure SSL boolean toggle. Defaults to 'false'.
	InsecureSSL *bool `json:"insecureSsl,omitempty" tf:"insecure_ssl,omitempty"`
}

type ConfigurationParameters struct {

	// The content type for the payload. Valid values are either form or json.
	// The content type for the payload. Valid values are either 'form' or 'json'.
	// +kubebuilder:validation:Optional
	ContentType *string `json:"contentType,omitempty" tf:"content_type,omitempty"`

	// Insecure SSL boolean toggle. Defaults to false.
	// Insecure SSL boolean toggle. Defaults to 'false'.
	// +kubebuilder:validation:Optional
	InsecureSSL *bool `json:"insecureSsl,omitempty" tf:"insecure_ssl,omitempty"`

	// The shared secret for the webhook. See API documentation.
	// The shared secret for the webhook
	// +kubebuilder:validation:Optional
	SecretSecretRef *v1.SecretKeySelector `json:"secretSecretRef,omitempty" tf:"-"`

	// The URL of the webhook.
	// The URL of the webhook.
	// +kubebuilder:validation:Required
	URLSecretRef v1.SecretKeySelector `json:"urlSecretRef" tf:"-"`
}

type WebhookInitParameters struct {

	// Indicate if the webhook should receive events. Defaults to true.
	// Indicate if the webhook should receive events. Defaults to 'true'.
	Active *bool `json:"active,omitempty" tf:"active,omitempty"`

	// Configuration block for the webhook. Detailed below.
	// Configuration for the webhook.
	Configuration []ConfigurationInitParameters `json:"configuration,omitempty" tf:"configuration,omitempty"`

	// A list of events which should trigger the webhook. See a list of available events.
	// A list of events which should trigger the webhook
	Events []*string `json:"events,omitempty" tf:"events,omitempty"`

	// The repository of the webhook.
	// The repository of the webhook.
	Repository *string `json:"repository,omitempty" tf:"repository,omitempty"`
}

type WebhookObservation struct {

	// Indicate if the webhook should receive events. Defaults to true.
	// Indicate if the webhook should receive events. Defaults to 'true'.
	Active *bool `json:"active,omitempty" tf:"active,omitempty"`

	// Configuration block for the webhook. Detailed below.
	// Configuration for the webhook.
	Configuration []ConfigurationObservation `json:"configuration,omitempty" tf:"configuration,omitempty"`

	Etag *string `json:"etag,omitempty" tf:"etag,omitempty"`

	// A list of events which should trigger the webhook. See a list of available events.
	// A list of events which should trigger the webhook
	Events []*string `json:"events,omitempty" tf:"events,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The repository of the webhook.
	// The repository of the webhook.
	Repository *string `json:"repository,omitempty" tf:"repository,omitempty"`

	// URL of the webhook.  This is a sensitive attribute because it may include basic auth credentials.
	// Configuration block for the webhook
	URL *string `json:"url,omitempty" tf:"url,omitempty"`
}

type WebhookParameters struct {

	// Indicate if the webhook should receive events. Defaults to true.
	// Indicate if the webhook should receive events. Defaults to 'true'.
	// +kubebuilder:validation:Optional
	Active *bool `json:"active,omitempty" tf:"active,omitempty"`

	// Configuration block for the webhook. Detailed below.
	// Configuration for the webhook.
	// +kubebuilder:validation:Optional
	Configuration []ConfigurationParameters `json:"configuration,omitempty" tf:"configuration,omitempty"`

	// A list of events which should trigger the webhook. See a list of available events.
	// A list of events which should trigger the webhook
	// +kubebuilder:validation:Optional
	Events []*string `json:"events,omitempty" tf:"events,omitempty"`

	// The repository of the webhook.
	// The repository of the webhook.
	// +kubebuilder:validation:Optional
	Repository *string `json:"repository,omitempty" tf:"repository,omitempty"`
}

// WebhookSpec defines the desired state of Webhook
type WebhookSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     WebhookParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider WebhookInitParameters `json:"initProvider,omitempty"`
}

// WebhookStatus defines the observed state of Webhook.
type WebhookStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        WebhookObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Webhook is the Schema for the Webhooks API. Creates and manages repository webhooks within GitHub organizations or personal accounts
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,github}
type Webhook struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.events) || (has(self.initProvider) && has(self.initProvider.events))",message="spec.forProvider.events is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.repository) || (has(self.initProvider) && has(self.initProvider.repository))",message="spec.forProvider.repository is a required parameter"
	Spec   WebhookSpec   `json:"spec"`
	Status WebhookStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// WebhookList contains a list of Webhooks
type WebhookList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Webhook `json:"items"`
}

// Repository type metadata.
var (
	Webhook_Kind             = "Webhook"
	Webhook_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Webhook_Kind}.String()
	Webhook_KindAPIVersion   = Webhook_Kind + "." + CRDGroupVersion.String()
	Webhook_GroupVersionKind = CRDGroupVersion.WithKind(Webhook_Kind)
)

func init() {
	SchemeBuilder.Register(&Webhook{}, &WebhookList{})
}
