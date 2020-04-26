package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	urllib "net/url"
	"strings"
	"time"

	"github.com/gojektech/heimdall/httpclient"
)

// RESTCaller is a utility class that makes calling REST api easier
type RESTCaller struct {
	root     string
	respType string
	timeout  time.Duration
}

// RESTCallerBuilder is a builder class that helps creating RESTCaller
type RESTCallerBuilder struct {
	root     string
	timeout  string
	respType string
}

// BuildRESTCaller returns a newly created RESTCaller builder object that helps you
// configure the RESTCaller
func BuildRESTCaller(root string) *RESTCallerBuilder {
	return &RESTCallerBuilder{
		root:     root,
		respType: "json",
		timeout:  "1s",
	}
}

// Timeout sets timeout for the REST api call
func (b *RESTCallerBuilder) Timeout(t string) *RESTCallerBuilder {
	b.timeout = t
	return b
}

// RespType specifies in which type the server returns its response
// Currently supports "json"
func (b *RESTCallerBuilder) RespType(t string) *RESTCallerBuilder {
	b.respType = t
	return b
}

// Build actually creates RESTCaller instance
func (b *RESTCallerBuilder) Build() (*RESTCaller, error) {
	timeout, err := time.ParseDuration(b.timeout)
	if err != nil {
		return nil, err
	}

	availRespTypes := []string{
		"json",
	}

	respTypeAllowed := false
	for _, typeName := range availRespTypes {
		if b.respType == strings.ToLower(typeName) {
			respTypeAllowed = true
		}
	}

	if !respTypeAllowed {
		supported := strings.Join(availRespTypes, ", ")
		return nil, fmt.Errorf("response type (%s) is not supported ([%s] supported)", b.respType, supported)
	}

	return &RESTCaller{
		root:     b.root,
		timeout:  timeout,
		respType: b.respType,
	}, nil
}

// Get calls REST API through GET method
func (rest *RESTCaller) Get(url string, params map[string]string, headers *http.Header, dest interface{}) error {
	if headers == nil {
		headers = &http.Header{}
	}

	if rest.respType == "json" {
		headers.Set("Content-Type", "application/json")
	}

	var getParam = ""
	if params != nil {
		first := true
		for prmName, prmValue := range params {
			if !first {
				getParam += "&"
			} else {
				getParam += "?"
			}
			getParam += fmt.Sprintf("%s=%s", urllib.QueryEscape(prmName), urllib.QueryEscape(prmValue))
		}
	}
	url += getParam

	client := httpclient.NewClient(httpclient.WithHTTPTimeout(rest.timeout))
	resp, err := client.Get(fmt.Sprintf("%s%s", rest.root, url), *headers)
	if err != nil {
		log.Printf("[GET] %s (error: %s)", url, err.Error())
		return err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("[GET] %s (error: %s)", url, err.Error())
		return err
	}
	log.Printf("[GET] %s (%d bytes)", url, len(body))

	if rest.respType == "json" {
		err = json.Unmarshal(body, dest)
	}
	return err
}

// Post calls REST API through POST method
func (rest *RESTCaller) Post(url string, data []byte, headers *http.Header, dest interface{}) error {

	if headers == nil {
		headers = &http.Header{}
	}

	// rest.respType
	if rest.respType == "json" {
		headers.Set("Content-Type", "application/json")
	}

	reqReader := bytes.NewReader(data)

	client := httpclient.NewClient(httpclient.WithHTTPTimeout(rest.timeout))
	resp, err := client.Post(fmt.Sprintf("%s%s", rest.root, url), reqReader, *headers)
	if err != nil {
		log.Printf("[POST] %s (error: %s)", url, err.Error())
		return err
	}

	body, err := ioutil.ReadAll(resp.Body)
	log.Printf("[POST] %s (%d bytes)", url, len(body))

	if err != nil {
		log.Printf("[POST] %s (error: %s)", url, err.Error())
		return err
	}

	if rest.respType == "json" {
		err = json.Unmarshal(body, dest)
	}
	return err
}
