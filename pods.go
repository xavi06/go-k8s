package kubernetes

import (
	"fmt"
	"strings"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

// GetPods func
func GetPods(clientset *kubernetes.Clientset, namespace string, labelSelector string) ([]Pod, error) {
	pods, err := clientset.CoreV1().Pods(namespace).List(metav1.ListOptions{LabelSelector: labelSelector})
	var data []Pod
	if err != nil {
		return data, err
	}

	for _, Item := range pods.Items {
		Name := Item.Name
		NodeName := Item.Spec.NodeName
		Labels := Item.Labels
		var labelList []string
		for key, val := range Labels {
			label := fmt.Sprintf("%s=%s", key, val)
			labelList = append(labelList, label)
		}
		LabelsFmt := strings.Join(labelList, ",")
		Namespace := Item.Namespace
		CreationTimestamp := Item.CreationTimestamp
		Phase := Item.Status.Phase
		PodIP := Item.Status.PodIP
		QOSClass := Item.Status.QOSClass
		StartTime := Item.Status.StartTime
		Containers := Item.Spec.Containers
		Conditions := Item.Status.Conditions
		ContainerStatuses := Item.Status.ContainerStatuses
		data = append(data, Pod{Name, NodeName, Labels, LabelsFmt, Namespace, CreationTimestamp, Phase, PodIP, QOSClass,
			StartTime, Containers, Conditions, ContainerStatuses})
	}
	return data, nil
}

//DeletePods func
func DeletePods(clientset *kubernetes.Clientset, namespace string, name string) error {
	err := clientset.CoreV1().Pods(namespace).Delete(name, &metav1.DeleteOptions{})
	if err != nil {
		return err
	}
	return nil

}
