package user

type Service interface {
	GetUserById(params *getUserByIDRequest) (*User, error)
	InsertUser(params *addUserRequest) (int32, error)
	UpdateUser(params *updateUserRequest) (int32, error)
	CancelUser(params *cancelUserRequest) (int32, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s *service) GetUserById(params *getUserByIDRequest) (*User, error) {
	return s.repo.GetUserById(params)
}

func (s *service) InsertUser(params *addUserRequest) (int32, error) {
	return s.repo.InsertUser(params)
}

func (s *service) UpdateUser(params *updateUserRequest) (int32, error) {
	return s.repo.UpdateUser(params)
}

func (s *service) CancelUser(params *cancelUserRequest) (int32, error) {
	return s.repo.CancelUser(params)
}
