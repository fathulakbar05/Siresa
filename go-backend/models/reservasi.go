 package models

 import "time"

 type Reservasi struct {
  IDReservasi  uint      `gorm:"primaryKey" json:"id_reservasi"`
  IDUser       uint      `json:"id_user"`
  IDRuangan    uint      `json:"id_ruangan"`
  Tanggal      string    `json:"tanggal"`      // Format: "YYYY-MM-DD"
  JamMulai     string    `json:"jam_mulai"`    // Format: "HH:MM:SS"
  JamSelesai   string    `json:"jam_selesai"`  // Format: "HH:MM:SS"
  Status       string    `json:"status"`       // "menunggu", "disetujui", "ditolak"
  WaktuPengajuan time.Time `json:"waktu_pengajuan"`
 }

 func (Reservasi) TableName() string {
  return "jadwal_reservasi" // Nama tabel yang benar
 }