package kscquery

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go/aws/request"
	"io/ioutil"
)

// UnmarshalHandler is a named request handler for unmarshaling query protocol requests
var UnmarshalHandler = request.NamedHandler{Name: "kscsdk.query.Unmarshal", Fn: Unmarshal}

// UnmarshalMetaHandler is a named request handler for unmarshaling query protocol request metadata
var UnmarshalMetaHandler = request.NamedHandler{Name: "kscsdk.query.UnmarshalMeta", Fn: UnmarshalMeta}

// Unmarshal unmarshals a response for an AWS Query service.
func Unmarshal(r *request.Request) {
	defer r.HTTPResponse.Body.Close()
	if r.DataFilled() {
		body, err := ioutil.ReadAll(r.HTTPResponse.Body)
		if err != nil {
			fmt.Printf("read body err, %v\n", err)
			return
		}
		if err = json.Unmarshal(body, &r.Data); err != nil {
			fmt.Printf("Unmarshal err, %v\n", err)
			return
		}
	}
}

// UnmarshalMeta unmarshals header response values for an AWS Query service.
func UnmarshalMeta(r *request.Request) {
	//r.RequestID = r.HTTPResponse.Header.Get("X-Amzn-Requestid")
}
