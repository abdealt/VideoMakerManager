package main

import (
	"log"
	"videomaker/app"
	"videomaker/database"
)

func main() {
	// Initialisation de la base de données
	if err := database.Init(); err != nil {
		log.Fatalf("Erreur lors de l'initialisation de la base de données: %v", err)
	}

	app.Run()
}
