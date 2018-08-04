package main

import (
	"fmt"
	"zj/kubernetes"
)

func main() {
	clientset, err := kubernetes.GetClientset("/Users/xavi/.kube/config")
	if err != nil {
		panic(err.Error())
	}
	//res, err := kubernetes.GetNodes(clientset)
	res, _ := kubernetes.GetPods(clientset, "default", "")
	fmt.Println(res[0])
	//res, err := kubernetes.GetNodeDetails(clientset, "kubernetes.io/hostname=wgkvm-k8s-master01")
	//fmt.Println(res.Conditions)
	//res, err := kubernetes.GetConfigmap(clientset, "ops", "app=projects")
	//fmt.Println(res)
	//res, _ := kubernetes.GetDeployments(clientset, "default", "")
	//fmt.Println(res)
}
