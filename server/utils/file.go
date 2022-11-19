package utils

import (
	"os"
)

/*
和文件操作相关方法
*/

//路径不存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		err = OutEr("文件%s,不存在", path)
		return false, err
	}
	return false, err
}

//创建目录
func CreateDir(dirs ...string) (err error) {
	for _, v := range dirs {
		exist, err := PathExists(v)
		if err != nil {
			return err
		}
		if !exist {
			if err := os.MkdirAll(v, os.ModePerm); err != nil {
				return err
			}
		}
	}
	return err
}

//异步日志分割
//func WriteSyncer() (zapcore.WriteSyncer, error) {
//	fileWriter, err := rotatelogfile.New(
//		//path.Join(admin.GV_SERVER.Zap.Director, "%Y-%m-%d.log"),
//		//rotatelogfile.WithLinkName(admin.GV_SERVER.Zap.LinkName),
//		rotatelogfile.WithMaxAge(7*24*time.Hour),
//		rotatelogfile.WithRotationTime(24*time.Hour),
//	)
//	if admin.GV_SERVER.Zap.LogInConsole {
//		return zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(fileWriter)), err
//	}
//	return zapcore.AddSync(fileWriter), err
//}
