package request

import (
	model "eirc.app/internal/v1/structure/requests"
)

func (e *entity) Created(input *model.Table) (err error) {
	err = e.db.Model(&model.Table{}).Create(&input).Error

	return err
}

func (e *entity) List(input *model.Fields) (amount int64, output []*model.Table, err error) {
	db := e.db.Model(&model.Table{})
	if input.DepartmentID != nil {
		db.Where("department_id = ?", input.DepartmentID)
	}
	if input.EmployeeID != nil {
		db.Where("employee_id = ?", input.EmployeeID)
	}
	if input.EquipmentID != nil {
		db.Where("equipment_id = ?", input.EquipmentID)
	}
	// if input.Price != nil {
	// 	db.Where("price = ?", input.Price)
	// }

	// if input.Quanity != nil {
	// 	db.Where("name like %?%", *input.Quanity)
	// }

	return amount, output, err
}

func (e *entity) GetByID(input *model.Field) (output *model.Table, err error) {
	db := e.db.Model(&model.Table{}).Where("request_id = ?", input.RequestID)

	err = db.First(&output).Error

	return output, err
}

func (e *entity) Deleted(input *model.Table) (err error) {
	err = e.db.Model(&model.Table{}).Delete(&input).Error

	return err
}

func (e *entity) Updated(input *model.Table) (err error) {
	err = e.db.Model(&model.Table{}).Save(&input).Error

	return err
}
