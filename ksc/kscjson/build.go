package kscjson

import (
	"encoding/json"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/private/protocol/query/queryutil"
	"github.com/KscSDK/ksc-sdk-go/ksc/kscbody"
	"net/url"
	"reflect"
	"strings"
)

var BuildHandler = request.NamedHandler{Name: "kscsdk.kscjson.Build", Fn: Build}

func Build(r *request.Request) {
	method := strings.ToUpper(r.HTTPRequest.Method)
	if v := r.HTTPRequest.Header.Get("Content-Type"); len(v) == 0 {
		r.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")
	}

	if v := r.HTTPRequest.Header.Get("Content-Type"); !strings.Contains(strings.ToLower(v), "application/json") || method == "GET" {
		kscbody.BodyParam(r)
		return
	}
	body := url.Values{
		"Action":  {r.Operation.Name},
		"Version": {r.ClientInfo.APIVersion},
	}

	if method == "GET" {
		if reflect.TypeOf(r.Params) == reflect.TypeOf(&map[string]interface{}{}) {
			m := *(r.Params).(*map[string]interface{})
			for k, v := range m {
				body.Add(k, v.(string))
			}
		} else if err := queryutil.Parse(body, r.Params, false); err != nil {
			r.Error = awserr.New("SerializationError", "failed encoding Query request", err)
			return
		}
	}
	r.HTTPRequest.URL.RawQuery = body.Encode()

	b, _ := json.Marshal(r.Params)
	str := string(b)
	r.SetStringBody(str)
	r.Params = nil
	r.HTTPRequest.Header.Set("Accept", "application/json")
}
