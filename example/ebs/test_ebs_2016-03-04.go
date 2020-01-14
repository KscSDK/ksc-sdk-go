package main

import (
	"encoding/json"
	"fmt"
	"github.com/KscSDK/ksc-sdk-go/ksc"
	"github.com/KscSDK/ksc-sdk-go/ksc/utils"
	"github.com/KscSDK/ksc-sdk-go/service/ebs"
)

func main()  {
    ak := "你的AK"
	sk := "你的SK"
	region := "cn-beijing-6"
	//debug模式的话 打开这个开关
	svc := ebs.SdkNew(ksc.NewClient(ak, sk /*,true*/), &ksc.Config{Region: &region}, &utils.UrlInfo{
		UseSSL: false,
	})
	var resp *map[string]interface{}
	var err error

    //	****************************************查询云盘(DescribeVolumes())***************************************************
	/*describeVolumes := make(map[string]interface{})
	describeVolumes["VolumeType"] = "SSD3.0"//SSD2.0/SSD3.0/SATA2.0/EHDD
	describeVolumes["VolumeCategory"] = "data"//data/system
	describeVolumes["VolumeStatus"] = "available"//creating|available|in-use|error|recycling
	describeVolumes["VolumeCreateDate"] = "2019-01-01"//yyyy-MM-dd
	describeVolumes["AvailabilityZone"] = "cn-beijing-6a"
	describeVolumes["VolumeId.1"] = "12b66795-f5ea-41a7-a211-2dda4f4511bf"

	resp, err = svc.DescribeVolumes(&describeVolumes)
	if err != nil {
		fmt.Println("error:", err.Error())
	}
	if resp != nil {
		str, _ := json.Marshal(&resp)
		fmt.Printf("%+v\n", string(str))
	}*/

	//	****************************************创建云盘(CreateVolume())***************************************************
	/*createVolume := make(map[string]interface{})
	createVolume["VolumeName"] = "test"
	createVolume["VolumeType"] = "SSD3.0"//必传 SSD2.0/SSD3.0/SATA2.0/EHDD
	createVolume["VolumeDesc"] = "test"
	createVolume["Size"] = "100"//必传 GB
	createVolume["ChargeType"] = "HourlyInstantSettlement"//必传
	createVolume["AvailabilityZone"] = "cn-beijing-6a"//必传
	createVolume["ProjectId"] = "0"

	resp, err = svc.CreateVolume(&createVolume)
	if err != nil {
		fmt.Println("error:", err.Error())
	}
	if resp != nil {
		str, _ := json.Marshal(&resp)
		fmt.Printf("%+v\n", string(str))
	}*/

	//	****************************************查询可挂载主机(DescribeEbsInstances())***************************************************
	/*describeEbsInstances := make(map[string]interface{})
	describeEbsInstances["AvailabilityZone"] = "cn-beijing-6a"//必传
	describeEbsInstances["VolumeType"] = "SSD3.0"//必传

	resp, err = svc.DescribeEbsInstances(&describeEbsInstances)
	if err != nil {
		fmt.Println("error:", err.Error())
	}
	if resp != nil {
		str, _ := json.Marshal(&resp)
		fmt.Printf("%+v\n", string(str))
	}*/

	//	****************************************校验挂载主机(ValidateAttachInstance())***************************************************
	/*validateAttachInstance := make(map[string]interface{})
	validateAttachInstance["VolumeType"] = "SSD3.0"//必传
	validateAttachInstance["InstanceId"] = "12b66795-f5ea-41a7-a211-2dda4f4511bf"//必传 主机ID

	resp, err = svc.ValidateAttachInstance(&validateAttachInstance)
	if err != nil {
		fmt.Println("error:", err.Error())
	}
	if resp != nil {
		str, _ := json.Marshal(&resp)
		fmt.Printf("%+v\n", string(str))
	}*/

	//	****************************************挂载云盘(AttachVolume())***************************************************
	/*attachVolume := make(map[string]interface{})
	attachVolume["VolumeId"] = "799e479d-e0a0-40ec-ad97-8f91babb4f26"//必传
	attachVolume["InstanceId"] = "12b66795-f5ea-41a7-a211-2dda4f4511bf"//必传
	attachVolume["DeleteWithInstance"] = "true"

	resp, err = svc.AttachVolume(&attachVolume)
	if err != nil {
		fmt.Println("error:", err.Error())
	}
	if resp != nil {
		str, _ := json.Marshal(&resp)
		fmt.Printf("%+v\n", string(str))
	}*/

	//	****************************************卸载云盘(DetachVolume())***************************************************
	/*detachVolume := make(map[string]interface{})
	detachVolume["VolumeId"] = "799e479d-e0a0-40ec-ad97-8f91babb4f26"//必传

	resp, err = svc.DetachVolume(&detachVolume)
	if err != nil {
		fmt.Println("error:", err.Error())
	}
	if resp != nil {
		str, _ := json.Marshal(&resp)
		fmt.Printf("%+v\n", string(str))
	}*/

	//	****************************************扩容云盘(ResizeVolume())***************************************************
	/*resizeVolume := make(map[string]interface{})
	resizeVolume["VolumeId"] = "799e479d-e0a0-40ec-ad97-8f91babb4f26"//必传
	resizeVolume["Size"] = "200"//必传

	resp, err = svc.ResizeVolume(&resizeVolume)
	if err != nil {
		fmt.Println("error:", err.Error())
	}
	if resp != nil {
		str, _ := json.Marshal(&resp)
		fmt.Printf("%+v\n", string(str))
	}*/

	//	****************************************修改云盘属性(ModifyVolume())***************************************************
	/*modifyVolume := make(map[string]interface{})
	modifyVolume["VolumeId"] = "799e479d-e0a0-40ec-ad97-8f91babb4f26"//必传
	modifyVolume["VolumeName"] = "test"
	modifyVolume["VolumeDesc"] = "test"
	modifyVolume["DeleteWithInstance"] = "true"

	resp, err = svc.ModifyVolume(&modifyVolume)
	if err != nil {
		fmt.Println("error:", err.Error())
	}
	if resp != nil {
		str, _ := json.Marshal(&resp)
		fmt.Printf("%+v\n", string(str))
	}*/

	//	****************************************修改云盘项目组(UpdateVolumeProject())***************************************************
	/*updateVolumeProject := make(map[string]interface{})
	updateVolumeProject["VolumeId.1"] = "799e479d-e0a0-40ec-ad97-8f91babb4f26"//必传
	updateVolumeProject["ProjectId"] = "2"//必传

	resp, err = svc.UpdateVolumeProject(&updateVolumeProject)
	if err != nil {
		fmt.Println("error:", err.Error())
	}
	if resp != nil {
		str, _ := json.Marshal(&resp)
		fmt.Printf("%+v\n", string(str))
	}*/

	//	****************************************查询主机下云盘(DescribeInstanceVolumes())***************************************************
	/*describeInstanceVolumes := make(map[string]interface{})
	describeInstanceVolumes["InstanceId"] = "799e479d-e0a0-40ec-ad97-8f91babb4f26"//必传

	resp, err = svc.DescribeInstanceVolumes(&describeInstanceVolumes)
	if err != nil {
		fmt.Println("error:", err.Error())
	}
	if resp != nil {
		str, _ := json.Marshal(&resp)
		fmt.Printf("%+v\n", string(str))
	}*/

	//	****************************************删除云盘(DeleteVolume())***************************************************
	/*deleteVolume := make(map[string]interface{})
	deleteVolume["VolumeId"] = "799e479d-e0a0-40ec-ad97-8f91babb4f26"//必传
	deleteVolume["ForceDelete"] = "false"

	resp, err = svc.DeleteVolume(&deleteVolume)
	if err != nil {
		fmt.Println("error:", err.Error())
	}
	if resp != nil {
		str, _ := json.Marshal(&resp)
		fmt.Printf("%+v\n", string(str))
	}*/

	//	****************************************恢复云盘(RecoveryVolume())***************************************************
	/*recoveryVolume := make(map[string]interface{})
	recoveryVolume["VolumeId"] = "799e479d-e0a0-40ec-ad97-8f91babb4f26"//必传

	resp, err = svc.RecoveryVolume(&recoveryVolume)
	if err != nil {
		fmt.Println("error:", err.Error())
	}
	if resp != nil {
		str, _ := json.Marshal(&resp)
		fmt.Printf("%+v\n", string(str))
	}*/

}
