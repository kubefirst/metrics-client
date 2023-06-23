/*
Copyright (C) 2021-2023, Kubefirst

This program is licensed under MIT.
See the LICENSE file for more details.
*/
package cmd

import (
	"github.com/kubefirst/metrics-client/internal/telemetryShim"
	"github.com/kubefirst/runtime/pkg/segment"
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
		segmentClient, err := telemetryShim.SetupInitialTelemetry(
			clusterId,
			clusterType,
			installMethod,
			kubefirstTeam,
			kubefirstTeamInfo,
		)
		if err != nil {
			log.Warn(err)
		}
		defer segmentClient.Client.Close()

		switch transmitType {
		case "cluster-zero":
			if installMethod == "" {
				log.Fatalf("when transmitting cluster-zero metric type, install-method is a required value")
			}
			telemetryShim.TransmitClusterZero(true, segmentClient, segment.MetricClusterInstallStarted, "")
			telemetryShim.TransmitClusterZero(true, segmentClient, segment.MetricClusterInstallCompleted, "")
			log.Infof(
				"metrics transmitted: %s/%s %s and %s",
				clusterId,
				clusterType,
				segment.MetricClusterInstallStarted,
				segment.MetricClusterInstallCompleted,
			)
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
	transmitCmd.Flags().StringVar(&clusterId, "cluster-id", "", "the kubefirst cluster id (required)")
	transmitCmd.MarkFlagRequired("cluster-id")
	transmitCmd.Flags().StringVar(&clusterType, "cluster-type", "", "the kubefirst cluster type (required)")
	transmitCmd.MarkFlagRequired("cluster-type")
	transmitCmd.Flags().StringVar(&installMethod, "install-method", "", "the installation method for the cluster")
	transmitCmd.Flags().StringVar(&kubefirstTeam, "kubefirst-team", "", "kubefirst team [true/false]")
	transmitCmd.Flags().StringVar(&kubefirstTeamInfo, "kubefirst-team-info", "", "kubefirst team info")
}
