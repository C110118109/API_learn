package equipment

import (
	"eirc.app/internal/v1/resolver/equipment"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Presenter interface {
	Created(ctx *gin.Context)
	List(ctx *gin.Context)
	GetByID(ctx *gin.Context)
	Updated(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type presenter struct {
	EquipmentResolver equipment.Resolver
}

func New(db *gorm.DB) Presenter {
	return &presenter{
		EquipmentResolver: equipment.New(db),
	}
}
