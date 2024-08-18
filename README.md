# go-crud-book-api-postgres

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