// Define configuration specifics (like database) here:

package config

import (
	"os"
)

func GetDatabaseURL() string {
	return os.Getenv("DB_URL")
}