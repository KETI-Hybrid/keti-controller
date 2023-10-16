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

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// PodMetricSpec defines the desired state of PodMetric
type PodMetricSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of PodMetric. Edit podmetric_types.go to remove/update

	CPUCoreGauge     PodMetricValue `json:"cpuCoreGauge,omitempty"`
	CPUCoreCounter   PodMetricValue `json:"cpuCoreCounter,omitempty"`
	MemoryGauge      PodMetricValue `json:"memoryGauge,omitempty"`
	MemoryCounter    PodMetricValue `json:"memoryCounter,omitempty"`
	StorageGauge     PodMetricValue `json:"storageGauge,omitempty"`
	StorageCounter   PodMetricValue `json:"storageCounter,omitempty"`
	NetworkRXCounter PodMetricValue `json:"networkRXCounter,omitempty"`
	NetworkTXCounter PodMetricValue `json:"networkTXCounter,omitempty"`
	NetworkGauge     PodMetricValue `json:"networkGauge,omitempty"`
}

type PodMetricValue struct {
	ClusterName  string `json:"clustername,omitempty"`
	PodNamespace string `json:"podnamespace,omitempty"`
	PodName      string `json:"podname,omitempty"`
	Value        string `json:"value,omitempty"`
}

// PodMetricStatus defines the observed state of PodMetric
type PodMetricStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// PodMetric is the Schema for the podmetrics API
type PodMetric struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   PodMetricSpec   `json:"spec,omitempty"`
	Status PodMetricStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// PodMetricList contains a list of PodMetric
type PodMetricList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PodMetric `json:"items"`
}

func init() {
	SchemeBuilder.Register(&PodMetric{}, &PodMetricList{})
}
