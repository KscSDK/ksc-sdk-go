package main

import (
	"encoding/json"
	"fmt"
	"github.com/KscSDK/ksc-sdk-go/ksc"
	"github.com/KscSDK/ksc-sdk-go/ksc/utils"
	"github.com/KscSDK/ksc-sdk-go/service/cdnv2"
)

func main() {
	ak := "your ak"
	sk := "your sk"

	region := "cn-shanghai-2"

	//debug模式的话 打开这个开关
	svc := cdnv2.SdkNew(ksc.NewClient(ak, sk ,true), &ksc.Config{Region: &region}, &utils.UrlInfo{
		UseSSL: true,
	})

	var resp *map[string]interface{}
	var err error

	//	****************************************访问数据查询GET（GetClientRequestData）***************************************************
	// Parameters:
	//        StartTime       String  获取数据起始时间点，日期格式按ISO8601表示法，北京时间，格式为：YYYY-MM-DDThh:mm+0800，例如： 2016-08-01T21:10+0800
	//        EndTime         String  结束时间需大于起始时间；获取日期格式按照ISO8601表示法，北京时间，格式为：YYYY-MM-DDThh:mm+0800，例如： 2016-08-01T21:15+0800
	//        Metric          String  指定查询指标，支持的类型有：flow：流量，单位为 byte；bandwidth：带宽，单位为 bps；request：请求数，单位为 次；qps：平均每秒请求次数，单位：次，支持批量数据查询，多个类型用逗号（半角）分隔，默认为全部类型
	//        Interval        Long    统计粒度，取值为 5：5分钟粒度；60：1小时粒度；1440：1天粒度；5分钟及以上粒度的带宽值均取该粒度时间段内的5分钟粒度的峰值点,流量、请求数取值为求和值。缺省为5分钟粒度。5分钟粒度查询范围最大跨度为31天，其余粒度查询范围最大跨度为90天
	//        CdnType         String  产品类型，只允许输入一种类型，取值为file：大文件下载，video：音视频点播，page：图片小文件；支持多个查询，多个类型用逗号分隔
	//        Domains         String  当前账户下选择时间段内正在运行状态的加速域名，可输入需要查询的域名，支持批量域名查询，多个域名用逗号（半角）分隔，缺省为全部域名
	//        Areas           String  网民分布区域简称，具体见【使用须知】，支持多区域查询，多个区域用逗号（半角）分隔，缺省为全部区域
	//        Provinces       String  省份区域名称， 枚举类型表见【使用须知】；支持省份批量查询，多个省份用逗号（半角）分隔，缺省为全部省份（当Areas选项有且仅有中国大陆选项时，此类型生效，且支持数据展开），具体Provinces见下附录说明
	//        Isps            String  运营商名称，枚举类型表见【使用须知】；只允许输入一种类型，缺省为全部运营商，（当Areas选项有且仅有中国大陆选项时，此类型生效，且支持数据展开），具体ISP见下附录说明
	//        IpType          String  IP类型，取值为ipv4:ipv4类型数据;ipv6:ipv类型数据；单选，缺省为全部IP类型
	//        Schema          String  协议类型， 取值为http:http协议数据; https:https协议数据；quic：quic协议数据，单选，缺省为全部协议类型
	//        ResultType      String  统计结果数据展示方式，支持不展开和展开，缺省为 不展开（unexpand）数据展开支持全部展开（expand）和按照指定查询维度展开，包含：域名（domain）、区域、省份（province）、运营商（isp）、IP类型（ip）、协议类型（schema） 支持多个条件组合展开，多个条件之间用半角逗号区分开来 —— PS:此处一共有31种组合数据包含关系如下：域名 > 省份 > 运营商 > IP 类型 > 协议类型


	getClientRequestData := make(map[string]interface{})
	getClientRequestData["StartTime"] = "2023-03-02T01:00+0800"
	getClientRequestData["EndTime"] = "2023-03-02T01:10+0800"
	getClientRequestData["Metric"] = "flow"
	getClientRequestData["Interval"] = "5"
	getClientRequestData["CdnType"] = "video"
	getClientRequestData["Areas"] = "CN"
	getClientRequestData["Provinces"] = "beijing"
	getClientRequestData["Isps"] = ""
	getClientRequestData["IpType"] = ""
	getClientRequestData["Schema"] = ""
	getClientRequestData["ResultType"] = ""

	resp, err = svc.GetClientRequestDataGet(&getClientRequestData)
	if err != nil {
		fmt.Println("error:", err.Error())
	}
	if resp != nil {
		str, _ := json.Marshal(&resp)
		fmt.Printf("%+v\n", string(str))
	}

	//	****************************************查询带宽数据POST（GetClientRequestData）***************************************************
	// Parameters:
	//        StartTime       String  获取数据起始时间点，日期格式按ISO8601表示法，北京时间，格式为：YYYY-MM-DDThh:mm+0800，例如： 2016-08-01T21:10+0800
	//        EndTime         String  结束时间需大于起始时间；获取日期格式按照ISO8601表示法，北京时间，格式为：YYYY-MM-DDThh:mm+0800，例如： 2016-08-01T21:15+0800
	//        Metric          String  指定查询指标，支持的类型有：flow：流量，单位为 byte；bandwidth：带宽，单位为 bps；request：请求数，单位为 次；qps：平均每秒请求次数，单位：次，支持批量数据查询，多个类型用逗号（半角）分隔，默认为全部类型
	//        Interval        Long    统计粒度，取值为 5：5分钟粒度；60：1小时粒度；1440：1天粒度；5分钟及以上粒度的带宽值均取该粒度时间段内的5分钟粒度的峰值点,流量、请求数取值为求和值。缺省为5分钟粒度。5分钟粒度查询范围最大跨度为31天，其余粒度查询范围最大跨度为90天
	//        CdnType         String  产品类型，只允许输入一种类型，取值为file：大文件下载，video：音视频点播，page：图片小文件；支持多个查询，多个类型用逗号分隔
	//        Domains         String  当前账户下选择时间段内正在运行状态的加速域名，可输入需要查询的域名，支持批量域名查询，多个域名用逗号（半角）分隔，缺省为全部域名
	//        Areas           String  网民分布区域简称，具体见【使用须知】，支持多区域查询，多个区域用逗号（半角）分隔，缺省为全部区域
	//        Provinces       String  省份区域名称， 枚举类型表见【使用须知】；支持省份批量查询，多个省份用逗号（半角）分隔，缺省为全部省份（当Areas选项有且仅有中国大陆选项时，此类型生效，且支持数据展开），具体Provinces见下附录说明
	//        Isps            String  运营商名称，枚举类型表见【使用须知】；只允许输入一种类型，缺省为全部运营商，（当Areas选项有且仅有中国大陆选项时，此类型生效，且支持数据展开），具体ISP见下附录说明
	//        IpType          String  IP类型，取值为ipv4:ipv4类型数据;ipv6:ipv类型数据；单选，缺省为全部IP类型
	//        Schema          String  协议类型， 取值为http:http协议数据; https:https协议数据；quic：quic协议数据，单选，缺省为全部协议类型
	//        ResultType      String  统计结果数据展示方式，支持不展开和展开，缺省为 不展开（unexpand）数据展开支持全部展开（expand）和按照指定查询维度展开，包含：域名（domain）、区域、省份（province）、运营商（isp）、IP类型（ip）、协议类型（schema） 支持多个条件组合展开，多个条件之间用半角逗号区分开来 —— PS:此处一共有31种组合数据包含关系如下：域名 > 省份 > 运营商 > IP 类型 > 协议类型



	getClientRequestDataPost := make(map[string]interface{})
	getClientRequestDataPost["StartTime"] = "2023-03-02T01:00+0800"
	getClientRequestDataPost["EndTime"] = "2023-03-02T01:10+0800"
	getClientRequestDataPost["Metric"] = "flow"
	getClientRequestDataPost["Interval"] = "5"
	getClientRequestDataPost["CdnType"] = "video"
	getClientRequestDataPost["Areas"] = "CN"
	getClientRequestDataPost["Provinces"] = "beijing"
	getClientRequestDataPost["Isps"] = ""
	getClientRequestDataPost["IpType"] = ""
	getClientRequestDataPost["Schema"] = ""
	getClientRequestDataPost["ResultType"] = ""
	
	resp, err = svc.GetClientRequestDataPost(&getClientRequestDataPost)
	if err != nil {
		fmt.Println("error:", err.Error())
	}
	if resp != nil {
		str, _ := json.Marshal(&resp)
		fmt.Printf("%+v\n", string(str))
	}

	//	****************************************访问数据查询GET（GetServerData）***************************************************
	// Parameters:
	//        StartTime       String  获取数据起始时间点，日期格式按ISO8601表示法，北京时间，格式为：YYYY-MM-DDThh:mm+0800，例如： 2016-08-01T21:10+0800
	//        EndTime         String  结束时间需大于起始时间；获取日期格式按照ISO8601表示法，北京时间，格式为：YYYY-MM-DDThh:mm+0800，例如： 2016-08-01T21:15+0800
	//        Metric          String  指定查询指标，支持的类型有：flow：流量，单位为 byte；bandwidth：带宽，单位为 bps；request：请求数，单位为 次；qps：平均每秒请求次数，单位：次，支持批量数据查询，多个类型用逗号（半角）分隔，默认为全部类型
	//        DataType        String  数据类型，取值为edge：服务数据； origin：回源数据；一次查询只允许查询一个类型，缺省为 edge
	//        Interval        Long    统计粒度，取值为 5：5分钟粒度；60：1小时粒度；1440：1天粒度；5分钟及以上粒度的带宽值均取该粒度时间段内的5分钟粒度的峰值点,流量、请求数取值为求和值。缺省为5分钟粒度。5分钟粒度查询范围最大跨度为31天，其余粒度查询范围最大跨度为90天
	//        CdnType         String  产品类型，只允许输入一种类型，取值为file：大文件下载，video：音视频点播，page：图片小文件；支持多个查询，多个类型用逗号分隔
	//        Domains         String  当前账户下选择时间段内正在运行状态的加速域名，可输入需要查询的域名，支持批量域名查询，多个域名用逗号（半角）分隔，缺省为全部域名
	//        Regions         String  网民分布区域简称，具体见【使用须知】，支持多区域查询，多个区域用逗号（半角）分隔，缺省为全部区域
	//        Schema          String  协议类型， 取值为http:http协议数据; https:https协议数据；quic：quic协议数据，单选，缺省为全部协议类型
	//        ResultType      String  统计结果数据展示方式，支持不展开和展开，缺省为 不展开（unexpand）数据展开支持全部展开（expand）和按照指定查询维度展开，包含：域名（domain）、区域、省份（province）、运营商（isp）、IP类型（ip）、协议类型（schema） 支持多个条件组合展开，多个条件之间用半角逗号区分开来 —— PS:此处一共有31种组合数据包含关系如下：域名 > 省份 > 运营商 > IP 类型 > 协议类型


	getServerDataData := make(map[string]interface{})
	getServerDataData["StartTime"] = "2023-03-02T01:00+0800"
	getServerDataData["EndTime"] = "2023-03-02T01:10+0800"
	getServerDataData["Metric"] = "flow"
	getServerDataData["DataType"] = "edge"
	getServerDataData["Interval"] = "5"
	getServerDataData["CdnType"] = "video"
	getServerDataData["Areas"] = "CN"
	getServerDataData["Schema"] = ""
	getServerDataData["ResultType"] = ""

	resp, err = svc.GetServerDataGet(&getClientRequestData)
	if err != nil {
		fmt.Println("error:", err.Error())
	}
	if resp != nil {
		str, _ := json.Marshal(&resp)
		fmt.Printf("%+v\n", string(str))
	}

	//	****************************************查询带宽数据POST（GetServerData）***************************************************
	// Parameters:
	//        StartTime       String  获取数据起始时间点，日期格式按ISO8601表示法，北京时间，格式为：YYYY-MM-DDThh:mm+0800，例如： 2016-08-01T21:10+0800
	//        EndTime         String  结束时间需大于起始时间；获取日期格式按照ISO8601表示法，北京时间，格式为：YYYY-MM-DDThh:mm+0800，例如： 2016-08-01T21:15+0800
	//        Metric          String  指定查询指标，支持的类型有：flow：流量，单位为 byte；bandwidth：带宽，单位为 bps；request：请求数，单位为 次；qps：平均每秒请求次数，单位：次，支持批量数据查询，多个类型用逗号（半角）分隔，默认为全部类型
	//        DataType        String  数据类型，取值为edge：服务数据； origin：回源数据；一次查询只允许查询一个类型，缺省为 edge
	//        Interval        Long    统计粒度，取值为 5：5分钟粒度；60：1小时粒度；1440：1天粒度；5分钟及以上粒度的带宽值均取该粒度时间段内的5分钟粒度的峰值点,流量、请求数取值为求和值。缺省为5分钟粒度。5分钟粒度查询范围最大跨度为31天，其余粒度查询范围最大跨度为90天
	//        CdnType         String  产品类型，只允许输入一种类型，取值为file：大文件下载，video：音视频点播，page：图片小文件；支持多个查询，多个类型用逗号分隔
	//        Domains         String  当前账户下选择时间段内正在运行状态的加速域名，可输入需要查询的域名，支持批量域名查询，多个域名用逗号（半角）分隔，缺省为全部域名
	//        Regions         String  网民分布区域简称，具体见【使用须知】，支持多区域查询，多个区域用逗号（半角）分隔，缺省为全部区域
	//        Schema          String  协议类型， 取值为http:http协议数据; https:https协议数据；quic：quic协议数据，单选，缺省为全部协议类型
	//        ResultType      String  统计结果数据展示方式，支持不展开和展开，缺省为 不展开（unexpand）数据展开支持全部展开（expand）和按照指定查询维度展开，包含：域名（domain）、区域、省份（province）、运营商（isp）、IP类型（ip）、协议类型（schema） 支持多个条件组合展开，多个条件之间用半角逗号区分开来 —— PS:此处一共有31种组合数据包含关系如下：域名 > 省份 > 运营商 > IP 类型 > 协议类型



	getServerDataDataPost := make(map[string]interface{})
	getServerDataDataPost["StartTime"] = "2023-03-02T01:00+0800"
	getServerDataDataPost["EndTime"] = "2023-03-02T01:10+0800"
	getServerDataDataPost["Metric"] = "flow"
	getServerDataDataPost["DataType"] = "edge"
	getServerDataDataPost["Interval"] = "5"
	getServerDataDataPost["CdnType"] = "video"
	getServerDataDataPost["Areas"] = "CN"
	getServerDataDataPost["Schema"] = ""
	getServerDataDataPost["ResultType"] = ""
	
	resp, err = svc.GetServerDataPost(&getClientRequestDataPost)
	if err != nil {
		fmt.Println("error:", err.Error())
	}
	if resp != nil {
		str, _ := json.Marshal(&resp)
		fmt.Printf("%+v\n", string(str))
	}
}
