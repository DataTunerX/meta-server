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
)

// SubTask defines a dataset task's subtask
type SubTask struct {
	// +kubebuilder:validation:MaxLength=63
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Enum=language-modeling;masked-language-modeling;natural-language-inference;extractive-qa;open-domain-qa;multi-choice-qa;closed-domain-qa;multi-class-classification;sentiment-classification;topic-classification;multi-label-classification;news-articles-summarization;
	Name string `json:"name"`
}

// Task defines a dataset task type
type Task struct {
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:MaxLength=63
	// +kubebuilder:validation:Enum=Text Generation;Question Answering;Text Classification;Summarization
	Name string `json:"name"`
	// SubTask defines a dataset task's subtask e.g. language-modeling of Text Generation, open-domain-qa of Question Answering etc.
	// It is corresponding to Task.
	SubTasks []SubTask `json:"subTasks,omitempty"`
}

// Train defines a dataset's subsets' train-splits file address
type Train struct {
	// +kubebuilder:validation:Required
	File string `json:"file"`
}

// Test defines a dataset's subsets' test-splits file address
type Test struct {
	File string `json:"file,omitempty"`
}

// Validate defines a dataset's subsets' validate-splits file address
type Validate struct {
	File string `json:"file,omitempty"`
}

// Splits defines a dataset's train-splits, test-splits, validate-splits address info
type Splits struct {
	Train    *Train    `json:"train,omitempty"`
	Test     *Test     `json:"test,omitempty"`
	Validate *Validate `json:"validate,omitempty"`
}

// Subset defines a dataset‘s subset
type Subset struct {
	// +kubebuilder:validation:MaxLength=63
	// +kubebuilder:validation:Required
	// Subset e.g. Random Sample Subset, Balanced Class Subset, Time Window Subset, Feature Subset, Cross-Validation Subset, Outlier Detection Subset etc.
	Name string `json:"name,omitempty"`
	// Splits describes subsets of training and testing and validation data splits
	Splits Splits `json:"splits,omitempty"`
}

// Feature defines a dataset's column name as a feature and its data type and relationship to finetune feature fields
type Feature struct {
	// +kubebuilder:validation:Enum=instruction;response
	// +kubebuilder:validation:Required
	Name string `json:"name"`
	// +kubebuilder:validation:MaxLength=63
	// +kubebuilder:validation:Required
	// MapTo describes the dataste feature field mapping to instruction or response.
	MapTo string `json:"mapTo"`
	// +kubebuilder:default:=string
	DataType string `json:"dataType,omitempty"`
}

// DatasetInfo defines the dataset description include subsets and features
type DatasetInfo struct {
	Subsets  []Subset  `json:"subsets,omitempty"`
	Features []Feature `json:"features,omitempty"`
}

// Plugin defines a plugin used by a dataset
type Plugin struct {

	// +kubebuilder:default:=false
	// LoadPlugin describes a Scoring CR whether uses plugin to do evaluation, if true then Plugin data will be needed,
	LoadPlugin bool `json:"loadPlugin"`
	// +kubebuilder:validation:MaxLength=63
	Name       string `json:"name,omitempty"`
	Parameters string `json:"parameters,omitempty"`
}

// DatasetMetadata defines the metadata fields in DatasetSpec
type DatasetMetadata struct {
	// +kubebuilder:validation:Required
	// Languages includes Chinese and English and etc.
	Languages []string `json:"languages"`
	// Tags describes a dataset, it's customized.
	Tags []string `json:"tags,omitempty"`
	// License includes CC BY, CC BY-SA, CC BY-ND, CC BY-NC, CC BY-NC-SA, CC BY-NC-ND, CC0, ODC-PDDL, ODC-BY
	// ODC-ODbl, CDLA-Permissive-2.0, CDLA-Sharing-1.0
	License string `json:"license,omitempty"`
	// +kubebuilder:validation:Required
	// Size describes dataset's entries number
	Size string `json:"size"`
	// +kubebuilder:validation:Required
	// Task describes the main task that the dataset can do, including Text Generation, Question Answering,
	// Translation, Conversational etc.
	Task *Task `json:"task"`
	// DatasetInfo describes a dataset's subsets and Features.
	DatasetInfo *DatasetInfo `json:"datasetInfo,omitempty"`
	// Plugin describes the plugin including parameters and whether uses a plugin.
	Plugin *Plugin `json:"plugin,omitempty"`
}

// DatasetCard defines a dataset's readme, in type of markdown
type DatasetCard struct {
	DatasetCardRef string `json:"datasetCardRef,omitempty"`
}

// DatasetFiles defines a dataset's source file address
type DatasetFiles struct {
	Source string `json:"source,omitempty"`
}

// DatasetSpec defines the desired state of Dataset
type DatasetSpec struct {
	// +kubebuilder:validation:Required
	// DatasetMetadata contains a dataset's Language, Tags, Size, License, Task, Plugin, and baseinfo
	DatasetMetadata *DatasetMetadata `json:"datasetMetadata"`
	// DatasetCard contains a dataset's README reference.
	DatasetCard *DatasetCard `json:"datasetCard,omitempty"`
	// DatasetFiles describes a dataset source address.
	DatasetFiles *DatasetFiles `json:"datasetFiles,omitempty"`
}

// DatasetState is an enum type for the State field
type DatasetState string

const (
	DatasetReady   DatasetState = "READY"
	DatasetUnready DatasetState = "UNREADY"
)

// DatasetStatus defines the observed state of Dataset
type DatasetStatus struct {
	// +kubebuilder:validation:Enum=READY;UNREADY
	State                 DatasetState `json:"state,omitempty"`
	ReferenceFinetuneName []string     `json:"referenceFinetuneName,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// Dataset is the Schema for the datasets API
type Dataset struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DatasetSpec   `json:"spec,omitempty"`
	Status DatasetStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// DatasetList contains a list of Dataset
type DatasetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Dataset `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Dataset{}, &DatasetList{})
}
