package user

import (
	"github.com/GoGooliveryProviderAPI/helper"
	"gorm.io/gorm"
)

type Repository interface {
	GetUserById(params *getUserByIDRequest) (*User, error)
	InsertUser(params *addUserRequest) (int32, error)
	UpdateUser(params *updateUserRequest) (int32, error)
	CancelUser(params *cancelUserRequest) (int32, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{
		db: &gorm.DB{},
	}
}

func (repo *repository) GetUserById(params *getUserByIDRequest) (*User, error) {
	var user User

	repo.db.Where("id_Pais = ? AND id = ?", params.CountryId, params.UserID).Find(&user)

	return &user, nil
}

func (repo *repository) InsertUser(params *addUserRequest) (int32, error) {

	user := User{

		FirstName: params.FirstName,
		LastName:  params.LastName,
		DocState:  params.DocState,
		Login:     params.Login,
		Password:  params.Password,
		Email:     params.Email,
		CelPhone:  params.CelPhone,
		Country:   params.Country,
		DtReg:     params.DtReg,
		IdUser:    params.IdUser,
		Cancel:    false,
		IdCancel:  0,
	}
	result := repo.db.Create(&user)

	helper.Catch(result.Error)
	//id, _ := result.LastInsertId()
	id := result.Statement.CurDestIndex
	return int32(id), nil

}

func (repo *repository) UpdateUser(params *updateUserRequest) (int32, error) {
	user := User{
		Id:        params.Id,
		FirstName: params.FirstName,
		LastName:  params.LastName,
		DocState:  params.DocState,
		Login:     params.Login,
		Password:  params.Password,
		Email:     params.Email,
		CelPhone:  params.CelPhone,
		Country:   params.Country,
		DtReg:     params.DtReg,
		IdUser:    params.IdUser,
		Cancel:    params.Cancel,
		IdCancel:  0,
	}
	result := repo.db.Save(&user)
	helper.Catch(result.Error)
	//id, _ := result.LastInsertId()
	id := result.Statement.CurDestIndex
	//id, _ := repo.db.Get()
	return int32(id), nil
}

func (repo *repository) CancelUser(params *cancelUserRequest) (int32, error) {
	user := User{
		Cancel: true,
	}
	result := repo.db.Save(&user)

	helper.Catch(result.Error)
	id := result.Statement.CurDestIndex
	return int32(id), nil
}
