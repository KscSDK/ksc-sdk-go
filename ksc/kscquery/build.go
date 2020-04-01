package kscquery

import (
	"github.com/KscSDK/ksc-sdk-go/ksc/kscbody"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/private/protocol/query/queryutil"
	"net/url"
	"reflect"
	"strings"
)

// BuildHandler is a named request handler for building query protocol requests
var BuildHandler = request.NamedHandler{Name: "kscsdk.query.Build", Fn: Build}

// Build builds a request for an AWS Query service.
func Build(r *request.Request) {
	body := url.Values{
		"Action":  {r.Operation.Name},
		"Version": {r.ClientInfo.APIVersion},
	}
	r.HTTPRequest.Header.Add("Accept", "application/json")
	//method := strings.ToUpper(r.HTTPRequest.Method)
	r.HTTPRequest.URL.RawQuery = body.Encode()
	v := r.HTTPRequest.Header.Get("Content-Type")
	if len(v) > 0 && strings.Contains(strings.ToLower(v), "application/json") {
		r.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")
		kscbody.BodyJson(r)
		return
	}
	if reflect.TypeOf(r.Params) == reflect.TypeOf(&map[string]interface{}{}) {
		m := *(r.Params).(*map[string]interface{})
		for k, v := range m {
			if reflect.TypeOf(v).String() == "string" {
				body.Add(k, v.(string))
			}
		}
	} else if err := queryutil.Parse(body, r.Params, false); err != nil {
		r.Error = awserr.New("SerializationError", "failed encoding Query request", err)
		return
	}
	if len(v) > 0 && strings.Contains(strings.ToLower(v), "x-www-form-urlencoded") {
		r.HTTPRequest.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=utf-8")
		r.SetBufferBody([]byte(body.Encode()))
		return
	}
	if len(v) > 0 {
		r.Error = awserr.New("SerializationError", "not support such content-type", nil)
		return
	}

	method := strings.ToUpper(r.HTTPRequest.Method)
	if method == "GET" || method == "DELETE" {
		r.HTTPRequest.URL.RawQuery = body.Encode()
	} else if method == "POST" || method == "PUT" {
		r.HTTPRequest.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=utf-8")
		r.SetBufferBody([]byte(body.Encode()))
	} else {
		r.HTTPRequest.URL.RawQuery = body.Encode()
	}
}
