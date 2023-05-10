package employee

import (
	"eirc.app/internal/v1/service/department"
	"eirc.app/internal/v1/service/employee"
	model "eirc.app/internal/v1/structure/employees"
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
	EmployeeService  employee.Service
	DeparmentService department.Service
}

func New(db *gorm.DB) Resolver {

	return &resolver{
		EmployeeService:  employee.New(db),
		DeparmentService: department.New(db),
	}
}
