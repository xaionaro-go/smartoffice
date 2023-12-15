package smartoffice

import (
	"bytes"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

type SmartOffice struct {
	url *url.URL
}

func New(urlString string) *SmartOffice {
	_url, err := url.Parse(urlString)
	if err != nil {
		panic(err)
	}
	return &SmartOffice{
		url: _url,
	}
}

func (office *SmartOffice) getURL(uri string) *url.URL {
	urlString := office.url.String()
	urlCopy, err := url.Parse(urlString)
	if err != nil {
		panic(err)
	}
	urlCopy.Path += uri
	return urlCopy
}

func (office *SmartOffice) EnableLightByRadio() {
	office.sendPOST("execute/", "text/plain", newEnableLightByRadioCommands())
}

func (office *SmartOffice) DisableLightByRadio() {
	office.sendPOST("execute/", "text/plain", newDisableLightByRadioCommands())
}

func (office *SmartOffice) sendPOST(path string, contentType string, data []byte) {
	req := &http.Request{
		Method:     "POST",
		URL:        office.getURL(path),
		Proto:      "HTTP/1.0",
		ProtoMajor: 1,
		ProtoMinor: 0,
		Header: map[string][]string{
			"Content-Type": {contentType},
		},
		Body:             &BufferCloser{Buffer: bytes.NewBuffer(data)},
		TransferEncoding: []string{"identity"},
		ContentLength:    int64(len(data)),
	}

	resp, err := (&http.Client{
		Timeout: 30 * time.Second,
	}).Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Printf("%s\n", resp.Body)
}
