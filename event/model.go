package event

import (
	"github.com/golang/protobuf/ptypes/timestamp"
	"gorm.io/gorm"
)

type Event struct {
	gorm.Model
	ID                  int32               `json:"id"`
	IdStatus            int32               `json:"idStatus"`
	Description         string              `json:"description"`
	UserReg             int64               `json:"userReg"`
	DtReg               timestamp.Timestamp `json:"dtReg"`
	Cancel              bool                `json:"cancel"`
	Success             bool                `json:"success"`
	Traffic             bool                `json:"traffic"`
	InCustommer         bool                `json:"inCustomer"`
	InSupplier          bool                `json:"inSupplier"`
	OnPicking           bool                `json:"inPicking"`
	OnPacking           bool                `json:"inPacking"`
	OnExpedition        bool                `json:"onExpedition"`
	Finisher            bool                `json:"finisher"`
	Going_to            string              `json:"goingTo"`
	IdPais              int32               `json:"idPais"`
	InLogisticOperation bool                `json:"inLogOper"`
	InAttend            bool                `json:"inAttend"`
}

type EventList struct {
	Data []*Event `json:"data"`
}
