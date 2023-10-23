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

func SendEventV2(segmentIOWriteKey string, event TelemetryEvent, metricName string, errMsg string) error {

	client := analytics.New(segmentIOWriteKey)
	defer client.Close()

	if event.MetricName == ClusterInstallStarted {
		err := client.Enqueue(analytics.Identify{
			UserId: event.UserId,
			Type:   "identify",
		})
		if err != nil {
			return err
		}
	}
	err := client.Enqueue(analytics.Track{
		UserId: event.UserId,
		Event:  metricName,
		Properties: analytics.NewProperties().
			Set("cli_version", event.CliVersion).
			Set("cloud_provider", event.CloudProvider).
			Set("cluster_id", event.ClusterID).
			Set("cluster_type", event.ClusterType).
			Set("domain", event.DomainName).
			Set("git_provider", event.GitProvider).
			Set("client", event.KubefirstClient).
			Set("kubefirst_team", event.KubefirstTeam).
			Set("kubefirst_team_info", event.KubefirstTeamInfo).
			Set("machine_id", event.MachineID).
			Set("error", errMsg).
			Set("install_method", event.InstallMethod),
	})
	if err != nil {
		return err
	}

	return nil
}
