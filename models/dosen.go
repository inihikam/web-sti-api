package models

type Jabatan struct {
	ID   uint   `json:"id"`
	Nama string `json:"nama"`
}

type Dosen struct {
	ID      uint    `json:"id"`
	Nama    string  `json:"nama"`
	NPP     string  `json:"npp"`
	Jabatan Jabatan `json:"jabatan"`
}

type LecturersByPosition struct {
	Position  string  `json:"position"`
	Lecturers []Dosen `json:"lecturers"`
}