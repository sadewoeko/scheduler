package duler

type Reader interface {
	Find(id int) (Duler, error)
	FindAll() ([]Duler, error)
}

type Writer interface {
	Store(duler Duler) (int64, error)
}

type Repository interface {
	Reader
	Writer
}

type UseCase interface {
	Reader
	Writer
}
