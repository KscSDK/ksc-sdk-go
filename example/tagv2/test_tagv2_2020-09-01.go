package main

import (
	"encoding/json"
	"fmt"
	"github.com/KscSDK/ksc-sdk-go/ksc"
	"github.com/KscSDK/ksc-sdk-go/ksc/utils"
	"github.com/KscSDK/ksc-sdk-go/service/tagv2"
)

func main() {
	ak := "您的AK"
	sk := "您的SK"
	region := "cn-beijing-6"
	svc := tagv2.SdkNew(ksc.NewClient(ak, sk), &ksc.Config{Region: &region}, &utils.UrlInfo{
		UseSSL: true,
	})

	var resp *map[string]interface{}
	var err error
	//	****************************************创建Tag(CreateTag())***************************************************
	createTag := make(map[string]interface{})
	createTag["Key"] = "xym-test"
	createTag["Value"] = "123"

	resp, err = svc.CreateTag(&createTag)
	if err != nil {
		fmt.Println("error:", err.Error())
	}
	if resp != nil {
		str, _ := json.Marshal(&resp)
		fmt.Printf("%+v\n", string(str))
	}
	//	****************************************查询tag(CreateTag())***************************************************
}
