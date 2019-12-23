package main

import (
	"encoding/json"
	"fmt"
	"github.com/KscSDK/ksc-sdk-go/ksc"
	"github.com/KscSDK/ksc-sdk-go/ksc/utils"
	"github.com/KscSDK/ksc-sdk-go/service/vpc"
)

func main() {
	ak := "你的ak"
	sk := "你的sk"
	region := "cn-shanghai-2"
	svc := vpc.SdkNew(ksc.NewClient(ak, sk), &ksc.Config{Region: &region}, &utils.UrlInfo{
		UseSSL: true,
	})

	var resp *map[string]interface{}
	var err error
	//	****************************************创建VPC(CreateVpc())***************************************************
	createVpc := make(map[string]interface{})
	createVpc["VpcName"] = "xuan"
	createVpc["CidrBlock"] = "10.0.0.0/8"

	resp, err = svc.CreateVpc(&createVpc)
	if err != nil {
		fmt.Println("error:", err.Error())
	}
	if resp != nil {
		str, _ := json.Marshal(&resp)
		fmt.Printf("%+v\n", string(str))
	}
	/*
	//	****************************************删除VPC(DeleteVpc())***************************************************
	deleteVpc := make(map[string]interface{})
	deleteVpc["LoadBalancerId"] = ""
	resp, err = svc.DeleteVpc(&deleteVpc)
	if err != nil {
		fmt.Println("error:", err.Error())
	}
	if resp != nil {
		str, _ := json.Marshal(&resp)
		fmt.Printf("%+v\n", string(str))
	}
	//	****************************************修改VPC(ModifyVpc())***************************************************
	modifyVpc := make(map[string]interface{})
	modifyVpc["LoadBalancerId"] = ""
	modifyVpc["LoadBalancerName"] = ""

	resp, err = svc.ModifyVpc(&modifyVpc)
	if err != nil {
		fmt.Println("error:", err.Error())
	}
	if resp != nil {
		str, _ := json.Marshal(&resp)
		fmt.Printf("%+v\n", string(str))
	}
	//	****************************************描述VPC信息(DescribeVpcs())***************************************************


	m_vpc := make(map[string]interface{})
	m_vpc["VpcId.1"] = "b2adb6b5-1575-4833-9a4f-536c152932f4"

	resp, err = svc.DescribeVpcs(&m_vpc)
	if err != nil {
		fmt.Println("there was an error listing instances in", err.Error())
	}
	if resp != nil {
		l := (*resp)["VpcSet"].([]interface{})
		for i := 0; i < len(l); i++ {
			item := l[i].(map[string]interface{})
			fmt.Printf("%+v:%+v\n", item["VpcId"], item["VpcName"])
		}
	}
	m_vnet := make(map[string]interface{})
	resp1, err1 := svc.DescribeSubnets(&m_vnet)
	if err1 != nil {
		fmt.Println("there was an error listing instances in", err1.Error())
	}
	if resp1 != nil {
		l := (*resp1)["SubnetSet"].([]interface{})
		for i := 0; i < len(l); i++ {
			item := l[i].(map[string]interface{})
			fmt.Printf("%+v:%+v\n", item["SubnetId"], item["SubnetName"])
		}
	}
	 */
}
