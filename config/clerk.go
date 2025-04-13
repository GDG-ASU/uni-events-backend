package config

import (
	"log"
	"os"
)

var ClerkSecret string

func InitClerk() {
	ClerkSecret = os.Getenv("CLERK_SECRET")
	if ClerkSecret == "" {
		log.Fatal("Missing CLERK_SECRET in environment")
	}
	log.Println("Clerk initialized")
}
