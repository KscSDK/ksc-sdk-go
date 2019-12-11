package kscjson

import (
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/KscSDK/ksc-sdk-go/ksc/kscquery"
)

// UnmarshalHandler is a named request handler for unmarshaling query protocol requests
var UnmarshalHandler = request.NamedHandler{Name: "kscsdk.query.Unmarshal", Fn: Unmarshal}

// UnmarshalMetaHandler is a named request handler for unmarshaling query protocol request metadata
var UnmarshalMetaHandler = request.NamedHandler{Name: "kscsdk.query.UnmarshalMeta", Fn: UnmarshalMeta}

// Unmarshal unmarshals a response for an AWS Query service.
func Unmarshal(r *request.Request) {
	kscquery.Unmarshal(r)
}

// UnmarshalMeta unmarshals header response values for an AWS Query service.
func UnmarshalMeta(r *request.Request) {
	//r.RequestID = r.HTTPResponse.Header.Get("X-Amzn-Requestid")
	kscquery.UnmarshalMeta(r)
}
