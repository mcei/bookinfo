Приложение BookInfo для хранения и обработки информации о книгах

Приложение позволяет:
- добавить информацию о книге в хранилище
- обновить информацию
- удалить информацию
- запросить информацию о книге
- запросить информацию о всех книгах

Использование:

```
curl -X GET localhost:8000/books

curl -X GET localhost:8000/books/1

curl -X POST localhost:8000/books --data '{"ID":1,"Title":"Mybook1","Author":"Author1","Year":"2021"}'

curl -X PUT localhost:8000/books --data '{"ID":1,"Title":"Mybook1","Author":"Author2","Year":"2022"}'

curl -X DELETE localhost:8000/books/1
```


Notes:

- Standard Go Layout: https://github.com/golang-standards/project-layout/
- Another example: https://github.com/herryg91/go-clean-architecture/tree/main/examples/book-rest-api
- https://github.com/herryg91/go-clean-architecture/blob/main/README.md#folder-structure
