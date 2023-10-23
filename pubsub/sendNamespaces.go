package pubsub

import (
	"strconv"

	"github.com/BESTSELLER/nightscaler/k8s"
	"github.com/BESTSELLER/nightscaler/logger"
	"github.com/BESTSELLER/nightscaler/timeutil"
)

func SendNamespaces() {
	namespaces, err := k8s.List()
	if err != nil {
		logger.Log.Fatal().Err(err).Msg("failed to list namespaces")
	}

	for _, ns := range namespaces {
		s := "Mon-Fri 13:30-23:59 Europe/Copenhagen"

		startTime, endTime := timeutil.GetUptimes(s)

		ns.Annotations["nightscaler/uptime-start"] = strconv.FormatInt(startTime.UnixMilli(), 10)
		ns.Annotations["nightscaler/uptime-end"] = strconv.FormatInt(endTime.UnixMilli(), 10)
	}

	err = Publish(namespaces, Attributes{
		Entity: "namespace",
		Action: "create",
	})
	if err != nil {
		logger.Log.Fatal().Err(err).Msg("failed to send Pub/Sub message")
	}
}
