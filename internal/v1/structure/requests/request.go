package requests

import (
	"time"

	model "eirc.app/internal/v1/structure"
)

// Table struct is database table struct
type Table struct {
	// 編號UUID
	RequestID string `gorm:"primaryKey;column:request_id;uuid_generate_v4()type:UUID;" json:"request_id,omitempty"`
	// 單號
	RequestCode string `gorm:"->;column:request_code;add_request_code()type:text;" json:"request_code,omitempty"` //->唯讀不寫 else:會是空白
	// 員工ID
	EmployeeID string `gorm:"column:employee_id;type:UUID;" json:"employee_id,omitempty"`
	// 部門ID
	DepartmentID string `gorm:"column:department_id;type:UUID;" json:"department_id,omitempty"`
	// 設備ID
	EquipmentID string `gorm:"column:equipment_id;type:UUID;" json:"equipment_id,omitempty"`
	// 數量
	Quanity int64 `gorm:"column:quanity;type:INT4;" json:"quanity,omitempty"`
	// 價格
	Price int64 `gorm:"column:price;type:INT4;" json:"price,omitempty"`
	// 創建時間
	CreatedTime time.Time `gorm:"column:created_time;type:TIMESTAMP;" json:"created_time"`
}

// Base struct is corresponding to table structure file
type Base struct {
	// 編號
	RequestID string `json:"request_id,omitempty"`
	// 單號
	RequestCode string `json:"request_code,omitempty"`
	// 員工ID
	EmployeeID string `json:"employee_id,omitempty"`
	// 部門ID
	DepartmentID string `json:"department_id,omitempty"`
	// 設備ID
	EquipmentID string `json:"equipment_id,omitempty"`
	// 數量
	Quanity int64 `json:"quanity,omitempty"`
	// 價格
	Price int64 `json:"price,omitempty"`
	// 創建時間
	CreatedTime time.Time `json:"created_time"`
}

// Single return structure file
type Single struct {
	// 編號
	RequestID string `json:"request_id,omitempty"`
	// 單號
	RequestCode string `json:"request_code,omitempty"`
	// 員工ID
	EmployeeID string `json:"employee_id,omitempty"`
	// 部門ID
	DepartmentID string `json:"department_id,omitempty"`
	// 設備ID
	EquipmentID string `json:"equipment_id,omitempty"`
	// 數量
	Quanity int64 `json:"quanity,omitempty"`
	// 價格
	Price int64 `json:"price,omitempty"`
	// 創建時間
	CreatedTime time.Time `json:"created_time"`
}

// Created struct is used to create
type Created struct {
	// 員工ID
	EmployeeID string `json:"employee_id" binding:"required" validate:"required"`
	// 部門ID
	DepartmentID string `json:"department_id" binding:"required,uuid4" validate:"required"`
	// 設備ID
	EquipmentID string `json:"equipment_id" binding:"required,uuid4" validate:"required"`
	// 數量
	Quanity int64 `json:"quanity,omitempty" binding:"required" validate:"required"`
	// 價格
	Price int64 `json:"price" binding:"required" validate:"required"`
}

// Field is structure file for search
type Field struct {
	// 編號
	RequestID string `json:"request_id,omitempty" binding:"omitempty" swaggerignore:"true"`
	// 單號
	RequestCode string `json:"request_code,omitempty" form:"request_code"`
	// 員工ID
	EmployeeID *string `json:"employee_id,omitempty" form:"employee_id" binding:"omitempty,uuid4"`
	// 部門ID
	DepartmentID *string `json:"department_id,omitempty" form:"department_id" binding:"omitempty,uuid4"`
	// 設備ID
	EquipmentID *string `json:"equipment_id,omitempty" form:"equipment_id" binding:"omitempty,uuid4"`
	// 數量
	Quanity *int64 `json:"quanity,omitempty" form:"quanity"`
	// 價格
	Price *int64 `json:"price,omitempty" form:"price"`
}

// Fields is the searched structure file (including pagination)
type Fields struct {
	Field
	model.InPage
}

// List is multiple return structure files
type List struct {
	Requests []*struct {
		// 編號
		RequestID string `json:"request_id,omitempty"`
		// 單號
		RequestCode string `json:"request_code,omitempty"`
		// 員工ID
		EmployeeID string `json:"employee_id,omitempty"`
		// 部門ID
		DepartmentID string `json:"department_id,omitempty"`
		// 設備ID
		EquipmentID string `json:"equipment_id,omitempty"`
		// 數量
		Quanity int64 `json:"quanity,omitempty"`
		// 價格
		Price int64 `json:"price,omitempty"`
	} `json:"requests"`
	model.OutPage
}

// Updated struct is used to update
type Updated struct {
	// 編號
	RequestID string `json:"request_id,omitempty" binding:"omitempty" swaggerignore:"true"`
	// 單號
	RequestCode *string `json:"request_code,omitempty"`
	// 員工ID
	EmployeeID *string `json:"employee_id,omitempty" binding:"omitempty,uuid4"`
	// 部門ID
	DepartmentID *string `json:"department_id,omitempty" binding:"omitempty,uuid4"`
	// 設備ID
	EquipmentID *string `json:"equipment_id,omitempty" binding:"omitempty,uuid4"`
	// 數量
	Quanity int64 `json:"quanity,omitempty"`
	// 價格
	Price int64 `json:"price,omitempty"`
}

// TableName sets the insert table name for this struct type
func (a *Table) TableName() string {
	return "requests"
}
