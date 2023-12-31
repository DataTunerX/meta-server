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

// ScoringPluginSpec defines the desired state of ScoringPlugin
type ScoringPluginSpec struct {
	Version string `json:"version,omitempty"`
	// +kubebuilder:validation:Required
	// ScoringClass describes the class type of scoring for example, the name of the plugin creator
	ScoringClass string `json:"scoringClass,omitempty"`
	// Provider describes the specific method of a DatasetClass e.g. MysqlServer, rabbitmq, s3, http etc.
	// Provider's value needs to correspond to the value of DatasetClass
	// +kubebuilder:validation:Required
	Provider string `json:"provider"`
	// +kubebuilder:validation:Required
	// Metrics describes the metrics for evaluate a LLMs performance, including Accuracy, Precision, Recall, F1-Score,
	// Log Loss, ROC & AUC, BLEU, ROUGE, Confusion Matrix, mAP, RMSE, MAE, R-squared, Jaccard Similarity, Contractive Ratio etc.
	Metrics []string `json:"metrics"`
	// Parameters describes data plugin parameters in key-value style with string type, e.g. "{'param1': 'value1', 'param2': 'value2'}"
	Parameters string `json:"parameters,omitempty"`
}

// DataPluginState is an enum type for the State field
type ScoringPluginState string

const (
	ScoringPluginReady   ScoringPluginState = "READY"
	ScoringPluginUnready ScoringPluginState = "UNREADY"
)

// ScoringPluginStatus defines the observed state of ScoringPlugin
type ScoringPluginStatus struct {
	// +kubebuilder:validation:Enum=READY;UNREADY
	State ScoringPluginState `json:"state,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// ScoringPlugin is the Schema for the scoringplugins API
type ScoringPlugin struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ScoringPluginSpec   `json:"spec,omitempty"`
	Status ScoringPluginStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// ScoringPluginList contains a list of ScoringPlugin
type ScoringPluginList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ScoringPlugin `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ScoringPlugin{}, &ScoringPluginList{})
}
