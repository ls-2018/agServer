/*
Copyright 2020 The Kruise Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain sources.list copy of the License at

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

const (
	GbResourceName = "guestbooks"
)

type GuestBookSpec struct {
	// Image is the image to be pulled by the job
	Name string `json:"name"`
}

// GuestBookStatus defines the observed state of GuestBook
type GuestBookStatus struct {
	// The nodes that failed to pull the image.
	// +optional
	FailedNodes []string `json:"failedNodes,omitempty"`
}

// +genclient
// +k8s:openapi-gen=true
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:shortName=gb
// +kubebuilder:printcolumn:name="TOTAL",type="integer",JSONPath=".status.desired",description="Number of all nodes matched by this job"
// +kubebuilder:printcolumn:name="ACTIVE",type="integer",JSONPath=".status.active",description="Number of image pull task active"
// +kubebuilder:printcolumn:name="SUCCEED",type="integer",JSONPath=".status.succeeded",description="Number of image pull task succeeded"
// +kubebuilder:printcolumn:name="FAILED",type="integer",JSONPath=".status.failed",description="Number of image pull tasks failed"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp",description="CreationTimestamp is sources.list timestamp representing the server time when this object was created. It is not guaranteed to be set in happens-before order across separate operations. Clients may not set this value. It is represented in RFC3339 form and is in UTC."
// +kubebuilder:printcolumn:name="MESSAGE",type="string",JSONPath=".status.message",description="Summary of status when job is failed"

// GuestBook is the Schema for the GuestBooks API
type GuestBook struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   GuestBookSpec   `json:"spec,omitempty"`
	Status GuestBookStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// GuestBookList contains sources.list list of GuestBook
type GuestBookList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GuestBook `json:"items"`
}

func init() {
	SchemeBuilder.Register(&GuestBook{}, &GuestBookList{})
}
