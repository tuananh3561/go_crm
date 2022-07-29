package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/streadway/amqp"
	"github.com/tuananh3561/go_crm/app/config"
	"github.com/tuananh3561/go_crm/app/controller"
	"github.com/tuananh3561/go_crm/app/middleware"
	"github.com/tuananh3561/go_crm/app/repository"
	"github.com/tuananh3561/go_crm/app/service"
	"github.com/tuananh3561/go_crm/app/usecase"
	"net/http"
)

// Router is routing settings
func Router(r *gin.Engine, db config.Database, channelRabbitMQ *amqp.Channel) {

	var (
		roleRepo = repository.NewRoleRepository(db.MysqlAuth)

		jwtService = service.NewJWTService()
		//mediaService           = service.NewMediaService()
		historyActivityService = service.NewHistoryActivityService()

		//textService            = service.NewTextService()
		//authConnectService     = connect_service.NewCrmConnectService()
		//productConnectService  = connect_service.NewProductConnectService()
		//webConnectService      = connect_service.NewWebConnectService()
		//deleteCacheService     = connect_service.NewDeleteCacheService()
		//

		role = usecase.NewRole(roleRepo, historyActivityService)

		roleController = controller.NewRoleController(role)
		//translateController   = controller.NewTranslateController(translateRepo)
	)

	r.Use(middleware.CORSMiddleware())

	routesWeb := r.Group("")
	{
		routesWeb.GET("/", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{
				"message": "Service crm",
			})
		})
	}

	routesApi := r.Group("api", middleware.AuthorizeJWT(jwtService))
	{
		routesApi.GET("/", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{
				"message": "Service crm api",
			})
		})
		//role
		routesApi.GET("/role/list", roleController.List)
		routesApi.GET("/role/create", roleController.Create)
		routesApi.GET("/role/update", roleController.Update)
		routesApi.GET("/role/update_status", roleController.UpdateStatus)

		////Translate
		//routesApi.GET("/get-list-translate-by-params", translateController.GetListTranslateByParams)
		//routesApi.POST("/add-translate-language", translateController.AddTranslateLanguage)

	}
}
