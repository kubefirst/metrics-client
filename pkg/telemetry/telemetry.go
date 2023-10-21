/*
Copyright (C) 2021-2023, Kubefirst

This program is licensed under MIT.
See the LICENSE file for more details.
*/
package telemetry

import (
	"strings"
	"time"

	"github.com/kubefirst/metrics-client/pkg/segment"
)

// Heartbeat
func Heartbeat(segmentClient *segment.SegmentClient) {
	// sent one heartbeat for the mgmt cluster
	segment.SendCountMetric(segmentClient)
	// Transmit(segmentClient, segment.MetricKubefirstHeartbeat, "")
	// // workload
	// HeartbeatWorkloadClusters()
	// //TODO! DIETZ - NO WAY
	for range time.Tick(time.Minute * 2) {
		segment.SendCountMetric(segmentClient)

		// // sent one heartbeat for the mgmt cluster
		// Transmit(segmentClient, segment.MetricKubefirstHeartbeat, "")
		// // workload
		// HeartbeatWorkloadClusters()
	}
}

func RemoveSubdomainV2(domainName string) (string, error) {

	domainName = strings.TrimRight(domainName, ".")
	domainSlice := strings.Split(domainName, ".")

	if len(domainSlice) < 2 {
		return "", nil
	}

	domainName = strings.Join([]string{domainSlice[len(domainSlice)-2], domainSlice[len(domainSlice)-1]}, ".")

	return domainName, nil
}
