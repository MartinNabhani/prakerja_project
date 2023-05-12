package controllers

import (
	"net/http"
	"prakerja/configs"
	"prakerja/models"

	"github.com/labstack/echo/v4"
)

func GetAllSiswa(c echo.Context) error {
	var students []models.Student
	if err := configs.DB.Find(&students).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{
			Status: false, Message: "Gagal get data siswa dari database", Data: nil,
		})
	}
	return c.JSON(http.StatusOK, models.BaseResponse{
		Status: true, Message: "Sukses Mendapatkan Semua Siswa", Data: students,
	})
}

func GetSiswaByID(c echo.Context) error {

	// ambil data siswa berdasarkan id
	var student models.Student
	if err := configs.DB.First(&student, c.Param("id")).Error; err != nil {
		return c.JSON(http.StatusBadRequest, models.BaseResponse{
			Status: false, Message: "Data Tidak DItemukan", Data: nil,
		})
	}

	// return data dalam format JSON
	return c.JSON(http.StatusOK, models.BaseResponse{
		Status: true, Message: "Data berhasil Ditemukan", Data: student,
	})
}

// create new student
func CreateSiswaController(c echo.Context) error {
	student := models.Student{}
	c.Bind(&student)
	if err := configs.DB.Save(&student).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{
			Status: false, Message: "Gagal Create User", Data: nil,
		})
	}
	return c.JSON(http.StatusOK, models.BaseResponse{
		Status: true, Message: "Berhasil Create User", Data: student,
	})
}

func DeleteSiswaController(c echo.Context) error {
	id := c.Param("id")
	var student models.Student
	result := configs.DB.First(&student, id)
	if result.Error != nil {
		return c.JSON(http.StatusNotFound, models.BaseResponse{
			Status: false, Message: "Data Tidak Berhasil ditemukan", Data: nil,
		})
	}
	configs.DB.Delete(&student)
	return c.JSON(http.StatusOK, models.BaseResponse{
		Status: true, Message: "Berhasil Hapus Data", Data: student,
	})
}
func UpdateSiswaController(c echo.Context) error {
	studentID := c.Param("id")

	// Mencari student dengan ID yang diberikan
	var student models.Student
	if err := configs.DB.First(&student, studentID).Error; err != nil {
		return c.JSON(http.StatusNotFound, models.BaseResponse{
			Status: false, Message: "Data Tidak Berhasil ditemukan", Data: nil,
		})
	}

	// Mendapatkan data inputan dari client
	var input struct {
		models.Student
	}
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	// Memperbarui informasi student
	student.Nama = input.Nama
	student.JenisKelamin = input.JenisKelamin
	student.TglLahir = input.TglLahir
	student.Alamat = input.Alamat
	student.NoHP = input.NoHP
	student.Email = input.Email
	student.Nilai = input.Nilai

	if err := configs.DB.Save(&student).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{
			Status: false, Message: "Gagal Memperbarui data", Data: nil,
		})
	}

	return c.JSON(http.StatusOK, models.BaseResponse{
		Status: true, Message: "Berhasil Update Data", Data: student,
	})
}
