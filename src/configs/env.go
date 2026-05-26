package configs

import (
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
)

var loadOnce sync.Once

// load reads .env once per process. A missing file is not fatal:
// in production, variables come from the environment, not a file.
func load() {
	loadOnce.Do(func() {
		if err := godotenv.Load(".env"); err != nil {
			log.Printf("no .env file loaded (%v); relying on process environment", err)
		}
	})
}

func GetEnvKeys(key string) string {
	load()
	return os.Getenv(key)
}
