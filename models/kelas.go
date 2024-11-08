package models

type Kelas struct {
	ID        int    `json:"id"`
	NamaKelas string `json:"nama_kelas"`
	Tingkat   int    `json:"tingkat"`
	IdPerson  int    `json:"id_person"`
}
