package client

import (
	"net/http"
	"testing"
)

func TestGetHttpClient(t *testing.T) {
	client := GetHttpClient(false)
	if client == nil {
		t.Error("GetHttpClient should always return a http.Client")
	}
}

func TestGetHttpClient_SkipVerify(t *testing.T) {
	client := GetHttpClient(true)
	tr, ok := client.Transport.(*http.Transport)
	if !ok {
		t.Error("Transport is not an http.Transport")
	}
	if !tr.TLSClientConfig.InsecureSkipVerify {
		t.Error("Expected InsecureSkipVerify to be true")
	}
}

func TestGetHttpClient_NotSkipVerify(t *testing.T) {
	client := GetHttpClient(false)
	tr, ok := client.Transport.(*http.Transport)
	if !ok {
		t.Error("Transport is not an http.Transport")
	}
	if tr.TLSClientConfig.InsecureSkipVerify {
		t.Error("Expected InsecureSkipVerify to be false")
	}
}
