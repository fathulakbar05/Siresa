package models

import "time"

type Reservasi struct {
    IDReservasi  uint      `gorm:"primaryKey" json:"id_reservasi"`
    IDUser       uint      `json:"id_user"`
    IDRuangan    uint      `json:"id_ruangan"`
    Tanggal      string    `json:"tanggal"`
    JamMulai     string    `json:"jam_mulai"`
    JamSelesai   string    `json:"jam_selesai"`
    Status       string    `json:"status"`
    WaktuPengajuan time.Time `json:"waktu_pengajuan"`
}

func (Reservasi) TableName() string {
    return "jadwal_reservasi"
}
