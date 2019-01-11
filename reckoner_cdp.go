package reckonercdp

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"net/http"
	"net/url"
	"time"
)

const (
	defaultStreamingHost = "http://localhost:8080"
)

type Client struct {
	accessKeyID     string
	secretAccessKey string
	streamingHost   string
	httpClient      *http.Client
}

type ClientSettings struct {
	accessKeyID     string
	secretAccessKey string
	streamingHost   string
	transport       http.RoundTripper
}

func (client *Client) signature(req *http.Request) string {
	stringToSign := req.Method + "\n" +
		req.Header.Get("Content-MD5") + "\n" +
		req.Header.Get("Content-Type") + "\n" +
		req.Header.Get("Date")
	hmacStr := hmac.New(sha256.New, []byte(client.secretAccessKey))
	hmacStr.Write([]byte(stringToSign))
	hash := hmacStr.Sum(nil)
	base64Str := base64.StdEncoding.EncodeToString(hash)
	return base64Str
}

func (client *Client) streamingGet(path string, params url.Values) (*http.Response, error) {
	requestURL := (&url.URL{
		Scheme:   "http",
		Host:     client.streamingHost,
		Path:     path,
		RawQuery: params.Encode(),
	}).String()

	req, err := http.NewRequest("GET", requestURL, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Date", time.Now().Format(time.RFC1123))
	authorizationHeader := "RECKONER-CDP " + client.accessKeyID + ":" + client.signature(req)
	req.Header.Set("Authorization", authorizationHeader)

	res, err := client.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func NewClient(settings ClientSettings) *Client {
	httpClient := &http.Client{
		Transport: settings.transport,
	}
	streamingHost := settings.streamingHost
	if streamingHost == "" {
		streamingHost = defaultStreamingHost
	}

	return &Client{
		accessKeyID:     settings.accessKeyID,
		secretAccessKey: settings.secretAccessKey,
		httpClient:      httpClient,
		streamingHost:   streamingHost,
	}
}
