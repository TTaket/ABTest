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
		experiments.GET("/:id", experimentHandler.GetExperiment)
		experiments.POST("", experimentHandler.CreateExperiment)
		experiments.DELETE("/:id", experimentHandler.DeleteExperiment)
	}
}

// setupUserbRoutes 设置用户相关路由
func setupUserbRoutes(router *gin.RouterGroup) {
	userbHandler := handlers.NewUserbHandler()
	users := router.Group("/users")
	{
		users.POST("/info", userbHandler.GetUserInfo)
		users.POST("", userbHandler.AddUser)
		users.PUT("", userbHandler.UpdateUser)
		users.DELETE("/:id", userbHandler.DeleteUser)
	}
}

// setupLayerRoutes 设置层相关路由
func setupLayerRoutes(router *gin.RouterGroup) {
	layerHandler := handlers.NewLayerHandler()
	layers := router.Group("/layers")
	{
		layers.GET("", layerHandler.ListLayers)
		layers.GET("/:id", layerHandler.GetLayer)
		layers.POST("", layerHandler.CreateLayer)
		layers.PUT("", layerHandler.UpdateLayer)
		layers.DELETE("/:id", layerHandler.DeleteLayer)
	}
}

// RegisterRoutes 注册路由
func setupDynamicRoutes(router *gin.RouterGroup) {
	dynamicHandler := handlers.NewDynamicHandler()

	dynamics := router.Group("/api/dynamic")
	{
		// Layer-User binding 路由
		dynamics.POST("/layer-user", dynamicHandler.CreateLayerUserBinding)
		dynamics.GET("/layer-user", dynamicHandler.GetLayerUserBinding)
		dynamics.DELETE("/layer-user", dynamicHandler.DeleteLayerUserBinding)

		// Experiment-User binding 路由
		dynamics.POST("/experiment-user", dynamicHandler.CreateExperimentUserBinding)
		dynamics.GET("/experiment-user", dynamicHandler.GetExperimentUserBinding)
		dynamics.DELETE("/experiment-user", dynamicHandler.DeleteExperimentUserBinding)

		// Experiment-Layer binding 路由
		dynamics.POST("/experiment-layer", dynamicHandler.CreateExperimentLayerBinding)
		dynamics.GET("/experiment-layer", dynamicHandler.GetExperimentLayerBinding)
		dynamics.DELETE("/experiment-layer", dynamicHandler.DeleteExperimentLayerBinding)

		// 其他路由
		dynamics.GET("/layer-experiments", dynamicHandler.GetLayerExperiments)
	}
}
