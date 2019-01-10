package cmd

import (
	"flag"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/wlan0/simple-ci/pkg/ci"

	_ "github.com/golang/glog"
)

var RootCmd = &cobra.Command{
	Use:   "simple-ci",
	Short: "easy to setup, highly scalable ci system",
	RunE: func(c *cobra.Command, args []string) error {
		return ci.Run()
	},
	SilenceErrors: true,
	SilenceUsage:  true,
}

func init() {
	viper.AutomaticEnv()
	replacer := strings.NewReplacer("-", "_")
	viper.SetEnvKeyReplacer(replacer)

	RootCmd.Flags().StringSlice("store", []string{}, "list of etcd nodes")
	RootCmd.Flags().String("id", "", "unique id of the simple-ci cluster")
	RootCmd.Flags().String("ip", "", "ip address of the simple-ci server")
	RootCmd.Flags().Int("port", 8888, "port at which simple-ci server listens")

	RootCmd.Flags().String("github-endpoint", "", "callback endpoint of the simple-ci server")
	RootCmd.Flags().String("github-client-id", "", "github client id of the simple-ci server")
	RootCmd.Flags().String("github-secret", "", "github secret of the simple-ci server")
	RootCmd.Flags().String("webhook-secret", "", "webhook secret of the simple-ci github application")

	// parse the go default flagset to get flags for glog and other packages in future
	RootCmd.Flags().AddGoFlagSet(flag.CommandLine)

	// defaulting this to true so that logs are printed to console
	flag.Set("logtostderr", "true")

	//suppress the incorrect prefix in glog output
	flag.CommandLine.Parse([]string{})

	viper.BindPFlags(RootCmd.Flags())
}