package client

import (
	"crypto/tls"
	"net/http"
	"time"
)

func GetHttpClient(insecureSkipVerify bool) *http.Client {
	return &http.Client{
		Timeout: time.Second * 10,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: insecureSkipVerify,
			},
		},
	}

}
