package main

import (
	"context"
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"

	"log"
)

func main() {
	//var kubeconfig *string
	//if home := homedir.HomeDir(); home != "" {
	//	kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "kubeconfig to connect to cluster")
	//} else {
	//	kubeconfig = flag.String("kubeconfig", "", "complete path for kubeconfig")
	//}
	//flag.Parse()

	//conf, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	conf, err := rest.InClusterConfig()
	if err != nil {
		log.Fatal("couldn't able to build the config ", err)
	}

	client, err := kubernetes.NewForConfig(conf)
	if err != nil {
		log.Fatal("client is not created:", err)
	}

	namespace, err := client.CoreV1().Namespaces().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		log.Fatalf("couldnt able to talk to cluster and the error is: %v", err)
	}
	//fmt.Println(reflect.TypeOf(namespace))
	//println(namespace.Items)
	for _, s := range namespace.Items {
		//fmt.Println("I is", i, "S is", s)
		fmt.Println(s.Name)
	}

}
