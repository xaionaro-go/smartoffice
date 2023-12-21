package main

import (
	"net/http"

	"github.com/spf13/cobra"
	"github.com/mcuadros/go-octoprint"
	"github.com/xaionaro-go/smartoffice/client/pkg/smartoffice"
)

func sendSuccessAndClose(res http.ResponseWriter) {
	res.WriteHeader(200)
	res.Write([]byte(`Done, now close the window :)`))
}

type smartOfficeEnable []string

func (args smartOfficeEnable) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	go smartoffice.New(args[0]).EnableLightByRadio()
	go smartoffice.New(args[1]).IRSend(smartoffice.IRCONTROL_TYPE_NEC, 0, 0xF740BF, 32, 10, 0)
	client := octoprint.NewClient(args[2], args[3])
	go (&octoprint.PauseRequest{
		Action: octoprint.Pause,
	}).Do(client)
	sendSuccessAndClose(res)
}

type smartOfficeDisable []string

func (args smartOfficeDisable) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	go smartoffice.New(args[0]).DisableLightByRadio()
	go smartoffice.New(args[1]).IRSend(smartoffice.IRCONTROL_TYPE_NEC, 0, 0xF7C03F, 32, 10, 0)
	client := octoprint.NewClient(args[2], args[3])
	go (&octoprint.PauseRequest{
		Action: octoprint.Resume,
	}).Do(client)
	sendSuccessAndClose(res)
}

var rootCmd = &cobra.Command{
	Use:  "smartofficed",
	Args: cobra.ExactArgs(4),
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
