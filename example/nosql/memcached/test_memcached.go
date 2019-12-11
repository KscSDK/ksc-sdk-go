package main

import (
	"github.com/KscSDK/ksc-sdk-go/example/nosql"
	"github.com/KscSDK/ksc-sdk-go/ksc"
	"github.com/KscSDK/ksc-sdk-go/ksc/utils"
	"github.com/KscSDK/ksc-sdk-go/service/memcached"
)

var m *memcached.Memcached

func init() {
	m = memcached.SdkNew(ksc.NewClient(nosql.AK, nosql.SK), &ksc.Config{Region: &nosql.Region}, &utils.UrlInfo{
		UseSSL: false,
		Locate: false,
	})
}

func CreateCacheCluster() {
	form := make(map[string]interface{})
	form["AvailableZone"] = "cn-beijing-6b"
	form["Engine"] = "memcached"
	form["Name"] = "myRedisSingle"
	form["Mode"] = 2
	form["Capacity"] = 1
	form["BillType"] = 5
	form["NetType"] = 2
	form["VpcId"] = "b2adb6b5-1575-4833-9a4f-536c152932f4"
	form["VnetId"] = "b2adb6b5-1575-4833-9a4f-536c152932f4"
	form["Protocol"] = "4.0"
	nosql.ProcessResult(m.CreateCacheCluster(&form))
}

func DeleteCacheCluster() {
	form := make(map[string]interface{})
	form["AvailableZone"] = "cn-beijing-6b"
	form["Engine"] = "memcached"
	form["CacheId"] = "b2adb6b5-1575-4833-9a4f-536c152932f4"
	nosql.ProcessResult(m.DeleteCacheCluster(&form))
}

func DeleteCacheSecurityRule() {
	form := make(map[string]interface{})
	form["AvailableZone"] = "cn-beijing-6b"
	form["Engine"] = "memcached"
	form["CacheId"] = "b2adb6b5-1575-4833-9a4f-536c152932f4"
	form["SecurityRuleId"] = "123123"
	nosql.ProcessResult(m.DeleteCacheSecurityRule(&form))
}

func DescribeAvailabilityZones() {
	nosql.ProcessResult(m.DescribeAvailabilityZones(nil))
}

func DescribeCacheCluster() {
	form := make(map[string]interface{})
	form["AvailableZone"] = "cn-beijing-6b"
	form["Engine"] = "memcached"
	form["CacheId"] = "b2adb6b5-1575-4833-9a4f-536c152932f4"
	nosql.ProcessResult(m.DescribeCacheCluster(&form))
}

func DescribeCacheClusters() {
	form := make(map[string]interface{})
	form["AvailableZone"] = "cn-shanghai-2a"
	form["Engine"] = "memcached"
	nosql.ProcessResult(m.DescribeCacheClusters(&form))
}

func DescribeCacheSecurityRules() {
	form := make(map[string]interface{})
	form["AvailableZone"] = "cn-beijing-6b"
	form["Engine"] = "memcached"
	form["CacheId"] = "b2adb6b5-1575-4833-9a4f-536c152932f4"
	nosql.ProcessResult(m.DescribeCacheSecurityRules(&form))
}

func DescribeRegions() {
	nosql.ProcessResult(m.DescribeRegions(nil))
}

func FlushCacheCluster() {
	form := make(map[string]interface{})
	form["AvailableZone"] = "cn-beijing-6b"
	form["Engine"] = "memcached"
	form["CacheId"] = "b2adb6b5-1575-4833-9a4f-536c152932f4"
	nosql.ProcessResult(m.FlushCacheCluster(&form))
}

func RenameCacheCluster() {
	form := make(map[string]interface{})
	form["AvailableZone"] = "cn-beijing-6b"
	form["Engine"] = "memcached"
	form["CacheId"] = "b2adb6b5-1575-4833-9a4f-536c152932f4"
	form["Name"] = "xxxx"
	nosql.ProcessResult(m.RenameCacheCluster(&form))
}

func ResizeCacheCluster() {
	form := make(map[string]interface{})
	form["AvailableZone"] = "cn-beijing-6b"
	form["Engine"] = "memcached"
	form["CacheId"] = "b2adb6b5-1575-4833-9a4f-536c152932f4"
	form["Capacity"] = 2
	nosql.ProcessResult(m.ResizeCacheCluster(&form))
}

func SetCacheSecurityRules() {
	form := make(map[string]interface{})
	form["AvailableZone"] = "cn-beijing-6b"
	form["Engine"] = "memcached"
	form["CacheId"] = "b2adb6b5-1575-4833-9a4f-536c152932f4"
	form["SecurityGroupRules.Cidr.1"] = "192.168.18.17/21"
	form["SecurityGroupRules.Cidr.2"] = "192.168.18.18/21"
	nosql.ProcessResult(m.SetCacheSecurityRules(&form))
}

func UpdatePassword() {
	form := make(map[string]interface{})
	form["AvailableZone"] = "cn-beijing-6b"
	form["Engine"] = "memcached"
	form["CacheId"] = "b2adb6b5-1575-4833-9a4f-536c152932f4"
	form["Password"] = "123123"
	form["Mode"] = 2
	nosql.ProcessResult(m.UpdatePassword(&form))
}

func main() {
	CreateCacheCluster()
	DeleteCacheCluster()
	DeleteCacheSecurityRule()
	DescribeAvailabilityZones()
	DescribeCacheCluster()
	DescribeCacheClusters()
	DescribeCacheSecurityRules()
	DescribeRegions()
	FlushCacheCluster()
	RenameCacheCluster()
	ResizeCacheCluster()
	SetCacheSecurityRules()
	UpdatePassword()
}
