package k8s

import (
	"context"
	"fmt"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func ScaleUp(namespace string, additionalTime time.Duration) error {
	// create the clientset
	clientset, err := GetClientset()
	if err != nil {
		return err
	}

	location, err := time.LoadLocation("UTC")
	if err != nil {
		return err
	}

	t := time.Now().In(location)
	now := t.Format("2006-01-02T15:04:05-07:00")
	end := t.Add(additionalTime).Format("2006-01-02T15:04:05-07:00")
	timeString := fmt.Sprintf("%s-%s", now, end)

	// Get namespace
	result, err := clientset.CoreV1().Namespaces().Get(context.TODO(), namespace, metav1.GetOptions{})
	if err != nil {
		return err
	}

	// Set annotation on namespace
	result.Annotations["downscaler/force-uptime"] = timeString

	// Update namespace
	_, err = clientset.CoreV1().Namespaces().Update(context.TODO(), result, metav1.UpdateOptions{})

	return err
}
