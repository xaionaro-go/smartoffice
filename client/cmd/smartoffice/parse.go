package main

import (
	"strconv"
	"strings"
)

func smartParseUint(in string, nBits int) (uint64, error) {
	switch {
	case strings.HasPrefix(in, "0x"):
		return strconv.ParseUint(in[2:], 16, nBits)
	}
	return strconv.ParseUint(in, 10, nBits)
}
