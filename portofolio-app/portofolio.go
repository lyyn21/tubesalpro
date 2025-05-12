package main

//Struktur data untuk menyimpan informasi proyek
type Portfolio struct {
	ID          int
	Name        string
	Description string
	Skills      []string
	Year        int
}

//Variabel global untuk menyimpan daftar proyek
var Portfolios []Portfolio

//Menambah proyek ke daftar
func Addproject(p Portofolio) {
	Portfolios = append(Portfolios, p)
}

//Mengembalikan semua proyek
func GetAllProjects() []Portofolio {
	return Portofolios
}
