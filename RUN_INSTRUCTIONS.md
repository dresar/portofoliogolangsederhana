# Cara Menjalankan Aplikasi

## ✅ Solusi Error CGO

**Error yang terjadi:**
```
cgo: C compiler "gcc" not found
```

**Solusi yang sudah diterapkan:**
- Menggunakan driver SQLite pure Go (`github.com/glebarez/sqlite`)
- **TIDAK PERLU** install gcc atau set CGO_ENABLED
- **TIDAK PERLU** environment variable khusus

## 🚀 Cara Menjalankan

### 1. Install Dependencies
```bash
go mod tidy
```

### 2. Jalankan Server
```bash
go run cmd/server/main.go
```

**Sederhana! Tidak perlu set CGO_ENABLED lagi!**

### 3. Akses Aplikasi
- Website: http://localhost:8080
- Admin Login: http://localhost:8080/admin/login
- Password Admin: `admin123`

## 📝 Catatan

- Database SQLite akan dibuat otomatis sebagai `portfolio.db` saat pertama kali dijalankan
- Server berjalan di port 8080
- Semua dependency sudah pure Go, tidak memerlukan C compiler

