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

// DataPluginSpec defines the desired state of DataPlugin
type DataPluginSpec struct {
	// +kubebuilder:validation:Required
	// DatasetClass describes the class type of dataset for example, the name of the plugin creator
	DatasetClass string `json:"datasetClass"`
	Version      string `json:"version,omitempty"`
	// Provider describes the specific method of a DatasetClass e.g. MysqlServer, rabbitmq, s3, http etc.
	// Provider's value needs to correspond to the value of DatasetClass
	// +kubebuilder:validation:Required
	Provider string `json:"provider"`
	// Parameters describes data plugin parameters in key-value style with string type, e.g. "{'param1': 'value1', 'param2': 'value2'}"
	Parameters string `json:"parameters,omitempty"`
}

// DataPluginState is an enum type for the State field
type DataPluginState string

const (
	DataPluginReady   DataPluginState = "Ready"
	DataPluginUnready DataPluginState = "Unready"
)

// DataPluginStatus defines the observed state of DataPlugin
type DataPluginStatus struct {
	// +kubebuilder:validation:Enum=READY;UNREADY
	State DataPluginState `json:"state,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// DataPlugin is the Schema for the dataplugins API
type DataPlugin struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DataPluginSpec   `json:"spec,omitempty"`
	Status DataPluginStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// DataPluginList contains a list of DataPlugin
type DataPluginList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DataPlugin `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DataPlugin{}, &DataPluginList{})
}
