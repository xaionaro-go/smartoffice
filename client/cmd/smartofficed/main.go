package main

import (
	"net/http"

	"github.com/spf13/cobra"
	"github.com/xaionaro-go/smartoffice/client/pkg/smartoffice"
)

type smartOfficeEnable string

func (arg smartOfficeEnable) ServeHTTP(http.ResponseWriter, *http.Request) {
	smartoffice.New(string(arg)).EnableLightByRadio()
}

type smartOfficeDisable string

func (arg smartOfficeDisable) ServeHTTP(http.ResponseWriter, *http.Request) {
	smartoffice.New(string(arg)).DisableLightByRadio()
}

var rootCmd = &cobra.Command{
	Use:  "smartofficed",
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		mux := http.NewServeMux()
		mux.Handle("/enable/", smartOfficeEnable(args[0]))
		mux.Handle("/disable/", smartOfficeDisable(args[0]))

		s := &http.Server{
			Addr:    "127.0.0.1:46388",
			Handler: mux,
		}

		err := s.ListenAndServe()
		if err != nil {
			panic(err)
		}
	},
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}
