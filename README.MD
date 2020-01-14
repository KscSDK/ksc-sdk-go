### KSC SDK for Go
####  概述
1. 金山云Go语言SDK
2. Go语言版本1.5+ 推荐是用1.12+ 这样方便使用govendor
3. service目录下的文件请不要修改，都是由代码生成器自动生成 代码生成器暂时不开源
4. 如果使用vendor 请确保所有ksc-sdk-go 在GOPATH下面/src/github.com/KscSDK目录下 ($GOPATH/src/github.com/KscSDK)

##### vendor
使用vendor(Go 1.5+)做依赖的话可以直接使用本sdk工具

##### Go 1.5
需要去[AWS-SDK-GO](https://github.com/aws/aws-sdk-go) 下载后，拷贝aws-sdk-go到在GOPATH下面/src/github.com/aws目录下 ($GOPATH/src/github.com/aws)目录下 
或者使用 go get github.com/aws/aws-sdk-go

##### 代码示例
UseInternal:true 代表使用金山云内网域名

代码示例在example目录下

	ak := "用户AK"
	sk := "用户SK"
	region := "cn-beijing-6"
	svc := vpc.SdkNew(ksc.NewClient(ak, sk), &ksc.Config{Region: &region}, &utils.UrlInfo{
    		UseSSL: true,
    		UseInternal:true,
    	})

	//创建入参 参考 https://docs.ksyun.com/
	m_vpc := make(map[string]interface{})
	m_vpc["VpcId.1"] = "VPCID"

	//查询VPC信息
	resp, err := svc.DescribeVpcs(&m_vpc)
	if err != nil {
		fmt.Println("there was an error listing instances in", err.Error())
	}
	if resp != nil {
		l := (*resp)["VpcSet"].([]interface{})
		for i := 0; i < len(l); i++ {
			item := l[i].(map[string]interface{})
			fmt.Printf("%+v:%+v\n", item["VpcId"], item["VpcName"])
		}
	}
	//创建入参 参考 https://docs.ksyun.com/
	m_vnet := make(map[string]interface{})
	
	//查询子网信息
	resp1, err1 := svc.DescribeSubnets(&m_vnet)
	if err1 != nil {
		fmt.Println("there was an error listing instances in", err1.Error())
	}
	if resp1 != nil {
		l := (*resp1)["SubnetSet"].([]interface{})
		for i := 0; i < len(l); i++ {
			item := l[i].(map[string]interface{})
			fmt.Printf("%+v:%+v\n", item["SubnetId"], item["SubnetName"])
		}
	}

