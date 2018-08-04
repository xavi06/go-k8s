package kubernetes

import (
	"errors"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	// Uncomment the following line to load the gcp plugin (only required to authenticate against GKE clusters).
	// _ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
)

// GetNodes func
func GetNodes(clientset *kubernetes.Clientset) ([]Node, error) {
	nodes, err := clientset.CoreV1().Nodes().List(metav1.ListOptions{})
	var data []Node
	if err != nil {
		return data, err
	}
	for _, Item := range nodes.Items {
		Name := Item.Name
		IP := Item.Annotations["flannel.alpha.coreos.com/public-ip"]
		Status := Item.Status.Conditions[3].Status
		CPUByte, err := Item.Status.Allocatable["cpu"].MarshalJSON()
		if err != nil {
			continue
		}
		// byte to string
		CPU := string(CPUByte[:])
		MemoryByte, err := Item.Status.Allocatable["memory"].MarshalJSON()
		if err != nil {
			continue
		}
		Memory := string(MemoryByte[:])
		Node := Node{Name, IP, Status, CPU, Memory}
		data = append(data, Node)
	}
	return data, nil
}

// GetNodeDetails func
func GetNodeDetails(clientset *kubernetes.Clientset, labelSelector string) (NodeDetails, error) {
	nodes, err := clientset.CoreV1().Nodes().List(metav1.ListOptions{LabelSelector: labelSelector})
	var data NodeDetails
	if err != nil {
		return data, err
	}
	count := len(nodes.Items)
	if count != 1 {
		return data, errors.New("error")
	}
	item := nodes.Items[0]
	Name := item.Name
	Labels := item.Labels
	Annotations := item.Annotations
	CreationTimestamp := item.CreationTimestamp
	Address := item.Status.Addresses
	PodCIDR := item.Spec.PodCIDR
	Unschedulable := item.Spec.Unschedulable
	NodeInfo := item.Status.NodeInfo
	Conditions := item.Status.Conditions
	Allocatable := item.Status.Allocatable
	Capacity := item.Status.Capacity

	return NodeDetails{Name, Labels, Annotations, CreationTimestamp, Address, PodCIDR, Unschedulable, NodeInfo, Conditions, Allocatable, Capacity}, nil
}
