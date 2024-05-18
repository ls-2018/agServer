package v1alpha1

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type GuestBook struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,name=metadata"`

	Spec   GuestBookSpec   `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`
	Status GuestBookStatus `json:"status,omitempty" protobuf:"bytes,3,opt,name=status"`
}

type GuestBookSpec struct {
	InstanceAmount int `json:"instanceamount,omitempty" protobuf:"int32,1,opt,name=instanceamount"`
	InstanceCpu    int `json:"metadata,omitempty" protobuf:"int32,2,opt,name=instancecpu"`
}

type GuestBookStatus struct {
	ApprovalStatus string              `json:"approvalstatus" protobuf:"bytes,1,name=approvalstatus"`
	Instances      []GuestBookInstance `json:"instances" protobuf:"bytes,2,rep,name=instances"`
}

type GuestBookInstance struct {
	Cpu     int  `json:"cpu" protobuf:"int32,1,name=cpu"`
	Running bool `json:"running" protobuf:"bool,2,name=running"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type GuestBookList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	Items []GuestBook `json:"items" protobuf:"bytes,2,rep,name=items"`
}
