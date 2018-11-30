package duler

import (
	"fmt"

	"github.com/sadewoeko/scheduler/driver"
)

type DulerRepository struct {
	driver.DBHandler
}

func (r *DulerRepository) Find(id int) (Duler, error) {
	row, err := r.Query(fmt.Sprintf("SELECT * FROM dulers WHERE id = '%d'", id))

	if err != nil {
		return Duler{}, err
	}

	var duler Duler
	row.Next()
	row.Scan(&duler.Id, &duler.Name)

	return duler, nil
}

func (r *DulerRepository) FindAll() ([]Duler, error) {

	rows, err := r.Query(fmt.Sprintf("SELECT * FROM dulers"))

	if err != nil {
		return []Duler{}, err
	}

	var (
		duler  Duler
		dulers []Duler
	)
	for rows.Next() {
		rows.Scan(&duler.Id, &duler.Name)

		dulers = append(dulers, duler)

	}

	return dulers, nil
}

func (r *DulerRepository) Store(duler Duler) (int64, error) {
	res, err := r.Execute(fmt.Sprintf("INSERT INTO dulers(name) VALUES ('%s')", duler.Name))

	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}
