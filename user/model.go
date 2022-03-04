package user

type User struct {
	Id        int64  `json:"id" gorm:"primary_key"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	DocState  string `json:"docState"`
	Login     string `json:"login"`
	Password  string `json:"password"`
	Email     string `json:"email"`
	CelPhone  string `json:"celphone"`
	Country   string `json:"country"`
	IdUser    int64  `json:"idUser"`
	Cancel    bool   `json:"cancel"`
	IdCancel  int64  `json:"idCancel"`
}
