package utils

import (
 "crypto/sha256"
 "encoding/hex"
)

// HashPasswordSHA256 melakukan hash pada password menggunakan SHA256
func HashPasswordSHA256(password string) string { // Ubah nama fungsi menjadi PascalCase
 hasher := sha256.New()
 hasher.Write([]byte(password))
 return hex.EncodeToString(hasher.Sum(nil))
}

// VerifyPasswordSHA256 memverifikasi apakah plainTextPassword cocok dengan hashedPassword
func VerifyPasswordSHA256(hashedPassword, plainTextPassword string) bool { // Ubah nama fungsi menjadi PascalCase
 return hashedPassword == HashPasswordSHA256(plainTextPassword)
}