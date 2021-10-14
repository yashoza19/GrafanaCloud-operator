//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2021.

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
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GrafanaCloudLogForwarder) DeepCopyInto(out *GrafanaCloudLogForwarder) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GrafanaCloudLogForwarder.
func (in *GrafanaCloudLogForwarder) DeepCopy() *GrafanaCloudLogForwarder {
	if in == nil {
		return nil
	}
	out := new(GrafanaCloudLogForwarder)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GrafanaCloudLogForwarder) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GrafanaCloudLogForwarderList) DeepCopyInto(out *GrafanaCloudLogForwarderList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]GrafanaCloudLogForwarder, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GrafanaCloudLogForwarderList.
func (in *GrafanaCloudLogForwarderList) DeepCopy() *GrafanaCloudLogForwarderList {
	if in == nil {
		return nil
	}
	out := new(GrafanaCloudLogForwarderList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GrafanaCloudLogForwarderList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GrafanaCloudLogForwarderSpec) DeepCopyInto(out *GrafanaCloudLogForwarderSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GrafanaCloudLogForwarderSpec.
func (in *GrafanaCloudLogForwarderSpec) DeepCopy() *GrafanaCloudLogForwarderSpec {
	if in == nil {
		return nil
	}
	out := new(GrafanaCloudLogForwarderSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GrafanaCloudLogForwarderStatus) DeepCopyInto(out *GrafanaCloudLogForwarderStatus) {
	*out = *in
	if in.PodNames != nil {
		in, out := &in.PodNames, &out.PodNames
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GrafanaCloudLogForwarderStatus.
func (in *GrafanaCloudLogForwarderStatus) DeepCopy() *GrafanaCloudLogForwarderStatus {
	if in == nil {
		return nil
	}
	out := new(GrafanaCloudLogForwarderStatus)
	in.DeepCopyInto(out)
	return out
}