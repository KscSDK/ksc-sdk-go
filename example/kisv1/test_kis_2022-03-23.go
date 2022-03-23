package main

import (
	"encoding/json"
	"fmt"
	"github.com/KscSDK/ksc-sdk-go/ksc"
	"github.com/KscSDK/ksc-sdk-go/ksc/utils"
	"github.com/KscSDK/ksc-sdk-go/service/kisv1"
)

func main() {
	ak := "你的ak"
	sk := "你的sk"
	region := "cn-beijing-6"
	//debug模式的话 打开这个开关
	svc := kisv1.SdkNew(ksc.NewClient(ak, sk /*,true*/), &ksc.Config{Region: &region}, &utils.UrlInfo{
		UseSSL: true,
	})
	var resp *map[string]interface{}
	var err error
	//	****************************************1.1.1列出所有region***************************************************

	//listRegion := make(map[string]interface{})
	//
	//resp, err = svc.ListRegion(&listRegion)
	//if err != nil {
	//	fmt.Println("error:", err.Error())
	//}
	//if resp != nil {
	//	str, _ := json.Marshal(&resp)
	//	fmt.Printf("%+v\n", string(str))
	//}

	//	****************************************1.1.2列出所有机房*************************************************
	//listIdc := make(map[string]interface{})
	//
	//resp, err = svc.ListIdc(&listIdc)
	//if err != nil {
	//	fmt.Println("error:", err.Error())
	//}
	//if resp != nil {
	//	str, _ := json.Marshal(&resp)
	//	fmt.Printf("%+v\n", string(str))
	//}
	//	****************************************1.1.3列出运营商**************************************************
	//listIsp := make(map[string]interface{})
	//
	//resp, err = svc.ListIsp(&listIsp)
	//if err != nil {
	//	fmt.Println("error:", err.Error())
	//}
	//if resp != nil {
	//	str, _ := json.Marshal(&resp)
	//	fmt.Printf("%+v\n", string(str))
	//}
	//	****************************************1.1.4列出高速通道类别**************************************************
	//listDpClassify := make(map[string]interface{})
	//
	//resp, err = svc.ListDpClassify(&listDpClassify)
	//if err != nil {
	//	fmt.Println("error:", err.Error())
	//}
	//if resp != nil {
	//	str, _ := json.Marshal(&resp)
	//	fmt.Printf("%+v\n", string(str))
	//}
	//	****************************************1.1.5列出高速通道接入地址**************************************************
	//listDpAddress := make(map[string]interface{})
	//
	//resp, err = svc.ListDpAddress(&listDpAddress)
	//if err != nil {
	//	fmt.Println("error:", err.Error())
	//}
	//if resp != nil {
	//	str, _ := json.Marshal(&resp)
	//	fmt.Printf("%+v\n", string(str))
	//}
	//	****************************************1.1.6列出裸金属服务类型**************************************************
	//listDeviceProductType := make(map[string]interface{})
	//
	//resp, err = svc.ListDeviceProductType(&listDeviceProductType)
	//if err != nil {
	//	fmt.Println("error:", err.Error())
	//}
	//if resp != nil {
	//	str, _ := json.Marshal(&resp)
	//	fmt.Printf("%+v\n", string(str))
	//}
	//	****************************************1.2.1查询机柜服务***************************************************
	//getCabinet := make(map[string]interface{})
	//getCabinet["Limit"] = 5
	//getCabinet["Offset"] = 0
	//resp, err = svc.GetCabinet(&getCabinet)
	//if err != nil {
	//	fmt.Println("error:", err.Error())
	//}
	//if resp != nil {
	//	str, _ := json.Marshal(&resp)
	//	fmt.Printf("%+v\n", string(str))
	//}
	//	****************************************1.2.2查询裸金属服务***************************************************
	//getDevice := make(map[string]interface{})
	//getDevice["Limit"] = 5
	//getDevice["Offset"] = 0
	//resp, err = svc.GetDevice(&getDevice)
	//if err != nil {
	//	fmt.Println("error:", err.Error())
	//}
	//if resp != nil {
	//	str, _ := json.Marshal(&resp)
	//	fmt.Printf("%+v\n", string(str))
	//}
	//	****************************************1.2.3查询互联网带宽服务***************************************************
	//getBandwidth := make(map[string]interface{})
	//getBandwidth["Limit"] = 5
	//getBandwidth["Offset"] = 0
	//resp, err = svc.GetBandwidth(&getBandwidth)
	//if err != nil {
	//	fmt.Println("error:", err.Error())
	//}
	//if resp != nil {
	//	str, _ := json.Marshal(&resp)
	//	fmt.Printf("%+v\n", string(str))
	//}
	//	****************************************1.2.4查询IP地址服务***************************************************
	//getSubnet := make(map[string]interface{})
	//getSubnet["Limit"] = 5
	//getSubnet["Offset"] = 0
	//resp, err = svc.GetSubnet(&getSubnet)
	//if err != nil {
	//	fmt.Println("error:", err.Error())
	//}
	//if resp != nil {
	//	str, _ := json.Marshal(&resp)
	//	fmt.Printf("%+v\n", string(str))
	//}
	//	****************************************1.2.5查询高速通道服务***************************************************
	//getDp := make(map[string]interface{})
	//getDp["Limit"] = 5
	//getDp["Offset"] = 0
	//resp, err = svc.GetDp(&getDp)
	//if err != nil {
	//	fmt.Println("error:", err.Error())
	//}
	//if resp != nil {
	//	str, _ := json.Marshal(&resp)
	//	fmt.Printf("%+v\n", string(str))
	//}
	//	****************************************1.2.6查询VPC接入服务***************************************************
	//getVpcAccess := make(map[string]interface{})
	//getVpcAccess["Limit"] = 5
	//getVpcAccess["Offset"] = 0
	//resp, err = svc.GetVpcAccess(&getVpcAccess)
	//if err != nil {
	//	fmt.Println("error:", err.Error())
	//}
	//if resp != nil {
	//	str, _ := json.Marshal(&resp)
	//	fmt.Printf("%+v\n", string(str))
	//}
	//	****************************************1.2.7查询上联出口服务务***************************************************
	//getInternetPort := make(map[string]interface{})
	//getInternetPort["Limit"] = 5
	//getInternetPort["Offset"] = 0
	//resp, err = svc.GetInternetPort(&getInternetPort)
	//if err != nil {
	//	fmt.Println("error:", err.Error())
	//}
	//if resp != nil {
	//	str, _ := json.Marshal(&resp)
	//	fmt.Printf("%+v\n", string(str))
	//}
	//	****************************************1.3.1查询带宽流量***************************************************
	//getMonitorData := make(map[string]interface{})
	//getMonitorData["InstanceId"] = "9fd83a5e-33cf-4196-b8df-59fc57478488"
	//resp, err = svc.GetMonitorData(&getMonitorData)
	//if err != nil {
	//	fmt.Println("error:", err.Error())
	//}
	//if resp != nil {
	//	str, _ := json.Marshal(&resp)
	//	fmt.Printf("%+v\n", string(str))
	//}
	//	****************************************1.3.2查询IP地址流量***************************************************
	//getIpaddrData := make(map[string]interface{})
	//getIpaddrData["Ips"] = "9fd83a5e-33cf-4196-b8df-59fc57478488"
	//resp, err = svc.GetIpaddrData(&getIpaddrData)
	//if err != nil {
	//	fmt.Println("error:", err.Error())
	//}
	//if resp != nil {
	//	str, _ := json.Marshal(&resp)
	//	fmt.Printf("%+v\n", string(str))
	//}
	//	****************************************1.3.3列出监控产品类型***************************************************
	//listMonitorProductType := make(map[string]interface{})
	//resp, err = svc.ListMonitorProductType(&listMonitorProductType)
	//if err != nil {
	//	fmt.Println("error:", err.Error())
	//}
	//if resp != nil {
	//	str, _ := json.Marshal(&resp)
	//	fmt.Printf("%+v\n", string(str))
	//}
	//	****************************************1.3.4列出监控粒度***************************************************
	listStep := make(map[string]interface{})
	resp, err = svc.ListStep(&listStep)
	if err != nil {
		fmt.Println("error:", err.Error())
	}
	if resp != nil {
		str, _ := json.Marshal(&resp)
		fmt.Printf("%+v\n", string(str))
	}

}
