package models

import (
	"database/sql"
	"fmt"
	"time"

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
func (p *PasienModel)FindAll()([]entities.Pasien,error){
	query, err := p.conn.Query("SELECT * FROM pasien")
	if err != nil{
		return []entities.Pasien{}, err
	}
	defer query.Close()
	
	var dataPasien []entities.Pasien
	for query.Next(){
		var pasien entities.Pasien
		query.Scan(
			&pasien.Id, 
			&pasien.NamaLengkap, 
			&pasien.NIK, 
			&pasien.JenisKelamin, 
			&pasien.TempatLahir, 
			&pasien.TanggaLahir, 
			&pasien.Alamat, 
			&pasien.NoHp)

			if pasien.JenisKelamin == "1"{
				pasien.JenisKelamin = "Laki - Laki"
			}else {
				pasien.JenisKelamin = "Perempuan"
			}
			//2006-01-02 sama kayak yyyy-mm-dd
			tgl_lahir, _ := time.Parse("2006-01-02", pasien.TanggaLahir)
			//02-01-2006 sama kayak dd-mm-yyyy dan merupakan format indo	
			pasien.TanggaLahir = tgl_lahir.Format("02-01-2006")
			dataPasien = append(dataPasien, pasien)
	}
	return dataPasien, nil
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