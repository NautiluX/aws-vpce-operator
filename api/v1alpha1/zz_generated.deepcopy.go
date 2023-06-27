//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AAOAccountAcceptanceCriteria) DeepCopyInto(out *AAOAccountAcceptanceCriteria) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AAOAccountAcceptanceCriteria.
func (in *AAOAccountAcceptanceCriteria) DeepCopy() *AAOAccountAcceptanceCriteria {
	if in == nil {
		return nil
	}
	out := new(AAOAccountAcceptanceCriteria)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AcceptanceCriteria) DeepCopyInto(out *AcceptanceCriteria) {
	*out = *in
	if in.AwsAccountOperatorAccount != nil {
		in, out := &in.AwsAccountOperatorAccount, &out.AwsAccountOperatorAccount
		*out = new(AAOAccountAcceptanceCriteria)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AcceptanceCriteria.
func (in *AcceptanceCriteria) DeepCopy() *AcceptanceCriteria {
	if in == nil {
		return nil
	}
	out := new(AcceptanceCriteria)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AvoConfig) DeepCopyInto(out *AvoConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ControllerManagerConfigurationSpec.DeepCopyInto(&out.ControllerManagerConfigurationSpec)
	if in.EnableVpcEndpointController != nil {
		in, out := &in.EnableVpcEndpointController, &out.EnableVpcEndpointController
		*out = new(bool)
		**out = **in
	}
	if in.EnableVpcEndpointAcceptanceController != nil {
		in, out := &in.EnableVpcEndpointAcceptanceController, &out.EnableVpcEndpointAcceptanceController
		*out = new(bool)
		**out = **in
	}
	if in.EnableVpcEndpointTemplateController != nil {
		in, out := &in.EnableVpcEndpointTemplateController, &out.EnableVpcEndpointTemplateController
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AvoConfig.
func (in *AvoConfig) DeepCopy() *AvoConfig {
	if in == nil {
		return nil
	}
	out := new(AvoConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AvoConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AvoConfigList) DeepCopyInto(out *AvoConfigList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AvoConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AvoConfigList.
func (in *AvoConfigList) DeepCopy() *AvoConfigList {
	if in == nil {
		return nil
	}
	out := new(AvoConfigList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AvoConfigList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExternalNameServiceSpec) DeepCopyInto(out *ExternalNameServiceSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExternalNameServiceSpec.
func (in *ExternalNameServiceSpec) DeepCopy() *ExternalNameServiceSpec {
	if in == nil {
		return nil
	}
	out := new(ExternalNameServiceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecurityGroup) DeepCopyInto(out *SecurityGroup) {
	*out = *in
	if in.IngressRules != nil {
		in, out := &in.IngressRules, &out.IngressRules
		*out = make([]SecurityGroupRule, len(*in))
		copy(*out, *in)
	}
	if in.EgressRules != nil {
		in, out := &in.EgressRules, &out.EgressRules
		*out = make([]SecurityGroupRule, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecurityGroup.
func (in *SecurityGroup) DeepCopy() *SecurityGroup {
	if in == nil {
		return nil
	}
	out := new(SecurityGroup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecurityGroupRule) DeepCopyInto(out *SecurityGroupRule) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecurityGroupRule.
func (in *SecurityGroupRule) DeepCopy() *SecurityGroupRule {
	if in == nil {
		return nil
	}
	out := new(SecurityGroupRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VpcEndpoint) DeepCopyInto(out *VpcEndpoint) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VpcEndpoint.
func (in *VpcEndpoint) DeepCopy() *VpcEndpoint {
	if in == nil {
		return nil
	}
	out := new(VpcEndpoint)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VpcEndpoint) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VpcEndpointAcceptance) DeepCopyInto(out *VpcEndpointAcceptance) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VpcEndpointAcceptance.
func (in *VpcEndpointAcceptance) DeepCopy() *VpcEndpointAcceptance {
	if in == nil {
		return nil
	}
	out := new(VpcEndpointAcceptance)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VpcEndpointAcceptance) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VpcEndpointAcceptanceList) DeepCopyInto(out *VpcEndpointAcceptanceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]VpcEndpointAcceptance, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VpcEndpointAcceptanceList.
func (in *VpcEndpointAcceptanceList) DeepCopy() *VpcEndpointAcceptanceList {
	if in == nil {
		return nil
	}
	out := new(VpcEndpointAcceptanceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VpcEndpointAcceptanceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VpcEndpointAcceptanceSpec) DeepCopyInto(out *VpcEndpointAcceptanceSpec) {
	*out = *in
	in.AcceptanceCriteria.DeepCopyInto(&out.AcceptanceCriteria)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VpcEndpointAcceptanceSpec.
func (in *VpcEndpointAcceptanceSpec) DeepCopy() *VpcEndpointAcceptanceSpec {
	if in == nil {
		return nil
	}
	out := new(VpcEndpointAcceptanceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VpcEndpointAcceptanceStatus) DeepCopyInto(out *VpcEndpointAcceptanceStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VpcEndpointAcceptanceStatus.
func (in *VpcEndpointAcceptanceStatus) DeepCopy() *VpcEndpointAcceptanceStatus {
	if in == nil {
		return nil
	}
	out := new(VpcEndpointAcceptanceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VpcEndpointList) DeepCopyInto(out *VpcEndpointList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]VpcEndpoint, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VpcEndpointList.
func (in *VpcEndpointList) DeepCopy() *VpcEndpointList {
	if in == nil {
		return nil
	}
	out := new(VpcEndpointList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VpcEndpointList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VpcEndpointSpec) DeepCopyInto(out *VpcEndpointSpec) {
	*out = *in
	in.SecurityGroup.DeepCopyInto(&out.SecurityGroup)
	out.ExternalNameService = in.ExternalNameService
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VpcEndpointSpec.
func (in *VpcEndpointSpec) DeepCopy() *VpcEndpointSpec {
	if in == nil {
		return nil
	}
	out := new(VpcEndpointSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VpcEndpointStatus) DeepCopyInto(out *VpcEndpointStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VpcEndpointStatus.
func (in *VpcEndpointStatus) DeepCopy() *VpcEndpointStatus {
	if in == nil {
		return nil
	}
	out := new(VpcEndpointStatus)
	in.DeepCopyInto(out)
	return out
}
