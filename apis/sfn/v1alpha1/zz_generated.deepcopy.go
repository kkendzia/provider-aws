//go:build !ignore_autogenerated

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
func (in *Activity) DeepCopyInto(out *Activity) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Activity.
func (in *Activity) DeepCopy() *Activity {
	if in == nil {
		return nil
	}
	out := new(Activity)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Activity) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ActivityList) DeepCopyInto(out *ActivityList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Activity, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ActivityList.
func (in *ActivityList) DeepCopy() *ActivityList {
	if in == nil {
		return nil
	}
	out := new(ActivityList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ActivityList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ActivityListItem) DeepCopyInto(out *ActivityListItem) {
	*out = *in
	if in.ActivityARN != nil {
		in, out := &in.ActivityARN, &out.ActivityARN
		*out = new(string)
		**out = **in
	}
	if in.CreationDate != nil {
		in, out := &in.CreationDate, &out.CreationDate
		*out = (*in).DeepCopy()
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ActivityListItem.
func (in *ActivityListItem) DeepCopy() *ActivityListItem {
	if in == nil {
		return nil
	}
	out := new(ActivityListItem)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ActivityObservation) DeepCopyInto(out *ActivityObservation) {
	*out = *in
	if in.ActivityARN != nil {
		in, out := &in.ActivityARN, &out.ActivityARN
		*out = new(string)
		**out = **in
	}
	if in.CreationDate != nil {
		in, out := &in.CreationDate, &out.CreationDate
		*out = (*in).DeepCopy()
	}
	out.CustomActivityObservation = in.CustomActivityObservation
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ActivityObservation.
func (in *ActivityObservation) DeepCopy() *ActivityObservation {
	if in == nil {
		return nil
	}
	out := new(ActivityObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ActivityParameters) DeepCopyInto(out *ActivityParameters) {
	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]*Tag, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(Tag)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	out.CustomActivityParameters = in.CustomActivityParameters
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ActivityParameters.
func (in *ActivityParameters) DeepCopy() *ActivityParameters {
	if in == nil {
		return nil
	}
	out := new(ActivityParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ActivityScheduledEventDetails) DeepCopyInto(out *ActivityScheduledEventDetails) {
	*out = *in
	if in.Resource != nil {
		in, out := &in.Resource, &out.Resource
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ActivityScheduledEventDetails.
func (in *ActivityScheduledEventDetails) DeepCopy() *ActivityScheduledEventDetails {
	if in == nil {
		return nil
	}
	out := new(ActivityScheduledEventDetails)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ActivitySpec) DeepCopyInto(out *ActivitySpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ActivitySpec.
func (in *ActivitySpec) DeepCopy() *ActivitySpec {
	if in == nil {
		return nil
	}
	out := new(ActivitySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ActivityStatus) DeepCopyInto(out *ActivityStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ActivityStatus.
func (in *ActivityStatus) DeepCopy() *ActivityStatus {
	if in == nil {
		return nil
	}
	out := new(ActivityStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudWatchLogsLogGroup) DeepCopyInto(out *CloudWatchLogsLogGroup) {
	*out = *in
	if in.LogGroupARN != nil {
		in, out := &in.LogGroupARN, &out.LogGroupARN
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudWatchLogsLogGroup.
func (in *CloudWatchLogsLogGroup) DeepCopy() *CloudWatchLogsLogGroup {
	if in == nil {
		return nil
	}
	out := new(CloudWatchLogsLogGroup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomActivityObservation) DeepCopyInto(out *CustomActivityObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomActivityObservation.
func (in *CustomActivityObservation) DeepCopy() *CustomActivityObservation {
	if in == nil {
		return nil
	}
	out := new(CustomActivityObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomActivityParameters) DeepCopyInto(out *CustomActivityParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomActivityParameters.
func (in *CustomActivityParameters) DeepCopy() *CustomActivityParameters {
	if in == nil {
		return nil
	}
	out := new(CustomActivityParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomStateMachineObservation) DeepCopyInto(out *CustomStateMachineObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomStateMachineObservation.
func (in *CustomStateMachineObservation) DeepCopy() *CustomStateMachineObservation {
	if in == nil {
		return nil
	}
	out := new(CustomStateMachineObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomStateMachineParameters) DeepCopyInto(out *CustomStateMachineParameters) {
	*out = *in
	if in.RoleARN != nil {
		in, out := &in.RoleARN, &out.RoleARN
		*out = new(string)
		**out = **in
	}
	if in.RoleARNRef != nil {
		in, out := &in.RoleARNRef, &out.RoleARNRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.RoleARNSelector != nil {
		in, out := &in.RoleARNSelector, &out.RoleARNSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomStateMachineParameters.
func (in *CustomStateMachineParameters) DeepCopy() *CustomStateMachineParameters {
	if in == nil {
		return nil
	}
	out := new(CustomStateMachineParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExecutionListItem) DeepCopyInto(out *ExecutionListItem) {
	*out = *in
	if in.ExecutionARN != nil {
		in, out := &in.ExecutionARN, &out.ExecutionARN
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.RedriveDate != nil {
		in, out := &in.RedriveDate, &out.RedriveDate
		*out = (*in).DeepCopy()
	}
	if in.StartDate != nil {
		in, out := &in.StartDate, &out.StartDate
		*out = (*in).DeepCopy()
	}
	if in.StateMachineAliasARN != nil {
		in, out := &in.StateMachineAliasARN, &out.StateMachineAliasARN
		*out = new(string)
		**out = **in
	}
	if in.StateMachineARN != nil {
		in, out := &in.StateMachineARN, &out.StateMachineARN
		*out = new(string)
		**out = **in
	}
	if in.StateMachineVersionARN != nil {
		in, out := &in.StateMachineVersionARN, &out.StateMachineVersionARN
		*out = new(string)
		**out = **in
	}
	if in.StopDate != nil {
		in, out := &in.StopDate, &out.StopDate
		*out = (*in).DeepCopy()
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExecutionListItem.
func (in *ExecutionListItem) DeepCopy() *ExecutionListItem {
	if in == nil {
		return nil
	}
	out := new(ExecutionListItem)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExecutionStartedEventDetails) DeepCopyInto(out *ExecutionStartedEventDetails) {
	*out = *in
	if in.RoleARN != nil {
		in, out := &in.RoleARN, &out.RoleARN
		*out = new(string)
		**out = **in
	}
	if in.StateMachineAliasARN != nil {
		in, out := &in.StateMachineAliasARN, &out.StateMachineAliasARN
		*out = new(string)
		**out = **in
	}
	if in.StateMachineVersionARN != nil {
		in, out := &in.StateMachineVersionARN, &out.StateMachineVersionARN
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExecutionStartedEventDetails.
func (in *ExecutionStartedEventDetails) DeepCopy() *ExecutionStartedEventDetails {
	if in == nil {
		return nil
	}
	out := new(ExecutionStartedEventDetails)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HistoryEvent) DeepCopyInto(out *HistoryEvent) {
	*out = *in
	if in.Timestamp != nil {
		in, out := &in.Timestamp, &out.Timestamp
		*out = (*in).DeepCopy()
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HistoryEvent.
func (in *HistoryEvent) DeepCopy() *HistoryEvent {
	if in == nil {
		return nil
	}
	out := new(HistoryEvent)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LambdaFunctionScheduledEventDetails) DeepCopyInto(out *LambdaFunctionScheduledEventDetails) {
	*out = *in
	if in.Resource != nil {
		in, out := &in.Resource, &out.Resource
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LambdaFunctionScheduledEventDetails.
func (in *LambdaFunctionScheduledEventDetails) DeepCopy() *LambdaFunctionScheduledEventDetails {
	if in == nil {
		return nil
	}
	out := new(LambdaFunctionScheduledEventDetails)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LogDestination) DeepCopyInto(out *LogDestination) {
	*out = *in
	if in.CloudWatchLogsLogGroup != nil {
		in, out := &in.CloudWatchLogsLogGroup, &out.CloudWatchLogsLogGroup
		*out = new(CloudWatchLogsLogGroup)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LogDestination.
func (in *LogDestination) DeepCopy() *LogDestination {
	if in == nil {
		return nil
	}
	out := new(LogDestination)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoggingConfiguration) DeepCopyInto(out *LoggingConfiguration) {
	*out = *in
	if in.Destinations != nil {
		in, out := &in.Destinations, &out.Destinations
		*out = make([]*LogDestination, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(LogDestination)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.IncludeExecutionData != nil {
		in, out := &in.IncludeExecutionData, &out.IncludeExecutionData
		*out = new(bool)
		**out = **in
	}
	if in.Level != nil {
		in, out := &in.Level, &out.Level
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoggingConfiguration.
func (in *LoggingConfiguration) DeepCopy() *LoggingConfiguration {
	if in == nil {
		return nil
	}
	out := new(LoggingConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MapIterationEventDetails) DeepCopyInto(out *MapIterationEventDetails) {
	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MapIterationEventDetails.
func (in *MapIterationEventDetails) DeepCopy() *MapIterationEventDetails {
	if in == nil {
		return nil
	}
	out := new(MapIterationEventDetails)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MapRunListItem) DeepCopyInto(out *MapRunListItem) {
	*out = *in
	if in.ExecutionARN != nil {
		in, out := &in.ExecutionARN, &out.ExecutionARN
		*out = new(string)
		**out = **in
	}
	if in.StartDate != nil {
		in, out := &in.StartDate, &out.StartDate
		*out = (*in).DeepCopy()
	}
	if in.StateMachineARN != nil {
		in, out := &in.StateMachineARN, &out.StateMachineARN
		*out = new(string)
		**out = **in
	}
	if in.StopDate != nil {
		in, out := &in.StopDate, &out.StopDate
		*out = (*in).DeepCopy()
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MapRunListItem.
func (in *MapRunListItem) DeepCopy() *MapRunListItem {
	if in == nil {
		return nil
	}
	out := new(MapRunListItem)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RoutingConfigurationListItem) DeepCopyInto(out *RoutingConfigurationListItem) {
	*out = *in
	if in.StateMachineVersionARN != nil {
		in, out := &in.StateMachineVersionARN, &out.StateMachineVersionARN
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RoutingConfigurationListItem.
func (in *RoutingConfigurationListItem) DeepCopy() *RoutingConfigurationListItem {
	if in == nil {
		return nil
	}
	out := new(RoutingConfigurationListItem)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StateEnteredEventDetails) DeepCopyInto(out *StateEnteredEventDetails) {
	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StateEnteredEventDetails.
func (in *StateEnteredEventDetails) DeepCopy() *StateEnteredEventDetails {
	if in == nil {
		return nil
	}
	out := new(StateEnteredEventDetails)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StateExitedEventDetails) DeepCopyInto(out *StateExitedEventDetails) {
	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StateExitedEventDetails.
func (in *StateExitedEventDetails) DeepCopy() *StateExitedEventDetails {
	if in == nil {
		return nil
	}
	out := new(StateExitedEventDetails)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StateMachine) DeepCopyInto(out *StateMachine) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StateMachine.
func (in *StateMachine) DeepCopy() *StateMachine {
	if in == nil {
		return nil
	}
	out := new(StateMachine)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *StateMachine) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StateMachineAliasListItem) DeepCopyInto(out *StateMachineAliasListItem) {
	*out = *in
	if in.CreationDate != nil {
		in, out := &in.CreationDate, &out.CreationDate
		*out = (*in).DeepCopy()
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StateMachineAliasListItem.
func (in *StateMachineAliasListItem) DeepCopy() *StateMachineAliasListItem {
	if in == nil {
		return nil
	}
	out := new(StateMachineAliasListItem)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StateMachineList) DeepCopyInto(out *StateMachineList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]StateMachine, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StateMachineList.
func (in *StateMachineList) DeepCopy() *StateMachineList {
	if in == nil {
		return nil
	}
	out := new(StateMachineList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *StateMachineList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StateMachineListItem) DeepCopyInto(out *StateMachineListItem) {
	*out = *in
	if in.CreationDate != nil {
		in, out := &in.CreationDate, &out.CreationDate
		*out = (*in).DeepCopy()
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.StateMachineARN != nil {
		in, out := &in.StateMachineARN, &out.StateMachineARN
		*out = new(string)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StateMachineListItem.
func (in *StateMachineListItem) DeepCopy() *StateMachineListItem {
	if in == nil {
		return nil
	}
	out := new(StateMachineListItem)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StateMachineObservation) DeepCopyInto(out *StateMachineObservation) {
	*out = *in
	if in.CreationDate != nil {
		in, out := &in.CreationDate, &out.CreationDate
		*out = (*in).DeepCopy()
	}
	if in.StateMachineARN != nil {
		in, out := &in.StateMachineARN, &out.StateMachineARN
		*out = new(string)
		**out = **in
	}
	if in.StateMachineVersionARN != nil {
		in, out := &in.StateMachineVersionARN, &out.StateMachineVersionARN
		*out = new(string)
		**out = **in
	}
	out.CustomStateMachineObservation = in.CustomStateMachineObservation
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StateMachineObservation.
func (in *StateMachineObservation) DeepCopy() *StateMachineObservation {
	if in == nil {
		return nil
	}
	out := new(StateMachineObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StateMachineParameters) DeepCopyInto(out *StateMachineParameters) {
	*out = *in
	if in.Definition != nil {
		in, out := &in.Definition, &out.Definition
		*out = new(string)
		**out = **in
	}
	if in.LoggingConfiguration != nil {
		in, out := &in.LoggingConfiguration, &out.LoggingConfiguration
		*out = new(LoggingConfiguration)
		(*in).DeepCopyInto(*out)
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Publish != nil {
		in, out := &in.Publish, &out.Publish
		*out = new(bool)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]*Tag, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(Tag)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.TracingConfiguration != nil {
		in, out := &in.TracingConfiguration, &out.TracingConfiguration
		*out = new(TracingConfiguration)
		(*in).DeepCopyInto(*out)
	}
	if in.VersionDescription != nil {
		in, out := &in.VersionDescription, &out.VersionDescription
		*out = new(string)
		**out = **in
	}
	in.CustomStateMachineParameters.DeepCopyInto(&out.CustomStateMachineParameters)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StateMachineParameters.
func (in *StateMachineParameters) DeepCopy() *StateMachineParameters {
	if in == nil {
		return nil
	}
	out := new(StateMachineParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StateMachineSpec) DeepCopyInto(out *StateMachineSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StateMachineSpec.
func (in *StateMachineSpec) DeepCopy() *StateMachineSpec {
	if in == nil {
		return nil
	}
	out := new(StateMachineSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StateMachineStatus) DeepCopyInto(out *StateMachineStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StateMachineStatus.
func (in *StateMachineStatus) DeepCopy() *StateMachineStatus {
	if in == nil {
		return nil
	}
	out := new(StateMachineStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StateMachineVersionListItem) DeepCopyInto(out *StateMachineVersionListItem) {
	*out = *in
	if in.CreationDate != nil {
		in, out := &in.CreationDate, &out.CreationDate
		*out = (*in).DeepCopy()
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StateMachineVersionListItem.
func (in *StateMachineVersionListItem) DeepCopy() *StateMachineVersionListItem {
	if in == nil {
		return nil
	}
	out := new(StateMachineVersionListItem)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Tag) DeepCopyInto(out *Tag) {
	*out = *in
	if in.Key != nil {
		in, out := &in.Key, &out.Key
		*out = new(string)
		**out = **in
	}
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Tag.
func (in *Tag) DeepCopy() *Tag {
	if in == nil {
		return nil
	}
	out := new(Tag)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TaskFailedEventDetails) DeepCopyInto(out *TaskFailedEventDetails) {
	*out = *in
	if in.Resource != nil {
		in, out := &in.Resource, &out.Resource
		*out = new(string)
		**out = **in
	}
	if in.ResourceType != nil {
		in, out := &in.ResourceType, &out.ResourceType
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TaskFailedEventDetails.
func (in *TaskFailedEventDetails) DeepCopy() *TaskFailedEventDetails {
	if in == nil {
		return nil
	}
	out := new(TaskFailedEventDetails)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TaskScheduledEventDetails) DeepCopyInto(out *TaskScheduledEventDetails) {
	*out = *in
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
	if in.Resource != nil {
		in, out := &in.Resource, &out.Resource
		*out = new(string)
		**out = **in
	}
	if in.ResourceType != nil {
		in, out := &in.ResourceType, &out.ResourceType
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TaskScheduledEventDetails.
func (in *TaskScheduledEventDetails) DeepCopy() *TaskScheduledEventDetails {
	if in == nil {
		return nil
	}
	out := new(TaskScheduledEventDetails)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TaskStartFailedEventDetails) DeepCopyInto(out *TaskStartFailedEventDetails) {
	*out = *in
	if in.Resource != nil {
		in, out := &in.Resource, &out.Resource
		*out = new(string)
		**out = **in
	}
	if in.ResourceType != nil {
		in, out := &in.ResourceType, &out.ResourceType
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TaskStartFailedEventDetails.
func (in *TaskStartFailedEventDetails) DeepCopy() *TaskStartFailedEventDetails {
	if in == nil {
		return nil
	}
	out := new(TaskStartFailedEventDetails)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TaskStartedEventDetails) DeepCopyInto(out *TaskStartedEventDetails) {
	*out = *in
	if in.Resource != nil {
		in, out := &in.Resource, &out.Resource
		*out = new(string)
		**out = **in
	}
	if in.ResourceType != nil {
		in, out := &in.ResourceType, &out.ResourceType
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TaskStartedEventDetails.
func (in *TaskStartedEventDetails) DeepCopy() *TaskStartedEventDetails {
	if in == nil {
		return nil
	}
	out := new(TaskStartedEventDetails)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TaskSubmitFailedEventDetails) DeepCopyInto(out *TaskSubmitFailedEventDetails) {
	*out = *in
	if in.Resource != nil {
		in, out := &in.Resource, &out.Resource
		*out = new(string)
		**out = **in
	}
	if in.ResourceType != nil {
		in, out := &in.ResourceType, &out.ResourceType
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TaskSubmitFailedEventDetails.
func (in *TaskSubmitFailedEventDetails) DeepCopy() *TaskSubmitFailedEventDetails {
	if in == nil {
		return nil
	}
	out := new(TaskSubmitFailedEventDetails)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TaskSubmittedEventDetails) DeepCopyInto(out *TaskSubmittedEventDetails) {
	*out = *in
	if in.Resource != nil {
		in, out := &in.Resource, &out.Resource
		*out = new(string)
		**out = **in
	}
	if in.ResourceType != nil {
		in, out := &in.ResourceType, &out.ResourceType
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TaskSubmittedEventDetails.
func (in *TaskSubmittedEventDetails) DeepCopy() *TaskSubmittedEventDetails {
	if in == nil {
		return nil
	}
	out := new(TaskSubmittedEventDetails)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TaskSucceededEventDetails) DeepCopyInto(out *TaskSucceededEventDetails) {
	*out = *in
	if in.Resource != nil {
		in, out := &in.Resource, &out.Resource
		*out = new(string)
		**out = **in
	}
	if in.ResourceType != nil {
		in, out := &in.ResourceType, &out.ResourceType
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TaskSucceededEventDetails.
func (in *TaskSucceededEventDetails) DeepCopy() *TaskSucceededEventDetails {
	if in == nil {
		return nil
	}
	out := new(TaskSucceededEventDetails)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TaskTimedOutEventDetails) DeepCopyInto(out *TaskTimedOutEventDetails) {
	*out = *in
	if in.Resource != nil {
		in, out := &in.Resource, &out.Resource
		*out = new(string)
		**out = **in
	}
	if in.ResourceType != nil {
		in, out := &in.ResourceType, &out.ResourceType
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TaskTimedOutEventDetails.
func (in *TaskTimedOutEventDetails) DeepCopy() *TaskTimedOutEventDetails {
	if in == nil {
		return nil
	}
	out := new(TaskTimedOutEventDetails)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TracingConfiguration) DeepCopyInto(out *TracingConfiguration) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TracingConfiguration.
func (in *TracingConfiguration) DeepCopy() *TracingConfiguration {
	if in == nil {
		return nil
	}
	out := new(TracingConfiguration)
	in.DeepCopyInto(out)
	return out
}
