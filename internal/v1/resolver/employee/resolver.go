package employee

import (
	"eirc.app/internal/v1/service/department"
	//"eirc.app/internal/v1/service/company"
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
	DepartmentServer department.Service
	//CompanyService company.Service
}

func New(db *gorm.DB) Resolver {

	return &resolver{
		DepartmentServer: department.New(db),
		//CompanyService: company.New(db),
	}
}
