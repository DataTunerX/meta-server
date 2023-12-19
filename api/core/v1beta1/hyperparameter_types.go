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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	finetunev1beta1 "github.com/DataTunerX/meta-server/api/finetune/v1beta1"
)

type HyperparameterObjectiveType string

const (
	SFT HyperparameterObjectiveType = "SFT"
)

// HyperparameterSpec defines the desired state of Hyperparameter
type HyperparameterSpec struct {
	// +kubebuilder:validation:Required
	// Over-senator adjustments to achieve targets.
	Objective Objective `json:"objective"`
	// +kubebuilder:validation:Required
	// Finetune paramenter config.
	Parameters Parameters `json:"parameters"`
}

// The goal is usually a specific objective that is desired to be achieved through hyperparameter tuning,
// such as minimising losses or maximising some evaluation metric (e.g. accuracy, recall, etc.)
type Objective struct {
	/*
		The type (Type) of fine-tuning may refer to a specific method or strategy of hyperparameter tuning. Common hyperparameter tuning methods include:
		1. SFT
	*/
	// +kubebuilder:validation:Required
	// +kubebuilder:default:=SFT
	Type HyperparameterObjectiveType `json:"type"`
}

type Parameters struct {
	// Scheduler specifies the learning rate scheduler.
	// +kubebuilder:default:=cosine
	Scheduler finetunev1beta1.HyperparameterScheduler `json:"scheduler"`
	// Optimizer specifies the optimization algorithm.
	// +kubebuilder:default:=adamw_torch
	Optimizer finetunev1beta1.HyperparameterOptimizer `json:"optimizer"`
	// Int4 indicates whether to use 4-bit integer quantization.
	// +kubebuilder:default:=false
	Int4 bool `json:"int4"`
	// Int8 indicates whether to use 8-bit integer quantization.
	// +kubebuilder:default:=false
	Int8 bool `json:"int8"`
	// LoRA_R represents the radius parameter for Localized Receptive Attention.
	// +kubebuilder:default:="4"
	LoRA_R string `json:"loRA_R"`
	// LoRA_Alpha represents the alpha parameter for Localized Receptive Attention.
	// +kubebuilder:default:="32.0"
	LoRA_Alpha string `json:"loRA_Alpha"`
	// LoRA_Dropout specifies the dropout rate for Localized Receptive Attention.
	// +kubebuilder:default:="0.1"
	LoRA_Dropout string `json:"loRA_Dropout"`
	// LearningRate specifies the initial learning rate.
	// +kubebuilder:default:="0.001"
	LearningRate string `json:"learningRate"`
	// Epochs specifies the number of training epochs.
	// +kubebuilder:default:=10
	Epochs int `json:"epochs"`
	// BlockSize specifies the block size.
	// +kubebuilder:default:=512
	BlockSize int `json:"blockSize"`
	// BatchSize specifies the size of mini-batches.
	// +kubebuilder:default:=32
	BatchSize int `json:"batchSize"`
	// WarmupRatio specifies the ratio of warmup steps.
	// +kubebuilder:default:="0.1"
	WarmupRatio string `json:"warmupRatio"`
	// WeightDecay specifies the weight decay factor.
	// +kubebuilder:default:="0.0001"
	WeightDecay string `json:"weightDecay"`
	// GradAccSteps specifies the number of gradient accumulation steps.
	// +kubebuilder:default:=1
	GradAccSteps int `json:"gradAccSteps"`
	// TrainerType specifies the type of trainer to use.
	// +kubebuilder:default:=Standard
	// +kubebuilder:validation:Enum=Standard
	TrainerType string `json:"trainerType"`
	// PEFT indicates whether to enable Performance Evaluation and Forecasting Tool.
	// +kubebuilder:validation:Optional
	// +kubebuilder:default:=true
	PEFT bool `json:"PEFT,omitempty"`
	// FP16 indicates whether to use 16-bit floating point precision.
	// +kubebuilder:default:=false
	FP16 bool `json:"FP16"`
}

// HyperparameterStatus defines the observed state of Hyperparameter
type HyperparameterStatus struct {
	ReferenceFinetuneName []string `json:"referenceFinetuneName,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// Hyperparameter is the Schema for the hyperparameters API
type Hyperparameter struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   HyperparameterSpec   `json:"spec,omitempty"`
	Status HyperparameterStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// HyperparameterList contains a list of Hyperparameter
type HyperparameterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Hyperparameter `json:"items"`
}
