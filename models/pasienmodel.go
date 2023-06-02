package models

import (
	"database/sql"
	"fmt"

	"github.com/lulungsatrioprayuda/go-crud/config"
	"github.com/lulungsatrioprayuda/go-crud/entities"
)

type PasienModel struct {
	conn *sql.DB
}

func NewPasienModel() *PasienModel {
	conn, err := config.DBConnection()
	if err != nil {
		panic(err)
	}
	return &PasienModel{
		conn: conn,
	}
}

func (p *PasienModel)Create(pasien entities.Pasien) bool{
	result, err := p.conn.Exec("INSERT INTO pasien (nama_lengkap, nik, jenis_kelamin, tempat_lahir, tanggal_lahir, alamat, no_hp) VALUES(?,?,?,?,?,?,?)", 
	pasien.NamaLengkap, pasien.NIK, pasien.JenisKelamin, pasien.TempatLahir, pasien.TanggaLahir, pasien.Alamat, pasien.NoHp)

	if err != nil {
		fmt.Println(err)
		return false
	}

	lastInsertId, _ := result.LastInsertId()

	return lastInsertId > 0
}