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

// HyperparameterSpec defines the desired state of Hyperparameter
type HyperparameterSpec struct {
}

type Objective struct {
	Type string
}

// HyperparameterStatus defines the observed state of Hyperparameter
type HyperparameterStatus struct {
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

func init() {
	SchemeBuilder.Register(&Hyperparameter{}, &HyperparameterList{})
}
