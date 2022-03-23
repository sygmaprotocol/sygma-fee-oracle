package client

import (
	"context"
	"io"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
)

const (
	requestTimeOut = time.Minute
)

type HttpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type HttpRequestMessage struct {
	log         *logrus.Entry
	client      HttpClient
	url         string
	header      map[string]string
	requestBody io.Reader
	method      string
}

func NewHttpRequestMessage(method, url string, header map[string]string, requestBody io.Reader, log *logrus.Entry) *HttpRequestMessage {
	return &HttpRequestMessage{
		log:         log.WithField("http client", "client"),
		client:      &http.Client{},
		url:         url,
		header:      header,
		requestBody: requestBody,
		method:      method,
	}
}

func (h *HttpRequestMessage) Request() (int, []byte, error) {
	req, err := http.NewRequest(h.method, h.url, h.requestBody)
	if err != nil {
		h.log.Errorf("create new %s http request err: %s", h.method, err.Error())
		return -1, nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	if h.header != nil {
		for key, value := range h.header {
			req.Header.Set(key, value)
		}
	}

	ctx, cancel := context.WithTimeout(context.Background(), requestTimeOut)
	defer cancel()

	req = req.WithContext(ctx)

	resp, err := h.client.Do(req)
	if err != nil {
		h.log.Error("send http request err ", err)
		return -1, nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		h.log.Errorf("parse http.%s response body err: %s(%s)", h.method, err.Error(), h.url)
		return -1, nil, err
	}
	return resp.StatusCode, body, nil
}
