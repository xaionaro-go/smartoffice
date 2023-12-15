package smartoffice

import (
	"time"
)

var (
	delayChangeBrightness = time.Millisecond * 10
	delayGroupSwitch      = time.Millisecond * 100
)

func newDisableLightByRadioCommands() []byte {
	var cmds Commands
	for group := 0; group < 6; group++ {
		cmds.SetPin(3, false)
		cmds.Delay(delayChangeBrightness)
		for value := 100; value > 0; value-- {
			// decrease the brightness by 1:
			cmds.SetPin(2, false)
			cmds.Delay(delayChangeBrightness)
			cmds.SetPin(2, true)
			cmds.Delay(delayChangeBrightness)
		}
		cmds.SetPin(3, true)
		cmds.Delay(delayChangeBrightness)

		// select the next group:
		cmds.SetPin(1, false)
		cmds.Delay(delayGroupSwitch)
		cmds.SetPin(1, true)
		cmds.Delay(delayGroupSwitch)
	}

	return cmds.Bytes()
}

func newEnableLightByRadioCommands() []byte {
	var cmds Commands
	for group := 0; group < 6; group++ {
		for value := 0; value < 100; value++ {
			// increase the brightness by 1:
			cmds.SetPin(3, false)
			cmds.Delay(delayChangeBrightness)
			cmds.SetPin(3, true)
			cmds.Delay(delayChangeBrightness)
		}

		// select the next group:
		cmds.SetPin(1, false)
		cmds.Delay(delayGroupSwitch)
		cmds.SetPin(1, true)
		cmds.Delay(delayGroupSwitch)
	}

	return cmds.Bytes()
}
