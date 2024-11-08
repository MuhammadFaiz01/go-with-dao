package services

import (
	"go-dao/dao"
	"go-dao/models"
)

type KelasService struct {
	KelasDao *dao.KelasDao
}

func NewKelasService(dao *dao.KelasDao) *KelasService {
	return &KelasService{KelasDao: dao}
}

func (s *KelasService) CreateKelas(k *models.Kelas) error {
	return s.KelasDao.CreateKelas(k)
}

func (s *KelasService) UpdateKelas(id int, k *models.Kelas) error {
	return s.KelasDao.UpdateKelas(id, k)
}

func (s *KelasService) GetKelasByName(namaKelas string) ([]models.Kelas, error) {
	return s.KelasDao.GetKelasByName(namaKelas)
}

func (s *KelasService) GetAllKelas() ([]models.Kelas, error) {
	return s.KelasDao.GetAllKelas()
}

func (s *KelasService) DeleteKelas(id int) error {
	return s.KelasDao.DeleteKelas(id)
}
