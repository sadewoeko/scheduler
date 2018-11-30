package duler

type DulerService struct {
	Repo Repository
}

func (s *DulerService) Find(id int) (Duler, error) {
	// logic here
	return s.Repo.Find(id)
}

func (s *DulerService) FindAll() ([]Duler, error) {
	// logic here
	return s.Repo.FindAll()
}

func (s *DulerService) Store(duler Duler) (int64, error) {
	// logic here
	return s.Repo.Store(duler)
}
