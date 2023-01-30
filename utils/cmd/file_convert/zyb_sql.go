package file_convert

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

type MsgRole struct {
	ID     int64  `json:"id"`     // 角色ID
	Name   string `json:"name"`   // 角色名称
	Status int8   `json:"status"` // 状态(0:停用,1:启用)
}

var roleSql = "insert into tblShuangJianTongRole(role_id, name, status, created_at, updated_at) value (%d, '%s', %d, now(), now()) on duplicate key update name = '%s', status = %d, updated_at = now();\n"

func RoleConvert(line string) (string, bool) {
	role := new(MsgRole)
	if err := json.Unmarshal([]byte(line), role); err != nil {
		log.Fatal(err)
	}
	return fmt.Sprintf(roleSql, role.ID, role.Name, role.Status, role.Name, role.Status), false
}

type MsgUser struct {
	ID       int64  `json:"id"`       // 用户ID
	RealName string `json:"realName"` // 用户姓名
	Password string `json:"password"` // md5加密后的密码
	Phone    string `json:"phone"`    // 用户联系方式
	Account  string `json:"account"`  // 用户账号
	Status   int8   `json:"status"`   // 态(0:停用,1:启用)
}

var userSql = "insert into tblShuangJianTongUser(user_id, account, real_name, phone, password, status, created_at, updated_at) value (%d,'%s','%s','%s','%s',%d,now(),now()) on duplicate key update "
var userDup = "account='%s', real_name='%s', phone='%s', password='%s', status=%d, updated_at=now();\n"

func UserConvert(i int, line string) (string, bool) {
	if i >= 50000 {
		return "", true
	}

	user := new(MsgUser)
	if err := json.Unmarshal([]byte(line), user); err != nil {
		log.Fatal(err)
	}
	return fmt.Sprintf(userSql+userDup,
		user.ID, user.Account, user.RealName, user.Phone, user.Password, user.Status,
		user.Account, user.RealName, user.Phone, user.Password, user.Status), false
}

type MsgOrgRelationWrap struct {
	Data *MsgOrgRelations `json:"data"`
}

type MsgOrgRelations struct {
	OrgID         int64            `json:"orgId"`
	Name          string           `json:"name"`
	Category      int8             `json:"category"`
	Status        int8             `json:"status"`        // 状态(0:停用,1:启用)
	SortOrder     int64            `json:"sortOrder"`     //组织排序码
	UnitAttribute int8             `json:"unitAttribute"` //组织层级
	Flag          int64            `json:"flag"`          // 年级枚举，只有年级有值（1-6小学一到六年级7-9初中七到九年级11-13高中一到三年级14-17幼儿园）有是年级但是无该值情况，例如历史数据用户未录入该值，多建的组织超过学制的年级数量
	ParentID      int64            `json:"parentId"`
	OrgRelations  *MsgOrgRelations `json:"orgRelations"`
}

var orgInsertSql = "insert into tblShuangJianTongOrg(org_id, parent_id, name, sort_order, category, status, unit_attribute, flag, created_at, updated_at) value(%d,%d,'%s',%d,%d,%d,%d,%d,now(),now()) on duplicate key update "
var orgDupSql = "parent_id = %d, name = '%s', sort_order = %d, category = %d, status = %d, unit_attribute = %d, flag = %d,updated_at = now();"

func OrgConvert(line string) (string, bool) {
	orgWrap := new(MsgOrgRelationWrap)
	if err := json.Unmarshal([]byte(line), orgWrap); err != nil {
		log.Fatal(err)
	}
	sql := orgInsertSql + orgDupSql + "\n"

	var b strings.Builder
	org := orgWrap.Data
	ct := 0
	for org != nil && ct < 100 {
		newString := fmt.Sprintf(sql,
			org.OrgID, org.ParentID, org.Name, org.SortOrder, org.Category, org.Status, org.UnitAttribute, org.Flag,
			org.ParentID, org.Name, org.SortOrder, org.Category, org.Status, org.UnitAttribute, org.Flag)
		b.WriteString(newString)
		ct += 1
		org = org.OrgRelations
	}
	return b.String(), false
}

type MsgRelation struct {
	ID           int64 `json:"id"`           //关联关系id
	UserID       int64 `json:"userId"`       //用户id
	RoleID       int64 `json:"roleId"`       //用户角色id
	OrgID        int64 `json:"orgId"`        //用户组织id
	ParentUserID int64 `json:"parentUserId"` //用户关联ID（监护人必须关联学生，0表示无关联关系）
	Status       int8  `json:"status"`       //状态(0:停用,1:启用)
}

var relationSql = "insert into tblShuangJianTongUserRole(id, role_id, org_id, user_id, parent_user_id, status, created_at, updated_at) value (%d,%d,%d,%d,%d,%d,now(),now()) on duplicate key update "
var relationDup = " role_id=%d, org_id=%d, user_id=%d, parent_user_id=%d, status=%d, updated_at=now(); \n"

func RelationConvert(line string) (string, bool) {
	relation := new(MsgRelation)
	if err := json.Unmarshal([]byte(line), relation); err != nil {
		log.Fatal(err)
	}
	sql := relationSql + relationDup

	return fmt.Sprintf(sql,
		relation.ID, relation.RoleID, relation.OrgID, relation.UserID, relation.ParentUserID, relation.Status,
		relation.RoleID, relation.OrgID, relation.UserID, relation.ParentUserID, relation.Status), false
}
