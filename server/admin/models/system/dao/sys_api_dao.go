package dao

import (
	"context"
	"errors"
	"gitee.com/moorewqk/antcom/server/cores/g"
	"gitee.com/moorewqk/antcom/server/cores/models/system"
	"gitee.com/moorewqk/antcom/server/utils/request"
	"gorm.io/gorm"
)

type SysApiDao struct {
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateApi
//@description: 新增基础api
//@param: api model.SysApi
//@return: err error
func (sysApiDao *SysApiDao) Create(ctx context.Context, sysApi system.SysApi) (err error) {
	//todo
	if !errors.Is(g.GV_DB.Where("path = ? AND method = ?", sysApi.Path, sysApi.Method).First(&system.SysApi{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("存在相同api")
	}
	return g.GV_DB.Create(&sysApi).Error
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateApi
//@description: 根据id更新api
//@param: api model.SysApi
//@return: err error

func (sysApiDao *SysApiDao) Update(ctx context.Context, sysApi system.SysApi) (err error) {
	var oldsysApi system.SysApi
	if sysApi.ID == 0 {
		return errors.New("实例ID入参缺失")
	}
	err = g.GV_DB.Where("id = ?", sysApi.ID).First(&oldsysApi).Error
	if oldsysApi.Path != sysApi.Path || oldsysApi.Method != sysApi.Method {
		if !errors.Is(g.GV_DB.Where("path = ? AND method = ?", sysApi.Path, sysApi.Method).First(&system.SysApi{}).Error, gorm.ErrRecordNotFound) {
			return errors.New("存在相同api路径")
		}
	}
	if err != nil {
		return err
	} else {
		//err = CasbinServiceApp.UpdateCasbinApi(oldA.Path, api.Path, oldA.Method, api.Method)
		//if err != nil {
		//	return err
		//} else {
		err = g.GV_DB.Save(&sysApi).Error
		//}
	}
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetApiById
//@description: 根据id获取api
//@param: id float64
//@return: err error, api model.SysApi

func (sysApiDao *SysApiDao) Detail(ctx context.Context, id float64) (err error, sysApi system.SysApi) {
	err = g.GV_DB.Where("id = ?", id).First(&sysApi).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetAPIInfoList
//@description: 分页获取数据,
//@param: api model.SysApi, info request.PageInfo, order string, desc bool
//@return: err error
func (sysApiDao *SysApiDao) List(ctx context.Context, sysApi system.SysApi, info request.PageInfo, order string, desc bool) (err error, list interface{}, total int64) {
	/*
		sysApi system.SysApi,  对象
		info request.PageInfo, 分页
		order string, 		   排序
		desc bool				降序

	*/

	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := g.GV_DB.Model(&system.SysApi{})
	var apiList []system.SysApi

	if sysApi.Path != "" {
		db = db.Where("path LIKE ?", "%"+sysApi.Path+"%")
	}

	if sysApi.Description != "" {
		db = db.Where("description LIKE ?", "%"+sysApi.Description+"%")
	}

	if sysApi.Method != "" {
		db = db.Where("method = ?", sysApi.Method)
	}

	if sysApi.ApiGroup != "" {
		db = db.Where("api_group = ?", sysApi.ApiGroup)
	}

	err = db.Count(&total).Error

	if err != nil {
		return err, apiList, total
	} else {
		db = db.Limit(limit).Offset(offset)
		if order != "" {
			var OrderStr string
			if desc {
				OrderStr = order + " desc"
			} else {
				OrderStr = order
			}
			err = db.Order(OrderStr).Find(&apiList).Error
		} else {
			err = db.Order("api_group").Find(&apiList).Error
		}
	}
	return err, apiList, total
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteApi
//@description: 删除基础api
//@param: api model.SysApi
//@return: err error
func (sysApiDao *SysApiDao) Delete(ctx context.Context, id int) (err error) {
	var (
		sysApi *system.SysApi
	)

	g.GV_DB.Where("id = ?", id).First(sysApi)
	err = g.GV_DB.Delete(&sysApi).Error
	//CasbinServiceApp.ClearCasbin(1, api.Path, api.Method)
	return err

}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetAllApis
//@description: 获取所有的api
//@return: err error, apis []model.SysApi
//
//func (sysApiDao *SysApiDao) GetAllApis() (err error, apis []system.SysApi) {
//	err = g.GV_DB.Find(&apis).Error
//	return
//}

////@author: [piexlmax](https://github.com/piexlmax)
////@function: DeleteApis
////@description: 删除选中API
////@param: apis []model.SysApi
////@return: err error
//
//func (sysApiDao *SysApiDao) DeleteApisByIds(ids request.IdsReq) (err error) {
//	err = g.GV_DB.Delete(&[]system.SysApi{}, "id in ?", ids.Ids).Error
//	return err
//}
//
//func (sysApiDao *SysApiDao) DeleteApiByIds(ids []string) (err error) {
//	return g.GV_DB.Delete(system.SysApi{}, ids).Error
//}
