package main

import (
	"BookServer/Database"
	"net/http"
)

func main() {
	// Initialize the central DataBase instance with the backed up data.
	// Previously backed up data will be restored from BackupFiles folder if any exits.
	Database.NewDB()

	// Start the server with Router Handler and Listen on port :3000
	http.ListenAndServe(":3000", Router())

}
