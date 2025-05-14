package main

import (
	"encoding/json"
	"fmt"

	"github.com/KscSDK/ksc-sdk-go/ksc"
	"github.com/KscSDK/ksc-sdk-go/ksc/utils"
	"github.com/KscSDK/ksc-sdk-go/service/kpfs"
)

func main() {
	ak := "you ak"
	sk := "you sk"
	region := "cn-beijing-6"
	//debug模式的话 打开这个开关
	svc := kpfs.SdkNew(ksc.NewClient(ak, sk /*,true*/), &ksc.Config{Region: &region}, &utils.UrlInfo{
		UseSSL: true,
	})
	var resp *map[string]interface{}
	var err error
	//	****************************************查询交换机批次号***************************************************

	queryAcls := make(map[string]interface{})

	resp, err = svc.DescribePerformanceOnePosixAclList(&queryAcls)
	if err != nil {
		fmt.Println("error:", err.Error())
	}
	if resp != nil {
		str, _ := json.Marshal(&resp)
		fmt.Printf("%+v\n", string(str))
	}
}
