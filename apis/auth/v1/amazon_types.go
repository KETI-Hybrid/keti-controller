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

// AmazonSpec defines the desired state of Amazon
type AmazonSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	AwsAccessKeyID     string `json:"awsAccessKeyID-id,omitempty"`
	AwsSecretAccessKey string `json:"awsSecretAccessKey,omitempty"`
	Region             string `json:"Region,omitempty"`
	DefaultARN         string `json:"DefaultARN,omitempty"`
}

// AmazonStatus defines the observed state of Amazon
type AmazonStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// Amazon is the Schema for the amazons API
type Amazon struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AmazonSpec   `json:"spec,omitempty"`
	Status AmazonStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// AmazonList contains a list of Amazon
type AmazonList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Amazon `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Amazon{}, &AmazonList{})
}
