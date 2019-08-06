package main

import (
	"encoding/json"
	"fmt"
	"github.com/ksc/ksc-sdk-go/ksc"
	"github.com/ksc/ksc-sdk-go/ksc/utils"
	"github.com/ksc/ksc-sdk-go/service/iam"
)

func main() {
	ak := "AKLT0tdvdrASSnGbxQbHSdeAbA"
	sk := "ODi8E02pdZSO1pGyAWh8+OHF6Uh9LIxv9xhYVtdgsN9IHetRJJvPWttQPxvIpJi+7Q=="
	region := "cn-beijing-6"

	/*
	ak := "AKLTWjcmVntESASmyr9QAGELMg"
	sk := "OKn4wzC1ncrDq5tYBFZWeESIuh5i1twyZm5+/lfNRdMr6ZlPwHNy0Eu3YUL+wACSIA=="
	region := "cn-shanghai-3"
	*/

	//debug模式的话 打开这个开关
	svc := iam.SdkNew(ksc.NewClient(ak, sk ,false), &ksc.Config{Region: &region}, &utils.UrlInfo{
		UseSSL: true,
	})

	var resp *map[string]interface{}
	var err error

	//	****************************************查询策略信息列表（ListPolicies）***************************************************
	resp, err = svc.ListPolicies(nil)
	if err != nil {
		fmt.Println("error:", err.Error())
	}
	if resp != nil {
		str, _ := json.Marshal(&resp)
		fmt.Printf("%+v\n", string(str))
	}

	//	****************************************查询用户基本信息（GetUser）***************************************************
	/*
	params := make(map[string]interface{})
	params["UserName"] = "one"

	resp, err = svc.GetUser(&params)
	if err != nil {
		fmt.Println("error:", err.Error())
	}
	if resp != nil {
		str, _ := json.Marshal(&resp)
		fmt.Printf("%+v\n", string(str))
	}
	*/

}
