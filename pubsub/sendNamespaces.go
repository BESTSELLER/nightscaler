package pubsub

import (
	"strconv"

	"github.com/BESTSELLER/nightscaler/config"
	"github.com/BESTSELLER/nightscaler/k8s"
	"github.com/BESTSELLER/nightscaler/logger"
	"github.com/BESTSELLER/nightscaler/timeutil"
	v1 "k8s.io/api/core/v1"
)

func SendNamespaces() {
	namespaces, err := k8s.List()
	if err != nil {
		logger.Log.Error().Err(err).Msg("failed to list namespaces")
	}

	tempNamespaces := []v1.Namespace{}

	for _, ns := range namespaces {
		s := ns.Annotations[config.NAMESPACE_ANNOTATION]

		if s == "" {
			continue
		}

		startTime, endTime, err := timeutil.GetUptimes(s)
		if err != nil {
			logger.Log.Error().Err(err).Msg("failed to parse uptime")
			return
		}

		ns.Annotations["nightscaler/uptime-start"] = strconv.FormatInt(startTime.UnixMilli(), 10)
		ns.Annotations["nightscaler/uptime-end"] = strconv.FormatInt(endTime.UnixMilli(), 10)

		tempNamespaces = append(tempNamespaces, ns)
	}

	err = Publish(tempNamespaces, Attributes{
		Entity: "namespace",
		Action: "create",
	})
	if err != nil {
		logger.Log.Error().Err(err).Msg("failed to send Pub/Sub message")
		return
	}
}
