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

package clients

import (
	"context"
	"encoding/json"

	"github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/upbound/upjet/pkg/terraform"

	"github.com/itera-io/provider-jet-taikun/apis/v1alpha1"
)

const (
	// error messages
	errNoProviderConfig     = "no providerConfigRef provided"
	errGetProviderConfig    = "cannot get referenced ProviderConfig"
	errTrackUsage           = "cannot track ProviderConfig usage"
	errExtractCredentials   = "cannot extract credentials"
	errUnmarshalCredentials = "cannot unmarshal taikun credentials as JSON"
)

const (
	taikunEmail            = "email"
	taikunPassword         = "password"
	taikunAPIHost          = "api_host"
	taikunKeycloakEmail    = "keycloak_email"
	taikunKeycloakPassword = "keycloak_password"
)

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
		pc := &v1alpha1.ProviderConfig{}
		if err := client.Get(ctx, types.NamespacedName{Name: configRef.Name}, pc); err != nil {
			return ps, errors.Wrap(err, errGetProviderConfig)
		}

		t := resource.NewProviderConfigUsageTracker(client, &v1alpha1.ProviderConfigUsage{})
		if err := t.Track(ctx, mg); err != nil {
			return ps, errors.Wrap(err, errTrackUsage)
		}

		data, err := resource.CommonCredentialExtractor(ctx, pc.Spec.Credentials.Source, client, pc.Spec.Credentials.CommonCredentialSelectors)
		if err != nil {
			return ps, errors.Wrap(err, errExtractCredentials)
		}
		taikunCreds := map[string]string{}
		if err := json.Unmarshal(data, &taikunCreds); err != nil {
			return ps, errors.Wrap(err, errUnmarshalCredentials)
		}

		// set credentials in Terraform provider configuration
		ps.Configuration = CredentialsConfiguration(taikunCreds)
		return ps, nil
	}
}

// CredentialsConfiguration set credentials in Terraform provider configuration
func CredentialsConfiguration(taikunCreds map[string]string) terraform.ProviderConfiguration {
	tmp := map[string]interface{}{}
	if v, ok := taikunCreds[taikunEmail]; ok {
		tmp[taikunEmail] = v
	}
	if v, ok := taikunCreds[taikunPassword]; ok {
		tmp[taikunPassword] = v
	}
	if v, ok := taikunCreds[taikunAPIHost]; ok {
		tmp[taikunAPIHost] = v
	}
	if v, ok := taikunCreds[taikunKeycloakEmail]; ok {
		tmp[taikunKeycloakEmail] = v
	}
	if v, ok := taikunCreds[taikunKeycloakPassword]; ok {
		tmp[taikunKeycloakPassword] = v
	}
	return tmp
}
