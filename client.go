package reckonercdp

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"net/http"
	"net/url"
	"time"
)

const (
	defaultHost            = "localhost:8888"
	defaultStreamingHost   = "localhost:8080"
	defaultScheme          = "http"
	defaultStreamingScheme = "http"
)

type Client struct {
	AccessKeyID         string
	SecretAccessKey     string
	Host                string
	StreamingHost       string
	HttpScheme          string
	StreamingHttpScheme string
	HttpClient          *http.Client
}

type ClientSettings struct {
	AccessKeyID         string
	SecretAccessKey     string
	Host                string
	StreamingHost       string
	HttpScheme          string
	StreamingHttpScheme string
	Transport           http.RoundTripper
}

func (client *Client) signature(req *http.Request) string {
	stringToSign := req.Method + "\n" +
		req.Header.Get("Content-MD5") + "\n" +
		req.Header.Get("Content-Type") + "\n" +
		req.Header.Get("Date")
	hmacStr := hmac.New(sha256.New, []byte(client.SecretAccessKey))
	hmacStr.Write([]byte(stringToSign))
	hash := hmacStr.Sum(nil)
	base64Str := base64.StdEncoding.EncodeToString(hash)
	return base64Str
}

func (client *Client) get(path string, params url.Values) (*http.Response, error) {
	requestURL := (&url.URL{
		Scheme:   "http",
		Host:     client.Host,
		Path:     path,
		RawQuery: params.Encode(),
	}).String()

	req, err := http.NewRequest("GET", requestURL, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Date", time.Now().Format(time.RFC1123))
	authorizationHeader := "RECKONER-CDP " + client.AccessKeyID + ":" + client.signature(req)
	req.Header.Set("Authorization", authorizationHeader)

	res, err := client.HttpClient.Do(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (client *Client) post(path string, body interface{}) (*http.Response, error) {
	requestURL := (&url.URL{
		Scheme: "http",
		Host:   client.Host,
		Path:   path,
	}).String()

	jsonStr, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", requestURL, bytes.NewBuffer(jsonStr))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Date", time.Now().Format(time.RFC1123))
	authorizationHeader := "RECKONER-CDP " + client.AccessKeyID + ":" + client.signature(req)
	req.Header.Set("Authorization", authorizationHeader)

	res, err := client.HttpClient.Do(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (client *Client) streamingGet(path string, params url.Values) (*http.Response, error) {
	requestURL := (&url.URL{
		Scheme:   "http",
		Host:     client.StreamingHost,
		Path:     path,
		RawQuery: params.Encode(),
	}).String()

	req, err := http.NewRequest("GET", requestURL, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Date", time.Now().Format(time.RFC1123))
	authorizationHeader := "RECKONER-CDP " + client.AccessKeyID + ":" + client.signature(req)
	req.Header.Set("Authorization", authorizationHeader)

	res, err := client.HttpClient.Do(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func NewClient(settings *ClientSettings) *Client {
	httpClient := &http.Client{
		Transport: settings.Transport,
	}

	host := settings.Host
	if host == "" {
		host = defaultHost
	}
	streamingHost := settings.StreamingHost
	if streamingHost == "" {
		streamingHost = defaultStreamingHost
	}

	httpScheme := settings.HttpScheme
	if httpScheme == "" {
		httpScheme = defaultScheme
	}
	streamingHttpScheme := settings.StreamingHttpScheme
	if streamingHttpScheme == "" {
		streamingHttpScheme = defaultStreamingScheme
	}

	return &Client{
		AccessKeyID:         settings.AccessKeyID,
		SecretAccessKey:     settings.SecretAccessKey,
		Host:                host,
		StreamingHost:       streamingHost,
		HttpScheme:          httpScheme,
		StreamingHttpScheme: streamingHttpScheme,
		HttpClient:          httpClient,
	}
}
