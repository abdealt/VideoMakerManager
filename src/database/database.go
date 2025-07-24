package database

import (
	"fmt"
	"os"
	"sync"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db    *gorm.DB
	once  sync.Once
	dbErr error
)

// Init initialise la connexion à la base de données
func Init() error {
	// Charger les variables d'environnement
	if err := godotenv.Load(); err != nil {
		return fmt.Errorf("erreur lors du chargement du fichier .env: %v", err)
	}

	// Utiliser DB_USERNAME au lieu de USERNAME pour éviter les conflits
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	database := os.Getenv("DB_DATABASE")
	port := os.Getenv("DB_PORT")

	// Valeurs par défaut si les variables ne sont pas définies
	if username == "" {
		username = "root"
	}
	if host == "" {
		host = "localhost"
	}
	if port == "" {
		port = "3306"
	}
	if database == "" {
		database = "vmm"
	}

	// Construction du DSN
	var dsn string
	if password == "" {
		// Si pas de mot de passe, ne pas inclure le ':'
		dsn = fmt.Sprintf("%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			username, host, port, database)
	} else {
		dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			username, password, host, port, database)
	}

	fmt.Printf("Tentative de connexion à la base de données avec l'utilisateur: %s\n", username)

	once.Do(func() {
		var err error
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			dbErr = fmt.Errorf("erreur lors de la connexion avec Gorm : %v", err)
		}
	})

	return dbErr
}

// Get retourne l'instance de la base de données
func Get() (*gorm.DB, error) {
	if db == nil {
		return nil, fmt.Errorf("base de données non initialisée")
	}
	return db, nil
}

// Connect est une méthode temporaire pour maintenir la compatibilité
func Connect() (*gorm.DB, error) {
	if db == nil {
		if err := Init(); err != nil {
			return nil, err
		}
	}
	return db, nil
}
