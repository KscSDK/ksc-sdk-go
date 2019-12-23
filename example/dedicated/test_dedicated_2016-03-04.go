package main

import (
	"encoding/json"
	"fmt"

	"github.com/KscSDK/ksc-sdk-go/ksc"
	"github.com/KscSDK/ksc-sdk-go/ksc/utils"
	"github.com/KscSDK/ksc-sdk-go/service/dedicated"
)

func main() {
	ak := "你的ak"
	sk := "你的sk"
	region := "cn-beijing-6"
	//debug模式的话 打开这个开关
	svc := dedicated.SdkNew(ksc.NewClient(ak, sk /*,true*/), &ksc.Config{Region: &region}, &utils.UrlInfo{
		UseSSL: true,
	})
	var resp *map[string]interface{}
	var err error

	//	****************************************创建专属虚机实例(RunInstances())***************************************************

	/*runInstances := make(map[string]interface{})
	runInstances["ImageId"] = "e00e5b9a-0914-4611-8730-3768bf08e035"
	runInstances["InstanceType"] = "DVM2.NONE"
	runInstances["SystemDisk.DiskSize"] = "20"
	runInstances["MaxCount"] = "1"
	runInstances["MinCount"] = "1"
	runInstances["SubnetId"] = "07e371be-f459-4c95-83c6-d2f03cc93d5a"
	runInstances["InstancePassword"] = "King@1234"
	runInstances["ChargeType"] = "Daily"
	runInstances["PurchaseTime"] = "1"
	runInstances["SecurityGroupId"] = "00806c62-ec68-4500-a5d2-9e8f9fe4d40a"
	runInstances["InstanceName"] = "sdk-test-dedicated"
	runInstances["InstanceNameSuffix"] = "1"
	runInstances["SriovNetSupport"] = "false"
	runInstances["ProjectId"] = "208"
	runInstances["InstanceConfigure.VCPU"] = "1"
	runInstances["InstanceConfigure.MemoryGb"] = "1"
	runInstances["InstanceConfigure.DataDiskGb"] = "10"
	runInstances["DedicatedHostId"] = "9bf5aa26-4aeb-414c-9208-c1061d6ba3de"

	resp, err = svc.RunInstances(&runInstances)
	if err != nil {
		fmt.Println("error:", err.Error())
	}
	if resp != nil {
		str, _ := json.Marshal(&resp)
		fmt.Printf("%+v\n", string(str))
	}*/

	//	****************************************创建专属宿主机(CreateDedicatedHosts())***************************************************

	/*createDedicatedHosts := make(map[string]interface{})
	createDedicatedHosts["DedicatedType"] = "DS2"
	createDedicatedHosts["Number"] = "1"
	createDedicatedHosts["Name"] = "test"
	createDedicatedHosts["ChargeType"] = "Daily"
	createDedicatedHosts["AvailabilityZone"] = "cn-shanghai-3b"

	resp, err = svc.CreateDedicatedHosts(&createDedicatedHosts)
	if err != nil {
		fmt.Println("error:", err.Error())
	}
	if resp != nil {
		str, _ := json.Marshal(&resp)
		fmt.Printf("%+v\n", string(str))
	}*/

	//	****************************************描述专属宿主机(DescribeDedicatedHosts())***************************************************

	/*describeDedicatedHosts := make(map[string]interface{})
	describeDedicatedHosts["DedicatedHostId"] = "1e804026-78ef-4a0a-b947-40d74ff7ca5d"

	resp, err = svc.DescribeDedicatedHosts(&describeDedicatedHosts)
	if err != nil {
		fmt.Println("error:", err.Error())
	}
	if resp != nil {
		str, _ := json.Marshal(&resp)
		fmt.Printf("%+v\n", string(str))
	}*/

	//	****************************************修改专属宿主机(RenameDedicatedHost())***************************************************

	/*renameDedicatedHost := make(map[string]interface{})
	renameDedicatedHost["DedicatedHostId"] = "1e804026-78ef-4a0a-b947-40d74ff7ca5d"
	renameDedicatedHost["NewDedicatedHostName"] = "wyn_test"

	resp, err = svc.RenameDedicatedHost(&renameDedicatedHost)
	if err != nil {
		fmt.Println("error:", err.Error())
	}
	if resp != nil {
		str, _ := json.Marshal(&resp)
		fmt.Printf("%+v\n", string(str))
	}*/

	//	****************************************设置专属宿主机虚拟核数(SetvCPU())***************************************************

	/*setvCPU := make(map[string]interface{})
	setvCPU["DedicatedHostId.1"] = "1e804026-78ef-4a0a-b947-40d74ff7ca5d"
	setvCPU["VCPU"] = "20"

	resp, err = svc.SetvCPU(&setvCPU)
	if err != nil {
		fmt.Println("error:", err.Error())
	}
	if resp != nil {
		str, _ := json.Marshal(&resp)
		fmt.Printf("%+v\n", string(str))
	}*/

	//	****************************************专属虚拟机迁移(InstanceMigrate())***************************************************

	/*instanceMigrate := make(map[string]interface{})
	instanceMigrate["DedicatedHostId"] = "1e804026-78ef-4a0a-b947-40d74ff7ca5d"
	instanceMigrate["InstanceId"] = "6867b808-feea-4416-88bb-361224d82f1b"

	resp, err = svc.InstanceMigrate(&instanceMigrate)
	if err != nil {
		fmt.Println("error:", err.Error())
	}
	if resp != nil {
		str, _ := json.Marshal(&resp)
		fmt.Printf("%+v\n", string(str))
	}*/

	//	****************************************修改专属宿主机所在集群(DedicatedHostMigrate())***************************************************

	/*dedicatedHostMigrate := make(map[string]interface{})
	dedicatedHostMigrate["DedicatedClusterId"] = "49c496f9-b836-40c7-8914-204ec860d30b"
	dedicatedHostMigrate["DedicatedHostId.1"] = "1e804026-78ef-4a0a-b947-40d74ff7ca5d"

	resp, err = svc.DedicatedHostMigrate(&dedicatedHostMigrate)
	if err != nil {
		fmt.Println("error:", err.Error())
	}
	if resp != nil {
		str, _ := json.Marshal(&resp)
		fmt.Printf("%+v\n", string(str))
	}*/

	//	****************************************创建专属集群(CreateDedicatedCluster())***************************************************

	/*createDedicatedCluster := make(map[string]interface{})
	createDedicatedCluster["DedicatedClusterName"] = "业务B-DN2"
	createDedicatedCluster["Model"] = "DN2"
	createDedicatedCluster["AvailabilityZone"] = "cn-shanghai-3b"

	resp, err = svc.CreateDedicatedCluster(&createDedicatedCluster)
	if err != nil {
		fmt.Println("error:", err.Error())
	}
	if resp != nil {
		str, _ := json.Marshal(&resp)
		fmt.Printf("%+v\n", string(str))
	}*/

	//	****************************************删除专属集群息(DeleteDedicatedCluster())***************************************************

	/*deleteDedicatedCluster := make(map[string]interface{})
	deleteDedicatedCluster["DedicatedClusterId.1"] = "f0627952-1282-419a-bc11-a25553ecc2c5"

	resp, err = svc.DeleteDedicatedCluster(&deleteDedicatedCluster)
	if err != nil {
		fmt.Println("error:", err.Error())
	}
	if resp != nil {
		str, _ := json.Marshal(&resp)
		fmt.Printf("%+v\n", string(str))
	}*/

	//	****************************************查看专属集群(DescribeDedicatedCluster())***************************************************

	/*describeDedicatedCluster := make(map[string]interface{})
	describeDedicatedCluster["DedicatedClusterId.1"] = "6867b808-feea-4416-88bb-361224d82f1b"

	resp, err = svc.DescribeDedicatedCluster(&describeDedicatedCluster)
	if err != nil {
		fmt.Println("error:", err.Error())
	}
	if resp != nil {
		str, _ := json.Marshal(&resp)
		fmt.Printf("%+v\n", string(str))
	}*/

	//	****************************************修改专属集群(ModifyDedicatedClusterName())***************************************************

	/*modifyDedicatedClusterName := make(map[string]interface{})
	modifyDedicatedClusterName["DedicatedClusterId"] = "6867b808-feea-4416-88bb-361224d82f1b"
	modifyDedicatedClusterName["DedicatedClusterName"] = "test"


	resp, err = svc.ModifyDedicatedClusterName(&modifyDedicatedClusterName)
	if err != nil {
		fmt.Println("error:", err.Error())
	}
	if resp != nil {
		str, _ := json.Marshal(&resp)
		fmt.Printf("%+v\n", string(str))
	}*/

	//	****************************************删除专属宿主机(DeleteDedicatedHost())***************************************************

	/*deleteDedicatedHost := make(map[string]interface{})
	deleteDedicatedHost["DedicatedHostId"] = "6867b808-feea-4416-88bb-361224d82f1b"

	resp, err = svc.DeleteDedicatedHost(&deleteDedicatedHost)
	if err != nil {
		fmt.Println("error:", err.Error())
	}
	if resp != nil {
		str, _ := json.Marshal(&resp)
		fmt.Printf("%+v\n", string(str))
	}*/

}
