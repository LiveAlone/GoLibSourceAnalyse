package _

import (
    "github.com/gin-gonic/gin"

    "database/sql"

    "gorm.io/datatypes"
	"gorm.io/gorm"
)

const TableNameOfModelTable = "tblModelTable"

// ModelTable 数据表模型
type ModelTable struct {
    gorm.Model
    Name string //姓名
    Age int32 //年龄
    Score float64 //分数
    Tage sql.NullString //标签类型
    Height sql.NullFloat64 //高度
    RewardInfo datatypes.JSON //奖励信息json
}

type ModelTableDao struct {
    Ctx *gin.Context
    Tx  *gorm.DB
}

func NewModelTableDao(ctx *gin.Context) (db *ModelTableDao) {
	db = new(ModelTableDao)
	// todo 初始化db，context
	db.Ctx = ctx
	return
}

func (entity *ModelTableDao) Insert(bean *ModelTable) error {
	err := entity.Tx.Table(TableNameOfModelTable).Create(bean).Error
	if err != nil {
		zlog.Errorf(entity.Ctx, "db insert tblModelTable error, bean:%v, cause:%v", bean, err)
		return outerror.ErrorDbInsert.Sprintf(err.Error())
	}
	return nil
}

func (entity *ModelTableDao) Update(info *ModelTable) (int64, error) {
    if info.ID == 0 {
  	    return 0, outerror.ErrorParamInvalid
  	}
  	rs := entity.Tx.Table(TableNameOfModelTable).Updates(info)
  	if rs.Error != nil {
  		zlog.Errorf(entity.Ctx, "db update tblModelTable error, info:%v, cause:%v", info, rs.Error)
  		return 0, outerror.ErrorDbUpdate.Sprintf(rs.Error.Error())
  	}
  	return rs.RowsAffected, nil
}

func (entity *ModelTableDao) QueryById(id int64) (*ModelTable, error) {
    if id == 0 {
		return nil, outerror.ErrorParamInvalid
	}
	info := new(ModelTable)
	rs := entity.Tx.Table(TableNameOfModelTable).Find(&info, id)
	if rs.Error == gorm.ErrRecordNotFound {
   		return nil, nil
   	}
	if rs.Error != nil {
		zlog.Errorf(entity.Ctx, "db query by id tblModelTable error, id:%v, cause:%v", id, rs.Error)
		return nil, outerror.ErrorDbSelect.Sprintf(rs.Error.Error())
	}
	return info, nil
}

func (entity *ModelTableDao) QueryByIds(ids []int64) (list []*ModelTable, err error) {
    if len(ids) == 0 {
		return nil, nil
	}
	rs := entity.Tx.Table(TableNameOfModelTable).Find(&list, ids)
	if rs.Error != nil {
		zlog.Errorf(entity.Ctx, "db query by ids tblModelTable error, ids:%v, cause:%v", ids, rs.Error)
		return nil, outerror.ErrorDbSelect.Sprintf(rs.Error.Error())
	}
	return list, nil
}