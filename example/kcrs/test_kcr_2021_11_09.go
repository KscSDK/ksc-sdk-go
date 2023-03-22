package main

import (
	"encoding/json"
	"fmt"
	"github.com/KscSDK/ksc-sdk-go/ksc"
	"github.com/KscSDK/ksc-sdk-go/ksc/utils"
	"github.com/KscSDK/ksc-sdk-go/service/kcrs"
)

func main() {
	ak := "用户AK"
	sk := "用户SK"
	region := "cn-beijing-6"
	//debug模式的话 打开这个开关
	svc := kcrs.SdkNew(ksc.NewClient(ak, sk /*,true*/), &ksc.Config{Region: &region}, &utils.UrlInfo{
		UseSSL:      true,
		UseInternal: false,
	})
	var resp *map[string]interface{}
	var err error

	//	****************************************创建镜像仓库实例(CreateInstance)***************************************************
	createInstance := map[string]interface{}{
		"InstanceName": "xxxxx",
		"ChargeType":   "HourlyInstantSettlement",
		"InstanceType": "basic",
		"PurchaseTime": "1",
		"ProjectId":    "0",
	}
	resp, err = svc.CreateInstance(&createInstance)
	if err != nil {
		fmt.Println("error:", err.Error())
	}
	if resp != nil {
		str, _ := json.Marshal(&resp)
		fmt.Printf("%+v\n", string(str))
	}
	//	****************************************查询镜像仓库列表(DescribeInstance)***************************************************
	describeInstance := map[string]interface{}{
		"ProjectId.1": "0",
		"MaxResults":  20,
		"Marker":      0,
	}
	resp, err = svc.DescribeInstance(&describeInstance)
	if err != nil {
		fmt.Println("error:", err.Error())
	}
	if resp != nil {
		str, _ := json.Marshal(&resp)
		fmt.Printf("%+v\n", string(str))
	}

	//	****************************************删除镜像仓库(DeleteInstance)***************************************************
	deleteInstance := map[string]interface{}{
		"InstanceId": "c3fdd2ac-xxx",
	}
	resp, err = svc.DeleteInstance(&deleteInstance)
	if err != nil {
		fmt.Println("error:", err.Error())
	}
	if resp != nil {
		str, _ := json.Marshal(&resp)
		fmt.Printf("%+v\n", string(str))
	}
	//	****************************************创建实例访问凭证(CreateInstanceToken)***************************************************
	createInstanceToken := map[string]interface{}{
		"InstanceId": "c3fdd2ac-xxxx",
		"TokenType":  "Hour",
		"TokenTime":  "24",
	}
	resp, err = svc.CreateInstanceToken(&createInstanceToken)
	if err != nil {
		fmt.Println("error:", err.Error())
	}
	if resp != nil {
		str, _ := json.Marshal(&resp)
		fmt.Printf("%+v\n", string(str))
	}
}
