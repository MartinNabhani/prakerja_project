package models

import (
	"gorm.io/gorm"
)

type Student struct {
	gorm.Model
	Nama         string  `json:"nama" form:"nama"`
	JenisKelamin string  `json:"jenis_kelamin" form:"jenis_kelamin"`
	TglLahir     string  `json:"tanggal_lahir" form:"tanggal_lahir" `
	Alamat       string  `json:"alamat" form:"alamat"`
	NoHP         string  `json:"no_hp" form:"no_hp"`
	Email        string  `json:"email" form:"email"`
	Nilai        float32 `json:"nilai" form:"nilai"`
	Keterangan   string  `json:"keterangan" form:"keterangan"`
}
