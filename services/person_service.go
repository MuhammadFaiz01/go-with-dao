package services

import (
	"go-dao/dao"
	"go-dao/models"
)

type PersonService struct {
	PersonDao dao.PersonDaoInterface
}

func NewPersonService(dao dao.PersonDaoInterface) *PersonService {
	return &PersonService{PersonDao: dao}
}

func (s *PersonService) CreatePerson(p *models.Person) error {
	return s.PersonDao.CreatePerson(p)
}

func (s *PersonService) UpdatePerson(id int, p *models.Person) error {
	return s.PersonDao.UpdatePerson(id, p)
}

func (s *PersonService) GetPersonByName(fullName string) ([]models.Person, error) {
	return s.PersonDao.GetPersonByName(fullName)
}

func (s *PersonService) GetAllPersons() ([]models.Person, error) {
	return s.PersonDao.GetAllPersons()
}

func (s *PersonService) DeletePerson(id int) error {
	return s.PersonDao.DeletePerson(id)
}
