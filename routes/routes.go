package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/streadway/amqp"
	"github.com/tuananh3561/go_crm/app/config"
	"github.com/tuananh3561/go_crm/app/controller"
	"github.com/tuananh3561/go_crm/app/job"
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

		role = usecase.NewRole(roleRepo, historyActivityService)

		roleController = controller.NewRoleController(role)
	)

	r.Use(middleware.CORSMiddleware())

	routesWeb := r.Group("")
	{
		routesWeb.GET("/", func(context *gin.Context) {
			job.PublishInsertLogMongo()
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
	}
}
