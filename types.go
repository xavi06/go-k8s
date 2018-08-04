package kubernetes

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Namespace struct
type Namespace struct {
	Name string
}

// Node struct
type Node struct {
	Name   string
	IP     string
	Status corev1.ConditionStatus
	CPU    string
	Memory string
}

// NodeDetails struct
type NodeDetails struct {
	Name              string
	Labels            map[string]string
	Annotations       map[string]string
	CreationTimestamp metav1.Time
	Address           []corev1.NodeAddress
	PodCIDR           string
	Unschedulable     bool
	NodeInfo          corev1.NodeSystemInfo
	Conditions        []corev1.NodeCondition
	Allocatable       corev1.ResourceList
	Capacity          corev1.ResourceList
}

// Deployment struct
type Deployment struct {
	Name              string
	Labels            map[string]string
	AvailableReplicas int32
	Replicas          int32
	CreationTimestamp metav1.Time
}

// Pod struct
type Pod struct {
	Name              string
	NodeName          string
	Labels            map[string]string
	LabelsFmt         string
	Namespace         string
	CreationTimestamp metav1.Time
	Phase             corev1.PodPhase
	PodIP             string
	QOSClass          corev1.PodQOSClass
	StartTime         *metav1.Time
	Containers        []corev1.Container
	Conditions        []corev1.PodCondition
	ContainerStatuses []corev1.ContainerStatus
}

// Configmap struct
type Configmap struct {
	Data map[string]string
}

// Service struct
type Service struct {
	Name              string
	Labels            map[string]string
	ClusterIP         string
	Ports             []corev1.ServicePort
	CreationTimestamp metav1.Time
}
