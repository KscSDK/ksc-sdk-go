package main

import (
	"github.com/KscSDK/ksc-sdk-go/example/nosql"
	"github.com/KscSDK/ksc-sdk-go/ksc"
	"github.com/KscSDK/ksc-sdk-go/ksc/utils"
	"github.com/KscSDK/ksc-sdk-go/service/mongodb"
)

var m *mongodb.Mongodb

func init() {
	m = mongodb.SdkNew(ksc.NewClient(nosql.AK, nosql.SK), &ksc.Config{Region: &nosql.Region}, &utils.UrlInfo{
		UseSSL: false,
		Locate: false,
	})
}

func CreateMongoDBInstance() {
	form := make(map[string]interface{})
	form["Name"] = "myMongo"
	form["DbVersion"] = "3.6"
	form["InstanceClass"] = "1C2G"
	form["Storage"] = "5"
	form["PayType"] = "byDay"
	form["VpcId"] = "vpcId"
	form["VnetId"] = "vnetId"
	form["InstancePassword"] = "123123"
	form["AvailabilityZone.1"] = "az"
	form["NodeNum"] = "3"
	nosql.ProcessResult(m.CreateMongoDBInstance(&form))
}

func DeleteMongoDBInstance() {
	form := make(map[string]interface{})
	form["InstanceId"] = "instanceId"
}

func DescribeMongoDBInstances() {
	nosql.ProcessResult(m.DescribeMongoDBInstances(nil))
}

func DescribeMongoDBInstance() {
	form := make(map[string]interface{})
	form["InstanceId"] = "instanceId"
	nosql.ProcessResult(m.DescribeMongoDBInstance(&form))
}

func RenameMongoDBInstance() {
	form := make(map[string]interface{})
	form["InstanceId"] = "instanceId"
	form["Name"] = "myMongo_rename"
	nosql.ProcessResult(m.RenameMongoDBInstance(&form))
}

func RestartMongoDBInstance() {
	form := make(map[string]interface{})
	form["InstanceId"] = "instanceId"
	nosql.ProcessResult(m.RestartMongoDBInstance(&form))
}

func DescribeMongoDBInstanceNode() {
	form := make(map[string]interface{})
	form["InstanceId"] = "instanceId"
	nosql.ProcessResult(m.DescribeMongoDBInstanceNode(&form))
}

func DescribeMongoDBShardNode() {
	form := make(map[string]interface{})
	form["InstanceId"] = "instanceId"
	nosql.ProcessResult(m.DescribeMongoDBShardNode(&form))
}

func AddSecondaryInstance() {
	form := make(map[string]interface{})
	form["InstanceId"] = "instanceId"
	form["NodeNum"] = "3"
	nosql.ProcessResult(m.AddSecondaryInstance(&form))
}

func CreateMongoDBShardInstance() {
	form := make(map[string]interface{})
	form["Name"] = "myMongo"
	form["InstancePassword"] = "123123"
	form["Storage"] = "5"
	form["VpcId"] = "vpcId"
	form["VnetId"] = "vnetId"
	form["DbVersion"] = "3.6"
	form["PayType"] = "byDay"
	form["ShardNum"] = "3"
	form["MongosNum"] = "3"
	form["ShardClass"] = "1C2G"
	form["MongosClass"] = "1C2G"
	form["NetMode"] = "Multiple_VIP"
	form["Mode"] = "Cluster"
	form["AvailabilityZone.1"] = "az"
	nosql.ProcessResult(m.CreateMongoDBShardInstance(&form))
}

func DescribeRegions() {
	nosql.ProcessResult(m.DescribeRegions(nil))
}

func AddSecurityGroupRule() {
	form := make(map[string]interface{})
	form["InstanceId"] = "instanceId"
	form["cidrs"] = "192.168.18.17/21"
	nosql.ProcessResult(m.AddSecurityGroupRule(&form))
}

func DeleteSecurityGroupRules() {
	form := make(map[string]interface{})
	form["InstanceId"] = "instanceId"
	form["cidrs"] = "192.168.18.17/21"
	nosql.ProcessResult(m.DeleteSecurityGroupRules(&form))
}

func ListSecurityGroupRules() {
	form := make(map[string]interface{})
	form["InstanceId"] = "instanceId"
	nosql.ProcessResult(m.ListSecurityGroupRules(&form))
}

func CreateMongoDBSnapshot() {
	form := make(map[string]interface{})
	form["InstanceId"] = "instanceId"
	form["Name"] = "myMongodb_snapshot"
	nosql.ProcessResult(m.CreateMongoDBSnapshot(&form))
}

func SetMongoDBTimingSnapshot() {
	form := make(map[string]interface{})
	form["InstanceId"] = "instanceId"
	form["TimingSwitch"] = "On"
	form["Timezone"] = "10:00-11:00"
	form["TimeCycle"] = "Mon,Tues,Wed"
	nosql.ProcessResult(m.SetMongoDBTimingSnapshot(&form))
}

func DeleteMongoDBSnapshot() {
	form := make(map[string]interface{})
	form["InstanceId"] = "instanceId"
	form["SnapshotId"] = "snapshotId"
	nosql.ProcessResult(m.DeleteMongoDBSnapshot(&form))
}

func DescribeMongoDBSnapshot() {
	form := make(map[string]interface{})
	form["InstanceId"] = "instanceId"
	nosql.ProcessResult(m.DescribeMongoDBSnapshot(&form))
}

func RenameMongoDBSnapshot() {
	form := make(map[string]interface{})
	form["SnapshotId"] = "snapshotId"
	form["Name"] = "myMongodb_snapshot_rename"
	nosql.ProcessResult(m.DeleteMongoDBSnapshot(&form))
}

func main() {

	CreateMongoDBInstance()
	DeleteMongoDBInstance()
	DescribeMongoDBInstances()
	DescribeMongoDBInstance()
	RenameMongoDBInstance()
	RestartMongoDBInstance()
	DescribeMongoDBInstanceNode()
	DescribeMongoDBShardNode()
	AddSecondaryInstance()
	CreateMongoDBShardInstance()
	DescribeRegions()
	AddSecurityGroupRule()
	DeleteSecurityGroupRules()
	ListSecurityGroupRules()
	CreateMongoDBSnapshot()
	SetMongoDBTimingSnapshot()
	DeleteMongoDBSnapshot()
	DescribeMongoDBSnapshot()
	RenameMongoDBSnapshot()
}
