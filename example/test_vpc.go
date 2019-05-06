package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/ksc/ksc-sdk-go/ksc/utils"
	"github.com/ksc/ksc-sdk-go/service/vpc"
)

func main() {
	ak := "你的AK"
	sk := "你的SK"
	region := "cn-beijing-6"
	_credentials := credentials.NewStaticCredentials(ak, sk, "")
	sess, err := session.NewSession(&aws.Config{Credentials: _credentials})
	svc := vpc.ExtraNew(&utils.UrlInfo{
		UseSSL: true,
		Locate: true,
	}, sess, &aws.Config{Region: &region})
	if err != nil {
		fmt.Printf("%v", err)
	}

	//var ids= []string{"ae19a823-e920-446b-a720-ceab512c7673"}

	m_vpc := make(map[string]interface{})
	m_vpc["VpcId.1"] = "b2adb6b5-1575-4833-9a4f-536c152932f4"

	resp, err := svc.DescribeVpcs(&m_vpc)
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

}
