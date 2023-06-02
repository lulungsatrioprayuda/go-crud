package pasiencontroller

import (
	"html/template"
	"net/http"

	"github.com/lulungsatrioprayuda/go-crud/entities"
	"github.com/lulungsatrioprayuda/go-crud/models"
)

var pasienModel = models.NewPasienModel()

func Index(response http.ResponseWriter, request *http.Request) {
	temp, err := template.ParseFiles("views/pasien/index.html")
	if err != nil {
		panic(err)
	}
	temp.Execute(response,nil)
	
}
func Add(response http.ResponseWriter, request *http.Request) {

	if request.Method == http.MethodGet{
		temp, err := template.ParseFiles("views/pasien/add.html")
		if err != nil {
			panic(err)
		}
		temp.Execute(response,nil)
	}else if request.Method == http.MethodPost{
		request.ParseForm()

		var pasien entities.Pasien
		pasien.NamaLengkap = request.Form.Get("nama_lengkap")
		pasien.NIK = request.Form.Get("nik")
		pasien.JenisKelamin = request.Form.Get("jenis_kelamin")
		pasien.TempatLahir = request.Form.Get("tempat_lahir")
		pasien.TanggaLahir = request.Form.Get("tanggal_lahir")
		pasien.Alamat = request.Form.Get("alamat")
		pasien.NoHp = request.Form.Get("no_hp")
		
		pasienModel.Create(pasien)
		data := map[string]interface{}{
			"pesan": "Data pasien baru disimpan",
		}
		
		temp, _ := template.ParseFiles("views/pasien/add.html")
		temp.Execute(response, data)
	}


}
func Edit(response http.ResponseWriter, request *http.Request) {

}
func Delete(response http.ResponseWriter, request *http.Request) {

}