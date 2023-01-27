package router

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/Imtiaz246/Book-Server/database"
	"github.com/stretchr/testify/require"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

type Response struct {
	Status  string `json:"status"`
	Message any    `json:"message""`
}

type Test struct {
	Method             string
	Path               string
	Body               io.Reader
	ExpectedStatusCode int
	ExpectedResponse   Response
}

var JWTToken string

func init() {
	database.NewTestDb()
}

func TestGetToken(t *testing.T) {
	// gets the response from http
	var testResponse Response
	jwtToken := ""
	tt := Test{
		Method: "POST",
		Path:   "/api/v1/users/get-token",
		Body: bytes.NewReader([]byte(`
					{
						"username": "imtiaz",
						"password": "1234"
					}`)),
		ExpectedStatusCode: 200,
		ExpectedResponse: Response{
			Status: "success",
		},
	}
	req, err := http.NewRequest(tt.Method, tt.Path, tt.Body)
	require.Equal(t, err, nil)

	res := httptest.NewRecorder()
	Router().ServeHTTP(res, req)

	body, _ := io.ReadAll(res.Body)
	err = json.Unmarshal(body, &testResponse)

	require.Equal(t, err, nil)
	// test code
	require.Equal(t, tt.ExpectedStatusCode, res.Code)

	jwtToken = fmt.Sprintf("%v", testResponse.Message)
	JWTToken = "Bearer " + jwtToken
}

func TestCreateUser(t *testing.T) {
	tests := []Test{
		{
			Method: "POST",
			Path:   "/api/v1/users",
			Body: bytes.NewReader([]byte(`{
    				"username" : "test_akkas",
    				"password" : "1234",
					"organization" : "Appscode Ltd",
    				"email": "xyz@gmail.com"}`)),
			ExpectedStatusCode: 200,
			ExpectedResponse: Response{
				Status: "success",
				Message: map[string]any{
					"id":           float64(101),
					"role":         "user",
					"username":     "test_akkas",
					"organization": "Appscode Ltd",
					"password":     "1234",
					"email":        "xyz@gmail.com",
					"name":         "",
				},
			},
		},
		{
			Method: "POST",
			Path:   "/api/v1/users",
			Body: bytes.NewReader([]byte(`
			{
					"username" : "test_akkas",
					"password" : "1234",
					"organization" : "Appscode Ltd",
					"email": "xyz@gmail.com"
			}`)),
			ExpectedStatusCode: 400,
			ExpectedResponse: Response{
				Status:  "failed",
				Message: "duplicate username",
			},
		},
	}

	// gets the response from http
	var testResponse Response
	for _, test := range tests {
		req, err := http.NewRequest(test.Method, test.Path, test.Body)
		if err != nil {
			t.Fatal(err)
		}
		res := httptest.NewRecorder()
		Router().ServeHTTP(res, req)

		body, _ := io.ReadAll(res.Body)
		err = json.Unmarshal(body, &testResponse)
		require.Equal(t, err, nil)
		// test code
		require.Equal(t, test.ExpectedStatusCode, res.Code)
		// test response body
		require.Equal(t, test.ExpectedResponse, testResponse)
	}
}

func TestGetUserList(t *testing.T) {
	tests := []Test{
		{
			Method:             "GET",
			Path:               "/api/v1/users",
			Body:               nil,
			ExpectedStatusCode: 200,
			ExpectedResponse: Response{
				Status: "success",
				Message: []any{
					map[string]any{
						"id":           float64(100),
						"role":         "admin",
						"username":     "imtiaz",
						"organization": "",
						"email":        "",
						"name":         "",
					},
					map[string]any{
						"id":           float64(101),
						"role":         "user",
						"username":     "test_akkas",
						"organization": "Appscode Ltd",
						"email":        "xyz@gmail.com",
						"name":         "",
					},
				},
			},
		},
	}

	// gets the response from http
	var testResponse Response
	for _, test := range tests {
		req, err := http.NewRequest(test.Method, test.Path, test.Body)
		if err != nil {
			t.Fatal(err)
		}
		res := httptest.NewRecorder()
		Router().ServeHTTP(res, req)

		body, _ := io.ReadAll(res.Body)
		err = json.Unmarshal(body, &testResponse)
		require.Equal(t, err, nil)
		// test code
		require.Equal(t, test.ExpectedStatusCode, res.Code)
		// test response body
		require.Equal(t, test.ExpectedResponse, testResponse)
	}
}

func TestGetUser(t *testing.T) {
	tests := []Test{
		{
			Method:             "GET",
			Path:               "/api/v1/users/imtiaz",
			Body:               nil,
			ExpectedStatusCode: 200,
			ExpectedResponse: Response{
				Status: "success",
			},
		},
	}
	// gets the response from http
	var testResponse Response

	for _, test := range tests {
		req, err := http.NewRequest(test.Method, test.Path, test.Body)
		req.Header.Add("Authorization", JWTToken)
		require.Equal(t, err, nil)

		res := httptest.NewRecorder()
		Router().ServeHTTP(res, req)

		body, _ := io.ReadAll(res.Body)
		err = json.Unmarshal(body, &testResponse)

		require.Equal(t, err, nil)
		// test code
		require.Equal(t, test.ExpectedStatusCode, res.Code)
		//// test response body
		//require.Equal(t, test.ExpectedResponse, testResponse)
	}
}

func TestUpdateUser(t *testing.T) {
	tests := []Test{
		{
			Method: "PUT",
			Path:   "/api/v1/users/test_akkas",
			Body: bytes.NewReader([]byte(`{
   				"username" : "test_akkas_updated",
   				"password" : "1234",
				"organization" : "Appscode Ltd",
   				"email": "xyz@gmail.com"}`)),
			ExpectedStatusCode: 200,
			ExpectedResponse: Response{
				Status:  "success",
				Message: "updated successfully",
			},
		},
		{
			Method: "PUT",
			Path:   "/api/v1/users/test_akkas_updated",
			Body: bytes.NewReader([]byte(`{
				"username" : "test_akkas",
				"password" : "1234",
				"organization" : "Appscode Ltd",
				"email": "xyz@gmail.com"}`)),
			ExpectedStatusCode: 200,
			ExpectedResponse: Response{
				Status:  "success",
				Message: "updated successfully",
			},
		},
		{
			Method: "PUT",
			Path:   "/api/v1/users/test_akkas_updated",
			Body: bytes.NewReader([]byte(`{
				"username" : "test_akkas",
				"password" : "1234",
				"organization" : "Appscode Ltd",
				"email": "xyz@gmail.com"}`)),
			ExpectedStatusCode: 406,
			ExpectedResponse: Response{
				Status:  "failed",
				Message: "username not found",
			},
		},
	}
	// gets the response from http
	var testResponse Response

	for _, test := range tests {
		req, err := http.NewRequest(test.Method, test.Path, test.Body)
		req.Header.Add("Authorization", JWTToken)
		require.Equal(t, err, nil)

		res := httptest.NewRecorder()
		Router().ServeHTTP(res, req)

		body, _ := io.ReadAll(res.Body)
		err = json.Unmarshal(body, &testResponse)

		require.Equal(t, err, nil)
		// test code
		require.Equal(t, test.ExpectedStatusCode, res.Code)
		// test response body
		require.Equal(t, test.ExpectedResponse, testResponse)
	}
}

func TestDeleteUser(t *testing.T) {
	tests := []Test{
		{
			Method:             "DELETE",
			Path:               "/api/v1/users/test_akkas",
			Body:               nil,
			ExpectedStatusCode: 202,
			ExpectedResponse: Response{
				Status: "success",
			},
		},
		{
			Method:             "DELETE",
			Path:               "/api/v1/users/test_akkas",
			Body:               nil,
			ExpectedStatusCode: 404,
			ExpectedResponse: Response{
				Status:  "failed",
				Message: "username not found",
			},
		},
	}
	// gets the response from http
	var testResponse Response

	for _, test := range tests {
		req, err := http.NewRequest(test.Method, test.Path, test.Body)
		req.Header.Add("Authorization", JWTToken)
		require.Equal(t, err, nil)

		res := httptest.NewRecorder()
		Router().ServeHTTP(res, req)

		body, _ := io.ReadAll(res.Body)
		err = json.Unmarshal(body, &testResponse)

		require.Equal(t, err, nil)
		// test code
		require.Equal(t, test.ExpectedStatusCode, res.Code)
		//// test response body
		//require.Equal(t, test.ExpectedResponse, testResponse)
	}
}

func TestCreateBook(t *testing.T) {
	tests := []Test{
		{
			Method: "POST",
			Path:   "/api/v1/books",
			Body: bytes.NewReader([]byte(`{
				"book-name": "Kire bhai ki obosta?",
				"price": 200,
				"isbn" : "4323-6456-4756-4564",
				"authors" : [
					{
						"username": "imtiaz"
					}
				],
				"book-content": {
					"over-view": "Haire over-view",
					"chapters" : [
						{
							"chapter-title": "chapter 1",
							"chapter-content": "chapter 1 content"
						},
						{
							"chapter-title": "chapter 2",
							"chapter-content": "chapter 2 content"
						}
					]
				}
			}`)),
			ExpectedStatusCode: 200,
			ExpectedResponse: Response{
				Status:  "failed",
				Message: "authentication required",
			},
		},
		{
			Method: "POST",
			Path:   "/api/v1/books",
			Body: bytes.NewReader([]byte(`{
				"book-name": "Kire bhai ki obosta?",
				"price": 200,
				"isbn" : "4323-6456-4756-4564",
				"authors" : [
					{
						"username": "imtiaz"
					}
				],
				"book-content": {
					"over-view": "Haire over-view",
					"chapters" : [
						{
							"chapter-title": "chapter 1",
							"chapter-content": "chapter 1 content"
						},
						{
							"chapter-title": "chapter 2",
							"chapter-content": "chapter 2 content"
						}
					]
				}
			}`)),
			ExpectedStatusCode: 200,
			ExpectedResponse: Response{
				Status: "success",
				Message: map[string]any{
					"id":          100,
					"price":       200,
					"sold-copies": 0,
					"isbn":        "4323-6456-4756-4564",
					"book-name":   "Kire bhai ki obosta?",
					"authors": []any{
						map[string]any{
							"id":           float64(100),
							"role":         "admin",
							"name":         "",
							"email":        "imtiazuddincho246@gmail.com",
							"username":     "imtiaz",
							"organization": "Appscode Ltd",
						},
						map[string]any{
							"book-content": map[string]any{
								"over-view": "Haire over-view",
								"chapters": []any{
									map[string]any{
										"chapter-title":   "chapter 1",
										"chapter-content": "chapter 1 content",
									},
									map[string]any{
										"chapter-title":   "chapter 2",
										"chapter-content": "chapter 2 content",
									},
								},
							},
						},
					},
				},
			},
		},
	}

	// gets the response from http
	var testResponse Response

	for _, test := range tests {
		req, err := http.NewRequest(test.Method, test.Path, test.Body)
		req.Header.Add("Authorization", JWTToken)
		require.Equal(t, err, nil)

		res := httptest.NewRecorder()
		Router().ServeHTTP(res, req)

		body, _ := io.ReadAll(res.Body)
		err = json.Unmarshal(body, &testResponse)

		require.Equal(t, err, nil)
		// test code
		require.Equal(t, test.ExpectedStatusCode, res.Code)
		//// test response body
		//require.Equal(t, test.ExpectedResponse, testResponse)

		// unauthorized access. Get the JWT token
		if res.Code == 401 {
			TestGetToken(t)
		}
	}
}

func TestGetBookList(t *testing.T) {
	tests := []Test{
		{
			Method:             "GET",
			Path:               "/api/v1/books",
			Body:               nil,
			ExpectedStatusCode: 200,
			ExpectedResponse: Response{
				Status: "success",
			},
		},
	}
	// gets the response from http
	var testResponse Response

	for _, test := range tests {
		req, err := http.NewRequest(test.Method, test.Path, test.Body)
		req.Header.Add("Authorization", JWTToken)
		require.Equal(t, err, nil)

		res := httptest.NewRecorder()
		Router().ServeHTTP(res, req)

		body, _ := io.ReadAll(res.Body)
		err = json.Unmarshal(body, &testResponse)

		require.Equal(t, err, nil)
		// test code
		require.Equal(t, test.ExpectedStatusCode, res.Code)
		//// test response body
		//require.Equal(t, test.ExpectedResponse, testResponse)
	}
}

func TestGetBook(t *testing.T) {
	tests := []Test{
		{
			Method:             "GET",
			Path:               "/api/v1/books/100",
			Body:               nil,
			ExpectedStatusCode: 406,
			ExpectedResponse: Response{
				Status: "success",
			},
		},
	}
	// gets the response from http
	var testResponse Response

	for _, test := range tests {
		req, err := http.NewRequest(test.Method, test.Path, test.Body)
		req.Header.Add("Authorization", JWTToken)
		require.Equal(t, err, nil)

		res := httptest.NewRecorder()
		Router().ServeHTTP(res, req)

		body, _ := io.ReadAll(res.Body)
		err = json.Unmarshal(body, &testResponse)

		require.Equal(t, err, nil)
		// test code
		require.Equal(t, test.ExpectedStatusCode, res.Code)
		//// test response body
		//require.Equal(t, test.ExpectedResponse, testResponse)
	}
}

func TestUpdateBook(t *testing.T) {
	tests := []Test{
		{
			Method: "PUT",
			Path:   "/api/v1/books/101",
			Body: bytes.NewReader([]byte(`{
				"book-name": "update",
				"price": 200,
				"isbn" : "4323-6456-4756-4564",
				"authors" : [
					{
						"username": "imtiaz"
					}
				],
				"book-content": {
					"over-view": "overview",
					"chapters" : [
						{
							"chapter-title": "chapter 1",
							"chapter-content": "chapter 1 content"
						},
						{
							"chapter-title": "chapter 2",
							"chapter-content": "chapter 2 content"
						}
					]
				}
			}`)),
			ExpectedStatusCode: 202,
			ExpectedResponse: Response{
				Status:  "success",
				Message: "updated successfully",
			},
		},
	}
	// gets the response from http
	var testResponse Response

	for _, test := range tests {
		req, err := http.NewRequest(test.Method, test.Path, test.Body)
		req.Header.Add("Authorization", JWTToken)
		require.Equal(t, err, nil)

		res := httptest.NewRecorder()
		Router().ServeHTTP(res, req)

		body, _ := io.ReadAll(res.Body)
		err = json.Unmarshal(body, &testResponse)

		require.Equal(t, err, nil)
		// test code
		require.Equal(t, test.ExpectedStatusCode, res.Code)
		// test response body
		require.Equal(t, test.ExpectedResponse, testResponse)
	}
}

func TestDeleteBook(t *testing.T) {
	tests := []Test{
		{
			Method:             "DELETE",
			Path:               "/api/v1/books/101",
			Body:               nil,
			ExpectedStatusCode: 202,
			ExpectedResponse: Response{
				Status:  "success",
				Message: "deleted successfully",
			},
		},
		{
			Method:             "DELETE",
			Path:               "/api/v1/books/101",
			Body:               nil,
			ExpectedStatusCode: 406,
			ExpectedResponse: Response{
				Status:  "failed",
				Message: "book doesn't exists",
			},
		},
	}
	// gets the response from http
	var testResponse Response

	for _, test := range tests {
		req, err := http.NewRequest(test.Method, test.Path, test.Body)
		req.Header.Add("Authorization", JWTToken)
		require.Equal(t, err, nil)

		res := httptest.NewRecorder()
		Router().ServeHTTP(res, req)

		body, _ := io.ReadAll(res.Body)
		err = json.Unmarshal(body, &testResponse)

		require.Equal(t, err, nil)
		// test code
		require.Equal(t, test.ExpectedStatusCode, res.Code)
		//// test response body
		//require.Equal(t, test.ExpectedResponse, testResponse)
	}
}

func TestBooksOfUser(t *testing.T) {
	tests := []Test{
		{
			Method:             "GET",
			Path:               "/api/v1/users/imtiaz/books",
			Body:               nil,
			ExpectedStatusCode: 200,
			ExpectedResponse: Response{
				Status: "success",
			},
		},
	}
	// gets the response from http
	var testResponse Response

	for _, test := range tests {
		req, err := http.NewRequest(test.Method, test.Path, test.Body)
		req.Header.Add("Authorization", JWTToken)
		require.Equal(t, err, nil)

		res := httptest.NewRecorder()
		Router().ServeHTTP(res, req)

		body, _ := io.ReadAll(res.Body)
		err = json.Unmarshal(body, &testResponse)

		require.Equal(t, err, nil)
		// test code
		require.Equal(t, test.ExpectedStatusCode, res.Code)
		//// test response body
		//require.Equal(t, test.ExpectedResponse, testResponse)
	}
}
