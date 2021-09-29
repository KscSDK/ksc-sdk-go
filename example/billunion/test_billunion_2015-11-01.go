package main

import (
	"encoding/json"
	"fmt"
	"github.com/KscSDK/ksc-sdk-go/ksc"
	"github.com/KscSDK/ksc-sdk-go/ksc/utils"
	"github.com/KscSDK/ksc-sdk-go/service/billunion"
)

func main() {
	ak := "你的ak"
	sk := "你的sk"
	region := "cn-beijing-6"

	//debug模式的话 打开这个开关
	svc := billunion.SdkNew(ksc.NewClient(ak, sk, false), &ksc.Config{Region: &region}, &utils.UrlInfo{
		UseSSL: false,
	})

	var resp *map[string]interface{}
	var err error

	//	****************************************查询实例账单汇总（DescribeInstanceSummaryBills）***************************************************
	params := make(map[string]interface{})
	params["BillMonth"] = "2021-06"
	//params["ProductCode"] = "xxxxx"
	params["Page"] = 1
	params["Size"] = 10

	resp, err = svc.DescribeInstanceSummaryBills(&params)
	if err != nil {
		fmt.Println("error:", err.Error())
	}
	if resp != nil {
		str, _ := json.Marshal(&resp)
		fmt.Printf("%+v\n", string(str))
	}

	//	****************************************获取计费类别汇总账单（DescribeBillSummaryByPayMode）***************************************************
	params1 := make(map[string]interface{})
	params1["BillBeginMonth"] = "2021-06"
	params1["BillEndMonth"] = "2021-06"

	resp, err = svc.DescribeBillSummaryByPayMode(&params1)
	if err != nil {
		fmt.Println("error:", err.Error())
	}
	if resp != nil {
		str, _ := json.Marshal(&resp)
		fmt.Printf("%+v\n", string(str))
	}

	params2 := make(map[string]interface{})
	params2["BillBeginMonth"] = "2021-06"
	params2["BillEndMonth"] = "2021-06"

	resp, err = svc.DescribeBillSummaryByProduct(&params2)
	if err != nil {
		fmt.Println("error:", err.Error())
	}
	if resp != nil {
		str, _ := json.Marshal(&resp)
		fmt.Printf("%+v\n", string(str))
	}

	params3 := make(map[string]interface{})
	params3["BillBeginMonth"] = "2021-06"
	params3["BillEndMonth"] = "2021-06"

	resp, err = svc.DescribeBillSummaryByProject(&params3)
	if err != nil {
		fmt.Println("error:", err.Error())
	}
	if resp != nil {
		str, _ := json.Marshal(&resp)
		fmt.Printf("%+v\n", string(str))
	}

	params4 := make(map[string]interface{})

	resp, err = svc.DescribeProductCode(&params4)
	if err != nil {
		fmt.Println("error:", err.Error())
	}
	if resp != nil {
		str, _ := json.Marshal(&resp)
		fmt.Printf("%+v\n", string(str))
	}
}
