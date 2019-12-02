package main

import (
	"encoding/json"
	"fmt"
	"github.com/KscSDK/ksc-sdk-go/ksc"
	"github.com/KscSDK/ksc-sdk-go/ksc/utils"
	"github.com/KscSDK/ksc-sdk-go/service/eip"
)

func main() {
	ak := "你的ak"
	sk := "你的sk"
	region := "cn-shanghai-2"
	//debug模式的话 打开这个开关
	svc := eip.SdkNew(ksc.NewClient(ak, sk /*,true*/), &ksc.Config{Region: &region}, &utils.UrlInfo{
		UseSSL:      true,
		UseInternal: true,
	})
	var resp *map[string]interface{}
	var err error
	//	****************************************获取用户可选链路信息(GetLines())***************************************************
	/*
		resp, err = svc.GetLines(nil)
		if err != nil {
			fmt.Println("error:", err.Error())
		}
		if resp != nil {
			str, _ := json.Marshal(&resp)
			fmt.Printf("%+v\n", string(str))
		}
	*/
	//	****************************************创建弹性IP(AllocateAddress())***************************************************

	/*
		   allocateAddress:=make(map[string]interface{})
		   allocateAddress["LineId"]="a2403858-2550-4612-850c-ea840fa343f9"
		   allocateAddress["BandWidth"]="10"
		   allocateAddress["ChargeType"]="PostPaidByDay"
		   allocateAddress["PurchaseTime"]="1"

		   resp, err = svc.AllocateAddress(&allocateAddress)
		   if err != nil {
			   fmt.Println("error:", err.Error())
		   }
		   if resp != nil {
			   str, _ := json.Marshal(&resp)
			   fmt.Printf("%+v\n", string(str))
		   }
		   // */
	//	****************************************删除弹性IP(ReleaseAddress())***************************************************

	/*
		releaseAddress:=make(map[string]interface{})
			releaseAddress["AllocationId"]="987afbec-b090-4529-8b7f-38c1c5cb1f21"

			resp, err = svc.ReleaseAddress(&releaseAddress)
			if err != nil {
				fmt.Println("error:", err.Error())
			}
			if resp != nil {
				str, _ := json.Marshal(&resp)
				fmt.Printf("%+v\n", string(str))
			}

	*/
	//	****************************************绑定弹性IP(AssociateAddress())***************************************************

	/*
		associateAddress:=make(map[string]interface{})
		associateAddress["AllocationId"]="a1df0fa0-ba1f-416a-a634-9465e5ab999f"
		associateAddress["InstanceType"]="Slb"
		associateAddress["InstanceId"]="c3dbfca9-d7f9-4902-9b3a-a6092ebf8e7e"
		//associateAddress["NetworkInterfaceId"]="25b3328d-eeef-4f0f-ba71-89d199d1ca32"（主机不可缺省）
		resp, err = svc.AssociateAddress(&associateAddress)
		if err != nil {
			fmt.Println("error:", err.Error())
		}
		if resp != nil {
			str, _ := json.Marshal(&resp)
			fmt.Printf("%+v\n", string(str))
		}



	*/
	//	****************************************解绑弹性IP(DisassociateAddress())***************************************************
	/*
		disassociateAddress:=make(map[string]interface{})
		disassociateAddress["AllocationId"]="a1df0fa0-ba1f-416a-a634-9465e5ab999f"

		resp, err = svc.DisassociateAddress(&disassociateAddress)
		if err != nil {
			fmt.Println("error:", err.Error())
		}
		if resp != nil {
			str, _ := json.Marshal(&resp)
			fmt.Printf("%+v\n", string(str))
		}

	*/
	//	****************************************查询弹性IP(DescribeAddresses())***************************************************

	describeAddresses := make(map[string]interface{})
	describeAddresses["Filter.1.Name"] = "instance-type"
	describeAddresses["Filter.1.Value.1"] = "Slb"
	describeAddresses["NextToken"] = "1"
	describeAddresses["MaxResults"] = "5"

	resp, err = svc.DescribeAddresses(nil)
	if err != nil {
		fmt.Println("error:", err.Error())
	}
	if resp != nil {
		str, _ := json.Marshal(&resp)
		fmt.Printf("%+v\n", string(str))
	}

	//	****************************************更新弹性IP配置(ModifyAddress())***************************************************
	/*
		modifyAddress:=make(map[string]interface{})
		modifyAddress["AllocationId"]="a1df0fa0-ba1f-416a-a634-9465e5ab999f"
		modifyAddress["BandWidth"]="2"
		resp, err = svc.ModifyAddress(&modifyAddress)
		if err != nil {
			fmt.Println("error:", err.Error())
		}
		if resp != nil {
			str, _ := json.Marshal(&resp)
			fmt.Printf("%+v\n", string(str))
		}
	*/
}
