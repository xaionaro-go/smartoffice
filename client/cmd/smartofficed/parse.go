package main

import (
	"net"
	"net/url"
)

func assertNoError(err error) {
	if err != nil {
		panic(err)
	}
}

func parseURL(in string) *url.URL {
	r, err := url.Parse(in)
	assertNoError(err)
	return r
}

func parseMAC(in string) net.HardwareAddr {
	r, err := net.ParseMAC(in)
	assertNoError(err)
	return r
}
