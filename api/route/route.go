package route

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/june21x/project-a-api/docs"
	initialization "github.com/june21x/project-a-api/init"
	"github.com/june21x/project-a-api/util"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Init(init *initialization.Initialization) *gin.Engine {

	router := gin.New()

	docs.SwaggerInfo.BasePath = "/api/v1"

	trans := util.InitErrorTranslator()
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		util.RegisterErrorTranslation(trans, v)
	}

	// router.Use(gin.Logger())
	router.Use(gin.Recovery())

	api := router.Group("/api")
	v1 := api.Group("/v1")

	{
		v1.GET("/ping", init.HealthCheckCtrl.HealthCheck)
	}
	{
		// 	auth := v1.Group("/auth")

		// 	auth.POST("/login", controller.Login)
		// 	auth.POST("/logout", controller.Logout)
	}
	{
		player := v1.Group("/players")
		player.POST("", init.PlayerCtrl.RegisterPlayer)
		player.GET("", init.PlayerCtrl.GetPlayers)
		player.GET("/:uuid", init.PlayerCtrl.GetPlayer)
		//  player.PUT("/:playerID", init.PlayerCtrl.UpdatePlayerData)
		//  player.DELETE("/:playerID", init.PlayerCtrl.DeletePlayer)
	}
	{
		worldMap := v1.Group("/world-map")

		area := worldMap.Group("/areas")
		area.POST("", init.AreaCtrl.CreateArea)
		area.GET("", init.AreaCtrl.GetAreas)
		area.GET("/:uuid", init.AreaCtrl.GetArea)
	}
	{
		router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	}

	return router
}
