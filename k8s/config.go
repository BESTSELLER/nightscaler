package k8s

import (
	"os"
	"path/filepath"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

var config *rest.Config
var clientset *kubernetes.Clientset

func GetConfig() (*rest.Config, error) {
	if config != nil {
		return config, nil
	}

	config, err := rest.InClusterConfig()
	if err == nil {
		return config, nil
	}

	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}

	kubeconfig := filepath.Join(homeDir, ".kube", "config")
	config, err = clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err == nil {
		return config, nil
	}
	return nil, err
}

// this will be removed when we run it in the cluster
func GetClientset() (*kubernetes.Clientset, error) {
	if clientset != nil {
		return clientset, nil
	}

	config, err := GetConfig()
	if err != nil {
		return nil, err
	}

	// create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}
	return clientset, nil
}
