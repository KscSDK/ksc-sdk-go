package ksc

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
)

func NewClient(ak, sk string, debug ...bool) *session.Session {
	flag := false
	if len(debug) > 0 && len(debug) == 1 {
		flag = debug[0]
	}
	_credentials := credentials.NewStaticCredentials(ak, sk, "")
	if flag {
		sess, err := session.NewSession(&aws.Config{Credentials: _credentials,
			LogLevel: aws.LogLevel(aws.LogDebugWithHTTPBody)})
		if err != nil {
			fmt.Printf("%v", err)
		}
		return sess
	} else {
		sess, err := session.NewSession(&aws.Config{Credentials: _credentials})
		if err != nil {
			fmt.Printf("%v", err)
		}
		return sess
	}

}
