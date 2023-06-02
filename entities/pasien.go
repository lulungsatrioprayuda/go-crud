package entities

//ini buat nampung data pasien dari database
type Pasien struct{
	Id int64
	NamaLengkap string
	NIK string
	JenisKelamin string
	TempatLahir string
	TanggaLahir string
	Alamat string
	NoHp string
}