package http

import (
	"strings"

	"github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"

	"architecture_go/pkg/type/logger"
)

func (d *Delivery) initRouter() *gin.Engine {

	if viper.GetBool("IS_PRODUCTION") {
		switch strings.ToUpper(strings.TrimSpace(viper.GetString("LOG_LEVEL"))) {
		case "DEBUG":
			gin.SetMode(gin.ReleaseMode)
		default:
			gin.SetMode(gin.ReleaseMode)
		}
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	var router = gin.New()

	router.Use(Tracer())

	// Logs all panic to error log
	//   - stack means whether output the stack info.
	router.Use(ginzap.RecoveryWithZap(logger.GetLogger(), true))

	d.routerClient(router.Group("/"))

	d.routerAdmin(router.Group("/admin"))

	return router
}

func (d *Delivery) routerClient(router *gin.RouterGroup) {
	router.GET("/", d.ShowVideos)
	router.POST("/add", d.AddVideo)
	router.DELETE("/:id", d.RemoveVideo)
}

func (d *Delivery) routerAdmin(router *gin.RouterGroup) {
	router.GET("/admin/", d.ShowGroups)
	router.POST("/admin/add", d.CreateGroup)
	router.DELETE("/admin/:id", d.DeleteGroup)
}
