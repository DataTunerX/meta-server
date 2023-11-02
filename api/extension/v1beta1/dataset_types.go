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

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.
// SubTask defines a dataset task's subtask
type SubTask struct {
	// +kubebuilder:validation:MaxLength=63
	Name string `json:"name,omitempty"`
}

// Task defines a dataset task type
type Task struct {
	// +kubebuilder:validation:MaxLength=63
	Name     string    `json:"name,omitempty"`
	SubTasks []SubTask `json:"subTasks,omitempty"`
}

// Train defines a dataset's subsets' train-splits file address
type Train struct {
	File string `json:"file,omitempty"`
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
	Train    Train    `json:"train,omitempty"`
	Test     Test     `json:"test,omitempty"`
	Validate Validate `json:"validate,omitempty"`
}

// Subset defines a dataset‘s subset
type Subset struct {
	// +kubebuilder:validation:MaxLength=63
	Name   string `json:"name,omitempty"`
	Splits Splits `json:"splits,omitempty"`
}

//

// Feature defines a dataset's column name as a feature and its data type and relationship to finetune feature fields
type Feature struct {
	// +kubebuilder:validation:Enum=instruction;response
	Name string `json:"name,omitempty"`
	// +kubebuilder:validation:MaxLength=63
	MapTo    string `json:"mapTo,omitempty"`
	DataType string `json:"dataType,omitempty"`
}

// DatasetInfo defines the dataset description include subsets and features
type DatasetInfo struct {
	Subsets  []Subset  `json:"subsets,omitempty"`
	Features []Feature `json:"features,omitempty"`
}

// Plugin defines a plugin used by a dataset
type Plugin struct {
	// +kubebuilder:default:=true
	LoadPlugin bool `json:"loadPlugin,omitempty"`
	// +kubebuilder:validation:MaxLength=63
	Name       string `json:"name,omitempty"`
	Parameters string `json:"parameters,omitempty"`
}

// DatasetMetadata defines the metadata fields in DatasetSpec
type DatasetMetadata struct {
	Languages   []string     `json:"languages,omitempty"`
	Tags        []string     `json:"tags,omitempty"`
	License     string       `json:"license,omitempty"`
	Size        string       `json:"size,omitempty"`
	Task        Task         `json:"task,omitempty"`
	DatasetInfo *DatasetInfo `json:"datasetInfo,omitempty"`
	Plugin      Plugin       `json:"plugin,omitempty"`
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
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	// +kubebuilder:validation:Required
	DatasetMetadata *DatasetMetadata `json:"datasetMetadata"`
	DatasetCard     *DatasetCard     `json:"datasetCard,omitempty"`
	DatasetFiles    *DatasetFiles    `json:"datasetFiles,omitempty"`
}

// DatasetState is an enum type for the State field
type DatasetState string

const (
	DatasetReady   DatasetState = "Ready"
	DatasetUnready DatasetState = "Unready"
)

// DatasetStatus defines the observed state of Dataset
type DatasetStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	State DatasetState `json:"state,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// Dataset is the Schema for the datasets API
type Dataset struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   *DatasetSpec   `json:"spec,omitempty"`
	Status *DatasetStatus `json:"status,omitempty"`
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
