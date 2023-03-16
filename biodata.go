package main

import (
	"fmt"
	"os"
)

type friend struct {
	name      string
	alamat    string
	pekerjaan string
	alasan    string
}

func main() {
	friend1 := friend{
		name:      "Rizar",
		alamat:    "bekasi barat",
		pekerjaan: "Software Engineer",
		alasan:    "Ingin memperdalam  bahasa Go.",
	}
	if len(os.Args) == 2 && os.Args[1] == "friend1" {
		fmt.Printf("Nama: %s\nAlamat: %s\nPekerjaan: %s\nAlasan: %s\n", friend1.name, friend1.alamat, friend1.pekerjaan, friend1.alasan)
	} else {
		fmt.Println("Data teman tidak ditemukan.")
	}

}
