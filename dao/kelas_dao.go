package dao

import (
	"context"
	"go-dao/models"

	"github.com/jackc/pgx/v5"
)

type KelasDao struct {
	DB *pgx.Conn
}

func NewKelasDao(db *pgx.Conn) *KelasDao {
	return &KelasDao{DB: db}
}

func (dao *KelasDao) CreateKelas(k *models.Kelas) error {
	_, err := dao.DB.Exec(context.Background(), "INSERT INTO kelas (nama_kelas, tingkat, id_person) VALUES ($1, $2, $3)", k.NamaKelas, k.Tingkat, k.IdPerson)
	return err
}

func (dao *KelasDao) UpdateKelas(id int, k *models.Kelas) error {
	_, err := dao.DB.Exec(context.Background(), "UPDATE kelas SET nama_kelas=$1, tingkat=$2, id_person=$3 WHERE id=$4", k.NamaKelas, k.Tingkat, k.IdPerson, id)
	return err
}

func (dao *KelasDao) GetKelasByName(namaKelas string) ([]models.Kelas, error) {
	rows, err := dao.DB.Query(context.Background(), "SELECT id, nama_kelas, tingkat, id_person FROM kelas WHERE nama_kelas ILIKE $1", "%"+namaKelas+"%")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var kelas []models.Kelas
	for rows.Next() {
		var k models.Kelas
		if err := rows.Scan(&k.ID, &k.NamaKelas, &k.Tingkat, &k.IdPerson); err != nil {
			return nil, err
		}
		kelas = append(kelas, k)
	}
	return kelas, nil
}

func (dao *KelasDao) GetAllKelas() ([]models.Kelas, error) {
	rows, err := dao.DB.Query(context.Background(), "SELECT id, nama_kelas, tingkat, id_person FROM kelas")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var kelas []models.Kelas
	for rows.Next() {
		var k models.Kelas
		if err := rows.Scan(&k.ID, &k.NamaKelas, &k.Tingkat, &k.IdPerson); err != nil {
			return nil, err
		}
		kelas = append(kelas, k)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return kelas, nil
}

func (dao *KelasDao) DeleteKelas(id int) error {
	_, err := dao.DB.Exec(context.Background(), "DELETE FROM kelas WHERE id=$1", id)
	return err
}
