package main

import "fmt"

const NMAX = 100

type Barang struct {
	Kode  string
	Nama  string
	Harga int
	Stok  int
}

type Transaksi struct {
	KodeBarang string
	NamaBarang string
	Tanggal    string
	Jenis      string
	Jumlah     int
}

var masterBarang [NMAX]Barang
var riwayat [NMAX]Transaksi

var nBarang int
var nTransaksi int

func main() {

	var menu int

	for {

		fmt.Println()
		fmt.Println("===== GUDANGIN =====")
		fmt.Println("1. Tambah Barang")
		fmt.Println("2. Tampilkan Barang")
		fmt.Println("3. Ubah Barang")
		fmt.Println("4. Hapus Barang")
		fmt.Println("5. Transaksi Masuk")
		fmt.Println("6. Transaksi Keluar")
		fmt.Println("7. Riwayat Transaksi Barang")
		fmt.Println("8. Sequential Search")
		fmt.Println("9. Binary Search")
		fmt.Println("10. Selection Sort Ascending")
		fmt.Println("11. Insertion Sort Descending")
		fmt.Println("12. Statistik Gudang")
		fmt.Println("0. Keluar")

		fmt.Print("Pilih : ")
		fmt.Scan(&menu)

		switch menu {

		case 1:
			tambahBarang()

		case 2:
			tampilBarang()

		case 3:
			ubahBarang()

		case 4:
			hapusBarang()

		case 5:
			transaksiMasuk()

		case 6:
			transaksiKeluar()

		case 7:
			riwayatPerBarang()

		case 8:
			sequentialSearch()

		case 9:
			binarySearch()

		case 10:
			selectionSort()

		case 11:
			insertionSort()

		case 12:
			statistikGudang()

		case 0:
			return

		default:
			fmt.Println("Menu tidak tersedia")
		}
	}
}

func tambahBarang() {

	fmt.Print("Kode Barang : ")
	fmt.Scan(&masterBarang[nBarang].Kode)

	fmt.Print("Nama Barang : ")
	fmt.Scan(&masterBarang[nBarang].Nama)

	fmt.Print("Harga Barang : ")
	fmt.Scan(&masterBarang[nBarang].Harga)

	fmt.Print("Stok Awal : ")
	fmt.Scan(&masterBarang[nBarang].Stok)

	nBarang++

	fmt.Println("Barang berhasil ditambah")
}

func tampilBarang() {

	if nBarang == 0 {
		fmt.Println("Data kosong")
		return
	}

	for i := 0; i < nBarang; i++ {

		fmt.Println("-------------------")
		fmt.Println("Kode :", masterBarang[i].Kode)
		fmt.Println("Nama :", masterBarang[i].Nama)
		fmt.Println("Harga :", masterBarang[i].Harga)
		fmt.Println("Stok :", masterBarang[i].Stok)
	}
}

func ubahBarang() {

	var kode string

	fmt.Print("Kode Barang : ")
	fmt.Scan(&kode)

	for i := 0; i < nBarang; i++ {

		if masterBarang[i].Kode == kode {

			fmt.Print("Nama Baru : ")
			fmt.Scan(&masterBarang[i].Nama)

			fmt.Print("Harga Baru : ")
			fmt.Scan(&masterBarang[i].Harga)

			fmt.Println("Data berhasil diubah")
			return
		}
	}

	fmt.Println("Barang tidak ditemukan")
}

func hapusBarang() {

	var kode string

	fmt.Print("Kode Barang : ")
	fmt.Scan(&kode)

	for i := 0; i < nBarang; i++ {

		if masterBarang[i].Kode == kode {

			for j := i; j < nBarang-1; j++ {
				masterBarang[j] = masterBarang[j+1]
			}

			nBarang--

			fmt.Println("Data berhasil dihapus")
			return
		}
	}

	fmt.Println("Barang tidak ditemukan")
}

func transaksiMasuk() {

	var kode string
	var tanggal string
	var jumlah int

	fmt.Print("Kode Barang : ")
	fmt.Scan(&kode)

	for i := 0; i < nBarang; i++ {

		if masterBarang[i].Kode == kode {

			fmt.Print("Tanggal : ")
			fmt.Scan(&tanggal)

			fmt.Print("Jumlah Masuk : ")
			fmt.Scan(&jumlah)

			masterBarang[i].Stok += jumlah

			riwayat[nTransaksi].KodeBarang = kode
			riwayat[nTransaksi].NamaBarang = masterBarang[i].Nama
			riwayat[nTransaksi].Tanggal = tanggal
			riwayat[nTransaksi].Jenis = "MASUK"
			riwayat[nTransaksi].Jumlah = jumlah

			nTransaksi++

			fmt.Println("Transaksi berhasil")
			return
		}
	}

	fmt.Println("Barang tidak ditemukan")
}

func transaksiKeluar() {

	var kode string
	var tanggal string
	var jumlah int

	fmt.Print("Kode Barang : ")
	fmt.Scan(&kode)

	for i := 0; i < nBarang; i++ {

		if masterBarang[i].Kode == kode {

			fmt.Print("Tanggal : ")
			fmt.Scan(&tanggal)

			fmt.Print("Jumlah Keluar : ")
			fmt.Scan(&jumlah)

			if jumlah > masterBarang[i].Stok {

				fmt.Println("Stok tidak cukup")
				return
			}

			masterBarang[i].Stok -= jumlah

			riwayat[nTransaksi].KodeBarang = kode
			riwayat[nTransaksi].NamaBarang = masterBarang[i].Nama
			riwayat[nTransaksi].Tanggal = tanggal
			riwayat[nTransaksi].Jenis = "KELUAR"
			riwayat[nTransaksi].Jumlah = jumlah

			nTransaksi++

			fmt.Println("Transaksi berhasil")
			return
		}
	}

	fmt.Println("Barang tidak ditemukan")
}

func riwayatPerBarang() {

	var kode string

	fmt.Print("Kode Barang : ")
	fmt.Scan(&kode)

	fmt.Println()
	fmt.Println("=== RIWAYAT BARANG ===")

	for i := 0; i < nTransaksi; i++ {

		if riwayat[i].KodeBarang == kode {

			fmt.Println("----------------")
			fmt.Println("Tanggal :", riwayat[i].Tanggal)
			fmt.Println("Jenis :", riwayat[i].Jenis)
			fmt.Println("Jumlah :", riwayat[i].Jumlah)
		}
	}
}

func sequentialSearch() {

	var nama string

	fmt.Print("Nama Barang : ")
	fmt.Scan(&nama)

	for i := 0; i < nBarang; i++ {

		if masterBarang[i].Nama == nama {

			fmt.Println("Data ditemukan")
			fmt.Println(masterBarang[i])

			return
		}
	}

	fmt.Println("Tidak ditemukan")
}

func urutKode() {

	for i := 0; i < nBarang-1; i++ {

		for j := i + 1; j < nBarang; j++ {

			if masterBarang[i].Kode > masterBarang[j].Kode {

				temp := masterBarang[i]
				masterBarang[i] = masterBarang[j]
				masterBarang[j] = temp
			}
		}
	}
}

func binarySearch() {

	urutKode()

	var kode string

	fmt.Print("Kode Barang : ")
	fmt.Scan(&kode)

	left := 0
	right := nBarang - 1

	for left <= right {

		mid := (left + right) / 2

		if masterBarang[mid].Kode == kode {

			fmt.Println("Data ditemukan")
			fmt.Println(masterBarang[mid])

			return

		} else if kode < masterBarang[mid].Kode {

			right = mid - 1

		} else {

			left = mid + 1
		}
	}

	fmt.Println("Tidak ditemukan")
}

func selectionSort() {

	for i := 0; i < nBarang-1; i++ {

		min := i

		for j := i + 1; j < nBarang; j++ {

			if masterBarang[j].Stok < masterBarang[min].Stok {
				min = j
			}
		}

		temp := masterBarang[i]
		masterBarang[i] = masterBarang[min]
		masterBarang[min] = temp
	}

	fmt.Println("Ascending berdasarkan stok")
	tampilBarang()
}

func insertionSort() {

	for i := 1; i < nBarang; i++ {

		key := masterBarang[i]
		j := i - 1

		for j >= 0 && masterBarang[j].Stok < key.Stok {

			masterBarang[j+1] = masterBarang[j]
			j--
		}

		masterBarang[j+1] = key
	}

	fmt.Println("Descending berdasarkan stok")
	tampilBarang()
}

func statistikGudang() {

	totalAset := 0

	for i := 0; i < nBarang; i++ {

		totalAset += masterBarang[i].Harga * masterBarang[i].Stok
	}

	fmt.Println()
	fmt.Println("Total Nilai Aset Gudang :", totalAset)

	if nBarang == 0 {
		return
	}

	min := 0

	for i := 1; i < nBarang; i++ {

		if masterBarang[i].Stok < masterBarang[min].Stok {
			min = i
		}
	}

	fmt.Println()
	fmt.Println("Barang dengan stok paling sedikit")
	fmt.Println("Kode :", masterBarang[min].Kode)
	fmt.Println("Nama :", masterBarang[min].Nama)
	fmt.Println("Stok :", masterBarang[min].Stok)
}
