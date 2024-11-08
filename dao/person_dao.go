package dao

import (
	"context"
	"go-dao/models"

	"github.com/jackc/pgx/v5"
)

type PersonDao struct {
	DB *pgx.Conn
}

func NewPersonDao(db *pgx.Conn) *PersonDao {
	return &PersonDao{DB: db}
}

func (dao *PersonDao) CreatePerson(p *models.Person) error {
	_, err := dao.DB.Exec(context.Background(), "INSERT INTO person (full_name, age, birth_date, address) VALUES ($1, $2, $3, $4)", p.FullName, p.Age, p.BirthDate, p.Address)
	return err
}

func (dao *PersonDao) UpdatePerson(id int, p *models.Person) error {
	_, err := dao.DB.Exec(context.Background(), "UPDATE person SET full_name=$1, age=$2, birth_date=$3, address=$4 WHERE id=$5", p.FullName, p.Age, p.BirthDate, p.Address, id)
	return err
}

func (dao *PersonDao) GetPersonByName(fullName string) ([]models.Person, error) {
	rows, err := dao.DB.Query(context.Background(), "SELECT id, full_name, age, birth_date, address FROM person WHERE full_name ILIKE $1", "%"+fullName+"%")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var persons []models.Person
	for rows.Next() {
		var p models.Person
		if err := rows.Scan(&p.ID, &p.FullName, &p.Age, &p.BirthDate, &p.Address); err != nil {
			return nil, err
		}
		persons = append(persons, p)
	}
	return persons, nil
}

func (dao *PersonDao) GetAllPersons() ([]models.Person, error) {
	rows, err := dao.DB.Query(context.Background(), "SELECT id, full_name, age, birth_date, address FROM person")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var persons []models.Person
	for rows.Next() {
		var p models.Person
		if err := rows.Scan(&p.ID, &p.FullName, &p.Age, &p.BirthDate, &p.Address); err != nil {
			return nil, err
		}
		persons = append(persons, p)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return persons, nil
}

func (dao *PersonDao) DeletePerson(id int) error {
	_, err := dao.DB.Exec(context.Background(), "DELETE FROM person WHERE id=$1", id)
	return err
}
