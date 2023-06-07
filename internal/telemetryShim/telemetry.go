/*
Copyright (C) 2021-2023, Kubefirst

This program is licensed under MIT.
See the LICENSE file for more details.
*/
package telemetryShim

import (
	"github.com/kubefirst/runtime/configs"
	"github.com/kubefirst/runtime/pkg/segment"
	log "github.com/sirupsen/logrus"
)

func SetupInitialTelemetry(
	clusterID string,
	clusterType string,
	installMethod string,
	kubefirstTeam string,
	kubefirstTeamInfo string,
) (*segment.SegmentClient, error) {
	// Segment Client
	segmentClient := &segment.SegmentClient{
		CliVersion:        configs.K1Version,
		ClusterID:         clusterID,
		ClusterType:       clusterType,
		KubefirstClient:   "api",
		KubefirstTeam:     kubefirstTeam,
		KubefirstTeamInfo: kubefirstTeamInfo,
		InstallMethod:     installMethod,
	}
	segmentClient.SetupClient()

	return segmentClient, nil
}

func TransmitClusterZero(useTelemetry bool, segmentClient *segment.SegmentClient, metricName string, errorMessage string) {
	if useTelemetry {
		segmentMsg := segmentClient.SendCountClusterZeroMetric(metricName, errorMessage)
		if segmentMsg != "" {
			log.Info(segmentMsg)
		}
	}
}
