package models

type BkAnnounce struct {
	ID             uint   `json:"id"`
	Title          string `json:"nama_kegiatan"`
	Content        string `json:"deskripsi"`
	TanggalMulai   string `json:"tanggal_mulai"`
	TanggalSelesai string `json:"tanggal_selesai"`
	Time           string `json:"waktu"`
	Location       string `json:"lokasi"`
	User           string `json:"mentor"`
}