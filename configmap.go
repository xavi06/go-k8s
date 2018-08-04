package kubernetes

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

// GetConfigmap func
func GetConfigmap(clientset *kubernetes.Clientset, namespace string, labelSelector string) ([]Configmap, error) {
	configmaps, err := clientset.CoreV1().ConfigMaps(namespace).List(metav1.ListOptions{LabelSelector: labelSelector})
	var data []Configmap
	if err != nil {
		return data, err
	}
	for _, Item := range configmaps.Items {
		Data := Item.Data
		data = append(data, Configmap{Data})

	}
	return data, nil
}
