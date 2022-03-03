package event

import (
	"github.com/GoGooliveryProviderAPI/helper"
	"gorm.io/gorm"
)

type Repository interface {
	GetEvents(params *getEventsRequest) ([]*Event, error)
	GetEventById(params *getEventByIDRequest) (*Event, error)
	InsertEvent(params *addEventRequest) (int32, error)
	UpdateEvent(params *updateEventRequest) (int32, error)
	CancelaEvent(params *cancelEventRequest) (int32, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{
		db: &gorm.DB{},
	}
}

func (repo *repository) GetEvents(params *getEventsRequest) ([]*Event, error) {
	var event []*Event

	results := repo.db.Table("goolivery_provider_global.events").Limit(10).Offset(5).Where("id_Pais = ? AND Cancel= ?", params.CountryId, false).Find(&event)
	helper.Catch(results.Error)
	return event, nil
}

func (repo *repository) GetEventById(params *getEventByIDRequest) (*Event, error) {
	var event Event

	repo.db.Where("id_Pais = ? AND id = ?", params.CountryId, params.EventID).Find(&event)

	return &event, nil
}

func (repo *repository) InsertEvent(params *addEventRequest) (int32, error) {

	event := Event{
		Description:         params.Description,
		IdStatus:            params.IdStatus,
		DtReg:               params.DtReg,
		Cancel:              params.Cancel,
		Success:             params.Success,
		Traffic:             params.Traffic,
		InCustommer:         params.InCustommer,
		InSupplier:          params.InSupplier,
		OnPicking:           params.OnPicking,
		OnPacking:           params.OnPacking,
		OnExpedition:        params.OnExpedition,
		Finisher:            params.Finisher,
		Going_to:            params.Going_to,
		InAttend:            params.InAttend,
		InLogisticOperation: params.InLogisticOperation,
	}
	result := repo.db.Create(&event)

	helper.Catch(result.Error)
	//id, _ := result.LastInsertId()
	id := result.Statement.CurDestIndex
	return int32(id), nil

}

func (repo *repository) UpdateEvent(params *updateEventRequest) (int32, error) {
	event := Event{
		ID:                  params.Id,
		Description:         params.Description,
		IdStatus:            params.IdStatus,
		DtReg:               params.DtReg,
		Cancel:              params.Cancel,
		Success:             params.Success,
		Traffic:             params.Traffic,
		InCustommer:         params.InCustommer,
		InSupplier:          params.InSupplier,
		OnPicking:           params.OnPicking,
		OnPacking:           params.OnPacking,
		OnExpedition:        params.OnExpedition,
		Finisher:            params.Finisher,
		Going_to:            params.Going_to,
		InAttend:            params.InAttend,
		InLogisticOperation: params.InLogisticOperation,
	}
	result := repo.db.Save(&event)
	helper.Catch(result.Error)
	//id, _ := result.LastInsertId()
	id := result.Statement.CurDestIndex
	//id, _ := repo.db.Get()
	return int32(id), nil
}

func (repo *repository) CancelaEvent(params *cancelEventRequest) (int32, error) {
	event := Event{
		Cancel: true,
	}
	result := repo.db.Save(&event)

	helper.Catch(result.Error)
	id := result.Statement.CurDestIndex
	return int32(id), nil
}
