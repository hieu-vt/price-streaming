package cmd

import (
	"fmt"
	goservice "github.com/200Lab-Education/go-sdk"
	"github.com/spf13/cobra"
	"os"
	"price-streaming/cmd/streaming"
	"price-streaming/common"
	"price-streaming/plugin/nats/pubsub"
)

func newService() goservice.Service {
	service := goservice.New(
		goservice.WithName("food-delivery-cart-categories"),
		goservice.WithVersion("1.0.0"),
		goservice.WithInitRunnable(pubsub.NewNatsPubSub(common.PluginNatsPubsub)),
	)

	return service
}

var rootCmd = &cobra.Command{
	Use:   "app",
	Short: "Start an food delivery service",
	Run: func(cmd *cobra.Command, args []string) {
		service := newService()
		serviceLogger := service.Logger("service")

		if err := service.Init(); err != nil {
			serviceLogger.Fatalln(err)
		}

		err := streaming.StreamingPrice(service)

		if err != nil {
			serviceLogger.Fatalln(err)
		}

		serviceLogger.Infof("Start streaming ...")

		if err := service.Start(); err != nil {
			serviceLogger.Fatalln(err)
		}
	},
}

func Execute() {
	rootCmd.AddCommand(outEnvCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)

		os.Exit(1)
	}
}
