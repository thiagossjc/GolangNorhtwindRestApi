package status

import (
	"gorm.io/gorm"
)

type Status struct {
	gorm.Model
	Id                  int32  `json:"id"`
	Description         string `json:"description"`
	UserReg             int64  `json:"userReg"`
	DateReg             string `json:"dataReg"`
	Cancel              bool   `json:"cancel"`
	Success             bool   `json:"success"`
	Traffic             bool   `json:"traffic"`
	InCustomer          bool   `json:"inCustomer"`
	InSupplier          bool   `json:"inSupplier"`
	OnPicking           bool   `json:"onPicking"`
	OnPacking           bool   `json:"onPacking"`
	OnExpedition        bool   `json:"onExpedition"`
	GoingTo             string `json:"goigTo"`
	IdCountry           int32  `json:"idCountry"`
	InAttenDance        bool   `json:"inAttend"`
	InLogisticOperation bool   `json:"inLogisticOperation"`
}

type StatusList struct {
	Data []*Status `json:"data"`
	//	TotalRecords int64       `json:"totalRecords"`
}
