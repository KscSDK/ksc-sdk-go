package main

import (
	"encoding/json"
	"fmt"
	"github.com/KscSDK/ksc-sdk-go/ksc"
	"github.com/KscSDK/ksc-sdk-go/ksc/utils"
	"github.com/KscSDK/ksc-sdk-go/service/cdn1"
)

func main() {
	ak := "yoursak"
	sk := "yourssk"
	region := "yoursregion"

	//debug模式的话 打开这个开关
	svc := cdn1.SdkNew(ksc.NewClient(ak, sk ,true), &ksc.Config{Region: &region}, &utils.UrlInfo{
		UseSSL: true,
	})

	var resp *map[string]interface{}
	var err error

	//	****************************************查询域名列表GET（GetCdnDomains）***************************************************
	// Parameters:
	//        PageSize        long    分页大小，默认20，最大500，取值1～500间整数
	//        PageNumber      long    取第几页。默认为1，取值1～10000
	//        DomainName      string  按域名过滤，默认为空，代表当前用户下所有域名
	//        ProjectId		String	查询指定的项目下面的域名，不指定默认为全部
	//        DomainStatus    string  按域名状态过滤，默认为空，代表当前用户下所有域名状态全部
	//        CdnType         string  产品类型，取值为video:音视频点播,file:大文件下载,live:流媒体直播，多个产品类型之间用逗号（半角）间隔，默认为空，代表当前用户下全部产品类型
	//        FuzzyMatch      string  域名过滤是否使用模糊匹配，取值为on：开启，off：关闭，默认为on

	//getCdnDomains := make(map[string]interface{})
	//getCdnDomains["PageSize"] = "20"
	//getCdnDomains["PageNumber"] = "1"
	//getCdnDomains["DomainName"] = ""
	//getCdnDomains["ProjectId"] = ""
	//getCdnDomains["DomainStatus"] = ""
	//getCdnDomains["CdnType"] = "live"
	//getCdnDomains["FuzzyMatch"] = ""
	//
	//resp, err = svc.GetCdnDomainsGet(&getCdnDomains)
	//if err != nil {
	//	fmt.Println("error:", err.Error())
	//}
	//if resp != nil {
	//	str, _ := json.Marshal(&resp)
	//	fmt.Printf("%+v\n", string(str))
	//}

	//	****************************************查询域名列表POST（GetCdnDomains）***************************************************
	// Parameters:
	//        PageSize        long    分页大小，默认20，最大500，取值1～500间整数
	//        PageNumber      long    取第几页。默认为1，取值1～10000
	//        DomainName      string  按域名过滤，默认为空，代表当前用户下所有域名
	//        ProjectId		String	查询指定的项目下面的域名，不指定默认为全部
	//        DomainStatus    string  按域名状态过滤，默认为空，代表当前用户下所有域名状态全部
	//        CdnType         string  产品类型，取值为video:音视频点播,file:大文件下载,live:流媒体直播，多个产品类型之间用逗号（半角）间隔，默认为空，代表当前用户下全部产品类型
	//        FuzzyMatch      string  域名过滤是否使用模糊匹配，取值为on：开启，off：关闭，默认为on

	//getCdnDomains := make(map[string]interface{})
	//getCdnDomains["PageSize"] = "20"
	//getCdnDomains["PageNumber"] = "1"
	//getCdnDomains["DomainName"] = ""
	//getCdnDomains["ProjectId"] = ""
	//getCdnDomains["DomainStatus"] = ""
	//getCdnDomains["CdnType"] = "live"
	//getCdnDomains["FuzzyMatch"] = ""
	//
	//resp, err = svc.GetCdnDomainsPost(&getCdnDomains)
	//if err != nil {
	//	fmt.Println("error:", err.Error())
	//}
	//if resp != nil {
	//	str, _ := json.Marshal(&resp)
	//	fmt.Printf("%+v\n", string(str))
	//}

	//	****************************************新增域名GET（AddCdnDomain ）***************************************************
	// Parameters:
	//        DomainName      string      需要接入CDN的域名
	//        CdnType         string  产品类型，取值为video:音视频点播,file:大文件下载
	//        ProjectId		String	    加速域名所属的项目，非必填项，默认归属为【默认项目】，若输入项目ID，可指定域名归属为已经创建好的项目ID下面
	//        CdnSubType      string      加速业务子类型（业务子类型是为了细分业务，默认不填写）
	//        CdnProtocol     string      客户访问边缘节点的协议。默认http，直播必须填写：http＋flv， hls，rtmp
	//        BillingRegions  string      加速区域，默认CN， 可以输入多个，以逗号间隔
	//        OriginType      string      源站类型 取值：ipaddr、 domain、KS3、ksvideo分别表示：IP源站、域名源站、KS3为源站、金山云视频云源站
	//        OriginProtocol  string      回源协议，取值：http，rtmp，hls，https（当前版本不支持https回源)
	//        OriginPort      integer     可以指定 443, 80。默认值80。
	//        Origin          string      回源地址，可以是IP或域名；IP支持最多20个，以逗号区分，域名只能输入一个
	//        SearchUrl	String	是	用于探测的url,有且只能输入一个。前提是当用户输入了泛域名,客户域名不允许出现kingsoftspark单词，精确域名忽略
	//

	//addCdnDomains := make(map[string]interface{})
	//addCdnDomains["DomainName"] = "ntj14123.test.com"
	//addCdnDomains["CdnType"] = "video"
	//addCdnDomains["ProjectId"] = "0"
	//addCdnDomains["CdnSubType"] = ""
	//addCdnDomains["CdnProtocol"] = "http"
	//addCdnDomains["BillingRegions"] = ""
	//addCdnDomains["OriginType"]="ipaddr"
	//addCdnDomains["OriginProtocol"] = "http"
	//addCdnDomains["OriginPort"] = ""
	//addCdnDomains["Origin"] = "127.111.12.11"
	//addCdnDomains["SearchUrl"] = ""
	//
	//resp, err = svc.AddCdnDomainGet(&addCdnDomains)
	//if err != nil {
	//	fmt.Println("error:", err.Error())
	//}
	//if resp != nil {
	//	str, _ := json.Marshal(&resp)
	//	fmt.Printf("%+v\n", string(str))
	//}

	//	****************************************新增域名POST（AddCdnDomain ）***************************************************
	// Parameters:
	//        DomainName      string      需要接入CDN的域名
	//        CdnType         string  产品类型，取值为video:音视频点播,file:大文件下载
	//        ProjectId		String	    加速域名所属的项目，非必填项，默认归属为【默认项目】，若输入项目ID，可指定域名归属为已经创建好的项目ID下面
	//        CdnSubType      string      加速业务子类型（业务子类型是为了细分业务，默认不填写）
	//        CdnProtocol     string      客户访问边缘节点的协议。默认http，直播必须填写：http＋flv， hls，rtmp
	//        BillingRegions  string      加速区域，默认CN， 可以输入多个，以逗号间隔
	//        OriginType      string      源站类型 取值：ipaddr、 domain、KS3、ksvideo分别表示：IP源站、域名源站、KS3为源站、金山云视频云源站
	//        OriginProtocol  string      回源协议，取值：http，rtmp，hls，https（当前版本不支持https回源)
	//        OriginPort      integer     可以指定 443, 80。默认值80。
	//        Origin          string      回源地址，可以是IP或域名；IP支持最多20个，以逗号区分，域名只能输入一个
	//        SearchUrl	String	是	用于探测的url,有且只能输入一个。前提是当用户输入了泛域名,客户域名不允许出现kingsoftspark单词，精确域名忽略
	//

	//addCdnDomains := make(map[string]interface{})
	//addCdnDomains["DomainName"] = "ntj14122.test.com"
	//addCdnDomains["CdnType"] = "video"
	//addCdnDomains["ProjectId"] = "0"
	//addCdnDomains["CdnSubType"] = ""
	//addCdnDomains["CdnProtocol"] = "http"
	//addCdnDomains["BillingRegions"] = ""
	//addCdnDomains["OriginType"]="ipaddr"
	//addCdnDomains["OriginProtocol"] = "http"
	//addCdnDomains["OriginPort"] = ""
	//addCdnDomains["Origin"] = "127.111.12.11"
	//addCdnDomains["SearchUrl"] = ""
	//
	//resp, err = svc.AddCdnDomainPost(&addCdnDomains)
	//if err != nil {
	//	fmt.Println("error:", err.Error())
	//}
	//if resp != nil {
	//	str, _ := json.Marshal(&resp)
	//	fmt.Printf("%+v\n", string(str))
	//}

	//	****************************************查询域名基础信息GET（GetCdnDomainBasicInfo ）***************************************************
	// Parameters:
	//        DomainId    String      域名ID，只允许输入单个域名ID

	//getCdnDomainBasicInfo := make(map[string]interface{})
	//getCdnDomainBasicInfo["DomainId"] = "2D07WU3"
	//
	//resp, err = svc.GetCdnDomainBasicInfoGet(&getCdnDomainBasicInfo)
	//if err != nil {
	//	fmt.Println("error:", err.Error())
	//}
	//if resp != nil {
	//	str, _ := json.Marshal(&resp)
	//	fmt.Printf("%+v\n", string(str))
	//}

	//	****************************************查询域名基础信息POST（GetCdnDomainBasicInfo ）***************************************************
	// Parameters:
	//        DomainId    String      域名ID，只允许输入单个域名ID

	//getCdnDomainBasicInfo := make(map[string]interface{})
	//getCdnDomainBasicInfo["DomainId"] = "2D07WU3"
	//
	//resp, err = svc.GetCdnDomainBasicInfoPost(&getCdnDomainBasicInfo)
	//if err != nil {
	//	fmt.Println("error:", err.Error())
	//}
	//if resp != nil {
	//	str, _ := json.Marshal(&resp)
	//	fmt.Printf("%+v\n", string(str))
	//}

	//	****************************************查询域名详细配置信息GET（GetDomainConfigs ）***************************************************
	// Parameters:
	//        DomainId    String  域名ID
	//        ConfigList  String  需要查询的配置，多个配置用逗号（半角）分隔，不填代表查询所有
	//                            当前支持  cache_expired、ip、error_page、http_header、optimize、page_compress、
	//                            ignore_query_string、range、referer、src_host、video_seek、waf,notify_url,
	//                            redirect_type,request_auth

	//getDomainConfigs := make(map[string]interface{})
	//getDomainConfigs["DomainId"] = "2D08065"
	//getDomainConfigs["ConfigList"] = "cache_expired,ip"
	//
	//resp, err = svc.GetDomainConfigsGet(&getDomainConfigs)
	//if err != nil {
	//	fmt.Println("error:", err.Error())
	//}
	//if resp != nil {
	//	str, _ := json.Marshal(&resp)
	//	fmt.Printf("%+v\n", string(str))
	//}

	//	****************************************查询域名详细配置信息POST（GetDomainConfigs ）***************************************************
	// Parameters:
	//        DomainId    String  域名ID
	//        ConfigList  String  需要查询的配置，多个配置用逗号（半角）分隔，不填代表查询所有
	//                            当前支持  cache_expired、ip、error_page、http_header、optimize、page_compress、
	//                            ignore_query_string、range、referer、src_host、video_seek、waf,notify_url,
	//                            redirect_type,request_auth

	//getDomainConfigs := make(map[string]interface{})
	//getDomainConfigs["DomainId"] = "2D08RF1"
	//getDomainConfigs["ConfigList"] = ""
	//
	//resp, err = svc.GetDomainConfigsPost(&getDomainConfigs)
	//if err != nil {
	//	fmt.Println("error:", err.Error())
	//}
	//if resp != nil {
	//	str, _ := json.Marshal(&resp)
	//	fmt.Printf("%+v\n", string(str))
	//}

	//	****************************************修改域名基础配置GET（ModifyCdnDomainBasicInfo ）***************************************************
	// Parameters:
	//        DomainId    String  域名ID
	//        Regions     String  加速区域，默认CN， 可以输入多个，以逗号间隔
	//        OriginType  String  源站类型 取值：ipaddr、 domain、KS3分别表示：IP源站、域名源站、KS3为源站。(此项目若输入，Origin必须输入)
	//        OriginPort  Integer 可以指定 443, 80。默认值80。443的话走https回源。当前版本只支持80.
	//        Origin      String  回源地址，可以是IP或域名；IP支持最多20个，以逗号区分，域名只能输入一个。IP与域名不能同时输入。 （此项目若输入，必须保证符合OriginType）
	//

	//modifyCdnDomainBasicInfo := make(map[string]interface{})
	//modifyCdnDomainBasicInfo["DomainId"] = "2D08RF1"
	//modifyCdnDomainBasicInfo["Regions"] = ""
	//modifyCdnDomainBasicInfo["OriginType"] = "ipaddr"
	//modifyCdnDomainBasicInfo["OriginPort"] = ""
	//modifyCdnDomainBasicInfo["Origin"] = "127.111.10.55"
	//
	//resp, err = svc.ModifyCdnDomainBasicInfoGet(&modifyCdnDomainBasicInfo)
	//if err != nil {
	//	fmt.Println("error:", err.Error())
	//}
	//if resp != nil {
	//	str, _ := json.Marshal(&resp)
	//	fmt.Printf("%+v\n", string(str))
	//}

	//	****************************************修改域名基础配置POST（ModifyCdnDomainBasicInfo ）***************************************************
	// Parameters:
	//        DomainId    String  域名ID
	//        Regions     String  加速区域，默认CN， 可以输入多个，以逗号间隔
	//        OriginType  String  源站类型 取值：ipaddr、 domain、KS3分别表示：IP源站、域名源站、KS3为源站。(此项目若输入，Origin必须输入)
	//        OriginPort  Integer 可以指定 443, 80。默认值80。443的话走https回源。当前版本只支持80.
	//        Origin      String  回源地址，可以是IP或域名；IP支持最多20个，以逗号区分，域名只能输入一个。IP与域名不能同时输入。 （此项目若输入，必须保证符合OriginType）
	//

	//modifyCdnDomainBasicInfo := make(map[string]interface{})
	//modifyCdnDomainBasicInfo["DomainId"] = "2D08RF1"
	//modifyCdnDomainBasicInfo["Regions"] = ""
	//modifyCdnDomainBasicInfo["OriginType"] = "ipaddr"
	//modifyCdnDomainBasicInfo["OriginPort"] = ""
	//modifyCdnDomainBasicInfo["Origin"] = "127.111.10.11"
	//
	//resp, err = svc.ModifyCdnDomainBasicInfoPost(&modifyCdnDomainBasicInfo)
	//if err != nil {
	//	fmt.Println("error:", err.Error())
	//}
	//if resp != nil {
	//	str, _ := json.Marshal(&resp)
	//	fmt.Printf("%+v\n", string(str))
	//}

	//	****************************************用于启用、停用某个加速域名GET（StartStopCdnDomain ）***************************************************
	// Parameters:
	//        ActionType  String  操作接口名，系统规定参数 取值：start：启用；stop：停用；
	//        DomainId    String  需要启用或停用CDN服务的域名ID，只允许输入一个域名ID

	//startStopCdnDomain := make(map[string]interface{})
	//startStopCdnDomain["DomainId"] = "2D08RF1"
	//startStopCdnDomain["ActionType"] = "start"
	//
	//resp, err = svc.StartStopCdnDomainGet(&startStopCdnDomain)
	//if err != nil {
	//	fmt.Println("error:", err.Error())
	//}
	//if resp != nil {
	//	str, _ := json.Marshal(&resp)
	//	fmt.Printf("%+v\n", string(str))
	//}

	//	****************************************用于启用、停用某个加速域名POST（StartStopCdnDomain ）***************************************************
	// Parameters:
	//        ActionType  String  操作接口名，系统规定参数 取值：start：启用；stop：停用；
	//        DomainId    String  需要启用或停用CDN服务的域名ID，只允许输入一个域名ID

	//startStopCdnDomain := make(map[string]interface{})
	//startStopCdnDomain["DomainId"] = "2D08RF1"
	//startStopCdnDomain["ActionType"] = "start"
	//
	//resp, err = svc.StartStopCdnDomainPost(&startStopCdnDomain)
	//if err != nil {
	//	fmt.Println("error:", err.Error())
	//}
	//if resp != nil {
	//	str, _ := json.Marshal(&resp)
	//	fmt.Printf("%+v\n", string(str))
	//}

	//	****************************************用于删除用户下已添加的加速域名  此操作只允许删除 DomainStatus 为已停止的域名；GET（DeleteCdnDomain ）***************************************************
	// Parameters:
	//        DomainId    String  域名ID

	//deleteCdnDomain := make(map[string]interface{})
	//deleteCdnDomain["DomainId"] = "2D08RF1"
	//
	//resp, err = svc.DeleteCdnDomainGet(&deleteCdnDomain)
	//if err != nil {
	//	fmt.Println("error:", err.Error())
	//}
	//if resp != nil {
	//	str, _ := json.Marshal(&resp)
	//	fmt.Printf("%+v\n", string(str))
	//}

	//	****************************************用于删除用户下已添加的加速域名  此操作只允许删除 DomainStatus 为已停止的域名；POST（DeleteCdnDomain ）***************************************************
	// Parameters:
	//        DomainId    String  域名ID

	//deleteCdnDomain := make(map[string]interface{})
	//deleteCdnDomain["DomainId"] = "2D08RF1"
	//
	//resp, err = svc.DeleteCdnDomainPost(&deleteCdnDomain)
	//if err != nil {
	//	fmt.Println("error:", err.Error())
	//}
	//if resp != nil {
	//	str, _ := json.Marshal(&resp)
	//	fmt.Printf("%+v\n", string(str))
	//}

	//	****************************************设置回源host功能 GET（SetBackOriginHostConfig）***************************************************
	//  Parameters:
	//        DomainId        String  域名ID
	//        BackOriginHost  String  是自定义回源域名，默认为空，表示不需要修改回源Host

	//setBackOriginHostConfig := make(map[string]interface{})
	//setBackOriginHostConfig["DomainId"] = "2D08RF1"
	//setBackOriginHostConfig["BackOriginHost"] = "www.xxnnxx.test.com"
	//
	//resp, err = svc.SetBackOriginHostConfigGet(&setBackOriginHostConfig)
	//if err != nil {
	//	fmt.Println("error:", err.Error())
	//}
	//if resp != nil {
	//	str, _ := json.Marshal(&resp)
	//	fmt.Printf("%+v\n", string(str))
	//}

	//	****************************************设置回源host功能 POST（SetBackOriginHostConfig）***************************************************
	//  Parameters:
	//        DomainId        String  域名ID
	//        BackOriginHost  String  是自定义回源域名，默认为空，表示不需要修改回源Host

	//setBackOriginHostConfig := make(map[string]interface{})
	//setBackOriginHostConfig["DomainId"] = "2D08RF1"
	//setBackOriginHostConfig["BackOriginHost"] = "www.xxnnxx.test.com"
	//
	//resp, err = svc.SetBackOriginHostConfigPost(&setBackOriginHostConfig)
	//if err != nil {
	//	fmt.Println("error:", err.Error())
	//}
	//if resp != nil {
	//	str, _ := json.Marshal(&resp)
	//	fmt.Printf("%+v\n", string(str))
	//}

	//	****************************************设置加速域名的Refer防盗链 加速域名创建后，默认不开启refer防盗链功能 GET（SetReferProtectionConfig）***************************************************
	// Parameters:
	//        DomainId    String  域名ID
	//        Enable      String  配置是否开启或关闭 取值：on、off，默认值为off关闭。开启时，下述必须项为必填项；关闭时，只更改此标识，忽略后面的项目。
	//        ReferType   String  refer类型，取值：block：黑名单；allow：白名单，开启后必填
	//        ReferList   String  逗号隔开的域名列表
	//        AllowEmpty  String  是否允许空refer访问,取值：on：允许；off：不允许；默认值：on。注：仅当选择白名单时，此项才生效

	//setReferProtectionConfig := make(map[string]interface{})
	//setReferProtectionConfig["DomainId"] = "2D08RF1"
	//setReferProtectionConfig["Enable"] = "off"
	//setReferProtectionConfig["ReferType"] = "block"
	//setReferProtectionConfig["ReferList"] = "www.baidu.com,www.sina.com"
	//setReferProtectionConfig["AllowEmpty"] = "off"
	//
	//resp, err = svc.SetReferProtectionConfigGet(&setReferProtectionConfig)
	//if err != nil {
	//	fmt.Println("error:", err.Error())
	//}
	//if resp != nil {
	//	str, _ := json.Marshal(&resp)
	//	fmt.Printf("%+v\n", string(str))
	//}

	//	****************************************设置加速域名的Refer防盗链 加速域名创建后，默认不开启refer防盗链功能 POST（SetReferProtectionConfig）***************************************************
	// Parameters:
	//        DomainId    String  域名ID
	//        Enable      String  配置是否开启或关闭 取值：on、off，默认值为off关闭。开启时，下述必须项为必填项；关闭时，只更改此标识，忽略后面的项目。
	//        ReferType   String  refer类型，取值：block：黑名单；allow：白名单，开启后必填
	//        ReferList   String  逗号隔开的域名列表
	//        AllowEmpty  String  是否允许空refer访问,取值：on：允许；off：不允许；默认值：on。注：仅当选择白名单时，此项才生效

	//setReferProtectionConfig := make(map[string]interface{})
	//setReferProtectionConfig["DomainId"] = "2D08RF1"
	//setReferProtectionConfig["Enable"] = "off"
	//setReferProtectionConfig["ReferType"] = "block"
	//setReferProtectionConfig["ReferList"] = "www.baidu.com,www.sina.com"
	//setReferProtectionConfig["AllowEmpty"] = "off"
	//
	//resp, err = svc.SetReferProtectionConfigPost(&setReferProtectionConfig)
	//if err != nil {
	//	fmt.Println("error:", err.Error())
	//}
	//if resp != nil {
	//	str, _ := json.Marshal(&resp)
	//	fmt.Printf("%+v\n", string(str))
	//}

	//	****************************************设置加速域名的Ip防盗链 加速域名创建后，默认不开启Ip防盗链功能 GET（SetIpProtectionConfig）***************************************************
	// Parameters:
	//        DomainId    String  域名ID
	//        Enable      String  配置是否开启或关闭 取值：on、off，默认值为off关闭。开启时，下述必须项为必填项；关闭时，只更改此标识，忽略后面的项目。
	//        IpType      String  refer类型，取值：block：黑名单；allow：白名单，开启后必填
	//        IpList      String  逗号隔开的Ip列表

	//setIpProtectionConfig := make(map[string]interface{})
	//setIpProtectionConfig["DomainId"] = "2D08RF1"
	//setIpProtectionConfig["Enable"] = "off"
	//setIpProtectionConfig["IpType"] = "block"
	//setIpProtectionConfig["IpList"] = "127.0.1.54,127.5.6.9"
	//
	//resp, err = svc.SetIpProtectionConfigGet(&setIpProtectionConfig)
	//if err != nil {
	//	fmt.Println("error:", err.Error())
	//}
	//if resp != nil {
	//	str, _ := json.Marshal(&resp)
	//	fmt.Printf("%+v\n", string(str))
	//}

	//	****************************************设置加速域名的Ip防盗链 加速域名创建后，默认不开启Ip防盗链功能 POST（SetIpProtectionConfig）***************************************************
	// Parameters:
	//        DomainId    String  域名ID
	//        Enable      String  配置是否开启或关闭 取值：on、off，默认值为off关闭。开启时，下述必须项为必填项；关闭时，只更改此标识，忽略后面的项目。
	//        IpType      String  refer类型，取值：block：黑名单；allow：白名单，开启后必填
	//        IpList      String  逗号隔开的Ip列表

	//setIpProtectionConfig := make(map[string]interface{})
	//setIpProtectionConfig["DomainId"] = "2D08RF1"
	//setIpProtectionConfig["Enable"] = "off"
	//setIpProtectionConfig["IpType"] = "block"
	//setIpProtectionConfig["IpList"] = "127.0.1.54,127.5.6.9"
	//
	//resp, err = svc.SetIpProtectionConfigPost(&setIpProtectionConfig)
	//if err != nil {
	//	fmt.Println("error:", err.Error())
	//}
	//if resp != nil {
	//	str, _ := json.Marshal(&resp)
	//	fmt.Printf("%+v\n", string(str))
	//}

	//	****************************************设置缓存规则。加速域名创建后，默认缓存规则为空; 更新加速域名的缓存规则为覆盖更新，需要对全部的规则进行修改，不能仅提交需要修改的部分 POST（SetCacheRuleConfig）***************************************************
	// Parameters:
	//        该接口需要输入json格式数据
	//            DomainId    String      域名ID
	//            CacheRules  CacheRule[] 是由CacheRule组成的数组，表示缓存规则列表。注意：该数组是有序的
	//                CacheRule：
	//                    CacheRuleType   String  缓存规则类型 file_suffix 文件后缀 directory 目录 exact 全路径 url_regex 正则表达式
	//                    Value           String  缓存规则的内容，当缓存规则类型为目录时仅允许单条输入，目录必须以/开头且以/结尾；当缓存规则类型为全路径时仅允许单条输入，全路径需输入完整路径，且必须以/开头；
	//                                            当缓存规则类型为正则表达式时仅允许单条输入，且必须输入标准正则表达式；当缓存规则为文件后缀时允许多个输入，文件后缀必须输入英文文件后缀名，多个文件后缀名以逗号（半角）间隔
	//                    CacheTime       Long    缓存时间，以秒为单位
	//                    RespectOrigin   String  是否遵循源站，off表示不遵循，on（默认)表示遵循
	//                    IgnoreNoCache   String  是否忽略源站的no－cache头，on表示忽略，off（默认）表示不忽略。 (本期暂不支持此选项)
	//

	//cacheRule1 := make(map[string]interface{})
	//cacheRule1["CacheRuleType"] = "directory"
	//cacheRule1["Value"] = "/XXX/"
	//cacheRule1["CacheTime"] = "11"
	//cacheRule1["RespectOrigin"] = ""
	//cacheRule1["IgnoreNoCache"] = ""
	//
	//cacheRules := make([]map[string]interface{},0)
	//cacheRules = append(cacheRules, cacheRule1)
	//
	//setCacheRuleConfig := make(map[string]interface{})
	//setCacheRuleConfig["DomainId"] = "2D08RF1"
	//setCacheRuleConfig["CacheRules"] = cacheRules
	//
	//resp, err = svc.SetIpProtectionConfigPost(&setCacheRuleConfig)
	//if err != nil {
	//	fmt.Println("error:", err.Error())
	//}
	//if resp != nil {
	//	str, _ := json.Marshal(&resp)
	//	fmt.Printf("%+v\n", string(str))
	//}

	//	****************************************设置加速域名的测试URL GET（SetTestUrlConfig）***************************************************
	// Parameters:
	//        DomainId    String  域名ID
	//        TestUrl     String  测试URL列表，逗号间隔，默认为空

	//setTestUrlConfig := make(map[string]interface{})
	//setTestUrlConfig["DomainId"] = "2D08RF1"
	//setTestUrlConfig["TestUrl"] = "www.xinfei.cn/1.html"
	//
	//resp, err = svc.SetTestUrlConfigGet(&setTestUrlConfig)
	//if err != nil {
	//	fmt.Println("error:", err.Error())
	//}
	//if resp != nil {
	//	str, _ := json.Marshal(&resp)
	//	fmt.Printf("%+v\n", string(str))
	//}

	//	****************************************设置加速域名的测试URL POST（SetTestUrlConfig）***************************************************
	// Parameters:
	//        DomainId    String  域名ID
	//        TestUrl     String  测试URL列表，逗号间隔，默认为空

	//setTestUrlConfig := make(map[string]interface{})
	//setTestUrlConfig["DomainId"] = "2D08RF1"
	//setTestUrlConfig["TestUrl"] = "www.xinfei.cn/1.html"
	//resp, err = svc.SetTestUrlConfigPost(&setTestUrlConfig)
	//
	//if err != nil {
	//	fmt.Println("error:", err.Error())
	//}
	//if resp != nil {
	//	str, _ := json.Marshal(&resp)
	//	fmt.Printf("%+v\n", string(str))
	//}

	//	****************************************设置高级回源策略 POST（SetOriginAdvancedConfig）***************************************************
	//	OriginLine为default默认源的线路，是必填项，其他几个源都是选填项。OriginLine不能重复填写。开启高级回源策略后，会关闭掉基础配置中的回源配置
	// Parameters:
	//        该接口需要输入json格式数据
	//            DomainId                String                  域名ID
	//            Enable                  String                  配置高级回源策略的开启或关闭 取值: on、off。注意：开启后会关闭掉基础配置中的的回源配置。默认值关闭。开启时，下述必须项为必填项；关闭时，只更改此标识，忽略后面的项目
	//            OriginType              String                  源站类型 取值：ipaddr、domain分别表示：IP源站、域名源站
	//            OriginPolicy            String                  rr: 轮询； quality: 按质量最优的topN来轮询回源
	//            OriginPolicyBestCount   Long                    当OriginPolicy是quality时，该项必填。取值1-10
	//

	//setOriginAdvancedConfig := make(map[string]interface{})
	//setOriginAdvancedConfig["DomainId"] = "2D08RF1"
	//setOriginAdvancedConfig["Enable"] = "on"
	//setOriginAdvancedConfig["OriginPolicy"] = "quality"
	//setOriginAdvancedConfig["OriginPolicyBestCount"] = 1
	//setOriginAdvancedConfig["OriginType"] = "domain"
	//setOriginAdvancedConfig["Origin"] = "www.b.xunfei.cn"
	//setOriginAdvancedConfig["BackupOriginType"] = "ipaddr"
	//setOriginAdvancedConfig["BackupOrigin"] = "1.1.1.1,2.2.2.2"
	//
	//resp, err = svc.SetOriginAdvancedConfigPost(&setOriginAdvancedConfig)
	//
	//if err != nil {
	//	fmt.Println("error:", err.Error())
	//}
	//if resp != nil {
	//	str, _ := json.Marshal(&resp)
	//	fmt.Printf("%+v\n", string(str))
	//}

	//	****************************************设置备注信息 GET（SetRemarkConfig）***************************************************
	// Parameters:
	//        DomainId    String  域名ID
	//        Remark      String  备注信息，默认为空

	//setRemarkConfig := make(map[string]interface{})
	//setRemarkConfig["DomainId"] = "2D08RF1"
	//setRemarkConfig["Remark"] = "xxxxxx"
	//
	//resp, err = svc.SetRemarkConfigGet(&setRemarkConfig)
	//if err != nil {
	//	fmt.Println("error:", err.Error())
	//}
	//if resp != nil {
	//	str, _ := json.Marshal(&resp)
	//	fmt.Printf("%+v\n", string(str))
	//}

	//	****************************************设置备注信息 POST（SetRemarkConfig）***************************************************
	// Parameters:
	//        DomainId    String  域名ID
	//        Remark      String  备注信息，默认为空

	//setRemarkConfig := make(map[string]interface{})
	//setRemarkConfig["DomainId"] = "2D08RF1"
	//setRemarkConfig["Remark"] = "xxxxxx"
	//
	//resp, err = svc.SetRemarkConfigPost(&setRemarkConfig)
	//
	//if err != nil {
	//	fmt.Println("error:", err.Error())
	//}
	//if resp != nil {
	//	str, _ := json.Marshal(&resp)
	//	fmt.Printf("%+v\n", string(str))
	//}

	//	****************************************设置时间戳+共享秘钥防盗链 GET（SetRequestAuthConfig）***************************************************
	// Parameters:
	//        DomainId        String  域名ID
	//        Enable          String  配置是否开启或关闭 取值：on、off，默认值为off关闭。开启时，下述必须项为必填项；关闭时，只更改此标识，忽略后面的项目。
	//        AuthType        String  类型，取值：typeA, typeB
	//        Key1            String  主密钥
	//        Key2            String  副密钥  多个逗号分隔
	//        ExpirationTime  Long    有效时间

	//setRequestAuthConfig := make(map[string]interface{})
	//setRequestAuthConfig["DomainId"] = "2D08RF1"
	//setRequestAuthConfig["Enable"] = "on"
	//setRequestAuthConfig["AuthType"] = "typeB"
	//setRequestAuthConfig["Key1"] = "111111"
	//setRequestAuthConfig["Key2"] = "222222,33333"
	//setRequestAuthConfig["ExpirationTime"] = "1000"
	//
	//resp, err = svc.SetRequestAuthConfigGet(&setRequestAuthConfig)
	//if err != nil {
	//	fmt.Println("error:", err.Error())
	//}
	//if resp != nil {
	//	str, _ := json.Marshal(&resp)
	//	fmt.Printf("%+v\n", string(str))
	//}

	//	****************************************设置时间戳+共享秘钥防盗链 POST（SetRequestAuthConfig）***************************************************
	// Parameters:
	//        DomainId        String  域名ID
	//        Enable          String  配置是否开启或关闭 取值：on、off，默认值为off关闭。开启时，下述必须项为必填项；关闭时，只更改此标识，忽略后面的项目。
	//        AuthType        String  类型，取值：typeA, typeB
	//        Key1            String  主密钥
	//        Key2            String  副密钥  多个逗号分隔
	//        ExpirationTime  Long    有效时间

	//setRequestAuthConfig := make(map[string]interface{})
	//setRequestAuthConfig["DomainId"] = "2D08RF1"
	//setRequestAuthConfig["Enable"] = "on"
	//setRequestAuthConfig["AuthType"] = "typeB"
	//setRequestAuthConfig["Key1"] = "111111"
	//setRequestAuthConfig["Key2"] = "222222,33333"
	//setRequestAuthConfig["ExpirationTime"] = "1000"
	//
	//resp, err = svc.SetRequestAuthConfigPost(&setRequestAuthConfig)
	//
	//if err != nil {
	//	fmt.Println("error:", err.Error())
	//}
	//if resp != nil {
	//	str, _ := json.Marshal(&resp)
	//	fmt.Printf("%+v\n", string(str))
	//}

	//	****************************************设置强制跳转 GET（SetForceRedirectConfig）***************************************************
	// Parameters:
	//        DomainId      string      域名ID
	//        RedirectType  string  配置强制跳转类型, 取值: off、 https，默认为off 。其中https表示http → https，当选择https时需保证域名已配置证书

	//setForceRedirectConfig := make(map[string]interface{})
	//setForceRedirectConfig["DomainId"] = "2D08RF1"
	//setForceRedirectConfig["RedirectType"] = "off"
	//
	//resp, err = svc.SetForceRedirectConfigGet(&setForceRedirectConfig)
	//if err != nil {
	//	fmt.Println("error:", err.Error())
	//}
	//if resp != nil {
	//	str, _ := json.Marshal(&resp)
	//	fmt.Printf("%+v\n", string(str))
	//}

	//	****************************************设置强制跳转 POST（SetForceRedirectConfig）***************************************************
	// Parameters:
	//        DomainId      string      域名ID
	//        RedirectType  string  配置强制跳转类型, 取值: off、 https，默认为off 。其中https表示http → https，当选择https时需保证域名已配置证书

	//setForceRedirectConfig := make(map[string]interface{})
	//setForceRedirectConfig["DomainId"] = "2D08RF1"
	//setForceRedirectConfig["RedirectType"] = "off"
	//
	//resp, err = svc.SetForceRedirectConfigPost(&setForceRedirectConfig)
	//
	//if err != nil {
	//	fmt.Println("error:", err.Error())
	//}
	//if resp != nil {
	//	str, _ := json.Marshal(&resp)
	//	fmt.Printf("%+v\n", string(str))
	//}

	//	****************************************设置HTTP 2.0 GET（SetHttp2OptionConfig）***************************************************
	// Parameters:
	//        DomainId      string      域名ID
	//        Enable  string  配置HTTP 2.0功能的开启或关闭 取值：on、off ，默认为off ；开启需保证域名已配置证书。

	//setHttp2OptionConfig := make(map[string]interface{})
	//setHttp2OptionConfig["DomainId"] = "2D08RF1"
	//setHttp2OptionConfig["Enable"] = "off"
	//
	//resp, err = svc.SetHttp2OptionConfigGet(&setHttp2OptionConfig)
	//if err != nil {
	//	fmt.Println("error:", err.Error())
	//}
	//if resp != nil {
	//	str, _ := json.Marshal(&resp)
	//	fmt.Printf("%+v\n", string(str))
	//}

	//	****************************************设置HTTP 2.0 POST（SetHttp2OptionConfig）***************************************************
	// Parameters:
	//        DomainId      string      域名ID
	//        Enable  string  配置HTTP 2.0功能的开启或关闭 取值：on、off ，默认为off ；开启需保证域名已配置证书。
	//
	//setHttp2OptionConfig := make(map[string]interface{})
	//setHttp2OptionConfig["DomainId"] = "2D08RF1"
	//setHttp2OptionConfig["Enable"] = "off"
	//
	//resp, err = svc.SetHttp2OptionConfigPost(&setHttp2OptionConfig)
	//
	//if err != nil {
	//	fmt.Println("error:", err.Error())
	//}
	//if resp != nil {
	//	str, _ := json.Marshal(&resp)
	//	fmt.Printf("%+v\n", string(str))
	//}

	//	****************************************设置智能压缩 GET（SetPageCompressConfig）***************************************************
	// Parameters:
	//        DomainId      string      域名ID
	//        Enable  string  	配置智能压缩的开启或关闭 取值：on、off ，默认为off 。

	//setPageCompressConfig := make(map[string]interface{})
	//setPageCompressConfig["DomainId"] = "2D08RF1"
	//setPageCompressConfig["Enable"] = "off"
	//
	//resp, err = svc.SetPageCompressConfigGet(&setPageCompressConfig)
	//if err != nil {
	//	fmt.Println("error:", err.Error())
	//}
	//if resp != nil {
	//	str, _ := json.Marshal(&resp)
	//	fmt.Printf("%+v\n", string(str))
	//}

	//	****************************************设置智能压缩 POST（SetPageCompressConfig）***************************************************
	// Parameters:
	//        DomainId      string      域名ID
	//        Enable  string  	配置智能压缩的开启或关闭 取值：on、off ，默认为off 。

	//setPageCompressConfig := make(map[string]interface{})
	//setPageCompressConfig["DomainId"] = "2D08RF1"
	//setPageCompressConfig["Enable"] = "off"
	//
	//resp, err = svc.SetPageCompressConfigPost(&setPageCompressConfig)
	//
	//if err != nil {
	//	fmt.Println("error:", err.Error())
	//}
	//if resp != nil {
	//	str, _ := json.Marshal(&resp)
	//	fmt.Printf("%+v\n", string(str))
	//}

	//	****************************************自定义错误页面 POST（SetErrorPageConfig）***************************************************
	// Parameters:
	//        DomainId      string      域名ID
	//        ErrorPages  List<ErrorPage>  	由ErrorPage组成的数组，表示自定义错误页面列表。注意：该数组是有序的，如果一个相同状态码在数组里有配置子集，则以最后面的子集为准。
	//	ErrorPage:
	//	    ErrorHttpCode	String	错误的状态码。
	//        CustomPageUrl	String	自定义发生错误后跳转的页面URL。注：需要检验URL的合法性，如果URL不是以https://或者http://开头，则报错，提示输入url有误。
	//
	//errorPage := make(map[string]interface{})
	//errorPage["ErrorHttpCode"] = "400"
	//errorPage["CustomPageUrl"] = "https://www.test.com/error400.html"
	//
	//errorPages := make([]map[string]interface{},0)
	//errorPages = append(errorPages, errorPage);
	//
	//setErrorPageConfig := make(map[string]interface{})
	//setErrorPageConfig["DomainId"] = "2D08RF1"
	//setErrorPageConfig["ErrorPages"] = errorPages
	//
	//resp, err = svc.SetErrorPageConfigPost(&setErrorPageConfig)
	//if err != nil {
	//	fmt.Println("error:", err.Error())
	//}
	//if resp != nil {
	//	str, _ := json.Marshal(&resp)
	//	fmt.Printf("%+v\n", string(str))
	//}

	//	****************************************查询带宽数据GET（GetBandwidthData）***************************************************
	// Parameters:
	//        DomainIds       String  域名ID，缺省为当前产品类型下的全部域名，可输入需要查询的域名ID，支持批量域名查询，多个域名ID用逗号（半角）分隔
	//        Regions         String  区域名称，缺省为 CN;  取值为CN:中国大陆，HK：香港，TW：台湾，AS：亚洲其他，NA：北美洲，SA：南美洲，EU：欧洲，AU：大洋洲，AF：非洲，支持多区域查询，多个区域用逗号（半角）分隔
	//        CdnType         string  产品类型，取值为video:音视频点播,file:大文件下载,live:流媒体直播
	//        StartTime       String  获取数据起始时间点，日期格式按ISO8601表示法，北京时间，格式为：YYYY-MM-DDThh:mm+0800，例如：2016-08-01T21:14+0800
	//        EndTime         String  结束时间需大于起始时间；获取日期格式按照ISO8601表示法，北京时间，格式为：YYYY-MM-DDThh:mm+0800，例如： 2016-08-01T21:14+0800
	//        ResultType      String    取值为0：多域名多区域数据做合并；1：每个域名每个区域的数据分别返回
	//        Granularity     String    统计粒度，取值为 5（默认）：5分钟粒度；10：10分钟粒度；20：20分钟粒度；60：1小时粒度；240：4小时粒度；480：8小时粒度；1440：1天粒度；以上粒度的带宽值均取该粒度时间段的峰值
	//        DataType        String  数据类型， 取值为edge:边缘数据; origin:回源数据; 支持多类型选择，多个类型用逗号（半角）分隔，缺省为 edge
	//		ProtocolType	否	String	协议类型， 取值为http:htts协议数据; https:https协议数据

	//getBandwidthData := make(map[string]interface{})
	//getBandwidthData["StartTime"] = "2021-05-02T01:00+0800"
	//getBandwidthData["EndTime"] = "2021-05-02T01:10+0800"
	//getBandwidthData["CdnType"] = "page"
	//getBandwidthData["ResultType"] = "0"
	//getBandwidthData["Granularity"] = "5"
	//getBandwidthData["DataType"] = "edge"
	//getBandwidthData["DomainIds"] = "2D083FU"
	//
	//resp, err = svc.GetBandwidthDataGet(&getBandwidthData)
	//if err != nil {
	//	fmt.Println("error:", err.Error())
	//}
	//if resp != nil {
	//	str, _ := json.Marshal(&resp)
	//	fmt.Printf("%+v\n", string(str))
	//}

	//	****************************************查询带宽数据POST（GetBandwidthData）***************************************************
	// Parameters:
	//        DomainIds       String  域名ID，缺省为当前产品类型下的全部域名，可输入需要查询的域名ID，支持批量域名查询，多个域名ID用逗号（半角）分隔
	//        Regions         String  区域名称，缺省为 CN;  取值为CN:中国大陆，HK：香港，TW：台湾，AS：亚洲其他，NA：北美洲，SA：南美洲，EU：欧洲，AU：大洋洲，AF：非洲，支持多区域查询，多个区域用逗号（半角）分隔
	//        CdnType         string  产品类型，取值为video:音视频点播,file:大文件下载,live:流媒体直播
	//        StartTime       String  获取数据起始时间点，日期格式按ISO8601表示法，北京时间，格式为：YYYY-MM-DDThh:mm+0800，例如：2016-08-01T21:14+0800
	//        EndTime         String  结束时间需大于起始时间；获取日期格式按照ISO8601表示法，北京时间，格式为：YYYY-MM-DDThh:mm+0800，例如： 2016-08-01T21:14+0800
	//        ResultType      String    取值为0：多域名多区域数据做合并；1：每个域名每个区域的数据分别返回
	//        Granularity     String    统计粒度，取值为 5（默认）：5分钟粒度；10：10分钟粒度；20：20分钟粒度；60：1小时粒度；240：4小时粒度；480：8小时粒度；1440：1天粒度；以上粒度的带宽值均取该粒度时间段的峰值
	//        DataType        String  数据类型， 取值为edge:边缘数据; origin:回源数据; 支持多类型选择，多个类型用逗号（半角）分隔，缺省为 edge
	//		ProtocolType	否	String	协议类型， 取值为http:htts协议数据; https:https协议数据

	//getBandwidthData := make(map[string]interface{})
	//getBandwidthData["StartTime"] = "2021-05-02T01:00+0800"
	//getBandwidthData["EndTime"] = "2021-05-02T01:10+0800"
	//getBandwidthData["CdnType"] = "page"
	//getBandwidthData["ResultType"] = "0"
	//getBandwidthData["Granularity"] = "5"
	//getBandwidthData["DataType"] = "edge"
	//getBandwidthData["DomainIds"] = "2D083FU"
	//
	//resp, err = svc.GetBandwidthDataPost(&getBandwidthData)
	//if err != nil {
	//	fmt.Println("error:", err.Error())
	//}
	//if resp != nil {
	//	str, _ := json.Marshal(&resp)
	//	fmt.Printf("%+v\n", string(str))
	//}

	//	****************************************查询流量数据GET（GetFlowData）***************************************************
	// Parameters:
	//        DomainIds       String  域名ID，缺省为当前产品类型下的全部域名，可输入需要查询的域名ID，支持批量域名查询，多个域名ID用逗号（半角）分隔
	//        Regions         String  区域名称，缺省为 CN;  取值为CN:中国大陆，HK：香港，TW：台湾，AS：亚洲其他，NA：北美洲，SA：南美洲，EU：欧洲，AU：大洋洲，AF：非洲，支持多区域查询，多个区域用逗号（半角）分隔
	//        CdnType         string  产品类型，取值为video:音视频点播,file:大文件下载,live:流媒体直播
	//        StartTime       String  获取数据起始时间点，日期格式按ISO8601表示法，北京时间，格式为：YYYY-MM-DDThh:mm+0800，例如：2016-08-01T21:14+0800
	//        EndTime         String  结束时间需大于起始时间；获取日期格式按照ISO8601表示法，北京时间，格式为：YYYY-MM-DDThh:mm+0800，例如： 2016-08-01T21:14+0800
	//        ResultType      String    取值为0：多域名多区域数据做合并；1：每个域名每个区域的数据分别返回
	//        Granularity     String    统计粒度，取值为 5（默认）：5分钟粒度；10：10分钟粒度；20：20分钟粒度；60：1小时粒度；240：4小时粒度；480：8小时粒度；1440：1天粒度；以上粒度的带宽值均取该粒度时间段的峰值
	//        DataType        String  数据类型， 取值为edge:边缘数据; origin:回源数据; 支持多类型选择，多个类型用逗号（半角）分隔，缺省为 edge
	//		  ProtocolType	否	String	协议类型， 取值为http:htts协议数据; https:https协议数据

	//getFlowData := make(map[string]interface{})
	//getFlowData["StartTime"] = "2021-05-02T01:00+0800"
	//getFlowData["EndTime"] = "2021-05-02T01:10+0800"
	//getFlowData["CdnType"] = "page"
	//getFlowData["ResultType"] = "0"
	//getFlowData["Granularity"] = "5"
	//getFlowData["DataType"] = "edge"
	//getFlowData["DomainIds"] = "2D083FU"
	//
	//resp, err = svc.GetFlowDataGet(&getFlowData)
	//if err != nil {
	//	fmt.Println("error:", err.Error())
	//}
	//if resp != nil {
	//	str, _ := json.Marshal(&resp)
	//	fmt.Printf("%+v\n", string(str))
	//}

	//	****************************************查询流量数据POST（GetFlowData）***************************************************
	// Parameters:
	//        DomainIds       String  域名ID，缺省为当前产品类型下的全部域名，可输入需要查询的域名ID，支持批量域名查询，多个域名ID用逗号（半角）分隔
	//        Regions         String  区域名称，缺省为 CN;  取值为CN:中国大陆，HK：香港，TW：台湾，AS：亚洲其他，NA：北美洲，SA：南美洲，EU：欧洲，AU：大洋洲，AF：非洲，支持多区域查询，多个区域用逗号（半角）分隔
	//        CdnType         string  产品类型，取值为video:音视频点播,file:大文件下载,live:流媒体直播
	//        StartTime       String  获取数据起始时间点，日期格式按ISO8601表示法，北京时间，格式为：YYYY-MM-DDThh:mm+0800，例如：2016-08-01T21:14+0800
	//        EndTime         String  结束时间需大于起始时间；获取日期格式按照ISO8601表示法，北京时间，格式为：YYYY-MM-DDThh:mm+0800，例如： 2016-08-01T21:14+0800
	//        ResultType      String    取值为0：多域名多区域数据做合并；1：每个域名每个区域的数据分别返回
	//        Granularity     String    统计粒度，取值为 5（默认）：5分钟粒度；10：10分钟粒度；20：20分钟粒度；60：1小时粒度；240：4小时粒度；480：8小时粒度；1440：1天粒度；以上粒度的带宽值均取该粒度时间段的峰值
	//        DataType        String  数据类型， 取值为edge:边缘数据; origin:回源数据; 支持多类型选择，多个类型用逗号（半角）分隔，缺省为 edge
	//		  ProtocolType	否	String	协议类型， 取值为http:htts协议数据; https:https协议数据

	//getFlowData := make(map[string]interface{})
	//getFlowData["StartTime"] = "2021-05-02T01:00+0800"
	//getFlowData["EndTime"] = "2021-05-02T01:10+0800"
	//getFlowData["CdnType"] = "page"
	//getFlowData["ResultType"] = "0"
	//getFlowData["Granularity"] = "5"
	//getFlowData["DataType"] = "edge"
	//getFlowData["DomainIds"] = "2D083FU"
	//
	//resp, err = svc.GetFlowDataPost(&getFlowData)
	//if err != nil {
	//	fmt.Println("error:", err.Error())
	//}
	//if resp != nil {
	//	str, _ := json.Marshal(&resp)
	//	fmt.Printf("%+v\n", string(str))
	//}

	//	****************************************请求数查询GET（GetPvData）***************************************************
	// Parameters:
	//        DomainIds       String  域名ID，缺省为当前产品类型下的全部域名，可输入需要查询的域名ID，支持批量域名查询，多个域名ID用逗号（半角）分隔
	//        Regions         String  区域名称，缺省为 CN;  取值为CN:中国大陆，HK：香港，TW：台湾，AS：亚洲其他，NA：北美洲，SA：南美洲，EU：欧洲，AU：大洋洲，AF：非洲，支持多区域查询，多个区域用逗号（半角）分隔
	//        CdnType         string  产品类型，取值为video:音视频点播,file:大文件下载,live:流媒体直播
	//        StartTime       String  获取数据起始时间点，日期格式按ISO8601表示法，北京时间，格式为：YYYY-MM-DDThh:mm+0800，例如：2016-08-01T21:14+0800
	//        EndTime         String  结束时间需大于起始时间；获取日期格式按照ISO8601表示法，北京时间，格式为：YYYY-MM-DDThh:mm+0800，例如： 2016-08-01T21:14+0800
	//        ResultType      String    取值为0：多域名多区域数据做合并；1：每个域名每个区域的数据分别返回
	//        Granularity     String    统计粒度，取值为 5（默认）：5分钟粒度；10：10分钟粒度；20：20分钟粒度；60：1小时粒度；240：4小时粒度；480：8小时粒度；1440：1天粒度；以上粒度的带宽值均取该粒度时间段的峰值
	//        DataType        String  数据类型， 取值为edge:边缘数据; origin:回源数据; 支持多类型选择，多个类型用逗号（半角）分隔，缺省为 edge
	//		ProtocolType	否	String	协议类型， 取值为http:htts协议数据; https:https协议数据

	//getPvData := make(map[string]interface{})
	//getPvData["StartTime"] = "2021-05-02T01:00+0800"
	//getPvData["EndTime"] = "2021-05-02T01:10+0800"
	//getPvData["CdnType"] = "page"
	//getPvData["ResultType"] = "0"
	//getPvData["Granularity"] = "5"
	//getPvData["DataType"] = "edge"
	//getPvData["DomainIds"] = "2D083FU"
	//
	//resp, err = svc.GetPvDataGet(&getPvData)
	//if err != nil {
	//	fmt.Println("error:", err.Error())
	//}
	//if resp != nil {
	//	str, _ := json.Marshal(&resp)
	//	fmt.Printf("%+v\n", string(str))
	//}

	//	****************************************请求数查询POST（GetPvData）***************************************************
	// Parameters:
	//        DomainIds       String  域名ID，缺省为当前产品类型下的全部域名，可输入需要查询的域名ID，支持批量域名查询，多个域名ID用逗号（半角）分隔
	//        Regions         String  区域名称，缺省为 CN;  取值为CN:中国大陆，HK：香港，TW：台湾，AS：亚洲其他，NA：北美洲，SA：南美洲，EU：欧洲，AU：大洋洲，AF：非洲，支持多区域查询，多个区域用逗号（半角）分隔
	//        CdnType         string  产品类型，取值为video:音视频点播,file:大文件下载,live:流媒体直播
	//        StartTime       String  获取数据起始时间点，日期格式按ISO8601表示法，北京时间，格式为：YYYY-MM-DDThh:mm+0800，例如：2016-08-01T21:14+0800
	//        EndTime         String  结束时间需大于起始时间；获取日期格式按照ISO8601表示法，北京时间，格式为：YYYY-MM-DDThh:mm+0800，例如： 2016-08-01T21:14+0800
	//        ResultType      String    取值为0：多域名多区域数据做合并；1：每个域名每个区域的数据分别返回
	//        Granularity     String    统计粒度，取值为 5（默认）：5分钟粒度；10：10分钟粒度；20：20分钟粒度；60：1小时粒度；240：4小时粒度；480：8小时粒度；1440：1天粒度；以上粒度的带宽值均取该粒度时间段的峰值
	//        DataType        String  数据类型， 取值为edge:边缘数据; origin:回源数据; 支持多类型选择，多个类型用逗号（半角）分隔，缺省为 edge
	//		ProtocolType	否	String	协议类型， 取值为http:htts协议数据; https:https协议数据

	//getPvData := make(map[string]interface{})
	//getPvData["StartTime"] = "2021-05-02T01:00+0800"
	//getPvData["EndTime"] = "2021-05-02T01:10+0800"
	//getPvData["CdnType"] = "page"
	//getPvData["ResultType"] = "0"
	//getPvData["Granularity"] = "5"
	//getPvData["DataType"] = "edge"
	//getPvData["DomainIds"] = "2D083FU"
	//
	//resp, err = svc.GetPvDataPost(&getPvData)
	//if err != nil {
	//	fmt.Println("error:", err.Error())
	//}
	//if resp != nil {
	//	str, _ := json.Marshal(&resp)
	//	fmt.Printf("%+v\n", string(str))
	//}

	//	**************************************** 命中率详情查询 GET（GetHitRateDetailedData）***************************************************
	// Parameters:
	//        StartTime       String  获取数据起始时间点，日期格式按ISO8601表示法，北京时间，格式为：YYYY-MM-DDThh:mm+0800，例如：2016-08-01T21:14+0800
	//        EndTime         String  结束时间需大于起始时间；获取日期格式按照ISO8601表示法，北京时间，格式为：YYYY-MM-DDThh:mm+0800，例如： 2016-08-01T21:14+0800
	//        CdnType         string  产品类型，取值为video:音视频点播,file:大文件下载,live:流媒体直播
	//        DomainIds       String  域名ID，缺省为当前产品类型下的全部域名，可输入需要查询的域名ID，支持批量域名查询，多个域名ID用逗号（半角）分隔
	//        ResultType      String  取值为0：多域名多区域数据做合并；1：每个域名每个区域的数据分别返回
	//        Granularity     String  统计粒度，取值为 5（默认）：5分钟粒度；10：10分钟粒度；20：20分钟粒度；60：1小时粒度；240：4小时粒度；480：8小时粒度；1440：1天粒度；以上粒度的带宽值均取该粒度时间段的峰值
	//        HitType         String  数据类型， 取值为flowhitrate:流量命中率; reqhitrate:请求数命中率; 支持多类型选择，多个类型用逗号（半角）分隔，缺省为reqhitrate
	//

	//getHitRateDetailedData := make(map[string]interface{})
	//getHitRateDetailedData["StartTime"] = "2021-05-02T01:00+0800"
	//getHitRateDetailedData["EndTime"] = "2021-05-02T01:10+0800"
	//getHitRateDetailedData["CdnType"] = "page"
	//getHitRateDetailedData["ResultType"] = "0"
	//getHitRateDetailedData["Granularity"] = "5"
	//getHitRateDetailedData["DomainIds"] = "2D083FU"
	//getHitRateDetailedData["HitType"] = "reqhitrate"
	//
	//resp, err = svc.GetHitRateDetailedDataGet(&getHitRateDetailedData)
	//if err != nil {
	//	fmt.Println("error:", err.Error())
	//}
	//if resp != nil {
	//	str, _ := json.Marshal(&resp)
	//	fmt.Printf("%+v\n", string(str))
	//}

	//	**************************************** 命中率详情查询 POST（GetHitRateDetailedData）***************************************************
	// Parameters:
	//        StartTime       String  获取数据起始时间点，日期格式按ISO8601表示法，北京时间，格式为：YYYY-MM-DDThh:mm+0800，例如：2016-08-01T21:14+0800
	//        EndTime         String  结束时间需大于起始时间；获取日期格式按照ISO8601表示法，北京时间，格式为：YYYY-MM-DDThh:mm+0800，例如： 2016-08-01T21:14+0800
	//        CdnType         string  产品类型，取值为video:音视频点播,file:大文件下载,live:流媒体直播
	//        DomainIds       String  域名ID，缺省为当前产品类型下的全部域名，可输入需要查询的域名ID，支持批量域名查询，多个域名ID用逗号（半角）分隔
	//        ResultType      String  取值为0：多域名多区域数据做合并；1：每个域名每个区域的数据分别返回
	//        Granularity     String  统计粒度，取值为 5（默认）：5分钟粒度；10：10分钟粒度；20：20分钟粒度；60：1小时粒度；240：4小时粒度；480：8小时粒度；1440：1天粒度；以上粒度的带宽值均取该粒度时间段的峰值
	//        HitType         String  数据类型， 取值为flowhitrate:流量命中率; reqhitrate:请求数命中率; 支持多类型选择，多个类型用逗号（半角）分隔，缺省为reqhitrate
	//

	//getHitRateDetailedData := make(map[string]interface{})
	//getHitRateDetailedData["StartTime"] = "2021-05-12T01:00+0800"
	//getHitRateDetailedData["EndTime"] = "2021-05-12T01:10+0800"
	//getHitRateDetailedData["CdnType"] = "page"
	//getHitRateDetailedData["ResultType"] = "0"
	//getHitRateDetailedData["Granularity"] = "5"
	//getHitRateDetailedData["DomainIds"] = "2D083FU"
	//getHitRateDetailedData["HitType"] = "reqhitrate"
	//
	//resp, err = svc.GetHitRateDetailedDataPost(&getHitRateDetailedData)
	//if err != nil {
	//	fmt.Println("error:", err.Error())
	//}
	//if resp != nil {
	//	str, _ := json.Marshal(&resp)
	//	fmt.Printf("%+v\n", string(str))
	//}

	//	**************************************** 命中率查询（饼图） GET（GetHitRateData）***************************************************
	// Parameters:
	//        StartTime   String  获取数据起始时间点，日期格式按ISO8601表示法，北京时间，格式为：YYYY-MM-DDThh:mm+0800，例如：2016-08-01T21:14+0800
	//        EndTime     String  结束时间需大于起始时间；获取日期格式按照ISO8601表示法，北京时间，格式为：YYYY-MM-DDThh:mm+0800，例如： 2016-08-01T21:14+0800
	//        CdnType     String  产品类型，只允许输入一种类型，取值为video:音视频点播,file:大文件下载,live:流媒体直播
	//        DomainIds   String  域名ID，缺省为当前产品类型下的全部域名，可输入需要查询的域名ID，支持批量域名查询，多个域名ID用逗号（半角）分隔
	//
	//
	//getHitRateData := make(map[string]interface{})
	//getHitRateData["StartTime"] = "2021-05-02T01:00+0800"
	//getHitRateData["EndTime"] = "2021-05-02T01:10+0800"
	//getHitRateData["CdnType"] = "page"
	//getHitRateData["DomainIds"] = "2D083FU"
	//
	//resp, err = svc.GetHitRateDataGet(&getHitRateData)
	//if err != nil {
	//	fmt.Println("error:", err.Error())
	//}
	//if resp != nil {
	//	str, _ := json.Marshal(&resp)
	//	fmt.Printf("%+v\n", string(str))
	//}

	//	**************************************** 命中率查询（饼图） POST（GetHitRateData）***************************************************
	// Parameters:
	//        StartTime   String  获取数据起始时间点，日期格式按ISO8601表示法，北京时间，格式为：YYYY-MM-DDThh:mm+0800，例如：2016-08-01T21:14+0800
	//        EndTime     String  结束时间需大于起始时间；获取日期格式按照ISO8601表示法，北京时间，格式为：YYYY-MM-DDThh:mm+0800，例如： 2016-08-01T21:14+0800
	//        CdnType     String  产品类型，只允许输入一种类型，取值为video:音视频点播,file:大文件下载,live:流媒体直播
	//        DomainIds   String  域名ID，缺省为当前产品类型下的全部域名，可输入需要查询的域名ID，支持批量域名查询，多个域名ID用逗号（半角）分隔
	//

	//getHitRateData := make(map[string]interface{})
	//getHitRateData["StartTime"] = "2021-05-02T01:00+0800"
	//getHitRateData["EndTime"] = "2021-05-02T01:10+0800"
	//getHitRateData["CdnType"] = "page"
	//getHitRateData["DomainIds"] = "2D083FU"
	//
	//resp, err = svc.GetHitRateDataPost(&getHitRateData)
	//if err != nil {
	//	fmt.Println("error:", err.Error())
	//}
	//if resp != nil {
	//	str, _ := json.Marshal(&resp)
	//	fmt.Printf("%+v\n", string(str))
	//}

	//	**************************************** 省份+运营商流量查询 GET（GetProvinceAndIspFlowData）***************************************************
	// Parameters:
	//        StartTime   String  获取数据起始时间点，日期格式按ISO8601表示法，北京时间，格式为：YYYY-MM-DDThh:mm+0800，例如：2016-08-01T21:14+0800
	//        EndTime     String  结束时间需大于起始时间；获取日期格式按照ISO8601表示法，北京时间，格式为：YYYY-MM-DDThh:mm+0800，例如： 2016-08-01T21:14+0800
	//        CdnType     String  产品类型，只允许输入一种类型，取值为video:音视频点播,file:大文件下载,live:流媒体直播
	//        DomainIds   String  域名ID，缺省为当前产品类型下的全部域名，可输入需要查询的域名ID，支持批量域名查询，多个域名ID用逗号（半角）分隔
	//        Provinces   String  省份区域名称， 取值详见枚举列表，支持多省份区域查询，多个省份区域用逗号（半角）分隔，缺省为全部省份区域
	//        Isps        String  运营商名称，取值详见枚举列表，支持多运营商查询，多个运营商用逗号（半角）分隔，缺省为全部运营商
	//        ResultType  String  取值为0：多域名多省份区域多运营商数据做合并；1：每个域名每个省份区域的每个运营商数据分别返回
	//        Granularity String  统计粒度，取值为 5（默认）：5分钟粒度；10：10分钟粒度；20：20分钟粒度；60：1小时粒度；240：4小时粒度；480：8小时粒度；1440：1天粒度
	//

	//getProvinceAndIspFlowData := make(map[string]interface{})
	//getProvinceAndIspFlowData["StartTime"] = "2021-05-02T01:00+0800"
	//getProvinceAndIspFlowData["EndTime"] = "2021-05-02T01:10+0800"
	//getProvinceAndIspFlowData["CdnType"] = "page"
	//getProvinceAndIspFlowData["DomainIds"] = "2D083FU"
	//getProvinceAndIspFlowData["ResultType"] = 0
	//getProvinceAndIspFlowData["Granularity"] = 5
	//getProvinceAndIspFlowData["Provinces"] = "beijing"
	//getProvinceAndIspFlowData["Isps"] = "CM"
	//
	//resp, err = svc.GetProvinceAndIspFlowDataGet(&getProvinceAndIspFlowData)
	//if err != nil {
	//	fmt.Println("error:", err.Error())
	//}
	//if resp != nil {
	//	str, _ := json.Marshal(&resp)
	//	fmt.Printf("%+v\n", string(str))
	//}

	//	**************************************** 省份+运营商流量查询 POST（GetProvinceAndIspFlowData）***************************************************
	// Parameters:
	//        StartTime   String  获取数据起始时间点，日期格式按ISO8601表示法，北京时间，格式为：YYYY-MM-DDThh:mm+0800，例如：2016-08-01T21:14+0800
	//        EndTime     String  结束时间需大于起始时间；获取日期格式按照ISO8601表示法，北京时间，格式为：YYYY-MM-DDThh:mm+0800，例如： 2016-08-01T21:14+0800
	//        CdnType     String  产品类型，只允许输入一种类型，取值为video:音视频点播,file:大文件下载,live:流媒体直播
	//        DomainIds   String  域名ID，缺省为当前产品类型下的全部域名，可输入需要查询的域名ID，支持批量域名查询，多个域名ID用逗号（半角）分隔
	//        Provinces   String  省份区域名称， 取值详见枚举列表，支持多省份区域查询，多个省份区域用逗号（半角）分隔，缺省为全部省份区域
	//        Isps        String  运营商名称，取值详见枚举列表，支持多运营商查询，多个运营商用逗号（半角）分隔，缺省为全部运营商
	//        ResultType  String  取值为0：多域名多省份区域多运营商数据做合并；1：每个域名每个省份区域的每个运营商数据分别返回
	//        Granularity String  统计粒度，取值为 5（默认）：5分钟粒度；10：10分钟粒度；20：20分钟粒度；60：1小时粒度；240：4小时粒度；480：8小时粒度；1440：1天粒度
	//

	//getProvinceAndIspFlowData := make(map[string]interface{})
	//getProvinceAndIspFlowData["StartTime"] = "2021-05-02T01:00+0800"
	//getProvinceAndIspFlowData["EndTime"] = "2021-05-02T01:10+0800"
	//getProvinceAndIspFlowData["CdnType"] = "page"
	//getProvinceAndIspFlowData["DomainIds"] = "2D083FU"
	//getProvinceAndIspFlowData["ResultType"] = 0
	//getProvinceAndIspFlowData["Granularity"] = 5
	//getProvinceAndIspFlowData["Provinces"] = "beijing,shandong"
	//getProvinceAndIspFlowData["Isps"] = "CM"
	//
	//resp, err = svc.GetProvinceAndIspFlowDataPost(&getProvinceAndIspFlowData)
	//if err != nil {
	//	fmt.Println("error:", err.Error())
	//}
	//if resp != nil {
	//	str, _ := json.Marshal(&resp)
	//	fmt.Printf("%+v\n", string(str))
	//}

	//	**************************************** 省份+运营商带宽查询 GET（GetProvinceAndIspBandwidthData）***************************************************
	// Parameters:
	//        StartTime   String  获取数据起始时间点，日期格式按ISO8601表示法，北京时间，格式为：YYYY-MM-DDThh:mm+0800，例如：2016-08-01T21:14+0800
	//        EndTime     String  结束时间需大于起始时间；获取日期格式按照ISO8601表示法，北京时间，格式为：YYYY-MM-DDThh:mm+0800，例如： 2016-08-01T21:14+0800
	//        CdnType     String  产品类型，只允许输入一种类型，取值为video:音视频点播,file:大文件下载,live:流媒体直播
	//        DomainIds   String  域名ID，缺省为当前产品类型下的全部域名，可输入需要查询的域名ID，支持批量域名查询，多个域名ID用逗号（半角）分隔
	//        Provinces   String  省份区域名称， 取值详见枚举列表，支持多省份区域查询，多个省份区域用逗号（半角）分隔，缺省为全部省份区域
	//        Isps        String  运营商名称，取值详见枚举列表，支持多运营商查询，多个运营商用逗号（半角）分隔，缺省为全部运营商
	//        ResultType  String  取值为0：多域名多省份区域多运营商数据做合并；1：每个域名每个省份区域的每个运营商数据分别返回
	//        Granularity String  统计粒度，取值为 5（默认）：5分钟粒度；10：10分钟粒度；20：20分钟粒度；60：1小时粒度；240：4小时粒度；480：8小时粒度；1440：1天粒度
	//

	//getProvinceAndIspBandwidthData := make(map[string]interface{})
	//getProvinceAndIspBandwidthData["StartTime"] = "2021-05-02T01:00+0800"
	//getProvinceAndIspBandwidthData["EndTime"] = "2021-05-02T01:10+0800"
	//getProvinceAndIspBandwidthData["CdnType"] = "page"
	//getProvinceAndIspBandwidthData["DomainIds"] = "2D083FU"
	//getProvinceAndIspBandwidthData["ResultType"] = 0
	//getProvinceAndIspBandwidthData["Granularity"] = 5
	//getProvinceAndIspBandwidthData["Provinces"] = "beijing"
	//getProvinceAndIspBandwidthData["Isps"] = "CM"
	//
	//resp, err = svc.GetProvinceAndIspBandwidthDataGet(&getProvinceAndIspBandwidthData)
	//if err != nil {
	//	fmt.Println("error:", err.Error())
	//}
	//if resp != nil {
	//	str, _ := json.Marshal(&resp)
	//	fmt.Printf("%+v\n", string(str))
	//}

	//	**************************************** 省份+运营商带宽查询 POST（GetProvinceAndIspBandwidthData）***************************************************
	// Parameters:
	//        StartTime   String  获取数据起始时间点，日期格式按ISO8601表示法，北京时间，格式为：YYYY-MM-DDThh:mm+0800，例如：2016-08-01T21:14+0800
	//        EndTime     String  结束时间需大于起始时间；获取日期格式按照ISO8601表示法，北京时间，格式为：YYYY-MM-DDThh:mm+0800，例如： 2016-08-01T21:14+0800
	//        CdnType     String  产品类型，只允许输入一种类型，取值为video:音视频点播,file:大文件下载,live:流媒体直播
	//        DomainIds   String  域名ID，缺省为当前产品类型下的全部域名，可输入需要查询的域名ID，支持批量域名查询，多个域名ID用逗号（半角）分隔
	//        Provinces   String  省份区域名称， 取值详见枚举列表，支持多省份区域查询，多个省份区域用逗号（半角）分隔，缺省为全部省份区域
	//        Isps        String  运营商名称，取值详见枚举列表，支持多运营商查询，多个运营商用逗号（半角）分隔，缺省为全部运营商
	//        ResultType  String  取值为0：多域名多省份区域多运营商数据做合并；1：每个域名每个省份区域的每个运营商数据分别返回
	//        Granularity String  统计粒度，取值为 5（默认）：5分钟粒度；10：10分钟粒度；20：20分钟粒度；60：1小时粒度；240：4小时粒度；480：8小时粒度；1440：1天粒度
	//

	//getProvinceAndIspBandwidthData := make(map[string]interface{})
	//getProvinceAndIspBandwidthData["StartTime"] = "2021-05-02T01:00+0800"
	//getProvinceAndIspBandwidthData["EndTime"] = "2021-05-02T01:10+0800"
	//getProvinceAndIspBandwidthData["CdnType"] = "page"
	//getProvinceAndIspBandwidthData["DomainIds"] = "2D083FU"
	//getProvinceAndIspBandwidthData["ResultType"] = 0
	//getProvinceAndIspBandwidthData["Granularity"] = 5
	//getProvinceAndIspBandwidthData["Provinces"] = "beijing"
	//getProvinceAndIspBandwidthData["Isps"] = "CM"
	//
	//resp, err = svc.GetProvinceAndIspBandwidthDataPost(&getProvinceAndIspBandwidthData)
	//if err != nil {
	//	fmt.Println("error:", err.Error())
	//}
	//if resp != nil {
	//	str, _ := json.Marshal(&resp)
	//	fmt.Printf("%+v\n", string(str))
	//}

	//	**************************************** 状态码统计(饼图) GET（GetHttpCodeData）***************************************************
	// Parameters:
	//        StartTime   String  获取数据起始时间点，日期格式按ISO8601表示法，北京时间，格式为：YYYY-MM-DDThh:mm+0800，例如：2016-08-01T21:14+0800
	//        EndTime     String  结束时间需大于起始时间；获取日期格式按照ISO8601表示法，北京时间，格式为：YYYY-MM-DDThh:mm+0800，例如： 2016-08-01T21:14+0800
	//        CdnType     String  产品类型，只允许输入一种类型，取值为video:音视频点播,file:大文件下载,live:流媒体直播
	//        DomainIds   String  域名ID，缺省为当前产品类型下的全部域名，可输入需要查询的域名ID，支持批量域名查询，多个域名ID用逗号（半角）分隔
	//

	//getHttpCodeData := make(map[string]interface{})
	//getHttpCodeData["StartTime"] = "2021-05-02T01:00+0800"
	//getHttpCodeData["EndTime"] = "2021-05-02T01:10+0800"
	//getHttpCodeData["CdnType"] = "page"
	//getHttpCodeData["DomainIds"] = "2D083FU"
	//
	//resp, err = svc.GetHttpCodeDataGet(&getHttpCodeData)
	//if err != nil {
	//	fmt.Println("error:", err.Error())
	//}
	//if resp != nil {
	//	str, _ := json.Marshal(&resp)
	//	fmt.Printf("%+v\n", string(str))
	//}

	//	**************************************** 状态码统计(饼图) POST（GetHttpCodeData）***************************************************
	// Parameters:
	//        StartTime   String  获取数据起始时间点，日期格式按ISO8601表示法，北京时间，格式为：YYYY-MM-DDThh:mm+0800，例如：2016-08-01T21:14+0800
	//        EndTime     String  结束时间需大于起始时间；获取日期格式按照ISO8601表示法，北京时间，格式为：YYYY-MM-DDThh:mm+0800，例如： 2016-08-01T21:14+0800
	//        CdnType     String  产品类型，只允许输入一种类型，取值为video:音视频点播,file:大文件下载,live:流媒体直播
	//        DomainIds   String  域名ID，缺省为当前产品类型下的全部域名，可输入需要查询的域名ID，支持批量域名查询，多个域名ID用逗号（半角）分隔
	//

	//getHttpCodeData := make(map[string]interface{})
	//getHttpCodeData["StartTime"] = "2021-05-02T01:00+0800"
	//getHttpCodeData["EndTime"] = "2021-05-02T01:10+0800"
	//getHttpCodeData["CdnType"] = "page"
	//getHttpCodeData["DomainIds"] = "2D083FU"
	//
	//resp, err = svc.GetHttpCodeDataPost(&getHttpCodeData)
	//if err != nil {
	//	fmt.Println("error:", err.Error())
	//}
	//if resp != nil {
	//	str, _ := json.Marshal(&resp)
	//	fmt.Printf("%+v\n", string(str))
	//}

	//	**************************************** 状态码详情统计 GET（GetHttpCodeDetailedData）***************************************************
	// Parameters:
	//        StartTime   String  获取数据起始时间点，日期格式按ISO8601表示法，北京时间，格式为：YYYY-MM-DDThh:mm+0800，例如：2016-08-01T21:14+0800
	//        EndTime     String  结束时间需大于起始时间；获取日期格式按照ISO8601表示法，北京时间，格式为：YYYY-MM-DDThh:mm+0800，例如： 2016-08-01T21:14+0800
	//        CdnType     String  产品类型，只允许输入一种类型，取值为video:音视频点播,file:大文件下载,live:流媒体直播
	//        DomainIds   String  域名ID，缺省为当前产品类型下的全部域名，可输入需要查询的域名ID，支持批量域名查询，多个域名ID用逗号（半角）分隔
	//        ResultType  String  取值为0：多域名多省份区域多运营商数据做合并；1：每个域名每个省份区域的每个运营商数据分别返回
	//        Granularity String  统计粒度，取值为 5（默认）：5分钟粒度；10：10分钟粒度；20：20分钟粒度；60：1小时粒度；240：4小时粒度；480：8小时粒度；1440：1天粒度
	//

	//getHttpCodeDetailedData := make(map[string]interface{})
	//getHttpCodeDetailedData["StartTime"] = "2021-05-02T01:00+0800"
	//getHttpCodeDetailedData["EndTime"] = "2021-05-02T01:10+0800"
	//getHttpCodeDetailedData["CdnType"] = "page"
	//getHttpCodeDetailedData["DomainIds"] = "2D083FU"
	//getHttpCodeDetailedData["ResultType"] = 0
	//getHttpCodeDetailedData["Granularity"] = 5
	//
	//resp, err = svc.GetHttpCodeDetailedDataGet(&getHttpCodeDetailedData)
	//if err != nil {
	//	fmt.Println("error:", err.Error())
	//}
	//if resp != nil {
	//	str, _ := json.Marshal(&resp)
	//	fmt.Printf("%+v\n", string(str))
	//}

	//	**************************************** 状态码详情统计 POST（GetHttpCodeDetailedData）***************************************************
	// Parameters:
	//        StartTime   String  获取数据起始时间点，日期格式按ISO8601表示法，北京时间，格式为：YYYY-MM-DDThh:mm+0800，例如：2016-08-01T21:14+0800
	//        EndTime     String  结束时间需大于起始时间；获取日期格式按照ISO8601表示法，北京时间，格式为：YYYY-MM-DDThh:mm+0800，例如： 2016-08-01T21:14+0800
	//        CdnType     String  产品类型，只允许输入一种类型，取值为video:音视频点播,file:大文件下载,live:流媒体直播
	//        DomainIds   String  域名ID，缺省为当前产品类型下的全部域名，可输入需要查询的域名ID，支持批量域名查询，多个域名ID用逗号（半角）分隔
	//        ResultType  String  取值为0：多域名多省份区域多运营商数据做合并；1：每个域名每个省份区域的每个运营商数据分别返回
	//        Granularity String  统计粒度，取值为 5（默认）：5分钟粒度；10：10分钟粒度；20：20分钟粒度；60：1小时粒度；240：4小时粒度；480：8小时粒度；1440：1天粒度
	//
	//
	//getHttpCodeDetailedData := make(map[string]interface{})
	//getHttpCodeDetailedData["StartTime"] = "2021-05-02T01:00+0800"
	//getHttpCodeDetailedData["EndTime"] = "2021-05-02T01:10+0800"
	//getHttpCodeDetailedData["CdnType"] = "page"
	//getHttpCodeDetailedData["DomainIds"] = "2D083FU"
	//getHttpCodeDetailedData["ResultType"] = 0
	//getHttpCodeDetailedData["Granularity"] = 5
	//
	//resp, err = svc.GetHttpCodeDetailedDataPost(&getHttpCodeDetailedData)
	//if err != nil {
	//	fmt.Println("error:", err.Error())
	//}
	//if resp != nil {
	//	str, _ := json.Marshal(&resp)
	//	fmt.Printf("%+v\n", string(str))
	//}

	//	**************************************** top url 查询 GET（GetTopUrlData）***************************************************
	// Parameters:
	//        StartTime   String  获取数据起始时间点，日期格式按ISO8601表示法，北京时间，格式为：YYYY-MM-DDThh:mm+0800，例如：2016-08-01T21:14+0800
	//        EndTime     String  结束时间需大于起始时间；获取日期格式按照ISO8601表示法，北京时间，格式为：YYYY-MM-DDThh:mm+0800，例如： 2016-08-01T21:14+0800
	//        CdnType     String  产品类型，只允许输入一种类型，取值为video:音视频点播,file:大文件下载,live:流媒体直播
	//        DomainIds   String  域名ID，缺省为当前产品类型下的全部域名，可输入需要查询的域名ID，支持批量域名查询，多个域名ID用逗号（半角）分隔
	//        LimitN      String  热门Url条数，取值为1-200，最大200，默认100

	//getTopUrlData := make(map[string]interface{})
	//getTopUrlData["StartTime"] = "2021-05-02T01:00+0800"
	//getTopUrlData["EndTime"] = "2021-05-02T01:10+0800"
	//getTopUrlData["CdnType"] = "page"
	//getTopUrlData["DomainIds"] = "2D083FU"
	//getTopUrlData["LimitN"] = 20
	//
	//resp, err = svc.GetTopUrlDataGet(&getTopUrlData)
	//if err != nil {
	//	fmt.Println("error:", err.Error())
	//}
	//if resp != nil {
	//	str, _ := json.Marshal(&resp)
	//	fmt.Printf("%+v\n", string(str))
	//}

	//	**************************************** top url 查询 POST（GetTopUrlData）***************************************************
	// Parameters:
	//        StartTime   String  获取数据起始时间点，日期格式按ISO8601表示法，北京时间，格式为：YYYY-MM-DDThh:mm+0800，例如：2016-08-01T21:14+0800
	//        EndTime     String  结束时间需大于起始时间；获取日期格式按照ISO8601表示法，北京时间，格式为：YYYY-MM-DDThh:mm+0800，例如： 2016-08-01T21:14+0800
	//        CdnType     String  产品类型，只允许输入一种类型，取值为video:音视频点播,file:大文件下载,live:流媒体直播
	//        DomainIds   String  域名ID，缺省为当前产品类型下的全部域名，可输入需要查询的域名ID，支持批量域名查询，多个域名ID用逗号（半角）分隔
	//        LimitN      String  热门Url条数，取值为1-200，最大200，默认100

	//getTopUrlData := make(map[string]interface{})
	//getTopUrlData["StartTime"] = "2021-05-02T01:00+0800"
	//getTopUrlData["EndTime"] = "2021-05-02T01:10+0800"
	//getTopUrlData["CdnType"] = "page"
	//getTopUrlData["DomainIds"] = "2D083FU"
	//getTopUrlData["LimitN"] = 20
	//
	//resp, err = svc.GetTopUrlDataPost(&getTopUrlData)
	//if err != nil {
	//	fmt.Println("error:", err.Error())
	//}
	//if resp != nil {
	//	str, _ := json.Marshal(&resp)
	//	fmt.Printf("%+v\n", string(str))
	//}

	//	**************************************** 用户区域统计 GET（GetAreaData）***************************************************
	// Parameters:
	//        StartTime   String  获取数据起始时间点，日期格式按ISO8601表示法，北京时间，格式为：YYYY-MM-DDThh:mm+0800，例如：2016-08-01T21:14+0800
	//        EndTime     String  结束时间需大于起始时间；获取日期格式按照ISO8601表示法，北京时间，格式为：YYYY-MM-DDThh:mm+0800，例如： 2016-08-01T21:14+0800
	//        CdnType     String  产品类型，只允许输入一种类型，取值为video:音视频点播,file:大文件下载,live:流媒体直播
	//        DomainIds   String  域名ID，缺省为当前产品类型下的全部域名，可输入需要查询的域名ID，支持批量域名查询，多个域名ID用逗号（半角）分隔
	//

	//getAreaData := make(map[string]interface{})
	//getAreaData["StartTime"] = "2021-05-02T01:00+0800"
	//getAreaData["EndTime"] = "2021-05-02T01:10+0800"
	//getAreaData["CdnType"] = "page"
	//getAreaData["DomainIds"] = "2D083FU"
	//
	//resp, err = svc.GetAreaDataGet(&getAreaData)
	//if err != nil {
	//	fmt.Println("error:", err.Error())
	//}
	//if resp != nil {
	//	str, _ := json.Marshal(&resp)
	//	fmt.Printf("%+v\n", string(str))
	//}

	//	**************************************** 用户区域统计 POST（GetAreaData）***************************************************
	// Parameters:
	//        StartTime   String  获取数据起始时间点，日期格式按ISO8601表示法，北京时间，格式为：YYYY-MM-DDThh:mm+0800，例如：2016-08-01T21:14+0800
	//        EndTime     String  结束时间需大于起始时间；获取日期格式按照ISO8601表示法，北京时间，格式为：YYYY-MM-DDThh:mm+0800，例如： 2016-08-01T21:14+0800
	//        CdnType     String  产品类型，只允许输入一种类型，取值为video:音视频点播,file:大文件下载,live:流媒体直播
	//        DomainIds   String  域名ID，缺省为当前产品类型下的全部域名，可输入需要查询的域名ID，支持批量域名查询，多个域名ID用逗号（半角）分隔
	//

	//getAreaData := make(map[string]interface{})
	//getAreaData["StartTime"] = "2021-05-02T01:00+0800"
	//getAreaData["EndTime"] = "2021-05-02T01:10+0800"
	//getAreaData["CdnType"] = "page"
	//getAreaData["DomainIds"] = "2D083FU"
	//
	//resp, err = svc.GetAreaDataPost(&getAreaData)
	//if err != nil {
	//	fmt.Println("error:", err.Error())
	//}
	//if resp != nil {
	//	str, _ := json.Marshal(&resp)
	//	fmt.Printf("%+v\n", string(str))
	//}

	//	**************************************** 运营商占比统计 GET（GetIspData）***************************************************
	// Parameters:
	//        StartTime   String  获取数据起始时间点，日期格式按ISO8601表示法，北京时间，格式为：YYYY-MM-DDThh:mm+0800，例如：2016-08-01T21:14+0800
	//        EndTime     String  结束时间需大于起始时间；获取日期格式按照ISO8601表示法，北京时间，格式为：YYYY-MM-DDThh:mm+0800，例如： 2016-08-01T21:14+0800
	//        CdnType     String  产品类型，只允许输入一种类型，取值为video:音视频点播,file:大文件下载,live:流媒体直播
	//        DomainIds   String  域名ID，缺省为当前产品类型下的全部域名，可输入需要查询的域名ID，支持批量域名查询，多个域名ID用逗号（半角）分隔
	//

	//getIspData := make(map[string]interface{})
	//getIspData["StartTime"] = "2021-05-02T01:00+0800"
	//getIspData["EndTime"] = "2021-05-02T01:10+0800"
	//getIspData["CdnType"] = "page"
	//getIspData["DomainIds"] = "2D083FU"
	//
	//resp, err = svc.GetIspDataGet(&getIspData)
	//if err != nil {
	//	fmt.Println("error:", err.Error())
	//}
	//if resp != nil {
	//	str, _ := json.Marshal(&resp)
	//	fmt.Printf("%+v\n", string(str))
	//}

	//	**************************************** 运营商占比统计 POST（GetIspData）***************************************************
	// Parameters:
	//        StartTime   String  获取数据起始时间点，日期格式按ISO8601表示法，北京时间，格式为：YYYY-MM-DDThh:mm+0800，例如：2016-08-01T21:14+0800
	//        EndTime     String  结束时间需大于起始时间；获取日期格式按照ISO8601表示法，北京时间，格式为：YYYY-MM-DDThh:mm+0800，例如： 2016-08-01T21:14+0800
	//        CdnType     String  产品类型，只允许输入一种类型，取值为video:音视频点播,file:大文件下载,live:流媒体直播
	//        DomainIds   String  域名ID，缺省为当前产品类型下的全部域名，可输入需要查询的域名ID，支持批量域名查询，多个域名ID用逗号（半角）分隔
	//

	//getIspData := make(map[string]interface{})
	//getIspData["StartTime"] = "2021-05-02T01:00+0800"
	//getIspData["EndTime"] = "2021-05-02T01:10+0800"
	//getIspData["CdnType"] = "page"
	//getIspData["DomainIds"] = "2D083FU"
	//
	//resp, err = svc.GetIspDataPost(&getIspData)
	//if err != nil {
	//	fmt.Println("error:", err.Error())
	//}
	//if resp != nil {
	//	str, _ := json.Marshal(&resp)
	//	fmt.Printf("%+v\n", string(str))
	//}

	//	**************************************** 域名排行查询 GET（GetDomainRankingListData）***************************************************
	// Parameters:
	//        StartTime   String  获取数据起始时间点，日期格式按ISO8601表示法，北京时间，格式为：YYYY-MM-DDThh:mm+0800，例如：2016-08-01T21:14+0800
	//        EndTime     String  结束时间需大于起始时间；获取日期格式按照ISO8601表示法，北京时间，格式为：YYYY-MM-DDThh:mm+0800，例如： 2016-08-01T21:14+0800
	//        CdnType     String  产品类型，只允许输入一种类型，取值为video:音视频点播,file:大文件下载,live:流媒体直播
	//

	//getDomainRankingListData := make(map[string]interface{})
	//getDomainRankingListData["StartTime"] = "2021-05-02T01:00+0800"
	//getDomainRankingListData["EndTime"] = "2021-05-02T01:10+0800"
	//getDomainRankingListData["CdnType"] = "video"
	//
	//resp, err = svc.GetDomainRankingListDataGet(&getDomainRankingListData)
	//if err != nil {
	//	fmt.Println("error:", err.Error())
	//}
	//if resp != nil {
	//	str, _ := json.Marshal(&resp)
	//	fmt.Printf("%+v\n", string(str))
	//}

	//	**************************************** 域名排行查询 POST（GetDomainRankingListData）***************************************************
	// Parameters:
	//        StartTime   String  获取数据起始时间点，日期格式按ISO8601表示法，北京时间，格式为：YYYY-MM-DDThh:mm+0800，例如：2016-08-01T21:14+0800
	//        EndTime     String  结束时间需大于起始时间；获取日期格式按照ISO8601表示法，北京时间，格式为：YYYY-MM-DDThh:mm+0800，例如： 2016-08-01T21:14+0800
	//        CdnType     String  产品类型，只允许输入一种类型，取值为video:音视频点播,file:大文件下载,live:流媒体直播
	//

	//getDomainRankingListData := make(map[string]interface{})
	//getDomainRankingListData["StartTime"] = "2021-05-02T01:00+0800"
	//getDomainRankingListData["EndTime"] = "2021-05-02T01:10+0800"
	//getDomainRankingListData["CdnType"] = "video"
	//
	//resp, err = svc.GetDomainRankingListDataPost(&getDomainRankingListData)
	//if err != nil {
	//	fmt.Println("error:", err.Error())
	//}
	//if resp != nil {
	//	str, _ := json.Marshal(&resp)
	//	fmt.Printf("%+v\n", string(str))
	//}

	//	**************************************** 直播按流维度查询流量 GET（GetLiveFlowDataByStream）***************************************************
	// Parameters:
	//        StartTime   String  获取数据起始时间点，日期格式按ISO8601表示法，北京时间，格式为：YYYY-MM-DDThh:mm+0800，例如：2016-08-01T21:14+0800
	//        EndTime     String  结束时间需大于起始时间；获取日期格式按照ISO8601表示法，北京时间，格式为：YYYY-MM-DDThh:mm+0800，例如： 2016-08-01T21:14+0800
	//        StreamUrls   String  流名，支持批量查询，多个流名用逗号（半角）分隔
	//        Regions     String  计费区域名称，取值为CN:中国大陆，HK：香港，TW：台湾，AS：亚洲其他，NA：北美洲，SA：南美洲，EU：欧洲，AU：大洋洲，AF：非洲，支持多计费区域查询，多个区域用逗号（半角）分隔，缺省为 CN
	//        ResultType  String  取值为0：多域名多省份区域多运营商数据做合并；1：每个域名每个省份区域的每个运营商数据分别返回
	//        Granularity String  统计粒度，取值为 5（默认）：5分钟粒度；10：10分钟粒度；20：20分钟粒度；60：1小时粒度；240：4小时粒度；480：8小时粒度；1440：1天粒度
	//

	//getLiveFlowDataByStream := make(map[string]interface{})
	//getLiveFlowDataByStream["StartTime"] = "2021-05-02T01:00+0800"
	//getLiveFlowDataByStream["EndTime"] = "2021-05-02T01:10+0800"
	//getLiveFlowDataByStream["StreamUrls"] = "rtmp://realflv3.plu.cn/live/ce781fecb2f6447d82d03590e520872f"
	//getLiveFlowDataByStream["Regions"] = "CN"
	//getLiveFlowDataByStream["ResultType"] = 0
	//getLiveFlowDataByStream["Granularity"] = 5
	//
	//resp, err = svc.GetLiveFlowDataByStreamGet(&getLiveFlowDataByStream)
	//if err != nil {
	//	fmt.Println("error:", err.Error())
	//}
	//if resp != nil {
	//	str, _ := json.Marshal(&resp)
	//	fmt.Printf("%+v\n", string(str))
	//}

	//	**************************************** 直播按流维度查询流量 POST（GetLiveFlowDataByStream）***************************************************
	// Parameters:
	//        StartTime   String  获取数据起始时间点，日期格式按ISO8601表示法，北京时间，格式为：YYYY-MM-DDThh:mm+0800，例如：2016-08-01T21:14+0800
	//        EndTime     String  结束时间需大于起始时间；获取日期格式按照ISO8601表示法，北京时间，格式为：YYYY-MM-DDThh:mm+0800，例如： 2016-08-01T21:14+0800
	//        StreamUrls   String  流名，支持批量查询，多个流名用逗号（半角）分隔
	//        Regions     String  计费区域名称，取值为CN:中国大陆，HK：香港，TW：台湾，AS：亚洲其他，NA：北美洲，SA：南美洲，EU：欧洲，AU：大洋洲，AF：非洲，支持多计费区域查询，多个区域用逗号（半角）分隔，缺省为 CN
	//        ResultType  String  取值为0：多域名多省份区域多运营商数据做合并；1：每个域名每个省份区域的每个运营商数据分别返回
	//        Granularity String  统计粒度，取值为 5（默认）：5分钟粒度；10：10分钟粒度；20：20分钟粒度；60：1小时粒度；240：4小时粒度；480：8小时粒度；1440：1天粒度
	//

	//getLiveFlowDataByStream := make(map[string]interface{})
	//getLiveFlowDataByStream["StartTime"] = "2021-05-12T01:00+0800"
	//getLiveFlowDataByStream["EndTime"] = "2021-05-12T01:10+0800"
	//getLiveFlowDataByStream["StreamUrls"] = "rtmp://realflv3.plu.cn/live/ce781fecb2f6447d82d03590e520872f"
	//getLiveFlowDataByStream["Regions"] = "CN"
	//getLiveFlowDataByStream["ResultType"] = 0
	//getLiveFlowDataByStream["Granularity"] = 5
	//
	//resp, err = svc.GetLiveFlowDataByStreamPost(&getLiveFlowDataByStream)
	//if err != nil {
	//	fmt.Println("error:", err.Error())
	//}
	//if resp != nil {
	//	str, _ := json.Marshal(&resp)
	//	fmt.Printf("%+v\n", string(str))
	//}

	//	**************************************** 直播按流维度查询带宽 GET（GetLiveBandwidthDataByStream）***************************************************
	// Parameters:
	//        StartTime   String  获取数据起始时间点，日期格式按ISO8601表示法，北京时间，格式为：YYYY-MM-DDThh:mm+0800，例如：2016-08-01T21:14+0800
	//        EndTime     String  结束时间需大于起始时间；获取日期格式按照ISO8601表示法，北京时间，格式为：YYYY-MM-DDThh:mm+0800，例如： 2016-08-01T21:14+0800
	//        StreamUrls   String  流名，支持批量查询，多个流名用逗号（半角）分隔
	//        Regions     String  计费区域名称，取值为CN:中国大陆，HK：香港，TW：台湾，AS：亚洲其他，NA：北美洲，SA：南美洲，EU：欧洲，AU：大洋洲，AF：非洲，支持多计费区域查询，多个区域用逗号（半角）分隔，缺省为 CN
	//        ResultType  String  取值为0：多域名多省份区域多运营商数据做合并；1：每个域名每个省份区域的每个运营商数据分别返回
	//        Granularity String  统计粒度，取值为 5（默认）：5分钟粒度；10：10分钟粒度；20：20分钟粒度；60：1小时粒度；240：4小时粒度；480：8小时粒度；1440：1天粒度
	//

	//getLiveBandwidthDataByStream := make(map[string]interface{})
	//getLiveBandwidthDataByStream["StartTime"] = "2021-05-12T01:00+0800"
	//getLiveBandwidthDataByStream["EndTime"] = "2021-05-12T01:10+0800"
	//getLiveBandwidthDataByStream["StreamUrls"] = "rtmp://realflv3.plu.cn/live/ce781fecb2f6447d82d03590e520872f"
	//getLiveBandwidthDataByStream["Regions"] = "CN"
	//getLiveBandwidthDataByStream["ResultType"] = 0
	//getLiveBandwidthDataByStream["Granularity"] = 5
	//
	//resp, err = svc.GetLiveBandwidthDataByStreamGet(&getLiveBandwidthDataByStream)
	//if err != nil {
	//	fmt.Println("error:", err.Error())
	//}
	//if resp != nil {
	//	str, _ := json.Marshal(&resp)
	//	fmt.Printf("%+v\n", string(str))
	//}

	//	**************************************** 直播按流维度查询带宽 POST（GetLiveBandwidthDataByStream）***************************************************
	// Parameters:
	//        StartTime   String  获取数据起始时间点，日期格式按ISO8601表示法，北京时间，格式为：YYYY-MM-DDThh:mm+0800，例如：2016-08-01T21:14+0800
	//        EndTime     String  结束时间需大于起始时间；获取日期格式按照ISO8601表示法，北京时间，格式为：YYYY-MM-DDThh:mm+0800，例如： 2016-08-01T21:14+0800
	//        StreamUrls   String  流名，支持批量查询，多个流名用逗号（半角）分隔
	//        Regions     String  计费区域名称，取值为CN:中国大陆，HK：香港，TW：台湾，AS：亚洲其他，NA：北美洲，SA：南美洲，EU：欧洲，AU：大洋洲，AF：非洲，支持多计费区域查询，多个区域用逗号（半角）分隔，缺省为 CN
	//        ResultType  String  取值为0：多域名多省份区域多运营商数据做合并；1：每个域名每个省份区域的每个运营商数据分别返回
	//        Granularity String  统计粒度，取值为 5（默认）：5分钟粒度；10：10分钟粒度；20：20分钟粒度；60：1小时粒度；240：4小时粒度；480：8小时粒度；1440：1天粒度
	//

	//getLiveBandwidthDataByStream := make(map[string]interface{})
	//getLiveBandwidthDataByStream["StartTime"] = "2021-05-12T01:00+0800"
	//getLiveBandwidthDataByStream["EndTime"] = "2021-05-12T01:10+0800"
	//getLiveBandwidthDataByStream["StreamUrls"] = "rtmp://realflv3.plu.cn/live/ce781fecb2f6447d82d03590e520872f"
	//getLiveBandwidthDataByStream["Regions"] = "CN"
	//getLiveBandwidthDataByStream["ResultType"] = 0
	//getLiveBandwidthDataByStream["Granularity"] = 5
	//
	//resp, err = svc.GetLiveBandwidthDataByStreamPost(&getLiveBandwidthDataByStream)
	//if err != nil {
	//	fmt.Println("error:", err.Error())
	//}
	//if resp != nil {
	//	str, _ := json.Marshal(&resp)
	//	fmt.Printf("%+v\n", string(str))
	//}

	//	**************************************** 直播按域名维度统计在线人数 GET（GetLiveOnlineUserDataByDomain）***************************************************
	// Parameters:
	//        StartTime   String  获取数据起始时间点，日期格式按ISO8601表示法，北京时间，格式为：YYYY-MM-DDThh:mm+0800，例如：2016-08-01T21:14+0800
	//        EndTime     String  结束时间需大于起始时间；获取日期格式按照ISO8601表示法，北京时间，格式为：YYYY-MM-DDThh:mm+0800，例如： 2016-08-01T21:14+0800
	//        DomainIds   String  域名ID，缺省为当前产品类型下的全部域名，可输入需要查询的域名ID，支持批量域名查询，多个域名ID用逗号（半角）分隔
	//        Regions     String  计费区域名称，取值为CN:中国大陆，HK：香港，TW：台湾，AS：亚洲其他，NA：北美洲，SA：南美洲，EU：欧洲，AU：大洋洲，AF：非洲，支持多计费区域查询，多个区域用逗号（半角）分隔，缺省为 CN
	//        ResultType  String  取值为0：多域名多省份区域多运营商数据做合并；1：每个域名每个省份区域的每个运营商数据分别返回
	//        Granularity String  统计粒度，取值为 5（默认）：5分钟粒度；10：10分钟粒度；20：20分钟粒度；60：1小时粒度；240：4小时粒度；480：8小时粒度；1440：1天粒度
	//

	//getLiveOnlineUserDataByDomain := make(map[string]interface{})
	//getLiveOnlineUserDataByDomain["StartTime"] = "2021-05-12T01:00+0800"
	//getLiveOnlineUserDataByDomain["EndTime"] = "2021-05-12T01:10+0800"
	//getLiveOnlineUserDataByDomain["DomainIds"] = "2D083FU"
	//getLiveOnlineUserDataByDomain["Regions"] = "CN"
	//getLiveOnlineUserDataByDomain["ResultType"] = 0
	//getLiveOnlineUserDataByDomain["Granularity"] = 5
	//
	//resp, err = svc.GetLiveOnlineUserDataByDomainGet(&getLiveOnlineUserDataByDomain)
	//if err != nil {
	//	fmt.Println("error:", err.Error())
	//}
	//if resp != nil {
	//	str, _ := json.Marshal(&resp)
	//	fmt.Printf("%+v\n", string(str))
	//}

	//	**************************************** 直播按域名维度统计在线人数 POST（GetLiveOnlineUserDataByDomain）***************************************************
	// Parameters:
	//        StartTime   String  获取数据起始时间点，日期格式按ISO8601表示法，北京时间，格式为：YYYY-MM-DDThh:mm+0800，例如：2016-08-01T21:14+0800
	//        EndTime     String  结束时间需大于起始时间；获取日期格式按照ISO8601表示法，北京时间，格式为：YYYY-MM-DDThh:mm+0800，例如： 2016-08-01T21:14+0800
	//        DomainIds   String  域名ID，缺省为当前产品类型下的全部域名，可输入需要查询的域名ID，支持批量域名查询，多个域名ID用逗号（半角）分隔
	//        Regions     String  计费区域名称，取值为CN:中国大陆，HK：香港，TW：台湾，AS：亚洲其他，NA：北美洲，SA：南美洲，EU：欧洲，AU：大洋洲，AF：非洲，支持多计费区域查询，多个区域用逗号（半角）分隔，缺省为 CN
	//        ResultType  String  取值为0：多域名多省份区域多运营商数据做合并；1：每个域名每个省份区域的每个运营商数据分别返回
	//        Granularity String  统计粒度，取值为 5（默认）：5分钟粒度；10：10分钟粒度；20：20分钟粒度；60：1小时粒度；240：4小时粒度；480：8小时粒度；1440：1天粒度
	//

	//getLiveOnlineUserDataByDomain := make(map[string]interface{})
	//getLiveOnlineUserDataByDomain["StartTime"] = "2021-05-12T01:00+0800"
	//getLiveOnlineUserDataByDomain["EndTime"] = "2021-05-12T01:10+0800"
	//getLiveOnlineUserDataByDomain["DomainIds"] = "2D083FU"
	//getLiveOnlineUserDataByDomain["Regions"] = "CN"
	//getLiveOnlineUserDataByDomain["ResultType"] = 0
	//getLiveOnlineUserDataByDomain["Granularity"] = 5
	//
	//resp, err = svc.GetLiveOnlineUserDataByDomainPost(&getLiveOnlineUserDataByDomain)
	//if err != nil {
	//	fmt.Println("error:", err.Error())
	//}
	//if resp != nil {
	//	str, _ := json.Marshal(&resp)
	//	fmt.Printf("%+v\n", string(str))
	//}

	//	**************************************** 直播按流维度统计在线人数 GET（GetLiveOnlineUserDataByStream）***************************************************
	// Parameters:
	//        StartTime   String  获取数据起始时间点，日期格式按ISO8601表示法，北京时间，格式为：YYYY-MM-DDThh:mm+0800，例如：2016-08-01T21:14+0800
	//        EndTime     String  结束时间需大于起始时间；获取日期格式按照ISO8601表示法，北京时间，格式为：YYYY-MM-DDThh:mm+0800，例如： 2016-08-01T21:14+0800
	//        StreamUrls   String  流名，支持批量查询，多个流名用逗号（半角）分隔
	//        Regions     String  计费区域名称，取值为CN:中国大陆，HK：香港，TW：台湾，AS：亚洲其他，NA：北美洲，SA：南美洲，EU：欧洲，AU：大洋洲，AF：非洲，支持多计费区域查询，多个区域用逗号（半角）分隔，缺省为 CN
	//        ResultType  String  取值为0：多域名多省份区域多运营商数据做合并；1：每个域名每个省份区域的每个运营商数据分别返回
	//        Granularity String  统计粒度，取值为 5（默认）：5分钟粒度；10：10分钟粒度；20：20分钟粒度；60：1小时粒度；240：4小时粒度；480：8小时粒度；1440：1天粒度
	//

	//getLiveOnlineUserDataByStream := make(map[string]interface{})
	//getLiveOnlineUserDataByStream["StartTime"] = "2021-05-12T01:00+0800"
	//getLiveOnlineUserDataByStream["EndTime"] = "2021-05-12T01:10+0800"
	//getLiveOnlineUserDataByStream["StreamUrls"] = "rtmp://realflv3.plu.cn/live/ce781fecb2f6447d82d03590e520872f"
	//getLiveOnlineUserDataByStream["Regions"] = "CN"
	//getLiveOnlineUserDataByStream["ResultType"] = 0
	//getLiveOnlineUserDataByStream["Granularity"] = 5
	//
	//resp, err = svc.GetLiveOnlineUserDataByStreamGet(&getLiveOnlineUserDataByStream)
	//if err != nil {
	//	fmt.Println("error:", err.Error())
	//}
	//if resp != nil {
	//	str, _ := json.Marshal(&resp)
	//	fmt.Printf("%+v\n", string(str))
	//}

	//	**************************************** 直播按流维度统计在线人数 POST（GetLiveOnlineUserDataByStream）***************************************************
	// Parameters:
	//        StartTime   String  获取数据起始时间点，日期格式按ISO8601表示法，北京时间，格式为：YYYY-MM-DDThh:mm+0800，例如：2016-08-01T21:14+0800
	//        EndTime     String  结束时间需大于起始时间；获取日期格式按照ISO8601表示法，北京时间，格式为：YYYY-MM-DDThh:mm+0800，例如： 2016-08-01T21:14+0800
	//        StreamUrls   String  流名，支持批量查询，多个流名用逗号（半角）分隔
	//        Regions     String  计费区域名称，取值为CN:中国大陆，HK：香港，TW：台湾，AS：亚洲其他，NA：北美洲，SA：南美洲，EU：欧洲，AU：大洋洲，AF：非洲，支持多计费区域查询，多个区域用逗号（半角）分隔，缺省为 CN
	//        ResultType  String  取值为0：多域名多省份区域多运营商数据做合并；1：每个域名每个省份区域的每个运营商数据分别返回
	//        Granularity String  统计粒度，取值为 5（默认）：5分钟粒度；10：10分钟粒度；20：20分钟粒度；60：1小时粒度；240：4小时粒度；480：8小时粒度；1440：1天粒度
	//

	//getLiveOnlineUserDataByStream := make(map[string]interface{})
	//getLiveOnlineUserDataByStream["StartTime"] = "2021-05-12T01:00+0800"
	//getLiveOnlineUserDataByStream["EndTime"] = "2021-05-12T01:10+0800"
	//getLiveOnlineUserDataByStream["StreamUrls"] = "rtmp://realflv3.plu.cn/live/ce781fecb2f6447d82d03590e520872f"
	//getLiveOnlineUserDataByStream["Regions"] = "CN"
	//getLiveOnlineUserDataByStream["ResultType"] = 0
	//getLiveOnlineUserDataByStream["Granularity"] = 5
	//
	//resp, err = svc.GetLiveOnlineUserDataByStreamPost(&getLiveOnlineUserDataByStream)
	//if err != nil {
	//	fmt.Println("error:", err.Error())
	//}
	//if resp != nil {
	//	str, _ := json.Marshal(&resp)
	//	fmt.Printf("%+v\n", string(str))
	//}

	//	**************************************** 获取按流维度的直播在线人数排行， 单位：每分钟的在线人数 GET（GetLiveTopOnlineUserData）***************************************************
	// Parameters:
	//        StartTime   String  获取数据起始时间点，日期格式按ISO8601表示法，北京时间，格式为：YYYY-MM-DDThh:mm+0800，例如：2016-08-01T21:14+0800
	//        DomainIds   String  域名ID，缺省为当前产品类型下的全部域名，可输入需要查询的域名ID，支持批量域名查询，多个域名ID用逗号（半角）分隔
	//        Regions     String  计费区域名称，取值为CN:中国大陆，HK：香港，TW：台湾，AS：亚洲其他，NA：北美洲，SA：南美洲，EU：欧洲，AU：大洋洲，AF：非洲，支持多计费区域查询，多个区域用逗号（半角）分隔，缺省为 CN
	//        ResultType  String  取值为0：多域名多省份区域多运营商数据做合并；1：每个域名每个省份区域的每个运营商数据分别返回
	//        LimitN 否 Int Top条数，取值为1-200，最大200，默认100

	//getLiveTopOnlineUserData := make(map[string]interface{})
	//getLiveTopOnlineUserData["StartTime"] = "2021-05-12T01:00+0800"
	//getLiveTopOnlineUserData["EndTime"] = "2021-05-12T01:10+0800"
	//getLiveTopOnlineUserData["Regions"] = "CN"
	//getLiveTopOnlineUserData["ResultType"] = 0
	//getLiveTopOnlineUserData["LimitN"] = 5
	//
	//resp, err = svc.GetLiveTopOnlineUserDataGet(&getLiveTopOnlineUserData)
	//if err != nil {
	//	fmt.Println("error:", err.Error())
	//}
	//if resp != nil {
	//	str, _ := json.Marshal(&resp)
	//	fmt.Printf("%+v\n", string(str))
	//}

	//	**************************************** 获取按流维度的直播在线人数排行， 单位：每分钟的在线人数 POST（GetLiveTopOnlineUserData）***************************************************
	// Parameters:
	//        StartTime   String  获取数据起始时间点，日期格式按ISO8601表示法，北京时间，格式为：YYYY-MM-DDThh:mm+0800，例如：2016-08-01T21:14+0800
	//        DomainIds   String  域名ID，缺省为当前产品类型下的全部域名，可输入需要查询的域名ID，支持批量域名查询，多个域名ID用逗号（半角）分隔
	//        Regions     String  计费区域名称，取值为CN:中国大陆，HK：香港，TW：台湾，AS：亚洲其他，NA：北美洲，SA：南美洲，EU：欧洲，AU：大洋洲，AF：非洲，支持多计费区域查询，多个区域用逗号（半角）分隔，缺省为 CN
	//        ResultType  String  取值为0：多域名多省份区域多运营商数据做合并；1：每个域名每个省份区域的每个运营商数据分别返回
	//        LimitN 否 Int Top条数，取值为1-200，最大200，默认100

	//getLiveTopOnlineUserData := make(map[string]interface{})
	//getLiveTopOnlineUserData["StartTime"] = "2021-05-12T01:00+0800"
	//getLiveTopOnlineUserData["EndTime"] = "2021-05-12T01:10+0800"
	//getLiveTopOnlineUserData["Regions"] = "CN"
	//getLiveTopOnlineUserData["ResultType"] = 0
	//getLiveTopOnlineUserData["LimitN"] = 5
	//
	//resp, err = svc.GetLiveTopOnlineUserDataPost(&getLiveTopOnlineUserData)
	//if err != nil {
	//	fmt.Println("error:", err.Error())
	//}
	//if resp != nil {
	//	str, _ := json.Marshal(&resp)
	//	fmt.Printf("%+v\n", string(str))
	//}

	//	**************************************** 日志下载接口 GET（GetDomainLogs）***************************************************
	// Parameters:
	//        PageSize        long    分页大小，默认20，最大500，取值1～500间整数
	//        PageNumber      long    取第几页。默认为1，取值1～10000
	//        DomainId      string  按域名过滤，默认为空，代表当前用户下所有域名
	//        StartTime    string  查询开始时间，格式yyyy-MM-dd，开始时间和结束时间均不指定时，默认是当天
	//		  EndTime    string  查询结束时间，格式yyyy-MM-dd，开始时间和结束时间均不指定时，默认是当天

	//getDomainLogs := make(map[string]interface{})
	//getDomainLogs["StartTime"] = "2021-05-12"
	//getDomainLogs["EndTime"] = "2021-05-13"
	//getDomainLogs["PageSize"] = 20
	//getDomainLogs["PageNumber"] = 1
	//getDomainLogs["DomainId"] = "2D083FU"
	//
	//resp, err = svc.GetDomainLogsGet(&getDomainLogs)
	//if err != nil {
	//	fmt.Println("error:", err.Error())
	//}
	//if resp != nil {
	//	str, _ := json.Marshal(&resp)
	//	fmt.Printf("%+v\n", string(str))
	//}

	//	**************************************** 日志下载接口 POST（GetDomainLogs）***************************************************
	// Parameters:
	//        PageSize        long    分页大小，默认20，最大500，取值1～500间整数
	//        PageNumber      long    取第几页。默认为1，取值1～10000
	//        DomainId      string  按域名过滤，默认为空，代表当前用户下所有域名
	//        StartTime    string  查询开始时间，格式yyyy-MM-dd，开始时间和结束时间均不指定时，默认是当天
	//		  EndTime    string  查询结束时间，格式yyyy-MM-dd，开始时间和结束时间均不指定时，默认是当天

	//getDomainLogs := make(map[string]interface{})
	//getDomainLogs["StartTime"] = "2021-05-12"
	//getDomainLogs["EndTime"] = "2021-05-13"
	//getDomainLogs["PageSize"] = 20
	//getDomainLogs["PageNumber"] = 1
	//getDomainLogs["DomainId"] = "2D083FU"
	//
	//resp, err = svc.GetDomainLogsPost(&getDomainLogs)
	//if err != nil {
	//	fmt.Println("error:", err.Error())
	//}
	//if resp != nil {
	//	str, _ := json.Marshal(&resp)
	//	fmt.Printf("%+v\n", string(str))
	//}

	//	**************************************** 刷新接口 POST（RefreshCaches）***************************************************
	// Parameters:
	//        Files        Url[]    需要文件类型刷新的Url列表
	//        Dirs         Url[]    需要文件类型刷新的Url列表
	//        其中url[]为：
	//		Url String 需要提交刷新的Url，单条

	//url := make(map[string]interface{})
	//url["Url"] = "http://ntj14123.test.com/abc.txt"
	//
	//files := make([]map[string]interface{}, 0)
	//files = append(files, url)
	//
	//refreshCaches := make(map[string]interface{})
	//refreshCaches["Files"] = files
	//
	//
	//resp, err = svc.RefreshCachesPost(&refreshCaches)
	//if err != nil {
	//	fmt.Println("error:", err.Error())
	//}
	//if resp != nil {
	//	str, _ := json.Marshal(&resp)
	//	fmt.Printf("%+v\n", string(str))
	//}

	//	**************************************** 预热接口 POST（PreloadCaches）***************************************************
	// Parameters:
	//        Urls         Url[]    需要文件类型预热的Url列表
	//        其中url[]为：
	//		Url String 需要提交预热的Url，单条

	//url := make(map[string]interface{})
	//url["Url"] = "http://ntj14123.test.com/abc.txt"
	//
	//urls := make([]map[string]interface{}, 0)
	//urls = append(urls, url)
	//
	//refreshCaches := make(map[string]interface{})
	//refreshCaches["Urls"] = urls
	//
	//
	//resp, err = svc.PreloadCachesPost(&refreshCaches)
	//if err != nil {
	//	fmt.Println("error:", err.Error())
	//}
	//if resp != nil {
	//	str, _ := json.Marshal(&resp)
	//	fmt.Printf("%+v\n", string(str))
	//}

	//	**************************************** 预热进度查询 POST（GetRefreshOrPreloadTask）***************************************************
	// Parameters:
	//		PageSize        long    分页大小，默认20，最大500，取值1～500间整数
	//      PageNumber      long    取第几页。默认为1，取值1～10000
	//      StartTime    string  查询开始时间，格式YYYY-MM-DDThh:mm+0800，开始时间和结束时间均不指定时，默认是当天
	//		EndTime    string  查询结束时间，格式YYYY-MM-DDThh:mm+0800，开始时间和结束时间均不指定时，默认是当天
	//		TaskId    string   支持按任务ID查询，只允许输入单个任务ID
	//      Urls         Url[]    需要文件类型预热的Url列表
	//      其中url[]为：
	//			Url String 需要提交预热的Url，单条

	//url := make(map[string]interface{})
	//url["Url"] = "http://ntj14123.test.com/abc.txt"
	//
	//urls := make([]map[string]interface{}, 0)
	//urls = append(urls, url)
	//
	//getRefreshOrPreloadTask := make(map[string]interface{})
	//getRefreshOrPreloadTask["Urls"] = urls
	//getRefreshOrPreloadTask["PageSize"] = 20
	//getRefreshOrPreloadTask["PageNumber"] = 2
	//getRefreshOrPreloadTask["StartTime"] = "2021-05-20T00:00+0800"
	//getRefreshOrPreloadTask["EndTime"] = "2021-05-21T00:00+0800"
	//getRefreshOrPreloadTask["TaskId"] = "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
	//
	//resp, err = svc.GetRefreshOrPreloadTaskPost(&getRefreshOrPreloadTask)
	//if err != nil {
	//	fmt.Println("error:", err.Error())
	//}
	//if resp != nil {
	//	str, _ := json.Marshal(&resp)
	//	fmt.Printf("%+v\n", string(str))
	//}

	//	**************************************** 获取刷新、预热URL及目录的最大限制数量 ，及当日剩余刷新、预热URL及目录的条数 GET（GetRefreshOrPreloadQuota）***************************************************

	//resp, err = svc.GetRefreshOrPreloadQuotaGet(nil)
	//if err != nil {
	//	fmt.Println("error:", err.Error())
	//}
	//if resp != nil {
	//	str, _ := json.Marshal(&resp)
	//	fmt.Printf("%+v\n", string(str))
	//}

	//	**************************************** 获取刷新、预热URL及目录的最大限制数量 ，及当日剩余刷新、预热URL及目录的条数 POST（GetRefreshOrPreloadQuota）***************************************************

	//resp, err = svc.GetRefreshOrPreloadQuotaPost(nil)
	//if err != nil {
	//	fmt.Println("error:", err.Error())
	//}
	//if resp != nil {
	//	str, _ := json.Marshal(&resp)
	//	fmt.Printf("%+v\n", string(str))
	//}

	//	**************************************** 设置日志服务接口 GET（SetDomainLogService）***************************************************
	//Parameters:
	//        ActionType      string  操作类型，取值为start：启用；stop：停用
	//        DomainIds    string  需要启用或停用日志服务的域名ID，支持批量域名开启或停用，多个域名ID用逗号（半角）分隔
	//		Granularity    Long  日志存储粒度，取值为60：按小时粒度存储；1440：按天粒度存储，当前暂不支持按小时粒度存储；开启时为必填，关闭时可不填

	//setDomainLogService := make(map[string]interface{})
	//setDomainLogService["ActionType"] = "stop"
	//setDomainLogService["DomainIds"] = "2D083FU"
	//setDomainLogService["Granularity"] = 1440
	//
	//resp, err = svc.SetDomainLogServiceGet(&setDomainLogService)
	//if err != nil {
	//	fmt.Println("error:", err.Error())
	//}
	//if resp != nil {
	//	str, _ := json.Marshal(&resp)
	//	fmt.Printf("%+v\n", string(str))
	//}

	//	**************************************** 设置日志服务接口 POST（SetDomainLogService）***************************************************
	//Parameters:
	//        ActionType      string  操作类型，取值为start：启用；stop：停用
	//        DomainIds    string  需要启用或停用日志服务的域名ID，支持批量域名开启或停用，多个域名ID用逗号（半角）分隔
	//		Granularity    Long  日志存储粒度，取值为60：按小时粒度存储；1440：按天粒度存储，当前暂不支持按小时粒度存储；开启时为必填，关闭时可不填

	//setDomainLogService := make(map[string]interface{})
	//setDomainLogService["ActionType"] = "start"
	//setDomainLogService["DomainIds"] = "2D083FU"
	//setDomainLogService["Granularity"] = 1440
	//
	//resp, err = svc.SetDomainLogServicePost(&setDomainLogService)
	//if err != nil {
	//	fmt.Println("error:", err.Error())
	//}
	//if resp != nil {
	//	str, _ := json.Marshal(&resp)
	//	fmt.Printf("%+v\n", string(str))
	//}

	//	**************************************** 获取域名日志服务状态 GET（GetDomainLogServiceStatus）***************************************************
	// Parameters:
	//        DomainIds    string  需要查询日志服务的域名ID，支持批量域名查询，多个域名ID用逗号（半角）分隔

	//getDomainLogServiceStatus := make(map[string]interface{})
	//getDomainLogServiceStatus["DomainIds"] = "2D083FU"

	//resp, err = svc.GetDomainLogServiceStatusGet(&getDomainLogServiceStatus)
	//if err != nil {
	//	fmt.Println("error:", err.Error())
	//}
	//if resp != nil {
	//	str, _ := json.Marshal(&resp)
	//	fmt.Printf("%+v\n", string(str))
	//}

	//	**************************************** 获取域名日志服务状态 POST（GetDomainLogServiceStatus）***************************************************
	// Parameters:
	//        DomainIds    string  需要查询日志服务的域名ID，支持批量域名查询，多个域名ID用逗号（半角）分隔

	//getDomainLogServiceStatus := make(map[string]interface{})
	//getDomainLogServiceStatus["DomainIds"] = "2D083FU"
	//
	//resp, err = svc.GetDomainLogServiceStatusPost(&getDomainLogServiceStatus)
	//if err != nil {
	//	fmt.Println("error:", err.Error())
	//}
	//if resp != nil {
	//	str, _ := json.Marshal(&resp)
	//	fmt.Printf("%+v\n", string(str))
	//}

	//	**************************************** 获取域名独立请求的IP个数 GET（GetUvData）***************************************************
	// Parameters:
	//        DomainIds       String  域名ID，缺省为当前产品类型下的全部域名，可输入需要查询的域名ID，支持批量域名查询，多个域名ID用逗号（半角）分隔
	//        CdnType     String  产品类型，只允许输入一种类型，取值为video:音视频点播,file:大文件下载,live:流媒体直播,当前不支持直播类型
	//        StartTime       String  获取数据起始时间点，日期格式按ISO8601表示法，北京时间，格式为：YYYY-MM-DDThh:mm+0800，例如：2016-08-01T21:14+0800
	//        EndTime         String  结束时间需大于起始时间；获取日期格式按照ISO8601表示法，北京时间，格式为：YYYY-MM-DDThh:mm+0800，例如： 2016-08-01T21:14+0800
	//        ResultType      Long    取值为0：多域名多区域数据做合并；1：每个域名每个区域的数据分别返回
	//        Granularity     Long     统计粒度，取值为 5（默认）：5分钟粒度；

	//getUvData := make(map[string]interface{})
	//getUvData["DomainIds"] = "2D083FU"
	//getUvData["CdnType"] = "page"
	//getUvData["StartTime"] = "2021-05-31T01:14+0800"
	//getUvData["EndTime"] = "2021-05-31T01:24+0800"
	//getUvData["ResultType"] = 0
	//getUvData["Granularity"] = 5
	//
	//resp, err = svc.GetUvDataGet(&getUvData)
	//if err != nil {
	//	fmt.Println("error:", err.Error())
	//}
	//if resp != nil {
	//	str, _ := json.Marshal(&resp)
	//	fmt.Printf("%+v\n", string(str))
	//}

	//	**************************************** 获取域名独立请求的IP个数 POST（GetUvData）***************************************************
	// Parameters:
	//        DomainIds       String  域名ID，缺省为当前产品类型下的全部域名，可输入需要查询的域名ID，支持批量域名查询，多个域名ID用逗号（半角）分隔
	//        CdnType     String  产品类型，只允许输入一种类型，取值为video:音视频点播,file:大文件下载,live:流媒体直播,当前不支持直播类型
	//        StartTime       String  获取数据起始时间点，日期格式按ISO8601表示法，北京时间，格式为：YYYY-MM-DDThh:mm+0800，例如：2016-08-01T21:14+0800
	//        EndTime         String  结束时间需大于起始时间；获取日期格式按照ISO8601表示法，北京时间，格式为：YYYY-MM-DDThh:mm+0800，例如： 2016-08-01T21:14+0800
	//        ResultType      Long    取值为0：多域名多区域数据做合并；1：每个域名每个区域的数据分别返回
	//        Granularity     Long     统计粒度，取值为 5（默认）：5分钟粒度；

	//getUvData := make(map[string]interface{})
	//getUvData["DomainIds"] = "2D083FU"
	//getUvData["CdnType"] = "page"
	//getUvData["StartTime"] = "2021-05-31T01:14+0800"
	//getUvData["EndTime"] = "2021-05-31T01:24+0800"
	//getUvData["ResultType"] = 0
	//getUvData["Granularity"] = 5
	//
	//resp, err = svc.GetUvDataPost(&getUvData)
	//if err != nil {
	//	fmt.Println("error:", err.Error())
	//}
	//if resp != nil {
	//	str, _ := json.Marshal(&resp)
	//	fmt.Printf("%+v\n", string(str))
	//}

	//	**************************************** 获取域名某天内某一时段的热门页面访问数据排名，仅包含Top200且访问数大于15次的热门页面的访问次数、访问流量，并按次数排名 GET（GetTopReferData）***************************************************
	// Parameters:
	//        DomainIds       String  域名ID，缺省为当前产品类型下的全部域名，可输入需要查询的域名ID，支持批量域名查询，多个域名ID用逗号（半角）分隔
	//        CdnType     String  产品类型，只允许输入一种类型，取值为video:音视频点播,file:大文件下载,live:流媒体直播,当前不支持直播类型
	//        StartTime       String  获取数据起始时间点，日期格式按ISO8601表示法，北京时间，格式为：YYYY-MM-DDThh:mm+0800，例如：2016-08-01T21:14+0800
	//        EndTime         String  结束时间需大于起始时间；获取日期格式按照ISO8601表示法，北京时间，格式为：YYYY-MM-DDThh:mm+0800，例如： 2016-08-01T21:14+0800
	//        LimitN     Long     热门Refer条数，取值为1-200，最大200，默认100

	//getTopReferData := make(map[string]interface{})
	//getTopReferData["DomainIds"] = "2D083FU"
	//getTopReferData["CdnType"] = "page"
	//getTopReferData["StartTime"] = "2021-05-31T01:14+0800"
	//getTopReferData["EndTime"] = "2021-05-31T01:24+0800"
	//getTopReferData["LimitN"] = 10
	//
	//resp, err = svc.GetTopReferDataGet(&getTopReferData)
	//if err != nil {
	//	fmt.Println("error:", err.Error())
	//}
	//if resp != nil {
	//	str, _ := json.Marshal(&resp)
	//	fmt.Printf("%+v\n", string(str))
	//}

	//	**************************************** 获取域名某天内某一时段的热门页面访问数据排名，仅包含Top200且访问数大于15次的热门页面的访问次数、访问流量，并按次数排名 POST（GetTopReferData）***************************************************
	// Parameters:
	//        DomainIds       String  域名ID，缺省为当前产品类型下的全部域名，可输入需要查询的域名ID，支持批量域名查询，多个域名ID用逗号（半角）分隔
	//        CdnType     String  产品类型，只允许输入一种类型，取值为video:音视频点播,file:大文件下载,live:流媒体直播,当前不支持直播类型
	//        StartTime       String  获取数据起始时间点，日期格式按ISO8601表示法，北京时间，格式为：YYYY-MM-DDThh:mm+0800，例如：2016-08-01T21:14+0800
	//        EndTime         String  结束时间需大于起始时间；获取日期格式按照ISO8601表示法，北京时间，格式为：YYYY-MM-DDThh:mm+0800，例如： 2016-08-01T21:14+0800
	//        LimitN     Long     热门Refer条数，取值为1-200，最大200，默认100

	//getTopReferData := make(map[string]interface{})
	//getTopReferData["DomainIds"] = "2D083FU"
	//getTopReferData["CdnType"] = "page"
	//getTopReferData["StartTime"] = "2021-05-31T01:14+0800"
	//getTopReferData["EndTime"] = "2021-05-31T01:24+0800"
	//getTopReferData["LimitN"] = 10
	//
	//resp, err = svc.GetTopReferDataPost(&getTopReferData)
	//if err != nil {
	//	fmt.Println("error:", err.Error())
	//}
	//if resp != nil {
	//	str, _ := json.Marshal(&resp)
	//	fmt.Printf("%+v\n", string(str))
	//}

	//	**************************************** 本接口用于获取域名某天内某一时段的TOP IP访问数据，仅包含Top200且访问次数大于15次的独立请求的IP的访问次数、访问流量，并按次数排序 GET（GetTopIpData）***************************************************
	// Parameters:
	//        DomainIds       String  域名ID，缺省为当前产品类型下的全部域名，可输入需要查询的域名ID，支持批量域名查询，多个域名ID用逗号（半角）分隔
	//        CdnType     String  产品类型，只允许输入一种类型，取值为video:音视频点播,file:大文件下载,live:流媒体直播,当前不支持直播类型
	//        StartTime       String  获取数据起始时间点，日期格式按ISO8601表示法，北京时间，格式为：YYYY-MM-DDThh:mm+0800，例如：2016-08-01T21:14+0800
	//        EndTime         String  结束时间需大于起始时间；获取日期格式按照ISO8601表示法，北京时间，格式为：YYYY-MM-DDThh:mm+0800，例如： 2016-08-01T21:14+0800
	//        LimitN     Long     热门Refer条数，取值为1-200，最大200，默认100

	//getTopIpData := make(map[string]interface{})
	//getTopIpData["DomainIds"] = "2D083FU"
	//getTopIpData["CdnType"] = "page"
	//getTopIpData["StartTime"] = "2021-05-31T01:14+0800"
	//getTopIpData["EndTime"] = "2021-05-31T01:24+0800"
	//getTopIpData["LimitN"] = 10
	//
	//resp, err = svc.GetTopIpDataGet(&getTopIpData)
	//if err != nil {
	//	fmt.Println("error:", err.Error())
	//}
	//if resp != nil {
	//	str, _ := json.Marshal(&resp)
	//	fmt.Printf("%+v\n", string(str))
	//}

	//	**************************************** 本接口用于获取域名某天内某一时段的TOP IP访问数据，仅包含Top200且访问次数大于15次的独立请求的IP的访问次数、访问流量，并按次数排序 POST（GetTopIpData）***************************************************
	// Parameters:
	//        DomainIds       String  域名ID，缺省为当前产品类型下的全部域名，可输入需要查询的域名ID，支持批量域名查询，多个域名ID用逗号（半角）分隔
	//        CdnType     String  产品类型，只允许输入一种类型，取值为video:音视频点播,file:大文件下载,live:流媒体直播,当前不支持直播类型
	//        StartTime       String  获取数据起始时间点，日期格式按ISO8601表示法，北京时间，格式为：YYYY-MM-DDThh:mm+0800，例如：2016-08-01T21:14+0800
	//        EndTime         String  结束时间需大于起始时间；获取日期格式按照ISO8601表示法，北京时间，格式为：YYYY-MM-DDThh:mm+0800，例如： 2016-08-01T21:14+0800
	//        LimitN     Long     热门Refer条数，取值为1-200，最大200，默认100

	//getTopIpData := make(map[string]interface{})
	//getTopIpData["DomainIds"] = "2D083FU"
	//getTopIpData["CdnType"] = "page"
	//getTopIpData["StartTime"] = "2021-05-31T01:14+0800"
	//getTopIpData["EndTime"] = "2021-05-31T01:24+0800"
	//getTopIpData["LimitN"] = 10
	//
	//resp, err = svc.GetTopIpDataPost(&getTopIpData)
	//if err != nil {
	//	fmt.Println("error:", err.Error())
	//}
	//if resp != nil {
	//	str, _ := json.Marshal(&resp)
	//	fmt.Printf("%+v\n", string(str))
	//}

	//	**************************************** 获取域名流量命中率、请求数命中率数据 GET（GetProvinceAndIspHitRateDetailedData）***************************************************
	// Parameters:
	//        DomainIds       String  域名ID，缺省为当前产品类型下的全部域名，可输入需要查询的域名ID，支持批量域名查询，多个域名ID用逗号（半角）分隔
	//        CdnType     String  产品类型，只允许输入一种类型，取值为video:音视频点播,file:大文件下载,live:流媒体直播
	//        StartTime       String  获取数据起始时间点，日期格式按ISO8601表示法，北京时间，格式为：YYYY-MM-DDThh:mm+0800，例如：2016-08-01T21:14+0800
	//        EndTime         String  结束时间需大于起始时间；获取日期格式按照ISO8601表示法，北京时间，格式为：YYYY-MM-DDThh:mm+0800，例如： 2016-08-01T21:14+0800
	//        Provinces     String     省份区域名称，取值详见枚举列表，支持多省份区域查询，多个省份区域用逗号（半角）分隔，缺省为全部省份区域
	//		Isps     String     运营商名称，取值详见枚举列表，支持多运营商查询，多个运营商用逗号（半角）分隔，缺省为全部运营商
	//		ResultType     Long     取值为0：多域名数据做合并；1：每个域名的数据分别返回
	//		Granularity     Long     热统计粒度，取值为 5（默认）：5分钟粒度；10：10分钟粒度；20：20分钟粒度；60：1小时粒度；240：4小时粒度；480：8小时粒度；1440：1天粒度；以上粒度均取该粒度时间段的流量之和、请求数之和
	//		HitType     String     数据类型， 取值为flowhitrate:流量命中率;reqhitrate:请求数命中率; 支持多类型选择，多个类型用逗号（半角）分隔，缺省为reqhitrate
	//

	//getProvinceAndIspHitRateDetailedData := make(map[string]interface{})
	//getProvinceAndIspHitRateDetailedData["DomainIds"] = "2D083FU"
	//getProvinceAndIspHitRateDetailedData["CdnType"] = "page"
	//getProvinceAndIspHitRateDetailedData["StartTime"] = "2021-05-31T01:14+0800"
	//getProvinceAndIspHitRateDetailedData["EndTime"] = "2021-05-31T01:24+0800"
	//getProvinceAndIspHitRateDetailedData["Provinces"] = "beijing"
	//getProvinceAndIspHitRateDetailedData["Isps"] = "CM"
	//getProvinceAndIspHitRateDetailedData["ResultType"] = 0
	//getProvinceAndIspHitRateDetailedData["Granularity"] = 5
	//getProvinceAndIspHitRateDetailedData["HitType"] = "reqhitrate"
	//
	//resp, err = svc.GetProvinceAndIspHitRateDetailedDataGet(&getProvinceAndIspHitRateDetailedData)
	//if err != nil {
	//	fmt.Println("error:", err.Error())
	//}
	//if resp != nil {
	//	str, _ := json.Marshal(&resp)
	//	fmt.Printf("%+v\n", string(str))
	//}

	//	**************************************** 获取域名流量命中率、请求数命中率数据 POST（GetProvinceAndIspHitRateDetailedData）***************************************************
	// Parameters:
	//        DomainIds       String  域名ID，缺省为当前产品类型下的全部域名，可输入需要查询的域名ID，支持批量域名查询，多个域名ID用逗号（半角）分隔
	//        CdnType     String  产品类型，只允许输入一种类型，取值为video:音视频点播,file:大文件下载,live:流媒体直播
	//        StartTime       String  获取数据起始时间点，日期格式按ISO8601表示法，北京时间，格式为：YYYY-MM-DDThh:mm+0800，例如：2016-08-01T21:14+0800
	//        EndTime         String  结束时间需大于起始时间；获取日期格式按照ISO8601表示法，北京时间，格式为：YYYY-MM-DDThh:mm+0800，例如： 2016-08-01T21:14+0800
	//        Provinces     String     省份区域名称，取值详见枚举列表，支持多省份区域查询，多个省份区域用逗号（半角）分隔，缺省为全部省份区域
	//		Isps     String     运营商名称，取值详见枚举列表，支持多运营商查询，多个运营商用逗号（半角）分隔，缺省为全部运营商
	//		ResultType     Long     取值为0：多域名数据做合并；1：每个域名的数据分别返回
	//		Granularity     Long     热统计粒度，取值为 5（默认）：5分钟粒度；10：10分钟粒度；20：20分钟粒度；60：1小时粒度；240：4小时粒度；480：8小时粒度；1440：1天粒度；以上粒度均取该粒度时间段的流量之和、请求数之和
	//		HitType     String     数据类型， 取值为flowhitrate:流量命中率;reqhitrate:请求数命中率; 支持多类型选择，多个类型用逗号（半角）分隔，缺省为reqhitrate
	//

	//getProvinceAndIspHitRateDetailedData := make(map[string]interface{})
	//getProvinceAndIspHitRateDetailedData["DomainIds"] = "2D083FU"
	//getProvinceAndIspHitRateDetailedData["CdnType"] = "page"
	//getProvinceAndIspHitRateDetailedData["StartTime"] = "2021-05-31T01:14+0800"
	//getProvinceAndIspHitRateDetailedData["EndTime"] = "2021-05-31T01:24+0800"
	//getProvinceAndIspHitRateDetailedData["Provinces"] = "beijing"
	//getProvinceAndIspHitRateDetailedData["Isps"] = "CM"
	//getProvinceAndIspHitRateDetailedData["ResultType"] = 0
	//getProvinceAndIspHitRateDetailedData["Granularity"] = 5
	//getProvinceAndIspHitRateDetailedData["HitType"] = "reqhitrate"
	//
	//resp, err = svc.GetProvinceAndIspHitRateDetailedDataPost(&getProvinceAndIspHitRateDetailedData)
	//if err != nil {
	//	fmt.Println("error:", err.Error())
	//}
	//if resp != nil {
	//	str, _ := json.Marshal(&resp)
	//	fmt.Printf("%+v\n", string(str))
	//}

	//	**************************************** 获取域名一段时间内在中国大陆地区各省份及各运营商的Http状态码访问次数及占比数据 GET（GetProvinceAndIspHttpCodeData）***************************************************
	// Parameters:
	//        DomainIds       String  域名ID，缺省为当前产品类型下的全部域名，可输入需要查询的域名ID，支持批量域名查询，多个域名ID用逗号（半角）分隔
	//        CdnType     String  产品类型，只允许输入一种类型，取值为video:音视频点播,file:大文件下载,live:流媒体直播,当前不支持直播类型
	//        StartTime       String  获取数据起始时间点，日期格式按ISO8601表示法，北京时间，格式为：YYYY-MM-DDThh:mm+0800，例如：2016-08-01T21:14+0800
	//        EndTime         String  结束时间需大于起始时间；获取日期格式按照ISO8601表示法，北京时间，格式为：YYYY-MM-DDThh:mm+0800，例如： 2016-08-01T21:14+0800
	//        Provinces     String     省份区域名称，取值详见枚举列表，支持多省份区域查询，多个省份区域用逗号（半角）分隔，缺省为全部省份区域
	//		Isps     String     运营商名称，取值详见枚举列表，支持多运营商查询，多个运营商用逗号（半角）分隔，缺省为全部运营商
	//

	//getProvinceAndIspHttpCodeData := make(map[string]interface{})
	//getProvinceAndIspHttpCodeData["DomainIds"] = "2D083FU"
	//getProvinceAndIspHttpCodeData["CdnType"] = "page"
	//getProvinceAndIspHttpCodeData["StartTime"] = "2021-05-31T01:14+0800"
	//getProvinceAndIspHttpCodeData["EndTime"] = "2021-05-31T01:24+0800"
	//getProvinceAndIspHttpCodeData["Provinces"] = "beijing"
	//getProvinceAndIspHttpCodeData["Isps"] = "CM"
	//
	//resp, err = svc.GetProvinceAndIspHttpCodeDataGet(&getProvinceAndIspHttpCodeData)
	//if err != nil {
	//	fmt.Println("error:", err.Error())
	//}
	//if resp != nil {
	//	str, _ := json.Marshal(&resp)
	//	fmt.Printf("%+v\n", string(str))
	//}

	//	**************************************** 获取域名一段时间内在中国大陆地区各省份及各运营商的Http状态码访问次数及占比数据 POST（GetProvinceAndIspHttpCodeData）***************************************************
	// Parameters:
	//        DomainIds       String  域名ID，缺省为当前产品类型下的全部域名，可输入需要查询的域名ID，支持批量域名查询，多个域名ID用逗号（半角）分隔
	//        CdnType     String  产品类型，只允许输入一种类型，取值为video:音视频点播,file:大文件下载,live:流媒体直播,当前不支持直播类型
	//        StartTime       String  获取数据起始时间点，日期格式按ISO8601表示法，北京时间，格式为：YYYY-MM-DDThh:mm+0800，例如：2016-08-01T21:14+0800
	//        EndTime         String  结束时间需大于起始时间；获取日期格式按照ISO8601表示法，北京时间，格式为：YYYY-MM-DDThh:mm+0800，例如： 2016-08-01T21:14+0800
	//        Provinces     String     省份区域名称，取值详见枚举列表，支持多省份区域查询，多个省份区域用逗号（半角）分隔，缺省为全部省份区域
	//		Isps     String     运营商名称，取值详见枚举列表，支持多运营商查询，多个运营商用逗号（半角）分隔，缺省为全部运营商
	//

	//getProvinceAndIspHttpCodeData := make(map[string]interface{})
	//getProvinceAndIspHttpCodeData["DomainIds"] = "2D083FU"
	//getProvinceAndIspHttpCodeData["CdnType"] = "page"
	//getProvinceAndIspHttpCodeData["StartTime"] = "2021-05-31T01:14+0800"
	//getProvinceAndIspHttpCodeData["EndTime"] = "2021-05-31T01:24+0800"
	//getProvinceAndIspHttpCodeData["Provinces"] = "beijing"
	//getProvinceAndIspHttpCodeData["Isps"] = "CM"
	//
	//resp, err = svc.GetProvinceAndIspHttpCodeDataPost(&getProvinceAndIspHttpCodeData)
	//if err != nil {
	//	fmt.Println("error:", err.Error())
	//}
	//if resp != nil {
	//	str, _ := json.Marshal(&resp)
	//	fmt.Printf("%+v\n", string(str))
	//}

	//	**************************************** 获取域名在中国大陆地区各省份及各运营商的Http状态码详细访问次数及占比数据 GET（GetProvinceAndIspHttpCodeDetailedData）***************************************************
	// Parameters:
	//        DomainIds       String  域名ID，缺省为当前产品类型下的全部域名，可输入需要查询的域名ID，支持批量域名查询，多个域名ID用逗号（半角）分隔
	//        CdnType     String  产品类型，只允许输入一种类型，取值为video:音视频点播,file:大文件下载,live:流媒体直播，当前不支持直播类型
	//        StartTime       String  获取数据起始时间点，日期格式按ISO8601表示法，北京时间，格式为：YYYY-MM-DDThh:mm+0800，例如：2016-08-01T21:14+0800
	//        EndTime         String  结束时间需大于起始时间；获取日期格式按照ISO8601表示法，北京时间，格式为：YYYY-MM-DDThh:mm+0800，例如： 2016-08-01T21:14+0800
	//        Provinces     String     省份区域名称，取值详见枚举列表，支持多省份区域查询，多个省份区域用逗号（半角）分隔，缺省为全部省份区域
	//		Isps     String     运营商名称，取值详见枚举列表，支持多运营商查询，多个运营商用逗号（半角）分隔，缺省为全部运营商
	//		Granularity     Long     统计粒度，取值为 5（默认）：5分钟粒度；10：10分钟粒度；20：20分钟粒度；60：1小时粒度；240：4小时粒度；480：8小时粒度；1440：1天粒度
	//		ResultType     Long     取值为0：多域名数据做合并；1：每个域名的数据分别返回

	//getProvinceAndIspHttpCodeDetailedData := make(map[string]interface{})
	//getProvinceAndIspHttpCodeDetailedData["DomainIds"] = "2D083FU"
	//getProvinceAndIspHttpCodeDetailedData["CdnType"] = "page"
	//getProvinceAndIspHttpCodeDetailedData["StartTime"] = "2021-05-31T01:14+0800"
	//getProvinceAndIspHttpCodeDetailedData["EndTime"] = "2021-05-31T01:24+0800"
	//getProvinceAndIspHttpCodeDetailedData["Provinces"] = "beijing"
	//getProvinceAndIspHttpCodeDetailedData["Isps"] = "CM"
	//getProvinceAndIspHttpCodeDetailedData["Granularity"] = 5
	//getProvinceAndIspHttpCodeDetailedData["ResultType"] = 1
	//
	//resp, err = svc.GetProvinceAndIspHttpCodeDetailedDataGet(&getProvinceAndIspHttpCodeDetailedData)
	//if err != nil {
	//	fmt.Println("error:", err.Error())
	//}
	//if resp != nil {
	//	str, _ := json.Marshal(&resp)
	//	fmt.Printf("%+v\n", string(str))
	//}

	//	**************************************** 获取域名在中国大陆地区各省份及各运营商的Http状态码详细访问次数及占比数据 POST（GetProvinceAndIspHttpCodeDetailedData）***************************************************
	// Parameters:
	//        DomainIds       String  域名ID，缺省为当前产品类型下的全部域名，可输入需要查询的域名ID，支持批量域名查询，多个域名ID用逗号（半角）分隔
	//        CdnType     String  产品类型，只允许输入一种类型，取值为video:音视频点播,file:大文件下载,live:流媒体直播，当前不支持直播类型
	//        StartTime       String  获取数据起始时间点，日期格式按ISO8601表示法，北京时间，格式为：YYYY-MM-DDThh:mm+0800，例如：2016-08-01T21:14+0800
	//        EndTime         String  结束时间需大于起始时间；获取日期格式按照ISO8601表示法，北京时间，格式为：YYYY-MM-DDThh:mm+0800，例如： 2016-08-01T21:14+0800
	//        Provinces     String     省份区域名称，取值详见枚举列表，支持多省份区域查询，多个省份区域用逗号（半角）分隔，缺省为全部省份区域
	//		Isps     String     运营商名称，取值详见枚举列表，支持多运营商查询，多个运营商用逗号（半角）分隔，缺省为全部运营商
	//		Granularity     Long     统计粒度，取值为 5（默认）：5分钟粒度；10：10分钟粒度；20：20分钟粒度；60：1小时粒度；240：4小时粒度；480：8小时粒度；1440：1天粒度
	//		ResultType     Long     取值为0：多域名数据做合并；1：每个域名的数据分别返回

	//getProvinceAndIspHttpCodeDetailedData := make(map[string]interface{})
	//getProvinceAndIspHttpCodeDetailedData["DomainIds"] = "2D083FU"
	//getProvinceAndIspHttpCodeDetailedData["CdnType"] = "page"
	//getProvinceAndIspHttpCodeDetailedData["StartTime"] = "2021-05-31T01:14+0800"
	//getProvinceAndIspHttpCodeDetailedData["EndTime"] = "2021-05-31T01:24+0800"
	//getProvinceAndIspHttpCodeDetailedData["Provinces"] = "beijing"
	//getProvinceAndIspHttpCodeDetailedData["Isps"] = "CM"
	//getProvinceAndIspHttpCodeDetailedData["Granularity"] = 5
	//getProvinceAndIspHttpCodeDetailedData["ResultType"] = 1
	//
	//
	//resp, err = svc.GetProvinceAndIspHttpCodeDetailedDataPost(&getProvinceAndIspHttpCodeDetailedData)
	//if err != nil {
	//	fmt.Println("error:", err.Error())
	//}
	//if resp != nil {
	//	str, _ := json.Marshal(&resp)
	//	fmt.Printf("%+v\n", string(str))
	//}

	//	**************************************** 获取域名在中国大陆地区各省份及各运营商的请求数数据，包括边缘请求数 GET（GetProvinceAndIspPvData）***************************************************
	// Parameters:
	//        DomainIds       String  域名ID，缺省为当前产品类型下的全部域名，可输入需要查询的域名ID，支持批量域名查询，多个域名ID用逗号（半角）分隔
	//        CdnType     String  产品类型，只允许输入一种类型，取值为video:音视频点播,file:大文件下载,live:流媒体直播
	//        StartTime       String  获取数据起始时间点，日期格式按ISO8601表示法，北京时间，格式为：YYYY-MM-DDThh:mm+0800，例如：2016-08-01T21:14+0800
	//        EndTime         String  结束时间需大于起始时间；获取日期格式按照ISO8601表示法，北京时间，格式为：YYYY-MM-DDThh:mm+0800，例如： 2016-08-01T21:14+0800
	//        Provinces     String     省份区域名称，取值详见枚举列表，支持多省份区域查询，多个省份区域用逗号（半角）分隔，缺省为全部省份区域
	//		Isps     String     运营商名称，取值详见枚举列表，支持多运营商查询，多个运营商用逗号（半角）分隔，缺省为全部运营商
	//		Granularity     Long     统计粒度，取值为 5（默认）：5分钟粒度；10：10分钟粒度；20：20分钟粒度；60：1小时粒度；240：4小时粒度；480：8小时粒度；1440：1天粒度
	//		ResultType     Long     取值为0：多域名数据做合并；1：每个域名的数据分别返回

	//getProvinceAndIspPvData := make(map[string]interface{})
	//getProvinceAndIspPvData["DomainIds"] = "2D083FU"
	//getProvinceAndIspPvData["CdnType"] = "page"
	//getProvinceAndIspPvData["StartTime"] = "2021-05-31T01:14+0800"
	//getProvinceAndIspPvData["EndTime"] = "2021-05-31T01:24+0800"
	//getProvinceAndIspPvData["Provinces"] = "beijing"
	//getProvinceAndIspPvData["Isps"] = "CM"
	//getProvinceAndIspPvData["Granularity"] = 5
	//getProvinceAndIspPvData["ResultType"] = 1
	//
	//resp, err = svc.GetProvinceAndIspPvDataGet(&getProvinceAndIspPvData)
	//if err != nil {
	//	fmt.Println("error:", err.Error())
	//}
	//if resp != nil {
	//	str, _ := json.Marshal(&resp)
	//	fmt.Printf("%+v\n", string(str))
	//}

	//	**************************************** 获取域名在中国大陆地区各省份及各运营商的请求数数据，包括边缘请求数 POST（GetProvinceAndIspPvData）***************************************************
	// Parameters:
	//        DomainIds       String  域名ID，缺省为当前产品类型下的全部域名，可输入需要查询的域名ID，支持批量域名查询，多个域名ID用逗号（半角）分隔
	//        CdnType     String  产品类型，只允许输入一种类型，取值为video:音视频点播,file:大文件下载,live:流媒体直播
	//        StartTime       String  获取数据起始时间点，日期格式按ISO8601表示法，北京时间，格式为：YYYY-MM-DDThh:mm+0800，例如：2016-08-01T21:14+0800
	//        EndTime         String  结束时间需大于起始时间；获取日期格式按照ISO8601表示法，北京时间，格式为：YYYY-MM-DDThh:mm+0800，例如： 2016-08-01T21:14+0800
	//        Provinces     String     省份区域名称，取值详见枚举列表，支持多省份区域查询，多个省份区域用逗号（半角）分隔，缺省为全部省份区域
	//		Isps     String     运营商名称，取值详见枚举列表，支持多运营商查询，多个运营商用逗号（半角）分隔，缺省为全部运营商
	//		Granularity     Long     统计粒度，取值为 5（默认）：5分钟粒度；10：10分钟粒度；20：20分钟粒度；60：1小时粒度；240：4小时粒度；480：8小时粒度；1440：1天粒度
	//		ResultType     Long     取值为0：多域名数据做合并；1：每个域名的数据分别返回
	//
	//getProvinceAndIspPvData := make(map[string]interface{})
	//getProvinceAndIspPvData["DomainIds"] = "2D083FU"
	//getProvinceAndIspPvData["CdnType"] = "page"
	//getProvinceAndIspPvData["StartTime"] = "2021-05-31T01:14+0800"
	//getProvinceAndIspPvData["EndTime"] = "2021-05-31T01:24+0800"
	//getProvinceAndIspPvData["Provinces"] = "beijing"
	//getProvinceAndIspPvData["Isps"] = "CM"
	//getProvinceAndIspPvData["Granularity"] = 5
	//getProvinceAndIspPvData["ResultType"] = 1
	//
	//resp, err = svc.GetProvinceAndIspPvDataPost(&getProvinceAndIspPvData)
	//if err != nil {
	//	fmt.Println("error:", err.Error())
	//}
	//if resp != nil {
	//	str, _ := json.Marshal(&resp)
	//	fmt.Printf("%+v\n", string(str))
	//}

	//	**************************************** 获取域名一段时间内的回源Http状态码访问次数及占比数据 GET（GetSrcHttpCodeData）***************************************************
	// Parameters:
	//        DomainIds       String  域名ID，缺省为当前产品类型下的全部域名，可输入需要查询的域名ID，支持批量域名查询，多个域名ID用逗号（半角）分隔
	//        CdnType     String  产品类型，只允许输入一种类型，取值为video:音视频点播,file:大文件下载,live:流媒体直播，当前暂不支持直播类型
	//        StartTime       String  获取数据起始时间点，日期格式按ISO8601表示法，北京时间，格式为：YYYY-MM-DDThh:mm+0800，例如：2016-08-01T21:14+0800
	//        EndTime         String  结束时间需大于起始时间；获取日期格式按照ISO8601表示法，北京时间，格式为：YYYY-MM-DDThh:mm+0800，例如： 2016-08-01T21:14+0800
	//

	//getSrcHttpCodeData := make(map[string]interface{})
	//getSrcHttpCodeData["DomainIds"] = "2D083FU"
	//getSrcHttpCodeData["CdnType"] = "page"
	//getSrcHttpCodeData["StartTime"] = "2021-05-31T01:14+0800"
	//getSrcHttpCodeData["EndTime"] = "2021-05-31T01:24+0800"
	//
	//resp, err = svc.GetSrcHttpCodeDataGet(&getSrcHttpCodeData)
	//if err != nil {
	//	fmt.Println("error:", err.Error())
	//}
	//if resp != nil {
	//	str, _ := json.Marshal(&resp)
	//	fmt.Printf("%+v\n", string(str))
	//}

	//	**************************************** 获取域名一段时间内的回源Http状态码访问次数及占比数据 POST（GetSrcHttpCodeData）***************************************************
	// Parameters:
	//        DomainIds       String  域名ID，缺省为当前产品类型下的全部域名，可输入需要查询的域名ID，支持批量域名查询，多个域名ID用逗号（半角）分隔
	//        CdnType     String  产品类型，只允许输入一种类型，取值为video:音视频点播,file:大文件下载,live:流媒体直播，当前暂不支持直播类型
	//        StartTime       String  获取数据起始时间点，日期格式按ISO8601表示法，北京时间，格式为：YYYY-MM-DDThh:mm+0800，例如：2016-08-01T21:14+0800
	//        EndTime         String  结束时间需大于起始时间；获取日期格式按照ISO8601表示法，北京时间，格式为：YYYY-MM-DDThh:mm+0800，例如： 2016-08-01T21:14+0800

	//getSrcHttpCodeData := make(map[string]interface{})
	//getSrcHttpCodeData["DomainIds"] = "2D083FU"
	//getSrcHttpCodeData["CdnType"] = "page"
	//getSrcHttpCodeData["StartTime"] = "2021-05-31T01:14+0800"
	//getSrcHttpCodeData["EndTime"] = "2021-05-31T01:24+0800"
	//resp, err = svc.GetSrcHttpCodeDataPost(&getSrcHttpCodeData)
	//if err != nil {
	//	fmt.Println("error:", err.Error())
	//}
	//if resp != nil {
	//	str, _ := json.Marshal(&resp)
	//	fmt.Printf("%+v\n", string(str))
	//}

	//	**************************************** 获取域名的回源Http状态码详细访问次数及占比数据 GET（GetSrcHttpCodeDetailedData）***************************************************
	// Parameters:
	//        DomainIds       String  域名ID，缺省为当前产品类型下的全部域名，可输入需要查询的域名ID，支持批量域名查询，多个域名ID用逗号（半角）分隔
	//        CdnType     String  产品类型，只允许输入一种类型，取值为video:音视频点播,file:大文件下载,live:流媒体直播,当前暂不支持直播类型
	//        StartTime       String  获取数据起始时间点，日期格式按ISO8601表示法，北京时间，格式为：YYYY-MM-DDThh:mm+0800，例如：2016-08-01T21:14+0800
	//        EndTime         String  结束时间需大于起始时间；获取日期格式按照ISO8601表示法，北京时间，格式为：YYYY-MM-DDThh:mm+0800，例如： 2016-08-01T21:14+0800
	//		Granularity     Long     统计粒度，取值为 5（默认）：5分钟粒度；10：10分钟粒度；20：20分钟粒度；60：1小时粒度；240：4小时粒度；480：8小时粒度；1440：1天粒度
	//		ResultType     Long     取值为0：多域名数据做合并；1：每个域名的数据分别返回

	//getSrcHttpCodeDetailedData := make(map[string]interface{})
	//getSrcHttpCodeDetailedData["DomainIds"] = "2D083FU"
	//getSrcHttpCodeDetailedData["CdnType"] = "page"
	//getSrcHttpCodeDetailedData["StartTime"] = "2021-05-31T01:14+0800"
	//getSrcHttpCodeDetailedData["EndTime"] = "2021-05-31T01:24+0800"
	//getSrcHttpCodeDetailedData["Granularity"] = 5
	//getSrcHttpCodeDetailedData["ResultType"] = 0
	//
	//resp, err = svc.GetSrcHttpCodeDetailedDataGet(&getSrcHttpCodeDetailedData)
	//if err != nil {
	//	fmt.Println("error:", err.Error())
	//}
	//if resp != nil {
	//	str, _ := json.Marshal(&resp)
	//	fmt.Printf("%+v\n", string(str))
	//}

	//	**************************************** 获取域名的回源Http状态码详细访问次数及占比数据 POST（GetSrcHttpCodeDetailedData）***************************************************
	// Parameters:
	//        DomainIds       String  域名ID，缺省为当前产品类型下的全部域名，可输入需要查询的域名ID，支持批量域名查询，多个域名ID用逗号（半角）分隔
	//        CdnType     String  产品类型，只允许输入一种类型，取值为video:音视频点播,file:大文件下载,live:流媒体直播,当前暂不支持直播类型
	//        StartTime       String  获取数据起始时间点，日期格式按ISO8601表示法，北京时间，格式为：YYYY-MM-DDThh:mm+0800，例如：2016-08-01T21:14+0800
	//        EndTime         String  结束时间需大于起始时间；获取日期格式按照ISO8601表示法，北京时间，格式为：YYYY-MM-DDThh:mm+0800，例如： 2016-08-01T21:14+0800
	//		Granularity     Long     统计粒度，取值为 5（默认）：5分钟粒度；10：10分钟粒度；20：20分钟粒度；60：1小时粒度；240：4小时粒度；480：8小时粒度；1440：1天粒度
	//		ResultType     Long     取值为0：多域名数据做合并；1：每个域名的数据分别返回

	//getSrcHttpCodeDetailedData := make(map[string]interface{})
	//getSrcHttpCodeDetailedData["DomainIds"] = "2D083FU"
	//getSrcHttpCodeDetailedData["CdnType"] = "page"
	//getSrcHttpCodeDetailedData["StartTime"] = "2021-05-31T01:14+0800"
	//getSrcHttpCodeDetailedData["EndTime"] = "2021-05-31T01:24+0800"
	//getSrcHttpCodeDetailedData["Granularity"] = 5
	//getSrcHttpCodeDetailedData["ResultType"] = 0
	//
	//resp, err = svc.GetSrcHttpCodeDetailedDataPost(&getSrcHttpCodeDetailedData)
	//if err != nil {
	//	fmt.Println("error:", err.Error())
	//}
	//if resp != nil {
	//	str, _ := json.Marshal(&resp)
	//	fmt.Printf("%+v\n", string(str))
	//}

	//	**************************************** 本接口用于获取某段时间内按一级目录为维度下消耗的带宽 GET（GetBandwidthDataByDir）***************************************************
	// Parameters:
	//        DomainId       String  输入需要查询的域名ID，只允许输入一个
	//        StartTime       String  获取数据起始时间点，日期格式按ISO8601表示法，北京时间，格式为：YYYY-MM-DDThh:mm+0800，例如：2016-08-01T21:14+0800
	//        EndTime         String  结束时间需大于起始时间；获取日期格式按照ISO8601表示法，北京时间，格式为：YYYY-MM-DDThh:mm+0800，例如： 2016-08-01T21:14+0800
	//		Granularity     Long     统计粒度，取值为 5（默认）：5分钟粒度；10：10分钟粒度；20：20分钟粒度；60：1小时粒度；240：4小时粒度；480：8小时粒度；1440：1天粒度
	//		ResultType     Long     取值为0：多域名数据做合并；1：每个域名的数据分别返回
	//		Regions       String  区域名称， 取值为CN:中国大陆，HK：香港，TW：台湾，AS：亚洲其他，NA：北美洲，SA：南美洲，EU：欧洲，AU：大洋洲，AF：非洲，支持多区域查询，多个区域用逗号（半角）分隔，缺省为 CN
	//		Dirs       String  目录名称，支持统计域名下一级目录，即请求URL中域名后的第一个“\/”和第二个“\/”之间的内容;支持批量查询，多个目录用逗号（半角）分隔，缺省为该域名下所有一级目录及“\/”；若输入\/，则查询该域名下所有无一级目录的URL带宽合并值

	//getBandwidthDataByDir := make(map[string]interface{})
	//getBandwidthDataByDir["DomainId"] = "2D083FU"
	//getBandwidthDataByDir["StartTime"] = "2021-05-31T01:14+0800"
	//getBandwidthDataByDir["EndTime"] = "2021-05-31T01:24+0800"
	//getBandwidthDataByDir["Granularity"] = 5
	//getBandwidthDataByDir["ResultType"] = 0
	//getBandwidthDataByDir["Regions"] = "CN"
	//getBandwidthDataByDir["Dirs"] = "/"
	//
	//resp, err = svc.GetBandwidthDataByDirGet(&getBandwidthDataByDir)
	//if err != nil {
	//	fmt.Println("error:", err.Error())
	//}
	//if resp != nil {
	//	str, _ := json.Marshal(&resp)
	//	fmt.Printf("%+v\n", string(str))
	//}

	//	**************************************** 本接口用于获取某段时间内按一级目录为维度下消耗的带宽 POST（GetBandwidthDataByDir）***************************************************
	// Parameters:
	//        DomainId       String  输入需要查询的域名ID，只允许输入一个
	//        StartTime       String  获取数据起始时间点，日期格式按ISO8601表示法，北京时间，格式为：YYYY-MM-DDThh:mm+0800，例如：2016-08-01T21:14+0800
	//        EndTime         String  结束时间需大于起始时间；获取日期格式按照ISO8601表示法，北京时间，格式为：YYYY-MM-DDThh:mm+0800，例如： 2016-08-01T21:14+0800
	//		Granularity     Long     统计粒度，取值为 5（默认）：5分钟粒度；10：10分钟粒度；20：20分钟粒度；60：1小时粒度；240：4小时粒度；480：8小时粒度；1440：1天粒度
	//		ResultType     Long     取值为0：多域名数据做合并；1：每个域名的数据分别返回
	//		Regions       String  区域名称， 取值为CN:中国大陆，HK：香港，TW：台湾，AS：亚洲其他，NA：北美洲，SA：南美洲，EU：欧洲，AU：大洋洲，AF：非洲，支持多区域查询，多个区域用逗号（半角）分隔，缺省为 CN
	//		Dirs       String  目录名称，支持统计域名下一级目录，即请求URL中域名后的第一个“\/”和第二个“\/”之间的内容;支持批量查询，多个目录用逗号（半角）分隔，缺省为该域名下所有一级目录及“\/”；若输入\/，则查询该域名下所有无一级目录的URL带宽合并值

	//getBandwidthDataByDir := make(map[string]interface{})
	//getBandwidthDataByDir["DomainId"] = "2D083FU"
	//getBandwidthDataByDir["StartTime"] = "2021-05-31T01:14+0800"
	//getBandwidthDataByDir["EndTime"] = "2021-05-31T01:24+0800"
	//getBandwidthDataByDir["Granularity"] = 5
	//getBandwidthDataByDir["ResultType"] = 0
	//getBandwidthDataByDir["Regions"] = "CN"
	//getBandwidthDataByDir["Dirs"] = "/"
	//
	//resp, err = svc.GetBandwidthDataByDirPost(&getBandwidthDataByDir)
	//if err != nil {
	//	fmt.Println("error:", err.Error())
	//}
	//if resp != nil {
	//	str, _ := json.Marshal(&resp)
	//	fmt.Printf("%+v\n", string(str))
	//}

	//	**************************************** 本接口用于获取某段时间内按一级目录为维度下消耗的流量 GET（GetFlowDataByDir）***************************************************
	// Parameters:
	//        DomainId       String  输入需要查询的域名ID，只允许输入一个
	//        StartTime       String  获取数据起始时间点，日期格式按ISO8601表示法，北京时间，格式为：YYYY-MM-DDThh:mm+0800，例如：2016-08-01T21:14+0800
	//        EndTime         String  结束时间需大于起始时间；获取日期格式按照ISO8601表示法，北京时间，格式为：YYYY-MM-DDThh:mm+0800，例如： 2016-08-01T21:14+0800
	//		Granularity     Long     统计粒度，取值为 5（默认）：5分钟粒度；10：10分钟粒度；20：20分钟粒度；60：1小时粒度；240：4小时粒度；480：8小时粒度；1440：1天粒度
	//		ResultType     Long     取值为0：多域名数据做合并；1：每个域名的数据分别返回
	//		Regions       String  区域名称， 取值为CN:中国大陆，HK：香港，TW：台湾，AS：亚洲其他，NA：北美洲，SA：南美洲，EU：欧洲，AU：大洋洲，AF：非洲，支持多区域查询，多个区域用逗号（半角）分隔，缺省为 CN
	//		Dirs       String  目录名称，支持统计域名下一级目录，即请求URL中域名后的第一个“\/”和第二个“\/”之间的内容;支持批量查询，多个目录用逗号（半角）分隔，缺省为该域名下所有一级目录及“\/”；若输入\/，则查询该域名下所有无一级目录的URL带宽合并值

	//getFlowDataByDir := make(map[string]interface{})
	//getFlowDataByDir["DomainId"] = "2D083FU"
	//getFlowDataByDir["StartTime"] = "2021-05-31T01:14+0800"
	//getFlowDataByDir["EndTime"] = "2021-05-31T01:24+0800"
	//getFlowDataByDir["Granularity"] = 5
	//getFlowDataByDir["ResultType"] = 0
	//getFlowDataByDir["Regions"] = "CN"
	//getFlowDataByDir["Dirs"] = "/"
	//
	//resp, err = svc.GetFlowDataByDirGet(&getFlowDataByDir)
	//if err != nil {
	//	fmt.Println("error:", err.Error())
	//}
	//if resp != nil {
	//	str, _ := json.Marshal(&resp)
	//	fmt.Printf("%+v\n", string(str))
	//}

	//	**************************************** 本接口用于获取某段时间内按一级目录为维度下消耗的流量 POST（GetFlowDataByDir）***************************************************
	// Parameters:
	//        DomainId       String  输入需要查询的域名ID，只允许输入一个
	//        StartTime       String  获取数据起始时间点，日期格式按ISO8601表示法，北京时间，格式为：YYYY-MM-DDThh:mm+0800，例如：2016-08-01T21:14+0800
	//        EndTime         String  结束时间需大于起始时间；获取日期格式按照ISO8601表示法，北京时间，格式为：YYYY-MM-DDThh:mm+0800，例如： 2016-08-01T21:14+0800
	//		Granularity     Long     统计粒度，取值为 5（默认）：5分钟粒度；10：10分钟粒度；20：20分钟粒度；60：1小时粒度；240：4小时粒度；480：8小时粒度；1440：1天粒度
	//		ResultType     Long     取值为0：多域名数据做合并；1：每个域名的数据分别返回
	//		Regions       String  区域名称， 取值为CN:中国大陆，HK：香港，TW：台湾，AS：亚洲其他，NA：北美洲，SA：南美洲，EU：欧洲，AU：大洋洲，AF：非洲，支持多区域查询，多个区域用逗号（半角）分隔，缺省为 CN
	//		Dirs       String  目录名称，支持统计域名下一级目录，即请求URL中域名后的第一个“\/”和第二个“\/”之间的内容;支持批量查询，多个目录用逗号（半角）分隔，缺省为该域名下所有一级目录及“\/”；若输入\/，则查询该域名下所有无一级目录的URL带宽合并值

	//getFlowDataByDir := make(map[string]interface{})
	//getFlowDataByDir["DomainId"] = "2D083FU"
	//getFlowDataByDir["StartTime"] = "2021-05-31T01:14+0800"
	//getFlowDataByDir["EndTime"] = "2021-05-31T01:24+0800"
	//getFlowDataByDir["Granularity"] = 5
	//getFlowDataByDir["ResultType"] = 0
	//getFlowDataByDir["Regions"] = "CN"
	//getFlowDataByDir["Dirs"] = "/"
	//
	//resp, err = svc.GetFlowDataByDirPost(&getFlowDataByDir)
	//if err != nil {
	//	fmt.Println("error:", err.Error())
	//}
	//if resp != nil {
	//	str, _ := json.Marshal(&resp)
	//	fmt.Printf("%+v\n", string(str))
	//}

	//	**************************************** 本接口用于获取直播流维度的平均观看时长数据 GET（GetPlayTimeDataByStream）***************************************************
	// Parameters:
	//        StreamUrls      String  流名，支持批量查询，多个流名用逗号（半角）分隔
	//        StartTime       String  获取数据起始时间点，日期格式按ISO8601表示法，北京时间，格式为：YYYY-MM-DDThh:mm+0800，例如：2016-08-01T21:14+0800
	//        EndTime         String  结束时间需大于起始时间；获取日期格式按照ISO8601表示法，北京时间，格式为：YYYY-MM-DDThh:mm+0800，例如： 2016-08-01T21:14+0800
	//		Granularity     integer     统计粒度，取值为 5（默认）：5分钟粒度；10：10分钟粒度；20：20分钟粒度；60：1小时粒度；240：4小时粒度；480：8小时粒度；1440：1天粒度
	//		ResultType      integer     取值为0：多域名数据做合并；1：每个域名的数据分别返回
	//		Regions         String  区域名称， 取值为CN:中国大陆，HK：香港，TW：台湾，AS：亚洲其他，NA：北美洲，SA：南美洲，EU：欧洲，AU：大洋洲，AF：非洲，支持多区域查询，多个区域用逗号（半角）分隔，缺省为 CN

	//getPlayTimeDataByStream := make(map[string]interface{})
	//getPlayTimeDataByStream["StreamUrls"] = "http://c.gdown.baidu.com/live/m_defa5e0dd0d324101472363734966100.flv"
	//getPlayTimeDataByStream["StartTime"] = "2021-05-31T01:14+0800"
	//getPlayTimeDataByStream["EndTime"] = "2021-05-31T01:24+0800"
	//getPlayTimeDataByStream["Granularity"] = 5
	//getPlayTimeDataByStream["ResultType"] = 0
	//getPlayTimeDataByStream["Regions"] = "CN"
	//
	//resp, err = svc.GetPlayTimeDataByStreamGet(&getPlayTimeDataByStream)
	//if err != nil {
	//	fmt.Println("error:", err.Error())
	//}
	//if resp != nil {
	//	str, _ := json.Marshal(&resp)
	//	fmt.Printf("%+v\n", string(str))
	//}

	//	**************************************** 本接口用于获取直播流维度的平均观看时长数据 POST（GetPlayTimeDataByStream）***************************************************
	// Parameters:
	//        StreamUrls      String  流名，支持批量查询，多个流名用逗号（半角）分隔
	//        StartTime       String  获取数据起始时间点，日期格式按ISO8601表示法，北京时间，格式为：YYYY-MM-DDThh:mm+0800，例如：2016-08-01T21:14+0800
	//        EndTime         String  结束时间需大于起始时间；获取日期格式按照ISO8601表示法，北京时间，格式为：YYYY-MM-DDThh:mm+0800，例如： 2016-08-01T21:14+0800
	//		Granularity     integer     统计粒度，取值为 5（默认）：5分钟粒度；10：10分钟粒度；20：20分钟粒度；60：1小时粒度；240：4小时粒度；480：8小时粒度；1440：1天粒度
	//		ResultType      integer     取值为0：多域名数据做合并；1：每个域名的数据分别返回
	//		Regions         String  区域名称， 取值为CN:中国大陆，HK：香港，TW：台湾，AS：亚洲其他，NA：北美洲，SA：南美洲，EU：欧洲，AU：大洋洲，AF：非洲，支持多区域查询，多个区域用逗号（半角）分隔，缺省为 CN

	//getPlayTimeDataByStream := make(map[string]interface{})
	//getPlayTimeDataByStream["StreamUrls"] = "http://c.gdown.baidu.com/live/m_defa5e0dd0d324101472363734966100.flv"
	//getPlayTimeDataByStream["StartTime"] = "2021-05-31T01:14+0800"
	//getPlayTimeDataByStream["EndTime"] = "2021-05-31T01:24+0800"
	//getPlayTimeDataByStream["Granularity"] = 5
	//getPlayTimeDataByStream["ResultType"] = 0
	//getPlayTimeDataByStream["Regions"] = "CN"
	//
	//resp, err = svc.GetPlayTimeDataByStreamPost(&getPlayTimeDataByStream)
	//if err != nil {
	//	fmt.Println("error:", err.Error())
	//}
	//if resp != nil {
	//	str, _ := json.Marshal(&resp)
	//	fmt.Printf("%+v\n", string(str))
	//}

	//	**************************************** 本接口用于获取直播域名维度的观看时长数据 GET（GetPlayTimeDataByDomain）***************************************************
	// Parameters:
	//        DomainIds      String  流名，支持批量查询，多个流名用逗号（半角）分隔
	//        StartTime       String  获取数据起始时间点，日期格式按ISO8601表示法，北京时间，格式为：YYYY-MM-DDThh:mm+0800，例如：2016-08-01T21:14+0800
	//        EndTime         String  结束时间需大于起始时间；获取日期格式按照ISO8601表示法，北京时间，格式为：YYYY-MM-DDThh:mm+0800，例如： 2016-08-01T21:14+0800
	//		Granularity     integer     统计粒度，取值为 5（默认）：5分钟粒度；10：10分钟粒度；20：20分钟粒度；60：1小时粒度；240：4小时粒度；480：8小时粒度；1440：1天粒度
	//		ResultType      integer     取值为0：多域名数据做合并；1：每个域名的数据分别返回
	//		Regions         String  区域名称， 取值为CN:中国大陆，HK：香港，TW：台湾，AS：亚洲其他，NA：北美洲，SA：南美洲，EU：欧洲，AU：大洋洲，AF：非洲，支持多区域查询，多个区域用逗号（半角）分隔，缺省为 CN

	//getPlayTimeDataByDomain := make(map[string]interface{})
	//getPlayTimeDataByDomain["DomainIds"] = "2D07W0T"
	//getPlayTimeDataByDomain["StartTime"] = "2021-05-31T01:14+0800"
	//getPlayTimeDataByDomain["EndTime"] = "2021-05-31T01:24+0800"
	//getPlayTimeDataByDomain["Granularity"] = 5
	//getPlayTimeDataByDomain["ResultType"] = 0
	//getPlayTimeDataByDomain["Regions"] = "CN"
	//
	//resp, err = svc.GetPlayTimeDataByDomainGet(&getPlayTimeDataByDomain)
	//if err != nil {
	//	fmt.Println("error:", err.Error())
	//}
	//if resp != nil {
	//	str, _ := json.Marshal(&resp)
	//	fmt.Printf("%+v\n", string(str))
	//}

	//	**************************************** 本接口用于获取直播域名维度的观看时长数据 POST（GetPlayTimeDataByDomain）***************************************************
	// Parameters:
	//      DomainIds      String  流名，支持批量查询，多个流名用逗号（半角）分隔
	//      StartTime       String  获取数据起始时间点，日期格式按ISO8601表示法，北京时间，格式为：YYYY-MM-DDThh:mm+0800，例如：2016-08-01T21:14+0800
	//      EndTime         String  结束时间需大于起始时间；获取日期格式按照ISO8601表示法，北京时间，格式为：YYYY-MM-DDThh:mm+0800，例如： 2016-08-01T21:14+0800
	//		Granularity     integer     统计粒度，取值为 5（默认）：5分钟粒度；10：10分钟粒度；20：20分钟粒度；60：1小时粒度；240：4小时粒度；480：8小时粒度；1440：1天粒度
	//		ResultType      integer     取值为0：多域名数据做合并；1：每个域名的数据分别返回
	//		Regions         String  区域名称， 取值为CN:中国大陆，HK：香港，TW：台湾，AS：亚洲其他，NA：北美洲，SA：南美洲，EU：欧洲，AU：大洋洲，AF：非洲，支持多区域查询，多个区域用逗号（半角）分隔，缺省为 CN

	//getPlayTimeDataByDomain := make(map[string]interface{})
	//getPlayTimeDataByDomain["DomainIds"] = "2D07W0T"
	//getPlayTimeDataByDomain["StartTime"] = "2021-05-31T01:14+0800"
	//getPlayTimeDataByDomain["EndTime"] = "2021-05-31T01:24+0800"
	//getPlayTimeDataByDomain["Granularity"] = 5
	//getPlayTimeDataByDomain["ResultType"] = 0
	//getPlayTimeDataByDomain["Regions"] = "CN"
	//
	//resp, err = svc.GetPlayTimeDataByDomainPost(&getPlayTimeDataByDomain)
	//if err != nil {
	//	fmt.Println("error:", err.Error())
	//}
	//if resp != nil {
	//	str, _ := json.Marshal(&resp)
	//	fmt.Printf("%+v\n", string(str))
	//}

	//	**************************************** 获取用户当前的计费方式 GET（GetBillingMode）***************************************************
	// Parameters:
	//        CdnType   String   产品类型，只允许输入一种类型，取值为download:下载类加速,；live:直播加速

	//getBillingMode := make(map[string]interface{})
	//getBillingMode["CdnType"] = "download"
	//
	//resp, err = svc.GetBillingModeGet(&getBillingMode)
	//if err != nil {
	//	fmt.Println("error:", err.Error())
	//}
	//if resp != nil {
	//	str, _ := json.Marshal(&resp)
	//	fmt.Printf("%+v\n", string(str))
	//}

	//	**************************************** 获取用户当前的计费方式 POST（GetBillingMode）***************************************************
	// Parameters:
	//        CdnType   String   产品类型，只允许输入一种类型，取值为download:下载类加速,；live:直播加速

	//getBillingMode := make(map[string]interface{})
	//getBillingMode["CdnType"] = "download"
	//
	//resp, err = svc.GetBillingModePost(&getBillingMode)
	//if err != nil {
	//	fmt.Println("error:", err.Error())
	//}
	//if resp != nil {
	//	str, _ := json.Marshal(&resp)
	//	fmt.Printf("%+v\n", string(str))
	//}

	//	**************************************** 获取域名的计费数据 GET（GetBillingData）***************************************************
	// Parameters:
	//        DomainIds       String  域名ID，缺省为当前产品类型下的全部域名，可输入需要查询的域名ID，支持批量域名查询，多个域名ID用逗号（半角）分隔
	//        StartTime       String  获取数据起始时间点，日期格式按ISO8601表示法，北京时间，格式为：YYYY-MM-DDThh:mm+0800，例如：2016-08-01T21:14+0800
	//        EndTime         String  结束时间需大于起始时间；获取日期格式按照ISO8601表示法，北京时间，格式为：YYYY-MM-DDThh:mm+0800，例如： 2016-08-01T21:14+0800
	//        CdnType     String  产品类型，只允许输入一种类型，取值为video:音视频点播,file:大文件下载,live:流媒体直播
	//        Regions         String  区域名称， 取值为CN:中国大陆，HK：香港，TW：台湾，AS：亚洲其他，NA：北美洲，SA：南美洲，EU：欧洲，AU：大洋洲，AF：非洲，支持多区域查询，多个区域用逗号（半角）分隔，缺省为 CN
	//        BillingMode     String  计费方式， 取值为 peakbw:峰值计费;peak95bw:95峰值计费;averagebw：日峰值平均值计费；monthflow：流量按月，只允许输入一种计费方式，缺省为 peakbw ；
	//

	//getBillingData := make(map[string]interface{})
	//getBillingData["DomainIds"] = "2D07WRH"
	//getBillingData["StartTime"] = "2021-05-31T01:14+0800"
	//getBillingData["EndTime"] = "2021-05-31T01:24+0800"
	//getBillingData["CdnType"] = "video"
	//getBillingData["BillingMode"] = "peakbw"
	//getBillingData["Regions"] = "CN"
	//
	//resp, err = svc.GetBillingDataGet(&getBillingData)
	//if err != nil {
	//	fmt.Println("error:", err.Error())
	//}
	//if resp != nil {
	//	str, _ := json.Marshal(&resp)
	//	fmt.Printf("%+v\n", string(str))
	//}

	//	**************************************** 获取域名的计费数据 POST（GetBillingData）***************************************************
	// Parameters:
	//        DomainIds       String  域名ID，缺省为当前产品类型下的全部域名，可输入需要查询的域名ID，支持批量域名查询，多个域名ID用逗号（半角）分隔
	//        StartTime       String  获取数据起始时间点，日期格式按ISO8601表示法，北京时间，格式为：YYYY-MM-DDThh:mm+0800，例如：2016-08-01T21:14+0800
	//        EndTime         String  结束时间需大于起始时间；获取日期格式按照ISO8601表示法，北京时间，格式为：YYYY-MM-DDThh:mm+0800，例如： 2016-08-01T21:14+0800
	//        CdnType     String  产品类型，只允许输入一种类型，取值为video:音视频点播,file:大文件下载,live:流媒体直播
	//        Regions         String  区域名称， 取值为CN:中国大陆，HK：香港，TW：台湾，AS：亚洲其他，NA：北美洲，SA：南美洲，EU：欧洲，AU：大洋洲，AF：非洲，支持多区域查询，多个区域用逗号（半角）分隔，缺省为 CN
	//        BillingMode     String  计费方式， 取值为 peakbw:峰值计费;peak95bw:95峰值计费;averagebw：日峰值平均值计费；monthflow：流量按月，只允许输入一种计费方式，缺省为 peakbw ；
	//

	//getBillingData := make(map[string]interface{})
	//getBillingData["DomainIds"] = "2D07WRH"
	//getBillingData["StartTime"] = "2021-05-31T01:14+0800"
	//getBillingData["EndTime"] = "2021-05-31T01:24+0800"
	//getBillingData["CdnType"] = "video"
	//getBillingData["BillingMode"] = "peakbw"
	//getBillingData["Regions"] = "CN"
	//
	//resp, err = svc.GetBillingDataPost(&getBillingData)
	//if err != nil {
	//	fmt.Println("error:", err.Error())
	//}
	//if resp != nil {
	//	str, _ := json.Marshal(&resp)
	//	fmt.Printf("%+v\n", string(str))
	//}

	//	**************************************** 获取域名当前的服务节点IP列表，用于分析域名服务节点运行状况，便于故障排查 GET（GetServiceIpData）***************************************************
	// Parameters:
	//        DomainId       String  域名ID，输入需要查询的域名ID，仅支持单个域名ID

	//getServiceIpData := make(map[string]interface{})
	//getServiceIpData["DomainId"] = "2D07WRH"
	//
	//resp, err = svc.GetServiceIpDataGet(&getServiceIpData)
	//if err != nil {
	//	fmt.Println("error:", err.Error())
	//}
	//if resp != nil {
	//	str, _ := json.Marshal(&resp)
	//	fmt.Printf("%+v\n", string(str))
	//}

	//	**************************************** 获取域名当前的服务节点IP列表，用于分析域名服务节点运行状况，便于故障排查 POST（GetServiceIpData）***************************************************
	// Parameters:
	//        DomainId       String  域名ID，输入需要查询的域名ID，仅支持单个域名ID

	//getServiceIpData := make(map[string]interface{})
	//getServiceIpData["DomainId"] = "2D07WRH"
	//
	//resp, err = svc.GetServiceIpDataPost(&getServiceIpData)
	//if err != nil {
	//	fmt.Println("error:", err.Error())
	//}
	//if resp != nil {
	//	str, _ := json.Marshal(&resp)
	//	fmt.Printf("%+v\n", string(str))
	//}

	//	**************************************** 获取域名带宽峰值，峰值时间点 GET（GetPeakBandwidthData）***************************************************
	//  Parameters:
	//        DomainIds       String  域名ID，缺省为当前产品类型下的全部域名，可输入需要查询的域名ID，支持批量域名查询，多个域名ID用逗号（半角）分隔
	//        StartTime       String  获取数据起始时间点，日期格式按ISO8601表示法，北京时间，格式为：YYYY-MM-DDThh:mm+0800，例如：2016-08-01T21:14+0800
	//        EndTime         String  结束时间需大于起始时间；获取日期格式按照ISO8601表示法，北京时间，格式为：YYYY-MM-DDThh:mm+0800，例如： 2016-08-01T21:14+0800
	//        CdnType     String  产品类型，只允许输入一种类型，取值为video:音视频点播,file:大文件下载,live:流媒体直播
	//        Regions         String  区域名称， 取值为CN:中国大陆，HK：香港，TW：台湾，AS：亚洲其他，NA：北美洲，SA：南美洲，EU：欧洲，AU：大洋洲，AF：非洲，支持多区域查询，多个区域用逗号（半角）分隔，缺省为 CN

	//getPeakBandwidthData := make(map[string]interface{})
	//getPeakBandwidthData["DomainIds"] = "2D07WRH"
	//getPeakBandwidthData["StartTime"] = "2021-05-31T01:14+0800"
	//getPeakBandwidthData["EndTime"] = "2021-05-31T01:24+0800"
	//getPeakBandwidthData["CdnType"] = "video"
	//getPeakBandwidthData["Regions"] = "CN"
	//
	//resp, err = svc.GetPeakBandwidthDataGet(&getPeakBandwidthData)
	//if err != nil {
	//	fmt.Println("error:", err.Error())
	//}
	//if resp != nil {
	//	str, _ := json.Marshal(&resp)
	//	fmt.Printf("%+v\n", string(str))
	//}

	//	**************************************** 获取域名带宽峰值，峰值时间点 POST（GetPeakBandwidthData）***************************************************
	//  Parameters:
	//        DomainIds       String  域名ID，缺省为当前产品类型下的全部域名，可输入需要查询的域名ID，支持批量域名查询，多个域名ID用逗号（半角）分隔
	//        StartTime       String  获取数据起始时间点，日期格式按ISO8601表示法，北京时间，格式为：YYYY-MM-DDThh:mm+0800，例如：2016-08-01T21:14+0800
	//        EndTime         String  结束时间需大于起始时间；获取日期格式按照ISO8601表示法，北京时间，格式为：YYYY-MM-DDThh:mm+0800，例如： 2016-08-01T21:14+0800
	//        CdnType     String  产品类型，只允许输入一种类型，取值为video:音视频点播,file:大文件下载,live:流媒体直播
	//        Regions         String  区域名称， 取值为CN:中国大陆，HK：香港，TW：台湾，AS：亚洲其他，NA：北美洲，SA：南美洲，EU：欧洲，AU：大洋洲，AF：非洲，支持多区域查询，多个区域用逗号（半角）分隔，缺省为 CN

	//getPeakBandwidthData := make(map[string]interface{})
	//getPeakBandwidthData["DomainIds"] = "2D07WRH"
	//getPeakBandwidthData["StartTime"] = "2021-05-31T01:14+0800"
	//getPeakBandwidthData["EndTime"] = "2021-05-31T01:24+0800"
	//getPeakBandwidthData["CdnType"] = "video"
	//getPeakBandwidthData["Regions"] = "CN"
	//
	//resp, err = svc.GetPeakBandwidthDataPost(&getPeakBandwidthData)
	//if err != nil {
	//	fmt.Println("error:", err.Error())
	//}
	//if resp != nil {
	//	str, _ := json.Marshal(&resp)
	//	fmt.Printf("%+v\n", string(str))
	//}

	//	**************************************** 本接口用于屏蔽、解除屏蔽URL。 POST（BlockDomainUrl）***************************************************
	//  Parameters:
	//        BlockType         string    操作接口名，系统规定参数 取值：block：屏蔽URL；unblock：解除屏蔽；
	//        Urls              Url[]    URL列表
	//        其中url[]为：
	//		Url String 需要提交屏蔽/解屏蔽的Url，单条

	//url := make(map[string]interface{});
	//url["Url"] = "http://ntj14122.test.com/abc.txt";
	//
	//urls := make([]map[string]interface{}, 0)
	//urls = append(urls, url)
	//blockDomainUrl := make(map[string]interface{})
	//blockDomainUrl["BlockType"] = "unblock"
	//blockDomainUrl["Urls"] = urls
	//
	//resp, err = svc.BlockDomainUrlPost(&blockDomainUrl)
	//if err != nil {
	//	fmt.Println("error:", err.Error())
	//}
	//if resp != nil {
	//	str, _ := json.Marshal(&resp)
	//	fmt.Printf("%+v\n", string(str))
	//}

	//	**************************************** 本接口用于获取屏蔽URL任务进度百分比及状态，查看任务是否在全网生效。 POST（GetBlockUrlTask）***************************************************
	//  Parameters:
	//            Urls              Url[]    URL列表
	//            PageSize          Long     分页大小，取值为1-50，最大50，默认20
	//            PageNumber        String   取得第几页，取值为：1-100000，最大100000，默认1
	//            其中url[]为：
	//		        Url String 需要提交查询的Url，单条

	//url := make(map[string]interface{});
	//url["Url"] = "http://ntj14122.test.com/abc.txt";
	//
	//urls := make([]map[string]interface{}, 0)
	//urls = append(urls, url)
	//getBlockUrlTask := make(map[string]interface{})
	//getBlockUrlTask["PageSize"] = 1
	//getBlockUrlTask["PageNumber"] = 10
	//getBlockUrlTask["Urls"] = urls
	//
	//resp, err = svc.GetBlockUrlTaskPost(&getBlockUrlTask)
	//if err != nil {
	//	fmt.Println("error:", err.Error())
	//}
	//if resp != nil {
	//	str, _ := json.Marshal(&resp)
	//	fmt.Printf("%+v\n", string(str))
	//}

	//	**************************************** 获取屏蔽URL最大限制数量，及剩余的条数 GET（GetBlockUrlQuota）***************************************************
	//resp, err = svc.GetBlockUrlQuotaGet(nil)
	//if err != nil {
	//	fmt.Println("error:", err.Error())
	//}
	//if resp != nil {
	//	str, _ := json.Marshal(&resp)
	//	fmt.Printf("%+v\n", string(str))
	//}

	//	**************************************** 为单域名或者多域名配置证书，多域名用英文半角逗号隔开 GET（ConfigCertificate）***************************************************
	//  Parameters:
	//            Enable            String    开启、关闭设置服务证书，取值：on：开启，off：关闭，默认为off，当选择开启时，以下为必填
	//            DomainIds         String    domainid列表，英文半角逗号隔开
	//            CertificateId     String    金山云生成的证书唯一性ID，若输入证书ID，则以下内容可不填写，若无证书ID，则以下内容为必填
	//            CertificateName   String    证书名称
	//            ServerCertificate String    域名对应的证书内容
	//            PrivateKey        String    证书对应的私钥内容

	//configCertificate := make(map[string]interface{})
	//configCertificate["Enable"] = "on"
	//configCertificate["DomainIds"] = "2D09VN9"
	//configCertificate["ServerCertificate"] = "-----BEGIN CERTIFICATE-----\\n MIIDiTCCAnGgAwIBAgIJANF7PDOT9Wm8MA0GCSqGSIb3DQEBCwUAMFsxCzAJBgNV\\n BAYTAmFoMQswCQYDVQQIDAJiajELMAkGA1UEBwwCYmoxETAPBgNVBAoMCGtpbmdz\\n b2Z0MQowCAYDVQQLDAFzMRMwEQYDVQQDDAphYWEuY29tLmNuMB4XDTE3MDcxMzA4\\n MTA1NFoXDTE3MDgxMjA4MTA1NFowWzELMAkGA1UEBhMCYWgxCzAJBgNVBAgMAmJq\\n MQswCQYDVQQHDAJiajERMA8GA1UECgwIa2luZ3NvZnQxCjAIBgNVBAsMAXMxEzAR\\n BgNVBAMMCmFhYS5jb20uY24wggEiMA0GCSqGSIb3DQEBAQUAA4IBDwAwggEKAoIB\\n AQCwp3KNrIW288L93DGyuBQ1eVL2XVW2IsCW3SKuntiTtFxA5+xecKN1xoTbNrJp\\n oWjIHUow/DRyH2k0j1zHrWs7mHfx/1LR2G6PfCGUgvA2EDNaoDia3gIPmmt6QUVh\\n l2XFqUrjmcaDk870dy8o5VMwblaC/mFhEFwlxGAjM6GqHHVlZ0AeXR+NJ7F3Fu85\\n 8dP/zFfb2kSzpjgsBuAQ2W5+fxtvlKHZlMfLE4OuQiME8JhvL/czzOVK9q2wGUzl\\n /IkbyG3SwLoLFrTry2TsSibdfZe8ZrpS3gKaNqwSg1RVvShDlyvLfjm6nLgV7l2i\\n fO9vHsaHjdmb9mUwK/soRvNFAgMBAAGjUDBOMB0GA1UdDgQWBBTK3MSboh5vOtTF\\n /i7ly75c7uHPJjAfBgNVHSMEGDAWgBTK3MSboh5vOtTF/i7ly75c7uHPJjAMBgNV\\n HRMEBTADAQH/MA0GCSqGSIb3DQEBCwUAA4IBAQCtVHjp0C94iHqXS2E1e2gdP8eU\\n qg2naXxMRCXo2hSLEcPmNhv2Z8SfCpu54CInEyINxZg58z7KxNFD3AmNcj1/r67L\\n lVmkZzZrNsLT3Ti3SYiPgRUOSTiABZZrr0HRRbD+N32mLJ/2IeuBneBM6o0ER3Lp\\n HDD+hVcB15+MmC0uj68kWCbYcfd3go6XiPpmVyaPkTkhYQDHufW0SZ+dILvDNIVy\\n rAh7ZEUiP81LxYZcsf8fPKA/MmscQ4wZ2Owxt0KecgrkXcPTeOBdFUQAanmdsyLt\\n +U+EGOF31pUgw7FLg5Tb0ejhE1aOmjYybWCt6FTldPBHHEDCcQZjOEXn56YG\\n -----END CERTIFICATE-----\\n "
	//configCertificate["CertificateName"] = "2D09VN9_test1"
	//configCertificate["PrivateKey"] = "-----BEGIN PRIVATE KEY-----\\n MIIEvwIBADANBgkqhkiG9w0BAQEFAASCBKkwggSlAgEAAoIBAQCwp3KNrIW288L9\\n 3DGyuBQ1eVL2XVW2IsCW3SKuntiTtFxA5+xecKN1xoTbNrJpoWjIHUow/DRyH2k0\\n j1zHrWs7mHfx/1LR2G6PfCGUgvA2EDNaoDia3gIPmmt6QUVhl2XFqUrjmcaDk870\\n dy8o5VMwblaC/mFhEFwlxGAjM6GqHHVlZ0AeXR+NJ7F3Fu858dP/zFfb2kSzpjgs\\n BuAQ2W5+fxtvlKHZlMfLE4OuQiME8JhvL/czzOVK9q2wGUzl/IkbyG3SwLoLFrTr\\n y2TsSibdfZe8ZrpS3gKaNqwSg1RVvShDlyvLfjm6nLgV7l2ifO9vHsaHjdmb9mUw\\n K/soRvNFAgMBAAECggEBAIOw0Jz899GjdsF43TO2NpqGj2pJuhPFZH0S7T/v+tRh\\n qERaoMLmhXTPQUuKQwar5UkJTL2nxhEtiWg9V5UjmsUarJAjHsKA7irZBs+HrTsg\\n aKguuQP6bN7k5yqEbgyKqLvpsIJrqKl+DtH/55A9JP79wlB1AnMxlwAwnNqhKut8\\n flwaeO6719IexNXb4WF9TZtKoW/XTqr0HOeiasaltJ/3tZBEXrIoC0Qvxa0534Nf\\n RfrBc2Kq86wFVl1P47JoaKv7obpZndAVrRWif1CVrtzQzIG+WhIfIhNPEYJniAWH\\n 8UtVhcGT3Eiiq8HfZEiqklNnG++oWKGfsx98mkb2SgECgYEA3Ohjxyhb8bFazC9h\\n hvF/GytKzbhCKfGM2cjD+OvGT17pEUma7QvLMA20Sm3f27e/Sw5NZtyGx9a3fGGf\\n dY0jEQ6y3R+4dhcf4ZfKkkBL4R8jA29wRiACRc3rpYFqYB+GsiI3A6IZHfnKUU2S\\n DTOjxQJJvuzM5Ppjs0OmLToXQDECgYEAzLdnC8FfAG3P4JpT0trrnxax7nBCo6d3\\n RKvz320W1dGtfECVe60DaaYB0W8AFEIsmakdWBsWNdCBaYxQ1spC1mNUyo7sIwme\\n OSyZ+Cxdg/r4e02zyFPtHKd7ar7o3CpkSQIcM0Zk/m2eMG6C7m0ChjJY8934NBYb\\n PhsN6BG7E1UCgYEAwdWDn3/xVVzan+k/OSnz7sII7AOuwqD5hysbkfJH2uMbvJiK\\n QU8k5bBQrzJDx8YuKsyM7CG6feUQsSnzwjCqQVBVb6NitvPJfKg1Dikuq4Unst74\\n c/+oHtn12A57aYagKPPOs/hq85t3g+l9qunR3I8KaGXdz1lJXEWSrYKYXjECgYB7\\n YQWf1hk1nvksOpbOe9aJ+RmfxNTE4UdGggPm4k5i644NVrdA5JMr9zsdSDLaAs/y\\n hDQFR73pDRMR09lcumXx48fUlLLIoyFTAAiDw+lQg8+CMOBrmflLzbzaJtkc6Aes\\n 4LKyTHjNxq8SLWiH3fcpfeqSf3L5oWEl2xRUi2seSQKBgQCvlZnes+BmikpJTUsY\\n ZeRym1Ojw/Dlmf90sHQRPOn+SXq3Nxi4Qye+TO0aLsey4+6S1CZ8q8iwzV8RZoLS\\n Uh7RZC32iKP+TdxeXlTpHKxfCfp5NvXRoH5BBnI2kcwvmEMKKToWYppv99/jmXQA\\n bT7gY3gln0TZI6WyJTxCZ7NJIg==\\n -----END PRIVATE KEY-----\\n "
	//
	//resp, err = svc.ConfigCertificateGet(&configCertificate)
	//if err != nil {
	//	fmt.Println("error:", err.Error())
	//}
	//if resp != nil {
	//	str, _ := json.Marshal(&resp)
	//	fmt.Printf("%+v\n", string(str))
	//}

	//	**************************************** 为单域名或者多域名配置证书，多域名用英文半角逗号隔开 POST（ConfigCertificate）***************************************************
	//  Parameters:
	//            Enable            String    开启、关闭设置服务证书，取值：on：开启，off：关闭，默认为off，当选择开启时，以下为必填
	//            DomainIds         String    domainid列表，英文半角逗号隔开
	//            CertificateId     String    金山云生成的证书唯一性ID，若输入证书ID，则以下内容可不填写，若无证书ID，则以下内容为必填
	//            CertificateName   String    证书名称
	//            ServerCertificate String    域名对应的证书内容
	//            PrivateKey        String    证书对应的私钥内容

	//configCertificate := make(map[string]interface{})
	//configCertificate["Enable"] = "on"
	//configCertificate["DomainIds"] = "2D09VN9"
	//configCertificate["ServerCertificate"] = "-----BEGIN CERTIFICATE-----\\n MIIDiTCCAnGgAwIBAgIJANF7PDOT9Wm8MA0GCSqGSIb3DQEBCwUAMFsxCzAJBgNV\\n BAYTAmFoMQswCQYDVQQIDAJiajELMAkGA1UEBwwCYmoxETAPBgNVBAoMCGtpbmdz\\n b2Z0MQowCAYDVQQLDAFzMRMwEQYDVQQDDAphYWEuY29tLmNuMB4XDTE3MDcxMzA4\\n MTA1NFoXDTE3MDgxMjA4MTA1NFowWzELMAkGA1UEBhMCYWgxCzAJBgNVBAgMAmJq\\n MQswCQYDVQQHDAJiajERMA8GA1UECgwIa2luZ3NvZnQxCjAIBgNVBAsMAXMxEzAR\\n BgNVBAMMCmFhYS5jb20uY24wggEiMA0GCSqGSIb3DQEBAQUAA4IBDwAwggEKAoIB\\n AQCwp3KNrIW288L93DGyuBQ1eVL2XVW2IsCW3SKuntiTtFxA5+xecKN1xoTbNrJp\\n oWjIHUow/DRyH2k0j1zHrWs7mHfx/1LR2G6PfCGUgvA2EDNaoDia3gIPmmt6QUVh\\n l2XFqUrjmcaDk870dy8o5VMwblaC/mFhEFwlxGAjM6GqHHVlZ0AeXR+NJ7F3Fu85\\n 8dP/zFfb2kSzpjgsBuAQ2W5+fxtvlKHZlMfLE4OuQiME8JhvL/czzOVK9q2wGUzl\\n /IkbyG3SwLoLFrTry2TsSibdfZe8ZrpS3gKaNqwSg1RVvShDlyvLfjm6nLgV7l2i\\n fO9vHsaHjdmb9mUwK/soRvNFAgMBAAGjUDBOMB0GA1UdDgQWBBTK3MSboh5vOtTF\\n /i7ly75c7uHPJjAfBgNVHSMEGDAWgBTK3MSboh5vOtTF/i7ly75c7uHPJjAMBgNV\\n HRMEBTADAQH/MA0GCSqGSIb3DQEBCwUAA4IBAQCtVHjp0C94iHqXS2E1e2gdP8eU\\n qg2naXxMRCXo2hSLEcPmNhv2Z8SfCpu54CInEyINxZg58z7KxNFD3AmNcj1/r67L\\n lVmkZzZrNsLT3Ti3SYiPgRUOSTiABZZrr0HRRbD+N32mLJ/2IeuBneBM6o0ER3Lp\\n HDD+hVcB15+MmC0uj68kWCbYcfd3go6XiPpmVyaPkTkhYQDHufW0SZ+dILvDNIVy\\n rAh7ZEUiP81LxYZcsf8fPKA/MmscQ4wZ2Owxt0KecgrkXcPTeOBdFUQAanmdsyLt\\n +U+EGOF31pUgw7FLg5Tb0ejhE1aOmjYybWCt6FTldPBHHEDCcQZjOEXn56YG\\n -----END CERTIFICATE-----\\n "
	//configCertificate["CertificateName"] = "2D09VN9_test1"
	//configCertificate["PrivateKey"] = "-----BEGIN PRIVATE KEY-----\\n MIIEvwIBADANBgkqhkiG9w0BAQEFAASCBKkwggSlAgEAAoIBAQCwp3KNrIW288L9\\n 3DGyuBQ1eVL2XVW2IsCW3SKuntiTtFxA5+xecKN1xoTbNrJpoWjIHUow/DRyH2k0\\n j1zHrWs7mHfx/1LR2G6PfCGUgvA2EDNaoDia3gIPmmt6QUVhl2XFqUrjmcaDk870\\n dy8o5VMwblaC/mFhEFwlxGAjM6GqHHVlZ0AeXR+NJ7F3Fu858dP/zFfb2kSzpjgs\\n BuAQ2W5+fxtvlKHZlMfLE4OuQiME8JhvL/czzOVK9q2wGUzl/IkbyG3SwLoLFrTr\\n y2TsSibdfZe8ZrpS3gKaNqwSg1RVvShDlyvLfjm6nLgV7l2ifO9vHsaHjdmb9mUw\\n K/soRvNFAgMBAAECggEBAIOw0Jz899GjdsF43TO2NpqGj2pJuhPFZH0S7T/v+tRh\\n qERaoMLmhXTPQUuKQwar5UkJTL2nxhEtiWg9V5UjmsUarJAjHsKA7irZBs+HrTsg\\n aKguuQP6bN7k5yqEbgyKqLvpsIJrqKl+DtH/55A9JP79wlB1AnMxlwAwnNqhKut8\\n flwaeO6719IexNXb4WF9TZtKoW/XTqr0HOeiasaltJ/3tZBEXrIoC0Qvxa0534Nf\\n RfrBc2Kq86wFVl1P47JoaKv7obpZndAVrRWif1CVrtzQzIG+WhIfIhNPEYJniAWH\\n 8UtVhcGT3Eiiq8HfZEiqklNnG++oWKGfsx98mkb2SgECgYEA3Ohjxyhb8bFazC9h\\n hvF/GytKzbhCKfGM2cjD+OvGT17pEUma7QvLMA20Sm3f27e/Sw5NZtyGx9a3fGGf\\n dY0jEQ6y3R+4dhcf4ZfKkkBL4R8jA29wRiACRc3rpYFqYB+GsiI3A6IZHfnKUU2S\\n DTOjxQJJvuzM5Ppjs0OmLToXQDECgYEAzLdnC8FfAG3P4JpT0trrnxax7nBCo6d3\\n RKvz320W1dGtfECVe60DaaYB0W8AFEIsmakdWBsWNdCBaYxQ1spC1mNUyo7sIwme\\n OSyZ+Cxdg/r4e02zyFPtHKd7ar7o3CpkSQIcM0Zk/m2eMG6C7m0ChjJY8934NBYb\\n PhsN6BG7E1UCgYEAwdWDn3/xVVzan+k/OSnz7sII7AOuwqD5hysbkfJH2uMbvJiK\\n QU8k5bBQrzJDx8YuKsyM7CG6feUQsSnzwjCqQVBVb6NitvPJfKg1Dikuq4Unst74\\n c/+oHtn12A57aYagKPPOs/hq85t3g+l9qunR3I8KaGXdz1lJXEWSrYKYXjECgYB7\\n YQWf1hk1nvksOpbOe9aJ+RmfxNTE4UdGggPm4k5i644NVrdA5JMr9zsdSDLaAs/y\\n hDQFR73pDRMR09lcumXx48fUlLLIoyFTAAiDw+lQg8+CMOBrmflLzbzaJtkc6Aes\\n 4LKyTHjNxq8SLWiH3fcpfeqSf3L5oWEl2xRUi2seSQKBgQCvlZnes+BmikpJTUsY\\n ZeRym1Ojw/Dlmf90sHQRPOn+SXq3Nxi4Qye+TO0aLsey4+6S1CZ8q8iwzV8RZoLS\\n Uh7RZC32iKP+TdxeXlTpHKxfCfp5NvXRoH5BBnI2kcwvmEMKKToWYppv99/jmXQA\\n bT7gY3gln0TZI6WyJTxCZ7NJIg==\\n -----END PRIVATE KEY-----\\n "
	//
	//resp, err = svc.ConfigCertificatePost(&configCertificate)
	//if err != nil {
	//	fmt.Println("error:", err.Error())
	//}
	//if resp != nil {
	//	str, _ := json.Marshal(&resp)
	//	fmt.Printf("%+v\n", string(str))
	//}

	//	**************************************** 更新证书信息 GET（SetCertificate）***************************************************
	//  Parameters:
	//            CertificateId     String  证书对应的唯一ID
	//            CertificateName   String  安全证书名称
	//            ServerCertificate String  域名对应的安全证书内容
	//            PrivateKey        String  安全证书对应的私钥内容

	//setCertificate := make(map[string]interface{})
	//setCertificate["CertificateId"] = "917"
	//setCertificate["CertificateName"] = "2D09VN9_test1"
	//setCertificate["ServerCertificate"] = "-----BEGIN CERTIFICATE-----\\n MIIDiTCCAnGgAwIBAgIJANF7PDOT9Wm8MA0GCSqGSIb3DQEBCwUAMFsxCzAJBgNV\\n BAYTAmFoMQswCQYDVQQIDAJiajELMAkGA1UEBwwCYmoxETAPBgNVBAoMCGtpbmdz\\n b2Z0MQowCAYDVQQLDAFzMRMwEQYDVQQDDAphYWEuY29tLmNuMB4XDTE3MDcxMzA4\\n MTA1NFoXDTE3MDgxMjA4MTA1NFowWzELMAkGA1UEBhMCYWgxCzAJBgNVBAgMAmJq\\n MQswCQYDVQQHDAJiajERMA8GA1UECgwIa2luZ3NvZnQxCjAIBgNVBAsMAXMxEzAR\\n BgNVBAMMCmFhYS5jb20uY24wggEiMA0GCSqGSIb3DQEBAQUAA4IBDwAwggEKAoIB\\n AQCwp3KNrIW288L93DGyuBQ1eVL2XVW2IsCW3SKuntiTtFxA5+xecKN1xoTbNrJp\\n oWjIHUow/DRyH2k0j1zHrWs7mHfx/1LR2G6PfCGUgvA2EDNaoDia3gIPmmt6QUVh\\n l2XFqUrjmcaDk870dy8o5VMwblaC/mFhEFwlxGAjM6GqHHVlZ0AeXR+NJ7F3Fu85\\n 8dP/zFfb2kSzpjgsBuAQ2W5+fxtvlKHZlMfLE4OuQiME8JhvL/czzOVK9q2wGUzl\\n /IkbyG3SwLoLFrTry2TsSibdfZe8ZrpS3gKaNqwSg1RVvShDlyvLfjm6nLgV7l2i\\n fO9vHsaHjdmb9mUwK/soRvNFAgMBAAGjUDBOMB0GA1UdDgQWBBTK3MSboh5vOtTF\\n /i7ly75c7uHPJjAfBgNVHSMEGDAWgBTK3MSboh5vOtTF/i7ly75c7uHPJjAMBgNV\\n HRMEBTADAQH/MA0GCSqGSIb3DQEBCwUAA4IBAQCtVHjp0C94iHqXS2E1e2gdP8eU\\n qg2naXxMRCXo2hSLEcPmNhv2Z8SfCpu54CInEyINxZg58z7KxNFD3AmNcj1/r67L\\n lVmkZzZrNsLT3Ti3SYiPgRUOSTiABZZrr0HRRbD+N32mLJ/2IeuBneBM6o0ER3Lp\\n HDD+hVcB15+MmC0uj68kWCbYcfd3go6XiPpmVyaPkTkhYQDHufW0SZ+dILvDNIVy\\n rAh7ZEUiP81LxYZcsf8fPKA/MmscQ4wZ2Owxt0KecgrkXcPTeOBdFUQAanmdsyLt\\n +U+EGOF31pUgw7FLg5Tb0ejhE1aOmjYybWCt6FTldPBHHEDCcQZjOEXn56YG\\n -----END CERTIFICATE-----\\n "
	//setCertificate["PrivateKey"] = "-----BEGIN PRIVATE KEY-----\\n MIIEvwIBADANBgkqhkiG9w0BAQEFAASCBKkwggSlAgEAAoIBAQCwp3KNrIW288L9\\n 3DGyuBQ1eVL2XVW2IsCW3SKuntiTtFxA5+xecKN1xoTbNrJpoWjIHUow/DRyH2k0\\n j1zHrWs7mHfx/1LR2G6PfCGUgvA2EDNaoDia3gIPmmt6QUVhl2XFqUrjmcaDk870\\n dy8o5VMwblaC/mFhEFwlxGAjM6GqHHVlZ0AeXR+NJ7F3Fu858dP/zFfb2kSzpjgs\\n BuAQ2W5+fxtvlKHZlMfLE4OuQiME8JhvL/czzOVK9q2wGUzl/IkbyG3SwLoLFrTr\\n y2TsSibdfZe8ZrpS3gKaNqwSg1RVvShDlyvLfjm6nLgV7l2ifO9vHsaHjdmb9mUw\\n K/soRvNFAgMBAAECggEBAIOw0Jz899GjdsF43TO2NpqGj2pJuhPFZH0S7T/v+tRh\\n qERaoMLmhXTPQUuKQwar5UkJTL2nxhEtiWg9V5UjmsUarJAjHsKA7irZBs+HrTsg\\n aKguuQP6bN7k5yqEbgyKqLvpsIJrqKl+DtH/55A9JP79wlB1AnMxlwAwnNqhKut8\\n flwaeO6719IexNXb4WF9TZtKoW/XTqr0HOeiasaltJ/3tZBEXrIoC0Qvxa0534Nf\\n RfrBc2Kq86wFVl1P47JoaKv7obpZndAVrRWif1CVrtzQzIG+WhIfIhNPEYJniAWH\\n 8UtVhcGT3Eiiq8HfZEiqklNnG++oWKGfsx98mkb2SgECgYEA3Ohjxyhb8bFazC9h\\n hvF/GytKzbhCKfGM2cjD+OvGT17pEUma7QvLMA20Sm3f27e/Sw5NZtyGx9a3fGGf\\n dY0jEQ6y3R+4dhcf4ZfKkkBL4R8jA29wRiACRc3rpYFqYB+GsiI3A6IZHfnKUU2S\\n DTOjxQJJvuzM5Ppjs0OmLToXQDECgYEAzLdnC8FfAG3P4JpT0trrnxax7nBCo6d3\\n RKvz320W1dGtfECVe60DaaYB0W8AFEIsmakdWBsWNdCBaYxQ1spC1mNUyo7sIwme\\n OSyZ+Cxdg/r4e02zyFPtHKd7ar7o3CpkSQIcM0Zk/m2eMG6C7m0ChjJY8934NBYb\\n PhsN6BG7E1UCgYEAwdWDn3/xVVzan+k/OSnz7sII7AOuwqD5hysbkfJH2uMbvJiK\\n QU8k5bBQrzJDx8YuKsyM7CG6feUQsSnzwjCqQVBVb6NitvPJfKg1Dikuq4Unst74\\n c/+oHtn12A57aYagKPPOs/hq85t3g+l9qunR3I8KaGXdz1lJXEWSrYKYXjECgYB7\\n YQWf1hk1nvksOpbOe9aJ+RmfxNTE4UdGggPm4k5i644NVrdA5JMr9zsdSDLaAs/y\\n hDQFR73pDRMR09lcumXx48fUlLLIoyFTAAiDw+lQg8+CMOBrmflLzbzaJtkc6Aes\\n 4LKyTHjNxq8SLWiH3fcpfeqSf3L5oWEl2xRUi2seSQKBgQCvlZnes+BmikpJTUsY\\n ZeRym1Ojw/Dlmf90sHQRPOn+SXq3Nxi4Qye+TO0aLsey4+6S1CZ8q8iwzV8RZoLS\\n Uh7RZC32iKP+TdxeXlTpHKxfCfp5NvXRoH5BBnI2kcwvmEMKKToWYppv99/jmXQA\\n bT7gY3gln0TZI6WyJTxCZ7NJIg==\\n -----END PRIVATE KEY-----\\n "
	//
	//resp, err = svc.SetCertificateGet(&setCertificate)
	//if err != nil {
	//	fmt.Println("error:", err.Error())
	//}
	//if resp != nil {
	//	str, _ := json.Marshal(&resp)
	//	fmt.Printf("%+v\n", string(str))
	//}

	//	**************************************** 更新证书信息 POST（SetCertificate）***************************************************
	//  Parameters:
	//            CertificateId     String  证书对应的唯一ID
	//            CertificateName   String  安全证书名称
	//            ServerCertificate String  域名对应的安全证书内容
	//            PrivateKey        String  安全证书对应的私钥内容

	//setCertificate := make(map[string]interface{})
	//setCertificate["CertificateId"] = "917"
	//setCertificate["CertificateName"] = "2D09VN9_test1"
	//setCertificate["ServerCertificate"] = "-----BEGIN CERTIFICATE-----\\n MIIDiTCCAnGgAwIBAgIJANF7PDOT9Wm8MA0GCSqGSIb3DQEBCwUAMFsxCzAJBgNV\\n BAYTAmFoMQswCQYDVQQIDAJiajELMAkGA1UEBwwCYmoxETAPBgNVBAoMCGtpbmdz\\n b2Z0MQowCAYDVQQLDAFzMRMwEQYDVQQDDAphYWEuY29tLmNuMB4XDTE3MDcxMzA4\\n MTA1NFoXDTE3MDgxMjA4MTA1NFowWzELMAkGA1UEBhMCYWgxCzAJBgNVBAgMAmJq\\n MQswCQYDVQQHDAJiajERMA8GA1UECgwIa2luZ3NvZnQxCjAIBgNVBAsMAXMxEzAR\\n BgNVBAMMCmFhYS5jb20uY24wggEiMA0GCSqGSIb3DQEBAQUAA4IBDwAwggEKAoIB\\n AQCwp3KNrIW288L93DGyuBQ1eVL2XVW2IsCW3SKuntiTtFxA5+xecKN1xoTbNrJp\\n oWjIHUow/DRyH2k0j1zHrWs7mHfx/1LR2G6PfCGUgvA2EDNaoDia3gIPmmt6QUVh\\n l2XFqUrjmcaDk870dy8o5VMwblaC/mFhEFwlxGAjM6GqHHVlZ0AeXR+NJ7F3Fu85\\n 8dP/zFfb2kSzpjgsBuAQ2W5+fxtvlKHZlMfLE4OuQiME8JhvL/czzOVK9q2wGUzl\\n /IkbyG3SwLoLFrTry2TsSibdfZe8ZrpS3gKaNqwSg1RVvShDlyvLfjm6nLgV7l2i\\n fO9vHsaHjdmb9mUwK/soRvNFAgMBAAGjUDBOMB0GA1UdDgQWBBTK3MSboh5vOtTF\\n /i7ly75c7uHPJjAfBgNVHSMEGDAWgBTK3MSboh5vOtTF/i7ly75c7uHPJjAMBgNV\\n HRMEBTADAQH/MA0GCSqGSIb3DQEBCwUAA4IBAQCtVHjp0C94iHqXS2E1e2gdP8eU\\n qg2naXxMRCXo2hSLEcPmNhv2Z8SfCpu54CInEyINxZg58z7KxNFD3AmNcj1/r67L\\n lVmkZzZrNsLT3Ti3SYiPgRUOSTiABZZrr0HRRbD+N32mLJ/2IeuBneBM6o0ER3Lp\\n HDD+hVcB15+MmC0uj68kWCbYcfd3go6XiPpmVyaPkTkhYQDHufW0SZ+dILvDNIVy\\n rAh7ZEUiP81LxYZcsf8fPKA/MmscQ4wZ2Owxt0KecgrkXcPTeOBdFUQAanmdsyLt\\n +U+EGOF31pUgw7FLg5Tb0ejhE1aOmjYybWCt6FTldPBHHEDCcQZjOEXn56YG\\n -----END CERTIFICATE-----\\n "
	//setCertificate["PrivateKey"] = "-----BEGIN PRIVATE KEY-----\\n MIIEvwIBADANBgkqhkiG9w0BAQEFAASCBKkwggSlAgEAAoIBAQCwp3KNrIW288L9\\n 3DGyuBQ1eVL2XVW2IsCW3SKuntiTtFxA5+xecKN1xoTbNrJpoWjIHUow/DRyH2k0\\n j1zHrWs7mHfx/1LR2G6PfCGUgvA2EDNaoDia3gIPmmt6QUVhl2XFqUrjmcaDk870\\n dy8o5VMwblaC/mFhEFwlxGAjM6GqHHVlZ0AeXR+NJ7F3Fu858dP/zFfb2kSzpjgs\\n BuAQ2W5+fxtvlKHZlMfLE4OuQiME8JhvL/czzOVK9q2wGUzl/IkbyG3SwLoLFrTr\\n y2TsSibdfZe8ZrpS3gKaNqwSg1RVvShDlyvLfjm6nLgV7l2ifO9vHsaHjdmb9mUw\\n K/soRvNFAgMBAAECggEBAIOw0Jz899GjdsF43TO2NpqGj2pJuhPFZH0S7T/v+tRh\\n qERaoMLmhXTPQUuKQwar5UkJTL2nxhEtiWg9V5UjmsUarJAjHsKA7irZBs+HrTsg\\n aKguuQP6bN7k5yqEbgyKqLvpsIJrqKl+DtH/55A9JP79wlB1AnMxlwAwnNqhKut8\\n flwaeO6719IexNXb4WF9TZtKoW/XTqr0HOeiasaltJ/3tZBEXrIoC0Qvxa0534Nf\\n RfrBc2Kq86wFVl1P47JoaKv7obpZndAVrRWif1CVrtzQzIG+WhIfIhNPEYJniAWH\\n 8UtVhcGT3Eiiq8HfZEiqklNnG++oWKGfsx98mkb2SgECgYEA3Ohjxyhb8bFazC9h\\n hvF/GytKzbhCKfGM2cjD+OvGT17pEUma7QvLMA20Sm3f27e/Sw5NZtyGx9a3fGGf\\n dY0jEQ6y3R+4dhcf4ZfKkkBL4R8jA29wRiACRc3rpYFqYB+GsiI3A6IZHfnKUU2S\\n DTOjxQJJvuzM5Ppjs0OmLToXQDECgYEAzLdnC8FfAG3P4JpT0trrnxax7nBCo6d3\\n RKvz320W1dGtfECVe60DaaYB0W8AFEIsmakdWBsWNdCBaYxQ1spC1mNUyo7sIwme\\n OSyZ+Cxdg/r4e02zyFPtHKd7ar7o3CpkSQIcM0Zk/m2eMG6C7m0ChjJY8934NBYb\\n PhsN6BG7E1UCgYEAwdWDn3/xVVzan+k/OSnz7sII7AOuwqD5hysbkfJH2uMbvJiK\\n QU8k5bBQrzJDx8YuKsyM7CG6feUQsSnzwjCqQVBVb6NitvPJfKg1Dikuq4Unst74\\n c/+oHtn12A57aYagKPPOs/hq85t3g+l9qunR3I8KaGXdz1lJXEWSrYKYXjECgYB7\\n YQWf1hk1nvksOpbOe9aJ+RmfxNTE4UdGggPm4k5i644NVrdA5JMr9zsdSDLaAs/y\\n hDQFR73pDRMR09lcumXx48fUlLLIoyFTAAiDw+lQg8+CMOBrmflLzbzaJtkc6Aes\\n 4LKyTHjNxq8SLWiH3fcpfeqSf3L5oWEl2xRUi2seSQKBgQCvlZnes+BmikpJTUsY\\n ZeRym1Ojw/Dlmf90sHQRPOn+SXq3Nxi4Qye+TO0aLsey4+6S1CZ8q8iwzV8RZoLS\\n Uh7RZC32iKP+TdxeXlTpHKxfCfp5NvXRoH5BBnI2kcwvmEMKKToWYppv99/jmXQA\\n bT7gY3gln0TZI6WyJTxCZ7NJIg==\\n -----END PRIVATE KEY-----\\n "
	//
	//resp, err = svc.SetCertificatePost(&setCertificate)
	//if err != nil {
	//	fmt.Println("error:", err.Error())
	//}
	//if resp != nil {
	//	str, _ := json.Marshal(&resp)
	//	fmt.Printf("%+v\n", string(str))
	//}

	//	**************************************** 批量删除证书列表 GET（RemoveCertificates）***************************************************
	//  Parameters:
	//            CertificateIds   String  证书id列表，英文半角逗号隔开

	//removeCertificates := make(map[string]interface{})
	//removeCertificates["CertificateIds"] = "910,911"
	//
	//resp, err = svc.RemoveCertificatesGet(&removeCertificates)
	//if err != nil {
	//	fmt.Println("error:", err.Error())
	//}
	//if resp != nil {
	//	str, _ := json.Marshal(&resp)
	//	fmt.Printf("%+v\n", string(str))
	//}

	//	**************************************** 批量删除证书列表 POST（RemoveCertificates）***************************************************
	//  Parameters:
	//            CertificateIds   String  证书id列表，英文半角逗号隔开

	//removeCertificates := make(map[string]interface{})
	//removeCertificates["CertificateIds"] = "910,911"
	//
	//resp, err = svc.RemoveCertificatesPost(&removeCertificates)
	//if err != nil {
	//	fmt.Println("error:", err.Error())
	//}
	//if resp != nil {
	//	str, _ := json.Marshal(&resp)
	//	fmt.Printf("%+v\n", string(str))
	//}

	//	**************************************** 分页获得用户证书列表参数 GET（GetCertificates）***************************************************
	//  Parameters:
	//            PageSize        Long    一页展示条数
	//            PageNum         Long    当前页码

	//getCertificates := make(map[string]interface{})
	//getCertificates["PageSize"] = 1
	//getCertificates["PageNum"] = 1
	//
	//resp, err = svc.GetCertificatesGet(&getCertificates)
	//if err != nil {
	//	fmt.Println("error:", err.Error())
	//}
	//if resp != nil {
	//	str, _ := json.Marshal(&resp)
	//	fmt.Printf("%+v\n", string(str))
	//}

	//	**************************************** 分页获得用户证书列表参数 POST（GetCertificates）***************************************************
	//  Parameters:
	//            PageSize        Long    一页展示条数
	//            PageNum         Long    当前页码

	//getCertificates := make(map[string]interface{})
	//getCertificates["PageSize"] = 1
	//getCertificates["PageNum"] = 2
	//
	//resp, err = svc.GetCertificatesPost(&getCertificates)
	//if err != nil {
	//	fmt.Println("error:", err.Error())
	//}
	//if resp != nil {
	//	str, _ := json.Marshal(&resp)
	//	fmt.Printf("%+v\n", string(str))
	//}

	//	**************************************** 更新对加速域名批量修改的配置项，并且支持精确域名和泛域名对应的修改 POST（SetDomainConfigs）***************************************************
	//******  Parameters:
	//******		DomainId	String	是	域名ID
	//******		IgnoreQueryStringConfig	IgnoreQueryStringConfig	否	表示设置过滤参数
	//******		BackOriginHostConfig	BackOriginHostConfig	否	表示设置回源host
	//******		ReferProtectionConfig	ReferProtectionConfig	否	表示设置refer防盗链
	//******		CacheRuleConfig	CacheRuleConfig	否	表示设置缓存策略
	//******		IpProtectionConfig	IpProtectionConfig	否	表示设置IP防盗链

	//cacheRuleDir := make(map[string]interface{})
	//cacheRuleDir["CacheRuleType"]="directory"
	//cacheRuleDir["Value"]="/xxxx/"
	//cacheRuleDir["CacheTime"]=11
	//cacheRuleDir["RespectOrigin"]="off"
	//
	//cacheRuleExact := make(map[string]interface{})
	//cacheRuleExact["CacheRuleType"] = "exact"
	//cacheRuleExact["Value"] = "/sdfsf/sdf.text"
	//cacheRuleExact["CacheTime"] = 120
	//cacheRuleExact["RespectOrigin"] = "off"
	//
	//cacheRules := make([]map[string]interface{}, 0)
	//cacheRules = append(cacheRules, cacheRuleDir, cacheRuleExact)
	//
	//cacheRuleConfig := make(map[string]interface{})
	//cacheRuleConfig["CacheRules"] = cacheRules
	//
	//ignoreQueryStringConfig := make(map[string]interface{})
	//ignoreQueryStringConfig["Enable"] = "on"
	//
	//referProtectionConfig := make(map[string]interface{})
	//referProtectionConfig["Enable"] = "on"
	//referProtectionConfig["ReferType"] = "block"
	//referProtectionConfig["ReferList"] = "www.baidu.com,www.sina.com"
	//referProtectionConfig["AllowEmpty"] = "on"
	//
	//backOriginHostConfig := make(map[string]interface{})
	//backOriginHostConfig["BackOriginHost"] = "www.a.qunar.com"
	//
	//ipProtectionConfig := make(map[string]interface{})
	//ipProtectionConfig["Enable"] = "on"
	//ipProtectionConfig["IpType"] = "allow"
	//ipProtectionConfig["IpList"] = "100.101.111.15"
	//
	//setDomainConfigs := make(map[string]interface{})
	//setDomainConfigs["DomainId"] = "2D07WRH"
	//setDomainConfigs["IgnoreQueryStringConfig"] = ignoreQueryStringConfig
	//setDomainConfigs["BackOriginHostConfig"] = backOriginHostConfig
	//setDomainConfigs["ReferProtectionConfig"] = referProtectionConfig
	//setDomainConfigs["CacheRuleConfig"] = cacheRuleConfig
	//setDomainConfigs["IpProtectionConfig"] = ipProtectionConfig
	//
	//resp, err = svc.SetDomainConfigsPost(&setDomainConfigs)
	//if err != nil {
	//	fmt.Println("error:", err.Error())
	//}
	//if resp != nil {
	//	str, _ := json.Marshal(&resp)
	//	fmt.Printf("%+v\n", string(str))
	//}

	//	**************************************** 获取泛域名次级域名带宽数据，包括边缘带宽、回源带宽数据 GET（GetSubDomainsBandwidthData）***************************************************
	//  Parameters:
	//      DomainId	是	String	表示一个泛域名
	//		Domains	是	String	表示泛域名的次级域名，但查询次级域名的个数≤100个
	//		StartTime	是	String	获取数据起始时间点，日期格式按ISO8601表示法，北京时间，格式为：YYYY-MM-DDThh:mm+0800，例如： 2016-08-01T21:14+0800
	//		EndTime	是	String	结束时间需大于起始时间；获取日期格式按照ISO8601表示法，北京时间，格式为：YYYY-MM-DDThh:mm+0800，例如： 2016-08-01T21:14+0800
	//		Regions	否	String	区域名称， 取值为CN:中国大陆，HK：香港，TW：台湾，AS：亚洲其他，NA：北美洲，SA：南美洲，EU：欧洲，AU：大洋洲，AF：非洲，支持多区域查询，多个区域用逗号（半角）分隔，缺省为 CN
	//		ResultType	是	Long	取值为0：多域名多区域数据做合并；1：每个域名每个区域的数据分别返回。
	//		Granularity	否	Long	统计粒度，取值为 5（默认）：5分钟粒度；10：10分钟粒度；20：20分钟粒度；60：1小时粒度；240：4小时粒度；480：8小时粒度；1440：1天粒度；以上粒度的带宽值均取该粒度时间段的峰值
	//		DataType	否	String	数据类型， 取值为edge:边缘数据; origin:回源数据; 支持多类型选择，多个类型用逗号（半角）分隔，缺省为 edge
	//		ProtocolType	否	String	协议类型， 取值为http:http协议数据; https:https协议数据

	//getSubDomainsBandwidthData := make(map[string]interface{})
	//getSubDomainsBandwidthData["DomainId"] = "2D07WRH"
	//getSubDomainsBandwidthData["Domains"] = ""
	//getSubDomainsBandwidthData["StartTime"] = "2021-06-01T21:14+0800"
	//getSubDomainsBandwidthData["EndTime"] = "2021-06-01T21:34+0800"
	//getSubDomainsBandwidthData["Regions"] = "CN"
	//getSubDomainsBandwidthData["ResultType"] = 0
	//getSubDomainsBandwidthData["Granularity"] = 5
	//getSubDomainsBandwidthData["DataType"] = "edge"
	//getSubDomainsBandwidthData["ProtocolType"] = "http"
	//
	//resp, err = svc.GetSubDomainsBandwidthDataGet(&getSubDomainsBandwidthData)
	//if err != nil {
	//	fmt.Println("error:", err.Error())
	//}
	//if resp != nil {
	//	str, _ := json.Marshal(&resp)
	//	fmt.Printf("%+v\n", string(str))
	//}

	//	**************************************** 获取泛域名次级域名带宽数据，包括边缘带宽、回源带宽数据 POST（GetSubDomainsBandwidthData）***************************************************
	//  Parameters:
	//      DomainId	是	String	表示一个泛域名
	//		Domains	是	String	表示泛域名的次级域名，但查询次级域名的个数≤100个
	//		StartTime	是	String	获取数据起始时间点，日期格式按ISO8601表示法，北京时间，格式为：YYYY-MM-DDThh:mm+0800，例如： 2016-08-01T21:14+0800
	//		EndTime	是	String	结束时间需大于起始时间；获取日期格式按照ISO8601表示法，北京时间，格式为：YYYY-MM-DDThh:mm+0800，例如： 2016-08-01T21:14+0800
	//		Regions	否	String	区域名称， 取值为CN:中国大陆，HK：香港，TW：台湾，AS：亚洲其他，NA：北美洲，SA：南美洲，EU：欧洲，AU：大洋洲，AF：非洲，支持多区域查询，多个区域用逗号（半角）分隔，缺省为 CN
	//		ResultType	是	Long	取值为0：多域名多区域数据做合并；1：每个域名每个区域的数据分别返回。
	//		Granularity	否	Long	统计粒度，取值为 5（默认）：5分钟粒度；10：10分钟粒度；20：20分钟粒度；60：1小时粒度；240：4小时粒度；480：8小时粒度；1440：1天粒度；以上粒度的带宽值均取该粒度时间段的峰值
	//		DataType	否	String	数据类型， 取值为edge:边缘数据; origin:回源数据; 支持多类型选择，多个类型用逗号（半角）分隔，缺省为 edge
	//		ProtocolType	否	String	协议类型， 取值为http:http协议数据; https:https协议数据

	//getSubDomainsBandwidthData := make(map[string]interface{})
	//getSubDomainsBandwidthData["DomainId"] = "2D07WRH"
	//getSubDomainsBandwidthData["Domains"] = ""
	//getSubDomainsBandwidthData["StartTime"] = "2021-06-01T21:14+0800"
	//getSubDomainsBandwidthData["EndTime"] = "2021-06-01T21:34+0800"
	//getSubDomainsBandwidthData["Regions"] = "CN"
	//getSubDomainsBandwidthData["ResultType"] = 0
	//getSubDomainsBandwidthData["Granularity"] = 5
	//getSubDomainsBandwidthData["DataType"] = "edge"
	//getSubDomainsBandwidthData["ProtocolType"] = "http"
	//
	//resp, err = svc.GetSubDomainsBandwidthDataPost(&getSubDomainsBandwidthData)
	//if err != nil {
	//	fmt.Println("error:", err.Error())
	//}
	//if resp != nil {
	//	str, _ := json.Marshal(&resp)
	//	fmt.Printf("%+v\n", string(str))
	//}

	//	**************************************** 获取泛域名次级域名流量数据，包括边缘流量、回源流量数据 GET（GetSubDomainsFlowData）***************************************************
	//  Parameters:
	//      DomainId	是	String	表示一个泛域名
	//		Domains	是	String	表示泛域名的次级域名，但查询次级域名的个数≤100个
	//		StartTime	是	String	获取数据起始时间点，日期格式按ISO8601表示法，北京时间，格式为：YYYY-MM-DDThh:mm+0800，例如： 2016-08-01T21:14+0800
	//		EndTime	是	String	结束时间需大于起始时间；获取日期格式按照ISO8601表示法，北京时间，格式为：YYYY-MM-DDThh:mm+0800，例如： 2016-08-01T21:14+0800
	//		Regions	否	String	区域名称， 取值为CN:中国大陆，HK：香港，TW：台湾，AS：亚洲其他，NA：北美洲，SA：南美洲，EU：欧洲，AU：大洋洲，AF：非洲，支持多区域查询，多个区域用逗号（半角）分隔，缺省为 CN
	//		ResultType	是	Long	取值为0：多域名多区域数据做合并；1：每个域名每个区域的数据分别返回。
	//		Granularity	否	Long	统计粒度，取值为 5（默认）：5分钟粒度；10：10分钟粒度；20：20分钟粒度；60：1小时粒度；240：4小时粒度；480：8小时粒度；1440：1天粒度；以上粒度的带宽值均取该粒度时间段的峰值
	//		DataType	否	String	数据类型， 取值为edge:边缘数据; origin:回源数据; 支持多类型选择，多个类型用逗号（半角）分隔，缺省为 edge
	//		ProtocolType	否	String	协议类型， 取值为http:http协议数据; https:https协议数据

	//getSubDomainsFlowData := make(map[string]interface{})
	//getSubDomainsFlowData["DomainId"] = "2D07WRH"
	//getSubDomainsFlowData["Domains"] = ""
	//getSubDomainsFlowData["StartTime"] = "2021-06-01T21:14+0800"
	//getSubDomainsFlowData["EndTime"] = "2021-06-01T21:34+0800"
	//getSubDomainsFlowData["Regions"] = "CN"
	//getSubDomainsFlowData["ResultType"] = 0
	//getSubDomainsFlowData["Granularity"] = 5
	//getSubDomainsFlowData["DataType"] = "edge"
	//getSubDomainsFlowData["ProtocolType"] = "http"
	//
	//resp, err = svc.GetSubDomainsFlowDataGet(&getSubDomainsFlowData)
	//if err != nil {
	//	fmt.Println("error:", err.Error())
	//}
	//if resp != nil {
	//	str, _ := json.Marshal(&resp)
	//	fmt.Printf("%+v\n", string(str))
	//}

	//	****************************************获取泛域名次级域名流量数据，包括边缘流量、回源流量数据 POST（GetSubDomainsFlowData）***************************************************
	//  Parameters:
	//      DomainId	是	String	表示一个泛域名
	//		Domains	是	String	表示泛域名的次级域名，但查询次级域名的个数≤100个
	//		StartTime	是	String	获取数据起始时间点，日期格式按ISO8601表示法，北京时间，格式为：YYYY-MM-DDThh:mm+0800，例如： 2016-08-01T21:14+0800
	//		EndTime	是	String	结束时间需大于起始时间；获取日期格式按照ISO8601表示法，北京时间，格式为：YYYY-MM-DDThh:mm+0800，例如： 2016-08-01T21:14+0800
	//		Regions	否	String	区域名称， 取值为CN:中国大陆，HK：香港，TW：台湾，AS：亚洲其他，NA：北美洲，SA：南美洲，EU：欧洲，AU：大洋洲，AF：非洲，支持多区域查询，多个区域用逗号（半角）分隔，缺省为 CN
	//		ResultType	是	Long	取值为0：多域名多区域数据做合并；1：每个域名每个区域的数据分别返回。
	//		Granularity	否	Long	统计粒度，取值为 5（默认）：5分钟粒度；10：10分钟粒度；20：20分钟粒度；60：1小时粒度；240：4小时粒度；480：8小时粒度；1440：1天粒度；以上粒度的带宽值均取该粒度时间段的峰值
	//		DataType	否	String	数据类型， 取值为edge:边缘数据; origin:回源数据; 支持多类型选择，多个类型用逗号（半角）分隔，缺省为 edge
	//		ProtocolType	否	String	协议类型， 取值为http:http协议数据; https:https协议数据

	//getSubDomainsFlowData := make(map[string]interface{})
	//getSubDomainsFlowData["DomainId"] = "2D07WRH"
	//getSubDomainsFlowData["Domains"] = ""
	//getSubDomainsFlowData["StartTime"] = "2021-06-01T21:14+0800"
	//getSubDomainsFlowData["EndTime"] = "2021-06-01T21:34+0800"
	//getSubDomainsFlowData["Regions"] = "CN"
	//getSubDomainsFlowData["ResultType"] = 0
	//getSubDomainsFlowData["Granularity"] = 5
	//getSubDomainsFlowData["DataType"] = "edge"
	//getSubDomainsFlowData["ProtocolType"] = "http"
	//
	//resp, err = svc.GetSubDomainsFlowDataPost(&getSubDomainsFlowData)
	//if err != nil {
	//	fmt.Println("error:", err.Error())
	//}
	//if resp != nil {
	//	str, _ := json.Marshal(&resp)
	//	fmt.Printf("%+v\n", string(str))
	//}

	//	**************************************** 获取泛域名次级域名请求数数据，包括边缘请求数、回源请求数数据 GET（GetSubDomainsPvData）***************************************************
	//  Parameters:
	//      DomainId	是	String	表示一个泛域名
	//		Domains	是	String	表示泛域名的次级域名，但查询次级域名的个数≤100个
	//		StartTime	是	String	获取数据起始时间点，日期格式按ISO8601表示法，北京时间，格式为：YYYY-MM-DDThh:mm+0800，例如： 2016-08-01T21:14+0800
	//		EndTime	是	String	结束时间需大于起始时间；获取日期格式按照ISO8601表示法，北京时间，格式为：YYYY-MM-DDThh:mm+0800，例如： 2016-08-01T21:14+0800
	//		Regions	否	String	区域名称， 取值为CN:中国大陆，HK：香港，TW：台湾，AS：亚洲其他，NA：北美洲，SA：南美洲，EU：欧洲，AU：大洋洲，AF：非洲，支持多区域查询，多个区域用逗号（半角）分隔，缺省为 CN
	//		ResultType	是	Long	取值为0：多域名多区域数据做合并；1：每个域名每个区域的数据分别返回。
	//		Granularity	否	Long	统计粒度，取值为 5（默认）：5分钟粒度；10：10分钟粒度；20：20分钟粒度；60：1小时粒度；240：4小时粒度；480：8小时粒度；1440：1天粒度；以上粒度的带宽值均取该粒度时间段的峰值
	//		DataType	否	String	数据类型， 取值为edge:边缘数据; origin:回源数据; 支持多类型选择，多个类型用逗号（半角）分隔，缺省为 edge
	//		ProtocolType	否	String	协议类型， 取值为http:http协议数据; https:https协议数据

	//getSubDomainsPvData := make(map[string]interface{})
	//getSubDomainsPvData["DomainId"] = "2D07WRH"
	//getSubDomainsPvData["Domains"] = ""
	//getSubDomainsPvData["StartTime"] = "2021-06-01T21:14+0800"
	//getSubDomainsPvData["EndTime"] = "2021-06-01T21:34+0800"
	//getSubDomainsPvData["Regions"] = "CN"
	//getSubDomainsPvData["ResultType"] = 0
	//getSubDomainsPvData["Granularity"] = 5
	//getSubDomainsPvData["DataType"] = "edge"
	//getSubDomainsPvData["ProtocolType"] = "http"
	//
	//resp, err = svc.GetSubDomainsPvDataGet(&getSubDomainsPvData)
	//if err != nil {
	//	fmt.Println("error:", err.Error())
	//}
	//if resp != nil {
	//	str, _ := json.Marshal(&resp)
	//	fmt.Printf("%+v\n", string(str))
	//}

	//	**************************************** 获取泛域名次级域名请求数数据，包括边缘请求数、回源请求数数据 POST（GetSubDomainsPvData）***************************************************
	//  Parameters:
	//      DomainId	是	String	表示一个泛域名
	//		Domains	是	String	表示泛域名的次级域名，但查询次级域名的个数≤100个
	//		StartTime	是	String	获取数据起始时间点，日期格式按ISO8601表示法，北京时间，格式为：YYYY-MM-DDThh:mm+0800，例如： 2016-08-01T21:14+0800
	//		EndTime	是	String	结束时间需大于起始时间；获取日期格式按照ISO8601表示法，北京时间，格式为：YYYY-MM-DDThh:mm+0800，例如： 2016-08-01T21:14+0800
	//		Regions	否	String	区域名称， 取值为CN:中国大陆，HK：香港，TW：台湾，AS：亚洲其他，NA：北美洲，SA：南美洲，EU：欧洲，AU：大洋洲，AF：非洲，支持多区域查询，多个区域用逗号（半角）分隔，缺省为 CN
	//		ResultType	是	Long	取值为0：多域名多区域数据做合并；1：每个域名每个区域的数据分别返回。
	//		Granularity	否	Long	统计粒度，取值为 5（默认）：5分钟粒度；10：10分钟粒度；20：20分钟粒度；60：1小时粒度；240：4小时粒度；480：8小时粒度；1440：1天粒度；以上粒度的带宽值均取该粒度时间段的峰值
	//		DataType	否	String	数据类型， 取值为edge:边缘数据; origin:回源数据; 支持多类型选择，多个类型用逗号（半角）分隔，缺省为 edge
	//		ProtocolType	否	String	协议类型， 取值为http:http协议数据; https:https协议数据

	//getSubDomainsPvData := make(map[string]interface{})
	//getSubDomainsPvData["DomainId"] = "2D07WRH"
	//getSubDomainsPvData["Domains"] = ""
	//getSubDomainsPvData["StartTime"] = "2021-06-01T21:14+0800"
	//getSubDomainsPvData["EndTime"] = "2021-06-01T21:34+0800"
	//getSubDomainsPvData["Regions"] = "CN"
	//getSubDomainsPvData["ResultType"] = 0
	//getSubDomainsPvData["Granularity"] = 5
	//getSubDomainsPvData["DataType"] = "edge"
	//getSubDomainsPvData["ProtocolType"] = "http"
	//
	//resp, err = svc.GetSubDomainsPvDataPost(&getSubDomainsPvData)
	//if err != nil {
	//	fmt.Println("error:", err.Error())
	//}
	//if resp != nil {
	//	str, _ := json.Marshal(&resp)
	//	fmt.Printf("%+v\n", string(str))
	//}

	//	**************************************** 此接口用于设置视频拖拽 GET（SetVideoSeekConfig）***************************************************
	//  Parameters:
	//      DomainId	是	String	表示域名
	//		Enable	是	枚举值为:on,off 表示开关

	//setVideoSeekConfig := make(map[string]interface{})
	//setVideoSeekConfig["DomainId"] = "2D07WRH"
	//setVideoSeekConfig["Enable"] = "off"
	//
	//resp, err = svc.SetVideoSeekConfigGet(&setVideoSeekConfig)
	//if err != nil {
	//	fmt.Println("error:", err.Error())
	//}
	//if resp != nil {
	//	str, _ := json.Marshal(&resp)
	//	fmt.Printf("%+v\n", string(str))
	//}

	//	**************************************** 此接口用于设置视频拖拽 POST（SetVideoSeekConfig）***************************************************
	//  Parameters:
	//      DomainId	是	String	表示域名
	//		Enable	是	枚举值为:on,off 表示开关

	//setVideoSeekConfig := make(map[string]interface{})
	//setVideoSeekConfig["DomainId"] = "2D07WRH"
	//setVideoSeekConfig["Enable"] = "on"
	//
	//resp, err = svc.SetVideoSeekConfigPost(&setVideoSeekConfig)
	//if err != nil {
	//	fmt.Println("error:", err.Error())
	//}
	//if resp != nil {
	//	str, _ := json.Marshal(&resp)
	//	fmt.Printf("%+v\n", string(str))
	//}

	//	**************************************** 此接口用于获取域名拖拽信息 GET（GetVideoSeekConfig）***************************************************
	//  Parameters:
	//        DomainId	是	String	表示域名

	//getVideoSeekConfig := make(map[string]interface{})
	//getVideoSeekConfig["DomainId"] = "2D07WRH"
	//
	//resp, err = svc.GetVideoSeekConfigGet(&getVideoSeekConfig)
	//if err != nil {
	//	fmt.Println("error:", err.Error())
	//}
	//if resp != nil {
	//	str, _ := json.Marshal(&resp)
	//	fmt.Printf("%+v\n", string(str))
	//}

	//	**************************************** 此接口用于获取域名拖拽信息 POST（GetVideoSeekConfig）***************************************************
	//  Parameters:
	//        DomainId	是	String	表示域名

	//getVideoSeekConfig := make(map[string]interface{})
	//getVideoSeekConfig["DomainId"] = "2D07WRH"
	//
	//resp, err = svc.GetVideoSeekConfigPost(&getVideoSeekConfig)
	//if err != nil {
	//	fmt.Println("error:", err.Error())
	//}
	//if resp != nil {
	//	str, _ := json.Marshal(&resp)
	//	fmt.Printf("%+v\n", string(str))
	//}

	//	**************************************** 此接口用于设置Http响应头 GET（SetHttpHeadersConfig）***************************************************
	//  Parameters:
	//        DomainId	是	String 表示域名
	//        HeaderKey   是  String 表示设置的http头 只支持 Content-Type，Cache-Control，Content-Disposition，Content-Language，Expires，Access-Control-Allow-Origin，Access-Control-Allow-Methods，Access-Control-Max-Age，Access-Control-Expose-Headers 这9种
	//        HeaderValue 是  String 表示设置http头vaule值

	//setHttpHeadersConfig := make(map[string]interface{})
	//setHttpHeadersConfig["DomainId"] = "2D07WRH"
	//setHttpHeadersConfig["HeaderKey"] = "Expires"
	//setHttpHeadersConfig["HeaderValue"] = 20
	//
	//resp, err = svc.SetHttpHeadersConfigGet(&setHttpHeadersConfig)
	//if err != nil {
	//	fmt.Println("error:", err.Error())
	//}
	//if resp != nil {
	//	str, _ := json.Marshal(&resp)
	//	fmt.Printf("%+v\n", string(str))
	//}

	//	**************************************** 此接口用于设置Http响应头 POST（SetHttpHeadersConfig）***************************************************
	//  Parameters:
	//        DomainId	是	String 表示域名
	//        HeaderKey   是  String 表示设置的http头 只支持 Content-Type，Cache-Control，Content-Disposition，Content-Language，Expires，Access-Control-Allow-Origin，Access-Control-Allow-Methods，Access-Control-Max-Age，Access-Control-Expose-Headers 这9种
	//        HeaderValue 是  String 表示设置http头vaule值

	//setHttpHeadersConfig := make(map[string]interface{})
	//setHttpHeadersConfig["DomainId"] = "2D07WRH"
	//setHttpHeadersConfig["HeaderKey"] = "Expires"
	//setHttpHeadersConfig["HeaderValue"] = 40
	//
	//resp, err = svc.SetHttpHeadersConfigPost(&setHttpHeadersConfig)
	//if err != nil {
	//	fmt.Println("error:", err.Error())
	//}
	//if resp != nil {
	//	str, _ := json.Marshal(&resp)
	//	fmt.Printf("%+v\n", string(str))
	//}

	//	**************************************** 此接口用于删除Http响应头 GET（DeleteHttpHeadersConfig）***************************************************
	//  Parameters:
	//        DomainId	是	String 表示域名
	//        HeaderKey   是  String 表示设置的http头 只支持 Content-Type，Cache-Control，Content-Disposition，Content-Language，Expires，Access-Control-Allow-Origin，Access-Control-Allow-Methods，Access-Control-Max-Age，Access-Control-Expose-Headers 这9种

	//deleteHttpHeadersConfig := make(map[string]interface{})
	//deleteHttpHeadersConfig["DomainId"] = "2D07WRH"
	//deleteHttpHeadersConfig["HeaderKey"] = "Expires"
	//
	//resp, err = svc.DeleteHttpHeadersConfigGet(&deleteHttpHeadersConfig)
	//if err != nil {
	//	fmt.Println("error:", err.Error())
	//}
	//if resp != nil {
	//	str, _ := json.Marshal(&resp)
	//	fmt.Printf("%+v\n", string(str))
	//}

	//	**************************************** 此接口用于删除Http响应头 POST（DeleteHttpHeadersConfig）***************************************************
	//  Parameters:
	//        DomainId	是	String 表示域名
	//        HeaderKey   是  String 表示设置的http头 只支持 Content-Type，Cache-Control，Content-Disposition，Content-Language，Expires，Access-Control-Allow-Origin，Access-Control-Allow-Methods，Access-Control-Max-Age，Access-Control-Expose-Headers 这9种

	//deleteHttpHeadersConfig := make(map[string]interface{})
	//deleteHttpHeadersConfig["DomainId"] = "2D07WRH"
	//deleteHttpHeadersConfig["HeaderKey"] = "Expires"
	//
	//resp, err = svc.DeleteHttpHeadersConfigPost(&deleteHttpHeadersConfig)
	//if err != nil {
	//	fmt.Println("error:", err.Error())
	//}
	//if resp != nil {
	//	str, _ := json.Marshal(&resp)
	//	fmt.Printf("%+v\n", string(str))
	//}

	//	**************************************** 此接口用于获取http响应头 GET（GetHttpHeaderList）***************************************************
	//  Parameters:
	//        DomainId	是	String 表示域名

	//getHttpHeaderList := make(map[string]interface{})
	//getHttpHeaderList["DomainId"] = "2D07WRH"
	//
	//resp, err = svc.GetHttpHeaderListGet(&getHttpHeaderList)
	//if err != nil {
	//	fmt.Println("error:", err.Error())
	//}
	//if resp != nil {
	//	str, _ := json.Marshal(&resp)
	//	fmt.Printf("%+v\n", string(str))
	//}

	//	**************************************** 此接口用于获取http响应头 POST（GetHttpHeaderList）***************************************************
	//  Parameters:
	//        DomainId	是	String 表示域名

	//getHttpHeaderList := make(map[string]interface{})
	//getHttpHeaderList["DomainId"] = "2D07WRH"
	//
	//resp, err = svc.GetHttpHeaderListPost(&getHttpHeaderList)
	//if err != nil {
	//	fmt.Println("error:", err.Error())
	//}
	//if resp != nil {
	//	str, _ := json.Marshal(&resp)
	//	fmt.Printf("%+v\n", string(str))
	//}

	//	**************************************** 本接口用于获取某个时间点的播放统计信息，包括带宽、流量、在线人数，包括流维度和域名维度的数据 GET（GetLivePlayStatData）***************************************************
	//  Parameters:
	//        StartTime   String  获取数据起始时间点，日期格式按ISO8601表示法，北京时间，格式为：YYYY-MM-DDThh:mm+0800，例如：2016-08-01T21:14+0800
	//        DomainIds   String  域名ID，缺省为当前产品类型下的全部域名，可输入需要查询的域名ID，支持批量域名查询，多个域名ID用逗号（半角）分隔
	//        Regions     String  计费区域名称，取值为CN:中国大陆，HK：香港，TW：台湾，AS：亚洲其他，NA：北美洲，SA：南美洲，EU：欧洲，AU：大洋洲，AF：非洲，支持多计费区域查询，多个区域用逗号（半角）分隔，缺省为 CN
	//        ResultType  String  取值为0：只返回域名级别的汇总数据；1：返回域名级别+流维度的详细数据；
	//        LimitN 否 Int Top条数，取值为1-200，最大200，默认100

	//getLivePlayStatData := make(map[string]interface{})
	//getLivePlayStatData["DomainIds"] = "2D07WRH"
	//getLivePlayStatData["StartTime"] = "2021-06-01T21:14+0800"
	//getLivePlayStatData["EndTime"] = "2021-06-01T21:34+0800"
	//getLivePlayStatData["Regions"] = "CN"
	//getLivePlayStatData["ResultType"] = 0
	//getLivePlayStatData["LimitN"] = 10
	//
	//resp, err = svc.GetLivePlayStatDataGet(&getLivePlayStatData)
	//if err != nil {
	//	fmt.Println("error:", err.Error())
	//}
	//if resp != nil {
	//	str, _ := json.Marshal(&resp)
	//	fmt.Printf("%+v\n", string(str))
	//}

	//	**************************************** 本接口用于获取某个时间点的播放统计信息，包括带宽、流量、在线人数，包括流维度和域名维度的数据 POST（GetLivePlayStatData）***************************************************
	//  Parameters:
	//        StartTime   String  获取数据起始时间点，日期格式按ISO8601表示法，北京时间，格式为：YYYY-MM-DDThh:mm+0800，例如：2016-08-01T21:14+0800
	//        EndTime   String  获取数据结束时间点，日期格式按ISO8601表示法，北京时间，格式为：YYYY-MM-DDThh:mm+0800，例如：2016-08-01T21:14+0800
	//        DomainIds   String  域名ID，缺省为当前产品类型下的全部域名，可输入需要查询的域名ID，支持批量域名查询，多个域名ID用逗号（半角）分隔
	//        Regions     String  计费区域名称，取值为CN:中国大陆，HK：香港，TW：台湾，AS：亚洲其他，NA：北美洲，SA：南美洲，EU：欧洲，AU：大洋洲，AF：非洲，支持多计费区域查询，多个区域用逗号（半角）分隔，缺省为 CN
	//        ResultType  String  取值为0：只返回域名级别的汇总数据；1：返回域名级别+流维度的详细数据；
	//        LimitN 否 Int Top条数，取值为1-200，最大200，默认100

	getLivePlayStatData := make(map[string]interface{})
	getLivePlayStatData["DomainIds"] = "2D07WRH"
	getLivePlayStatData["StartTime"] = "2021-06-01T21:14+0800"
	getLivePlayStatData["EndTime"] = "2021-06-01T21:34+0800"
	getLivePlayStatData["Regions"] = "CN"
	getLivePlayStatData["ResultType"] = 0
	getLivePlayStatData["LimitN"] = 10

	resp, err = svc.GetLivePlayStatDataPost(&getLivePlayStatData)
	if err != nil {
		fmt.Println("error:", err.Error())
	}
	if resp != nil {
		str, _ := json.Marshal(&resp)
		fmt.Printf("%+v\n", string(str))
	}
}
