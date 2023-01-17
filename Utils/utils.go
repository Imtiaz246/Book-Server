package Utils

import (
	"encoding/json"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"io"
	"os"
	"time"
)

// RestoreDataFromBackupFiles restores the backed up data
// and store those data to the central database.
func RestoreDataFromBackupFiles() ([]byte, []byte) {
	// Restore User data
	uf, err := os.Open("./BackupFiles/Users.json")
	defer uf.Close()
	usersJsonData, err := io.ReadAll(uf)

	// Restore Book data
	bf, err := os.Open("./BackupFiles/Books.json")
	defer bf.Close()
	booksJsonData, err := io.ReadAll(bf)

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	if len(usersJsonData) == 0 {
		usersJsonData = []byte("[]")
	}
	if len(booksJsonData) == 0 {
		booksJsonData = []byte("[]")
	}
	return usersJsonData, booksJsonData
}

// StoreDataToBackupFiles gets data from the central database
// and store those data to the backup files.
func StoreDataToBackupFiles(userJsonData, booksJsonData []byte) error {
	uf, err := os.Open("./BackupFiles/Users.json")
	defer uf.Close()
	// Store usersJsonData to the backup file
	_, err = uf.Write(userJsonData)
	if err != nil {
		return err
	}
	bf, err := os.Open("./BackupFiles/Books.json")
	defer bf.Close()
	// Store booksJsonData to the backup file
	_, err = bf.Write(booksJsonData)
	return err
}

// CreateErrorJson creates json error object.
func CreateErrorJson(err error) []byte {
	type ErrorMsg struct {
		Status  string `json:"status"`
		Message string `json:"message"`
	}
	em := ErrorMsg{
		Status:  "failed",
		Message: err.Error(),
	}
	eJson, err := json.Marshal(em)
	if err != nil {
		return []byte(err.Error())
	}
	return eJson
}

// CreateSuccessJson creates json success message with payload data.
func CreateSuccessJson(msg any) ([]byte, error) {
	type SuccessMsg struct {
		Status  string `json:"status"`
		Message any    `json:"message"`
	}
	sm := SuccessMsg{
		Status:  "success",
		Message: msg,
	}
	sJson, err := json.Marshal(sm)

	return sJson, err
}

// GenerateJwtToken creates a jwt token for a user with the
// username and password. Returns the (token, error) tuple.
func GenerateJwtToken(username string) (string, error) {
	SecretKey := []byte("mysecret")
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(10 * time.Minute)
	claims["authorized"] = true
	claims["user"] = username

	tokenStr, err := token.SignedString(SecretKey)
	if err != nil {
		return "", err
	}
	return tokenStr, nil
}
