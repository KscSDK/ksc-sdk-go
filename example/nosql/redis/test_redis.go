package main

import (
	"github.com/KscSDK/ksc-sdk-go/example/nosql"
	"github.com/KscSDK/ksc-sdk-go/ksc"
	"github.com/KscSDK/ksc-sdk-go/ksc/utils"
	"github.com/KscSDK/ksc-sdk-go/service/kcsv1"
	"github.com/KscSDK/ksc-sdk-go/service/kcsv2"
)

var v1 *kcsv1.Kcsv1
var v2 *kcsv2.Kcsv2

func init() {
	v1 = kcsv1.SdkNew(ksc.NewClient(nosql.AK, nosql.SK), &ksc.Config{Region: &nosql.Region}, &utils.UrlInfo{
		UseSSL: false,
		Locate: false,
	})

	v2 = kcsv2.SdkNew(ksc.NewClient(nosql.AK, nosql.SK), &ksc.Config{Region: &nosql.Region}, &utils.UrlInfo{
		UseSSL: false,
		Locate: false,
	})
}

func CreateCacheCluster() {
	form := make(map[string]interface{})
	form["AvailableZone"] = "az"
	form["Engine"] = "redis"
	form["Name"] = "myRedisSingle"
	form["Mode"] = 2
	form["Capacity"] = 1
	form["BillType"] = 5
	form["NetType"] = 2
	form["VpcId"] = "vpcId"
	form["VnetId"] = "vnetId"
	form["Protocol"] = "4.0"
	nosql.ProcessResult(v1.CreateCacheCluster(&form))
}

func CreateCacheParameterGroup() {
	form := make(map[string]interface{})
	form["AvailableZone"] = "az"
	form["Engine"] = "redis"
	form["Name"] = "myRedisParamGroup"
	form["Description"] = "myRedisParamGroup"
	form["Parameters.ParameterName.1"] = "maxmemory-policy"
	form["Parameters.ParameterValue.1"] = "volatile-random"
	form["Parameters.ParameterName.2"] = "appendonly"
	form["Parameters.ParameterValue.2"] = "yes"
	nosql.ProcessResult(v1.CreateCacheParameterGroup(&form))
}

func CreateCacheSecurityGroup() {
	form := make(map[string]interface{})
	form["AvailableZone"] = "az"
	form["Engine"] = "redis"
	form["CacheId"] = "cacheId"
	form["Name"] = "myRedisParamGroup"
	form["Description"] = "myRedisParamGroup"
	form["SecurityGroupRules.Cidr.1"] = "192.168.18.17/21"
	form["SecurityGroupRules.Cidr.2"] = "192.168.18.18/21"
	nosql.ProcessResult(v1.CreateCacheSecurityGroup(&form))
}

func CreateSnapshot() {
	form := make(map[string]interface{})
	form["AvailableZone"] = "az"
	form["Engine"] = "redis"
	form["CacheId"] = "cacheId"
	form["Name"] = "myRedisSnapshot"
	nosql.ProcessResult(v1.CreateSnapshot(&form))
}

func DeleteCacheCluster() {
	form := make(map[string]interface{})
	form["AvailableZone"] = "az"
	form["Engine"] = "redis"
	form["CacheId"] = "cacheId"
	nosql.ProcessResult(v1.DeleteCacheCluster(&form))
}

func DeleteCacheParameterGroup() {
	form := make(map[string]interface{})
	form["AvailableZone"] = "az"
	form["Engine"] = "redis"
	form["CacheParameterGroupId"] = "123123"
	nosql.ProcessResult(v1.DeleteCacheParameterGroup(&form))
}

func DeleteCacheSecurityGroup() {
	form := make(map[string]interface{})
	form["AvailableZone"] = "az"
	form["Engine"] = "redis"
	form["CacheSecurityGroupId"] = "123123"
	nosql.ProcessResult(v1.DeleteCacheSecurityGroup(&form))
}

func DeleteCacheSecurityGroupRule() {
	form := make(map[string]interface{})
	form["AvailableZone"] = "az"
	form["Engine"] = "redis"
	form["CacheSecurityGroupId"] = "123123"
	form["SecurityRuleId"] = "123123"
	nosql.ProcessResult(v1.DeleteCacheSecurityGroupRule(&form))
}

func DeleteCacheSecurityRule() {
	form := make(map[string]interface{})
	form["AvailableZone"] = "az"
	form["Engine"] = "redis"
	form["CacheId"] = "cacheId"
	form["SecurityRuleId"] = "123123"
	nosql.ProcessResult(v1.DeleteCacheSecurityRule(&form))
}

func DeleteSnapshot() {
	form := make(map[string]interface{})
	form["AvailableZone"] = "az"
	form["Engine"] = "redis"
	form["SnapshotId"] = "123123"
	nosql.ProcessResult(v1.DeleteSnapshot(&form))
}

func DescribeAvailabilityZones() {
	nosql.ProcessResult(v1.DescribeAvailabilityZones(nil))
}

func DescribeCacheCluster() {
	form := make(map[string]interface{})
	form["AvailableZone"] = "az"
	form["Engine"] = "redis"
	form["CacheId"] = "cacheId"
	nosql.ProcessResult(v1.DescribeCacheCluster(&form))
}

func DescribeCacheClusters() {
	form := make(map[string]interface{})
	form["AvailableZone"] = "az"
	form["Engine"] = "redis"
	nosql.ProcessResult(v1.DescribeCacheClusters(&form))
}

func DescribeCacheDefaultParameters() {
	form := make(map[string]interface{})
	form["AvailableZone"] = "az"
	form["Engine"] = "redis"
	form["ParamVersion"] = "4.0"
	nosql.ProcessResult(v1.DescribeCacheDefaultParameters(&form))
}

func DescribeCacheParameterGroup() {
	form := make(map[string]interface{})
	form["AvailableZone"] = "az"
	form["Engine"] = "redis"
	form["CacheParameterGroupId"] = "123"
	nosql.ProcessResult(v1.DescribeCacheParameterGroup(&form))
}

func DescribeCacheParameterGroups() {
	form := make(map[string]interface{})
	form["AvailableZone"] = "az"
	form["Engine"] = "redis"
	form["CacheParameterGroupId"] = "123"
	form["ParamVersion"] = "4.0"
	nosql.ProcessResult(v1.DescribeCacheParameterGroups(&form))
}

func DescribeCacheParameters() {
	form := make(map[string]interface{})
	form["AvailableZone"] = "az"
	form["Engine"] = "redis"
	form["CacheId"] = "cacheId"
	nosql.ProcessResult(v1.DescribeCacheParameters(&form))
}

func DescribeCacheSecurityGroup() {
	form := make(map[string]interface{})
	form["AvailableZone"] = "az"
	form["Engine"] = "redis"
	form["CacheSecurityGroupId"] = "1212"
	nosql.ProcessResult(v1.DescribeCacheSecurityGroup(&form))
}

func DescribeCacheSecurityGroups() {
	form := make(map[string]interface{})
	form["AvailableZone"] = "az"
	form["Engine"] = "redis"
	nosql.ProcessResult(v1.DescribeCacheSecurityGroups(&form))
}

func DescribeCacheSecurityRules() {
	form := make(map[string]interface{})
	form["AvailableZone"] = "az"
	form["Engine"] = "redis"
	form["CacheId"] = "cacheId"
	nosql.ProcessResult(v1.DescribeCacheSecurityRules(&form))
}

func DescribeRegions() {
	nosql.ProcessResult(v1.DescribeRegions(nil))
}

func DescribeSnapshots() {
	form := make(map[string]interface{})
	form["AvailableZone"] = "az"
	form["Engine"] = "redis"
	form["CacheId"] = "cacheId"
	nosql.ProcessResult(v1.DescribeSnapshots(&form))
}

func DownloadSnapshot() {
	form := make(map[string]interface{})
	form["AvailableZone"] = "az"
	form["Engine"] = "redis"
	form["SnapshotId"] = "121"
	nosql.ProcessResult(v1.DownloadSnapshot(&form))
}

func ExportSnapshot() {
	form := make(map[string]interface{})
	form["AvailableZone"] = "az"
	form["Engine"] = "redis"
	form["SnapshotId"] = "121"
	form["BucketName"] = "xxxx"
	nosql.ProcessResult(v1.ExportSnapshot(&form))
}

func FlushCacheCluster() {
	form := make(map[string]interface{})
	form["AvailableZone"] = "az"
	form["Engine"] = "redis"
	form["CacheId"] = "cacheId"
	nosql.ProcessResult(v1.FlushCacheCluster(&form))
}

func ModifyCacheParameterGroup() {
	form := make(map[string]interface{})
	form["AvailableZone"] = "az"
	form["Engine"] = "redis"
	form["Name"] = "myRedisParamGroup"
	form["Description"] = "myRedisParamGroup"
	form["ParamVersion"] = "4.0"
	form["CacheParameterGroupId"] = "1231"
	form["Parameters.ParameterName.1"] = "maxmemory-policy"
	form["Parameters.ParameterValue.1"] = "volatile-random"
	form["Parameters.ParameterName.2"] = "appendonly"
	form["Parameters.ParameterValue.2"] = "yes"
	nosql.ProcessResult(v1.ModifyCacheParameterGroup(&form))
}

func ModifyCacheSecurityGroup() {
	form := make(map[string]interface{})
	form["AvailableZone"] = "az"
	form["Engine"] = "redis"
	form["CacheId"] = "cacheId"
	form["Name"] = "myRedisParamGroup"
	form["CacheSecurityGroupId"] = "1231"
	form["Description"] = "myRedisParamGroup"
	form["SecurityGroupRules.Cidr.1"] = "192.168.18.17/21"
	form["SecurityGroupRules.Cidr.2"] = "192.168.18.18/21"
	nosql.ProcessResult(v1.ModifyCacheSecurityGroup(&form))
}

func RenameCacheCluster() {
	form := make(map[string]interface{})
	form["AvailableZone"] = "az"
	form["Engine"] = "redis"
	form["CacheId"] = "cacheId"
	form["Name"] = "xxxx"
	nosql.ProcessResult(v1.RenameCacheCluster(&form))
}

func RenameSnapshot() {
	form := make(map[string]interface{})
	form["AvailableZone"] = "az"
	form["Engine"] = "redis"
	form["SnapshotId"] = "123123"
	form["Name"] = "xxxx"
	nosql.ProcessResult(v1.RenameSnapshot(&form))
}

func ResizeCacheCluster() {
	form := make(map[string]interface{})
	form["AvailableZone"] = "az"
	form["Engine"] = "redis"
	form["CacheId"] = "cacheId"
	form["Capacity"] = 2
	nosql.ProcessResult(v1.ResizeCacheCluster(&form))
}

func RestoreSnapshot() {
	form := make(map[string]interface{})
	form["AvailableZone"] = "az"
	form["Engine"] = "redis"
	form["SnapshotId"] = "123123"
	nosql.ProcessResult(v1.RestoreSnapshot(&form))
}

func SetCacheParameterGroup() {
	form := make(map[string]interface{})
	form["AvailableZone"] = "az"
	form["Engine"] = "redis"
	form["CacheId"] = "cacheId"
	form["CacheParameterGroupId"] = "122"
	nosql.ProcessResult(v1.SetCacheParameterGroup(&form))
}

func SetCacheParameters() {
	form := make(map[string]interface{})
	form["AvailableZone"] = "az"
	form["Engine"] = "redis"
	form["CacheId"] = "cacheId"
	form["Protocol"] = "4.0"
	form["Parameters.ParameterName.1"] = "maxmemory-policy"
	form["Parameters.ParameterValue.1"] = "volatile-random"
	form["Parameters.ParameterName.2"] = "appendonly"
	form["Parameters.ParameterValue.2"] = "yes"
	nosql.ProcessResult(v1.SetCacheParameters(&form))
}

func SetCacheSecurityGroup() {
	form := make(map[string]interface{})
	form["AvailableZone"] = "az"
	form["Engine"] = "redis"
	form["CacheId"] = "cacheId"
	form["CacheSecurityGroupId"] = "122"
	nosql.ProcessResult(v1.SetCacheSecurityGroup(&form))
}

func SetCacheSecurityRules() {
	form := make(map[string]interface{})
	form["AvailableZone"] = "az"
	form["Engine"] = "redis"
	form["CacheId"] = "cacheId"
	form["SecurityGroupRules.Cidr.1"] = "192.168.18.17/21"
	form["SecurityGroupRules.Cidr.2"] = "192.168.18.18/21"
	nosql.ProcessResult(v1.SetCacheSecurityRules(&form))
}

func SetTimingSnapshot() {
	form := make(map[string]interface{})
	form["AvailableZone"] = "az"
	form["Engine"] = "redis"
	form["CacheId"] = "cacheId"
	form["TimingSwitch"] = "on"
	form["Timezone"] = "01:00-02:00"
	nosql.ProcessResult(v1.SetTimingSnapshot(&form))
}

func AddCacheSlaveNode() {
	form := make(map[string]interface{})
	form["AvailableZone"] = "az"
	form["Engine"] = "redis"
	form["CacheId"] = "cacheId"
	nosql.ProcessResult(v2.AddCacheSlaveNode(&form))
}

func DeleteCacheSlaveNode() {
	form := make(map[string]interface{})
	form["AvailableZone"] = "az"
	form["Engine"] = "redis"
	form["CacheId"] = "cacheId"
	form["NodeId"] = "a1adb6b5-1575-4833-9a4f-536c152912314"
	nosql.ProcessResult(v2.DeleteCacheSlaveNode(&form))
}

func DescribeCacheReadonlyNode() {
	form := make(map[string]interface{})
	form["AvailableZone"] = "az"
	form["Engine"] = "redis"
	form["CacheId"] = "cacheId"
	nosql.ProcessResult(v2.DescribeCacheReadonlyNode(&form))
}

func UpdatePassword() {
	form := make(map[string]interface{})
	form["AvailableZone"] = "az"
	form["Engine"] = "redis"
	form["CacheId"] = "cacheId"
	form["Password"] = "ShiWo1101"
	form["Mode"] = "1"
	nosql.ProcessResult(v1.UpdatePassword(&form))
}

func main() {
	CreateCacheCluster()
	CreateCacheParameterGroup()
	CreateCacheSecurityGroup()
	CreateSnapshot()
	DeleteCacheCluster()
	DeleteCacheParameterGroup()
	DeleteCacheSecurityGroup()
	DeleteCacheSecurityGroupRule()
	DeleteCacheSecurityRule()
	DeleteSnapshot()
	DescribeAvailabilityZones()
	DescribeCacheCluster()
	DescribeCacheClusters()
	DescribeCacheDefaultParameters()
	DescribeCacheParameterGroup()
	DescribeCacheParameterGroups()
	DescribeCacheParameters()
	DescribeCacheSecurityGroup()
	DescribeCacheSecurityGroups()
	DescribeCacheSecurityRules()
	DescribeRegions()
	DescribeSnapshots()
	DownloadSnapshot()
	ExportSnapshot()
	FlushCacheCluster()
	ModifyCacheParameterGroup()
	ModifyCacheSecurityGroup()
	RenameCacheCluster()
	RenameSnapshot()
	ResizeCacheCluster()
	RestoreSnapshot()
	SetCacheParameterGroup()
	SetCacheParameters()
	SetCacheSecurityGroup()
	SetCacheSecurityRules()
	SetTimingSnapshot()
	AddCacheSlaveNode()
	DeleteCacheSlaveNode()
	DescribeCacheReadonlyNode()
}
