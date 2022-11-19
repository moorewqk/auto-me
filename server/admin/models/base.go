package models

import (
	"context"
)

type Dao interface {
	//创建
	Create(ctx context.Context, ob *interface{}) error
	//更新
	Update(ctx context.Context, newob *interface{}) (ob *interface{}, err error)
	//查询单个对象详情
	Detail(ctx context.Context, id int) (ob *interface{}, err error)
	//查询多个对象列表
	List(ctx context.Context, filter interface{}) ([]*interface{}, error)
	//删除单个对象
	Delete(ctx context.Context, id int) error
}
