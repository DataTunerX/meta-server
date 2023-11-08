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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type FinetuneState string

const (
	FinetuneInit       FinetuneState = "INIT"
	FinetunePending    FinetuneState = "PENDING"
	FinetuneRunning    FinetuneState = "RUNNING"
	FinetuneFailed     FinetuneState = "FAILED"
	FinetuneSuccessful FinetuneState = "SUCCESSFUL"
)

// FinetuneSpec defines the desired state of Finetune
type FinetuneSpec struct {
	// LLM specifies the large model CR used for fine-tuning.
	// +kubebuilder:validation:Required
	LLM string `json:"llm"`

	// Dataset specifies the dataset CR used for fine-tuning.
	// +kubebuilder:validation:Required
	Dataset string `json:"dataset"`

	// Hyperparameter specifies the hyperparameter CR used for fine-tuning.
	// +kubebuilder:validation:Optional
	Hyperparameter Hyperparameter `json:"hyperparameter"`

	// Resource specifies the resource configuration for fine-tuning.
	// +kubebuilder:validation:Optional
	Resource *Resource `json:"resource,omitempty"`
	// +kubebuilder:validation:Optional
	Node int `json:"node,omitempty"`
	// +kubebuilder:validation:Optional
	Image *ImageSetting `json:"image,omitempty"`
}

type Hyperparameter struct {
	HyperparameterRef string            `json:"hyperparameterRef,omitempty"`
	Overrides         map[string]string `json:"overrides,omitempty"`
}

type ImageSetting struct {
	// Name describe the image name.
	Name *string `json:"name"`
	// ImagePullPolicy describes a policy for if/when to pull a container image.
	// +kubebuilder:default:=IfNotPresent
	// +kubebuilder:validation:Enum=Always;IfNotPresent;Never
	// +kubebuilder:validation:Optional
	ImagePullPolicy *corev1.PullPolicy `json:"imagePullPolicy,omitempty"`
	// Path description of the model file path.
	// +kubebuilder:validation:Optional
	Path string `json:"path,omitempty"`
}

// FinetuneStatus defines the observed state of Finetune
type FinetuneStatus struct {
	// +kubebuilder:validation:Enum=INIT;PENDING;RUNNING;FAILED;SUCCESSFUL
	State FinetuneState `json:"state"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// Finetune is the Schema for the finetunes API
type Finetune struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   FinetuneSpec   `json:"spec,omitempty"`
	Status FinetuneStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// FinetuneList contains a list of Finetune
type FinetuneList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Finetune `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Finetune{}, &FinetuneList{})
}
