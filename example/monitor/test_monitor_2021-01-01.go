package main

import (
	"encoding/json"
	"fmt"

	"github.com/KscSDK/ksc-sdk-go/ksc"
	"github.com/KscSDK/ksc-sdk-go/ksc/utils"
	"github.com/KscSDK/ksc-sdk-go/service/monitorv4"
)

//只适用于docker产品线
func main() {
	ak := "ak"
	sk := "sk"
	region := "cn-beijing-6"
	//debug模式的话 打开这个开关
	svc := monitorv4.SdkNew(ksc.NewClient(ak, sk /*true*/), &ksc.Config{Region: &region}, &utils.UrlInfo{
		UseSSL: true,
	})

	//	****************************************列出告警策略(ListAlarmPolicy())***************************************************

	list_alarm_policy := make(map[string]interface{})
	list_alarm_policy["PageIndex"] = "1"
	list_alarm_policy["PageSize"] = "10"
	resp, err := svc.ListAlarmPolicy(&list_alarm_policy)

	if err != nil {
		fmt.Println("error:", err.Error())
	}
	if resp != nil {
		str, _ := json.MarshalIndent(&resp, "", "    ")
		fmt.Printf("%+v\n", string(str))
	}

	fmt.Println("----------------------")

	//	****************************************查询告警策略详细信息(DescribeAlarmPolicy())************************************************

	desc_alarm_policy := make(map[string]interface{})
	desc_alarm_policy["PolicyId"] = "25232"
	resp, err = svc.DescribeAlarmPolicy(&desc_alarm_policy)

	if err != nil {
		fmt.Println("error:", err.Error())
	}
	if resp != nil {
		str, _ := json.MarshalIndent(&resp, "", "    ")
		fmt.Printf("%+v\n", string(str))
	}

	fmt.Println("----------------------")

	//	****************************************查询告警策略关联实例明细(DescribePolicyObject())*********************************************

	desc_policy_object := make(map[string]interface{})
	desc_policy_object["PolicyId"] = "25232"
	desc_policy_object["PageIndex"] = "1"
	desc_policy_object["PageSize"] = "10"
	resp, err = svc.DescribePolicyObject(&desc_policy_object)

	if err != nil {
		fmt.Println("error:", err.Error())
	}
	if resp != nil {
		str, _ := json.MarshalIndent(&resp, "", "    ")
		fmt.Printf("%+v\n", string(str))
	}

	fmt.Println("----------------------")

	//	****************************************获取告警策略关联的接收人(DescribeAlarmReceives())*********************************************

	desc_alarm_receives := make(map[string]interface{})
	desc_alarm_receives["PolicyId"] = "25232"
	resp, err = svc.DescribeAlarmReceives(&desc_alarm_receives)

	if err != nil {
		fmt.Println("error:", err.Error())
	}
	if resp != nil {
		str, _ := json.MarshalIndent(&resp, "", "    ")
		fmt.Printf("%+v\n", string(str))
	}

	fmt.Println("----------------------")

	//****************************************添加告警策略关联的接收人(AddAlarmReceives())*********************************************

	add_alarm_receives := make(map[string]interface{})
	add_alarm_receives["PolicyId"] = 25232
	add_alarm_receives["ContactFlag"] = 2
	add_alarm_receives["ContactWay"] = 3
	add_alarm_receives["ContactId"] = []int{1985, 3607}
	resp, err = svc.AddAlarmReceives(&add_alarm_receives)

	if err != nil {
		fmt.Println("error:", err.Error())
	}
	if resp != nil {
		str, _ := json.MarshalIndent(&resp, "", "    ")
		fmt.Printf("%+v\n", string(str))
	}

	fmt.Println("----------------------")

	//	****************************************删除告警策略关联的接收人(DeleteAlarmReceives())*********************************************

	delete_alarm_receives := make(map[string]interface{})
	delete_alarm_receives["PolicyId"] = 25232
	delete_alarm_receives["ContactFlag"] = 2
	delete_alarm_receives["ContactId"] = []int{1985, 3607}
	resp, err = svc.DeleteAlarmReceives(&delete_alarm_receives)

	if err != nil {
		fmt.Println("error:", err.Error())
	}
	if resp != nil {
		str, _ := json.MarshalIndent(&resp, "", "    ")
		fmt.Printf("%+v\n", string(str))
	}

	fmt.Println("----------------------")

	//	****************************************获取联系组(GetUserGroup())*********************************************

	resp, err = svc.GetUserGroup(nil)

	if err != nil {
		fmt.Println("error:", err.Error())
	}
	if resp != nil {
		str, _ := json.MarshalIndent(&resp, "", "    ")
		fmt.Printf("%+v\n", string(str))
	}

	fmt.Println("----------------------")

	//	****************************************获取联系人(GetAlertUser())*********************************************

	get_alert_user := make(map[string]interface{})
	get_alert_user["UserGrpId.1"] = "879"
	get_alert_user["UserGrpId.2"] = "1484"
	resp, err = svc.GetAlertUser(&get_alert_user)

	if err != nil {
		fmt.Println("error:", err.Error())
	}
	if resp != nil {
		str, _ := json.MarshalIndent(&resp, "", "    ")
		fmt.Printf("%+v\n", string(str))
	}

	fmt.Println("----------------------")

	//	****************************************启用或禁用联系人(UpdateAlertUserStatus())*********************************************

	update_alert_user := make(map[string]interface{})
	update_alert_user["UserId"] = []int{1985, 3607}
	update_alert_user["UserStatus"] = 2
	resp, err = svc.UpdateAlertUserStatus(&update_alert_user)

	if err != nil {
		fmt.Println("error:", err.Error())
	}
	if resp != nil {
		str, _ := json.MarshalIndent(&resp, "", "    ")
		fmt.Printf("%+v\n", string(str))
	}

	fmt.Println("----------------------")
}
