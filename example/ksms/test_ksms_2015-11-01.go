package main

import (
	"encoding/json"
	"fmt"
	"github.com/KscSDK/ksc-sdk-go/ksc"
	"github.com/KscSDK/ksc-sdk-go/ksc/utils"
	"github.com/KscSDK/ksc-sdk-go/service/ksms"
)

func main() {
	ak := "你的ak"
	sk := "你的sk"
	region := "cn-shanghai-3"

	//debug模式的话 打开这个开关
	svc := ksms.SdkNew(ksc.NewClient(ak, sk ,false), &ksc.Config{Region: &region}, &utils.UrlInfo{
		UseSSL: false,
	})

	var resp *map[string]interface{}
	var err error

	//	****************************************发送单条短信（SendSms）***************************************************
	params := make(map[string]interface{})
	params["Mobile"] = "130xx"
	params["SignName"] = "xxxxx"
	params["TplId"] = "7"
	params["TplParams"] = "{\"param1\":\"test\"}"

	resp, err = svc.SendSms(&params)
	if err != nil {
		fmt.Println("error:", err.Error())
	}
	if resp != nil {
		str, _ := json.Marshal(&resp)
		fmt.Printf("%+v\n", string(str))
	}

	//	****************************************发送多条短信（BatchSendSms）***************************************************
	/*
	params := make(map[string]interface{})
	params["Mobile"] = "130xxxx,153xxxxx"
	params["SignName"] = "xxxxx"
	params["TplId"] = "7"
	params["TplParams"] = "{\"param1\":\"test\"}"

	resp, err = svc.BatchSendSms(&params)
	if err != nil {
		fmt.Println("error:", err.Error())
	}
	if resp != nil {
		str, _ := json.Marshal(&resp)
		fmt.Printf("%+v\n", string(str))
	}
	*/
}
