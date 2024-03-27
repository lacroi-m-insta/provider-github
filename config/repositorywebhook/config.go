package repositorywebhook

import "github.com/crossplane/upjet/pkg/config"

// Configure github_repository_webhook resource.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("github_repository_webhook", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "github"
		r.ShortGroup = "hook"

		// r.LateInitializer = config.LateInitializer{IgnoredFields: []string{"private", "default_branch"}}
	})
}
