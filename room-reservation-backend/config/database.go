package config

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"fmt"
	"log"
)

var DB *gorm.DB

func InitDB() {
	// Gunakan sqlite tanpa konfigurasi user/pass/host, cukup file database
	var err error
	DB, err = gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Gagal koneksi ke database:", err)
	}

	// Mengambil koneksi raw dari DB untuk memastikan bisa terhubung
	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatal("Gagal mendapatkan koneksi database:", err)
	}

	// Lakukan ping untuk memverifikasi koneksi ke database
	err = sqlDB.Ping()
	if err != nil {
		log.Fatal("Database tidak merespons:", err)
	}

	// Jika semuanya berhasil
	fmt.Println("Koneksi ke database berhasil")
}
