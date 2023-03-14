// Package database implements a basic database module for the Book-server.
// It provides a global 'db' object of type DataBase
// which is the central data storage for the Book-server.
// The db object can be created from backup files using NewDB function,
// and the pointer instance of the db object can be got by GetDB function.
package database
