/*
Copyright 2021 Syntasso.

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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

const WorkerResourceReplicas = -1
const ResourceRequestReplicas = 1

// WorkStatus defines the observed state of Work
type WorkStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// Work is the Schema for the works API
type Work struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   WorkSpec   `json:"spec,omitempty"`
	Status WorkStatus `json:"status,omitempty"`
}

// WorkSpec defines the desired state of Work
type WorkSpec struct {
	// Workload represents the manifest workload to be deployed on worker
	Workload WorkloadTemplate `json:"workload,omitempty"`

	// Scheduling is used for selecting the worker
	Scheduling WorkScheduling `json:"scheduling,omitempty"`

	// -1 denotes dependencies, 1 denotes Resource Request
	Replicas int `json:"replicas,omitempty"`
}

type WorkScheduling struct {
	Promise  []SchedulingConfig `json:"promise,omitempty"`
	Resource []SchedulingConfig `json:"resource,omitempty"`
}

func (w *Work) IsResourceRequest() bool {
	return w.Spec.Replicas == ResourceRequestReplicas
}

func (w *Work) IsWorkerResource() bool {
	return w.Spec.Replicas == WorkerResourceReplicas
}

func (w *Work) HasScheduling() bool {
	// Work has scheduling if either (or both) Promise or Resource has scheduling set
	return len(w.Spec.Scheduling.Resource) > 0 && len(w.Spec.Scheduling.Resource[0].Target.MatchLabels) > 0 ||
		len(w.Spec.Scheduling.Promise) > 0 && len(w.Spec.Scheduling.Promise[0].Target.MatchLabels) > 0
}

func (w *Work) GetSchedulingSelectors() map[string]string {
	return generateLabelSelectorsFromScheduling(append(w.Spec.Scheduling.Promise, w.Spec.Scheduling.Resource...))
}

// WorkloadTemplate represents the manifest workload to be deployed on worker
type WorkloadTemplate struct {
	// Manifests represents a list of resources to be deployed on the worker
	// +optional
	Manifests []Manifest `json:"manifests,omitempty"`
}

// Manifest represents a resource to be deployed on worker
type Manifest struct {
	// +kubebuilder:pruning:PreserveUnknownFields
	unstructured.Unstructured `json:",inline"`
}

//+kubebuilder:object:root=true

// WorkList contains a list of Work
type WorkList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Work `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Work{}, &WorkList{})
}
