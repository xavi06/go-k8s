package kubernetes

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

// GetServices func
func GetServices(clientset *kubernetes.Clientset, namespace string, labelSelector string) ([]Service, error) {
	var data []Service
	services, err := clientset.CoreV1().Services(namespace).List(metav1.ListOptions{LabelSelector: labelSelector})
	if err != nil {
		return data, err
	}
	for _, Item := range services.Items {
		Name := Item.Name
		Labels := Item.Labels
		ClusterIP := Item.Spec.ClusterIP
		Ports := Item.Spec.Ports
		CreationTimestamp := Item.CreationTimestamp
		data = append(data, Service{Name, Labels, ClusterIP, Ports, CreationTimestamp})
	}
	return data, nil
}
