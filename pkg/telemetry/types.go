/*
Copyright (C) 2021-2023, Kubefirst

This program is licensed under MIT.
See the LICENSE file for more details.
*/
package telemetry

import "github.com/segmentio/analytics-go"

type SegmentClient struct {
	TelemetryEvent TelemetryEvent
	Client         analytics.Client
}

type TelemetryEvent struct {
	CliVersion        string
	CloudProvider     string
	ClusterID         string
	ClusterType       string
	DomainName        string
	GitProvider       string
	InstallMethod     string
	KubefirstClient   string
	KubefirstTeam     string
	KubefirstTeamInfo string
	MachineID         string
	////////////////////////
	ErrorMessage string
	UserId       string
	MetricName   string
}
