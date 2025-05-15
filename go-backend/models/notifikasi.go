package models

import "time"

type Notifikasi struct {
    IDNotifikasi     uint      `gorm:"primaryKey" json:"id_notifikasi"`
    IDUser           uint      `json:"id_user"`
    IDRuangan        uint      `json:"id_ruangan"`
    IDReservasi      uint      `json:"id_reservasi"`
    StatusReservasi  string    `json:"status_reservasi"`
    WaktuDibuat      time.Time `json:"waktu_dibuat"`
}

func (Notifikasi) TableName() string {
    return "jadwal_reservasi"
}
