package main

import (
	"strconv"

	"github.com/spf13/cobra"
	"github.com/xaionaro-go/smartoffice/client/pkg/smartoffice"
)

var rootCmd = &cobra.Command{
	Use: "smartoffice",
}

var setPinValues = &cobra.Command{
	Use:  "set-pin-values",
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		values, err := strconv.ParseUint(args[1], 10, 64)
		if err != nil {
			panic(err)
		}
		smartoffice.New(args[0]).SetPinValues(values)
	},
}

var setPinValue = &cobra.Command{
	Use:  "set-pin-value",
	Args: cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		_pinID, err := strconv.ParseUint(args[1], 10, 8)
		if err != nil {
			panic(err)
		}
		pinID := uint8(_pinID)
		value, err := strconv.ParseBool(args[2])
		if err != nil {
			panic(err)
		}
		smartoffice.New(args[0]).SetPinValue(pinID, value)
	},
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
	rootCmd.AddCommand(setPinValue)
	rootCmd.AddCommand(setPinValues)
	rootCmd.AddCommand(enableLightByRadio)
	rootCmd.AddCommand(disableLightByRadio)
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}
