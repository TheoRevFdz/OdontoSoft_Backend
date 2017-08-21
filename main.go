package main

import "github.com/TheoRev/OdontoSoft_Backend/config"

func main() {
	config.Connect()
	config.CreateTables()
	config.CloseConnection()
}
