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

	extensionv1beta1 "github.com/DataTunerX/meta-server/api/extension/v1beta1"
	finetunev1beta1 "github.com/DataTunerX/meta-server/api/finetune/v1beta1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

type LLMCheckpointState string

const (
	LLMCheckpointInit       LLMCheckpointState = "INIT"
	LLMCheckpointPending    LLMCheckpointState = "PENDING"
	LLMCheckpointRunning    LLMCheckpointState = "RUNNING"
	LLMCheckpointFailed     LLMCheckpointState = "FAILED"
	LLMCheckpointSuccessful LLMCheckpointState = "SUCCESSFUL"
)

// LLMCheckpointSpec defines the desired state of LLMCheckpoint
type LLMCheckpointSpec struct {
	// LLM specifies the large model CR used for fine-tuning.
	// +kubebuilder:validation:Required
	LLM *LLMRefSpec `json:"llm"`

	// Dataset specifies the dataset CR used for fine-tuning.
	// +kubebuilder:validation:Required
	Dataset *DatasetRefSpec `json:"dataset"`

	// Hyperparameter specifies the hyperparameter CR used for fine-tuning.
	// +kubebuilder:validation:Required
	Hyperparameter *HyperparameterRefSpec `json:"hyperparameter"`

	// Image specifies the image info used for fine-tuning.
	// +kubebuilder:validation:Required
	Image *finetunev1beta1.ImageSetting `json:"image"`

	// CheckpointImage specifies the checkpointImage info.
	// +kubebuilder:validation:Optional
	CheckpointImage *CheckpointImage `json:"checkpointImage,omitempty"`

	// Checkpoint specifies the checkpoint file.
	// +kubebuilder:validation:Required
	Checkpoint string `json:"checkpoint"`
}

// LLMRefSpec defines the ref of LLM
type LLMRefSpec struct {
	// LLMRef describe the llm ref.
	LLMRef string `json:"llmRef"`
	// Spec describe the llm spec.
	// +kubebuilder:validation:Required
	Spec *LLMSpec `json:"spec"`
}

// DatasetRefSpec defines the ref of Dataset
type DatasetRefSpec struct {
	// DatasetRef describe the dataset ref.
	DatasetRef string `json:"datasetRef"`
	// Spec describe the dataset spec.
	// +kubebuilder:validation:Required
	Spec *extensionv1beta1.DatasetSpec `json:"spec"`
}

// HyperparameterRefSpec defines the ref of Hyperparameter
type HyperparameterRefSpec struct {
	// HyperparameterRef describe the hyperparameter ref.
	HyperparameterRef string `json:"hyperparameterRef"`
	// Spec describe the hyperparameter spec.
	// +kubebuilder:validation:Required
	Spec *HyperparameterSpec `json:"spec"`
}

// CheckpointImage defines the checkpointImage
type CheckpointImage struct {
	// Name describe the image name.
	Name *string `json:"name"`
	// ImagePullPolicy describes a policy for if/when to pull a container image.
	// +kubebuilder:default:=IfNotPresent
	// +kubebuilder:validation:Enum=Always;IfNotPresent;Never
	// +kubebuilder:validation:Optional
	ImagePullPolicy *corev1.PullPolicy `json:"imagePullPolicy,omitempty"`
	// LLMPath description of the model file path.
	// +kubebuilder:validation:Optional
	LLMPath string `json:"llmPath,omitempty"`
	// CheckPointPath description of the checkpoint file path.
	// +kubebuilder:validation:Optional
	CheckPointPath string `json:"checkPointPath,omitempty"`
}

// LLMCheckpointStatus defines the observed state of LLMCheckpoint
type LLMCheckpointStatus struct {
	// +kubebuilder:validation:Enum=INIT;PENDING;RUNNING;FAILED;SUCCESSFUL
	State LLMCheckpointState `json:"state"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// LLMCheckpoint is the Schema for the llmcheckpoints API
type LLMCheckpoint struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   LLMCheckpointSpec   `json:"spec,omitempty"`
	Status LLMCheckpointStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// LLMCheckpointList contains a list of LLMCheckpoint
type LLMCheckpointList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LLMCheckpoint `json:"items"`
}

func init() {
	SchemeBuilder.Register(&LLMCheckpoint{}, &LLMCheckpointList{})
}
