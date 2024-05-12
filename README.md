# Rest API Book Store

Simple Login and CRUD Book Store

## Setup

To deploy this project run

copy `.env.example` to `.env`

set Database Connection in `.env`

run

```bash
migrate -database "mysql://root@tcp(localhost:3306)/{db_name}" -path migrations up
```

if you dont install `golang-migrate` you can copy sql from `migrations` folder and run it

then

```bash
go get
```

then

```bash
go run main.go
```

## 🛠 Tech Stack

Go-Lang, Gin Framework, GORM, MySQL

## Authors

- [Radhitka Adha](https://www.github.com/radhitka)
