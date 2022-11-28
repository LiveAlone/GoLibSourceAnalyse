package _

import (
	"git.zuoyebang.cc/huixuexi/classflow/layer"
	"gorm.io/gorm"
)

const TableNameOfActivityBasic = "tblActivityBasic"

type TblActivityBasic struct {
    Id  int64  `gorm:"primarykey" json:"id"`      // 活动id 
    Name  string  `gorm:"column:name" json:"name"`      // 活动名称 
    ActivityType  int8  `gorm:"column:activity_type" json:"activity_type"`      // 活动类型 
    Division  string  `gorm:"column:division" json:"division"`      // 多学段，位标识 
    UserType  int8  `gorm:"column:user_type" json:"user_type"`      // 参与活动用户类型, 位标识教师学生 
    Banner  string  `gorm:"column:banner" json:"banner"`      // 宣传图 
    Content  string  `gorm:"column:content" json:"content"`      // 活动介绍 
    Appendix  string  `gorm:"column:appendix" json:"appendix"`      // 附件文件列表 
    Level  int8  `gorm:"column:level" json:"level"`      // 活动级别 1省 2市 3区 4校 
    ProvinceOrgId  int64  `gorm:"column:provinceOrgId" json:"provinceOrgId"`     
    ProvinceCode  int64  `gorm:"column:provinceCode" json:"provinceCode"`      // 省编码 
    CityOrgId  int64  `gorm:"column:cityOrgId" json:"cityOrgId"`     
    CityCode  int64  `gorm:"column:cityCode" json:"cityCode"`      // 市编码 
    AreaOrgId  int64  `gorm:"column:areaOrgId" json:"areaOrgId"`     
    AreaCode  int64  `gorm:"column:areaCode" json:"areaCode"`      // 区编码 
    SchoolId  int64  `gorm:"column:school_id" json:"school_id"`      // 校id 
    Status  int8  `gorm:"column:status" json:"status"`      // 活动状态 1 草稿  2 发布中 3 取消发布 4 删除 
    MaxWorkCount  int32  `gorm:"column:max_work_count" json:"max_work_count"`      // 每个用户最大作品数量 
    RewardGenerate  int8  `gorm:"column:reward_generate" json:"reward_generate"`      // 奖项生成方式 1 评审组  2 自动生成 
    RewardMaxScore  int32  `gorm:"column:reward_max_score" json:"reward_max_score"`      // 评分最大分数 
    RewardDetail  string  `gorm:"column:reward_detail" json:"reward_detail"`      // 奖项配置内容{} 奖项类型，数据结构 
    StartTime  time.Time  `gorm:"column:start_time" json:"start_time"`      // 活动整体开始时间 
    EndTime  time.Time  `gorm:"column:end_time" json:"end_time"`      // 活动整体结束时间 
    CreatedTime  time.Time  `gorm:"column:created_time" json:"created_time"`     
    UpdatedTime  time.Time  `gorm:"column:updated_time" json:"updated_time"`     
    
}

type ActivityBasicDao struct {
	models.Dao
}

func NewActivityBasicDao(ctx *gin.Context) (db *ActivityBasicDao) {
	db = new(ActivityBasicDao)
	db.Dao = models.NewDao(ctx, helpers.MysqlClientapps.WithContext(ctx))
	return
}

func (entity *ActivityBasicDao) OnCreate(param layer.IFlowParam) {
	entity.Dao.OnCreate(param)
	entity.SetTable(TableNameOfActivityBasic)
}

func (entity *ActivityBasicDao) Insert(bean *TblActivityBasic) error {
	err := entity.GetDB().Create(bean).Error
	if err != nil {
		entity.LogErrorf("db insert tblActivityBasic error, bean:%v, cause:%v", bean, err)
		return components.ErrorDbInsert.Sprintf(err.Error())
	}
	return nil
}

func (entity *ActivityBasicDao) Update(info *TblActivityBasic) error {
    if info.Id == 0 {
		return 0, components.ErrorParamInvalid
	}
	rs := entity.Dao.GetDB().Updates(info)
	if rs.Error != nil {
		entity.LogErrorf("db update tblActivityBasic error, info:%v, cause:%v", info, rs.Error)
		return 0, components.ErrorDbUpdate.Sprintf(err.Error())
	}
	return rs.RowsAffected, nil
}

func (entity *ActivityBasicDao) QueryById(id int64) (*TblActivityBasic, error) {
    if id == 0 {
		return 0, components.ErrorParamInvalid
	}
	info := new(TblActivityBasic)
	rs := entity.Dao.GetDB().Find(&info, id)
	if rs.Error == gorm.ErrRecordNotFound {
   		return nil, nil
   	}
	if rs.Error != nil {
		entity.LogErrorf("db query by id tblActivityBasic error, id:%v, cause:%v", id, rs.Error)
		return nil, components.ErrorDbSelect.Sprintf(err.Error())
	}
	return info, nil
}

