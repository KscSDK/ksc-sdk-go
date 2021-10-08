package main
import (
	"encoding/json"
	"fmt"
	"github.com/KscSDK/ksc-sdk-go/ksc"
	"github.com/KscSDK/ksc-sdk-go/ksc/utils"
	"github.com/KscSDK/ksc-sdk-go/service/kcev2"
)

func main() {
	ak := "用户AK"
	sk := "用户SK"
	region := "cn-beijing-6"
	//debug模式的话 打开这个开关
	svc := kcev2.SdkNew(ksc.NewClient(ak, sk /*,true*/), &ksc.Config{Region: &region}, &utils.UrlInfo{
		UseSSL:      true,
		UseInternal: false,
	})
	var resp *map[string]interface{}
	var err error

	//	****************************************创建集群(CreateCluster())***************************************************
		createCluster := map[string]interface{}{
			"ClusterName": "xxxxx",
			"ClusterManageMode": "DedicatedCluster",
			"VpcId": "15552848-026b-4ad6-a3e3-xxxxx",
			"PodCidr": "10.32.0.0/14",
			"ServiceCidr": "10.254.0.0/16",
			"NetworkType": "Flannel",
			"K8sVersion": "v1.19.3",
			"ReserveSubnetId": "9729f4c0-93ee-4e2a-9b2a-xxxxx",
			"PublicApiServer": "{\"LineId\":\"5fc2595f-1bfd-481b-bf64-2d08f116d800\",\"BandWidth\": \"10\",\"ChargeType\": \"PostPaidByDay\"}",
			"InstanceForNode.1.NodeRole": "Master_Etcd",
			"InstanceForNode.1.NodeConfig.1.Para": "{\"MaxCount\":3,\"MinCount\":3,\"ImageId\":\"a0699172-76c6-4885-a4e9-0799a9cb811d\",\"SubnetId\":\"4a077fa8-a239-47bf-a18f-xxxxx\",\"InstancePassword\":\"Root23123\",\"SecurityGroupId\":\"0dcd8356-9e7e-43ae-897b-xxxxx\",\"DataDiskGb\":0,\"ChargeType\":\"Daily\",\"InstanceType\":\"I3.2A\",\"PurchaseTime\":0,\"InstanceName\":\"kce-py\",\"InstanceNameSuffix\":1}",
			"InstanceForNode.1.NodeConfig.1.AdvancedSetting.DockerPath": "/data1/docker",
			"InstanceForNode.1.NodeConfig.1.AdvancedSetting.DataDisk.FileSystem": "ext3",
			"InstanceForNode.1.NodeConfig.1.AdvancedSetting.DataDisk.MountTarget": "/data1",
			"InstanceForNode.2.NodeRole": "Worker",
			"InstanceForNode.2.NodeConfig.1.Para": "{\"MaxCount\":1,\"MinCount\":1,\"ImageId\":\"a0699172-76c6-4885-a4e9-0799a9cb811d\",\"SubnetId\":\"4a077fa8-a239-47bf-a18f-xxxxx\",\"InstancePassword\":\"Root23123\",\"SecurityGroupId\":\"0dcd8356-9e7e-43ae-897b-xxxxx\",\"DataDiskGb\":0,\"ChargeType\":\"Daily\",\"InstanceType\":\"I3.2A\",\"PurchaseTime\":0,\"InstanceName\":\"kce-py\",\"InstanceNameSuffix\":1}",
			"InstanceForNode.2.NodeConfig.1.AdvancedSetting.DockerPath": "/data1/docker",
			"InstanceForNode.2.NodeConfig.1.AdvancedSetting.Label.1.Key": "label",
			"InstanceForNode.2.NodeConfig.1.AdvancedSetting.Label.1.Value": "worker",
			"InstanceForNode.2.NodeConfig.1.AdvancedSetting.DataDisk.FileSystem": "ext4",
			"InstanceForNode.2.NodeConfig.1.AdvancedSetting.DataDisk.MountTarget": "/data1",
		}
		resp, err = svc.CreateCluster(&createCluster)
		if err != nil {
			fmt.Println("error:", err.Error())
		}
		if resp != nil {
			str, _ := json.Marshal(&resp)
			fmt.Printf("%+v\n", string(str))
		}
}