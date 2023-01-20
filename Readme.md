# Rest Book API Server in golang
- The Book API Server provides endpoints to create, show, read, update & delete users and books.
## Running the server
```$ git clone https://github.com/Imtiaz246/Book-Server```

```$ cd Book-Server```

```$ go install```

```$ go run . start```  `[run the api server]`
## Available API Endpoints
| Method | API Endpoint                   | Authentication Type | Payload      | Description                                     |
|------|--------------------------------|---------------------|--------------|-------------------------------------------------|
| `POST` | `/api/v1/users/get-token`        | `not required`        | `payload`      | Return jwt token in response  
| `POST` | `/api/v1/users`                 | `not required`        |` payload `     | Creates a user if the user information is valid
| `GET`  | `/api/v1/users`                 | `not required`        | `not required` | Returns the list of all user
| `POST` | `/api/v1/users/{username}/books` | `not required`        | `not required` | Retuns the book of a specific user
| `GET`  | `/api/v1/users/{username}`      | `not required`        | `payload`      | Return the user specified by `{username}`
| `DELETE` | `/api/v1/users/{username}`      | `JWT`                 | `payload `     | Deletes the user specified by `{username}`
| `POST` | `/api/v1/books`                 | `JWT`                 | `payload`      | Creates a books if the book information is valid
| `GET` | `/api/v1/books`                 | `not required`        | `not required` | Returns all the book list
| `GET` | `/api/v1/books/{id}`            | `not required`        | `not required` | Returns the book specified by the `{id}`
| `DELETE` | `/api/v1/books/{id}`            | `JWT`                 |` not required `| Deletes the book specified by the `{id}`

## Basic curl commands

1. Create a user
- `curl -v -H "Content-type: application/json" -X POST -d '{"username":"lakkas", "password":"1234", "organization":"Appscode Ltd", "email": "xyz@gmail.com"}' http://localhost:3000/api/v1/users`

2. Get all user list
- `curl -v -H "Content-type: application/json" http://localhost:3000/api/v1/users`

3. Get book list of a user
- `curl -v -H "Content-type: application/json" http://localhost:3000/api/v1/users/imtiaz/books`

4. Get a user by username
- `curl -v -H "Content-type: application/json" http://localhost:3000/api/v1/users/lakkas`

5. Delete a user by username
- `curl -v -H "Content-type: application/json" -H "Authorization: Bearer <jwt token>" -X DELETE http://localhost:3000/api/v1/users/lakkas`

6. Create a book
- `curl -v -H "Content-type: application/json" -H "Authorization: Bearer <jwt token>" -X POST -d '{ "book-name": "new book", "price": 200, "isbn": "4323-6456-4756-4564", "authors": [ { "username": "imtiaz" } ], "book-content": { "over-view": "overview", "chapters": [ { "chapter-title": "chapter 1", "chapter-content": "chapter 1 content" }, { "chapter-title": "chapter 2", "chapter-content": "chapter 2 content" } ] } }' http://localhost:3000/api/v1/books`

7. Get all book list
- `curl -v -H "Content-type: application/json" http://localhost:3000/api/v1/books`

8. Get a book by book-id
- `curl -v -H "Content-type: application/json" http://localhost:3000/api/v1/books/1000`

9. Delete a book by book-id
- `curl -v -H "Content-type: application/json" -H "Authorization: Bearer <jwt token>" -X DELETE http://localhost:3000/api/v1/books/1000`

10. Get a jwt token
- `curl -v -H "Content-type: application/json" -d '{ "username": "imtiaz", "password": "1234" }' http://localhost:3000/api/v1/users/get-token`







