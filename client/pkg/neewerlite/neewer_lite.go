package neewerlite

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"net/url"
	"time"
)

type NeewerLite struct {
	url *url.URL
}

func NewNeewerLite(url *url.URL) *NeewerLite {
	return &NeewerLite{
		url: url,
	}
}

func (t NeewerLite) Bri(light net.HardwareAddr, brightness uint8) {
	t.sendAction(map[string]any{"light": light.String(), "bri": brightness})
}

func (t NeewerLite) getURL(args map[string]any) *url.URL {
	urlString := t.url.String()
	urlCopy, err := url.Parse(urlString)
	if err != nil {
		panic(err)
	}
	urlCopy.Path += "NeewerLite-Python/doAction"
	values := url.Values{
		"nopage": []string{"true"},
	}
	for k, v := range args {
		values[k] = []string{fmt.Sprint(v)}
	}
	urlCopy.RawQuery = values.Encode()
	return urlCopy
}

func (t NeewerLite) sendAction(args map[string]any) {
	req := &http.Request{
		Method: "GET",
		URL:    t.getURL(args),
	}
	log.Printf("sending GET to <%s>\n", req.URL.String())

	resp, err := (&http.Client{
		Timeout: 60 * time.Second,
	}).Do(req)
	if err != nil {
		time.Sleep(2 * time.Second)
		resp, err = (&http.Client{
			Timeout: 60 * time.Second,
		}).Do(req)
	}
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	ioutil.ReadAll(resp.Body)

	log.Printf("%s\n", resp.Body)
	time.Sleep(5 * time.Second)
}
