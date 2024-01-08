package main

import (
	"net/http"

	"github.com/IrfanBain/controllers/pasiencontroller"
)

func main() {

	http.HandleFunc("/", pasiencontroller.index)
	http.HandleFunc("/pasien", pasiencontroller.index)
	http.HandleFunc("/pasien/index", pasiencontroller.index)
	http.HandleFunc("/pasien/add", pasiencontroller.add)
	http.HandleFunc("/pasien/edit", pasiencontroller.edit)
	http.HandleFunc("/pasien/delete", pasiencontroller.delete)

	http.ListenAndServe(":8080", nil)
}