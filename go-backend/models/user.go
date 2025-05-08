package models

type User struct {
    IDUser   uint   `gorm:"primaryKey" json:"id_user"`
    Nama     string `json:"nama"`
    NIM_NIP  string `json:"nim_nip"`
    Password string `json:"password"`
    Role     string `json:"role"` // e.g. "admin", "user"
}

func (User) TableName() string {
    return "jadwal_reservasi"
}
