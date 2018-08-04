package kubernetes

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

// GetDeployments func
func GetDeployments(clientset *kubernetes.Clientset, namespace string, labelSelector string) ([]Deployment, error) {
	deployments, err := clientset.AppsV1().Deployments(namespace).List(metav1.ListOptions{LabelSelector: labelSelector})
	var data []Deployment
	if err != nil {
		return data, err
	}
	for _, Item := range deployments.Items {
		Name := Item.Name
		Labels := Item.Labels
		AvailableReplicas := Item.Status.AvailableReplicas
		Replicas := Item.Status.Replicas
		CreationTimestamp := Item.CreationTimestamp
		data = append(data, Deployment{Name, Labels, AvailableReplicas, Replicas, CreationTimestamp})
	}
	return data, nil

}
