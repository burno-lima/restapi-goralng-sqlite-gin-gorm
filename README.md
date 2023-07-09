# REST API IN GOLANG FOR BOOKS

* Execute
```bash
go run .
```

* Curl's

get_books

```curl
curl --request GET \
  --url http://localhost:8080/books
```

post_books
```curl
curl --request POST \
  --url http://localhost:8080/books \
  --header 'Content-Type: application/json' \
  --data '{
	"title": "Domain-Driven Design", 
	"author": "Eric Evann"
}'
```

get_books_id
```curl
curl --request GET \
  --url http://localhost:8080/books/1
```

ptch_books_id
```curl
curl --request PATCH \
  --url http://localhost:8080/books/4 \
  --header 'Content-Type: application/json' \
  --data '{
	"title": "Domain-Driven Design", 
	"author": "Eric Evan"
}'
```

delete_books_id
```curl
curl --request DELETE \
  --url http://localhost:8080/books/3
```
