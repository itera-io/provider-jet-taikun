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

type DiskObservation struct {

	// ID of the disk.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type DiskParameters struct {

	// Name of the device (required with AWS).
	// +kubebuilder:validation:Optional
	DeviceName *string `json:"deviceName,omitempty" tf:"device_name,omitempty"`

	// LUN ID (required with Azure).
	// +kubebuilder:validation:Optional
	LunID *float64 `json:"lunId,omitempty" tf:"lun_id,omitempty"`

	// Name of the disk.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// The disk size in GBs.
	// +kubebuilder:validation:Required
	Size *float64 `json:"size" tf:"size,omitempty"`

	// Type of the volume (only valid with OpenStack).
	// +kubebuilder:validation:Optional
	VolumeType *string `json:"volumeType,omitempty" tf:"volume_type,omitempty"`
}

type KubernetesNodeLabelObservation struct {
}

type KubernetesNodeLabelParameters struct {

	// Kubernetes node label key.
	// +kubebuilder:validation:Required
	Key *string `json:"key" tf:"key,omitempty"`

	// Kubernetes node label value.
	// +kubebuilder:validation:Required
	Value *string `json:"value" tf:"value,omitempty"`
}

type ProjectObservation struct {

	// Public IP address of the bastion.
	AccessIP *string `json:"accessIp,omitempty" tf:"access_ip,omitempty"`

	// Name of the project's alerting profile.
	AlertingProfileName *string `json:"alertingProfileName,omitempty" tf:"alerting_profile_name,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Bastion server. Required with: `server_kubemaster`, `server_kubeworker`.
	// +kubebuilder:validation:Optional
	ServerBastion []ServerBastionObservation `json:"serverBastion,omitempty" tf:"server_bastion,omitempty"`

	// Kubemaster server. Required with: `server_bastion`, `server_kubeworker`.
	// +kubebuilder:validation:Optional
	ServerKubemaster []ServerKubemasterObservation `json:"serverKubemaster,omitempty" tf:"server_kubemaster,omitempty"`

	// Kubeworker server. Required with: `server_bastion`, `server_kubemaster`.
	// +kubebuilder:validation:Optional
	ServerKubeworker []ServerKubeworkerObservation `json:"serverKubeworker,omitempty" tf:"server_kubeworker,omitempty"`

	// Virtual machines.
	// +kubebuilder:validation:Optional
	VM []VMObservation `json:"vm,omitempty" tf:"vm,omitempty"`
}

type ProjectParameters struct {

	// ID of the project's access profile. Defaults to the default access profile of the project's organization.
	// +crossplane:generate:reference:type=github.com/itera-io/provider-jet-taikun/apis/accessprofile/v1alpha1.Profile
	// +kubebuilder:validation:Optional
	AccessProfileID *string `json:"accessProfileId,omitempty" tf:"access_profile_id,omitempty"`

	// Reference to a Profile in accessprofile to populate accessProfileId.
	// +kubebuilder:validation:Optional
	AccessProfileIDRef *v1.Reference `json:"accessProfileIdRef,omitempty" tf:"-"`

	// Selector for a Profile in accessprofile to populate accessProfileId.
	// +kubebuilder:validation:Optional
	AccessProfileIDSelector *v1.Selector `json:"accessProfileIdSelector,omitempty" tf:"-"`

	// ID of the project's alerting profile.
	// +crossplane:generate:reference:type=github.com/itera-io/provider-jet-taikun/apis/alertingprofile/v1alpha1.Profile
	// +kubebuilder:validation:Optional
	AlertingProfileID *string `json:"alertingProfileId,omitempty" tf:"alerting_profile_id,omitempty"`

	// Reference to a Profile in alertingprofile to populate alertingProfileId.
	// +kubebuilder:validation:Optional
	AlertingProfileIDRef *v1.Reference `json:"alertingProfileIdRef,omitempty" tf:"-"`

	// Selector for a Profile in alertingprofile to populate alertingProfileId.
	// +kubebuilder:validation:Optional
	AlertingProfileIDSelector *v1.Selector `json:"alertingProfileIdSelector,omitempty" tf:"-"`

	// If enabled, the Kubespray version will be automatically upgraded when a new version is available. Defaults to `false`.
	// +kubebuilder:validation:Optional
	AutoUpgrade *bool `json:"autoUpgrade,omitempty" tf:"auto_upgrade,omitempty"`

	// ID of the backup credential. If unspecified, backups are disabled.
	// +crossplane:generate:reference:type=github.com/itera-io/provider-jet-taikun/apis/backupcredential/v1alpha1.Credential
	// +kubebuilder:validation:Optional
	BackupCredentialID *string `json:"backupCredentialId,omitempty" tf:"backup_credential_id,omitempty"`

	// Reference to a Credential in backupcredential to populate backupCredentialId.
	// +kubebuilder:validation:Optional
	BackupCredentialIDRef *v1.Reference `json:"backupCredentialIdRef,omitempty" tf:"-"`

	// Selector for a Credential in backupcredential to populate backupCredentialId.
	// +kubebuilder:validation:Optional
	BackupCredentialIDSelector *v1.Selector `json:"backupCredentialIdSelector,omitempty" tf:"-"`

	// ID of the cloud credential used to create the project's servers.
	// +crossplane:generate:reference:type=github.com/itera-io/provider-jet-taikun/apis/cloudcredential/v1alpha1.Credential
	// +kubebuilder:validation:Optional
	CloudCredentialID *string `json:"cloudCredentialId,omitempty" tf:"cloud_credential_id,omitempty"`

	// Reference to a Credential in cloudcredential to populate cloudCredentialId.
	// +kubebuilder:validation:Optional
	CloudCredentialIDRef *v1.Reference `json:"cloudCredentialIdRef,omitempty" tf:"-"`

	// Selector for a Credential in cloudcredential to populate cloudCredentialId.
	// +kubebuilder:validation:Optional
	CloudCredentialIDSelector *v1.Selector `json:"cloudCredentialIdSelector,omitempty" tf:"-"`

	// If enabled, the project will be deleted on the expiration date and it will not be possible to recover it. Defaults to `false`. Required with: `expiration_date`.
	// +kubebuilder:validation:Optional
	DeleteOnExpiration *bool `json:"deleteOnExpiration,omitempty" tf:"delete_on_expiration,omitempty"`

	// Project's expiration date in the format: 'dd/mm/yyyy'.
	// +kubebuilder:validation:Optional
	ExpirationDate *string `json:"expirationDate,omitempty" tf:"expiration_date,omitempty"`

	// List of flavors bound to the project.
	// +kubebuilder:validation:Optional
	Flavors []*string `json:"flavors,omitempty" tf:"flavors,omitempty"`

	// List of images bound to the project.
	// +kubebuilder:validation:Optional
	Images []*string `json:"images,omitempty" tf:"images,omitempty"`

	// ID of the project's Kubernetes profile. Defaults to the default Kubernetes profile of the project's organization.
	// +crossplane:generate:reference:type=github.com/itera-io/provider-jet-taikun/apis/kubernetesprofile/v1alpha1.Profile
	// +kubebuilder:validation:Optional
	KubernetesProfileID *string `json:"kubernetesProfileId,omitempty" tf:"kubernetes_profile_id,omitempty"`

	// Reference to a Profile in kubernetesprofile to populate kubernetesProfileId.
	// +kubebuilder:validation:Optional
	KubernetesProfileIDRef *v1.Reference `json:"kubernetesProfileIdRef,omitempty" tf:"-"`

	// Selector for a Profile in kubernetesprofile to populate kubernetesProfileId.
	// +kubebuilder:validation:Optional
	KubernetesProfileIDSelector *v1.Selector `json:"kubernetesProfileIdSelector,omitempty" tf:"-"`

	// Kubernetes Version at project creation. Use the meta-argument `ignore_changes` to ignore future upgrades.
	// +kubebuilder:validation:Optional
	KubernetesVersion *string `json:"kubernetesVersion,omitempty" tf:"kubernetes_version,omitempty"`

	// Indicates whether to lock the project. Defaults to `false`.
	// +kubebuilder:validation:Optional
	Lock *bool `json:"lock,omitempty" tf:"lock,omitempty"`

	// Kubernetes cluster monitoring. Defaults to `false`.
	// +kubebuilder:validation:Optional
	Monitoring *bool `json:"monitoring,omitempty" tf:"monitoring,omitempty"`

	// Project name.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// ID of the organization which owns the project.
	// +crossplane:generate:reference:type=github.com/itera-io/provider-jet-taikun/apis/organization/v1alpha1.Organization
	// +kubebuilder:validation:Optional
	OrganizationID *string `json:"organizationId,omitempty" tf:"organization_id,omitempty"`

	// Reference to a Organization in organization to populate organizationId.
	// +kubebuilder:validation:Optional
	OrganizationIDRef *v1.Reference `json:"organizationIdRef,omitempty" tf:"-"`

	// Selector for a Organization in organization to populate organizationId.
	// +kubebuilder:validation:Optional
	OrganizationIDSelector *v1.Selector `json:"organizationIdSelector,omitempty" tf:"-"`

	// ID of the Policy profile. If unspecified, Gatekeeper is disabled.
	// +crossplane:generate:reference:type=github.com/itera-io/provider-jet-taikun/apis/policyprofile/v1alpha1.Profile
	// +kubebuilder:validation:Optional
	PolicyProfileID *string `json:"policyProfileId,omitempty" tf:"policy_profile_id,omitempty"`

	// Reference to a Profile in policyprofile to populate policyProfileId.
	// +kubebuilder:validation:Optional
	PolicyProfileIDRef *v1.Reference `json:"policyProfileIdRef,omitempty" tf:"-"`

	// Selector for a Profile in policyprofile to populate policyProfileId.
	// +kubebuilder:validation:Optional
	PolicyProfileIDSelector *v1.Selector `json:"policyProfileIdSelector,omitempty" tf:"-"`

	// Maximum CPU units. Defaults to `1000000`.
	// +kubebuilder:validation:Optional
	QuotaCPUUnits *float64 `json:"quotaCpuUnits,omitempty" tf:"quota_cpu_units,omitempty"`

	// Maximum disk size in GBs. Defaults to `102400`.
	// +kubebuilder:validation:Optional
	QuotaDiskSize *float64 `json:"quotaDiskSize,omitempty" tf:"quota_disk_size,omitempty"`

	// Maximum RAM size in GBs. Defaults to `102400`.
	// +kubebuilder:validation:Optional
	QuotaRAMSize *float64 `json:"quotaRamSize,omitempty" tf:"quota_ram_size,omitempty"`

	// Maximum CPU units for standalone VMs. Defaults to `1000000`.
	// +kubebuilder:validation:Optional
	QuotaVMCPUUnits *float64 `json:"quotaVmCpuUnits,omitempty" tf:"quota_vm_cpu_units,omitempty"`

	// Maximum RAM size in GBs for standalone VMs. Defaults to `102400`.
	// +kubebuilder:validation:Optional
	QuotaVMRAMSize *float64 `json:"quotaVmRamSize,omitempty" tf:"quota_vm_ram_size,omitempty"`

	// Maximum volume size in GBs for standalone VMs. Defaults to `102400`.
	// +kubebuilder:validation:Optional
	QuotaVMVolumeSize *float64 `json:"quotaVmVolumeSize,omitempty" tf:"quota_vm_volume_size,omitempty"`

	// Router ID end range (specify only if using OpenStack cloud credentials with Taikun Load Balancer enabled). Required with: `router_id_start_range`, `taikun_lb_flavor`.
	// +kubebuilder:validation:Optional
	RouterIDEndRange *float64 `json:"routerIdEndRange,omitempty" tf:"router_id_end_range,omitempty"`

	// Router ID start range (specify only if using OpenStack cloud credentials with Taikun Load Balancer enabled). Required with: `router_id_end_range`, `taikun_lb_flavor`.
	// +kubebuilder:validation:Optional
	RouterIDStartRange *float64 `json:"routerIdStartRange,omitempty" tf:"router_id_start_range,omitempty"`

	// Bastion server. Required with: `server_kubemaster`, `server_kubeworker`.
	// +kubebuilder:validation:Optional
	ServerBastion []ServerBastionParameters `json:"serverBastion,omitempty" tf:"server_bastion,omitempty"`

	// Kubemaster server. Required with: `server_bastion`, `server_kubeworker`.
	// +kubebuilder:validation:Optional
	ServerKubemaster []ServerKubemasterParameters `json:"serverKubemaster,omitempty" tf:"server_kubemaster,omitempty"`

	// Kubeworker server. Required with: `server_bastion`, `server_kubemaster`.
	// +kubebuilder:validation:Optional
	ServerKubeworker []ServerKubeworkerParameters `json:"serverKubeworker,omitempty" tf:"server_kubeworker,omitempty"`

	// OpenStack flavor for the Taikun load balancer (specify only if using OpenStack cloud credentials with Taikun Load Balancer enabled). Required with: `router_id_end_range`, `router_id_start_range`.
	// +kubebuilder:validation:Optional
	TaikunLBFlavor *string `json:"taikunLbFlavor,omitempty" tf:"taikun_lb_flavor,omitempty"`

	// Virtual machines.
	// +kubebuilder:validation:Optional
	VM []VMParameters `json:"vm,omitempty" tf:"vm,omitempty"`
}

type ServerBastionObservation struct {

	// The creator of the server.
	CreatedBy *string `json:"createdBy,omitempty" tf:"created_by,omitempty"`

	// ID of the server.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// IP of the server.
	IP *string `json:"ip,omitempty" tf:"ip,omitempty"`

	// The time and date of last modification.
	LastModified *string `json:"lastModified,omitempty" tf:"last_modified,omitempty"`

	// The last user to have modified the server.
	LastModifiedBy *string `json:"lastModifiedBy,omitempty" tf:"last_modified_by,omitempty"`

	// Server status.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`
}

type ServerBastionParameters struct {

	// The server's disk size in GBs. Defaults to `30`.
	// +kubebuilder:validation:Optional
	DiskSize *float64 `json:"diskSize,omitempty" tf:"disk_size,omitempty"`

	// The server's flavor.
	// +kubebuilder:validation:Required
	Flavor *string `json:"flavor" tf:"flavor,omitempty"`

	// Name of the server.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`
}

type ServerKubemasterObservation struct {

	// The creator of the server.
	CreatedBy *string `json:"createdBy,omitempty" tf:"created_by,omitempty"`

	// ID of the server.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// IP of the server.
	IP *string `json:"ip,omitempty" tf:"ip,omitempty"`

	// The time and date of last modification.
	LastModified *string `json:"lastModified,omitempty" tf:"last_modified,omitempty"`

	// The last user to have modified the server.
	LastModifiedBy *string `json:"lastModifiedBy,omitempty" tf:"last_modified_by,omitempty"`

	// Server status.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`
}

type ServerKubemasterParameters struct {

	// The server's disk size in GBs. Defaults to `30`.
	// +kubebuilder:validation:Optional
	DiskSize *float64 `json:"diskSize,omitempty" tf:"disk_size,omitempty"`

	// The server's flavor.
	// +kubebuilder:validation:Required
	Flavor *string `json:"flavor" tf:"flavor,omitempty"`

	// Attach Kubernetes node labels.
	// +kubebuilder:validation:Optional
	KubernetesNodeLabel []KubernetesNodeLabelParameters `json:"kubernetesNodeLabel,omitempty" tf:"kubernetes_node_label,omitempty"`

	// Name of the server.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`
}

type ServerKubeworkerKubernetesNodeLabelObservation struct {
}

type ServerKubeworkerKubernetesNodeLabelParameters struct {

	// Kubernetes node label key.
	// +kubebuilder:validation:Required
	Key *string `json:"key" tf:"key,omitempty"`

	// Kubernetes node label value.
	// +kubebuilder:validation:Required
	Value *string `json:"value" tf:"value,omitempty"`
}

type ServerKubeworkerObservation struct {

	// The creator of the server.
	CreatedBy *string `json:"createdBy,omitempty" tf:"created_by,omitempty"`

	// ID of the server.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// IP of the server.
	IP *string `json:"ip,omitempty" tf:"ip,omitempty"`

	// The time and date of last modification.
	LastModified *string `json:"lastModified,omitempty" tf:"last_modified,omitempty"`

	// The last user to have modified the server.
	LastModifiedBy *string `json:"lastModifiedBy,omitempty" tf:"last_modified_by,omitempty"`

	// Server status.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`
}

type ServerKubeworkerParameters struct {

	// The server's disk size in GBs. Defaults to `30`.
	// +kubebuilder:validation:Optional
	DiskSize *float64 `json:"diskSize,omitempty" tf:"disk_size,omitempty"`

	// The server's flavor.
	// +kubebuilder:validation:Required
	Flavor *string `json:"flavor" tf:"flavor,omitempty"`

	// Attach Kubernetes node labels.
	// +kubebuilder:validation:Optional
	KubernetesNodeLabel []ServerKubeworkerKubernetesNodeLabelParameters `json:"kubernetesNodeLabel,omitempty" tf:"kubernetes_node_label,omitempty"`

	// Name of the server.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`
}

type TagObservation struct {
}

type TagParameters struct {

	// Key of the tag.
	// +kubebuilder:validation:Required
	Key *string `json:"key" tf:"key,omitempty"`

	// Value of the tag.
	// +kubebuilder:validation:Required
	Value *string `json:"value" tf:"value,omitempty"`
}

type VMObservation struct {

	// Access IP of the VM.
	AccessIP *string `json:"accessIp,omitempty" tf:"access_ip,omitempty"`

	// The creator of the VM.
	CreatedBy *string `json:"createdBy,omitempty" tf:"created_by,omitempty"`

	// Disks associated with the VM.
	// +kubebuilder:validation:Optional
	Disk []DiskObservation `json:"disk,omitempty" tf:"disk,omitempty"`

	// ID of the VM.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// IP of the VM.
	IP *string `json:"ip,omitempty" tf:"ip,omitempty"`

	// The VM's image name.
	ImageName *string `json:"imageName,omitempty" tf:"image_name,omitempty"`

	// The time and date of last modification.
	LastModified *string `json:"lastModified,omitempty" tf:"last_modified,omitempty"`

	// The last user to have modified the VM.
	LastModifiedBy *string `json:"lastModifiedBy,omitempty" tf:"last_modified_by,omitempty"`

	// VM status.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`
}

type VMParameters struct {

	// Cloud init (updating this field will recreate the VM). Defaults to ` `.
	// +kubebuilder:validation:Optional
	CloudInit *string `json:"cloudInit,omitempty" tf:"cloud_init,omitempty"`

	// Disks associated with the VM.
	// +kubebuilder:validation:Optional
	Disk []DiskParameters `json:"disk,omitempty" tf:"disk,omitempty"`

	// The VM's flavor.
	// +kubebuilder:validation:Required
	Flavor *string `json:"flavor" tf:"flavor,omitempty"`

	// The VM's image ID (updating this field will recreate the VM).
	// +kubebuilder:validation:Required
	ImageID *string `json:"imageId" tf:"image_id,omitempty"`

	// Name of the VM (updating this field will recreate the VM).
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// Whether a public IP will be available (updating this field will recreate the VM if the project isn't hosted on OpenStack). Defaults to `false`.
	// +kubebuilder:validation:Optional
	PublicIP *bool `json:"publicIp,omitempty" tf:"public_ip,omitempty"`

	// Standalone profile ID bound to the VM (updating this field will recreate the VM).
	// +kubebuilder:validation:Required
	StandaloneProfileID *string `json:"standaloneProfileId" tf:"standalone_profile_id,omitempty"`

	// Tags linked to the VM (updating this field will recreate the VM).
	// +kubebuilder:validation:Optional
	Tag []TagParameters `json:"tag,omitempty" tf:"tag,omitempty"`

	// The VM's username (required for Azure).
	// +kubebuilder:validation:Optional
	Username *string `json:"username,omitempty" tf:"username,omitempty"`

	// The VM's volume size in GBs (updating this field will recreate the VM).
	// +kubebuilder:validation:Required
	VolumeSize *float64 `json:"volumeSize" tf:"volume_size,omitempty"`

	// Volume type (updating this field will recreate the VM).
	// +kubebuilder:validation:Optional
	VolumeType *string `json:"volumeType,omitempty" tf:"volume_type,omitempty"`
}

// ProjectSpec defines the desired state of Project
type ProjectSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ProjectParameters `json:"forProvider"`
}

// ProjectStatus defines the observed state of Project.
type ProjectStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ProjectObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Project is the Schema for the Projects API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,taikunjet}
type Project struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ProjectSpec   `json:"spec"`
	Status            ProjectStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ProjectList contains a list of Projects
type ProjectList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Project `json:"items"`
}

// Repository type metadata.
var (
	Project_Kind             = "Project"
	Project_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Project_Kind}.String()
	Project_KindAPIVersion   = Project_Kind + "." + CRDGroupVersion.String()
	Project_GroupVersionKind = CRDGroupVersion.WithKind(Project_Kind)
)

func init() {
	SchemeBuilder.Register(&Project{}, &ProjectList{})
}
