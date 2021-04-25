package beer

type Reader interface {
	GetAll() ([]*Beer, error)
	Get(ID int) (*Beer, error)
}

type Writer interface {
	Store(b *Beer) error
	Update(b *Beer) error
	Remove(ID int) error
}

type UseCase interface {
	Reader
	Writer
}

type Service struct{}

//This function  will act as a constructor
func NewService() *Service {
	return &Service{}
}

func (s *Service) GetAll() ([]*Beer, error) {
	return nil, nil
}

func (s *Service) Get(ID int) (*Beer, error) {
	return nil, nil
}

func (s *Service) Store(b *Beer) error {
	return nil
}

func (s *Service) Update(b *Beer) error {
	return nil
}

func (s *Service) Remove(ID int) error {
	return nil
}
