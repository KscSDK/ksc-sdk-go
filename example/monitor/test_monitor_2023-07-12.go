package main

import (
	"encoding/json"
	"fmt"

	"github.com/KscSDK/ksc-sdk-go/ksc"
	"github.com/KscSDK/ksc-sdk-go/ksc/utils"
	"github.com/KscSDK/ksc-sdk-go/service/monitorv5"
)

func main() {
	ak := "ak"
	sk := "sk"
	region := "cn-beijing-6"
	//debug模式的话 打开这个开关
	svc := monitorv5.SdkNew(ksc.NewClient(ak, sk /*true*/), &ksc.Config{Region: &region}, &utils.UrlInfo{
		UseSSL: true,
	})
	var resp *map[string]interface{}
	var err error

	//	****************************************获取监控实例(DescribeInstances)***************************************************

	m_instances := make(map[string]interface{})
	m_instances["Namespace"] = "kcm"
	m_instances["Dimensions.0.Name"] = "cluster"
	m_instances["Dimensions.0.Value"] = "mock-cluster-01"
	resp, err = svc.DescribeInstances(&m_instances)
	if err != nil {
		fmt.Println(err.Error())
	}
	if resp != nil {
		str, _ := json.Marshal(&resp)
		fmt.Printf("%+v\n", string(str))
	}

	//	****************************************获取监控项(ListMetrics)***************************************************

	m_metrics_l := make(map[string]interface{})
	m_metrics_l["Namespace"] = "kcm"
	m_metrics_l["InstanceID"] = "127.0.0.1:9100"
	resp, err = svc.ListMetrics(&m_metrics_l)
	if err != nil {
		fmt.Println(err.Error())
	}
	if resp != nil {
		str, _ := json.Marshal(&resp)
		fmt.Printf("%+v\n", string(str))
	}

	//	****************************************获取监控数据(GetMetricStatistics)***************************************************

	m_metrics_g := make(map[string]interface{})
	m_metrics_g["Namespace"] = "kcm"
	m_metrics_g["InstanceID"] = "127.0.0.1:9100"
	m_metrics_g["MetricName"] = "node_cpu_seconds_total"
	m_metrics_g["StartTime"] = "2023-07-12T16:20:00Z"
	m_metrics_g["EndTime"] = "2023-07-12T16:30:00Z"
	m_metrics_g["Period"] = "120"
	m_metrics_g["Aggregate"] = "Max,Min,Avg,Sum,Count"
	m_metrics_g["Dimensions.0.Name"] = "cluster"
	m_metrics_g["Dimensions.0.Value"] = "mock-cluster-01"
	m_metrics_g["Dimensions.1.Name"] = "cpu"
	m_metrics_g["Dimensions.1.Value"] = "0"
	m_metrics_g["Dimensions.2.Name"] = "mode"
	m_metrics_g["Dimensions.2.Value"] = "user"

	resp, err = svc.GetMetricStatistics(&m_metrics_g)
	if err != nil {
		fmt.Println(err.Error())
	}
	if resp != nil {
		str, _ := json.Marshal(&resp)
		fmt.Printf("%+v\n", string(str))
	}

	//	****************************************批量获取监控数据(GetMetricStatisticsBatch)***************************************************

	m_metrics_cpu := make(map[string]interface{})
	m_metrics_cpu["InstanceID"] = "127.0.0.1:9100"
	m_metrics_cpu["MetricName"] = "node_cpu_seconds_total"
	m_metrics_cpu["Dimensions"] = []map[string]string{
		{
			"Name":  "cluster",
			"Value": "mock-cluster-01",
		},
		{
			"Name":  "cpu",
			"Value": "0",
		},
		{
			"Name":  "mode",
			"Value": "user",
		},
	}

	m_metrics_mem := make(map[string]interface{})
	m_metrics_mem["InstanceID"] = "127.0.0.1:9100"
	m_metrics_mem["MetricName"] = "node_memory_MemFree_bytes"

	l1 := []interface{}{m_metrics_cpu, m_metrics_mem}

	m_metrics_b := make(map[string]interface{})
	m_metrics_b["Namespace"] = "kcm"
	m_metrics_b["Metrics"] = l1
	m_metrics_b["StartTime"] = "2023-07-12T16:20:00Z"
	m_metrics_b["EndTime"] = "2023-07-12T16:30:00Z"
	m_metrics_b["Period"] = 120
	m_metrics_b["Aggregate"] = []string{"Max", "Min", "Avg", "Sum", "Count"}

	resp, err = svc.GetMetricStatisticsBatch(&m_metrics_b)
	if err != nil {
		fmt.Println(err.Error())
	}
	if resp != nil {
		str, _ := json.Marshal(&resp)
		fmt.Printf("%+v\n", string(str))
	}
}
