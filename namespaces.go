package kubernetes

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

// GetNamespaces func
func GetNamespaces(clientset *kubernetes.Clientset) ([]Namespace, error) {
	namespaces, err := clientset.CoreV1().Namespaces().List(metav1.ListOptions{})
	var data []Namespace
	if err != nil {
		return data, err
	}
	for _, Item := range namespaces.Items {
		Name := Item.Name
		data = append(data, Namespace{Name})
	}
	return data, nil

}
