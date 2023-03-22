package main

import (
	"encoding/json"
	"fmt"
	"github.com/KscSDK/ksc-sdk-go/ksc"
	"github.com/KscSDK/ksc-sdk-go/ksc/utils"
	"github.com/KscSDK/ksc-sdk-go/service/sqlserver"
)

func main() {
	ak := "your ak"
	sk := "your sk"
	region := "cn-shanghai-3"
	//debug模式的话 打开这个开关
	ss := sqlserver.SdkNew(ksc.NewClient(ak, sk, true), &ksc.Config{Region: &region}, &utils.UrlInfo{
		UseSSL: false,
		Locate: false,
	})
	var resp *map[string]interface{}
	var err error

	describeAddresses := make(map[string]interface{})
	describeAddresses["DBInstanceStatus"] = "active"
	resp, err = ss.DescribeDBInstances(&describeAddresses)
	fmt.Println(resp)
	if err != nil {
		fmt.Println("error:", err.Error())
	}
	if resp != nil {
		str, _ := json.Marshal(&resp)
		fmt.Printf("%+v\n", string(str))
	}
}
