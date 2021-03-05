package main

import (
	"fmt"
	"github.com/KscSDK/ksc-sdk-go/ksc"
	"github.com/KscSDK/ksc-sdk-go/ksc/utils"
	"github.com/KscSDK/ksc-sdk-go/service/tag"
)

func main() {
	ak := "您的AK"
	sk := "您的SK"
	region := "cn-beijing-6"
	svc := tag.SdkNew(ksc.NewClient(ak, sk), &ksc.Config{Region: &region}, &utils.UrlInfo{
		UseSSL: true,
	})

	var resp *map[string]interface{}
	var err error
	//	****************************************查询Tag(DescribeTags())***************************************************
	describeTags := make(map[string]interface{})
	//查询云主机的标签
	describeTags["Filter.1.Name"] = "resource-type"
	describeTags["Filter.1.Value.1"] = "kec-instance"
	//使用标签的KEY作为固定条件
	describeTags["Filter.2.Name"] = "key"
	//自定义的标签KEY
	describeTags["Filter.2.Value.1"] = "xym222"
	//使用标签的Value作为固定条件
	describeTags["Filter.3.Name"] = "value"
	//自定义的标签value
	describeTags["Filter.3.Value.1"] = "11"
	resp, err = svc.DescribeTags(&describeTags)
	if err != nil {
		fmt.Println("error:", err.Error())
	}
	if resp != nil {
		//获取主机实例
		v, _ := utils.GetSdkValue("TagSet", *resp)
		if v != nil {
			tagSet := v.([]interface{})
			for _, j := range tagSet {
				fmt.Printf("%+v\n", j.(map[string]interface{})["ResourceId"])
			}
		}
	}
	//	****************************************查询tag(CreateTag())***************************************************
}
