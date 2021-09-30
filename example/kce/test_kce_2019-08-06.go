package main

import (
	"encoding/json"
	"fmt"
	"github.com/KscSDK/ksc-sdk-go/ksc"
	"github.com/KscSDK/ksc-sdk-go/ksc/utils"
	"github.com/KscSDK/ksc-sdk-go/service/kce"
)

func main() {
	ak := "用户AK"
	sk := "用户SK"
	region := "cn-beijing-6"
	//debug模式的话 打开这个开关
	svc := kce.SdkNew(ksc.NewClient(ak, sk /*,true*/), &ksc.Config{Region: &region}, &utils.UrlInfo{
		UseSSL:      true,
		UseInternal: false,
	})
	var resp *map[string]interface{}
	var err error
	//	****************************************创建节点池(CreateNodePool())***************************************************
/*	createNodePool := map[string]interface{}{
        "NodePoolName": "xuan-create",
        "ClusterId": "d9a0adf0-a8f3-xxx-xxxxx",
        "EnableAutoScale": true,
        "MinSize": 0,
        "DesiredCapacity": 0,
        "MaxSize": 2,
        "Label.0.Key": "labelk",
        "Label.0.Value": "labelv",
        "Taint.0.Key": "taintk",
        "Taint.0.Value": "taintv",
        "Taint.0.Effect": "NoSchedule",
        "NodeTemplate.ImageId": "13107fc5-0dd8-xxx-xxxxx",
        //"NodeTemplate.KeyId.1":"c409d431-xxx-xxx-9834-xxxxx",
        "NodeTemplate.DataDiskGb": 0,
        "NodeTemplate.SubnetId.1": "2c83f2d6-8160-xxx-xxxxx",
        "NodeTemplate.InstanceType": "E1.4B",
        "NodeTemplate.SecurityGroupId": "8728f772-6186-xxx-xxxxx",
        "NodeTemplate.ChargeType": "HourlyInstantSettlement",
        "NodeTemplate.SystemDisk.DiskType": "Local_SSD",
        "NodeTemplate.SubnetStrategy": "balanced-distribution",
        "NodeTemplate.AdvancedSetting.DockerPath": "/data/docker",
        "NodeTemplate.AdvancedSetting.Schedulable": true,
        "NodeTemplate.AdvancedSetting.UserScript": "ZWNobyAicG9zIiA+cG9z",
        "NodeTemplate.AdvancedSetting.PreUserScript": "ZWNobyAicHJlIiA+cHJl",
        "NodeTemplate.AdvancedSetting.ContainerLogMaxSize": 100,
        "NodeTemplate.AdvancedSetting.ContainerLogMaxFiles": 10,
        "NodeTemplate.SystemDisk.DiskSize": 20,
        "NodeTemplate.ProjectId": 0,
        "NodeTemplate.Password": "Xuanxxx",
        "NodeTemplate.DataDisk.1.Type": "EHDD",
        "NodeTemplate.DataDisk.1.Size": 30,
	}
		resp, err = svc.CreateNodePool(&createNodePool)
		if err != nil {
			fmt.Println("error:", err.Error())
		}
		if resp != nil {
			str, _ := json.Marshal(&resp)
			fmt.Printf("%+v\n", string(str))
		}
	//	****************************************查询节点池(DescribeNodePool())***************************************************
	describeNodePool := map[string]interface{}{
		 "ClusterId": "d9a0adf0-a8f3-xxx-xxxxx",
         "NodePoolId.1": "64xxxxxxxx",
	}
	resp, err = svc.DescribeNodePool(&describeNodePool)
	if err != nil {
		fmt.Println("error:", err.Error())
	}
	if resp != nil {
		str, _ := json.Marshal(&resp)
		fmt.Printf("%+v\n", string(str))
	}*/

	//	****************************************删除节点池(DeleteNodePool())***************************************************
	deleteNodePool := map[string]interface{}{
		"ClusterId": "d9a0adf0-a8f3-xxx-xxxxx",
		"NodePoolId.1": "64xxxxxxxx",
	}
	resp, err = svc.DeleteNodePool(&deleteNodePool)
	if err != nil {
		fmt.Println("error:", err.Error())
	}
	if resp != nil {
		str, _ := json.Marshal(&resp)
		fmt.Printf("%+v\n", string(str))
	}
/*	//	****************************************修改节点池(ModifyNodePool())***************************************************
	modifyNodePool := map[string]interface{}{
	  	"NodePoolId": "64xxxxxxxxx",
        "ClusterId": "d9a0adf0-a8f3-xxx-xxxxx",
        "NodePoolName": "xuan-update",
        "EnableAutoScale": false,
        "MinSize": 1,
        "DesiredCapacity": 0,
        "MaxSize": 3,
        "Label.0.Key": "labelkupdate",
        "Label.0.Value": "labelvupdate",
        "Taint.0.Key": "taintkupdate",
        "Taint.0.Value": "taintvupdate",
        "Taint.0.Effect": "NoSchedule",
        "UpdateExistingNodes": true,
	}
	resp, err = svc.ModifyNodePool(&modifyNodePool)
	if err != nil {
		fmt.Println("error:", err.Error())
	}
	if resp != nil {
		str, _ := json.Marshal(&resp)
		fmt.Printf("%+v\n", string(str))
	}

	//	****************************************修改节点池模板(ModifyNodeTemplate())***************************************************
	modifyNodeTemplate := map[string]interface{}{
        "ClusterId": "d9a0adf0-a8f3-xxx-xxxxx",
        "NodePoolId": "64xxxxxxxxx",
        "NodeTemplate.ImageId": "13107fc5-0dd8-xxx-xxxxx",
        //"NodeTemplate.KeyId.1":"c409d431-xxx-xxx-9834-xxxxx",
        "NodeTemplate.DataDiskGb": 0,
        "NodeTemplate.SubnetId.1": "2c83f2d6-8160-xxx-xxxxx",
        "NodeTemplate.InstanceType": "E1.2B",
        "NodeTemplate.SecurityGroupId": "8728f772-6186-xxx-xxxxx",
        "NodeTemplate.ChargeType": "HourlyInstantSettlement",
        "NodeTemplate.SystemDisk.DiskType": "Local_SSD",
        "NodeTemplate.SubnetStrategy": "balanced-distribution",
        "NodeTemplate.AdvancedSetting.DockerPath": "/data/docker/xuan",
        "NodeTemplate.AdvancedSetting.Schedulable": true,
        "NodeTemplate.AdvancedSetting.UserScript": "ZWNobyAicG9zIiA+cG9z",
        "NodeTemplate.AdvancedSetting.PreUserScript": "ZWNobyAicHJlIiA+cHJl",
        "NodeTemplate.AdvancedSetting.ContainerLogMaxSize": 110,
        "NodeTemplate.AdvancedSetting.ContainerLogMaxFiles": 11,
        "NodeTemplate.SystemDisk.DiskSize": 20,
        "NodeTemplate.ProjectId": 0,
        "NodeTemplate.Password": "Xuanxxx",
        "NodeTemplate.DataDisk.1.Type": "EHDD",
        "NodeTemplate.DataDisk.1.Size": 30,
	}
	resp, err = svc.ModifyNodeTemplate(&modifyNodeTemplate)
	if err != nil {
		fmt.Println("error:", err.Error())
	}
	if resp != nil {
		str, _ := json.Marshal(&resp)
		fmt.Printf("%+v\n", string(str))
	}
	//	****************************************修改节点池扩容策略(ModifyNodePoolScaleUpPolicy())***************************************************
	modifyNodePoolScaleUpPolicy := map[string]interface{}{
		 "ClusterId": "d9a0adf0-a8f3-xxx-xxxxx",
         "ScaleUpPolicy": "most-pods",
	}
	resp, err = svc.ModifyNodePoolScaleUpPolicy(&modifyNodePoolScaleUpPolicy)
	if err != nil {
		fmt.Println("error:", err.Error())
	}
	if resp != nil {
		str, _ := json.Marshal(&resp)
		fmt.Printf("%+v\n", string(str))
	}
	//	****************************************修改节点池缩容策略(ModifyNodePoolScaleDownPolicy())***************************************************
	modifyNodePoolScaleDownPolicy := map[string]interface{}{
		"ClusterId": "d9a0adf0-a8f3-xxx-xxxxx",
        "MaxEmptyBulkDelete": 1,
        "ScaleDownDelayAfterAdd": 1,
        "ScaleDownEnabled": true,
        "ScaleDownUnneededTime": 1,
        "ScaleDownUtilizationThreshold": 60,
	}
	/*resp, err = svc.ModifyNodePoolScaleDownPolicy(&modifyNodePoolScaleDownPolicy)
	if err != nil {
		fmt.Println("error:", err.Error())
	}
	if resp != nil {
		str, _ := json.Marshal(&resp)
		fmt.Printf("%+v\n", string(str))
	}*/
	//	****************************************添加节点至节点池(AddClusterInstanceToNodePool())***************************************************
	/*addClusterInstanceToNodePool := map[string]interface{}{
		  "ClusterId": "d9a0adf0-a8f3-xxx-xxxxx",
	      "NodePoolId": "64xxxxxxxxx",
	      "InstanceIds.1": "7654649f-bbc4-xxx-xxxxx",
	}
	resp, err = svc.AddClusterInstanceToNodePool(&addClusterInstanceToNodePool)
	if err != nil {
		fmt.Println("error:", err.Error())
	}
	if resp != nil {
		str, _ := json.Marshal(&resp)
		fmt.Printf("%+v\n", string(str))
	}*/
	//	****************************************设置缩容保护(ProtectedFromScaleDown())***************************************************
	/*protectedFromScaleDown := map[string]interface{}{
		 "ClusterId": "d9a0adf0-a8f3-xxx-xxxxx",
	     "NodePoolId": "64xxxxxxxxx",
	     "InstanceIds.1": "7654649f-bbc4-xxx-xxxxx",
	     "ProtectedFromScaleDown": true,
	}
	resp, err = svc.ProtectedFromScaleDown(&protectedFromScaleDown)
	if err != nil {
		fmt.Println("error:", err.Error())
	}
	if resp != nil {
		str, _ := json.Marshal(&resp)
		fmt.Printf("%+v\n", string(str))
	}*/
	//	****************************************节点池移除节点(DeleteClusterInstancesFromNodePool())***************************************************
/*
	deleteClusterInstancesFromNodePool := map[string]interface{}{
		"ClusterId": "d9a0adf0-a8f3-xxx-xxxxx",
        "NodePoolId": "64xxxxxxxxx",
        "InstanceIds.1": "40b1fee5-5d0d-xxx-xxxxx",
        "InstanceDeleteMode": "Terminate",
	}
	resp, err = svc.DeleteClusterInstancesFromNodePool(&deleteClusterInstancesFromNodePool)
	if err != nil {
		fmt.Println("error:", err.Error())
	}
	if resp != nil {
		str, _ := json.Marshal(&resp)
		fmt.Printf("%+v\n", string(str))
	}*/

}
