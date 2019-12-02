package main

import (
	"encoding/json"
	"fmt"
	"github.com/KscSDK/ksc-sdk-go/ksc"
	"github.com/KscSDK/ksc-sdk-go/ksc/utils"
	"github.com/KscSDK/ksc-sdk-go/service/monitor"
)

func main() {
	ak := "你的 AK"
	sk := "你的 SK"
	region := "cn-shanghai-2"
	//debug模式的话 打开这个开关
	svc := monitor.SdkNew(ksc.NewClient(ak, sk /*true*/), &ksc.Config{Region: &region}, &utils.UrlInfo{
		UseSSL: true,
	})
	/*
		 *
		****************************************获取监控数据(GetMetricStatistics())***************************************************

	*/
	//通用产品线（docker 除外)
	/*
		m_metrics := make(map[string]interface{})
		m_metrics["Namespace"] = "kec"
		m_metrics["InstanceID"] = "df2bd78e-8453-4fea-aabf-a0d5a0832ff8"
		m_metrics["MetricName"] = "cpu.utilizition.total"
		m_metrics["StartTime"] = "2019-05-05T15:00:00Z"
		m_metrics["EndTime"] = "2019-05-05T15:09:00Z"
		m_metrics["Period"] = "180"
		m_metrics["Aggregate"] ="Max,min,avg"

		resp, err := svc.GetMetricStatistics(&m_metrics)

	*/
	//docker产品线
	/*
		m_metrics := make(map[string]interface{})
		m_metrics["Namespace"] = "kce"
		m_metrics["MetricName"] = "container.cpu.usage"
		m_metrics["StartTime"] = "2019-05-05T15:00:00Z"
		m_metrics["EndTime"] = "2019-05-05T15:09:00Z"
		m_metrics["Period"] = "180"
		m_metrics["Aggregate"] ="Max,min,avg"

		m_metrics["Dimensions.0.Name"] = "ClusterId"
		m_metrics["Dimensions.0.Value"] = "b0b69e65-fcc0-4e63-a51c-1b7eb191ec66"
		m_metrics["Dimensions.1.Name"] = "NamespaceName"
		m_metrics["Dimensions.1.Value"] = "kube-system.canal.canal-wvvcj"
		m_metrics["Dimensions.2.Name"] = "DeploymentName"
		m_metrics["Dimensions.2.Value"] = "dsaddada"

		resp, err := svc.GetMetricStatistics(&m_metrics)

	*/
	/*
		 *
		****************************************获取监控项(ListMetrics)***************************************************

	*/
	//通用产品线（docker 除外)

	m_metrics := make(map[string]interface{})
	m_metrics["Namespace"] = "kec"
	m_metrics["InstanceID"] = "df2bd78e-8453-4fea-aabf-a0d5a0832ff8"
	m_metrics["PageIndex"] = "1"
	resp, err := svc.ListMetrics(&m_metrics)

	//docker
	/*
		m_metrics := make(map[string]interface{})
		m_metrics["Namespace"] = "kce"
		m_metrics["PageIndex"] = "1"
		m_metrics["Dimensions.0.Name"] = "ClusterId"
		m_metrics["Dimensions.0.Value"] = "b0b69e65-fcc0-4e63-a51c-1b7eb191ec66"
		m_metrics["Dimensions.1.Name"] = "NamespaceName"
		m_metrics["Dimensions.1.Value"] = "kube-system.canal.canal-wvvcj"
		m_metrics["Dimensions.2.Name"] = "DeploymentName"
		m_metrics["Dimensions.2.Value"] = "dsaddada"

		resp, err := monitorSvc.ListMetrics(m_metrics)
	*/
	if err != nil {
		fmt.Println(err.Error())
	}
	if resp != nil {
		str, _ := json.Marshal(&resp)
		fmt.Printf("%+v\n", string(str))
	}

}
