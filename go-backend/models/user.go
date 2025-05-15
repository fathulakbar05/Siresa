package models

 type User struct {
  IDUser   uint   `gorm:"primaryKey" json:"id_user"`
  Nama     string `json:"nama"`
  NIM_NIP  string `json:"nim_nip" gorm:"column:nim_nip"`
  Password string `json:"password"` // Pastikan ini persis sama dengan nama kolom di DB
  Role     string `json:"role"`
 }

 func (User) TableName() string {
  return "users" // Pastikan ini benar
 }