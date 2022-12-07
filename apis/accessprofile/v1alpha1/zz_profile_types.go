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

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type AllowedHostObservation struct {

	// ID of the host.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type AllowedHostParameters struct {

	// IPv4 address of the host
	// +kubebuilder:validation:Required
	Address *string `json:"address" tf:"address,omitempty"`

	// Description of the host.
	// +kubebuilder:validation:Required
	Description *string `json:"description" tf:"description,omitempty"`

	// Number of bits in the network mask.
	// +kubebuilder:validation:Required
	MaskBits *float64 `json:"maskBits" tf:"mask_bits,omitempty"`
}

type DNSServerObservation struct {

	// ID of the DNS server.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type DNSServerParameters struct {

	// Address of the DNS server.
	// +kubebuilder:validation:Required
	Address *string `json:"address" tf:"address,omitempty"`
}

type NtpServerObservation struct {

	// ID of the NTP server.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type NtpServerParameters struct {

	// Address of the NTP server.
	// +kubebuilder:validation:Required
	Address *string `json:"address" tf:"address,omitempty"`
}

type ProfileObservation struct {

	// List of allowed hosts.
	// +kubebuilder:validation:Optional
	AllowedHost []AllowedHostObservation `json:"allowedHost,omitempty" tf:"allowed_host,omitempty"`

	// The creator of the access profile.
	CreatedBy *string `json:"createdBy,omitempty" tf:"created_by,omitempty"`

	// List of DNS servers.
	// +kubebuilder:validation:Optional
	DNSServer []DNSServerObservation `json:"dnsServer,omitempty" tf:"dns_server,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The time and date of last modification.
	LastModified *string `json:"lastModified,omitempty" tf:"last_modified,omitempty"`

	// The last user to have modified the profile.
	LastModifiedBy *string `json:"lastModifiedBy,omitempty" tf:"last_modified_by,omitempty"`

	// List of NTP servers.
	// +kubebuilder:validation:Optional
	NtpServer []NtpServerObservation `json:"ntpServer,omitempty" tf:"ntp_server,omitempty"`

	// The name of the organization which owns the access profile.
	OrganizationName *string `json:"organizationName,omitempty" tf:"organization_name,omitempty"`

	// List of SSH users.
	// +kubebuilder:validation:Optional
	SSHUser []SSHUserObservation `json:"sshUser,omitempty" tf:"ssh_user,omitempty"`
}

type ProfileParameters struct {

	// List of allowed hosts.
	// +kubebuilder:validation:Optional
	AllowedHost []AllowedHostParameters `json:"allowedHost,omitempty" tf:"allowed_host,omitempty"`

	// List of DNS servers.
	// +kubebuilder:validation:Optional
	DNSServer []DNSServerParameters `json:"dnsServer,omitempty" tf:"dns_server,omitempty"`

	// HTTP proxy of the access profile.
	// +kubebuilder:validation:Optional
	HTTPProxy *string `json:"httpProxy,omitempty" tf:"http_proxy,omitempty"`

	// Indicates whether to lock the access profile. Defaults to `false`.
	// +kubebuilder:validation:Optional
	Lock *bool `json:"lock,omitempty" tf:"lock,omitempty"`

	// The name of the access profile.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// List of NTP servers.
	// +kubebuilder:validation:Optional
	NtpServer []NtpServerParameters `json:"ntpServer,omitempty" tf:"ntp_server,omitempty"`

	// The ID of the organization which owns the access profile.
	// +crossplane:generate:reference:type=github.com/itera-io/provider-jet-taikun/apis/organization/v1alpha1.Organization
	// +kubebuilder:validation:Optional
	OrganizationID *string `json:"organizationId,omitempty" tf:"organization_id,omitempty"`

	// Reference to a Organization in organization to populate organizationId.
	// +kubebuilder:validation:Optional
	OrganizationIDRef *v1.Reference `json:"organizationIdRef,omitempty" tf:"-"`

	// Selector for a Organization in organization to populate organizationId.
	// +kubebuilder:validation:Optional
	OrganizationIDSelector *v1.Selector `json:"organizationIdSelector,omitempty" tf:"-"`

	// List of SSH users.
	// +kubebuilder:validation:Optional
	SSHUser []SSHUserParameters `json:"sshUser,omitempty" tf:"ssh_user,omitempty"`
}

type SSHUserObservation struct {

	// ID of the SSH user.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type SSHUserParameters struct {

	// Name of the SSH user.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// Public key of the SSH user.
	// +kubebuilder:validation:Required
	PublicKey *string `json:"publicKey" tf:"public_key,omitempty"`
}

// ProfileSpec defines the desired state of Profile
type ProfileSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ProfileParameters `json:"forProvider"`
}

// ProfileStatus defines the observed state of Profile.
type ProfileStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ProfileObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Profile is the Schema for the Profiles API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,taikunjet}
type Profile struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ProfileSpec   `json:"spec"`
	Status            ProfileStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ProfileList contains a list of Profiles
type ProfileList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Profile `json:"items"`
}

// Repository type metadata.
var (
	Profile_Kind             = "Profile"
	Profile_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Profile_Kind}.String()
	Profile_KindAPIVersion   = Profile_Kind + "." + CRDGroupVersion.String()
	Profile_GroupVersionKind = CRDGroupVersion.WithKind(Profile_Kind)
)

func init() {
	SchemeBuilder.Register(&Profile{}, &ProfileList{})
}
