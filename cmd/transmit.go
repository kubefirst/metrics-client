/*
Copyright (C) 2021-2023, Kubefirst

This program is licensed under MIT.
See the LICENSE file for more details.
*/
package cmd

import (
	"os"

	"github.com/denisbrodbeck/machineid"
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

		segmentClient := analytics.New(telemetry.SegmentIOWriteKey)
		defer segmentClient.Close()

		event := telemetry.TelemetryEvent{
			CliVersion:        os.Getenv("CLI_VERSION"),
			CloudProvider:     os.Getenv("CLOUD_PROVIDER"),
			ClusterID:         os.Getenv("CLUSTER_ID"),
			ClusterType:       os.Getenv("CLUSTER_TYPE"),
			DomainName:        strippedDomainName, // done
			GitProvider:       os.Getenv("GIT_PROVIDER"),
			InstallMethod:     os.Getenv("INSTALL_METHOD"),
			KubefirstClient:   os.Getenv("KUBEFIRST_CLIENT"),
			KubefirstTeam:     os.Getenv("KUBEFIRST_TEAM"),
			KubefirstTeamInfo: os.Getenv("KUBEFIRST_TEAM_INFO"),
			MachineID:         machineID, // done
			ErrorMessage:      err.Error(),
			UserId:            machineID,
			MetricName:        telemetry.MetricClusterInstallStarted,
		}

		switch transmitType {
		case "cluster-zero":
			//started event
			err := telemetry.SendCountMetric(event)
			if err != nil {
				log.Error(err)
			}
			log.Infof("metrics transmitted: %s", event.MetricName)

			//completed event
			event.MetricName = telemetry.MetricClusterInstallCompleted
			err = telemetry.SendCountMetric(event)
			if err != nil {
				log.Error(err)
			}
			log.Infof("metrics transmitted: %s", event.MetricName)
		default:
			log.Errorf("%s is not an allowed option", transmitType)
		}
	},
}

func init() {
	rootCmd.AddCommand(transmitCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// transmitCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// transmitCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	transmitCmd.Flags().StringVar(&transmitType, "type", "", "the type of metric to transmit [cluster-zero] (required)")
	transmitCmd.MarkFlagRequired("type")
}
