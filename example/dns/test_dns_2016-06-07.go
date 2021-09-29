package main

import (
	"encoding/json"
	"fmt"
	"github.com/KscSDK/ksc-sdk-go/ksc"
	"github.com/KscSDK/ksc-sdk-go/ksc/utils"
	"github.com/KscSDK/ksc-sdk-go/service/dns"
)

func main() {
	ak := "你的AK"
	sk := "你的SK"
	region := "cn-beijing-6"
	svc := dns.SdkNew(ksc.NewClient(ak, sk), &ksc.Config{Region: &region}, &utils.UrlInfo{
		UseSSL: true,
	})

	var resp *map[string]interface{}
	var err error
	//	****************************************创建域名(CreateHostedZone)***************************************************
	createHostedZone := make(map[string]interface{})
	createHostedZone["HostedZoneName"] = "HostedZoneName"

	resp, err = svc.CreateHostedZone(&createHostedZone)

	if err != nil {
		fmt.Println("error:", err.Error())
	}
	if resp != nil {
		str, _ := json.Marshal(&resp)
		fmt.Printf("%+v\n", string(str))
	}

	//	****************************************删除域名(DeleteHostedZone)***************************************************
	//deleteHostedZone := make(map[string]interface{})
	//deleteHostedZone["HostedZoneId"] = "HostedZoneId"
	//
	//resp, err = svc.DeleteHostedZone(&deleteHostedZone)
	//
	//if err != nil {
	//	fmt.Println("error:", err.Error())
	//}
	//if resp != nil {
	//	str, _ := json.Marshal(&resp)
	//	fmt.Printf("%+v\n", string(str))
	//}

	//	****************************************描述域名(DescribeHostedZones)***************************************************
	//resp, err = svc.DescribeHostedZones(nil)
	//
	//if err != nil {
	//	fmt.Println("error:", err.Error())
	//}
	//if resp != nil {
	//	str, _ := json.Marshal(&resp)
	//	fmt.Printf("%+v\n", string(str))
	//}
	//

	//	****************************************创建域名记录(CreateResourceRecord)***************************************************
	//createResourceRecord := make(map[string]interface{})
	//createResourceRecord["HostedZoneId"] = "HostedZoneId"
	//createResourceRecord["ResourceRecordName"] = "ResourceRecordName"
	//createResourceRecord["ResourceRecordType"] = "A"
	//createResourceRecord["GeoLocationId"] = "GeoLocationId"
	//createResourceRecord["TTL"] = "60"
	//createResourceRecord["Value"] = "Value"
	//createResourceRecord["Weight"] = "50"
	//
	//resp, err = svc.CreateResourceRecord(&createResourceRecord)
	//
	//if err != nil {
	//	fmt.Println("error:", err.Error())
	//}
	//if resp != nil {
	//	str, _ := json.Marshal(&resp)
	//	fmt.Printf("%+v\n", string(str))
	//}

	//	****************************************修改域名记录(ModifyResourceRecord)***************************************************
	//modifyResourceRecord := make(map[string]interface{})
	//modifyResourceRecord["HostedZoneId"] = "HostedZoneId"
	//modifyResourceRecord["ResourceRecordId"] = "ResourceRecordId"
	//modifyResourceRecord["ResourceRecordName"] = "ResourceRecordName"
	//modifyResourceRecord["ResourceRecordType"] = "A"
	//modifyResourceRecord["GeoLocationId"] = "GeoLocationId"
	//modifyResourceRecord["TTL"] = "60"
	//modifyResourceRecord["Value"] = "Value"
	//modifyResourceRecord["Weight"] = "50"
	//
	//resp, err = svc.ModifyResourceRecord(&modifyResourceRecord)
	//
	//if err != nil {
	//	fmt.Println("error:", err.Error())
	//}
	//if resp != nil {
	//	str, _ := json.Marshal(&resp)
	//	fmt.Printf("%+v\n", string(str))
	//}

	//	****************************************删除域名记录(DeleteResourceRecord)***************************************************
	//deleteResourceRecord := make(map[string]interface{})
	//deleteResourceRecord["HostedZoneId"] = "HostedZoneId"
	//deleteResourceRecord["ResourceRecordId"] = "ResourceRecordId"
	//
	//resp, err = svc.DeleteResourceRecord(&deleteResourceRecord)
	//
	//if err != nil {
	//	fmt.Println("error:", err.Error())
	//}
	//if resp != nil {
	//	str, _ := json.Marshal(&resp)
	//	fmt.Printf("%+v\n", string(str))
	//}

	//	****************************************描述域名记录信息(DescribeResourceRecords)***************************************************
	//describeResourceRecords := make(map[string]interface{})
	//describeResourceRecords["HostedZoneId"] = "HostedZoneId"
	//
	//resp, err = svc.DescribeResourceRecords(&describeResourceRecords)
	//
	//if err != nil {
	//	fmt.Println("error:", err.Error())
	//}
	//if resp != nil {
	//	str, _ := json.Marshal(&resp)
	//	fmt.Printf("%+v\n", string(str))
	//}

	//	****************************************更换域名记录类型(ResourceRecordChangeType)***************************************************
	//resourceRecordChangeType := make(map[string]interface{})
	//resourceRecordChangeType["HostedZoneId"] = "HostedZoneId"
	//resourceRecordChangeType["ResourceRecordName"] = "ResourceRecordName"
	//resourceRecordChangeType["ResourceRecordType"] = "A"
	//resourceRecordChangeType["ResourceRecordTargetType"] = "A"
	//resourceRecordChangeType["GeoLocationId"] = "GeoLocationId"
	//
	//resp, err = svc.ResourceRecordChangeType(&resourceRecordChangeType)
	//
	//if err != nil {
	//	fmt.Println("error:", err.Error())
	//}
	//if resp != nil {
	//	str, _ := json.Marshal(&resp)
	//	fmt.Printf("%+v\n", string(str))
	//}

	//	****************************************描述线路信息(GetGeolocations)***************************************************

	//resp, err = svc.GetGeolocations(nil)
	//
	//if err != nil {
	//	fmt.Println("error:", err.Error())
	//}
	//if resp != nil {
	//	str, _ := json.Marshal(&resp)
	//	fmt.Printf("%+v\n", string(str))
	//}

}
