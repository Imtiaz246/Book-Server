package cmd

import (
	"fmt"
	"github.com/Imtiaz246/Book-Server/database"
	"github.com/Imtiaz246/Book-Server/router"
	"github.com/spf13/cobra"
	"log"
	"net/http"
	"os"
	"os/signal"
)

var (
	port                     int
	adminName, adminPassword string

	rootCmd = &cobra.Command{
		Use:   "starts the server",
		Short: "starts the server",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(`Please add admin credentials: 
			--admin-username <username> --admin-password <password>
			-p <port no> by default app will run on port 3000`)
		},
	}
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {

}

func startServer() {
	// Initialize the central DataBase instance with the backed up data.
	// Previously backed up data will be restored from BackupFiles folder if any exits.
	db, _ := database.NewDB()

	// Catch the os signal
	osSignalChan := make(chan os.Signal)
	signal.Notify(osSignalChan, os.Interrupt, os.Kill)

	// Start the server with router Handler and Listen on port :3000
	go http.ListenAndServe(":3000", router.Router())

	<-osSignalChan
	err := db.DbBackup()
	if err != nil {
		log.Println("DB backup has not been done..")
	} else {
		log.Println("DB backup taken successfully..")
	}
}
