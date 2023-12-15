package smartoffice

import (
	"bytes"
	"strconv"
	"time"
)

type Commands struct {
	bytes.Buffer
}

func (c *Commands) SetValue(value uint64) {
	(*bytes.Buffer)(&c.Buffer).WriteString("S" + strconv.FormatUint(value, 36) + "\n")
}

func (c *Commands) SetPin(pin uint8, value bool) {
	valueStr := "0"
	if value {
		valueStr = "1"
	}
	(*bytes.Buffer)(&c.Buffer).WriteString("s" + strconv.FormatUint(uint64(pin), 36) + valueStr + "\n")
}

func (c *Commands) Delay(delay time.Duration) {
	(*bytes.Buffer)(&c.Buffer).WriteString("d" + strconv.FormatUint(uint64(delay.Milliseconds()), 36) + "\n")
}

func (c *Commands) Close() error {
	return nil
}
