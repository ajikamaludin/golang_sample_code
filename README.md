# Golang Projects

Repositori ini berisi dua project Golang:

1. `country_group` untuk mengelompokan country daru file cities.csv kedalam output.json dengan nama negara dan jumlah kota
2. `root_words` untuk melakukan replacement akar kata dari kalimat yang diberikan

Setiap project memiliki file `main.go` yang dapat dijalankan secara independen.

## Prasyarat

Pastikan Anda telah menginstal:

-   Go (minimal versi 1.19)
-   Docker

## Menjalankan dengan `go run`

Jika tidak ingin membuat binary, Anda dapat menjalankan langsung dengan `go run`:

### 1. `country_group`

```sh
cd country_group
# Jalankan dengan go run
go run main.go
```

### 2. `root_words`

```sh
cd root_words
# Jalankan dengan go run
go run main.go
```

## Menjalankan dengan Docker

Sebagai alternatif, Anda juga dapat menjalankan setiap project menggunakan Docker.

### 1. `country_group`

```sh
docker run --rm -v "$PWD/country_group":/app -w /app golang:latest go run main.go
```

### 2. `root_words`

```sh
docker run --rm -v "$PWD/root_words":/app -w /app golang:latest go run main.go
```

## Struktur Folder

```
repo/
│-- country_group/
│   ├── main.go
│   ├── go.mod
│   ├── go.sum
│   └── ...
│
│-- root_words/
│   ├── main.go
│   ├── go.mod
│   ├── go.sum
│   └── ...
```

Dengan petunjuk di atas, Anda dapat menjalankan setiap project dengan cara yang sesuai dengan kebutuhan Anda.
