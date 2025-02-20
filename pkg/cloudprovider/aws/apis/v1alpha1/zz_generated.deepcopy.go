//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
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
	"github.com/aws/karpenter/pkg/apis/provisioning/v1alpha5"
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AWS) DeepCopyInto(out *AWS) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	if in.LaunchTemplate != nil {
		in, out := &in.LaunchTemplate, &out.LaunchTemplate
		*out = new(string)
		**out = **in
	}
	if in.SubnetSelector != nil {
		in, out := &in.SubnetSelector, &out.SubnetSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.SecurityGroupSelector != nil {
		in, out := &in.SecurityGroupSelector, &out.SecurityGroupSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.MetadataOptions != nil {
		in, out := &in.MetadataOptions, &out.MetadataOptions
		*out = new(MetadataOptions)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AWS.
func (in *AWS) DeepCopy() *AWS {
	if in == nil {
		return nil
	}
	out := new(AWS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AWS) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Constraints) DeepCopyInto(out *Constraints) {
	*out = *in
	if in.Constraints != nil {
		in, out := &in.Constraints, &out.Constraints
		*out = new(v1alpha5.Constraints)
		(*in).DeepCopyInto(*out)
	}
	if in.AWS != nil {
		in, out := &in.AWS, &out.AWS
		*out = new(AWS)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Constraints.
func (in *Constraints) DeepCopy() *Constraints {
	if in == nil {
		return nil
	}
	out := new(Constraints)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MetadataOptions) DeepCopyInto(out *MetadataOptions) {
	*out = *in
	if in.HTTPEndpoint != nil {
		in, out := &in.HTTPEndpoint, &out.HTTPEndpoint
		*out = new(string)
		**out = **in
	}
	if in.HTTPProtocolIPv6 != nil {
		in, out := &in.HTTPProtocolIPv6, &out.HTTPProtocolIPv6
		*out = new(string)
		**out = **in
	}
	if in.HTTPPutResponseHopLimit != nil {
		in, out := &in.HTTPPutResponseHopLimit, &out.HTTPPutResponseHopLimit
		*out = new(int64)
		**out = **in
	}
	if in.HTTPTokens != nil {
		in, out := &in.HTTPTokens, &out.HTTPTokens
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MetadataOptions.
func (in *MetadataOptions) DeepCopy() *MetadataOptions {
	if in == nil {
		return nil
	}
	out := new(MetadataOptions)
	in.DeepCopyInto(out)
	return out
}
