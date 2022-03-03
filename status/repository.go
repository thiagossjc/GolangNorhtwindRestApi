package status

import (
	"github.com/GolangNorhtwindRestApi/helper"
	"gorm.io/gorm"
)

type Repository interface {
	GetStatus(params *getStatusRequest) ([]*Status, error)
	GetStatusById(params *getStatusByIDRequest) (*Status, error)
	InsertStatu(params *addStatusRequest) (int32, error)
	UpdateStatu(params *updateStatusRequest) (int32, error)
	CancelaStatu(params *cancelStatusRequest) (int32, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{
		db: &gorm.DB{},
	}
}

func (repo *repository) GetStatus(params *getStatusRequest) ([]*Status, error) {
	var status []*Status

	// SELECT * FROM users OFFSET 5 LIMIT 10;

	results := repo.db.Table("goolivery_provider_global.status").Limit(10).Offset(5).Where("id_Pais = ? AND Cancel= ?", params.CountryId, false).Find(&status)
	helper.Catch(results.Error)
	return status, nil
}

func (repo *repository) GetStatusById(params *getStatusByIDRequest) (*Status, error) {
	var status Status

	repo.db.Where("id_Pais = ? AND id >= ?", params.CountryId, params.StatusID).Find(&status)

	return &status, nil
}

func (repo *repository) InsertStatu(params *addStatusRequest) (int32, error) {

	status := Status{Description: params.Description,
		UserReg:             params.UsereReg,
		Success:             params.Success,
		Traffic:             params.Traffic,
		InCustomer:          params.InCustomer,
		InSupplier:          params.InSupplier,
		OnPicking:           params.OnPicking,
		OnPacking:           params.OnPacking,
		OnExpedition:        params.OnExpedition,
		GoingTo:             params.GoingTo,
		InAttenDance:        params.InAttend,
		InLogisticOperation: params.InLogisticOperation,
		IdCountry:           params.IdCountry,
	}
	result := repo.db.Create(&status)

	helper.Catch(result.Error)
	//id, _ := result.LastInsertId()
	id := result.Statement.CurDestIndex
	return int32(id), nil

}

func (repo *repository) UpdateStatu(params *updateStatusRequest) (int32, error) {
	status := Status{
		Id:                  params.Id,
		UserReg:             params.UsereReg,
		Success:             params.Success,
		Traffic:             params.Traffic,
		InCustomer:          params.InCustomer,
		InSupplier:          params.InSupplier,
		OnPicking:           params.OnPicking,
		OnPacking:           params.OnPacking,
		OnExpedition:        params.OnExpedition,
		GoingTo:             params.GoingTo,
		InAttenDance:        params.InAttend,
		InLogisticOperation: params.InLogisticOperation,
		IdCountry:           params.IdCountry,
	}
	result := repo.db.Save(&status)
	helper.Catch(result.Error)
	//id, _ := result.LastInsertId()
	id := result.Statement.CurDestIndex
	//id, _ := repo.db.Get()
	return int32(id), nil
}

func (repo *repository) CancelaStatu(params *cancelStatusRequest) (int32, error) {
	status := Status{
		Cancel: true,
	}
	result := repo.db.Save(&status)

	helper.Catch(result.Error)
	id := result.Statement.CurDestIndex
	return int32(id), nil
}
