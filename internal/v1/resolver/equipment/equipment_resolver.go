package equipment

import (
	"encoding/json"
	"errors"

	"eirc.app/internal/pkg/code"
	"eirc.app/internal/pkg/log"
	"eirc.app/internal/pkg/util"
	equipmentModel "eirc.app/internal/v1/structure/equipments"
	"gorm.io/gorm"
)

func (r *resolver) Created(trx *gorm.DB, input *equipmentModel.Created) interface{} {
	defer trx.Rollback()
	// Todo 檢查重複
	equipment, err := r.EquipmentService.WithTrx(trx).Created(input)
	if err != nil {
		log.Error(err)

		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	trx.Commit()
	return code.GetCodeMessage(code.Successful, equipment.EquipmentID)
}

func (r *resolver) List(input *equipmentModel.Fields) interface{} {
	//input.IsDeleted = util.PointerBool(false)
	output := &equipmentModel.List{}
	output.Limit = input.Limit
	output.Page = input.Page
	quantity, equipments, err := r.EquipmentService.List(input)
	equipmentsByte, err := json.Marshal(equipments)
	if err != nil {
		log.Error(err)

		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	output.Pages = util.Pagination(quantity, output.Limit)
	err = json.Unmarshal(equipmentsByte, &output.Equipments)
	if err != nil {
		log.Error(err)

		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	return code.GetCodeMessage(code.Successful, output)
}

func (r *resolver) GetByID(input *equipmentModel.Field) interface{} {
	//input.IsDeleted = util.PointerBool(false)
	base, err := r.EquipmentService.GetByID(input)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {

			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)

		return code.GetCodeMessage(code.InternalServerError, err)
	}

	frontEquipment := &equipmentModel.Single{}
	equipmentsByte, _ := json.Marshal(base)
	err = json.Unmarshal(equipmentsByte, &frontEquipment)
	if err != nil {
		log.Error(err)

		return code.GetCodeMessage(code.InternalServerError, err)
	}

	return code.GetCodeMessage(code.Successful, frontEquipment)
}

func (r *resolver) Deleted(input *equipmentModel.Updated) interface{} {
	_, err := r.EquipmentService.GetByID(&equipmentModel.Field{EquipmentID: input.EquipmentID})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {

			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)

		return code.GetCodeMessage(code.InternalServerError, err)
	}

	err = r.EquipmentService.Deleted(input)
	if err != nil {
		log.Error(err)

		return code.GetCodeMessage(code.InternalServerError, err)
	}

	return code.GetCodeMessage(code.Successful, "Delete ok!")
}

func (r *resolver) Updated(input *equipmentModel.Updated) interface{} {
	equipment, err := r.EquipmentService.GetByID(&equipmentModel.Field{EquipmentID: input.EquipmentID})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {

			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)

		return code.GetCodeMessage(code.InternalServerError, err)
	}

	err = r.EquipmentService.Updated(input)
	if err != nil {
		log.Error(err)

		return code.GetCodeMessage(code.InternalServerError, err)
	}

	return code.GetCodeMessage(code.Successful, equipment.EquipmentID)
}
