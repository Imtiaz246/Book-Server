package Utils

import (
	"fmt"
	"io"
	"os"
)

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

func StoreDataToBackupFiles() {

}
