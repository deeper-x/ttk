package settings

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

// GetDBURI return database uri string
func GetDBURI() (string, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return "", err
	}

	user := os.Getenv("DB_USER")
	name := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")
	passwd := os.Getenv("DB_PASSWD")
	host := os.Getenv("DB_HOST")

	URI := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", user, passwd, host, port, name)

	return URI, nil
}
