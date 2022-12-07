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
// Code generated by angryjet. DO NOT EDIT.

package v1alpha1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	v1alpha1 "github.com/itera-io/provider-jet-taikun/apis/organization/v1alpha1"
	errors "github.com/pkg/errors"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this CredentialOpenstack.
func (mg *CredentialOpenstack) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.OrganizationID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.OrganizationIDRef,
		Selector:     mg.Spec.ForProvider.OrganizationIDSelector,
		To: reference.To{
			List:    &v1alpha1.OrganizationList{},
			Managed: &v1alpha1.Organization{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.OrganizationID")
	}
	mg.Spec.ForProvider.OrganizationID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.OrganizationIDRef = rsp.ResolvedReference

	return nil
}
