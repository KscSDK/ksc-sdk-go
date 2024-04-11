package main

import (
	"encoding/json"
	"fmt"
	"github.com/KscSDK/ksc-sdk-go/ksc"
	"github.com/KscSDK/ksc-sdk-go/ksc/utils"
	"github.com/KscSDK/ksc-sdk-go/service/cdnv3"
)

func main() {
	ak := "AKLT7QA2mrRWTTG9glDtSshA"
	sk := "OO1Rb37LorWEyi1BUJqviGkvPq1krhClVeKfFdoz"

	region := "cn-shanghai-2"

	//debug模式的话 打开这个开关
	svc := cdnv3.SdkNew(ksc.NewClient(ak, sk, true), &ksc.Config{Region: &region}, &utils.UrlInfo{
		UseSSL: true,
	})

	var resp *map[string]interface{}
	var err error

	//	**************************************** 本接口用于屏蔽、解除屏蔽URL。 POST（BlockDomainUrl）***************************************************
	// Parameters:
	//       BlockType         string    操作接口名，系统规定参数 取值：block：屏蔽URL；unblock：解除屏蔽；
	//       Urls              Url[]    URL列表
	//       其中url[]为：
	//		Url String 需要提交屏蔽/解屏蔽的Url，单条
	//
	url := make(map[string]interface{})
	url["Url"] = "http://ntj14122.test.com/abc.txt"

	urls := make([]map[string]interface{}, 0)
	urls = append(urls, url)
	blockDomainUrl := make(map[string]interface{})
	blockDomainUrl["BlockType"] = "unblock"
	blockDomainUrl["Urls"] = urls
	blockDomainUrl["BlockTime"] = 604800

	resp, err = svc.BlockDomainUrlPost(&blockDomainUrl)
	if err != nil {
		fmt.Println("error:", err.Error())
	}
	if resp != nil {
		str, _ := json.Marshal(&resp)
		fmt.Printf("%+v\n", string(str))
	}
}
