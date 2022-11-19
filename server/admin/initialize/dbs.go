package initialize

import (
	"errors"
	"fmt"
	"gitee.com/moorewqk/antcom/server/cores/g"
	"gitee.com/moorewqk/antcom/server/cores/models/system"
	"gitee.com/moorewqk/antcom/server/cores/types"
	"gitee.com/moorewqk/antcom/server/utils"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

/*
数据库连接
数据库迁移
*/

var (
	Tables []interface{}
)

type GromManger struct {
	*gorm.DB
}

/*加载应用数据

 */
func (mg *GromManger) LoadAppModels() {
	mg.PushTable(system.SysUser{})
	mg.PushTable(system.SysApi{})

}

func (mg *GromManger) PushTable(tb interface{}) {
	Tables = append(Tables, tb)
}

//@author: SliverHorn
//@function: Gorm
//@description: 初始化数据库并产生数据库全局变量
//@return: *gorm.DB

func (mg *GromManger) MigrateTable() {
	mg.LoadAppModels()
	st := utils.SliceType{SliceInterface: Tables}
	st.LoopAndCount()
	if st.LoopRet["interface"] == 0 {
		g.GV_LOG.Warn("未加载到模型数据")
		return
	}
	err := mg.AutoMigrate(Tables...)
	if err != nil {
		g.GV_LOG.Warn("数据表迁移失败")
	} else {
		g.GV_LOG.Info("数据表迁移成功")
	}
	return
}

func makeDsn(m types.Mysql) (string, error) {

	if m.Dbname == "" || m.Username == "" || m.Password == "" || m.Path == "" {
		mes := "数据库信息未配置"
		g.GV_LOG.Errorf(mes)
		return "", errors.New(mes)
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local&timeout=%s", m.Username, m.Password, m.Path, m.Dbname, "10s")
	return dsn, nil
}

func NewMySQL() *gorm.DB {
	m := g.GV_SERVER.Mysql
	dsn, err := makeDsn(m)
	if err != nil {
		os.Exit(1)
		return nil
	}

	mysqlConfig := mysql.Config{
		DSN:                       dsn,   // DSN data source name
		DefaultStringSize:         191,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据版本自动配置
	}
	db, err := gorm.Open(mysql.New(mysqlConfig), &gorm.Config{})
	if err != nil {
		g.GV_LOG.Errorf("数据库%s,连接失败,", m.Path)
		os.Exit(1)
		return nil
	} else {
		g.GV_LOG.Infof("数据库%s,连接成功,", m.Path)
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(m.MaxIdleConns)
		sqlDB.SetMaxOpenConns(m.MaxOpenConns)
		return db
	}
}
