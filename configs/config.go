package configs

import (
	"log"

	"github.com/joho/godotenv"
)

type Config struct {
	NOTION_SECRET string 
	NOTION_PAGEID string
	OPEN_AI string
}

var Env = Config{
	NOTION_SECRET: "NOTION_SECRET",
	NOTION_PAGEID: "NOTION_PAGEID",
	OPEN_AI: "OPEN_AI",
}


func StartConfig() error {
	err := godotenv.Load()
	if err != nil {
	  log.Fatal("Error loading .env file")
	}

	return nil
}