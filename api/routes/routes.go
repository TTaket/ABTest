package routes

import (
	"ABTest/api/handlers"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// SetupRouter 设置所有路由
func SetupRouter() *gin.Engine {
	r := gin.Default()

	// 添加Swagger文档路由
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// API路由组
	api := r.Group("/api")
	{
		// 实验相关路由
		setupExperimentRoutes(api)

		// 用户相关路由
		setupUserbRoutes(api)

		// 层相关路由
		setupLayerRoutes(api)

		// 动态绑定的路由
		setupDynamicRoutes(api)
	}

	return r
}

// setupExperimentRoutes 设置实验相关路由
func setupExperimentRoutes(router *gin.RouterGroup) {
	experimentHandler := handlers.NewExperimentHandler()
	experiments := router.Group("/experiments")
	{
		// 按照Swagger定义修正路由
		experiments.GET("/getinfo", experimentHandler.GetExperiment)
		experiments.POST("/create-experiment", experimentHandler.CreateExperiment)
		experiments.DELETE("/delete-experiment", experimentHandler.DeleteExperiment)
		experiments.PUT("/baseinfo/update-info", experimentHandler.UpdateExperimentBaseInfo)
		experiments.POST("/groups/create-groups", experimentHandler.AddExperimentGroup)
		experiments.DELETE("/delete-gropup", experimentHandler.DeleteExperimentGroup)
	}
}

// setupUserbRoutes 设置用户相关路由
func setupUserbRoutes(router *gin.RouterGroup) {
	userbHandler := handlers.NewUserbHandler()
	users := router.Group("/users")
	{
		// 按照Swagger定义修正路由
		users.POST("/info/getinfo", userbHandler.GetUserInfo)
		users.POST("/adduser", userbHandler.AddUser)
		users.PUT("/update-userinfo", userbHandler.UpdateUser)
		users.DELETE("/delete-user", userbHandler.DeleteUser)

		// 批量操作相关路由修正
		users.POST("/batch-add-user", userbHandler.BatchAddUser)
		users.DELETE("/batch-del-user", userbHandler.BatchDeleteUser)
		users.PUT("/batch-update-user", userbHandler.BatchUpdateUser)
		users.POST("/batch-get-info", userbHandler.BatchGetUserInfo)
	}

	// 流量分散路由
	traffic := router.Group("/traffic")
	{
		traffic.POST("/scatter", userbHandler.ScatterTraffic)
	}
}

// setupLayerRoutes 设置层相关路由
func setupLayerRoutes(router *gin.RouterGroup) {
	layerHandler := handlers.NewLayerHandler()
	layers := router.Group("/layers")
	{
		// 按照Swagger定义修正路由
		layers.GET("/get-all", layerHandler.ListLayers)
		layers.GET("/getinfo", layerHandler.GetLayer)
		layers.POST("/create", layerHandler.CreateLayer)
		layers.PUT("/updateinfo", layerHandler.UpdateLayer)
		layers.DELETE("/delete-layer", layerHandler.DeleteLayer)
	}
}

// setupDynamicRoutes 设置动态绑定相关路由
func setupDynamicRoutes(router *gin.RouterGroup) {
	dynamicHandler := handlers.NewDynamicHandler()

	dynamics := router.Group("/dynamic")
	{
		// 按照Swagger定义修正Layer-User绑定路由
		dynamics.POST("/layer-user-bindings/create-bind", dynamicHandler.CreateLayerUserBinding)
		dynamics.GET("/layer-user-bindings/get-bind", dynamicHandler.GetLayerUserBinding)
		dynamics.DELETE("/layer-user-bindings/delete-bind", dynamicHandler.DeleteLayerUserBinding)

		// 按照Swagger定义修正Experiment-User绑定路由
		dynamics.POST("/experiment-user-bindings/create-bind", dynamicHandler.CreateExperimentUserBinding)
		dynamics.GET("/experiment-user-bindings/get-bind", dynamicHandler.GetExperimentUserBinding)
		dynamics.DELETE("/experiment-user-bindings/delete-bind", dynamicHandler.DeleteExperimentUserBinding)

		// 按照Swagger定义修正Experiment-Layer绑定路由
		dynamics.POST("/experiment-layer-bindings/create-bind", dynamicHandler.CreateExperimentLayerBinding)
		dynamics.GET("/experiment-layer-bindings/get-bind", dynamicHandler.GetExperimentLayerBinding)
		dynamics.DELETE("/experiment-layer-bindings/delete-bind", dynamicHandler.DeleteExperimentLayerBinding)

		// 其他路由
		dynamics.GET("/layer-experiments/get-all-bylayer", dynamicHandler.GetLayerExperiments)
	}
}
