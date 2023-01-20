# Rest Book API Server in golang
- The Book API Server provides endpoints to create, show, read, update & delete users and books.
## Running the server
```$ git clone git@github.com:Imtiaz246/Book-Server.git```

```$ cd Book-Server```

```$ go install```

```$ go run . start```  `[run the api server]`
## To Test the API endpoints
```$ go test ./...```  `[test the API endpoints]`
## Available API Endpoints
| Method   | API Endpoint                 | Authentication Type | Payload | Description                                     |
|----------|------------------------------|---------------------|---------|-------------------------------------------------|
| `POST`   | `/api/v1/users/get-token`|                     |`payload` | Returns jwt token in response 
| `POST`   | `/api/v1/users` |                     | `payload` | Creates a user if the user information is valid
| `GET`    | `/api/v1/users`|                     |         | Returns the list of all user
| `GET`    | `/api/v1/users/{username}/books`|                     |         | Returns the book list of a specific user
| `GET`    | `/api/v1/users/{username}`|                     |  | Returns the user specified by `{username}`
| `DELETE` | `/api/v1/users/{username}`| `JWT`               | `payload` | Deletes the user specified by `{username}`
| `POST`   | `/api/v1/books`| `JWT`               | `payload` | Creates a books if the book information is valid
| `GET`    | `/api/v1/books`|                     |                 | Returns all the book list
| `GET`    | `/api/v1/books/{id}`|                     |                     | Returns the book specified by the `{id}`
| `DELETE` | `/api/v1/books/{id}` | `JWT`               |         | Deletes the book specified by the `{id}`

## Basic curl commands

1. Create a user
```shell
curl -v -H "Content-type: application/json" -X POST -d '{"username":"lakkas", "password":"1234", "organization":"Appscode Ltd", "email": "xyz@gmail.com"}' http://localhost:3000/api/v1/users
```
2. Get all user list

```shell
curl -v -H "Content-type: application/json" http://localhost:3000/api/v1/users
```
3. Get book list of a user 
```shell
 curl -v -H "Content-type: application/json" http://localhost:3000/api/v1/users/imtiaz/books
```
4. Get a user by username 
```shell
curl -v -H "Content-type: application/json" http://localhost:3000/api/v1/users/lakkas
```

5. Delete a user by username
```shell
curl -v -H "Content-type: application/json" -H "Authorization: Bearer <jwt token>" -X DELETE http://localhost:3000/api/v1/users/lakkas
```
6. Create a book
```shell
curl -v -H "Content-type: application/json" -H "Authorization: Bearer <jwt token>" -X POST -d '{ "book-name": "new book", "price": 200, "isbn": "4323-6456-4756-4564", "authors": [ { "username": "imtiaz" } ], "book-content": { "over-view": "overview", "chapters": [ { "chapter-title": "chapter 1", "chapter-content": "chapter 1 content" }, { "chapter-title": "chapter 2", "chapter-content": "chapter 2 content" } ] } }' http://localhost:3000/api/v1/books`
```
7. Get all book list
```shell
curl -v -H "Content-type: application/json" http://localhost:3000/api/v1/books
```
8. Get a book by book-id
```shell
curl -v -H "Content-type: application/json" http://localhost:3000/api/v1/books/1000
```
9. Delete a book by book-id
```shell
curl -v -H "Content-type: application/json" -H "Authorization: Bearer <jwt token>" -X DELETE http://localhost:3000/api/v1/books/1000
```
10. Get a jwt token
```shell
curl -v -H "Content-type: application/json" -d '{ "username": "imtiaz", "password": "1234" }' http://localhost:3000/api/v1/users/get-token`
```
