//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2023.

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

package v1beta1

import (
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BestVersion) DeepCopyInto(out *BestVersion) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BestVersion.
func (in *BestVersion) DeepCopy() *BestVersion {
	if in == nil {
		return nil
	}
	out := new(BestVersion)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FineTune) DeepCopyInto(out *FineTune) {
	*out = *in
	if in.Resource != nil {
		in, out := &in.Resource, &out.Resource
		*out = new(Resource)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FineTune.
func (in *FineTune) DeepCopy() *FineTune {
	if in == nil {
		return nil
	}
	out := new(FineTune)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FinetuneExperiment) DeepCopyInto(out *FinetuneExperiment) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FinetuneExperiment.
func (in *FinetuneExperiment) DeepCopy() *FinetuneExperiment {
	if in == nil {
		return nil
	}
	out := new(FinetuneExperiment)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FinetuneExperiment) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FinetuneExperimentList) DeepCopyInto(out *FinetuneExperimentList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]FinetuneExperiment, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FinetuneExperimentList.
func (in *FinetuneExperimentList) DeepCopy() *FinetuneExperimentList {
	if in == nil {
		return nil
	}
	out := new(FinetuneExperimentList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FinetuneExperimentList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FinetuneExperimentSpec) DeepCopyInto(out *FinetuneExperimentSpec) {
	*out = *in
	if in.FinetuneJobs != nil {
		in, out := &in.FinetuneJobs, &out.FinetuneJobs
		*out = make([]FinetuneJob, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	out.ScoringConfig = in.ScoringConfig
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FinetuneExperimentSpec.
func (in *FinetuneExperimentSpec) DeepCopy() *FinetuneExperimentSpec {
	if in == nil {
		return nil
	}
	out := new(FinetuneExperimentSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FinetuneExperimentStatus) DeepCopyInto(out *FinetuneExperimentStatus) {
	*out = *in
	out.BestVersion = in.BestVersion
	if in.JobStates != nil {
		in, out := &in.JobStates, &out.JobStates
		*out = make([]FinetuneJobStatus, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FinetuneExperimentStatus.
func (in *FinetuneExperimentStatus) DeepCopy() *FinetuneExperimentStatus {
	if in == nil {
		return nil
	}
	out := new(FinetuneExperimentStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FinetuneJob) DeepCopyInto(out *FinetuneJob) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FinetuneJob.
func (in *FinetuneJob) DeepCopy() *FinetuneJob {
	if in == nil {
		return nil
	}
	out := new(FinetuneJob)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FinetuneJob) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FinetuneJobList) DeepCopyInto(out *FinetuneJobList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]FinetuneJob, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FinetuneJobList.
func (in *FinetuneJobList) DeepCopy() *FinetuneJobList {
	if in == nil {
		return nil
	}
	out := new(FinetuneJobList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FinetuneJobList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FinetuneJobResult) DeepCopyInto(out *FinetuneJobResult) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FinetuneJobResult.
func (in *FinetuneJobResult) DeepCopy() *FinetuneJobResult {
	if in == nil {
		return nil
	}
	out := new(FinetuneJobResult)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FinetuneJobSpec) DeepCopyInto(out *FinetuneJobSpec) {
	*out = *in
	in.FineTune.DeepCopyInto(&out.FineTune)
	if in.ScoringConfig != nil {
		in, out := &in.ScoringConfig, &out.ScoringConfig
		*out = new(ScoringConfig)
		**out = **in
	}
	out.ServeConfig = in.ServeConfig
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FinetuneJobSpec.
func (in *FinetuneJobSpec) DeepCopy() *FinetuneJobSpec {
	if in == nil {
		return nil
	}
	out := new(FinetuneJobSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FinetuneJobStatus) DeepCopyInto(out *FinetuneJobStatus) {
	*out = *in
	out.Result = in.Result
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FinetuneJobStatus.
func (in *FinetuneJobStatus) DeepCopy() *FinetuneJobStatus {
	if in == nil {
		return nil
	}
	out := new(FinetuneJobStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Resource) DeepCopyInto(out *Resource) {
	*out = *in
	if in.Limits != nil {
		in, out := &in.Limits, &out.Limits
		*out = new(ResourceLimits)
		(*in).DeepCopyInto(*out)
	}
	if in.Requests != nil {
		in, out := &in.Requests, &out.Requests
		*out = new(ResourceLimits)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Resource.
func (in *Resource) DeepCopy() *Resource {
	if in == nil {
		return nil
	}
	out := new(Resource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceLimits) DeepCopyInto(out *ResourceLimits) {
	*out = *in
	out.CPU = in.CPU.DeepCopy()
	out.Memory = in.Memory.DeepCopy()
	if in.GPU != nil {
		in, out := &in.GPU, &out.GPU
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceLimits.
func (in *ResourceLimits) DeepCopy() *ResourceLimits {
	if in == nil {
		return nil
	}
	out := new(ResourceLimits)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScoringConfig) DeepCopyInto(out *ScoringConfig) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScoringConfig.
func (in *ScoringConfig) DeepCopy() *ScoringConfig {
	if in == nil {
		return nil
	}
	out := new(ScoringConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServeConfig) DeepCopyInto(out *ServeConfig) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServeConfig.
func (in *ServeConfig) DeepCopy() *ServeConfig {
	if in == nil {
		return nil
	}
	out := new(ServeConfig)
	in.DeepCopyInto(out)
	return out
}
