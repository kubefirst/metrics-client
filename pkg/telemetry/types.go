/*
Copyright (C) 2021-2023, Kubefirst

This program is licensed under MIT.
See the LICENSE file for more details.
*/
package telemetry

type TelemetryEvent struct {
	CliVersion        string
	CloudProvider     string
	ClusterID         string
	ClusterType       string
	DomainName        string
	ErrorMessage      string
	GitProvider       string
	InstallMethod     string
	KubefirstClient   string
	KubefirstTeam     string
	KubefirstTeamInfo string
	MachineID         string
	MetricName        string
	ParentClusterId   string
	UserId            string
}
