# Fintrack - Catatan Belajar Go

## Tujuan

Belajar Go backend dari sudut pandang frontend developer sambil membangun project finance tracker sederhana.

Target:
- paham dasar Go yang kepakai di backend
- bisa bikin HTTP server sederhana
- bisa kirim dan terima JSON
- bisa bikin CRUD transaksi
- siap untuk interview junior backend/fullstack

---

## Kondisi Project Saat Ini

Saat ini project masih berbentuk CLI sederhana.

Yang sudah ada:
- struct `Transaction`
- const `Income` dan `Expense`
- function `applyTransaction`
- function `printTransactions`
- data transaksi masih hardcoded di slice
- program sudah bisa hitung saldo akhir

---

## Project Todolist

### 0. CLI Finance Tracker Dasar

- [x] Buat struct `Transaction`
  Objective: punya model data untuk satu transaksi.
  Output: ada field `Nama`, `Tipe`, dan `Amount`.

- [x] Buat tipe transaksi `income` dan `expense`
  Objective: transaksi bisa dibedakan sebagai pemasukan atau pengeluaran.
  Output: ada const `Income` dan `Expense`.

- [x] Hitung saldo akhir dari daftar transaksi
  Objective: saldo bertambah kalau income dan berkurang kalau expense.
  Output: program menampilkan saldo akhir.

- [x] Tampilkan daftar transaksi di terminal
  Objective: user bisa melihat semua transaksi yang diproses.
  Output: setiap transaksi tampil dengan nama, tipe, dan jumlah.

### 1. Validasi Transaksi

- [ ] Tolak transaksi dengan tipe tidak valid
  Objective: program tidak menerima tipe selain `income` atau `expense`.
  Output: jika tipe salah, tampilkan error yang jelas.

- [ ] Tolak transaksi dengan amount kurang dari atau sama dengan 0
  Objective: mencegah data transaksi tidak masuk akal.
  Output: transaksi invalid tidak ikut dihitung.

- [ ] Ubah `applyTransaction` agar bisa return error
  Objective: mulai belajar error handling Go lewat kebutuhan project.
  Output: function return `(int, error)`.

- [ ] Buat custom type untuk tipe transaksi
  Objective: belajar `type` di Go agar nilai `income` dan `expense` lebih aman dipakai.
  Output: ada type seperti `TransactionType`.

- [ ] Tambahkan method validasi pada `Transaction`
  Objective: mulai kenal method pada struct dan merapikan logic validasi.
  Output: ada method seperti `func (t Transaction) Validate() error`.

### 2. Testing Dasar

- [ ] Tambahkan unit test untuk `applyTransaction`
  Objective: mulai belajar testing function kecil sebelum masuk HTTP.
  Output: ada test untuk income, expense, dan invalid input.

- [ ] Tambahkan unit test untuk validasi transaksi
  Objective: memastikan rule validasi tidak gampang rusak saat code berubah.
  Output: invalid type dan invalid amount punya test.

- [ ] Kenali command `go test`
  Objective: terbiasa menjalankan test dari terminal.
  Output: test bisa dijalankan dengan `go test ./...`.

### 3. Summary Keuangan

- [ ] Hitung total income
  Objective: tahu total pemasukan yang masuk.
  Output: terminal menampilkan `Total income`.

- [ ] Hitung total expense
  Objective: tahu total pengeluaran yang keluar.
  Output: terminal menampilkan `Total expense`.

- [ ] Tampilkan ringkasan saldo
  Objective: saldo akhir bisa diverifikasi dari summary.
  Output: tampil `income - expense = saldo akhir`.

### 4. Kategori Transaksi

- [ ] Tambahkan field `Category` ke `Transaction`
  Objective: transaksi bisa dikelompokkan.
  Output: tiap transaksi punya kategori seperti `food`, `salary`, `transport`.

- [ ] Tampilkan transaksi per kategori
  Objective: tahu uang paling banyak keluar di kategori apa.
  Output: ada summary per kategori.

- [ ] Buat summary kategori dengan `map`
  Objective: belajar `map[string]int` lewat kebutuhan nyata.
  Output: kategori dipetakan ke total amount.

### 5. Struct Dan JSON

- [ ] Buat struct response sederhana
  Objective: mulai menyiapkan data agar bisa dikirim sebagai JSON.
  Output: ada struct seperti `GreetingResponse` atau `TransactionResponse`.

- [ ] Tambahkan JSON tags pada struct
  Objective: paham kenapa field Go perlu kapital tapi output JSON biasanya huruf kecil.
  Output: ada tag seperti ``json:"message"``.

- [ ] Encode struct ke JSON
  Objective: bisa mengubah data Go menjadi JSON.
  Output: data tampil dalam format JSON string.

- [ ] Decode JSON ke struct
  Objective: belajar menerima input JSON, bukan hanya mengirim response JSON.
  Output: data JSON bisa dibaca ke struct Go.

### 6. Dasar Project Go

- [ ] Pahami `go mod init` dan `go mod tidy`
  Objective: tahu cara Go mengelola module dan dependency.
  Output: paham fungsi `go.mod` dan `go.sum`.

- [ ] Rapikan file jadi lebih dari satu file bila perlu
  Objective: belajar bahwa satu package bisa punya beberapa file.
  Output: code tidak harus menumpuk di satu file.

- [ ] Kenali `defer`
  Objective: paham pattern umum untuk cleanup resource.
  Output: tahu kapan pakai `defer`, misalnya saat menutup `rows` nanti.

- [ ] Kenali pointer dasar
  Objective: paham beda value dan pointer untuk kebutuhan backend nanti.
  Output: tahu kapan function atau method perlu menerima pointer.

### 7. HTTP Server Pertama

- [ ] Buat endpoint `GET /hello`
  Objective: belajar server HTTP paling dasar di Go.
  Output: browser atau Postman menampilkan response dari server.

- [ ] Jalankan server di port `8080`
  Objective: bisa menjalankan backend lokal.
  Output: `localhost:8080` bisa diakses.

- [ ] Handle error dari `http.ListenAndServe`
  Objective: tidak mengabaikan error startup server.
  Output: jika server gagal start, error tampil jelas.

- [ ] Pastikan route ditulis benar di `http.HandleFunc`
  Objective: memahami pattern route dasar di `net/http`.
  Output: route memakai path seperti `"/hello"`, bukan `"GET /hello"`.

- [ ] Cek method request di dalam handler
  Objective: paham bahwa `http.HandleFunc` cocokkan path, bukan method.
  Output: handler bisa menolak method yang tidak sesuai.

### 8. JSON Response Dari HTTP

- [ ] Ubah response plain text jadi JSON
  Objective: endpoint tidak lagi kirim string biasa.
  Output: endpoint return object JSON valid.

- [ ] Set header `Content-Type: application/json`
  Objective: client tahu bahwa response adalah JSON.
  Output: header response benar.

- [ ] Tambahkan timestamp ke response
  Objective: belajar pakai `time.Now().UTC().Format(time.RFC3339)`.
  Output: JSON punya field `timestamp`.

- [ ] Encode JSON langsung ke `ResponseWriter`
  Objective: belajar `json.NewEncoder(w).Encode(...)`.
  Output: handler bisa kirim JSON response langsung.

- [ ] Buat format error response JSON sederhana
  Objective: client dapat bentuk error yang konsisten.
  Output: error API dikirim sebagai JSON, bukan plain text acak.

### 9. API Transaksi In-Memory

- [ ] Buat endpoint `GET /transactions`
  Objective: expose daftar transaksi sebagai API.
  Output: endpoint return semua transaksi dalam JSON.

- [ ] Buat endpoint `POST /transactions`
  Objective: bisa menambah transaksi baru lewat request.
  Output: transaksi baru masuk ke slice.

- [ ] Validasi body request
  Objective: data yang masuk tidak asal diterima.
  Output: request invalid dikembalikan sebagai error.

- [ ] Return status code yang sesuai
  Objective: mulai terbiasa dengan pola API yang benar.
  Output: pakai `200`, `201`, `400`, `405`, atau `500` sesuai kondisi.

- [ ] Buat endpoint `GET /transactions/{id}`
  Objective: belajar ambil satu resource secara spesifik.
  Output: satu transaksi bisa diambil berdasarkan id.

- [ ] Buat endpoint `DELETE /transactions/{id}`
  Objective: mulai melengkapi CRUD dasar.
  Output: transaksi bisa dihapus.

- [ ] Buat endpoint `PUT /transactions/{id}` atau `PATCH /transactions/{id}`
  Objective: mulai belajar update data lewat API.
  Output: transaksi bisa diubah.

### 10. Query Params Dan Filter

- [ ] Baca query params dari URL
  Objective: bisa ambil nilai seperti `?type=income`.
  Output: handler bisa membaca filter dari request.

- [ ] Filter transaksi berdasarkan tipe
  Objective: user bisa lihat income saja atau expense saja.
  Output: hasil API berubah sesuai query.

- [ ] Filter transaksi berdasarkan kategori
  Objective: API lebih berguna untuk kebutuhan finance tracker.
  Output: transaksi bisa difilter per kategori.

### 11. Rapikan Project

- [ ] Pecah code `main.go` yang terlalu besar
  Objective: codebase lebih gampang dibaca dan dikembangkan.
  Output: logic dibagi ke file atau package yang rapi.

- [ ] Pisahkan model, handler, dan business logic
  Objective: mulai belajar separation of concerns.
  Output: project tidak menumpuk semua logic di satu file.

- [ ] Kenali `context.Context` dari request
  Objective: paham dasar context sebelum dipakai ke database.
  Output: tahu fungsi `r.Context()` pada handler.

### 12. Database

- [ ] Kenali package `database/sql`
  Objective: paham jalur standar akses database di Go.
  Output: tahu peran `sql.DB`, `Query`, `QueryRow`, dan `Exec`.

- [ ] Tambahkan konfigurasi database dari environment variable
  Objective: belajar pattern config backend sederhana.
  Output: connection string dibaca dari env, bukan hardcoded.

- [ ] Sambungkan project ke PostgreSQL
  Objective: mulai memakai penyimpanan data sungguhan.
  Output: aplikasi bisa membuka koneksi database.

- [ ] Buat tabel `transactions`
  Objective: mulai belajar schema database untuk project.
  Output: ada tabel dengan kolom yang sesuai.

- [ ] Pindahkan data dari hardcoded slice ke PostgreSQL
  Objective: data tidak hilang saat program berhenti.
  Output: transaksi disimpan di database.

- [ ] Implementasi `GET /transactions` dari database
  Objective: API membaca data langsung dari DB.
  Output: endpoint tidak lagi bergantung pada slice.

- [ ] Implementasi `POST /transactions` ke database
  Objective: transaksi baru tersimpan permanen.
  Output: insert ke PostgreSQL berhasil.

- [ ] Implementasi `GET /transactions/{id}` dari database
  Objective: satu transaksi bisa diambil langsung dari DB.
  Output: query by id berhasil.

- [ ] Implementasi update dan delete ke database
  Objective: CRUD transaksi lengkap saat sudah pakai PostgreSQL.
  Output: update dan delete berjalan dari API ke DB.

### 13. Lanjutan Testing

- [ ] Tambahkan test untuk handler HTTP
  Objective: belajar test API tanpa harus selalu menjalankan server manual.
  Output: handler penting punya test dengan `httptest`.

- [ ] Tambahkan test untuk logic database bila sudah stabil
  Objective: menjaga perubahan besar tetap aman saat project makin kompleks.
  Output: area penting di layer data punya test.

---

## Catatan Dari Latihan

- `http.HandleFunc` untuk route dasar cukup pakai path, misalnya `"/hello"`.
- method request dicek di dalam handler, bukan di string route.
- `http.ListenAndServe` itu blocking, jadi jangan dipanggil dua kali sembarangan.
- error dari server startup harus dicek.
- `ResponseWriter` hanya hidup di scope handler.
- JSON response butuh header `Content-Type: application/json`.
- field struct perlu huruf kapital agar bisa di-encode oleh package `encoding/json`.
- `go test` enak dipakai sejak logic masih kecil, jangan tunggu project besar.
- `defer` dan `context` akan makin penting saat mulai pakai database.

---

## Fokus Terdekat

Fokus sekarang:
1. selesaikan `Validasi Transaksi`
2. tambah `Testing Dasar` untuk `applyTransaction`
3. lanjut ke `Summary Keuangan`
4. baru masuk `Struct Dan JSON` lalu HTTP

---

*Last updated: 2026-05-12*
