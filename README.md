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

## API Reference

#### Register

```http
  POST /api/register
```

| Body       | Type     | Description   |
| :--------- | :------- | :------------ |
| `name`     | `string` | **Required**. |
| `username` | `string` | **Required**. |
| `password` | `string` | **Required**. |

#### Login

```http
  POST /api/login
```

| Body       | Type     | Description   |
| :--------- | :------- | :------------ |
| `username` | `string` | **Required**. |
| `password` | `string` | **Required**. |

#### List Books

```http
  GET /api/books
```

| Parameter      | Type     | Description      |
| :------------- | :------- | :--------------- |
| `query`        | `string` |                  |
| `is_published` | `string` | **Only 1 and 0** |

#### Detail Book

```http
  GET /api/books/detail/{id}
```

| Parameter | Type     | Description  |
| :-------- | :------- | :----------- |
| `id`      | `string` | **Required** |

#### Delete Book

```http
  DELETE /api/books/delete/{id}
```

| Parameter | Type     | Description  |
| :-------- | :------- | :----------- |
| `id`      | `string` | **Required** |

#### Add Book

```http
  POST /api/books/add
```

| Body           | Type      | Description  |
| :------------- | :-------- | :----------- |
| `title`        | `string`  | **Required** |
| `total_pages`  | `int`     | **Required** |
| `cover`        | `file`    | **Required** |
| `author`       | `string`  | **Required** |
| `publisher`    | `string`  | **Required** |
| `is_published` | `boolean` | **Required** |

#### Update Book

```http
  POST /api/books/update/{id}
```

| Parameter | Type     | Description  |
| :-------- | :------- | :----------- |
| `id`      | `string` | **Required** |

| Body           | Type      | Description  |
| :------------- | :-------- | :----------- |
| `title`        | `string`  | **Required** |
| `total_pages`  | `int`     | **Required** |
| `cover`        | `file`    | **optional** |
| `author`       | `string`  | **Required** |
| `publisher`    | `string`  | **Required** |
| `is_published` | `boolean` | **Required** |

## Authors

- [Radhitka Adha](https://www.github.com/radhitka)
