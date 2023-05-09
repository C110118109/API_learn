package equipment

import (
	"eirc.app/internal/v1/service/equipment"
	model "eirc.app/internal/v1/structure/equipments"
	"gorm.io/gorm"
)

type Resolver interface {
	Created(trx *gorm.DB, input *model.Created) interface{}
	List(input *model.Fields) interface{}
	GetByID(input *model.Field) interface{}
	Deleted(input *model.Updated) interface{}
	Updated(input *model.Updated) interface{}
}

type resolver struct {
	EquipmentService equipment.Service
}

func New(db *gorm.DB) Resolver {
	return &resolver{
		EquipmentService: equipment.New(db),
	}
}
