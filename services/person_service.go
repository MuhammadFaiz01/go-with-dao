package services

import (
	"go-dao/dao"
	"go-dao/models"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/360EntSecGroup-Skylar/excelize"
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

func (s *PersonService) ExportPersons(filePath string) error {

	persons, err := s.PersonDao.GetAllPersons()
	if err != nil {
		return err
	}

	f := excelize.NewFile()
	sheet := "Sheet1"
	f.NewSheet(sheet)

	f.SetCellValue(sheet, "A1", "ID")
	f.SetCellValue(sheet, "B1", "Full Name")
	f.SetCellValue(sheet, "C1", "Age")
	f.SetCellValue(sheet, "D1", "Birth Date")
	f.SetCellValue(sheet, "E1", "Address")

	for i, person := range persons {
		row := strconv.Itoa(i + 2)
		f.SetCellValue(sheet, "A"+row, person.ID)
		f.SetCellValue(sheet, "B"+row, person.FullName)
		f.SetCellValue(sheet, "C"+row, person.Age)
		f.SetCellValue(sheet, "D"+row, person.BirthDate.Format("2006-01-02"))
		f.SetCellValue(sheet, "E"+row, person.Address)
	}

	dir := filepath.Dir(filePath)
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return err
	}

	if err := f.SaveAs(filePath); err != nil {
		return err
	}

	return nil
}

func (s *PersonService) ImportFromExcel(filePath string) error {
	f, err := excelize.OpenFile(filePath)
	if err != nil {
		return err
	}

	rows := f.GetRows("Sheet1")
	for i, row := range rows {

		if i == 0 {
			continue
		}

		var person models.Person
		if len(row) > 1 {
			person.FullName = row[1]
		}
		if len(row) > 2 {
			age, _ := strconv.Atoi(row[2])
			person.Age = age
		}
		if len(row) > 3 {
			person.BirthDate, _ = time.Parse("2006-01-02", row[3])
		}
		if len(row) > 4 {
			person.Address = row[4]
		}

		if err := s.PersonDao.CreatePerson(&person); err != nil {
			return err
		}
	}

	return nil
}
