package kscquery

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/request"
	"io/ioutil"
)

type xmlErrorResponse struct {
	Error Error `json:"Error"`
	RequestID string   `json:"RequestID"`
}

type Error struct {
	Code string `json:"Code"`
	Message string `json:"Message"`
}

// UnmarshalErrorHandler is a name request handler to unmarshal request errors
var UnmarshalErrorHandler = request.NamedHandler{Name: "kscsdk.query.UnmarshalError", Fn: UnmarshalError}

// UnmarshalError unmarshals an error response for an AWS Query service.
func UnmarshalError(r *request.Request) {
	defer r.HTTPResponse.Body.Close()
	resp := xmlErrorResponse{}
	if r.DataFilled() {
		body, err := ioutil.ReadAll(r.HTTPResponse.Body)
		if err != nil {
			fmt.Printf("read body err, %v\n", err)
			return
		}
		if err = json.Unmarshal(body,&resp); err != nil {
			fmt.Printf("Unmarshal err, %v\n", err)
			return
		}
		r.Error = awserr.NewRequestFailure(
			awserr.New(resp.Error.Code, resp.Error.Message, nil),
			r.HTTPResponse.StatusCode,
			resp.RequestID,
		)
		return
	}else {
		r.Error = awserr.NewRequestFailure(
			awserr.New("ServiceUnavailableException", "service is unavailable", nil),
			r.HTTPResponse.StatusCode,
			r.RequestID,
		)
		return
	}
}
