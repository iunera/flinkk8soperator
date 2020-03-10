// +build !ignore_autogenerated

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1beta2

import (
	v1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnvironmentConfig) DeepCopyInto(out *EnvironmentConfig) {
	*out = *in
	if in.EnvFrom != nil {
		in, out := &in.EnvFrom, &out.EnvFrom
		*out = make([]v1.EnvFromSource, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Env != nil {
		in, out := &in.Env, &out.Env
		*out = make([]v1.EnvVar, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnvironmentConfig.
func (in *EnvironmentConfig) DeepCopy() *EnvironmentConfig {
	if in == nil {
		return nil
	}
	out := new(EnvironmentConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FlinkApplication) DeepCopyInto(out *FlinkApplication) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FlinkApplication.
func (in *FlinkApplication) DeepCopy() *FlinkApplication {
	if in == nil {
		return nil
	}
	out := new(FlinkApplication)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FlinkApplication) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FlinkApplicationError) DeepCopyInto(out *FlinkApplicationError) {
	*out = *in
	if in.LastErrorUpdateTime != nil {
		in, out := &in.LastErrorUpdateTime, &out.LastErrorUpdateTime
		*out = (*in).DeepCopy()
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FlinkApplicationError.
func (in *FlinkApplicationError) DeepCopy() *FlinkApplicationError {
	if in == nil {
		return nil
	}
	out := new(FlinkApplicationError)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FlinkApplicationList) DeepCopyInto(out *FlinkApplicationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]FlinkApplication, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FlinkApplicationList.
func (in *FlinkApplicationList) DeepCopy() *FlinkApplicationList {
	if in == nil {
		return nil
	}
	out := new(FlinkApplicationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FlinkApplicationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FlinkApplicationSpec) DeepCopyInto(out *FlinkApplicationSpec) {
	*out = *in
	if in.ImagePullSecrets != nil {
		in, out := &in.ImagePullSecrets, &out.ImagePullSecrets
		*out = make([]v1.LocalObjectReference, len(*in))
		copy(*out, *in)
	}
	if in.SecurityContext != nil {
		in, out := &in.SecurityContext, &out.SecurityContext
		*out = new(v1.PodSecurityContext)
		(*in).DeepCopyInto(*out)
	}
	in.FlinkConfig.DeepCopyInto(&out.FlinkConfig)
	in.TaskManagerConfig.DeepCopyInto(&out.TaskManagerConfig)
	in.JobManagerConfig.DeepCopyInto(&out.JobManagerConfig)
	out.SavepointInfo = in.SavepointInfo
	if in.RPCPort != nil {
		in, out := &in.RPCPort, &out.RPCPort
		*out = new(int32)
		**out = **in
	}
	if in.BlobPort != nil {
		in, out := &in.BlobPort, &out.BlobPort
		*out = new(int32)
		**out = **in
	}
	if in.QueryPort != nil {
		in, out := &in.QueryPort, &out.QueryPort
		*out = new(int32)
		**out = **in
	}
	if in.UIPort != nil {
		in, out := &in.UIPort, &out.UIPort
		*out = new(int32)
		**out = **in
	}
	if in.MetricsQueryPort != nil {
		in, out := &in.MetricsQueryPort, &out.MetricsQueryPort
		*out = new(int32)
		**out = **in
	}
	if in.Volumes != nil {
		in, out := &in.Volumes, &out.Volumes
		*out = make([]v1.Volume, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.VolumeMounts != nil {
		in, out := &in.VolumeMounts, &out.VolumeMounts
		*out = make([]v1.VolumeMount, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.MaxCheckpointRestoreAgeSeconds != nil {
		in, out := &in.MaxCheckpointRestoreAgeSeconds, &out.MaxCheckpointRestoreAgeSeconds
		*out = new(int32)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FlinkApplicationSpec.
func (in *FlinkApplicationSpec) DeepCopy() *FlinkApplicationSpec {
	if in == nil {
		return nil
	}
	out := new(FlinkApplicationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FlinkApplicationStatus) DeepCopyInto(out *FlinkApplicationStatus) {
	*out = *in
	if in.StartedAt != nil {
		in, out := &in.StartedAt, &out.StartedAt
		*out = (*in).DeepCopy()
	}
	if in.LastUpdatedAt != nil {
		in, out := &in.LastUpdatedAt, &out.LastUpdatedAt
		*out = (*in).DeepCopy()
	}
	if in.VersionStatuses != nil {
		in, out := &in.VersionStatuses, &out.VersionStatuses
		*out = make([]FlinkApplicationVersionStatus, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.LastSeenError != nil {
		in, out := &in.LastSeenError, &out.LastSeenError
		*out = new(FlinkApplicationError)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FlinkApplicationStatus.
func (in *FlinkApplicationStatus) DeepCopy() *FlinkApplicationStatus {
	if in == nil {
		return nil
	}
	out := new(FlinkApplicationStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FlinkApplicationVersionStatus) DeepCopyInto(out *FlinkApplicationVersionStatus) {
	*out = *in
	out.ClusterStatus = in.ClusterStatus
	in.JobStatus.DeepCopyInto(&out.JobStatus)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FlinkApplicationVersionStatus.
func (in *FlinkApplicationVersionStatus) DeepCopy() *FlinkApplicationVersionStatus {
	if in == nil {
		return nil
	}
	out := new(FlinkApplicationVersionStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FlinkClusterStatus) DeepCopyInto(out *FlinkClusterStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FlinkClusterStatus.
func (in *FlinkClusterStatus) DeepCopy() *FlinkClusterStatus {
	if in == nil {
		return nil
	}
	out := new(FlinkClusterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FlinkJobStatus) DeepCopyInto(out *FlinkJobStatus) {
	*out = *in
	if in.StartTime != nil {
		in, out := &in.StartTime, &out.StartTime
		*out = (*in).DeepCopy()
	}
	if in.RestoreTime != nil {
		in, out := &in.RestoreTime, &out.RestoreTime
		*out = (*in).DeepCopy()
	}
	if in.LastFailingTime != nil {
		in, out := &in.LastFailingTime, &out.LastFailingTime
		*out = (*in).DeepCopy()
	}
	if in.LastCheckpointTime != nil {
		in, out := &in.LastCheckpointTime, &out.LastCheckpointTime
		*out = (*in).DeepCopy()
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FlinkJobStatus.
func (in *FlinkJobStatus) DeepCopy() *FlinkJobStatus {
	if in == nil {
		return nil
	}
	out := new(FlinkJobStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JobManagerConfig) DeepCopyInto(out *JobManagerConfig) {
	*out = *in
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = new(v1.ResourceRequirements)
		(*in).DeepCopyInto(*out)
	}
	in.EnvConfig.DeepCopyInto(&out.EnvConfig)
	if in.Replicas != nil {
		in, out := &in.Replicas, &out.Replicas
		*out = new(int32)
		**out = **in
	}
	if in.OffHeapMemoryFraction != nil {
		in, out := &in.OffHeapMemoryFraction, &out.OffHeapMemoryFraction
		*out = new(float64)
		**out = **in
	}
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Tolerations != nil {
		in, out := &in.Tolerations, &out.Tolerations
		*out = make([]v1.Toleration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JobManagerConfig.
func (in *JobManagerConfig) DeepCopy() *JobManagerConfig {
	if in == nil {
		return nil
	}
	out := new(JobManagerConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SavepointInfo) DeepCopyInto(out *SavepointInfo) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SavepointInfo.
func (in *SavepointInfo) DeepCopy() *SavepointInfo {
	if in == nil {
		return nil
	}
	out := new(SavepointInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TaskManagerConfig) DeepCopyInto(out *TaskManagerConfig) {
	*out = *in
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = new(v1.ResourceRequirements)
		(*in).DeepCopyInto(*out)
	}
	in.EnvConfig.DeepCopyInto(&out.EnvConfig)
	if in.TaskSlots != nil {
		in, out := &in.TaskSlots, &out.TaskSlots
		*out = new(int32)
		**out = **in
	}
	if in.OffHeapMemoryFraction != nil {
		in, out := &in.OffHeapMemoryFraction, &out.OffHeapMemoryFraction
		*out = new(float64)
		**out = **in
	}
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Tolerations != nil {
		in, out := &in.Tolerations, &out.Tolerations
		*out = make([]v1.Toleration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TaskManagerConfig.
func (in *TaskManagerConfig) DeepCopy() *TaskManagerConfig {
	if in == nil {
		return nil
	}
	out := new(TaskManagerConfig)
	in.DeepCopyInto(out)
	return out
}
