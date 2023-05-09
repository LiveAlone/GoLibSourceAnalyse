package main

import (
	"fmt"
	"github.com/LiveAlone/GoLibSourceAnalyse/utils/util"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

type ShuangJianTongOrg struct {
	OrgId         int64  `gorm:"column:org_id" json:"orgId" form:"orgId"`                         //组织id
	ParentId      int64  `gorm:"column:parent_id" json:"parentId" form:"parentId"`                // 父组织id
	Name          string `gorm:"column:name" json:"name" form:"name"`                             // 名称
	Status        int8   `gorm:"column:status" json:"status" form:"status"`                       // 状态(-1:停用,0:未生效,1:启用)
	UnitAttribute int8   `gorm:"column:unit_attribute" json:"unitAttribute" form:"unitAttribute"` // 单位层级除市本级(1：市教育局2:区县,3 学段4:学校,5:年级,6:班级)'  市本级：(1：市教育局2:区县,3:学校,4:年级,5:班级)
}

type MatchRecord struct {
	Id          int64  `gorm:"column:id" json:"id" form:"id"`
	Type        int8   `gorm:"column:type" json:"type" form:"type"`
	ThirdSource string `gorm:"column:third_source" json:"thirdSource" form:"thirdSource"`
	ThirdOrgId  string `gorm:"column:third_org_id" json:"thirdOrgId" form:"thirdOrgId"`
	YunsiOrgId  string `gorm:"column:yunsi_org_id" json:"yunsiOrgId" form:"yunsiOrgId"`
}

type SchCodeEntity struct {
	SchoolId string
	AreaCode int64
	AreaName string
}

var areaCollect = map[string]bool{
	"开发区": true,
	"武义县": true,
	"永康市": true,
	"婺城区": true,
	"磐安县": true,
	"义乌市": true,
	"金东区": true,
	"东阳市": true,
	"兰溪市": true,
	"市本级": true,
	"浦江县": true,
}

func main() {
	db, err := buildDb(hxxSchool)
	if err != nil {
		fmt.Println(err)
		return
	}

	schList, err := GenerateSchArea(db)
	if err != nil {
		fmt.Println(err)
		return
	}
	for i, entity := range schList {
		if !areaCollect[entity.AreaName] {
			log.Fatalf("entity not found area i:%v e:%v", i, entity)
		}
	}

	var rs []string
	for i, entity := range schList {
		record := &MatchRecord{}
		tx := db.Table("tblMatchRecord").Where("type = 1 and third_source = 'shuangjiantong' and third_org_id = ?", entity.SchoolId).Take(record)
		if tx.Error == gorm.ErrRecordNotFound {
			continue
		}
		if tx.Error != nil {
			log.Fatalf("query match record error:%v", tx.Error)
		}
		sql := fmt.Sprintf("insert into tblSjtSchoolArea(school_id, area_code, area_name) values (%s, %d, '%s');\n", record.YunsiOrgId, entity.AreaCode, entity.AreaName)
		rs = append(rs, sql)
		fmt.Println(i, "finish with sql", sql)
	}
	for i, r := range rs {
		fmt.Println(i, r)
	}
	_ = util.WriteFileLines("dest/sch_area.sql", rs)
}

func GenerateSchArea(db *gorm.DB) ([]SchCodeEntity, error) {
	schList := make([]ShuangJianTongOrg, 0)
	tx := db.Table("tblShuangJianTongOrg").Where("unit_attribute = 4 and status = 1").Scan(&schList)
	if tx.Error != nil {
		return nil, tx.Error
	}

	division := make([]int64, 0)
	mp := make(map[int64]bool)
	for _, org := range schList {
		if !mp[org.ParentId] {
			mp[org.ParentId] = true
			division = append(division, org.ParentId)
		}
	}
	divisionList := make([]ShuangJianTongOrg, 0)
	tx = db.Table("tblShuangJianTongOrg").Where("org_id in ? and status = 1", division).Scan(&divisionList)
	if tx.Error != nil {
		return nil, tx.Error
	}

	areaIdList := make([]int64, 0)
	divisionMap := make(map[int64]ShuangJianTongOrg)
	for _, org := range divisionList {
		areaIdList = append(areaIdList, org.ParentId)
		divisionMap[org.OrgId] = org
	}

	areaList := make([]ShuangJianTongOrg, 0)
	tx = db.Table("tblShuangJianTongOrg").Where("org_id in ? and status = 1", areaIdList).Scan(&areaList)
	if tx.Error != nil {
		return nil, tx.Error
	}

	areaMap := make(map[int64]ShuangJianTongOrg)
	for _, org := range areaList {
		areaMap[org.OrgId] = org
	}

	schAreaList := make([]SchCodeEntity, 0)
	for _, org := range schList {
		schAreaList = append(schAreaList, SchCodeEntity{
			SchoolId: fmt.Sprintf("%d", org.OrgId),
			AreaCode: areaMap[divisionMap[org.ParentId].ParentId].OrgId,
			AreaName: areaMap[divisionMap[org.ParentId].ParentId].Name,
		})
	}
	return schAreaList, nil
}

// ssh proxy tunnel
// ssh -L 8888:172.16.1.173:3306 root@123.156.228.100
const hxxSchool = "bawu_mysql:bawu#123@tcp(127.0.0.1:8888)/hxx_trans?charset=utf8mb4&parseTime=True&loc=Local"

//const hxxSchool = "homework:homework@tcp(10.117.0.4:3306)/hxx_school_qa?charset=utf8mb4&parseTime=True&loc=Local"

func buildDb(url string) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(url), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
