package request

import (
	//"eirc.app/internal/v1/entity/equipment"
	//"eirc.app/internal/v1/presenter/employee"
	"eirc.app/internal/v1/service/department"
	"eirc.app/internal/v1/service/employee"
	"eirc.app/internal/v1/service/equipment"
	"eirc.app/internal/v1/service/request"

	model "eirc.app/internal/v1/structure/requests"
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
	RequestService    request.Service
	DepartmentService department.Service
	EmployeeService   employee.Service
	EquipmentService  equipment.Service
}

func New(db *gorm.DB) Resolver {

	return &resolver{
		RequestService:    request.New(db),
		DepartmentService: department.New(db),
		EmployeeService:   employee.New(db),
		EquipmentService:  equipment.New(db),
	}
}
