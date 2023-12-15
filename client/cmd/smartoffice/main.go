package main

import (
	"github.com/spf13/cobra"
	"github.com/xaionaro-go/smartoffice/client/pkg/smartoffice"
)

var rootCmd = &cobra.Command{
	Use: "smartoffice",
}

var enableLightByRadio = &cobra.Command{
	Use:  "enable-light-by-radio",
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		smartoffice.New(args[0]).EnableLightByRadio()
	},
}

var disableLightByRadio = &cobra.Command{
	Use:  "disable-light-by-radio",
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		smartoffice.New(args[0]).DisableLightByRadio()
	},
}

func main() {
	rootCmd.AddCommand(enableLightByRadio)
	rootCmd.AddCommand(disableLightByRadio)
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}
