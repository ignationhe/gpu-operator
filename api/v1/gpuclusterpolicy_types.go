/*
Copyright 2024 NVIDIA CORPORATION & AFFILIATES.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS, either express or implied.
S License.
*/

.io/apimachinery/ of a component managed by the GPU Operator.
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
	// +kubebuilder:validation:Enum=docker;containerd
	// +kubebuilder:default=containerd
	DefaultRuntime string `json:"defaultRuntime,omitempty"`

	// InitContainerImage is the image used for init containers.
	// +optional
	InitContainerImage string `json:"initContainerImage,omitempty"`
}

// DriverSpec describes configuration for the NVIDIA driver component.
type DriverSpec struct {
	// Enabled controls whether the driver daemonset is deployed.
	// +kubebuilder:default=true
	Enabled *bool `json:"enabled,omitempty"`

	// Image is the NVIDIA driver container image.
	Image string `json:"image,omitempty"`

	// Version is the NVIDIA driver version to deploy.
	Version string `json:"version,omitempty"`
}

// ToolkitSpec d
