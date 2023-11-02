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
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +kubebuilder:validation:Enum=SupervisionMethod;LearningMethod;MachineLearning;Modality
type LLMType string

// +kubebuilder:validation:Enum=Multimodal;ComputerVision;NLP;Audio;Tabular;ReinforcementLearning
type LLMDomain string

const (
	SupervisionMethod     LLMType   = "SupervisionMethod"
	LearningMethod        LLMType   = "LearningMethod"
	MachineLearning       LLMType   = "MachineLearning"
	Modality              LLMType   = "Modality"
	Multimodal            LLMDomain = "Multimodal"
	ComputerVision        LLMDomain = "ComputerVision"
	NLP                   LLMDomain = "NLP"
	Audio                 LLMDomain = "Audio"
	Tabular               LLMDomain = "Tabular"
	ReinforcementLearning LLMDomain = "ReinforcementLearning"
)

// LLMSpec defines the desired state of LLM
type LLMSpec struct {
	// +kubebuilder:validation:Required
	// Describe metadata information for llm.
	LLMMetdata LLMMetdata `json:"llmMetdata"`
	// +kubebuilder:validation:Optional
	// Describe readme information for llm.
	LLMCard *LLMCard `json:"llmCard,omitempty"`
	// +kubebuilder:validation:Optional
	// Describe file information for llm.
	LLMFiles *LLMFiles `json:"llmFiles,omitempty"`
}

type LLMMetdata struct {
	// +kubebuilder:validation:Required
	// Describe the name of the llm.
	Name string `json:"name"`
	// +kubebuilder:validation:Optional
	// Describe the domain of the llm, the following fields are optional:
	// 1. Multimodal
	// 2. ComputerVision
	// 3. NLP
	// 4. Audio
	// 5. Tabular
	// 6. ReinforcementLearning
	Domain []LLMDomain `json:"domain,omitempty"`
	// +kubebuilder:validation:Required
	// Describe the types of tasks in the pipeline corresponding to the llm.
	// example:
	/*
			...
			tasks:
			  - feature-extraction
		      - text-to-image
		    ...
	*/
	Tasks []string `json:"tasks"`
	// +kubebuilder:validation:Required
	// Describe the AI framework used by the llm.
	Frameworks []string `json:"frameworks"`
	// +kubebuilder:validation:Required
	// Describe list of natural languages supported by the llm.
	Languages []string `json:"languages"`
	// +kubebuilder:validation:Required
	// Describe the open source licences to which the llm adheres.
	License []string `json:"license"`
	// +kubebuilder:validation:Optional
	// Custom labels for llm.
	Tags []string `json:"tags,omitempty"`
	// +kubebuilder:validation:Optional
	// Describe the datasets used in the llm.
	Datasets []string `json:"datasets,omitempty"`
	// +kubebuilder:validation:Optional
	// Describe the llm on which the model was fine-tuned.
	BaseLLM string `json:"baseLlm,omitempty"`
	// +kubebuilder:validation:Required
	// Describe the llm image config.
	LLMImageConfig LLMImageConfig `json:"llmImageConfig"`
	// +kubebuilder:validation:Optional
	// Describe Hardware and software information the llm's operation
	ComputeInfrastructure *ComputeInfrastructure `json:"computeInfrastructure,omitempty"`
}

type LLMCard struct {
	// +kubebuilder:validation:Optional
	// Readme information of the model, parsed in markdown format,
	// is presented as information on the model page to help model users understand the model and use it correctly.
	// Describe references to readme files mounted by configmap.
	LLMCardRef string `json:"llmCardRef"`
}

type LLMFiles struct {
	// +kubebuilder:validation:Optional
	// Describe llm source file citation address.
	Source string `json:"source"`
}

type ComputeInfrastructure struct {
	// +kubebuilder:validation:Optional
	// Describe the hardware information required for the llm to operate
	Hardware *Hardware `json:"hardware,omitempty"`
}

type Hardware struct {
	// Describes the size of the video memory required by the llm.
	VRam resource.Quantity `json:"vRam"`
	// Description of the number of cpu cores needed for llm.
	Cpu resource.Quantity `json:"cpu"`
	// Description of the memory size required by llm
	Memory resource.Quantity `json:"memory"`
}

type LLMImageConfig struct {
	// +kubebuilder:validation:Required
	Image string `json:"image"`
	// +kubebuilder:validation:Optional
	Path string `json:"path"`
}

// LLMStatus defines the observed state of LLM
type LLMStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// LLM is the Schema for the llms API
type LLM struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   LLMSpec   `json:"spec,omitempty"`
	Status LLMStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// LLMList contains a list of LLM
type LLMList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LLM `json:"items"`
}
