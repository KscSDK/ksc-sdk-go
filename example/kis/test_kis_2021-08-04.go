package main

import (
	"encoding/json"
	"fmt"
	"github.com/KscSDK/ksc-sdk-go/ksc"
	"github.com/KscSDK/ksc-sdk-go/ksc/utils"
	"github.com/KscSDK/ksc-sdk-go/service/kis"
)

func main() {
	ak := "你的ak"
	sk := "你的sk"
	region := "cn-beijing-6"
	//debug模式的话 打开这个开关
	svc := kis.SdkNew(ksc.NewClient(ak, sk /*,true*/), &ksc.Config{Region: &region}, &utils.UrlInfo{
		UseSSL: true,
	})
	var resp *map[string]interface{}
	var err error
	//	****************************************查询交换机批次号***************************************************

	querySwitchBatchIds := make(map[string]interface{})
	querySwitchBatchIds["Limit"] = 5
	querySwitchBatchIds["Offset"] = 0

	resp, err = svc.QuerySwitchBatchIds(&querySwitchBatchIds)
	if err != nil {
		fmt.Println("error:", err.Error())
	}
	if resp != nil {
		str, _ := json.Marshal(&resp)
		fmt.Printf("%+v\n", string(str))
	}

	//	****************************************查询交换机信息**************************************************
	//querySwitchOrderInfos := make(map[string]interface{})
	//querySwitchOrderInfos["BatchId"]  = "QYBD_内网核心交换机_375689"
	//querySwitchOrderInfos["Limit"] = 5
	//querySwitchOrderInfos["Offset"] = 0
	//resp, err = svc.QuerySwitchOrderInfos(&querySwitchOrderInfos)
	//if err != nil {
	//	fmt.Println("error:", err.Error())
	//}
	//if resp != nil {
	//	str, _ := json.Marshal(&resp)
	//	fmt.Printf("%+v\n", string(str))
	//}
	//	****************************************查询服务器批次号**************************************************
	//queryServerBatchIds := make(map[string]interface{})
	//queryServerBatchIds["Limit"] = 5
	//queryServerBatchIds["Offset"] = 0
	//resp, err = svc.QueryServerBatchIds(&queryServerBatchIds)
	//if err != nil {
	//	fmt.Println("error:", err.Error())
	//}
	//if resp != nil {
	//	str, _ := json.Marshal(&resp)
	//	fmt.Printf("%+v\n", string(str))
	//}
	//	****************************************查询服务器信息***************************************************
	//queryServerOrderInfos := make(map[string]interface{})
	//queryServerOrderInfos["BatchId"] = "projectNumber11-newService11-375588-MI-S7-1-3"
	//queryServerOrderInfos["Limit"] = 5
	//queryServerOrderInfos["Offset"] = 0
	//resp, err = svc.QueryServerOrderInfos(&queryServerOrderInfos)
	//if err != nil {
	//	fmt.Println("error:", err.Error())
	//}
	//if resp != nil {
	//	str, _ := json.Marshal(&resp)
	//	fmt.Printf("%+v\n", string(str))
	//}
	//	****************************************查询线号信息***************************************************
	//queryPortLineInfos := make(map[string]interface{})
	//queryPortLineInfos["SwitchSn"] = "G1Q515N001190"
	//queryPortLineInfos["Limit"] = 5
	//queryPortLineInfos["Offset"] = 0
	//resp, err = svc.QueryPortLineInfos(&queryPortLineInfos)
	//if err != nil {
	//	fmt.Println("error:", err.Error())
	//}
	//if resp != nil {
	//	str, _ := json.Marshal(&resp)
	//	fmt.Printf("%+v\n", string(str))
	//}
}
