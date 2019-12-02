package main

import (
	"encoding/json"
	"fmt"
	"github.com/KscSDK/ksc-sdk-go/ksc"
	"github.com/KscSDK/ksc-sdk-go/ksc/utils"
	"github.com/KscSDK/ksc-sdk-go/service/monitorv3"
)
//只适用于docker产品线
func main() {
	ak := "ak"
	sk := "sk"
	region := "cn-beijing-6"
	//debug模式的话 打开这个开关
	svc := monitorv3.SdkNew(ksc.NewClient(ak, sk /*true*/), &ksc.Config{Region: &region}, &utils.UrlInfo{
		UseSSL: true,
	})

	//	****************************************获取监控数据(GetMetricStatistics())***************************************************

	m_metrics_g := make(map[string]interface{})
	m_metrics_g["Namespace"] = "kce"
	m_metrics_g["MetricName"] = "pod.network.tx"
	m_metrics_g["StartTime"] = "2019-11-26T14:50:00Z"
	m_metrics_g["EndTime"] = "2019-11-26T15:00:00Z"
	m_metrics_g["Period"] = "180"
	m_metrics_g["Aggregate"] = "Max,min,avg"
	m_metrics_g["Dimensions.0.Name"] = "ClusterId"
	m_metrics_g["Dimensions.0.Value"] = "807a4149-b7e2-4e05-8a35-b77221ce5bb8"
	m_metrics_g["Dimensions.1.Name"] = "NamespaceName"
	m_metrics_g["Dimensions.1.Value"] = "kube-system"
	m_metrics_g["Dimensions.2.Name"] = "WorkloadType"
	m_metrics_g["Dimensions.2.Value"] = "deployment"
	m_metrics_g["Dimensions.3.Name"] = "WorkloadName"
	m_metrics_g["Dimensions.3.Value"] = "system-monitor"
	m_metrics_g["Dimensions.4.Name"] = "PodName"
	m_metrics_g["Dimensions.4.Value"] = "system-monitor-69f6d456bf-r7khs"
	m_metrics_g["Dimensions.5.Name"] = "ContainerName"
	m_metrics_g["Dimensions.5.Value"] = "deployment"

	resp, err := svc.GetMetricStatistics(&m_metrics_g)
	if err != nil {
		fmt.Println(err.Error())
	}
	if resp != nil {
		str, _ := json.Marshal(&resp)
		fmt.Printf("%+v\n", string(str))
	}
	fmt.Println("----------------------")
	//****************************************获取监控项(ListMetrics)***************************************************

	m_metrics_l := make(map[string]interface{})
	m_metrics_l["Namespace"] = "kce"
	m_metrics_l["PageIndex"] = "1"

	m_metrics_l["Dimensions.0.Name"] = "ClusterId"
	m_metrics_l["Dimensions.0.Value"] = "807a4149-b7e2-4e05-8a35-b77221ce5bb8"
	m_metrics_l["Dimensions.1.Name"] = "NamespaceName"
	m_metrics_l["Dimensions.1.Value"] = "kube-system"
	m_metrics_l["Dimensions.2.Name"] = "WorkloadType"
	m_metrics_l["Dimensions.2.Value"] = "deployment"
	m_metrics_l["Dimensions.3.Name"] = "WorkloadName"
	m_metrics_l["Dimensions.3.Value"] = "system-monitor"
	m_metrics_l["Dimensions.4.Name"] = "PodName"
	m_metrics_l["Dimensions.4.Value"] = "system-monitor-69f6d456bf-r7khs"
	m_metrics_l["Dimensions.5.Name"] = "ContainerName"
	m_metrics_l["Dimensions.5.Value"] = "deployment"

	resp, err = svc.ListMetrics(&m_metrics_l)
	if err != nil {
		fmt.Println(err.Error())
	}
	if resp != nil {
		str, _ := json.Marshal(&resp)
		fmt.Printf("%+v\n", string(str))
	}

}
