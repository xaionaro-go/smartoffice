package main

import (
	"net/http"

	"github.com/spf13/cobra"
	"github.com/xaionaro-go/smartoffice/client/pkg/smartoffice"
)

type smartOfficeEnable []string

func (args smartOfficeEnable) ServeHTTP(http.ResponseWriter, *http.Request) {
	go smartoffice.New(args[0]).EnableLightByRadio()
	go smartoffice.New(args[1]).IRSend(smartoffice.IRCONTROL_TYPE_NEC, 0, 0xF7C03F, 32, 10, 0)
}

type smartOfficeDisable []string

func (args smartOfficeDisable) ServeHTTP(http.ResponseWriter, *http.Request) {
	go smartoffice.New(args[0]).DisableLightByRadio()
	go smartoffice.New(args[1]).IRSend(smartoffice.IRCONTROL_TYPE_NEC, 0, 0xF740BF, 32, 10, 0)
}

var rootCmd = &cobra.Command{
	Use:  "smartofficed",
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		mux := http.NewServeMux()
		mux.Handle("/enable/", smartOfficeEnable(args))
		mux.Handle("/disable/", smartOfficeDisable(args))

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
