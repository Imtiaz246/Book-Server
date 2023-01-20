package main

import (
	"BookServer/Database"
	"BookServer/Router"
	"net/http"
	"os"
	"os/signal"
)

func main() {
	// Initialize the central DataBase instance with the backed up data.
	// Previously backed up data will be restored from BackupFiles folder if any exits.
	db := Database.NewDB()

	// Catch the os signal
	osSignalChan := make(chan os.Signal)
	signal.Notify(osSignalChan)

	// Start the server with Router Handler and Listen on port :3000
	go http.ListenAndServe(":3000", Router.Router())

	<-osSignalChan
	db.DbBackup()
}
