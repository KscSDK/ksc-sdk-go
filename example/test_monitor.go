package main

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/ksc/ksc-sdk-go/ksc/utils"
	"github.com/ksc/ksc-sdk-go/service/monitor"
)

func main() {
	ak := "你的AK"
	sk := "你的SK"
	region := "cn-beijing-6"
	_credentials := credentials.NewStaticCredentials(ak, sk, "")
	sess, err := session.NewSession(&aws.Config{Credentials: _credentials})
	svc := monitor.ExtraNew(&utils.UrlInfo{
		UseSSL: true,
		Locate: true,
	}, sess, &aws.Config{Region: &region})
	if err != nil {
		fmt.Printf("%v", err)
	}

	//var ids= []string{"ae19a823-e920-446b-a720-ceab512c7673"}

	m_vpc := make(map[string]interface{})
	m_metrics := make(map[string]interface{})
	m_metrics["InstanceID"] = "f4ba1a48-7538-4ce9-a628-26ab9889d1bb"
	m_metrics["MetricName"] = "eip.bps.in"
	l := []string{"Max", "Min", "Avg"}
	l1 := []interface{}{m_metrics}
	m_vpc["Namespace"] = "eip"
	m_vpc["StartTime"] = "2019-05-05T15:00:00Z"
	m_vpc["EndTime"] = "2019-05-05T0115:59:00Z"
	m_vpc["Period"] = 1800
	m_vpc["Aggregate"] = l
	m_vpc["Metrics"] = l1

	resp, err := svc.GetMetricStatisticsBatch(&m_vpc)
	if err != nil {
		fmt.Println("there was an error listing instances in", err.Error())
	}
	if resp != nil {
		str, _ := json.Marshal(&resp)
		fmt.Printf("%+v\n", string(str))
	}

}
