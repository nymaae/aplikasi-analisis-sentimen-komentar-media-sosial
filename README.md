# Aplikasi Analisis Sentimen Komentar Media Sosial

Aplikasi ini ini dirancang sebagai alat bantu untuk mengelola dan menganalisis komentar pengguna dari media sosial berdasarkan tiga kategori sentiment yaitu positif, netral, dan negatif. Data utama yang digunakan adalah daftar komentar dan analisis sentimennya. Fokus utama aplikasi adalah melakukan analisis sentimen terhadap komentar yang masuk, sehingga komentar yang mengandung unsur negatif atau tidak pantas dapat diidentifikasi lebih awal.

Aplikasi ini ditujukan bagi pengguna media sosial yang bertugas untuk menyaring dan memantau kualitas percakapan di platform. Dengan menggunakan aplikasi ini, pengguna dapat secara aktif menambahkan komentar baru, mengedit komentar yang ada, menghapus komentar yang tidak relevan, serta melakukan pencarian dan pengurutan komentar berdasarkan kriteria tertentu. Selain itu, aplikasi juga menyediakan statistik sentimen untuk membantu pengguna dalam menilai interaksi sosial di platform tersebut.

Adapun fitur-fitur yang dikembangkan pada aplikasi ini, yaitu:

1.  **Fitur Login**
    Sebelum mengakses menu utama, pengguna diwajibkan untuk melakukan proses login terlebih dahulu. Langkah ini bertujuan untuk menjaga keamanan serta memastikan bahwa penggunaan aplikasi dilakukan secara tertib dan terkontrol.

2.  **Fitur Menu Utama**
    Setelah proses login berhasil, aplikasi akan menampilkan menu utama yang memuat berbagai fitur yang dapat diakses oleh pengguna. Melalui menu ini, pengguna dapat memilih berbagai fungsi sesuai kebutuhan, seperti menambahkan komentar, menampilkan komentar, mengedit komentar, menghapus komentar, melakukan pencarian komentar, mengurutkan komentar, serta melihat statistik komentar berdasarkan analisis sentimen.

3.  **Fitur Tambah Komentar**
    Pengguna dapat menambahkan komentar, setelah pengguna menambahkan komentar kemudian sistem akan secara otomatis menganalisis isi komentar tersebut dan mengkategorikannya ke dalam salah satu dari tiga jenis sentimen, yaitu: positif, netral, atau negatif.

4.  **Fitur Tampilkan Komentar**
    Komentar yang telah ditambahkan dapat ditampilkan kembali sesuai urutan saat dimasukkan ke dalam sistem, komentar yang dimasukkan lebih awal akan muncul lebih dahulu. Selain itu, pengguna juga dapat mengurutkan komentar berdasarkan panjang teks komentar dalam urutan menaik (ascending) dari komentar terpendek ke terpanjang, atau menurun (descending) dari komentar terpanjang ke terpendek. Untuk proses pengurutan ini, aplikasi menggunakan algoritma Selection Sort dan Insertion Sort.

5.  **Fitur Ubah Komentar (Edit)**
    Pengguna diberikan opsi untuk mengubah komentar yang telah dimasukkan sebelumnya. Caranya, pengguna terlebih dahulu memasukkan komentar lengkap yang ingin diubah, lalu mengisi kolom baru dengan komentar pengganti. Proses ini melibatkan dua langkah:
    * Masukkan komentar yang ingin diubah :
    * Masukkan komentar baru :

6.  **Fitur Hapus Komentar**
    Pengguna dapat menghapus komentar tertentu dengan memasukkan teks komentar lengkap yang ingin dihapus. Langkah penghapusan dilakukan dengan memasukkan komentar yang ingin dihapus, lalu setelah itu sistem akan menghapus komentar tersebut dan menampilkan pesan `komentar berhasil dihapus`.

7.  **Fitur Cari Komentar**
    Melalui fitur ini, pengguna dapat melakukan pencarian komentar secara spesifik berdasarkan kata kunci yang dimasukkan. Setelah memasukkan kata kunci pencarian, sistem akan memberikan pilihan metode pencarian yang dapat dipilih oleh pengguna, yaitu:
    1.  Sequential Search
    2.  Binary Search (data harus diurutkan secara alfabet terlebih dahulu)
    Pengguna diminta untuk memasukkan pilihan metode pencarian yang diinginkan. Setelah metode pencarian dipilih, sistem akan menampilkan hasil komentar yang sesuai dengan kata kunci berdasarkan metode pencarian yang digunakan.

8.  **Fitur Urutkan Panjang Komentar**
    Sistem akan menampilkan pilihan untuk mengurutkan komentar berdasarkan sentimen dengan opsi ascending (menaik) atau descending (menurun). Pengguna memilih salah satu opsi tersebut, kemudian sistem akan mengurutkan komentar sesuai pilihan dan menampilkan semua komentar yang sudah terurut. Urutan komentar berdasarkan panjang teks, baik secara menaik (dari komentar terpendek ke terpanjang) maupun menurun (dari komentar terpanjang ke terpendek).

9.  **Fitur Urutkan Sentimen Komentar**
    Sistem akan menampilkan pilihan untuk mengurutkan komentar berdasarkan sentimen dengan opsi ascending (menaik) atau descending (menurun). Pengguna memilih salah satu opsi tersebut, kemudian sistem akan mengurutkan komentar sesuai pilihan dan menampilkan semua komentar yang sudah terurut berdasarkan sentiment komentarnya.
    Jika memilih ascending (asc), urutan komentar dimulai dari sentimen negatif, diikuti netral, dan terakhir positif. Sebaliknya, jika memilih descending (desc), urutan dimulai dari positif, kemudian netral, dan terakhir negatif.

10. **Fitur Statistik Sentimen Komentar**
    Fitur ini menampilkan jumlah komentar yang termasuk ke dalam masing-masing kategori sentimen secara keseluruhan, yaitu komentar positif, negatif dan netral. Statistik ini akan diperbarui secara otomatis setiap kali komentar ditambahkan, diubah, atau dihapus, sehingga pengguna dapat memantau kondisi percakapan secara real-time.

## Deskripsi Penggunaan Aplikasi Analisis Sentimen Komentar Media Sosial

Aplikasi ini dirancang untuk membantu pengguna dalam mengelola dan menganalisis komentar dari media sosial melalui menu-menu yang terstruktur dan mudah digunakan. Berikut adalah panduan langkah demi langkah penggunaan aplikasi:

1.  **Login Pengguna**
    - Untuk memulai, pengguna harus masuk ke dalam sistem.
    - Masukkan Username: `kamari`
    - Masukkan Password: `kamari`
    - Tekan enter untuk melanjutkan dan keterangan yang akan muncul yaitu `login berhasil`.

2.  **Tampilan Menu Utama**
    - Setelah berhasil login, pengguna akan diarahkan ke tampilan menu utama.
    - Pada menu ini, pengguna dapat memilih berbagai fungsi berikut:
        * Tambah Komentar
        * Tampilkan Semua Komentar
        * Edit Komentar
        * Hapus Komentar
        * Cari Komentar
        * Urutkan Panjang Komentar
        * Urutkan Sentimen
        * Statistik Sentimen

3.  **Menu: Tambah Komentar**
    - Pilih opsi "Tambah Komentar" pada menu utama.
    - Sistem akan menampilkan pesan `"Masukkan komentar:"`
    - Pengguna memasukkan komentar yang diinginkan, lalu tekan Enter.
    - Sistem akan menampilkan pesan `"Komentar berhasil ditambahkan."`
    - Setelah itu, pengguna diberikan pilihan untuk menambahkan komentar lagi, atau kembali ke menu utama.

    Untuk menjaga konsistensi pemrosesan, apabila komentar mengandung lebih dari satu kata, gunakan tanda garis bawah (_) sebagai pengganti spasi.

    Contoh:
    `Layanan_sangat_memuaskan`
    `Produk_tidak_sesuai_deskripsi`

4.  **Menu: Tampilkan Semua Komentar**
    - Pilih opsi "Tampilkan Semua Komentar" pada menu utama.
    - Sistem akan menampilkan seluruh komentar yang telah dimasukkan oleh pengguna beserta kategori sentimennya (positif, netral, atau negatif).

5.  **Menu: Edit Komentar**
    - Pilih opsi "Edit Komentar" pada menu utama.
    - Sistem akan meminta pengguna untuk memasukkan komentar yang ingin diubah secara lengkap agar dapat dikenali dengan tepat.
    - Setelah itu, pengguna diminta memasukkan komentar baru sebagai pengganti komentar lama.
    - Jika sudah selesai, perubahan akan disimpan dan sistem menampilkan pesan bahwa komentar berhasil diubah.

6.  **Menu: Hapus Komentar**
    - Pilih opsi "Hapus Komentar" pada menu utama.
    - Sistem akan menampilkan perintah: `"Masukkan komentar yang ingin dihapus:"`
    - Pengguna harus memasukkan komentar secara lengkap agar sistem dapat mencocokkannya dengan komentar yang ada.
    - Setelah komentar ditemukan dan dihapus, sistem akan menampilkan pesan: `"Komentar berhasil dihapus."`

7.  **Menu: Cari Komentar**
    - Pilih opsi "Cari Komentar" pada menu utama.
    - Sistem akan meminta pengguna untuk memasukkan kata kunci pencarian dengan perintah: `"Masukkan kata kunci:"`
    - Setelah itu, sistem akan menampilkan pilihan metode pencarian:
        1.  Sequential Search
        2.  Binary Search (urutan alfabet)
    - Sistem akan menampilkan perintah: `"Masukkan pilihan:"`
    - Pengguna memilih metode pencarian sesuai keinginan. Setelah pilihan dimasukkan, sistem akan menampilkan hasil komentar yang sesuai dengan kata kunci dan metode pencarian yang dipilih.

8.  **Menu: Urutkan Panjang Teks**
    - Pilih opsi "Urutkan Panjang Teks" pada menu utama.
    - Sistem akan menampilkan perintah: `"Urutkan berdasarkan panjang komentar (asc/desc):"`
    - Jika pengguna memilih berdasarkan ascending, sistem akan mengurutkan dari komentar terpendek ke terpanjang dan descending untuk mengurutkan dari komentar terpanjang ke terpendek.
    - Setelah pilihan dimasukkan, sistem akan menampilkan pesan: `"Komentar berhasil diurutkan."`
    - Sistem kemudian menampilkan semua komentar yang telah diurutkan lengkap dengan kategori sentimennya masing-masing.

9.  **Menu: Urutkan Sentimen**
    - Pilih opsi "Urutkan Sentimen Komentar" pada menu utama.
    - Sistem akan menampilkan perintah: `"Urutkan berdasarkan sentimen (asc/desc):"`
    - Pengguna memilih opsi:
        * Ascending untuk mengurutkan komentar berdasarkan sentimen dari negatif → netral → positif.
        * Descending untuk mengurutkan dari positif → netral → negatif.
    - Setelah pilihan dimasukkan, sistem akan menampilkan pesan: `"Komentar berhasil diurutkan."`
    - Sistem kemudian menampilkan semua komentar yang telah diurutkan sesuai kategori sentimen, lengkap dengan isi komentar dan label sentimennya masing-masing.

10. **Menu: Statistik Sentimen**
    - Pilih opsi "Statistik Sentimen Komentar" pada menu utama.
    - Sistem akan secara otomatis menghitung dan menampilkan jumlah komentar berdasarkan masing-masing kategori sentimen.
    - Statistik yang ditampilkan mencakup:
        * Positif: `[jumlah komentar positif]`
        * Netral: `[jumlah komentar netral]`
        * Negatif: `[jumlah komentar negatif]`
    - Informasi ini membantu pengguna untuk memahami proporsi sentimen dalam seluruh komentar yang telah dianalisis.

Dengan mengikuti panduan ini, pengguna dapat menggunakan seluruh fitur yang tersedia dalam aplikasi melalui menu-menu yang telah disediakan untuk mengelola komentar secara efisien dan terstruktur.
