package main

import (
	"encoding/json"
	"fmt"

	"github.com/KscSDK/ksc-sdk-go/ksc"
	"github.com/KscSDK/ksc-sdk-go/ksc/utils"
	"github.com/KscSDK/ksc-sdk-go/service/kec"
)

func main() {
	ak := "你的ak"
	sk := "你的sk"
	region := "cn-beijing-6"
	//debug模式的话 打开这个开关
	svc := kec.SdkNew(ksc.NewClient(ak, sk /*,true*/), &ksc.Config{Region: &region}, &utils.UrlInfo{
		UseSSL: true,
	})
	var resp *map[string]interface{}
	var err error

	//	****************************************创建实例(RunInstances())***************************************************

	/*runInstances := make(map[string]interface{})
	runInstances["ImageId"] = "b2e78146-58f1-4298-9397-ebf942246a2b"
	runInstances["InstanceType"] = "N2.1A"
	runInstances["DataDiskGb"] = "50"
	runInstances["MaxCount"] = "1"
	runInstances["MinCount"] = "1"
	runInstances["SubnetId"] = "9fd83a5e-33cf-4196-b8df-59fc57478488"
	runInstances["InstancePassword"] = "Qwer@1234"
	runInstances["ChargeType"] = "Daily"
	runInstances["PurchaseTime"] = "1"
	runInstances["SecurityGroupId"] = "31a5484d-8077-4aca-8f79-e093f9d4bbc1"
	runInstances["InstanceName"] = "sdk-test-ebs"
	runInstances["InstanceNameSuffix"] = "1"
	runInstances["SriovNetSupport"] = "false"
	runInstances["ProjectId"] = "208"
	runInstances["MaxCount"] = "1"
	runInstances["MaxCount"] = "1"

	runInstances["DataDisk.1.Type"] = "SSD3.0"
	runInstances["DataDisk.1.Size"] = "10"
	runInstances["DataDisk.1.DeleteWithInstance"] = "true"

	resp, err = svc.RunInstances(&runInstances)
	if err != nil {
		fmt.Println("error:", err.Error())
	}
	if resp != nil {
		str, _ := json.Marshal(&resp)
		fmt.Printf("%+v\n", string(str))
	}*/

	//	****************************************关闭实例(StopInstances())***************************************************

	/*stopInstances := make(map[string]interface{})
	stopInstances["InstanceId.1"] = "6867b808-feea-4416-88bb-361224d82f1b"
	stopInstances["ForceStop"] = "true"

	resp, err = svc.StopInstances(&stopInstances)
	if err != nil {
		fmt.Println("error:", err.Error())
	}
	if resp != nil {
		str, _ := json.Marshal(&resp)
		fmt.Printf("%+v\n", string(str))
	}*/

	//	****************************************开启实例(StartInstances())***************************************************

	/*startInstances := make(map[string]interface{})
	startInstances["InstanceId.1"] = "6867b808-feea-4416-88bb-361224d82f1b"

	resp, err = svc.StartInstances(&startInstances)
	if err != nil {
		fmt.Println("error:", err.Error())
	}
	if resp != nil {
		str, _ := json.Marshal(&resp)
		fmt.Printf("%+v\n", string(str))
	}*/

	//	****************************************重启实例(RebootInstances())***************************************************

	/*rebootInstances := make(map[string]interface{})
	rebootInstances["InstanceId.1"] = "6867b808-feea-4416-88bb-361224d82f1b"

	resp, err = svc.RebootInstances(&rebootInstances)
	if err != nil {
		fmt.Println("error:", err.Error())
	}
	if resp != nil {
		str, _ := json.Marshal(&resp)
		fmt.Printf("%+v\n", string(str))
	}*/

	//	****************************************描述实例信息(DescribeInstances())***************************************************

	/*describeInstances := make(map[string]interface{})
	describeInstances["InstanceId.1"] = "6867b808-feea-4416-88bb-361224d82f1b"
	describeInstances["Marker"] = "0"
	describeInstances["MaxResults"] = "5"

	resp, err = svc.DescribeInstances(nil)
	if err != nil {
		fmt.Println("error:", err.Error())
	}
	if resp != nil {
		str, _ := json.Marshal(&resp)
		fmt.Printf("%+v\n", string(str))
	}*/

	//	****************************************获取VNC信息(DescribeInstanceVnc())***************************************************

	/*describeInstanceVnc := make(map[string]interface{})
	describeInstanceVnc["InstanceId"] = "6867b808-feea-4416-88bb-361224d82f1b"
	resp, err = svc.DescribeInstanceVnc(&describeInstanceVnc)
	if err != nil {
		fmt.Println("error:", err.Error())
	}
	if resp != nil {
		str, _ := json.Marshal(&resp)
		fmt.Printf("%+v\n", string(str))
	}*/

	//	****************************************查询机型套餐配置信息(DescribeInstanceTypeConfigs())***************************************************

	/*describeInstanceTypeConfigs := make(map[string]interface{})
	describeInstanceTypeConfigs["Filter.1.Name"] = "instance-family"
	describeInstanceTypeConfigs["Filter.1.Value.1"] = "N1"
	describeInstanceTypeConfigs["Filter.2.Name"] = "availability-zone"
	describeInstanceTypeConfigs["Filter.2.Value.1"] = "cn-shanghai-3"
	resp, err = svc.DescribeInstanceTypeConfigs(&describeInstanceTypeConfigs)
	if err != nil {
		fmt.Println("error:", err.Error())
	}
	if resp != nil {
		str, _ := json.Marshal(&resp)
		fmt.Printf("%+v\n", string(str))
	}*/

	//	****************************************查询机型配置信息(DescribeInstanceFamilys())***************************************************

	/*resp, err = svc.DescribeInstanceFamilys(nil)
	if err != nil {
		fmt.Println("error:", err.Error())
	}
	if resp != nil {
		str, _ := json.Marshal(&resp)
		fmt.Printf("%+v\n", string(str))
	}*/

	//	****************************************修改实例属性信息(ModifyInstanceAttribute())***************************************************

	/*modifyInstanceAttribute := make(map[string]interface{})
	modifyInstanceAttribute["InstanceId"] = "6867b808-feea-4416-88bb-361224d82f1b"
	modifyInstanceAttribute["InstanceName"] = "go-sdk-test-勿动"
	//modifyInstanceAttribute["InstancePassword"] = "King_123"
	resp, err = svc.ModifyInstanceAttribute(&modifyInstanceAttribute)
	if err != nil {
		fmt.Println("error:", err.Error())
	}
	if resp != nil {
		str, _ := json.Marshal(&resp)
		fmt.Printf("%+v\n", string(str))
	}*/

	//	****************************************升级实例套餐类型(ModifyInstanceType())***************************************************

	/*modifyInstanceType := make(map[string]interface{})
	modifyInstanceType["InstanceId"] = "6867b808-feea-4416-88bb-361224d82f1b"
	modifyInstanceType["InstanceType"] = "N2.1B"
	modifyInstanceType["DataDiskGb"] = "50"
	resp, err = svc.ModifyInstanceType(&modifyInstanceType)
	if err != nil {
		fmt.Println("error:", err.Error())
	}
	if resp != nil {
		str, _ := json.Marshal(&resp)
		fmt.Printf("%+v\n", string(str))
	}*/

	//	****************************************更换或者重新安装实例操作系统(ModifyInstanceImage())***************************************************

	/*modifyInstanceImage := make(map[string]interface{})
	modifyInstanceImage["InstanceId"] = "6867b808-feea-4416-88bb-361224d82f1b"
	modifyInstanceImage["ImageId"] = "e6664c07-39e4-4d6b-9d44-1d3721143584"
	modifyInstanceImage["InstancePassword"] = "Qwer@1234"

	resp, err = svc.ModifyInstanceImage(&modifyInstanceImage)
	if err != nil {
		fmt.Println("error:", err.Error())
	}
	if resp != nil {
		str, _ := json.Marshal(&resp)
		fmt.Printf("%+v\n", string(str))
	}*/

	//	****************************************销毁实例(TerminateInstances())***************************************************

	/*terminateInstances := make(map[string]interface{})
	terminateInstances["InstanceId.1"] = "6867b808-feea-4416-88bb-361224d82f1b"

	resp, err = svc.TerminateInstances(&terminateInstances)
	if err != nil {
		fmt.Println("error:", err.Error())
	}
	if resp != nil {
		str, _ := json.Marshal(&resp)
		fmt.Printf("%+v\n", string(str))
	}*/

	//	****************************************查询本地磁盘(DescribeLocalVolumes())***************************************************

	/*describeLocalVolumes := make(map[string]interface{})
	describeLocalVolumes["InstanceName"] = "test_ARM"

	//resp, err = svc.DescribeLocalVolumes(&describeLocalVolumes)
	resp, err = svc.DescribeLocalVolumes(nil)
	if err != nil {
		fmt.Println("error:", err.Error())
	}
	if resp != nil {
		str, _ := json.Marshal(&resp)
		fmt.Printf("%+v\n", string(str))
	}*/

	//	****************************************查询地域列表(DescribeRegions())***************************************************

	/*resp, err = svc.DescribeRegions(nil)
	if err != nil {
		fmt.Println("error:", err.Error())
	}
	if resp != nil {
		str, _ := json.Marshal(&resp)
		fmt.Printf("%+v\n", string(str))
	}*/

	//	****************************************查询可用区列表(DescribeAvailabilityZones())***************************************************

	/*resp, err = svc.DescribeAvailabilityZones(nil)
	if err != nil {
		fmt.Println("error:", err.Error())
	}
	if resp != nil {
		str, _ := json.Marshal(&resp)
		fmt.Printf("%+v\n", string(str))
	}*/

	//	****************************************为主机添加网卡(AttachNetworkInterface())***************************************************

	/*attachNetworkInterface := make(map[string]interface{})
	attachNetworkInterface["InstanceId"] = "ff80b325-f724-4150-a6ac-bfa42ec8f6bf"
	attachNetworkInterface["SecurityGroupId.1"] = "e7a70523-d5cb-41dd-9ccb-badfe187bf3c"
	attachNetworkInterface["SubnetId"] = "78916229-506b-45ba-8b7f-b00c7f62ead7"

	resp, err = svc.AttachNetworkInterface(&attachNetworkInterface)
	if err != nil {
		fmt.Println("error:", err.Error())
	}
	if resp != nil {
		str, _ := json.Marshal(&resp)
		fmt.Printf("%+v\n", string(str))
	}*/

	//	****************************************修改网络接口属性信息(ModifyNetworkInterfaceAttribute())***************************************************

	/*modifyNetworkInterfaceAttribute := make(map[string]interface{})
	modifyNetworkInterfaceAttribute["InstanceId"] = "ff80b325-f724-4150-a6ac-bfa42ec8f6bf"
	modifyNetworkInterfaceAttribute["SecurityGroupId.1"] = "31a5484d-8077-4aca-8f79-e093f9d4bbc1"
	modifyNetworkInterfaceAttribute["SubnetId"] = "9fd83a5e-33cf-4196-b8df-59fc57478488"
	modifyNetworkInterfaceAttribute["NetworkInterfaceId"] = "c27e5675-8e95-4308-94e7-8a6fdb350ac0"

	resp, err = svc.ModifyNetworkInterfaceAttribute(&modifyNetworkInterfaceAttribute)
	if err != nil {
		fmt.Println("error:", err.Error())
	}
	if resp != nil {
		str, _ := json.Marshal(&resp)
		fmt.Printf("%+v\n", string(str))
	}*/

	//	****************************************删除主机网络接口(DetachNetworkInterface())***************************************************

	/*detachNetworkInterface := make(map[string]interface{})
	detachNetworkInterface["InstanceId"] = "ff80b325-f724-4150-a6ac-bfa42ec8f6bf"
	detachNetworkInterface["NetworkInterfaceId"] = "866fca0d-02f6-4b02-946b-1e434ae9024c"

	resp, err = svc.DetachNetworkInterface(&detachNetworkInterface)
	if err != nil {
		fmt.Println("error:", err.Error())
	}
	if resp != nil {
		str, _ := json.Marshal(&resp)
		fmt.Printf("%+v\n", string(str))
	}*/

	//	****************************************创建本地盘快照(CreateLocalVolumeSnapshot())***************************************************

	/*createLocalVolumeSnapshot := make(map[string]interface{})
	createLocalVolumeSnapshot["LocalVolumeId"] = "ff80b325-f724-4150-a6ac-bfa42ec8f6bf-a"
	createLocalVolumeSnapshot["LocalVolumeSnapshotName"] = "MyTestSnapshot"

	resp, err = svc.CreateLocalVolumeSnapshot(&createLocalVolumeSnapshot)
	if err != nil {
		fmt.Println("error:", err.Error())
	}
	if resp != nil {
		str, _ := json.Marshal(&resp)
		fmt.Printf("%+v\n", string(str))
	}*/

	//	****************************************回滚本地盘快照(RollbackLocalVolume())***************************************************

	/*rollbackLocalVolume := make(map[string]interface{})
	rollbackLocalVolume["LocalVolumeId"] = "ff80b325-f724-4150-a6ac-bfa42ec8f6bf-a"
	rollbackLocalVolume["LocalVolumeSnapshotId"] = "f7b3407a-7a4a-4810-9b1d-edf36bf384ca"

	resp, err = svc.RollbackLocalVolume(&rollbackLocalVolume)
	if err != nil {
		fmt.Println("error:", err.Error())
	}
	if resp != nil {
		str, _ := json.Marshal(&resp)
		fmt.Printf("%+v\n", string(str))
	}*/

	//	****************************************查询本地盘快照(DescribeLocalVolumeSnapshots())***************************************************

	/*describeLocalVolumeSnapshots := make(map[string]interface{})
	describeLocalVolumeSnapshots["LocalVolumeName"] = "ksc-ff80b325-f724-4150-a6ac-bfa42ec8f6bf-vda"
	describeLocalVolumeSnapshots["InstanceName"] = "test"

	//resp, err = svc.DescribeLocalVolumeSnapshots(&describeLocalVolumeSnapshots)
	resp, err = svc.DescribeLocalVolumeSnapshots(nil)
	if err != nil {
		fmt.Println("error:", err.Error())
	}
	if resp != nil {
		str, _ := json.Marshal(&resp)
		fmt.Printf("%+v\n", string(str))
	}*/

	//	****************************************删除本地盘快照(DeleteLocalVolumeSnapshot())***************************************************

	/*deleteLocalVolumeSnapshot := make(map[string]interface{})
	deleteLocalVolumeSnapshot["LocalVolumeSnapshotId.1"] = "f7b3407a-7a4a-4810-9b1d-edf36bf384ca"

	resp, err = svc.DeleteLocalVolumeSnapshot(&deleteLocalVolumeSnapshot)
	if err != nil {
		fmt.Println("error:", err.Error())
	}
	if resp != nil {
		str, _ := json.Marshal(&resp)
		fmt.Printf("%+v\n", string(str))
	}*/

	//	****************************************主机绑定密钥(AttachKey())***************************************************

	/*attachKey := make(map[string]interface{})
	attachKey["InstanceId.1"] = "ff80b325-f724-4150-a6ac-bfa42ec8f6bf"
	attachKey["KeyId.1"] = "40b9a22f-effa-42f8-8a55-2b4f2e3652b9"

	resp, err = svc.AttachKey(&attachKey)
	if err != nil {
		fmt.Println("error:", err.Error())
	}
	if resp != nil {
		str, _ := json.Marshal(&resp)
		fmt.Printf("%+v\n", string(str))
	}*/

	//	****************************************主机解绑密钥(DetachKey())***************************************************

	/*detachKey := make(map[string]interface{})
	detachKey["InstanceId.1"] = "ff80b325-f724-4150-a6ac-bfa42ec8f6bf"
	detachKey["KeyId.1"] = "40b9a22f-effa-42f8-8a55-2b4f2e3652b9"

	resp, err = svc.DetachKey(&detachKey)
	if err != nil {
		fmt.Println("error:", err.Error())
	}
	if resp != nil {
		str, _ := json.Marshal(&resp)
		fmt.Printf("%+v\n", string(str))
	}*/

	//	****************************************创建镜像(CreateImage())***************************************************

	/*createImage := make(map[string]interface{})
	createImage["InstanceId"] = "ff80b325-f724-4150-a6ac-bfa42ec8f6bf"
	createImage["Name"] = "test_SDK_image"

	resp, err = svc.CreateImage(&createImage)
	if err != nil {
		fmt.Println("error:", err.Error())
	}
	if resp != nil {
		str, _ := json.Marshal(&resp)
		fmt.Printf("%+v\n", string(str))
	}*/

	//	****************************************描述镜像信息(DescribeImages())***************************************************

	/*describeImages := make(map[string]interface{})
	describeImages["ImageId"] = "9139a745-d077-4df5-b66b-45ec04de5d05"

	resp, err = svc.DescribeImages(&describeImages)
	if err != nil {
		fmt.Println("error:", err.Error())
	}
	if resp != nil {
		str, _ := json.Marshal(&resp)
		fmt.Printf("%+v\n", string(str))
	}*/

	//	****************************************删除镜像(RemoveImages())***************************************************

	/*removeImages := make(map[string]interface{})
	removeImages["ImageId"] = "f203b92d-956c-4fa4-9609-f5591f33d557"

	resp, err = svc.RemoveImages(&removeImages)
	if err != nil {
		fmt.Println("error:", err.Error())
	}
	if resp != nil {
		str, _ := json.Marshal(&resp)
		fmt.Printf("%+v\n", string(str))
	}*/

	//	****************************************修改镜像属性信息(ModifyImageAttribute())***************************************************

	/*modifyImageAttribute := make(map[string]interface{})
	modifyImageAttribute["ImageId"] = "9139a745-d077-4df5-b66b-45ec04de5d05"
	modifyImageAttribute["Name"] = "SDK_image"

	resp, err = svc.ModifyImageAttribute(&modifyImageAttribute)
	if err != nil {
		fmt.Println("error:", err.Error())
	}
	if resp != nil {
		str, _ := json.Marshal(&resp)
		fmt.Printf("%+v\n", string(str))
	}*/

	//	****************************************镜像导入(ImportImage())***************************************************

	/*importImage := make(map[string]interface{})
	importImage["ImageName"] = "测试导入镜像"
	importImage["Architecture"] = "x86_64"
	importImage["Platform"] = "centos-6"
	importImage["ImageUrl"] = "http://mybucket.ks3-cn-beijing.ksyun.com/centos6_5_64.raw"
	importImage["ImageFormat"] = "raw"

	resp, err = svc.ImportImage(&importImage)
	if err != nil {
		fmt.Println("error:", err.Error())
	}
	if resp != nil {
		str, _ := json.Marshal(&resp)
		fmt.Printf("%+v\n", string(str))
	}*/

	//	****************************************镜像复制(CopyImage())***************************************************

	/*copyImage := make(map[string]interface{})
	copyImage["ImageId.1"] = "7753e731-ea67-45c0-9745-ffa19c43c7d5"
	copyImage["DestinationRegion.1"] = "cn-shanghai-2"

	resp, err = svc.CopyImage(&copyImage)
	if err != nil {
		fmt.Println("error:", err.Error())
	}
	if resp != nil {
		str, _ := json.Marshal(&resp)
		fmt.Printf("%+v\n", string(str))
	}*/

	//	****************************************查看镜像分享信息(DescribeImageSharePermission())***************************************************

	/*describeImageSharePermission := make(map[string]interface{})
	describeImageSharePermission["ImageId"] = "9139a745-d077-4df5-b66b-45ec04de5d05"

	resp, err = svc.DescribeImageSharePermission(&describeImageSharePermission)
	if err != nil {
		fmt.Println("error:", err.Error())
	}
	if resp != nil {
		str, _ := json.Marshal(&resp)
		fmt.Printf("%+v\n", string(str))
	}*/

	//	****************************************修改镜像分享信息(ModifyImageSharePermission())***************************************************

	/*modifyImageSharePermission := make(map[string]interface{})
	modifyImageSharePermission["ImageId"] = "9139a745-d077-4df5-b66b-45ec04de5d05"
	modifyImageSharePermission["AccountId.1"] = "73400974"
	modifyImageSharePermission["Permission"] = "share"

	resp, err = svc.ModifyImageSharePermission(&modifyImageSharePermission)
	if err != nil {
		fmt.Println("error:", err.Error())
	}
	if resp != nil {
		str, _ := json.Marshal(&resp)
		fmt.Printf("%+v\n", string(str))
	}*/
}
