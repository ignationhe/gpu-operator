/*
Copyright 2024 NVIDIA CORPORATION & AFFILIATES.

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

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/scheme"
)

// State represents the state of a component managed by the GPU Operator.
type State string

const (
	// StateReady indicates the component is fully operational.
	StateReady State = "ready"
	// StateNotReady indicates the component is not yet operational.
	StateNotReady State = "notReady"
	// StateDisabled indicates the component is explicitly disabled.
	StateDisabled State = "disabled"
	// StateError indicates the component encountered an error.
	StateError State = "error"
	// StateIgnored indicates the component state is not being tracked.
	StateIgnored State = "ignored"
)

// GPUClusterPolicySpec defines the desired state of GPUClusterPolicy.
type GPUClusterPolicySpec struct {
	// Operator contains configuration for the GPU Operator itself.
	// +optional
	Operator OperatorSpec `json:"operator,omitempty"`

	// Driver contains configuration for the NVIDIA driver daemonset.
	// +optional
	Driver DriverSpec `json:"driver,omitempty"`

	// Toolkit contains configuration for the NVIDIA container toolkit daemonset.
	// +optional
	Toolkit ToolkitSpec `json:"toolkit,omitempty"`

	// DevicePlugin contains configuration for the NVIDIA device plugin daemonset.
	// +optional
	DevicePlugin DevicePluginSpec `json:"devicePlugin,omitempty"`

	// DCGMExporter contains configuration for the DCGM exporter daemonset.
	// +optional
	DCGMExporter DCGMExporterSpec `json:"dcgmExporter,omitempty"`
}

// OperatorSpec describes configuration options for the operator itself.
type OperatorSpec struct {
	// DefaultRuntime is the default container runtime on the cluster nodes.
	// +kubebuilder:validation:Enum;c;containerd
	// +kubebuilder:default=containerd
	DefaultRuntime string `json:"defaultRuntime,omitempty"`

	// InitContainerImage is the image used for init containers.
	// +optional
	InitContainerImage string `json:"initContainerImage,omitempty"`
}

// DriverSpec describes configuration for the NVIDIA driver component{
	// Enabled controls whether the driver daemonset is deployed.
	bebuilder:default=true
	Enabled *bool `json:"enabled,omitempty"`

	// Image is the NVIDIA driver container image.
	Image string `json:"image,omitempty"`

	// Version is the NVIDIA driver version to deploy.
	Version string `json:"version,omitempty"`
}

// ToolkitSpec describes configuration for the NVIDIA container toolkit component.
type ToolkitSpec struct {
	// Enabled controls whether the toolkit daemonset is deployed.
	// +kubebuilder:default=true
	Enabled *bool `json:"enabled,omitempty"`

	// Image is the container toolkit image.
	Image string `json:"image,omitempty"`

	// Version is the container toolkit version to deploy.
	Version string `json:"version,omitempty"`
}

// DevicePluginSpec describes configuration for the NVIDIA device plugin component.
type DevicePluginSpec struct {
	// Enabled controls whether the device plugin daemonset is deployed.
	// +kubebuilder:default=true
	Enabled *bool `json:"enabled,omitempty"`

	// Image is the device plugin container image.
	Image string `json:"image,omitempty"`

	// Version is the device plugin version to deploy.
	Version string `json:"version,omitempty"`
}

// DCGMExporterSpec describes configuration for the DCGM exporter component.
type DCGMExporterSpec struct {
	// Enabled controls whether the DCGM exporter daemonset is deployed.
	// +kubebuilder:default=true
	Enabled *bool `json:"enabled,omitempty"`

	// Image is the DCGM exporter container image.
	Image string `json:"image,omitempty"`

	// Version is the DCGM exporter version to deploy.
	Version string `json:"version,omitempty"`
}

// GPUClusterPolicyStatus defines the observed state of GPUClusterPolicy.
type GPUClusterPolicyStatus struct {
	// State is the overall state of the GPU cluster policy.
	State State `json:"state,omitempty"`

	// Namespace is the namespace where GPU operator components are deployed.
	Namespace string `json:"namespace,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
// +kubebuilder:printcolumn:name="Status",type=string,JSONPath=`.status.state`,priority=0
// +kubebuilder:printcolumn:name="Age",type=string,JSONPath=`.metadata.creationTimestamp`,priority=0

// GPUClusterPolicy is the Schema for the gpuclusterpolicies API.
type GPUClusterPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   GPUClusterPolicySpec   `json:"spec,omitempty"`
	Status GPUClusterPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// GPUClusterPolicyList contains a list of GPUClusterPolicy.
type GPUClusterPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GPUClusterPolicy `json:"items"`
}

var (
	// SchemeBuilder is used to add functions to this group's scheme.
	SchemeBuilder = &scheme.Builder{GroupVersion: GroupVersion}
	// AddToScheme adds the types in this group-version to the given scheme.
	AddToScheme = SchemeBuilder.AddToScheme
)

func init() {
	SchemeBuilder.Register(&GPUClusterPolicy{}, &GPUClusterPolicyList{})
}

// DeepCopyObject implements the runtime.Object interface.
func (in *GPUClusterPolicy) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopy creates a deep copy of GPUClusterPolicy.
func (in *GPUClusterPolicy) DeepCopy() *GPUClusterPolicy {
	if in == nil {
		return nil
	}
	out := new(GPUClusterPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto copies all properties of GPUClusterPolicy into another instance.
func (in *GPUClusterPolicy) DeepCopyInto(out *GPUClusterPolicy) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
}
