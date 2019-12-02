package main

import (
	"encoding/json"
	"fmt"
	"github.com/KscSDK/ksc-sdk-go/ksc"
	"github.com/KscSDK/ksc-sdk-go/ksc/utils"
	"github.com/KscSDK/ksc-sdk-go/service/sks"
)

func main() {
	ak := "你的ak"
	sk := "你的sk"
	region := "cn-beijing-6"
	//debug模式的话 打开这个开关
	svc := sks.SdkNew(ksc.NewClient(ak, sk /*,true*/), &ksc.Config{Region: &region}, &utils.UrlInfo{
		UseSSL: true,
	})
	var resp *map[string]interface{}
	var err error
	//	****************************************查询密钥 DescribeKeys***************************************************
	describeKeys := make(map[string]interface{})
	describeKeys["Filter.1.Name"] = "key-name"
	describeKeys["Filter.1.Value.1"] = "new-test"

	resp, err = svc.DescribeKeys(&describeKeys)
	if err != nil {
		fmt.Println("error:", err.Error())
	}
	if resp != nil {
		str, _ := json.Marshal(&resp)
		fmt.Printf("%+v\n", string(str))
	}
	//	****************************************创建密钥 CreateKey***************************************************
	//createKey := make(map[string]interface{})
	//createKey["KeyName"] = "key-test"
	//
	//resp, err = svc.CreateKey(&createKey)
	//if err != nil {
	//	fmt.Println("error:", err.Error())
	//}
	//if resp != nil {
	//	str, _ := json.Marshal(&resp)
	//	fmt.Printf("%+v\n", string(str))
	//}
	//	****************************************更新密钥 ModifyKey***************************************************
	//modifyKey := make(map[string]interface{})
	//modifyKey["KeyId"] = "2d7b9ce0-af3e-475d-8c0d-265e084679bb"
	//
	//modifyKey["KeyName"] = "key-test-ttt"
	//
	//	resp, err = svc.ModifyKey(&modifyKey)
	//if err != nil {
	//	fmt.Println("error:", err.Error())
	//}
	//if resp != nil {
	//	str, _ := json.Marshal(&resp)
	//	fmt.Printf("%+v\n", string(str))
	//}
	//	****************************************导入密钥 ImportKey***************************************************
	//importKey := make(map[string]interface{})
	//importKey["PublicKey"] = "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAAAgQCKdMMmtsycUm1bRcltzGYBWKTzm27MlH//P23cycAf6bHFqWCpXahwxvILr7xjB7f43JYDWTz8oTeR4cfjo4UOG652O+KY37m+k98lCv8PPhrRNqctFh1W6E/94d0uO+P3noQAwGW8ljMfM5Q3TgM13C8oSYY6qLHIbj2Yz49PZw== root\n"
	//
	//importKey["KeyName"] = "key-test-ttt"
	//
	//resp, err = svc.ImportKey(&importKey)
	//if err != nil {
	//	fmt.Println("error:", err.Error())
	//}
	//if resp != nil {
	//	str, _ := json.Marshal(&resp)
	//	fmt.Printf("%+v\n", string(str))
	//}
}
