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
)

// NVIDIADriverSpec defines the desired state of NVIDIADriver
type NVIDIADriverSpec struct {
	// DriverType defines the type of NVIDIA driver to deploy.
	// +kubebuilder:validation:Enum=gpu;vgpu;vgpu-host-manager
	// +kubebuilder:default=gpu
	DriverType DriverType `json:"driverType,omitempty"`

	// Repository is the NVIDIA driver container image repository.
	// +optional
	Repository string `json:"repository,omitempty"`

	// Image is the NVIDIA driver container image name.
	// +optional
	Image string `json:"image,omitempty"`

	// Version is the NVIDIA driver version.
	// +optional
	Version string `json:"version,omitempty"`

	// NodeSelector specifies a selector for the nodes on which the driver should be deployed.
	// +optional
	NodeSelector map[string]string `json:"nodeSelector,omitempty"`

	// UsePrecompiled indicates whether to use precompiled driver modules.
	// +optional
	// +kubebuilder:default=false
	UsePrecompiled *bool `json:"usePrecompiled,omitempty"`
}

// DriverType defines the type of NVIDIA driver to be deployed
// +kubebuilder:validation:Enum=gpu;vgpu;vgpu-host-manager
type DriverType string

const (
	// GPU is the standard NVIDIA GPU driver type
	GPU DriverType = "gpu"
	// VGPU is the NVIDIA vGPU guest driver type
	VGPU DriverType = "vgpu"
	// VGPUHostManager is the NVIDIA vGPU host manager driver type
	VGPUHostManager DriverType = "vgpu-host-manager"
)

// NVIDIADriverStatus defines the observed state of NVIDIADriver
type NVIDIADriverStatus struct {
	// State indicates the current state of the NVIDIADriver deployment.
	// +optional
	State State `json:"state,omitempty"`

	// Namespace is the namespace where the NVIDIADriver DaemonSet is deployed.
	// +optional
	Namespace string `json:"namespace,omitempty"`

	// Conditions contains the list of conditions for the NVIDIADriver resource.
	// +optional
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}

// State represents the current state of a component
type State string

const (
	// Ready indicates the operational
	Ready State = "ready"
	// NotReady indicates the component is not yet operational
	NotReady State = "notReady"
	// indicates the component has been disabled
	Disabled State = "disabled"
	// Failed indicates the component has encountered a failure
	Failed State
// +kubebuilder:object:rootbuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
// +kubebuilder:printcolumn:name="Status",type=string,JSONPath=`.status.state`,priority=0
// +kubebuilder:printcolumn:name="Age",type=string,JSONPath=`.metadata.creationTimestamp`,priority=0

// NVIDIADriver is the Schema for the nvidiadrivers API
type NVIDIADriver struct {
	metav1.TypeMeta   `json:",av1.ObjectMeta `json:"metadata,omitempty"`

	Spec   NVIDIADriverSpec   `json:"spec,omitempty"`
	Status NVIDIADriverStatus `json:"status,omitemn// +kubebuilder:object:root=true

// NVIDIADriverList contains a list of NVIDIADriver
type NVIDIADriverList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NVIDIADriver `json:"items"`
}

func init() {
	SchemeBuilder.Register(&NVIDIADriver{}, &NVIDIADriverList{})
}
