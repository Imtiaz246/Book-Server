package cmd

import (
	"fmt"
	"github.com/Imtiaz246/Book-Server/database"
	"github.com/Imtiaz246/Book-Server/router"
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
)

var (
	port                     int
	adminName, adminPassword string

	rootCmd = &cobra.Command{
		Use:   "starts the server",
		Short: "starts the server",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			startServer()
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
	rootCmd.PersistentFlags().IntVarP(&port, "port", "p", 0, "port no for the server to run")
}

func startServer() {
	// Load .env files
	err := godotenv.Load(".env")
	if err != nil {
		log.Println(err.Error())
	}

	// Initialize the central DataBase instance with the backed up data.
	// Previously backed up data will be restored from BackupFiles folder if any exits.
	db, err := database.NewDB()
	if err != nil {
		log.Println(err.Error())
	}
	err = db.AddAdmin(os.Getenv("ADMIN_USERNAME"), os.Getenv("ADMIN_PASSWORD"))
	if err != nil {
		log.Println(err.Error())
	}

	// Catch the os signal
	osSignalChan := make(chan os.Signal)
	signal.Notify(osSignalChan, os.Interrupt, os.Kill)

	// Start the server with router Handler and Listen on port :3000 by default if -p flag is not given
	log.Println("Listening on port : ", port)
	go func() {
		err := http.ListenAndServe(":"+strconv.Itoa(port), router.Router())
		if err != nil {
			log.Println(err.Error())
			os.Exit(1)
		}
	}()

	<-osSignalChan
	err = db.DbBackup()
	if err != nil {
		log.Println(err.Error())
	} else {
		log.Println("DB backup taken successfully..")
	}
}
