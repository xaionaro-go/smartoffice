package smartoffice

import (
	"time"
)

var (
	delayChangeBrightness = time.Millisecond * 10
	delayGroupSwitch      = time.Millisecond * 200
	brightnessClickStep   = 1.9
)

func newNoise() []byte {
	var cmds Commands
	for value := 0; float64(value) < float64(100)/float64(brightnessClickStep); value++ {
		cmds.SetPin(3, false)
		cmds.Delay(delayChangeBrightness)
		cmds.SetPin(2, false)
		cmds.Delay(delayChangeBrightness)
		cmds.SetPin(2, true)
		cmds.Delay(delayChangeBrightness)
		cmds.SetPin(3, true)
		cmds.Delay(delayChangeBrightness)
		cmds.SetPin(2, false)
		cmds.Delay(delayChangeBrightness)
		cmds.SetPin(2, true)
		cmds.Delay(delayChangeBrightness)
	}
	return cmds.Bytes()
}

func newDisableLightByRadioCommands() []byte {
	var cmds Commands
	cmds.SetPin(4, true)
	for group := 0; group < 6; group++ {
		cmds.SetPin(3, false)
		cmds.Delay(delayChangeBrightness)
		for value := 0; float64(value) < float64(100)/float64(brightnessClickStep); value++ {
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

	// re-iterate all groups to make sure no settings were lost:
	for group := 0; group < 6; group++ {
		// select the next group:
		cmds.SetPin(1, false)
		cmds.Delay(delayGroupSwitch)
		cmds.SetPin(1, true)
		cmds.Delay(delayGroupSwitch)
	}

	cmds.SetPin(4, false)
	return cmds.Bytes()
}

func newEnableLightByRadioCommands() []byte {
	var cmds Commands
	cmds.SetPin(4, true)
	for group := 0; group < 6; group++ {
		for value := 0; float64(value) < float64(100)/float64(brightnessClickStep); value++ {
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

	// re-iterate all groups to make sure no settings were lost:
	for group := 0; group < 6; group++ {
		// select the next group:
		cmds.SetPin(1, false)
		cmds.Delay(delayGroupSwitch)
		cmds.SetPin(1, true)
		cmds.Delay(delayGroupSwitch)
	}

	cmds.SetPin(4, false)
	return cmds.Bytes()
}
