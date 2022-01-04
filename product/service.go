package product

type Service interface {
	GetProductoById(param *getProductBYIDRequest) (*Product, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}

func (s *service) GetProductoById(param *getProductBYIDRequest) (*Product, error) {
	//Business logica
	return s.repo.GetProductoById(param.ProdutoID)
}
