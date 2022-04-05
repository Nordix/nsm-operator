/*


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

package v1alpha1

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type Forwarder struct {
	// Forwarder descriptive name
	Name string `json:"Name"`
	// Forwarder image string
	// (must be a complete image path with tag)
	Image string `json:"Image"`
}

type Registry struct {
	// Registry type
	// +kubebuilder:validation:Enum=k8s;memory
	Type string `json:"type"`
	// Registry Image with tag
	Image string `json:"image,omitempty"`
}

// NSMSpec defines the desired state of NSM
type NSMSpec struct {
	// tag represents the desired nsm version
	Version       string            `json:"version"`
	NsmPullPolicy corev1.PullPolicy `json:"nsmPullPolicy"`
	Registry      Registry          `json:"registry"`
	NsmgrImage    string            `json:"nsmgrImage,omitempty"`
	Forwarders    []Forwarder       `json:"forwarders"`
}

// NSMPhase is the type for the operator phases
type NSMPhase string

// Operator phases
const (
	NSMPhaseInitial     NSMPhase = ""
	NSMPhasePending     NSMPhase = "Pending"
	NSMPhaseCreating    NSMPhase = "Creating"
	NSMPhaseRunning     NSMPhase = "Running"
	NSMPhaseTerminating NSMPhase = "Terminating"
)

// NSMStatus defines the observed state of NSM
type NSMStatus struct {
	// Operator phases during deployment
	Phase NSMPhase `json:"phase"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=nsms

// NSM is the Schema for the nsms API
type NSM struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   NSMSpec   `json:"spec,omitempty"`
	Status NSMStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// NSMList contains a list of NSM
type NSMList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NSM `json:"items"`
}

func init() {
	SchemeBuilder.Register(&NSM{}, &NSMList{})
}
