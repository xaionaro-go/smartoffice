package neewerlite

import (
	"fmt"
	"log"
	"net"
	"os/exec"
	"strings"
)

type NeewerLiteCommand struct {
	cmd []string
}

func NewNeewerLiteCommand(cmd ...string) *NeewerLiteCommand {
	return &NeewerLiteCommand{
		cmd: cmd,
	}
}

func (t NeewerLiteCommand) Bri(light net.HardwareAddr, brightness uint8) {
	cmd := []string{"bluetoothctl", "remove", light.String()}
	fmt.Printf("command: <%s %s>\n", cmd[0], strings.Join(cmd[1:], " "))
	cmdResult := exec.Command(cmd[0], cmd[1:]...)
	output, err := cmdResult.CombinedOutput()
	log.Printf("%s\n", output)
	if err != nil {
		log.Println("unable to reset the connection to the device: ", err)
	}
	t.sendAction(map[string]any{"light": light.String(), "bri": brightness})
}

func (t NeewerLiteCommand) sendAction(values map[string]any) {
	args := []string{"--cli", "--force_instance"}
	args = append(args, t.cmd[1:]...)
	for k, v := range values {
		args = append(args, fmt.Sprintf("--%s=%v", k, v))
	}
	fmt.Printf("command: <%s %s>\n", t.cmd[0], strings.Join(args, " "))
	cmd := exec.Command(t.cmd[0], args...)
	output, err := cmd.CombinedOutput()
	log.Printf("%s\n", output)
	if err != nil {
		log.Println("unable to send to command: ", err)
		return
	}
}
