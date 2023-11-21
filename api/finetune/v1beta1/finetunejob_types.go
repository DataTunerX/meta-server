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

package v1beta1

import (
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type FinetuneJobState string

const (
	FinetuneJobInit       FinetuneJobState = "INIT"
	FinetuneJobFailed     FinetuneJobState = "FAILED"
	FinetuneJobSuccessful FinetuneJobState = "SUCCESSFUL"
	FinetuneJobBuildImage FinetuneJobState = "BUILDIMAGE"
	FinetuneJobFinetune   FinetuneJobState = "FINETUNE"
	FinetuneJobServe      FinetuneJobState = "SERVE"
)

// FinetuneJobSpec defines the desired state of FinetuneJob
type FinetuneJobSpec struct {
	// Finetune cr config.
	// +kubebuilder:validation:Required
	FineTune FineTune `json:"finetune"`
	// Score plugin config.
	// +kubebuilder:validation:Optional
	ScoringConfig *ScoringConfig `json:"scoringConfig,omitempty"`
	// Serve config.
	// +kubebuilder:validation:Optional
	ServeConfig ServeConfig `json:"serveConfig"`
}

// ResourceLimits represents the resource limits for a task.
type ResourceLimits struct {
	// CPU specifies the CPU resource limit.
	// +kubebuilder:default="2"
	CPU resource.Quantity `json:"cpu"`

	// Memory specifies the memory resource limit.
	// +kubebuilder:default="4Gi"
	Memory resource.Quantity `json:"memory"`

	// GPU specifies the GPU resource limit.
	// +kubebuilder:default="1"
	GPU *string `json:"gpu"`
}

// Resource represents the resources configuration for a task.
type Resource struct {
	// Limits specifies the resource limits.
	Limits   *ResourceLimits `json:"limits"`
	Requests *ResourceLimits `json:"requests"`
}

// ServeConfig represents the configuration for serving with Ray.
type ServeConfig struct {
	// NodeSelector specifies the node where service will be deployed.
	// +kubebuilder:validation:Optional
	NodeSelector map[string]string `json:"nodeSelector"`

	// Tolerations specifies the tolerations for service.
	Tolerations []corev1.Toleration `json:"tolerations,omitempty"`
}

// ScoringConfig represents the configuration for scoring.
type ScoringConfig struct {
	// Name specifies the name of the scoring CR.
	Name       string `json:"name,omitempty"`
	Parameters string `json:"parameters,omitempty"`
}

type FineTune struct {
	// +kubebuilder:validation:Optional
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:Required
	FinetuneSpec FinetuneSpec `json:"finetuneSpec"`
}

// FinetuneJobStatus defines the observed state of FinetuneJob
type FinetuneJobStatus struct {
	// +kubebuilder:validation:Enum=INIT;FAILED;SUCCESSFUL;BUILDIMAGE;FINETUNE;SERVE
	// +kubebuilder:default=INIT
	State  FinetuneJobState   `json:"state"`
	Stats  string             `json:"stats,omitempty"`
	Result *FinetuneJobResult `json:"result,omitempty"`
	// +kubebuilder:validation:Enum=INIT;PENDING;RUNNING;FAILED;SUCCESSFUL
	FinetuneState FinetuneState      `json:"finetuneState,omitempty"`
	Conditions    []metav1.Condition `json:"conditions,omitempty"`
}

type FinetuneJobResult struct {
	ModelExportResult bool   `json:"modelExportResult"`
	Dashboard         string `json:"dashboard"`
	Serve             string `json:"serve"`
	Image             string `json:"image"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// FinetuneJob is the Schema for the finetunejobs API
type FinetuneJob struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   FinetuneJobSpec   `json:"spec,omitempty"`
	Status FinetuneJobStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// FinetuneJobList contains a list of FinetuneJob
type FinetuneJobList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FinetuneJob `json:"items"`
}
