package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

//// Maksimum jumlah komentar yang dapat disimpan
const NMAX int = 1000

// Tipe Sentimen
const (
	sentimenPositif = "Positif"
	sentimenNetral  = "Netral"
	sentimenNegatif = "Negatif"
)

// Tipe bentukan Komentar untuk menyimpan komentar dan hasil sentimennya
type komentar struct {
	teks      string
	sentimen  sentiment
}

// Alias Tipe bentukan Sentiment 
type sentiment string

// Array statis utama untuk menyimpan semua komentar 
var listKomentar [NMAX]komentar
var jumlahKomentar int = 0

// Kata kunci analisis sentimen
var positif, netral, negatif int 
var kataPositif = [13]string{"bagus", "baik", "cantik", "ganteng", "keren", "mantap", "hebat", "memuaskan", "indah", "menarik", "puas", "senang", "suka"}
var kataNegatif = [13]string{"buruk", "jahat", "jelek", "benci", "gagal", "tidak suka", "mengecewakan", "lebay", "aneh", "tidak suka", "marah", "sedih", "rusak"}


// Subprogram: Analisis Sentimen
/*
	I.S. : Belum ada sentimen yang ditentukan untuk sebuah komentar
	F.S. : Sentimen dari komentar dikembalikan dalam bentuk kategori: positif, negatif, atau netral
*/
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


// Subprogram: Bersih Layar
/*
	I.S. : Terminal dalam keadaan sebelumnya (berisi tampilan/output dari proses sebelumnya)
	F.S. : Layar terminal dibersihkan dan siap untuk menampilkan tampilan baru
*/
func bersihLayar() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}


// Subprogram: Header
/*
	I.S. : Belum ada tampilan header pada terminal
	F.S. : Tampilan header program Sentilytics muncul di terminal
*/
func header() {
	fmt.Println("==============================================================================")
	fmt.Println("                        SELAMAT DATANG DI SENTILYTICS                         ")
	fmt.Println("==============================================================================")
	fmt.Println("Aplikasi ini akan membantu menganalisis sentimen dari komentar di media sosial")
	fmt.Println("==============================================================================")
}


// Subprogram   : Login Pengguna
/*
	I.S. : Variabel pointer S dan P belum berisi username dan password pengguna.
	F.S. : Variabel S dan P telah berisi nilai "kamari", dan proses login berhasil dilakukan.
*/
func login(S, P *string) {
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



// Subprogram    : Menampilkan Menu Utama
/*
	I.S. : Program belum menunjukkan menu pilihan ke pengguna.
	F.S. : Menu ditampilkan dan aksi diambil sesuai pilihan pengguna hingga pengguna memilih keluar (exit).
*/
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


// Subprogram : Menambahkan Komentar
/*
	I.S. : listKomentar mungkin kosong atau berisi komentar lama, jumlahKomentar berisi nilai saat ini.
	F.S. : Sebuah komentar baru ditambahkan ke listKomentar, jumlahKomentar bertambah 1 untuk setiap komentar yang berhasil ditambahkan.
*/
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
		*jumlah = *jumlah + 1

		fmt.Println("Komentar berhasil ditambahkan.")
		fmt.Println("<<1. Tambah komentar lagi>>", "\t", "<<0. Kembali ke Menu>>")
		fmt.Scan(&pilih)
 	}
	    bersihLayar()
		menu()
}


// Subprogram: Menu untuk Menampilkan Semua Komentar
/*
	I.S. : Belum ada tampilan komentar yang diurutkan, variabel 'pilih' belum diisi oleh pengguna.
	F.S. : Komentar telah ditampilkan sesuai urutan (ascending/descending) atau pengguna kembali ke menu utama.
*/
func menutampilkanSemuaKomentar() {
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
		tampilkanSemuaKomentar()
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



// Subprogram: Mengurutkan komentar berdasarkan panjang teks secara ascending
/*
	I.S. : listKomentar belum terurut berdasarkan panjang teks.
	F.S. : listKomentar terurut naik berdasarkan panjang teks komentar.
*/
func sortKomentarAsc() { 
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


// Subprogram: Mengurutkan komentar berdasarkan panjang teks secara descending
/*
	I.S. : listKomentar belum terurut berdasarkan panjang teks.
	F.S. : listKomentar terurut turun berdasarkan panjang teks komentar.
*/
func sortKomentarDesc() {
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
/*
	I.S. : Komentar telah dianalisis dan disimpan dalam listKomentar, jumlahKomentar ≥ 0.
	F.S. : Komentar ditampilkan satu per satu ke layar jika ada, atau muncul pesan jika tidak ada komentar.
*/
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
/*
	I.S. : Komentar lama masih tersimpan di listKomentar
	F.S. : Komentar lama digantikan oleh komentar baru, dengan sentimen yang diperbarui
*/
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
/*
	I.S. : listKomentar berisi komentar, pengguna belum memilih komentar yang ingin dihapus
	F.S. : Jika ditemukan, komentar dihapus dari listKomentar dan jumlahKomentar berkurang 1
*/
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
/*
	I.S. : listKomentar terisi komentar, pengguna belum melakukan pencarian
	F.S. : Jika ditemukan, komentar akan dicetak dengan sentimennya
*/

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


// Subprogram: Pencarian Komentar dengan Sequential Search
/*
	I.S. : keyword merupakan kata kunci pencarian, listKomentar berisi komentar-komentar hasil input atau analisis, jumlahKomentar ≥ 0.
	F.S. : Jika ditemukan, ditampilkan komentar yang mengandung keyword (tidak case sensitive); jika tidak ditemukan, tampilkan pesan bahwa komentar tidak ditemukan.
*/
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



// Subprogram: Pencarian Komentar dengan Binary Search
/*
	I.S. : keyword merupakan kata kunci pencarian, listKomentar sudah terurut secara alfabet berdasarkan teks komentar, jumlahKomentar ≥ 0.
	F.S. : Jika ditemukan, ditampilkan komentar yang mengandung keyword (tidak case sensitive); jika tidak ditemukan, tampilkan pesan bahwa komentar tidak ditemukan.
*/
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



// Subprogram: Mengurutkan Komentar secara Alfabet
/*
	I.S. : listKomentar belum terurut berdasarkan urutan alfabet teks komentar, jumlahKomentar ≥ 0.
	F.S. : listKomentar terurut secara naik (ascending) berdasarkan urutan alfabet teks komentar.
*/
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
/*
	I.S. : listKomentar berisi komentar-komentar dengan panjang teks yang belum terurut, 
	       pengguna belum memilih mode pengurutan (ascending/descending).
	F.S. : listKomentar terurut berdasarkan panjang teks komentar sesuai mode yang dipilih 
	       ("asc" untuk ascending, "desc" untuk descending), dan semua komentar ditampilkan.
*/
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
/*
	I.S. : listKomentar berisi komentar-komentar dengan sentimen yang belum terurut,
	       pengguna belum memilih mode pengurutan ("asc" untuk ascending atau "desc" untuk descending).
	F.S. : listKomentar telah terurut berdasarkan sentimen komentar sesuai mode yang dipilih,
	       dan semua komentar berhasil ditampilkan.
*/
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


// Subprogram: Prioritas Sentimen
/*
	I.S. : Variabel 's' berisi nilai sentimen yang belum diurutkan berdasarkan prioritas.
	F.S. : Fungsi mengembalikan nilai integer yang merepresentasikan prioritas sentimen:
	       1 untuk sentimen positif, 2 untuk netral, 3 untuk negatif, dan 4 untuk lainnya.
*/
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
/*
	I.S. : listKomentar berisi kumpulan komentar dengan sentimen yang sudah dianalisis, 
	       jumlahKomentar ≥ 0, variabel penghitung statistik belum diinisialisasi.
	F.S. : Statistik jumlah komentar dengan sentimen positif, netral, dan negatif
	       ditampilkan ke layar terminal.
*/
func statistikSentimen(){
	var i int 
	var positif, netral, negatif int
	for i = 0; i < jumlahKomentar; i++ {
		switch listKomentar[i].sentimen {
		case "positif":
			positif++
		case "netral":
			netral++
		case "negatif":
			negatif++
		}
	}
	fmt.Println("Statistik Sentimen Komentar:")
	fmt.Printf("Positif: %d\n", positif)
	fmt.Printf("Netral : %d\n", netral)
	fmt.Printf("Negatif: %d\n", negatif)
}	



// Subprogram: inisialisasiDataDummy
/*
	I.S. : listKomentar dan variabel n belum berisi data komentar apapun.
	F.S. : listKomentar terisi dengan 12 komentar dummy beserta sentimennya, 
	       variabel n berisi nilai 12, dan variabel statistik positif, negatif, netral diupdate.
*/
func inisialisasiDataDummy(listKomentar *[NMAX]komentar, n *int) {
	listKomentar[0].teks, listKomentar[0].sentimen = "Desain tasnya bagus dan sangat menarik.", "positif"
        listKomentar[1].teks, listKomentar[1].sentimen = "Pelayanan customer service-nya sangat mengecewakan.", "negatif"
        listKomentar[2].teks, listKomentar[2].sentimen = "Saya puas banget belanja di sini, barangnya bagus dan designnya gemes!", "positif"
        listKomentar[3].teks, listKomentar[3].sentimen = "Sayangnya kualitas produk ini cukup buruk untuk harga segitu.", "negatif"
        listKomentar[4].teks, listKomentar[4].sentimen = "Tampilan websitenya indah dan navigasinya mudah digunakan.", "positif"
        listKomentar[5].teks, listKomentar[5].sentimen = "Makanan di resto itu tidak sesuai dengan tampilan di buku menunya.", "negatif"
        listKomentar[6].teks, listKomentar[6].sentimen = "Tempat ini sangat indah dan family friendly.", "positif"
        listKomentar[7].teks, listKomentar[7].sentimen = "Respon admin lebay dan tidak profesional saat ditanya.", "negatif"
        listKomentar[8].teks, listKomentar[8].sentimen = "Pengalaman pertama yang sangat menyenangkan, suka banget!", "positif"
        listKomentar[9].teks, listKomentar[9].sentimen = "Barang yang datang rusak dan tidak sesuai dengan pesanan saya.", "negatif"
	listKomentar[10].teks, listKomentar[10].sentimen = "Produk diterima dengan baik, namun belum sempat digunakan.", "netral"
	listKomentar[11].teks, listKomentar[11].sentimen = "Saya baru saja menerima paket ini tadi pagi.", "netral"

	*n = 12
	positif = 5
    negatif = 5
    netral = 2
}


/*
	I.S. : Program belum menerima input menu dari pengguna.
	F.S. : Menu dipilih dan prosedur yang sesuai dijalankan berdasarkan input pengguna.
*/
func main() {
	header()
	var username, password string
	login(&username, &password)
	inisialisasiDataDummy(&listKomentar, &jumlahKomentar)
	menu()
}
