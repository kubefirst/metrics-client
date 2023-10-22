/*
Copyright (C) 2021-2023, Kubefirst

This program is licensed under MIT.
See the LICENSE file for more details.
*/
package telemetry

import (
	"github.com/segmentio/analytics-go"
)

func SendEvent(segmentClient *SegmentClient, metricName string, errMsg string) error {
	if segmentClient.TelemetryEvent.MetricName == ClusterInstallStarted {
		err := segmentClient.Client.Enqueue(analytics.Identify{
			UserId: segmentClient.TelemetryEvent.UserId,
			Type:   "identify",
		})
		if err != nil {
			return err
		}
	}
	err := segmentClient.Client.Enqueue(analytics.Track{
		UserId: segmentClient.TelemetryEvent.UserId,
		Event:  metricName,
		Properties: analytics.NewProperties().
			Set("cli_version", segmentClient.TelemetryEvent.CliVersion).
			Set("cloud_provider", segmentClient.TelemetryEvent.CloudProvider).
			Set("cluster_id", segmentClient.TelemetryEvent.ClusterID).
			Set("cluster_type", segmentClient.TelemetryEvent.ClusterType).
			Set("domain", segmentClient.TelemetryEvent.DomainName).
			Set("git_provider", segmentClient.TelemetryEvent.GitProvider).
			Set("client", segmentClient.TelemetryEvent.KubefirstClient).
			Set("kubefirst_team", segmentClient.TelemetryEvent.KubefirstTeam).
			Set("kubefirst_team_info", segmentClient.TelemetryEvent.KubefirstTeamInfo).
			Set("machine_id", segmentClient.TelemetryEvent.MachineID).
			Set("error", errMsg).
			Set("install_method", segmentClient.TelemetryEvent.InstallMethod),
	})
	if err != nil {
		return err
	}

	return nil
}
