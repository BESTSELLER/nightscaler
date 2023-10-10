package k8s

import (
	"context"

	globalconfig "github.com/BESTSELLER/nightscaler/config"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func List() (result []v1.Namespace, err error) {
	// create the clientset
	clientset, err := GetClientset()
	if err != nil {
		return nil, err
	}

	// get all namespaces
	namespaces, err := clientset.CoreV1().Namespaces().List(context.Background(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

	// loop through namespaces and check for annotation
	for _, ns := range namespaces.Items {
		if _, ok := ns.Annotations[globalconfig.NAMESPACE_ANNOTATION]; ok {
			result = append(result, ns)
		}
	}
	return result, err
}
