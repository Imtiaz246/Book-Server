package Utils

import (
	"encoding/json"
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"io"
	"os"
	"time"
)

// RestoreDataFromBackupFiles restores the backed up data
// and store those data to the central database.
func RestoreDataFromBackupFiles() ([]byte, []byte, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return nil, nil, err
	}
	// Restore User data
	uf, err := os.Open(cwd + "/BackupFiles/Users.json")
	if err != nil {
		return nil, nil, err
	}
	defer uf.Close()
	usersJsonData, err := io.ReadAll(uf)
	if err != nil {
		return nil, nil, err
	}

	// Restore Book data
	bf, err := os.Open(cwd + "/BackupFiles/Books.json")
	if err != nil {
		return nil, nil, err
	}
	defer bf.Close()
	booksJsonData, err := io.ReadAll(bf)
	if err != nil {
		return nil, nil, err
	}
	if len(usersJsonData) == 0 {
		usersJsonData = []byte("{}")
	}
	if len(booksJsonData) == 0 {
		booksJsonData = []byte("{}")
	}
	return usersJsonData, booksJsonData, nil
}

// StoreDataToBackupFiles gets data from the central database
// and store those data to the backup files.
func StoreDataToBackupFiles(userJsonData, booksJsonData []byte) error {
	var err error
	cwd, err := os.Getwd()
	if err != nil {
		return err
	}
	err = os.WriteFile(cwd+"/BackupFiles/Users.json", userJsonData, 0644)
	if err != nil {
		return err
	}
	err = os.WriteFile(cwd+"/BackupFiles/Books.json", booksJsonData, 0644)
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
