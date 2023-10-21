/*
Copyright (C) 2021-2023, Kubefirst

This program is licensed under MIT.
See the LICENSE file for more details.
*/
package cmd

import (
	"os"

	"github.com/denisbrodbeck/machineid"
	"github.com/kubefirst/kubefirst-api/pkg/metrics"
	"github.com/kubefirst/kubefirst-api/pkg/segment"
	"github.com/kubefirst/metrics-client/pkg/telemetry"
	"github.com/kubefirst/metrics-client/pkg/utils"

	"github.com/segmentio/analytics-go"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	clusterId         string
	clusterType       string
	installMethod     string
	kubefirstTeam     string
	kubefirstTeamInfo string
	transmitType      string
)

// transmitCmd represents the transmit command
var transmitCmd = &cobra.Command{
	Use:   "transmit",
	Short: "transmit a metric",
	Long:  `transmit a metric`,
	Run: func(cmd *cobra.Command, args []string) {
		// Telemetry handler

		kubefirstVersion := os.Getenv("KUBEFIRST_VERSION")
		if kubefirstVersion == "" {
			kubefirstVersion = "development"
		}

		domainName := os.Getenv("DOMAIN_NAME")
		strippedDomainName, err := utils.RemoveSubdomainV2(domainName)
		if err != nil {
			log.Errorf("error encountered while reducing domain name. %s", err)
		}
		machineID, _ := machineid.ID()

		segmentClient := telemetry.SegmentClient{
			TelemetryEvent: telemetry.TelemetryEvent{
				CliVersion:        os.Getenv("CLI_VERSION"),
				CloudProvider:     os.Getenv("CLOUD_PROVIDER"),
				ClusterID:         os.Getenv("CLUSTER_ID"),
				ClusterType:       os.Getenv("CLUSTER_TYPE"),
				DomainName:        strippedDomainName,
				ErrorMessage:      err.Error(),
				GitProvider:       os.Getenv("GIT_PROVIDER"),
				InstallMethod:     os.Getenv("INSTALL_METHOD"),
				KubefirstClient:   os.Getenv("KUBEFIRST_CLIENT"),
				KubefirstTeam:     os.Getenv("KUBEFIRST_TEAM"),
				KubefirstTeamInfo: os.Getenv("KUBEFIRST_TEAM_INFO"),
				MachineID:         machineID,
				MetricName:        metrics.ClusterInstallStarted,
				UserId:            machineID,
			},
			Client: analytics.New(segment.SegmentIOWriteKey),
		}

		defer segmentClient.Client.Close()

		switch transmitType {
		case "cluster-zero":
			//started event
			err := telemetry.SendCountMetric(&segmentClient, metrics.ClusterInstallStarted)
			if err != nil {
				log.Error(err)
			}
			log.Infof("metrics transmitted: %s", segmentClient.TelemetryEvent.MetricName)

			//completed event
			segmentClient.TelemetryEvent.MetricName = metrics.ClusterInstallCompleted
			err = telemetry.SendCountMetric(&segmentClient, metrics.ClusterInstallCompleted)
			if err != nil {
				log.Error(err)
			}
			log.Infof("metrics transmitted: %s", segmentClient.TelemetryEvent.MetricName)
		default:
			log.Errorf("%s is not an allowed option", transmitType)
		}
	},
}

func init() {
	rootCmd.AddCommand(transmitCmd)

	transmitCmd.Flags().StringVar(&transmitType, "type", "", "the type of metric to transmit [cluster-zero] (required)")
	transmitCmd.MarkFlagRequired("type")
}
