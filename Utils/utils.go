package Utils

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

// RestoreDataFromBackupFiles restores the backed up data
// and store those data to the central database.
func RestoreDataFromBackupFiles() ([]byte, []byte) {
	// Restore User data
	file, err := os.Open("./BackupFiles/Users.json")
	usersJsonData, err := io.ReadAll(file)

	// Restore Book data
	file, err = os.Open("./BackupFiles/Books.json")
	booksJsonData, err := io.ReadAll(file)

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	return usersJsonData, booksJsonData
}

// StoreDataToBackupFiles gets data from the central database
// and store those data to the backup files.
func StoreDataToBackupFiles() {

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
