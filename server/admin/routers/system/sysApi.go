package system

import "github.com/gin-gonic/gin"

type SysApi struct{}

func (this *SysApi) LoadRouter(rg *gin.RouterGroup) {
	//apiRouter := rg.Group("api").Use()
	//var apiRouterApi = v1.ApiGroupApp.SystemApiGroup.SystemApiApi
	{
		//apiRouter.POST("createApi", apiRouterApi.CreateApi)               // 创建Api
		//apiRouter.POST("deleteApi", apiRouterApi.DeleteApi)               // 删除Api
		//apiRouter.POST("getApiList", apiRouterApi.GetApiList)             // 获取Api列表
		//apiRouter.POST("getApiById", apiRouterApi.GetApiById)             // 获取单条Api消息
		//apiRouter.POST("updateApi", apiRouterApi.UpdateApi)               // 更新api
		//apiRouter.POST("getAllApis", apiRouterApi.GetAllApis)             // 获取所有api
		//apiRouter.DELETE("deleteApisByIds", apiRouterApi.DeleteApisByIds) // 删除选中api
	}
}
