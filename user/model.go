package user

import (
	"github.com/golang/protobuf/ptypes/timestamp"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Id        int64               `json:"id"`
	FirstName string              `json:"firstName"`
	LastName  string              `json:"lastName"`
	DocState  string              `json:"docState"`
	Login     string              `json:"login"`
	Password  string              `json:"password"`
	Email     string              `json:"email"`
	CelPhone  string              `json:"celphone"`
	Country   string              `json:"country"`
	DtReg     timestamp.Timestamp `json:"dtReg"`
	IdUser    int64               `json:"idUser"`
	Cancel    bool                `json:"cancel"`
	IdCancel  int64               `json:"idCancel"`
}
