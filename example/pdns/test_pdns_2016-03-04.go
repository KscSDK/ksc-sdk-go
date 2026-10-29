package main

import (
	"encoding/json"
	"fmt"
	"github.com/KscSDK/ksc-sdk-go/ksc"
	"github.com/KscSDK/ksc-sdk-go/ksc/utils"
	"github.com/KscSDK/ksc-sdk-go/service/pdns"
)

func main() {
	ak := "xxx"
	sk := "xxx"
	region := "cn-beijing-6"
	svc := pdns.SdkNew(ksc.NewClient(ak, sk), &ksc.Config{Region: &region}, &utils.UrlInfo{
		UseSSL: true,
	})

	var resp *map[string]interface{}
	var err error
	//	****************************************查询解析纪录(DescribeZoneRecord)***************************************************
	//describeZoneRecord := make(map[string]interface{})
	//describeZoneRecord["ZoneId"] = "xxx"
	//
	//resp, err = svc.DescribeZoneRecord(&describeZoneRecord)
	//
	//if err != nil {
	//	fmt.Println("error:", err.Error())
	//}
	//if resp != nil {
	//	str, _ := json.Marshal(&resp)
	//	fmt.Printf("%+v\n", string(str))
	//}

	//	****************************************修改解析记录(ModifyZoneRecords)***************************************************
	// 该接口以 JSON body 提交，map 的 key 原样作为 JSON 字段名（不做大小写转换），
	// 故需使用接口约定的 camelCase 字段：zoneId、recordId、recordTtl、recordDataSet。
	modifyZoneRecords := make(map[string]interface{})
	modifyZoneRecords["ZoneId"] = "xxx"
	modifyZoneRecords["RecordId"] = "xxx"
	//modifyZoneRecords["recordTtl"] = 0

	recordDataSet := []map[string]interface{}{
		{
			"RecordValue": "1.1.1.5",
			"Priority":    0,
			"Weight":      0,
			"Port":        0,
		},
	}
	modifyZoneRecords["RecordDataSet"] = recordDataSet

	resp, err = svc.ModifyZoneRecords(&modifyZoneRecords)
	if err != nil {
		fmt.Println("error:", err.Error())
	}
	if resp != nil {
		str, _ := json.Marshal(&resp)
		fmt.Printf("%+v\n", string(str))
	}

}
