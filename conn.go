package kubernetes

import (
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	// Uncomment the following line to load the gcp plugin (only required to authenticate against GKE clusters).
	// _ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
)

// GetClientset func
func GetClientset(kubeconfigFile string) (*kubernetes.Clientset, error) {
	var kubeconfig *string
	kubeconfig = &kubeconfigFile
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err.Error())
		//return "", err
	}
	// create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		//return "", err
		panic(err.Error())
	}
	return clientset, err

}
