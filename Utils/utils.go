package Utils

import (
	"encoding/json"
	"errors"
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
		usersJsonData = []byte("{}")
	}
	if len(booksJsonData) == 0 {
		booksJsonData = []byte("{}")
	}
	return usersJsonData, booksJsonData
}

// StoreDataToBackupFiles gets data from the central database
// and store those data to the backup files.
func StoreDataToBackupFiles(userJsonData, booksJsonData []byte) error {
	var err error
	err = os.WriteFile("./BackupFiles/Users.json", userJsonData, 0644)
	if err != nil {
		return err
	}
	err = os.WriteFile("./BackupFiles/Books.json", booksJsonData, 0644)
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
	// Get the secret key from the environment variable
	jwtSecretKey := []byte(os.Getenv("jwt-secret"))
	// JWT claims struct
	type Claims struct {
		Username string `json:"username"`
		jwt.RegisteredClaims
	}
	expTime := time.Now().Add(time.Hour * 12)
	claims := &Claims{
		username,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString(jwtSecretKey)

	return tokenStr, err
}

// CheckJWTValidation checks if a jwt token is valid or not.
// Returns (username, error) tuple.
func CheckJWTValidation(tokenStr string) (string, error) {
	// Get the secret key from the environment variable
	jwtSecretKey := []byte(os.Getenv("jwt-secret"))
	// JWT claims struct
	type Claims struct {
		Username string `json:"username"`
		jwt.RegisteredClaims
	}
	claims := &Claims{}
	rToken, err := jwt.ParseWithClaims(tokenStr, claims,
		func(token *jwt.Token) (interface{}, error) {
			return jwtSecretKey, nil
		})
	if err != nil {
		return "", err
	}
	if !rToken.Valid {
		return "", errors.New("token is not valid")
	}
	return claims.Username, nil
}
