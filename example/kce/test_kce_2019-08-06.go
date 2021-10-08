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

	//	****************************************创建集群(CreateCluster)***************************************************
/*	createCluster := map[string]interface{}{
		"ClusterName": "xxxxx",
		"ClusterManageMode": "DedicatedCluster",
		"VpcId": "15552848-026b-4ad6-a3e3-xxxxx",
		"PodCidr": "10.32.0.0/14",
		"ServiceCidr": "10.254.0.0/16",
		"NetworkType": "Flannel",
		"K8sVersion": "v1.13.4",
		"ReserveSubnetId": "9729f4c0-93ee-4e2a-9b2a-xxxxx",
		"PublicApiServer": "{\"LineId\":\"5fc2595f-1bfd-481b-bf64-2d08f116d800\",\"BandWidth\": \"10\",\"ChargeType\": \"PostPaidByDay\"}",
		"InstanceSet.0.NodeRole": "Master_Etcd",
		"InstanceSet.0.NodePara.0": "{\"MaxCount\":3,\"MinCount\":3,\"ImageId\":\"a0699172-76c6-4885-a4e9-0799a9cb811d\",\"SubnetId\":\"4a077fa8-a239-47bf-a18f-xxxxx\",\"InstancePassword\":\"Root23123\",\"SecurityGroupId\":\"0dcd8356-9e7e-43ae-897b-xxxxx\",\"DataDiskGb\":0,\"ChargeType\":\"Daily\",\"InstanceType\":\"I3.2A\",\"PurchaseTime\":0,\"InstanceName\":\"kce-py\",\"InstanceNameSuffix\":1}",
		"InstanceSet.1.NodeRole": "Worker",
		"InstanceSet.1.NodePara.0": "{\"MaxCount\":1,\"MinCount\":1,\"ImageId\":\"a0699172-76c6-4885-a4e9-0799a9cb811d\",\"SubnetId\":\"4a077fa8-a239-47bf-a18f-xxxxx\",\"InstancePassword\":\"Root23123\",\"SecurityGroupId\":\"0dcd8356-9e7e-43ae-897b-xxxxx\",\"DataDiskGb\":0,\"ChargeType\":\"Daily\",\"InstanceType\":\"I3.2A\",\"PurchaseTime\":0,\"InstanceName\":\"kce-py\",\"InstanceNameSuffix\":1}",
	}
	resp, err = svc.CreateCluster(&createCluster)
	if err != nil {
		fmt.Println("error:", err.Error())
	}
	if resp != nil {
		str, _ := json.Marshal(&resp)
		fmt.Printf("%+v\n", string(str))
	}*/
	//	****************************************查询集群列表(DescribeCluster())***************************************************
	describeCluster := map[string]interface{}{
		"ClusterId": "d9a0adf0-a8f3-xxx-xxxxx",
	}
	resp, err = svc.DescribeCluster(&describeCluster)
	if err != nil {
		fmt.Println("error:", err.Error())
	}
	if resp != nil {
		str, _ := json.Marshal(&resp)
		fmt.Printf("%+v\n", string(str))
	}
	//	****************************************新增节点(AddClusterInstances())***************************************************
	addClusterInstances := map[string]interface{}{
		"ClusterId": "a77b437f-07c9-4ae7-ac8d-xxxxx",
		"InstanceSet.0.NodeRole": "Worker",
		"InstanceSet.0.NodePara.0": "{\"MaxCount\":1,\"MinCount\":1,\"ImageId\":\"a0699172-76c6-4885-a4e9-0799a9cb811d\",\"SubnetId\":\"4a077fa8-a239-47bf-a18f-xxxxx\",\"InstancePassword\":\"Root23123\",\"SecurityGroupId\":\"0dcd8356-9e7e-43ae-897b-xxxxx\",\"DataDiskGb\":0,\"ChargeType\":\"Daily\",\"InstanceType\":\"I3.2A\",\"PurchaseTime\":0,\"InstanceName\":\"kce-py\",\"InstanceNameSuffix\":1}",
	}
	resp, err = svc.AddClusterInstances(&addClusterInstances)
	if err != nil {
		fmt.Println("error:", err.Error())
	}
	if resp != nil {
		str, _ := json.Marshal(&resp)
		fmt.Printf("%+v\n", string(str))
	}
	//	****************************************删除集群中的节点(DeleteClusterInstances())***************************************************
	deleteClusterInstances := map[string]interface{}{
		"ClusterId": "a77b437f-07c9-4ae7-ac8d-xxxxx",
		"InstanceId.1": "0253d503-485e-4adc-8859-xxxxx",
	}
	resp, err = svc.DeleteClusterInstances(&deleteClusterInstances)
	if err != nil {
		fmt.Println("error:", err.Error())
	}
	if resp != nil {
		str, _ := json.Marshal(&resp)
		fmt.Printf("%+v\n", string(str))
	}
	//	*********************强制移除节点,该接口必须在DeleteClusterInstances后执行(ForceRemoveClusterInstance())*******************
	forceRemoveClusterInstance := map[string]interface{}{
		"ClusterId": "a77b437f-07c9-4ae7-ac8d-xxxxx",
		"InstanceId.1": "0253d503-485e-4adc-8859-xxxxx",
		"InstanceId.2": "0253d503-485e-4adc-8860-xxxxx",
	}
	resp, err = svc.ForceRemoveClusterInstance(&forceRemoveClusterInstance)
	if err != nil {
		fmt.Println("error:", err.Error())
	}
	if resp != nil {
		str, _ := json.Marshal(&resp)
		fmt.Printf("%+v\n", string(str))
	}
	//	****************************************修改集群基本信息(ModifyClusterInfo())***************************************************
	modifyClusterInfo := map[string]interface{}{
		"ClusterId": "d9a0adf0-a8f3-xxx-xxxxx",
		"ClusterName": "update",
		"ClusterDesc": "xxxxxxxx",
	}
	resp, err = svc.ModifyClusterInfo(&modifyClusterInfo)
	if err != nil {
		fmt.Println("error:", err.Error())
	}
	if resp != nil {
		str, _ := json.Marshal(&resp)
		fmt.Printf("%+v\n", string(str))
	}

	//	****************************************下载集群配置文件(DownloadClusterConfig())***************************************************
	downloadClusterConfig := map[string]interface{}{
		"ClusterId": "d9a0adf0-a8f3-xxx-xxxxx",
	}
	resp, err = svc.DownloadClusterConfig(&downloadClusterConfig)
	if err != nil {
		fmt.Println("error:", err.Error())
	}
	if resp != nil {
		str, _ := json.Marshal(&resp)
		fmt.Printf("%+v\n", string(str))
	}
	//	****************************************删除集群(DeleteCluster())***************************************************
	deleteCluster := map[string]interface{}{
		"ClusterId": "d9a0adf0-a8f3-xxx-xxxxx",
	}
	resp, err = svc.DeleteCluster(&deleteCluster)
	if err != nil {
		fmt.Println("error:", err.Error())
	}
	if resp != nil {
		str, _ := json.Marshal(&resp)
		fmt.Printf("%+v\n", string(str))
	}
	//	****************************************获取容器服务支持的节点操作系统(DescribeInstanceImage())***************************************************
	describeInstanceImage := map[string]interface{}{
		"ClusterId": "d9a0adf0-a8f3-xxx-xxxxx",
	}
	resp, err = svc.DescribeInstanceImage(&describeInstanceImage)
	if err != nil {
		fmt.Println("error:", err.Error())
	}
	if resp != nil {
		str, _ := json.Marshal(&resp)
		fmt.Printf("%+v\n", string(str))
	}
	//	****************************************获取支持移入物理机集群的epc列表(DescribeEpcForCluster())***************************************************
	describeEpcForCluster := map[string]interface{}{
		"ClusterId": "d9a0adf0-a8f3-xxx-xxxxx",
	}
	resp, err = svc.DescribeEpcForCluster(&describeEpcForCluster)
	if err != nil {
		fmt.Println("error:", err.Error())
	}
	if resp != nil {
		str, _ := json.Marshal(&resp)
		fmt.Printf("%+v\n", string(str))
	}
	//	****************************************移入物理机服务器至物理机集群(AddClusterEpcInstances())***************************************************
	addClusterEpcInstances := map[string]interface{}{
		"ClusterId": "d9a0adf0-a8f3-xxx-xxxxx",
		"InstanceId.1": "64xxxxxxxx",
	}
	resp, err = svc.AddClusterEpcInstances(&addClusterEpcInstances)
	if err != nil {
		fmt.Println("error:", err.Error())
	}
	if resp != nil {
		str, _ := json.Marshal(&resp)
		fmt.Printf("%+v\n", string(str))
	}
	//	****************************************查询已经存在的云服务器(DescribeExistedInstances())***************************************************
	describeExistedInstances := map[string]interface{}{
		"ClusterId": "d9a0adf0-a8f3-xxx-xxxxx",
		"InstanceId.1": "64xxxxxxxx",
	}
	resp, err = svc.DescribeExistedInstances(&describeExistedInstances)
	if err != nil {
		fmt.Println("error:", err.Error())
	}
	if resp != nil {
		str, _ := json.Marshal(&resp)
		fmt.Printf("%+v\n", string(str))
	}
	//	****************************************添加已有的服务器(AddExistedInstances())***************************************************
	addExistedInstances := map[string]interface{}{
		"ClusterId": "84d89f76-xxxx-47a2-b37e-xxxxx",
		"ExistedInstanceKecSet.1.NodeRole": "worker",
		"ExistedInstanceKecSet.1.KecPara.1": "{\"InstanceId\":\"8d1cae6a-23c3-47f6-8fe6-xxxxx\",\"ImageId\":\"81cc01c3-4d64-40fa-89af-xxxxx\",\"InstancePassword\":\"xxxxx\"}",
	}
	resp, err = svc.AddExistedInstances(&addExistedInstances)
	if err != nil {
		fmt.Println("error:", err.Error())
	}
	if resp != nil {
		str, _ := json.Marshal(&resp)
		fmt.Printf("%+v\n", string(str))
	}
	//	****************************************获取裸金属服务器支持的节点操作系统(DescribeEpcImage())***************************************************
	describeEpcImage := map[string]interface{}{
		//"ImageId": "64xxxxxxxx",
	}
	resp, err = svc.DescribeEpcImage(&describeEpcImage)
	if err != nil {
		fmt.Println("error:", err.Error())
	}
	if resp != nil {
		str, _ := json.Marshal(&resp)
		fmt.Printf("%+v\n", string(str))
	}
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
