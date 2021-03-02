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

package v1

import (
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Condition) DeepCopyInto(out *Condition) {
	*out = *in
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Condition.
func (in *Condition) DeepCopy() *Condition {
	if in == nil {
		return nil
	}
	out := new(Condition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConditionsMeta) DeepCopyInto(out *ConditionsMeta) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConditionsMeta.
func (in *ConditionsMeta) DeepCopy() *ConditionsMeta {
	if in == nil {
		return nil
	}
	out := new(ConditionsMeta)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KogitoCloudEventInfo) DeepCopyInto(out *KogitoCloudEventInfo) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KogitoCloudEventInfo.
func (in *KogitoCloudEventInfo) DeepCopy() *KogitoCloudEventInfo {
	if in == nil {
		return nil
	}
	out := new(KogitoCloudEventInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KogitoCloudEventsStatus) DeepCopyInto(out *KogitoCloudEventsStatus) {
	*out = *in
	if in.Consumes != nil {
		in, out := &in.Consumes, &out.Consumes
		*out = make([]KogitoCloudEventInfo, len(*in))
		copy(*out, *in)
	}
	if in.Produces != nil {
		in, out := &in.Produces, &out.Produces
		*out = make([]KogitoCloudEventInfo, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KogitoCloudEventsStatus.
func (in *KogitoCloudEventsStatus) DeepCopy() *KogitoCloudEventsStatus {
	if in == nil {
		return nil
	}
	out := new(KogitoCloudEventsStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KogitoProbe) DeepCopyInto(out *KogitoProbe) {
	*out = *in
	in.LivenessProbe.DeepCopyInto(&out.LivenessProbe)
	in.ReadinessProbe.DeepCopyInto(&out.ReadinessProbe)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KogitoProbe.
func (in *KogitoProbe) DeepCopy() *KogitoProbe {
	if in == nil {
		return nil
	}
	out := new(KogitoProbe)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KogitoRuntime) DeepCopyInto(out *KogitoRuntime) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KogitoRuntime.
func (in *KogitoRuntime) DeepCopy() *KogitoRuntime {
	if in == nil {
		return nil
	}
	out := new(KogitoRuntime)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KogitoRuntime) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KogitoRuntimeList) DeepCopyInto(out *KogitoRuntimeList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]KogitoRuntime, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KogitoRuntimeList.
func (in *KogitoRuntimeList) DeepCopy() *KogitoRuntimeList {
	if in == nil {
		return nil
	}
	out := new(KogitoRuntimeList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KogitoRuntimeList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KogitoRuntimeSpec) DeepCopyInto(out *KogitoRuntimeSpec) {
	*out = *in
	in.KogitoServiceSpec.DeepCopyInto(&out.KogitoServiceSpec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KogitoRuntimeSpec.
func (in *KogitoRuntimeSpec) DeepCopy() *KogitoRuntimeSpec {
	if in == nil {
		return nil
	}
	out := new(KogitoRuntimeSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KogitoRuntimeStatus) DeepCopyInto(out *KogitoRuntimeStatus) {
	*out = *in
	in.KogitoServiceStatus.DeepCopyInto(&out.KogitoServiceStatus)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KogitoRuntimeStatus.
func (in *KogitoRuntimeStatus) DeepCopy() *KogitoRuntimeStatus {
	if in == nil {
		return nil
	}
	out := new(KogitoRuntimeStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KogitoServiceSpec) DeepCopyInto(out *KogitoServiceSpec) {
	*out = *in
	if in.Replicas != nil {
		in, out := &in.Replicas, &out.Replicas
		*out = new(int32)
		**out = **in
	}
	if in.Env != nil {
		in, out := &in.Env, &out.Env
		*out = make([]corev1.EnvVar, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.Resources.DeepCopyInto(&out.Resources)
	if in.DeploymentLabels != nil {
		in, out := &in.DeploymentLabels, &out.DeploymentLabels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.ServiceLabels != nil {
		in, out := &in.ServiceLabels, &out.ServiceLabels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	out.Monitoring = in.Monitoring
	if in.Config != nil {
		in, out := &in.Config, &out.Config
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	in.Probes.DeepCopyInto(&out.Probes)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KogitoServiceSpec.
func (in *KogitoServiceSpec) DeepCopy() *KogitoServiceSpec {
	if in == nil {
		return nil
	}
	out := new(KogitoServiceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KogitoServiceStatus) DeepCopyInto(out *KogitoServiceStatus) {
	*out = *in
	in.ConditionsMeta.DeepCopyInto(&out.ConditionsMeta)
	if in.DeploymentConditions != nil {
		in, out := &in.DeploymentConditions, &out.DeploymentConditions
		*out = make([]appsv1.DeploymentCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.CloudEvents.DeepCopyInto(&out.CloudEvents)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KogitoServiceStatus.
func (in *KogitoServiceStatus) DeepCopy() *KogitoServiceStatus {
	if in == nil {
		return nil
	}
	out := new(KogitoServiceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Monitoring) DeepCopyInto(out *Monitoring) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Monitoring.
func (in *Monitoring) DeepCopy() *Monitoring {
	if in == nil {
		return nil
	}
	out := new(Monitoring)
	in.DeepCopyInto(out)
	return out
}
