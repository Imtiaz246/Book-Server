package main

import (
	"github.com/Imtiaz246/Book-Server/Database"
	"github.com/Imtiaz246/Book-Server/Router"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// Initialize the central DataBase instance with the backed up data.
	// Previously backed up data will be restored from BackupFiles folder if any exits.
	db, _ := Database.NewDB()

	// Catch the os signal
	osSignalChan := make(chan os.Signal)
	signal.Notify(osSignalChan, os.Interrupt)

	// Start the server with Router Handler and Listen on port :3000
	go http.ListenAndServe(":3000", Router.Router())

	for sigHandles(<-osSignalChan) {
		db.DbBackup()
		break
	}

}

func sigHandles(sig os.Signal) bool {
	return sig == syscall.SIGINT || sig == syscall.SIGKILL
}
