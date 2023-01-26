# Rest Book API Server in golang
- The Book API Server provides endpoints to create, show, read, update & delete users and books.
## Running the server
```$ git clone git@github.com:Imtiaz246/Book-Server.git```

```$ cd Book-Server```

```$ go mod download```

```$ go run . start```  `[run the api server]`
## To Test the API endpoints
```$ go test ./...```  `[test the API endpoints]`

## Run the server in Docker
```$ docker volume create bs-backup``` ```[creates a volume for backup]```

```$ docker build -t book-server-img .``` ```[creates the image named book-server-img]```

```$ docker run -p 3000:3000 -v bs-backup:/root/BackupFiles --name bookserver book-server-img```

## Get the Docker image
```$ docker pull imtiazcho/book-server:latest```

## Available API Endpoints
| Method   | API Endpoint                     | Authentication Type | Payload   | Description                                     |
|----------|----------------------------------|--------------------|-----------|-------------------------------------------------|
| `POST`   | `/api/v1/users/get-token`        |                    | `payload` | Returns jwt token in response 
| `POST`   | `/api/v1/users`                  |                    | `payload` | Creates a user if the user information is valid
| `GET`    | `/api/v1/users`                  |                    |           | Returns the list of all user
| `GET`    | `/api/v1/users/{username}/books` |                    |           | Returns the book list of a specific user
| `GET`    | `/api/v1/users/{username}`       |                    |           | Returns the user specified by `{username}`
| `DELETE` | `/api/v1/users/{username}`       | `JWT`              | `payload` | Deletes the user specified by `{username}`
|`PUT`     | `/api/v1/users/{username}`       | `JWT`              | `payload` | Updates the user information specified by `{username}`
| `POST`   | `/api/v1/books`                  | `JWT`              | `payload` | Creates a books if the book information is valid
| `GET`    | `/api/v1/books`                  |                    |           | Returns all the book list
| `GET`    | `/api/v1/books/{id}`             |                    |           | Returns the book specified by the `{id}`
| `DELETE` | `/api/v1/books/{id}`             | `JWT`              |           | Deletes the book specified by the `{id}`
|`PUT`     | `/api/v1/books/{id}`             | `JWT`              | `payload` | Updates the book information specified by `{id}`

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
6. Update a user by username
```shell
curl -v -H "Content-type: application/json" -H "Authorization: Bearer <jwt token>" -d '{ "username" : "lakkas updated", "password" : "1234", "organization" : "Appscode Ltd", "email": "xyz@gmail.com" }' -X PUT http://localhost:3000/api/v1/users/lakkas
```
7. Create a book
```shell
curl -v -H "Content-type: application/json" -H "Authorization: Bearer <jwt token>" -X POST -d '{ "book-name": "new book", "price": 200, "isbn": "4323-6456-4756-4564", "authors": [ { "username": "imtiaz" } ], "book-content": { "over-view": "overview", "chapters": [ { "chapter-title": "chapter 1", "chapter-content": "chapter 1 content" }, { "chapter-title": "chapter 2", "chapter-content": "chapter 2 content" } ] } }' http://localhost:3000/api/v1/books`
```
8. Get all book list
```shell
curl -v -H "Content-type: application/json" http://localhost:3000/api/v1/books
```
9. Get a book by book-id
```shell
curl -v -H "Content-type: application/json" http://localhost:3000/api/v1/books/1000
```
10. Delete a book by book-id
```shell
curl -v -H "Content-type: application/json" -H "Authorization: Bearer <jwt token>" -X DELETE http://localhost:3000/api/v1/books/1000
```
11. Update a book by book-id
```shell
curl -v -H "Content-type: application/json" -H "Authorization: Bearer <jwt token>" -X PUT -d '{ "book-name": "updated book curl", "price": 200, "isbn": "4323-6456-4756-4564", "authors": [ { "username": "imtiaz" } ], "book-content": { "over-view": "Haire over-view", "chapters": [ { "chapter-title": "chapter 1", "chapter-content": "chapter 1 content" }, { "chapter-title": "chapter 2", "chapter-content": "chapter 2 content" } ] } }' http://localhost:3000/api/v1/books/1000
```

12. Get a jwt token
```shell
curl -v -H "Content-type: application/json" -d '{ "username": "imtiaz", "password": "1234" }' http://localhost:3000/api/v1/users/get-token`
```
