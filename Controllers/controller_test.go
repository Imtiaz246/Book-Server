package Controllers

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"testing"
)

type Response struct {
	Status  string `json:"status"`
	Message any
}

var GlobalResponseUnMarshaller Response

func TestControllers(t *testing.T) {
	err := testBookControllers()
	if err != nil {
		t.Error(err.Error())
	}
}

func testBookControllers() error {
	// Testing {GET /api/v1/books} routes
	allBooks := struct {
		Method, Path, expStatus string
		expResponseStatus       string
	}{
		"GET",
		"http://localhost:3000/api/v1/books",
		"200 OK",
		"success",
	}
	resp, _ := http.Get(allBooks.Path)
	respBody, _ := io.ReadAll(resp.Body)
	if resp.Status != allBooks.expStatus {
		return errors.New("expected 200 found " + resp.Status)
	}
	err := json.Unmarshal(respBody, &GlobalResponseUnMarshaller)
	if err != nil {
		return err
	}
	if GlobalResponseUnMarshaller.Status != allBooks.expResponseStatus {
		return errors.New("response status doesn't match")
	}
	//---------------------------------------------------------------
	// Testing {POST /api/v1/books} get the id & {GET /api/v1/books/{id}} routes

	return nil
}
