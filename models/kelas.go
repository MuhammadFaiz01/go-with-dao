package models

import "github.com/google/uuid"

type Kelas struct {
	ID        uuid.UUID `json:"id"`
	NamaKelas string    `json:"nama_kelas"`
	Tingkat   int       `json:"tingkat"`
	IdPerson  int       `json:"id_person"`
}
