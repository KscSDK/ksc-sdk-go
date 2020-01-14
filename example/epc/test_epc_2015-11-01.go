package main

import (
	"encoding/json"
	"fmt"
	"github.com/KscSDK/ksc-sdk-go/ksc"
	"github.com/KscSDK/ksc-sdk-go/ksc/utils"
	"github.com/KscSDK/ksc-sdk-go/service/epc"
)

func main() {
	ak := "你的ak"
	sk := "你的sk"
	region := "cn-beijing-6"
	//debug模式的话 打开这个开关
	svc := epc.SdkNew(ksc.NewClient(ak, sk /*,true*/), &ksc.Config{Region: &region}, &utils.UrlInfo{
		UseSSL: true,
	})
	var resp *map[string]interface{}
	var err error
	//	****************************************查询物理机 DescribeEpcs***************************************************

	describeEpcs := make(map[string]interface{})
	describeEpcs["Filter.1.Name"] = "host-type"
	describeEpcs["Filter.1.Value.1"] = "CAL"

	resp, err = svc.DescribeEpcs(&describeEpcs)
	if err != nil {
		fmt.Println("error:", err.Error())
	}
	if resp != nil {
		str, _ := json.Marshal(&resp)
		fmt.Printf("%+v\n", string(str))
	}

	//	****************************************创建云物理机***************************************************

	//	createEpcs := make(map[string]interface{})
	//	createEpcs["HostType"] = "CAL"
	//	createEpcs["Raid"] = "Raid1"
	//	createEpcs["ImageId"] = "prt72bd5d05-69eb-4468-8087-ebbb62b95c29"
	//	createEpcs["NetworkInterfaceMode"] = "bond4"
	//	createEpcs["SubnetId"] = "6ff84276-e359-4ff1-a53f-59b243299079"
	//	createEpcs["SecurityGroupId.1"] = "0dcd8356-9e7e-43ae-897b-a8dac8914e62"
	//	createEpcs["KeyId"] = "8db54426-ed02-403b-9aca-470351e5a037"
	//	createEpcs["ChargeType"] = "Daily"
	//	createEpcs["Password"] = "ABCabc123"
	//	createEpcs["AvailabilityZone"] = "cn-shanghai-3b"
	//
	//	resp, err = svc.CreateEpc(&createEpcs)
	//	if err != nil {
	//		fmt.Println("error:", err.Error())
	//	}
	//	if resp != nil {
	//		str, _ := json.Marshal(&resp)
	//		fmt.Printf("%+v\n", string(str))
	//	}

	//	****************************************删除云物理机***************************************************

	//	deleteEpc := make(map[string]interface{})
	//	deleteEpc["HostId"] = ""
	//	resp, err = svc.DeleteEpc(&deleteEpc)
	//	if err != nil {
	//		fmt.Println("error:", err.Error())
	//	}
	//	if resp != nil {
	//		str, _ := json.Marshal(&resp)
	//		fmt.Printf("%+v\n", string(str))
	//	}

	//	****************************************更新云物理主机信息***************************************************

	//	modifyEpc := make(map[string]interface{})
	//	modifyEpc["HostId"] = "fb7dacb1-ba21-4d1f-9d83-63622f6a498d"
	//	modifyEpc["HostName"] = "modify"
	//	resp, err = svc.ModifyEpc(&modifyEpc)
	//	if err != nil {
	//		fmt.Println("error:", err.Error())
	//	}
	//	if resp != nil {
	//		str, _ := json.Marshal(&resp)
	//		fmt.Printf("%+v\n", string(str))
	//	}

	//******************************************修改超线程*********************************************************

	//	modifyHyperThreading := make(map[string]interface{})
	//	modifyHyperThreading["HostId"] = "fb7dacb1-ba21-4d1f-9d83-63622f6a498d"
	//	modifyHyperThreading["HyperThreadingStatus"] = "start"
	//	resp, err = svc.ModifyHyperThreading(&modifyHyperThreading)
	//	if err != nil {
	//		fmt.Println("error:", err.Error())
	//	}
	//	if resp != nil {
	//		str, _ := json.Marshal(&resp)
	//		fmt.Printf("%+v\n", string(str))
	//	}

	//******************************************重置密码*********************************************************

	//	resetPassword := make(map[string]interface{})
	//	resetPassword["HostId"] = "fb7dacb1-ba21-4d1f-9d83-63622f6a498d"
	//	resetPassword["Password"] = "AAAaaa111"
	//	resp, err = svc.ResetPassword(&resetPassword)
	//	if err != nil {
	//		fmt.Println("error:", err.Error())
	//	}
	//	if resp != nil {
	//		str, _ := json.Marshal(&resp)
	//		fmt.Printf("%+v\n", string(str))
	//	}

	//******************************************启动物理机*********************************************************

	//	startEpc := make(map[string]interface{})
	//	startEpc["HostId"] = "fb7dacb1-ba21-4d1f-9d83-63622f6a498d"
	//	resp, err = svc.StartEpc(&startEpc)
	//	if err != nil {
	//		fmt.Println("error:", err.Error())
	//	}
	//	if resp != nil {
	//		str, _ := json.Marshal(&resp)
	//		fmt.Printf("%+v\n", string(str))
	//	}

	//******************************************创建自定义镜像*********************************************************

	//	createImage := make(map[string]interface{})
	//	createImage["HostId"] = "fb7dacb1-ba21-4d1f-9d83-63622f6a498d"
	//	createImage["ImageName"]="createImage"
	//	resp, err = svc.CreateImage(&createImage)
	//	if err != nil {
	//		fmt.Println("error:", err.Error())
	//	}
	//	if resp != nil {
	//		str, _ := json.Marshal(&resp)
	//		fmt.Printf("%+v\n", string(str))
	//	}

	//******************************************修改自定义镜像*********************************************************

	//	modifyImage := make(map[string]interface{})
	//	modifyImage["ImageId"] = "id"
	//	modifyImage["ImageName"] = "modifyImage"
	//	resp, err = svc.ModifyImage(&modifyImage)
	//	if err != nil {
	//		fmt.Println("error:", err.Error())
	//	}
	//	if resp != nil {
	//		str, _ := json.Marshal(&resp)
	//		fmt.Printf("%+v\n", string(str))
	//	}

	//******************************************删除自定义镜像*********************************************************

	//	deleteImage := make(map[string]interface{})
	//	deleteImage["ImageId"] = "id"
	//	resp, err = svc.DeleteImage(&deleteImage)
	//	if err != nil {
	//		fmt.Println("error:", err.Error())
	//	}
	//	if resp != nil {
	//		str, _ := json.Marshal(&resp)
	//		fmt.Printf("%+v\n", string(str))
	//	}

	//******************************************查看云物理主机镜像信息*********************************************************

	//	resp, err = svc.DescribeImages(nil)
	//	if err != nil {
	//		fmt.Println("error:", err.Error())
	//	}
	//	if resp != nil {
	//		str, _ := json.Marshal(&resp)
	//		fmt.Printf("%+v\n", string(str))
	//	}

	//******************************************查看物理机巡检信息*********************************************************

	//	resp, err = svc.DescribeInspections(nil)
	//	if err != nil {
	//		fmt.Println("error:", err.Error())
	//	}
	//	if resp != nil {
	//		str, _ := json.Marshal(&resp)
	//		fmt.Printf("%+v\n", string(str))
	//	}

	//******************************************修改网卡信息*********************************************************

	//	modifyNetworkInterfaceAttribute:=make(map[string] interface{})
	//	modifyNetworkInterfaceAttribute["NetworkInterfaceId"]="fb7dacb1-ba21-4d1f-9d83-63622f6a498d"
	//	modifyNetworkInterfaceAttribute["HostId"]="fb7dacb1-ba21-4d1f-9d83-63622f6a498d"
	//	modifyNetworkInterfaceAttribute["SubnetId"]="6ff84276-e359-4ff1-a53f-59b243299079"
	//	//modifyNetworkInterfaceAttribute["DNS1"]="198.18.224.10"
	//	resp, err = svc.ModifyNetworkInterfaceAttribute(&modifyNetworkInterfaceAttribute)
	//	//resp, err = svc.ModifyDns(&modifyNetworkInterfaceAttribute)
	//	if err != nil {
	//		fmt.Println("error:", err.Error())
	//	}
	//	if resp != nil {
	//		str, _ := json.Marshal(&resp)
	//		fmt.Printf("%+v\n", string(str))
	//	}

	//******************************************导入秘钥*********************************************************

	//	importKey:=make(map[string] interface{})
	//	importKey["KeyName"]="testsdk"
	//	importKey["PublicKey"]=""
	//	resp, err = svc.ImportKey(&importKey)
	//	if err != nil {
	//		fmt.Println("error:", err.Error())
	//	}
	//	if resp != nil {
	//		str, _ := json.Marshal(&resp)
	//		fmt.Printf("%+v\n", string(str))
	//	}

	//******************************************删除秘钥*********************************************************

	//	deleteKey := make(map[string]interface{})
	//	deleteKey["KeyId"] = "testsdk"
	//	resp, err = svc.DeleteKey(&deleteKey)
	//	if err != nil {
	//		fmt.Println("error:", err.Error())
	//	}
	//	if resp != nil {
	//		str, _ := json.Marshal(&resp)
	//		fmt.Printf("%+v\n", string(str))
	//	}

	//******************************************修改秘钥信息*********************************************************

	//	modifyKey := make(map[string]interface{})
	//	modifyKey["KeyId"] = "8db54426-ed02-403b-9aca-470351e5a037"
	//	modifyKey["keyName"] = "modify"
	//	resp, err = svc.ModifyKey(&modifyKey)
	//	if err != nil {
	//		fmt.Println("error:", err.Error())
	//	}
	//	if resp != nil {
	//		str, _ := json.Marshal(&resp)
	//		fmt.Printf("%+v\n", string(str))
	//	}

	//******************************************创建秘钥*********************************************************

	//	createKey := make(map[string]interface{})
	//	createKey["keyName"] = "create"
	//	resp, err = svc.CreateKey(&createKey)
	//	if err != nil {
	//		fmt.Println("error:", err.Error())
	//	}
	//	if resp != nil {
	//		str, _ := json.Marshal(&resp)
	//		fmt.Printf("%+v\n", string(str))
	//	}

	//	****************************************获取秘钥列表信息***************************************************

	//	resp, err = svc.DescribeKeys(nil)
	//	if err != nil {
	//		fmt.Println("error:", err.Error())
	//	}
	//	if resp != nil {
	//		str, _ := json.Marshal(&resp)
	//		fmt.Printf("%+v\n", string(str))
	//	}

	//	****************************************获取硬盘监控信息***************************************************

	//	describePhysicalMonitor:=make(map[string] interface{})
	//	describePhysicalMonitor["HostId"]="fb7dacb1-ba21-4d1f-9d83-63622f6a498d"
	//	resp, err = svc.DescribePhysicalMonitor(&describePhysicalMonitor)
	//	if err != nil {
	//		fmt.Println("error:", err.Error())
	//	}
	//	if resp != nil {
	//		str, _ := json.Marshal(&resp)
	//		fmt.Printf("%+v\n", string(str))
	//	}

	//	****************************************获取云物理主机托管机柜列表信息***************************************************

	//	resp, err = svc.DescribeCabinets(nil)
	//	if err != nil {
	//		fmt.Println("error:", err.Error())
	//	}
	//	if resp != nil {
	//		str, _ := json.Marshal(&resp)
	//		fmt.Printf("%+v\n", string(str))
	//	}

	//	****************************************获取配件列表信息***************************************************

	//	resp, err = svc.DescribeAccessorys(nil)
	//	if err != nil {
	//		fmt.Println("error:", err.Error())
	//	}
	//	if resp != nil {
	//		str, _ := json.Marshal(&resp)
	//		fmt.Printf("%+v\n", string(str))
	//	}

	//	****************************************销售配件***************************************************

	//	buyAccessory := make(map[string]interface{})
	//	buyAccessory["AvailabilityZone"] = "cn-shanghai-3a"
	//	buyAccessory["AccessoryType"] = "NetworkCard"
	//	buyAccessory["AccessorySuit"] = "NetworkCard-100GB"
	//	buyAccessory["ChargeType"] = "PostpaidByTime"
	//
	//	resp, err = svc.BuyAccessory(&buyAccessory)
	//	if err != nil {
	//		fmt.Println("error:", err.Error())
	//	}
	//	if resp != nil {
	//		str, _ := json.Marshal(&resp)
	//		fmt.Printf("%+v\n", string(str))
	//	}

	//******************************************删除配件*****************************************************

	//	deleteAccessory:=make(map[string] interface{})
	//	deleteAccessory["AccessoryId"]="e658d9c8-835b-4e87-b2ab-c277c5c3e7c9"
	//	resp, err = svc.DeleteAccessory(&deleteAccessory)
	//	if err != nil {
	//		fmt.Println("error:", err.Error())
	//	}
	//	if resp != nil {
	//		str, _ := json.Marshal(&resp)
	//		fmt.Printf("%+v\n", string(str))
	//	}

	//****************************************创建带外管理****************************************************

	//	createRemoteManagement := make(map[string]interface{})
	//	createRemoteManagement["PhoneNumber"] = ""
	//	createRemoteManagement["pin"]=""
	//	resp, err = svc.CreateRemoteManagement(&createRemoteManagement)
	//	if err != nil {
	//		fmt.Println("error:", err.Error())
	//	}
	//	if resp != nil {
	//		str, _ := json.Marshal(&resp)
	//		fmt.Printf("%+v\n", string(str))
	//	}

	//	****************************************获取验证码***************************************************

	//	getDynamicCode:=make(map[string] interface{})
	//	getDynamicCode["RemoteManagementId"]=""
	//	resp, err = svc.GetDynamicCode(nil)
	//	if err != nil {
	//		fmt.Println("error:", err.Error())
	//	}
	//	if resp != nil {
	//		str, _ := json.Marshal(&resp)
	//		fmt.Printf("%+v\n", string(str))
	//	}
	//	****************************************查询云物理主机的库存***************************************************
	//	//describeEpcStocks:=make(map[string] interface{})
	//	//describeEpcStocks["Filter.1.Name"]="host-type"
	//	//describeEpcStocks["Filter.1.Value.1"]="CAL"
	//	resp, err = svc.DescribeEpcStocks(nil)
	//	if err != nil {
	//		fmt.Println("error:", err.Error())
	//	}
	//	if resp != nil {
	//		str, _ := json.Marshal(&resp)
	//		fmt.Printf("%+v\n", string(str))
	//	}

	//	****************************************查询云物理机型配置信息***************************************************
	////describeEpcDeviceAttributes:=make(map[string] interface{})
	////describeEpcDeviceAttributes["Filter.1.Name"]="host-type"
	////describeEpcDeviceAttributes["Filter.1.Value.1"]="CAL"
	//resp, err = svc.DescribeEpcDeviceAttributes(nil)
	//if err != nil {
	//	fmt.Println("error:", err.Error())
	//}
	//if resp != nil {
	//	str, _ := json.Marshal(&resp)
	//	fmt.Printf("%+v\n", string(str))
	//}
}
