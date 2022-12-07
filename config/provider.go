/*
Copyright 2021 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	tjconfig "github.com/upbound/upjet/pkg/config"

	resources "github.com/itera-io/provider-jet-taikun/config/resources"
)

const (
	resourcePrefix = "taikun"
	modulePath     = "github.com/itera-io/provider-jet-taikun"
)

//go:embed schema.json
var providerSchema string

//go:embed provider-metadata.yaml
var providerMetadata string

// GetProvider returns provider configuration
func GetProvider() *tjconfig.Provider {

	pc := tjconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		tjconfig.WithDefaultResourceOptions(),
		tjconfig.WithBasePackages(tjconfig.BasePackages{
			APIVersion: []string{
				// Default package for ProviderConfig APIs
				"apis/v1alpha1",
			},
			Controller: []string{
				// Default package for ProviderConfig controllers
				"internal/controller/providerconfig",
			},
		}),
		tjconfig.WithIncludeList([]string{

			// Resources
			"taikun_organization$",
			"taikun_access_profile$",
			"taikun_alerting_profile$",
			"taikun_kubernetes_profile$",
			"taikun_policy_profile$",
			"taikun_slack_configuration$",
			"taikun_cloud_credential",
			"taikun_cloud_credential_aws",
			"taikun_cloud_credential_azure",
			"taikun_cloud_credential_gcp",
			"taikun_cloud_credential_openstack",
			"taikun_showback_credential",
			"taikun_showback_rule",
			"taikun_billing_credential",
			"taikun_billing_rule",
			"taikun_user$",
			"taikun_standalone_profile$",
			"taikun_backup_policy$",
			"taikun_backup_credential$",
			"taikun_project$",
			"taikun_organization_billing_rule_attachment$",
			"taikun_kubeconfig",
			"taikun_project_user_attachment",

			// Data Sources
		}))

	for _, configure := range []func(provider *tjconfig.Provider){
		// add custom config functions

		// Resources
		resources.Configure,

		// Data
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
