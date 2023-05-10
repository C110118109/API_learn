package equipment

import (
	"eirc.app/internal/v1/middleware"
	"eirc.app/internal/v1/presenter/equipment"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetRoute(route *gin.Engine, db *gorm.DB) *gin.Engine {
	controller := equipment.New(db)
	v10 := route.Group("authority").Group("v1.0").Group("equipments")
	{
		v10.POST("", middleware.Transaction(db), controller.Created)
		v10.GET("", controller.List)
		v10.GET(":equipmentID", controller.GetByID)
		v10.DELETE(":equipmentID", controller.Delete)
		v10.PATCH(":equipmentID", controller.Updated)
	}

	return route
}
