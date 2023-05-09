package departments

import (
	"time"

	model "eirc.app/internal/v1/structure"
)

// Department struct is a row record of the companies table in the invoice database
// Table struct is database table struct
type Table struct {
	// 部門編號
	DepartmentID string `gorm:"primaryKey;uuid_generate_v4();column:department_id;type:uuid;" json:"department_id,omitempty"`
	// 部門名稱
	Name string `gorm:"column:name;type:TEXT;" json:"name,omitempty"`
	// 創建時間
	CreatedTime time.Time `gorm:"column:created_time;type:TIMESTAMP;" json:"created_time"`
}

// Base struct is corresponding to table structure file
type Base struct {
	// 部門編號
	DepartmentID string `json:"department_id,omitempty"`
	// 部門名稱
	Name string `json:"name,omitempty"`
	// 創建時間
	CreatedTime time.Time `json:"created_time"`
}

// Single return structure file
type Single struct {
	// 部門編號
	DepartmentID string `json:"department_id,omitempty"`
	// 部門名稱
	Name string `json:"name,omitempty"`
	// 創建時間
	CreatedTime time.Time `json:"created_time"`
}

// Created struct is used to create
type Created struct {
	// 部門名稱
	Name string `json:"name,omitempty" binding:"required" validate:"required"`

}

// Updated struct is used to update
type Updated struct {
	// 部門編號
	DepartmentID string `json:"department_id,omitempty" binding:"omitempty,uuid4" swaggerignore:"true"`
	// 部門名稱
	Name string `json:"name,omitempty"`
}

// Field is structure file for search
type Field struct {
	// 部門編號
	DepartmentID string `json:"department_id,omitempty"  binding:"omitempty,uuid4" swaggerignore:"true"`
	// 部門名稱
	Name *string `json:"name,omitempty" form:"name"`
}

// Fields is the searched structure file (including pagination)
type Fields struct {
	Field
	model.InPage
}

// List is multiple return structure files
type List struct {
	Departments []*struct {
		// 部門編號
		DepartmentID string `json:"department_id,omitempty"`
		// 部門名稱
		Name string `json:"name,omitempty"`
		// 創建時間
		CreatedTime time.Time `json:"created_time"`
		
	} `json:"departments"`
	model.OutPage
}

// TableName sets the insert table name for this struct type
func (c *Table) TableName() string {
	return "departments"
}
