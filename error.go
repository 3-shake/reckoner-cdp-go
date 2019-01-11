package reckonercdp

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	UnauthorizedError = "UnauthorizedError"
	ServerError       = "ServerError"
	ClientError       = "ClientError"
)

type ReckonerCdpError struct {
	Type     string
	Message  string
	Response *http.Response
}

func (e *ReckonerCdpError) Error() string {
	return fmt.Sprintf("%s: %s)", e.Type, e.Message)
}

func (client *Client) createError(res *http.Response, message string) error {
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	message = fmt.Sprintf("%s: %s", message, body)

	errorType := ""
	if res.StatusCode == http.StatusUnauthorized {
		errorType = UnauthorizedError
	} else if res.StatusCode/100 == 4 {
		errorType = ClientError
	} else if res.StatusCode/100 == 5 {
		errorType = ServerError
	}

	return &ReckonerCdpError{
		Type:     errorType,
		Message:  message,
		Response: res,
	}
}
