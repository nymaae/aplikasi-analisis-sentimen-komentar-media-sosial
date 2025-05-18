package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

const NMAX int = 1000

// Tipe Sentimen
const (
	sentimenPositif = "Positif"
	sentimenNetral  = "Netral"
	sentimenNegatif = "Negatif"
)

// Tipe bentukan Komentar
type komentar struct {
	teks      string
	sentimen  sentiment
}

type sentiment string

// Array statis utama
var listKomentar [NMAX]komentar
var jumlahKomentar int = 0


// Kata kunci analisis sentimen
var kataPositif = [13]string{"bagus", "baik", "cantik", "ganteng", "keren", "mantap", "hebat", "memuaskan", "indah", "menarik", "puas", "senang", "suka"}
var kataNegatif = [13]string{"buruk", "jahat", "jelek", "benci", "gagal", "tidak suka", "mengecewakan", "lebay", "aneh", "tidak suka", "marah", "sedih", "rusak"}

// Initial State: listKomentar belum ada, jumlahKomentar = 0 
// Final state: komentarList telah berisi komentar-komentar yang sudah dianalisis, diedit, dihapus, dicari, atau diurutkan sesuai interaksi pengguna.

func analisisSentimen(teks string) sentiment {
	var i int
	for i = 0; i < len(kataPositif); i++ {
		if strings.Contains(teks, kataPositif[i]) {
			return sentiment(sentimenPositif)
		}
	}
	for i = 0; i < len(kataNegatif); i++ {
		if strings.Contains(teks, kataNegatif[i]) {
			return sentiment(sentimenNegatif)
		}
	}
	return sentiment(sentimenNetral)
}

func bersihLayar() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func header() {
	fmt.Println("==============================================================================")
	fmt.Println("                        SELAMAT DATANG DI SENTILYTICS                         ")
	fmt.Println("==============================================================================")
	fmt.Println("Aplikasi ini akan membantu menganalisis sentimen dari komentar di media sosial")
	fmt.Println("==============================================================================")
}

func login(P, S *string) {
    fmt.Println("LOGIN")
	
	for {
		fmt.Print("Masukkan username: ")
		fmt.Scan(S)
		fmt.Print("Masukkan password: ")
		fmt.Scan(P)
		
		if *S == "kamari" && *P == "kamari" {
			break
		}
		fmt.Println("Login gagal!")
	}
	fmt.Println("Login berhasil!")
}

// Menu Utama
func menu() {
	var pilih int
	for pilih != 9 {
	        fmt.Println("------------------------------")
	        fmt.Println("           M E N U            ")
	        fmt.Println("------------------------------")
	        fmt.Println("1. Tambah Komentar")
	        fmt.Println("2. Tampilkan Semua Komentar")
	        fmt.Println("3. Edit Komentar")
	        fmt.Println("4. Hapus Komentar")
	        fmt.Println("5. Cari Komentar")
	        fmt.Println("6. Urutkan Panjang Teks")
	        fmt.Println("7. Urutkan Sentimen")
	        fmt.Println("8. Statistik Sentimen")
	        fmt.Println("9. Exit")
	        fmt.Println("------------------------------")
	        fmt.Print("Tentukan Pilihan Anda (1-9): ")
			fmt.Scan(&pilih)
		    switch pilih {
		    case 1:
				    bersihLayar()
			        tambahKomentar(&listKomentar, &jumlahKomentar)
			case 2:
				    bersihLayar()
			        tampilkanSemuaKomentar()
			case 3:
				    bersihLayar()
			        editKomentar()  
			case 4:
				    bersihLayar()
			        hapusKomentar()
			case 5:
				    bersihLayar()
			        cariKomentar()
			case 6:
				    bersihLayar()
			        urutkanPanjangKomentar()
			case 7:
				    bersihLayar()
			        urutkanSentimenKomentar()
			case 8:
				    bersihLayar()
			        statistikSentimen()
			case 9:
				    bersihLayar()
			        fmt.Println("Terima kasih telah menggunakan Sentilytics!")
			        os.Exit(0)
		}
    }
}

// Subprogram: Menambahkan Komentar
func tambahKomentar(list *[NMAX]komentar, jumlah *int) {
	var teks string
	var pilih int 
	pilih = 1
	for pilih != 0 && *jumlah < NMAX {
		fmt.Println("---------------------------------------")
		fmt.Println("         MENAMBAHKAN KOMENTAR         ")
		fmt.Println("---------------------------------------")

		fmt.Print("Masukkan komentar: ")
		fmt.Scan(&teks) 
		list[*jumlah].teks = teks
		list[*jumlah].sentimen = analisisSentimen(teks)
		*jumlah++

		fmt.Println("Komentar berhasil ditambahkan.")
		fmt.Println("<<1. Tambah komentar lagi>>", "\t", "<<0. Kembali ke Menu>>")
		fmt.Scan(&pilih)
 	}
	    bersihLayar()
		menu()
}

// Subprogram: Menu untuk Menampilkan Semua Komentar
func TampilkanSemuaKomentar() {
	var pilih int
    fmt.Println("-----------------------------------------------")
    fmt.Println("      MENU PILIHAN MENAMPILKAN SEMUA KOMENTAR  ")
    fmt.Println("-----------------------------------------------")
    fmt.Println("1. Tampilkan Komentar (Ascending)")
    fmt.Println("2. Tampilkan Komentar (Descending)")
    fmt.Println("3. Kembali ke Menu")
    fmt.Println("-----------------------------------------------")
    fmt.Print("Masukkan pilihan Anda: ")
    fmt.Scan(&pilih)

    switch pilih {
	case 1:
		bersihLayar()
		sortKomentarAsc()
		tampilkanSemuaKomentar()
	case 2:
		bersihLayar()
		sortKomentarDesc()
		TampilkanSemuaKomentar()
	case 3:
		for pilih != 3 {
				fmt.Println("Pilihan tidak valid.")
				fmt.Scan(&pilih)
			}
			bersihLayar()
			if pilih == 0 {
				return
			}
		bersihLayar()
		if pilih == 3 {
			return
		}
	}
	menu()
}

func sortKomentarAsc() { 
//Mengurutkan teks yang diinput dari yang paling sedikit
		var i, minIdx, j int 
		for i = 0; i < jumlahKomentar-1; i++ {
		minIdx = i
		for j = i + 1; j < jumlahKomentar; j++ {
			if len(listKomentar[j].teks) < len(listKomentar[minIdx].teks) {
				minIdx = j
			}
		}
		listKomentar[i], listKomentar[minIdx] = listKomentar[minIdx], listKomentar[i]
	}
}

func sortKomentarDesc() {
//Mengurutkan teks yang diinput dari yang paling banyak
	var i, maxIdx, j int 
	for i = 0; i < jumlahKomentar-1; i++ {
		maxIdx = i
		for j = i + 1; j < jumlahKomentar; j++ {
			if len(listKomentar[j].teks) > len(listKomentar[maxIdx].teks) {
				maxIdx = j
			}
		}
		listKomentar[i], listKomentar[maxIdx] = listKomentar[maxIdx], listKomentar[i]
	}
}


// Subprogram: Menampilkan Semua Komentar
func tampilkanSemuaKomentar(){
	var i int 
	fmt.Println("Tampilkan Semua Komentar: ")
	if jumlahKomentar == 0 {
		fmt.Println("Belum ada komentar.")
		return
	}
	for i= 0; i < jumlahKomentar; i++ {
		fmt.Printf("[%d] %s - Sentimen: %s\n", i+1, listKomentar[i].teks, listKomentar[i].sentimen)
	}
}

// Subprogram: Mengedit Komentar
func editKomentar(){
	var i int 
	var lama, baru string //komentar lama dan baru  
	var found bool //penanda komentar lama ditemukan 
	
	fmt.Print("Masukkan komentar yang ingin diubah: ")
	fmt.Scan(&lama)
	found = false
	i = 0
	
	for !found && i < jumlahKomentar{ 
	//selama komentar lamanya belum ketemu dan i nya kurang dari jumlah komentar    
			if listKomentar[i].teks == lama {
			fmt.Print("Masukkan komentar baru: ")
			fmt.Scan(&baru)
			listKomentar[i].teks = baru
			listKomentar[i].sentimen = analisisSentimen(baru)
			fmt.Println("Komentar berhasil diubah.")
			found = true
		}else{
			i++
		}
		
	}
	
	if !found {
		fmt.Println("Komentar tidak ditemukan.")
	}
}

// Subprogram: Menghapus Komentar
func hapusKomentar(){
	var i, j int //i, j buat looping
	var teks string //cari komentar yang mau dihapus
	var found bool //penanda komentar lama ditemukan
	fmt.Print("Masukkan komentar yang ingin dihapus: ")
	fmt.Scan(&teks)
	found = false
	i = 0
	for !found && i < jumlahKomentar {
		if listKomentar[i].teks == teks {
			for j = i; j < jumlahKomentar-1; j++ {
				listKomentar[j] = listKomentar[j+1]
			}
			jumlahKomentar--
			fmt.Println("Komentar berhasil dihapus.")
			found = true
		} else {
			i++
		}
	}
	if !found {
		fmt.Println("Komentar tidak ditemukan.")
	}
}


// Subprogram: Mencari Komentar
func cariKomentar(){
	var keyword string
	var metode int
	fmt.Println("---------------------------------------")
	fmt.Println("             CARI KOMENTAR             ")
	fmt.Println("---------------------------------------")
	fmt.Print("Masukkan kata kunci: ")
	fmt.Scan(&keyword)

	fmt.Println("Pilih metode pencarian:")
	fmt.Println("1. Sequential Search")
	fmt.Println("2. Binary Search (urutan alfabet)")
	fmt.Print("Masukkan pilihan: ")
	fmt.Scan(&metode)

	if metode == 1 {
		sequentialSearch(keyword)
	} else if metode == 2 {
		binarySearch(keyword)
	} else {
		fmt.Println("Pilihan tidak valid.")
	}
}


func sequentialSearch(keyword string) {
	var ditemukan bool = false
	var i int
	for i = 0; i < jumlahKomentar; i++ {
		if strings.Contains(strings.ToLower(listKomentar[i].teks), strings.ToLower(keyword)) {
			fmt.Printf("[%d] %s - Sentimen: %s\n", i+1, listKomentar[i].teks, listKomentar[i].sentimen)
			ditemukan = true
		}
	}
	if !ditemukan {
		fmt.Println("Komentar tidak ditemukan.")
	}
}


func binarySearch(keyword string) {
	var low, high, mid int
	sortKomentarAlphabet()
	low = 0
	high = jumlahKomentar - 1
	var found bool = false

	for low <= high {
		mid = (low + high) / 2
		if strings.Contains(strings.ToLower(listKomentar[mid].teks), strings.ToLower(keyword)) {
			fmt.Printf("[%d] %s - Sentimen: %s\n", mid+1, listKomentar[mid].teks, listKomentar[mid].sentimen)
			found = true
			break
		} else if strings.ToLower(keyword) < strings.ToLower(listKomentar[mid].teks) {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	if !found {
		fmt.Println("Komentar tidak ditemukan.")
	}
}


func sortKomentarAlphabet() {
	var i, j int
	for i = 0; i < jumlahKomentar-1; i++ {
		for j = i + 1; j < jumlahKomentar; j++ {
			if listKomentar[i].teks > listKomentar[j].teks {
				listKomentar[i], listKomentar[j] = listKomentar[j], listKomentar[i]
			}
		}
	}
}


// Subprogram: Mengurutkan Panjang Komentar
func urutkanPanjangKomentar(){
	var mode string 
	fmt.Print("Urutkan berdasarkan panjang komentar (asc/desc): ")
	fmt.Scan(&mode)
	var i, idx, j int
	for i = 0; i < jumlahKomentar-1; i++ {
		idx = i
		for j = i + 1; j < jumlahKomentar; j++ {
			if (mode == "asc" && len(listKomentar[j].teks) < len(listKomentar[idx].teks)) ||
				(mode == "desc" && len(listKomentar[j].teks) > len(listKomentar[idx].teks)) {
				idx = j
			}
		}
		listKomentar[i], listKomentar[idx] = listKomentar[idx], listKomentar[i]
	}
	fmt.Println("Komentar berhasil diurutkan.")
	tampilkanSemuaKomentar()
}


// Subprogram: Mengurutkan Sentimen Komentar
func urutkanSentimenKomentar(){
	var mode string
	var i, j int
	fmt.Print("Urutkan berdasarkan sentimen (asc/desc): ")
	fmt.Scan(&mode)
	for i = 1; i < jumlahKomentar; i++ {
		key := listKomentar[i]
		j = i - 1
		for j >= 0 && ((mode == "asc" && listKomentar[j].sentimen > key.sentimen) ||
			(mode == "desc" && listKomentar[j].sentimen < key.sentimen)) {
			listKomentar[j+1] = listKomentar[j]
			j--
		}
		listKomentar[j+1] = key
	}
	fmt.Println("Komentar berhasil diurutkan.")
	tampilkanSemuaKomentar()
}

//gatau pake ini engga
func prioritasSentimen(s sentiment) int {
	switch s {
	case sentiment(sentimenPositif):
		return 1
	case sentiment(sentimenNetral):
		return 2
	case sentiment(sentimenNegatif):
		return 3
	default:
		return 4
	}
}


// Subprogram: Statistik Komentar
func statistikSentimen(){
	var i int 
	var positif, netral, negatif int
	for i = 0; i < jumlahKomentar; i++ {
		switch listKomentar[i].sentimen {
		case sentiment(sentimenPositif):
			positif++
		case sentiment(sentimenNetral):
			netral++
		case sentiment(sentimenNegatif):
			negatif++
		}
	}

	fmt.Println("Statistik Sentimen Komentar:")
	fmt.Printf("Positif:", positif)
	fmt.Printf("Netral :", netral)
	fmt.Printf("Negatif:", negatif)
}


func main() {
	header()
	var username, password string
	login(&username, &password)
	menu()
}
