# go-crud-book-api-postgres

### Things todo list

1. Clone this repository: `git clone https://github.com/hendisantika/go-crud-book-api-postgres.git`
2. Navigate to the folder: `cd go-crud-book-api-postgres`
3. Run docker compose: `docker compose up`
4. Run the application: `go run main.go`
5. Open your terminal

#### Add New Book
```shell
curl --location 'http://127.0.0.1:8080/api/books' \
--header 'Content-Type: application/json' \
--data '{
    "Title": "Belajar GoLang",
    "Author": "Hendi Santika",
    "PublishedAt":"2024-08-18"
}'
```

### Get All Books

```shell
curl --location 'http://127.0.0.1:8080/api/books'
```

#### Get Book by ID

```shell
curl --location 'http://127.0.0.1:8080/api/books/1'
```

### Update Book

```shell
curl --location --request PUT 'http://127.0.0.1:8080/api/books/1' \
--header 'Content-Type: application/json' \
--data '{
    "Title": "Belajar GoLang Updated",
    "Author": "Hendi Santika",
    "PublishedAt":"2024-08-18"
}'
```

### Delete Book

```shell
curl --location --request DELETE 'http://127.0.0.1:8080/api/books/5'
```