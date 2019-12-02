package main

import (
	"encoding/json"
	"fmt"
	"github.com/KscSDK/ksc-sdk-go/ksc"
	"github.com/KscSDK/ksc-sdk-go/ksc/utils"
	"github.com/KscSDK/ksc-sdk-go/service/monitorv2"
)

func main() {
	ak := "你的 AK"
	sk := "你的 SK"
	region := "cn-shanghai-2"
	//debug模式的话 打开这个开关
	svc := monitorv2.SdkNew(ksc.NewClient(ak, sk /*true*/), &ksc.Config{Region: &region}, &utils.UrlInfo{
		UseSSL: true,
	})

	/*
		 *
		****************************************获取监控数据(GetMetricStatisticsBatch())***************************************************

	*/
	//通用产品线（docker 除外)
	l := []string{"Max", "Min", "Avg"}

	m_metrics := make(map[string]interface{})
	m_metrics["InstanceID"] = "df2bd78e-8453-4fea-aabf-a0d5a0832ff8"
	m_metrics["MetricName"] = "cpu.utilizition.total"
	l1 := []interface{}{m_metrics}

	m_vpc := make(map[string]interface{})
	m_vpc["Namespace"] = "kec"
	m_vpc["StartTime"] = "2019-05-05T15:00:00Z"
	m_vpc["EndTime"] = "2019-05-05T15:09:00Z"
	m_vpc["Period"] = 180
	m_vpc["Aggregate"] = l
	m_vpc["Metrics"] = l1

	//docker产品线
	/*
	   	dimensions:=make(map[string]interface{})
	   	dimensions["Dimensions.0.Name"] = "ClusterId"
	   	dimensions["Dimensions.0.Value"] = "b0b69e65-fcc0-4e63-a51c-1b7eb191ec66"
	   	dimensions["Dimensions.1.Name"] = "NamespaceName"
	   	dimensions["Dimensions.1.Value"] = "kube-system.canal.canal-wvvcj"
	   	dimensions["Dimensions.2.Name"] = "DeploymentName"
	   	dimensions["Dimensions.2.Value"] = "dsaddada"

	   	m_metrics := make(map[string]interface{})
	   	m_metrics["MetricName"] = "container.cpu.usage"
	   	m_metrics["Dimensions"] = dimensions
	       l1 := []interface{}{m_metrics}

	   	l := []string{"Max", "Min", "Avg"}

	   	m_vpc := make(map[string]interface{})
	   	m_vpc["Namespace"] = "kec"
	   	m_vpc["StartTime"] = "2019-05-05T15:00:00Z"
	   	m_vpc["EndTime"] = "2019-05-05T15:09:00Z"
	   	m_vpc["Period"] = 1800
	   	m_vpc["Aggregate"] = l
	   	m_vpc["Metrics"] = l1

	*/
	resp, err := svc.GetMetricStatisticsBatch(&m_vpc)
	if err != nil {
		fmt.Println(err.Error())
	}
	if resp != nil {
		str, _ := json.Marshal(&resp)
		fmt.Printf("%+v\n", string(str))
	}

}
