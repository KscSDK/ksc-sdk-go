package main

import (
	"encoding/json"
	"fmt"
	"github.com/ksc/ksc-sdk-go/ksc"
	"github.com/ksc/ksc-sdk-go/ksc/utils"
	"github.com/ksc/ksc-sdk-go/service/iam"
)

func main() {
	/*
	ak := "你的ak"
	sk := "你的sk"
	*/
	ak := "AKLT2mjX47yFSp2leXlY8h0dVA"
	sk := "ODFsox160Mo4Qz3PX5zyYmIbmxFx4qhfbRBxtBH7urMolwxLxUdUoN+oPt5lRaMt8w=="
	region := "cn-shanghai-3"

	//debug模式的话 打开这个开关
	svc := iam.SdkNew(ksc.NewClient(ak, sk /*,true*/), &ksc.Config{Region: &region}, &utils.UrlInfo{
		UseSSL: false,
	})

	var resp *map[string]interface{}
	var err error

	//	****************************************查询策略信息列表（ListPolicies）***************************************************
	/*
	resp, err = svc.ListPolicies(nil)
	if err != nil {
		fmt.Println("error:", err.Error())
	}
	if resp != nil {
		str, _ := json.Marshal(&resp)
		fmt.Printf("%+v\n", string(str))
	}
	*/

	//	****************************************查询用户基本信息（GetUser）***************************************************
	params := make(map[string]interface{})
	params["UserName"] = "zechao1"

	resp, err = svc.GetUser(&params)
	if err != nil {
		fmt.Println("error:", err.Error())
	}
	if resp != nil {
		str, _ := json.Marshal(&resp)
		fmt.Printf("%+v\n", string(str))
	}

}
