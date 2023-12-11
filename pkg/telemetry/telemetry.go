/*
Copyright (C) 2021-2023, Kubefirst

This program is licensed under MIT.
See the LICENSE file for more details.
*/
package telemetry

import (
	"os"

	"github.com/segmentio/analytics-go"
	log "github.com/sirupsen/logrus"
)

func SendEvent(event TelemetryEvent, metricName string, errMsg string) error {

	useTelemetryStr := os.Getenv("USE_TELEMETRY")
	useTelemetry := true
	if useTelemetryStr == "false" {
		useTelemetry = false
	}

	if !useTelemetry {
		log.Info("telemetry collection is disabled")
		return nil
	}

	client, err := analytics.NewWithConfig(SegmentIOWriteKey, analytics.Config{
		Interval:  3,
		BatchSize: 2,
	})

	defer client.Close()
	if err != nil {
		return err
	}

	if event.MetricName == ClusterInstallCompleted {
		err := client.Enqueue(analytics.Identify{
			UserId: event.UserId,
		})
		if err != nil {
			return err
		}
	}

	err = client.Enqueue(analytics.Track{
		UserId: event.UserId,
		Event:  metricName,
		Properties: analytics.NewProperties().
			Set("cli_version", event.CliVersion).
			Set("cloud_provider", event.CloudProvider).
			Set("cluster_id", event.ClusterID).
			Set("cluster_type", event.ClusterType).
			Set("domain", event.DomainName).
			Set("git_provider", event.GitProvider).
			Set("install_method", event.InstallMethod).
			Set("client", event.KubefirstClient).
			Set("kubefirst_team", event.KubefirstTeam).
			Set("kubefirst_team_info", event.KubefirstTeamInfo).
			Set("machine_id", event.MachineID).
			Set("parent_cluster_id", event.ParentClusterId).
			Set("error", errMsg),
	})
	if err != nil {
		return err
	}

	return nil
}
