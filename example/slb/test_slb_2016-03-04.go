package main

import (
	"encoding/json"
	"fmt"
	"github.com/KscSDK/ksc-sdk-go/ksc"
	"github.com/KscSDK/ksc-sdk-go/ksc/utils"
	"github.com/KscSDK/ksc-sdk-go/service/slb"
)

func main() {
	ak := "你的ak"
	sk := "你的sk"
	region := "cn-shanghai-2"
	//debug模式的话 打开这个开关
	svc := slb.SdkNew(ksc.NewClient(ak, sk /*true*/), &ksc.Config{Region: &region}, &utils.UrlInfo{
		UseSSL: true,
	})
	var resp *map[string]interface{}
	var err error
	//	****************************************创建负载均衡(CreateLoadBalancer())***************************************************
	/*
		createLoadBalancer := make(map[string]interface{})
			createLoadBalancer["VpcId"] = "813a1598-79ea-4922-a262-66e195fd873b"
			createLoadBalancer["LoadBalancerName"] = "xuantt"
			createLoadBalancer["Type"] = "public"

			resp, err = svc.CreateLoadBalancer(&createLoadBalancer)
			if err != nil {
				fmt.Println("error:", err.Error())
			}
			if resp != nil {
				str, _ := json.Marshal(&resp)
				fmt.Printf("%+v\n", string(str))
			}

	*/

	//	****************************************删除负载均衡(DeleteLoadBalancer())***************************************************
	/*

		deleteLoadBalancer := make(map[string]interface{})
		deleteLoadBalancer["LoadBalancerId"] = "b2e1c23e-2038-4b46-8885-2d5425d7823d"
		resp, err = svc.DeleteLoadBalancer(&deleteLoadBalancer)
		if err != nil {
			fmt.Println("error:", err.Error())
		}
		if resp != nil {
			str, _ := json.Marshal(&resp)
			fmt.Printf("%+v\n", string(str))
		}

	*/
	//	****************************************修改负载均衡(ModifyLoadBalancer())***************************************************
	/*

		modifyLoadBalancer := make(map[string]interface{})
		modifyLoadBalancer["LoadBalancerId"] = "312f5962-1519-4696-896b-251caeee8718"
		modifyLoadBalancer["LoadBalancerName"] = "xuantest2"
		modifyLoadBalancer["LoadBalancerState"] = "stop"

		resp, err = svc.ModifyLoadBalancer(&modifyLoadBalancer)
		if err != nil {
			fmt.Println("error:", err.Error())
		}
		if resp != nil {
			str, _ := json.Marshal(&resp)
			fmt.Printf("%+v\n", string(str))
		}

	*/

	//	****************************************查询负载均衡信息(DescribeLoadBalancers())***************************************************

	describeLoadBalancers := make(map[string]interface{})
	describeLoadBalancers["LoadBalancerId.1"] = "b2e1c23e-2038-4b46-8885-2d5425d7823d"
	describeLoadBalancers["Filter.1.Name"] = "vpc-id"
	describeLoadBalancers["Filter.1.Value.1"] = "26cdbe78-1943-44ff-9125-b7c1e81f02c4"

	resp, err = svc.DescribeLoadBalancers(&describeLoadBalancers)
	if err != nil {
		fmt.Println("error:", err.Error())
	}
	if resp != nil {
		str, _ := json.Marshal(&resp)
		fmt.Printf("%+v\n", string(str))
	}

	//	****************************************创建监听器(CreateListeners())***************************************************
	/*
		createListeners := make(map[string]interface{})
		createListeners["LoadBalancerId"] = "312f5962-1519-4696-896b-251caeee8718"
		createListeners["ListenerProtocol"] = "TCP"
		createListeners["ListenerState"] = "start"
		createListeners["ListenerName"] = "xuantest2"
		createListeners["ListenerPort"] = "8010"
		createListeners["Method"] = "RoundRobin"
		createListeners["SessionState"] = "start"
		createListeners["SessionPersistencePeriod"] = "3600"

		resp, err = svc.CreateListeners(&createListeners)
		if err != nil {
			fmt.Println("error:", err.Error())
		}
		if resp != nil {
			str, _ := json.Marshal(&resp)
			fmt.Printf("%+v\n", string(str))
		}
	*/
	//	****************************************更改监听器(ModifyListeners())***************************************************
	/*
		modifyListeners := make(map[string]interface{})
		modifyListeners["ListenerId"] = "f3518801-7ca0-44a5-a037-45e885829a3e"
		modifyListeners["SessionPersistencePeriod"] = "60"
		modifyListeners["CookieType"] = "ImplantCookie"
		modifyListeners["CookieName"] = "xuan"
		modifyListeners["SessionState"] = "start"

		resp, err = svc.ModifyListeners(&modifyListeners)
		if err != nil {
			fmt.Println("error:", err.Error())
		}
		if resp != nil {
			str, _ := json.Marshal(&resp)
			fmt.Printf("%+v\n", string(str))
		}

	*/
	//	****************************************删除监听器(DeleteListeners())***************************************************
	/*
		deleteListeners := make(map[string]interface{})
		deleteListeners["ListenerId"] = "b7032b53-2e65-4667-b582-4c4c0776dbaa"
		resp, err = svc.DeleteListeners(&deleteListeners)
		if err != nil {
			fmt.Println("error:", err.Error())
		}
		if resp != nil {
			str, _ := json.Marshal(&resp)
			fmt.Printf("%+v\n", string(str))
		}

	*/
	//	****************************************查询监听器信息(DescribeListeners())***************************************************
	/*
		describeListeners := make(map[string]interface{})
		describeListeners["Filter.1.Name"] = "load-balancer-id"
		describeListeners["Filter.1.Value.1"] = "a2401d6b-d706-49b3-a0b4-e2fdb0c38899"
		resp, err = svc.DescribeListeners(&describeListeners)
		if err != nil {
			fmt.Println("error:", err.Error())
		}
		if resp != nil {
			str, _ := json.Marshal(&resp)
			fmt.Printf("%+v\n", string(str))
		}
	*/
	//	****************************************创建健康检查(ConfigureHealthCheck())***************************************************
	/*
		configureHealthCheck := make(map[string]interface{})
		configureHealthCheck["ListenerId"] = "d61cc536-081a-47a6-9225-b3419b1d9736"
		configureHealthCheck["HealthyThreshold"] = "9"
		configureHealthCheck["Interval"] = "10"
		configureHealthCheck["Timeout"] = "10"
		configureHealthCheck["UnhealthyThreshold"] = "10"
		configureHealthCheck["HealthCheckState"] = "start"
		resp, err = svc.ConfigureHealthCheck(&configureHealthCheck)
		if err != nil {
			fmt.Println("error:", err.Error())
		}
		if resp != nil {
			str, _ := json.Marshal(&resp)
			fmt.Printf("%+v\n", string(str))
		}

	*/
	//	****************************************修改健康检查(ModifyHealthCheck())***************************************************
	/*
		modifyHealthCheck := make(map[string]interface{})
		modifyHealthCheck["HealthCheckId"] = "aaadf285-dc29-4165-9840-7df57992c10c"
		modifyHealthCheck["HealthyThreshold"] = "2"
		modifyHealthCheck["Interval"] = "20"
		modifyHealthCheck["Timeout"] = "30"
		modifyHealthCheck["UnhealthyThreshold"] = "3"
		modifyHealthCheck["HealthCheckState"] = "start"
		resp, err = svc.ModifyHealthCheck(&modifyHealthCheck)
		if err != nil {
			fmt.Println("error:", err.Error())
		}
		if resp != nil {
			str, _ := json.Marshal(&resp)
			fmt.Printf("%+v\n", string(str))
		}
	*/
	//	****************************************删除健康检查(DeleteHealthCheck())***************************************************
	/*
		   deleteHealthCheck := make(map[string]interface{})
		   deleteHealthCheck["HealthCheckId"] = "aaadf285-dc29-4165-9840-7df57992c10c"
		   resp, err = svc.DeleteHealthCheck(&deleteHealthCheck)
		   if err != nil {
			   fmt.Println("error:", err.Error())
		   }
		   if resp != nil {
			   str, _ := json.Marshal(&resp)
			   fmt.Printf("%+v\n", string(str))
		   }

	*/
	//	****************************************查询健康检查信息(DescribeHealthChecks())***************************************************
	/*
		   describeHealthChecks := make(map[string]interface{})
		   describeHealthChecks["Filter.1.Name"] = "listener-id"
		   describeHealthChecks["Filter.1.Value.1"] = "d61cc536-081a-47a6-9225-b3419b1d9736"
		   describeHealthChecks["HealthCheckId.1"] = "aaadf285-dc29-4165-9840-7df57992c10c"
		   resp, err = svc.DescribeHealthChecks(&describeHealthChecks)
		   if err != nil {
			   fmt.Println("error:", err.Error())
		   }
		   if resp != nil {
			   str, _ := json.Marshal(&resp)
			   fmt.Printf("%+v\n", string(str))
		   }
	*/
	//	****************************************添加后端服务(RegisterInstancesWithListener())***************************************************
	/*
		   registerInstancesWithListener := make(map[string]interface{})
		   registerInstancesWithListener["ListenerId"] = "b7032b53-2e65-4667-b582-4c4c0776dbaa"
		   registerInstancesWithListener["RealServerType"] = "host"
		   registerInstancesWithListener["RealServerIp"] = "70.0.0.21"
		   registerInstancesWithListener["RealServerPort"] = "80"
		   registerInstancesWithListener["Weight"] = "10"
		   registerInstancesWithListener["InstanceId"] = "5ba7eb8a-2e91-4d16-84a9-6112f3e7b1bb"
		   resp, err = svc.RegisterInstancesWithListener(&registerInstancesWithListener)
		   if err != nil {
			   fmt.Println("error:", err.Error())
		   }
		   if resp != nil {
			   str, _ := json.Marshal(&resp)
			   fmt.Printf("%+v\n", string(str))
		   }
	*/
	//	****************************************修改后端服务(ModifyInstancesWithListener())***************************************************
	/*
		modifyInstancesWithListener := make(map[string]interface{})
		modifyInstancesWithListener["RegisterId"] = "992a7593-bd72-42bb-91b4-36b9f9bf5f9d"
		modifyInstancesWithListener["RealServerPort"] = "8080"
		modifyInstancesWithListener["Weight"] = "10"
		resp, err = svc.ModifyInstancesWithListener(&modifyInstancesWithListener)
		if err != nil {
			fmt.Println("error:", err.Error())
		}
		if resp != nil {
			str, _ := json.Marshal(&resp)
			fmt.Printf("%+v\n", string(str))
		}

	*/
	//	****************************************解绑后端服务(DeregisterInstancesFromListener())***************************************************
	/*
		deregisterInstancesFromListener := make(map[string]interface{})
		deregisterInstancesFromListener["RegisterId"] = "992a7593-bd72-42bb-91b4-36b9f9bf5f9d"
		resp, err = svc.DeregisterInstancesFromListener(&deregisterInstancesFromListener)
		if err != nil {
			fmt.Println("error:", err.Error())
		}
		if resp != nil {
			str, _ := json.Marshal(&resp)
			fmt.Printf("%+v\n", string(str))
		}

	*/
	//	****************************************查询后端服务信息(DescribeInstancesWithListener())***************************************************
	/*
		describeInstancesWithListener := make(map[string]interface{})
		describeInstancesWithListener["RegisterId.1"] = "62b2823b-7137-403e-bba4-31c2a90f4fe2"
		describeInstancesWithListener["Filter.1.Name"] = "listener-id"
		describeInstancesWithListener["Filter.1.Value.1"] = "b7032b53-2e65-4667-b582-4c4c0776dbaa"
		resp, err = svc.DescribeInstancesWithListener(&describeInstancesWithListener)
		if err != nil {
			fmt.Println("error:", err.Error())
		}
		if resp != nil {
			str, _ := json.Marshal(&resp)
			fmt.Printf("%+v\n", string(str))
		}
	*/
	//	****************************************创建负载均衡ACL(CreateLoadBalancerAcl())***************************************************
	/*
		createLoadBalancerAcl := make(map[string]interface{})
		createLoadBalancerAcl["LoadBalancerAclName"] = "xuan"
		resp, err = svc.CreateLoadBalancerAcl(&createLoadBalancerAcl)
		if err != nil {
			fmt.Println("error:", err.Error())
		}
		if resp != nil {
			str, _ := json.Marshal(&resp)
			fmt.Printf("%+v\n", string(str))
		}

	*/
	//	****************************************删除负载均衡ACL(DeleteLoadBalancerAcl())***************************************************
	/*
		deleteLoadBalancerAcl := make(map[string]interface{})
		deleteLoadBalancerAcl["LoadBalancerAclId"] = "3aadee9e-e586-4943-862d-472e164a5e00"
		resp, err = svc.DeleteLoadBalancerAcl(&deleteLoadBalancerAcl)
		if err != nil {
			fmt.Println("error:", err.Error())
		}
		if resp != nil {
			str, _ := json.Marshal(&resp)
			fmt.Printf("%+v\n", string(str))
		}
	*/
	//	****************************************修改负载均衡ACL(ModifyLoadBalancerAcl())***************************************************
	/*
		modifyLoadBalancerAcl := make(map[string]interface{})
		modifyLoadBalancerAcl["LoadBalancerAclId"] = "3aadee9e-e586-4943-862d-472e164a5e00"
		modifyLoadBalancerAcl["LoadBalancerAclName"] = "xuanxuantest"
		resp, err = svc.ModifyLoadBalancerAcl(&modifyLoadBalancerAcl)
		if err != nil {
			fmt.Println("error:", err.Error())
		}
		if resp != nil {
			str, _ := json.Marshal(&resp)
			fmt.Printf("%+v\n", string(str))
		}

	*/
	//	****************************************查询负载均衡ACL信息(DescribeLoadBalancerAcls())***************************************************
	/*
		describeLoadBalancerAcls := make(map[string]interface{})
		resp, err = svc.DescribeLoadBalancerAcls(&describeLoadBalancerAcls)
		if err != nil {
			fmt.Println("error:", err.Error())
		}
		if resp != nil {
			str, _ := json.Marshal(&resp)
			fmt.Printf("%+v\n", string(str))
		}
	*/
	//	****************************************新建负载均衡ACL规则(CreateLoadBalancerAclEntry())***************************************************
	/*
		createLoadBalancerAclEntry := make(map[string]interface{})
		createLoadBalancerAclEntry["LoadBalancerAclId"] = "3aadee9e-e586-4943-862d-472e164a5e00"
		createLoadBalancerAclEntry["CidrBlock"] = "0.0.0.0/0"
		createLoadBalancerAclEntry["RuleNumber"] = "33"
		createLoadBalancerAclEntry["RuleAction"] = "allow"
		createLoadBalancerAclEntry["Protocol"] = "ip"
		resp, err = svc.CreateLoadBalancerAclEntry(&createLoadBalancerAclEntry)
		if err != nil {
			fmt.Println("error:", err.Error())
		}
		if resp != nil {
			str, _ := json.Marshal(&resp)
			fmt.Printf("%+v\n", string(str))
		}
	*/
	//	****************************************删除负载均衡ACL规则(DeleteLoadBalancerAclEntry())***************************************************
	/*
		deleteLoadBalancerAclEntry := make(map[string]interface{})
		deleteLoadBalancerAclEntry["LoadBalancerAclId"] = "3aadee9e-e586-4943-862d-472e164a5e00"
		deleteLoadBalancerAclEntry["LoadBalancerAclEntryId"] = "81f0ccd0-02e3-4eec-8ce0-f093d93212f3"
		resp, err = svc.DeleteLoadBalancerAclEntry(&deleteLoadBalancerAclEntry)
		if err != nil {
			fmt.Println("error:", err.Error())
		}
		if resp != nil {
			str, _ := json.Marshal(&resp)
			fmt.Printf("%+v\n", string(str))
		}

	*/
	//	****************************************监听器绑定ACL(AssociateLoadBalancerAcl())***************************************************
	/*

		associateLoadBalancerAcl := make(map[string]interface{})
		associateLoadBalancerAcl["ListenerId"] = "b7032b53-2e65-4667-b582-4c4c0776dbaa"
		associateLoadBalancerAcl["LoadBalancerAclId"] = "3aadee9e-e586-4943-862d-472e164a5e00"
		resp, err = svc.AssociateLoadBalancerAcl(&associateLoadBalancerAcl)
		if err != nil {
			fmt.Println("error:", err.Error())
		}
		if resp != nil {
			str, _ := json.Marshal(&resp)
			fmt.Printf("%+v\n", string(str))
		}

	*/
	//	****************************************监听器解绑ACL(DisassociateLoadBalancerAcl())***************************************************
	/*

		disassociateLoadBalancerAcl := make(map[string]interface{})
		disassociateLoadBalancerAcl["ListenerId"] = "b7032b53-2e65-4667-b582-4c4c0776dbaa"
		resp, err = svc.DisassociateLoadBalancerAcl(&disassociateLoadBalancerAcl)
		if err != nil {
			fmt.Println("error:", err.Error())
		}
		if resp != nil {
			str, _ := json.Marshal(&resp)
			fmt.Printf("%+v\n", string(str))
		}

	*/
}
