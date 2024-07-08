package env

import (
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

var (
	DISCORD_TOKEN *string
	PORT          *string
)

func LoadEnv() {
	err := godotenv.Load()

	if err != nil {
		log.Println("Error loading .env file")
	}

	setPort()
	setDiscordKey()
}

func setPort() {
	p := os.Getenv("PORT")
	finalVal := ""

	if !strings.HasPrefix(p, ":") {
		finalVal = ":"
	}

	if len(p) >= 1 {
		finalVal += p
	} else {
		finalVal += "3000"
	}

	PORT = &finalVal
}

func setDiscordKey() {
	key := os.Getenv("DISCORD_TOKEN")
	var finalKey string

	if strings.HasPrefix(key, "Bot") {
		finalKey = key
	} else {
		finalKey = "Bot " + key
	}

	DISCORD_TOKEN = &finalKey
}
