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
type HyperparameterScheduler string
type HyperparameterOptimizer string

const (
	FinetuneInit       FinetuneState           = "INIT"
	FinetunePending    FinetuneState           = "PENDING"
	FinetuneRunning    FinetuneState           = "RUNNING"
	FinetuneFailed     FinetuneState           = "FAILED"
	FinetuneSuccessful FinetuneState           = "SUCCESSFUL"
	Cosine             HyperparameterScheduler = "Cosine"
	Linear             HyperparameterScheduler = "Linear"
	Constant           HyperparameterScheduler = "Constant"
	AdamW              HyperparameterOptimizer = "AdamW"
	Adam               HyperparameterOptimizer = "Adam"
	SGD                HyperparameterOptimizer = "SGD"
)

// FinetuneSpec defines the desired state of Finetune
type FinetuneSpec struct {
	// LLM specifies the large model CR used for fine-tuning.
	// +kubebuilder:validation:Required
	LLM string `json:"llm"`

	// Dataset specifies the dataset CR used for fine-tuning.
	// +kubebuilder:validation:Required
	Dataset string `json:"dataset"`

	// Hyperparameter specifies the hyperparameter used for fine-tuning.
	// +kubebuilder:validation:Required
	Hyperparameter Hyperparameter `json:"hyperparameter"`

	// Resource specifies the resource configuration for fine-tuning.
	// +kubebuilder:validation:Optional
	Resource *Resource `json:"resource,omitempty"`
	// +kubebuilder:default:1
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Minimum=1
	Node int `json:"node,omitempty"`
	// +kubebuilder:validation:Optional
	Image ImageSetting `json:"image,omitempty"`
}

type Hyperparameter struct {
	// HyperparameterRef specifies the hyperparameter CR used for fine-tuning.
	// +kubebuilder:validation:Required
	HyperparameterRef string `json:"hyperparameterRef"`
	// Overrides is used to override some hyperparameter
	// +kubebuilder:validation:Optional
	Overrides *Parameters `json:"overrides,omitempty"`
}

type Parameters struct {
	// Scheduler specifies the learning rate scheduler.
	// +kubebuilder:validation:Enum=Cosine;Linear;Constant
	Scheduler HyperparameterScheduler `json:"scheduler,omitempty"`
	// Optimizer specifies the optimization algorithm.
	// +kubebuilder:validation:Enum=AdamW;Adam;SGD
	Optimizer HyperparameterOptimizer `json:"optimizer,omitempty"`
	// Int4 indicates whether to use 4-bit integer quantization.
	Int4 *bool `json:"int4,omitempty"`
	// Int8 indicates whether to use 8-bit integer quantization.
	Int8 *bool `json:"int8,omitempty"`
	// LoRA_R represents the radius parameter for Localized Receptive Attention.
	LoRA_R *string `json:"loRA_R,omitempty"`
	// LoRA_Alpha represents the alpha parameter for Localized Receptive Attention.
	LoRA_Alpha *string `json:"loRA_Alpha,omitempty"`
	// LoRA_Dropout specifies the dropout rate for Localized Receptive Attention.
	LoRA_Dropout *string `json:"loRA_Dropout,omitempty"`
	// LearningRate specifies the initial learning rate.
	LearningRate *string `json:"learningRate,omitempty"`
	// Epochs specifies the number of training epochs.
	Epochs int `json:"epochs,omitempty"`
	// BlockSize specifies the block size.
	BlockSize int `json:"blockSize,omitempty"`
	// BatchSize specifies the size of mini-batches.
	BatchSize int `json:"batchSize,omitempty"`
	// WarmupRatio specifies the ratio of warmup steps.
	WarmupRatio *string `json:"warmupRatio,omitempty"`
	// WeightDecay specifies the weight decay factor.
	WeightDecay *string `json:"weightDecay,omitempty"`
	// GradAccSteps specifies the number of gradient accumulation steps.
	GradAccSteps int `json:"gradAccSteps,omitempty"`
	// TrainerType specifies the type of trainer to use.
	// +kubebuilder:validation:Enum=Standard
	TrainerType *string `json:"trainerType,omitempty"`
	// PEFT indicates whether to enable Performance Evaluation and Forecasting Tool.
	PEFT *bool `json:"PEFT,omitempty"`
	// FP16 indicates whether to use 16-bit floating point precision.
	FP16 *bool `json:"FP16,omitempty"`
}

type ImageSetting struct {
	// Name describe the image name.
	Name string `json:"name,omitempty"`
	// ImagePullPolicy describes a policy for if/when to pull a container image.
	// +kubebuilder:default:=IfNotPresent
	// +kubebuilder:validation:Enum=Always;IfNotPresent;Never
	// +kubebuilder:validation:Optional
	ImagePullPolicy corev1.PullPolicy `json:"imagePullPolicy,omitempty"`
	// Path description of the model file path.
	// +kubebuilder:validation:Optional
	Path string `json:"path,omitempty"`
}

// FinetuneStatus defines the observed state of Finetune
type FinetuneStatus struct {
	// +kubebuilder:validation:Enum=INIT;PENDING;RUNNING;FAILED;SUCCESSFUL
	State FinetuneState `json:"state"`
	// LLMCheckpoint describes the llmcheckpoint.
	LLMCheckpoint *Checkpoint `json:"llmCheckpoint,omitempty"`
	// RayJobInfo describes the rayjob.
	RayJobInfo *RayJobInfo `json:"rayJobInfo,omitempty"`
}

type RayJobInfo struct {
	RayJobPodName          string `json:"rayJobPodName,omitempty"`
	RayJobPodContainerName string `json:"rayJobPodContainerName,omitempty"`
}

type Checkpoint struct {
	LLMCheckpointRef string `json:"llmCheckpointRef"`
	CheckpointPath   string `json:"checkpointPath"`
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
