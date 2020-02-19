package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// Workflow is specification for a Workflow resource
type Workflow struct {
	metav1.TypeMeta	`json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

tatus WorkflowStatus `json:"status"`
    Spec   WorkflowSpec   `json:"spec"`
}

type WorkflowStatus struct {
    Name string `json:"name"`
}

// WorkflowSpec is the spec for a Workflow resource
type WorkflowSpec struct {
    Name string `json:"name"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// WorkflowList is a list of Workflow resources
type WorkflowList struct {
    metav1.TypeMeta `json:",inline"`
    metav1.ListMeta `json:"metadata"`

    Items []Workflow `json:"items"`
}
