//go:build !ignore_autogenerated
// +build !ignore_autogenerated

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

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"github.com/crossplane/crossplane-runtime/apis/common/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Kubeconfig) DeepCopyInto(out *Kubeconfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Kubeconfig.
func (in *Kubeconfig) DeepCopy() *Kubeconfig {
	if in == nil {
		return nil
	}
	out := new(Kubeconfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Kubeconfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeconfigList) DeepCopyInto(out *KubeconfigList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Kubeconfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeconfigList.
func (in *KubeconfigList) DeepCopy() *KubeconfigList {
	if in == nil {
		return nil
	}
	out := new(KubeconfigList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KubeconfigList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeconfigObservation) DeepCopyInto(out *KubeconfigObservation) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.ProjectName != nil {
		in, out := &in.ProjectName, &out.ProjectName
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeconfigObservation.
func (in *KubeconfigObservation) DeepCopy() *KubeconfigObservation {
	if in == nil {
		return nil
	}
	out := new(KubeconfigObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeconfigParameters) DeepCopyInto(out *KubeconfigParameters) {
	*out = *in
	if in.AccessScope != nil {
		in, out := &in.AccessScope, &out.AccessScope
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Namespace != nil {
		in, out := &in.Namespace, &out.Namespace
		*out = new(string)
		**out = **in
	}
	if in.ProjectID != nil {
		in, out := &in.ProjectID, &out.ProjectID
		*out = new(string)
		**out = **in
	}
	if in.ProjectIDRef != nil {
		in, out := &in.ProjectIDRef, &out.ProjectIDRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.ProjectIDSelector != nil {
		in, out := &in.ProjectIDSelector, &out.ProjectIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.Role != nil {
		in, out := &in.Role, &out.Role
		*out = new(string)
		**out = **in
	}
	if in.UserID != nil {
		in, out := &in.UserID, &out.UserID
		*out = new(string)
		**out = **in
	}
	if in.UserIDRef != nil {
		in, out := &in.UserIDRef, &out.UserIDRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.UserIDSelector != nil {
		in, out := &in.UserIDSelector, &out.UserIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.ValidityPeriod != nil {
		in, out := &in.ValidityPeriod, &out.ValidityPeriod
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeconfigParameters.
func (in *KubeconfigParameters) DeepCopy() *KubeconfigParameters {
	if in == nil {
		return nil
	}
	out := new(KubeconfigParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeconfigSpec) DeepCopyInto(out *KubeconfigSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeconfigSpec.
func (in *KubeconfigSpec) DeepCopy() *KubeconfigSpec {
	if in == nil {
		return nil
	}
	out := new(KubeconfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeconfigStatus) DeepCopyInto(out *KubeconfigStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeconfigStatus.
func (in *KubeconfigStatus) DeepCopy() *KubeconfigStatus {
	if in == nil {
		return nil
	}
	out := new(KubeconfigStatus)
	in.DeepCopyInto(out)
	return out
}
