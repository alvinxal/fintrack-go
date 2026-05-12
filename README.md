# Fintrack Go

Project latihan Go backend sambil membangun finance tracker sederhana.

Fokus project ini bukan cuma hasil akhir aplikasi, tapi proses belajar fundamental Go yang kepakai di backend: `struct`, `const`, `type`, `slice`, `map`, `error`, `json`, `net/http`, testing, sampai database.

## Tujuan

- belajar Go backend dari sudut pandang frontend developer
- paham dasar Go yang sering kepakai di backend
- bisa bikin HTTP server sederhana
- bisa kirim dan terima JSON
- bisa bikin CRUD transaksi
- siap untuk interview junior backend/fullstack

## Kondisi Saat Ini

Saat ini project masih berbentuk CLI sederhana.

Yang sudah ada:
- struct `Transaction`
- const `Income` dan `Expense`
- function `applyTransaction`
- function `printTransactions`
- data transaksi masih hardcoded di slice
- program sudah bisa hitung saldo akhir

Yang belum ada:
- validasi transaksi
- unit test
- JSON encode/decode
- HTTP server
- API transaksi
- database PostgreSQL

## Materi Yang Dipelajari

Roadmap belajar ada di `PROGRESS.md`.

Urutan besar materinya:
1. CLI finance tracker dasar
2. validasi transaksi
3. testing dasar
4. summary keuangan
5. kategori transaksi
6. struct dan JSON
7. dasar project Go
8. HTTP server pertama
9. JSON response dari HTTP
10. API transaksi in-memory
11. query params dan filter
12. rapikan project
13. database PostgreSQL
14. lanjutan testing

## Cara Menjalankan

Pastikan Go sudah terpasang.

```bash
go run .
```

Atau:

```bash
go run main.go
```

## Contoh Output Saat Ini

```text
----
Transaksi: Gaji
Jenis transaksi: income
Jumlah: 4500000
----
Transaksi: Makan
Jenis transaksi: expense
Jumlah: 25000
----
Saldo akhir:  4475000
```

## Struktur Saat Ini

```text
.
├── go.mod
├── main.go
├── PROGRESS.md
└── README.md
```

## Catatan

- project ini sengaja dimulai dari versi kecil dan sederhana
- perubahan dilakukan bertahap supaya tiap langkah tetap mudah dipahami
- `PROGRESS.md` jadi sumber utama untuk roadmap belajar dan checklist progress
