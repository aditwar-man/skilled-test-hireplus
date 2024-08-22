package main

import (
	"fmt"
	"math"
	"time"
)

// Fungsi untuk menghitung apakah cuti diperbolehkan
func cekCuti(jumlahCutiBersama int, tanggalJoin string, tanggalRencanaCuti string, durasiCuti int) (bool, string) {
	// Format waktu
	layout := "2006-01-02"

	// Parsing tanggal join dan tanggal rencana cuti
	tJoin, err := time.Parse(layout, tanggalJoin)
	if err != nil {
		return false, "Format tanggal join tidak valid"
	}
	tRencanaCuti, err := time.Parse(layout, tanggalRencanaCuti)
	if err != nil {
		return false, "Format tanggal rencana cuti tidak valid"
	}

	// Hitung 180 hari setelah tanggal join
	tanggalMulaiBisaCuti := tJoin.AddDate(0, 0, 180)
	if tRencanaCuti.Before(tanggalMulaiBisaCuti) {
		return false, "Belum 180 hari sejak tanggal join karyawan"
	}

	// Hitung jumlah hari dari tanggal mulai bisa cuti sampai akhir tahun
	akhirTahun := time.Date(tRencanaCuti.Year(), 12, 31, 0, 0, 0, 0, time.UTC)
	jumlahHariTahunPertama := int(akhirTahun.Sub(tanggalMulaiBisaCuti).Hours()/24) + 1

	// Hitung jumlah cuti pribadi yang diperbolehkan
	cutiKantor := 14
	jumlahCutiPribadi := cutiKantor - jumlahCutiBersama
	jumlahCutiPribadiTahunPertama := int(math.Floor(float64(jumlahHariTahunPertama) * float64(jumlahCutiPribadi) / 365.0))

	// Cek apakah durasi cuti melebihi jumlah cuti pribadi yang diperbolehkan
	if durasiCuti > jumlahCutiPribadiTahunPertama {
		return false, fmt.Sprintf("Hanya boleh mengambil %d hari cuti", jumlahCutiPribadiTahunPertama)
	}

	// Cek apakah durasi cuti lebih dari 3 hari berturut-turut
	if durasiCuti > 3 {
		return false, "Cuti pribadi maksimal 3 hari berturut-turut"
	}

	return true, "Cuti diperbolehkan"
}

func main() {
	// Contoh penggunaan
	jumlahCutiBersama := 7
	tanggalJoin := "2021-01-05"
	tanggalRencanaCuti := "2021-12-18"
	durasiCuti := 3

	bolehCuti, alasan := cekCuti(jumlahCutiBersama, tanggalJoin, tanggalRencanaCuti, durasiCuti)
	fmt.Println(bolehCuti) // False
	fmt.Println(alasan)    // Karena hanya boleh mengambil 1 hari cuti
}
