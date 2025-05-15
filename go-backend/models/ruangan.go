package models

type Ruangan struct {
  IDRuangan    uint   `gorm:"primaryKey" json:"id_ruangan"`
  NamaRuangan  string `json:"nama_ruangan"`
  LokasiGedung string `json:"lokasi_gedung"`
  Lantai       int    `json:"lantai"`
  Kapasitas    int    `json:"kapasitas"`
  Fasilitas    string `json:"fasilitas"`
  Ukuran       string `json:"ukuran"`
  Tipe         string `json:"tipe"`
  Layout       string `json:"layout"`
 }

 func (Ruangan) TableName() string {
  return "ruangan" // Nama tabel yang benar
 }