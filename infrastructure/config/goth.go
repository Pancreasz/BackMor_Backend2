package config

import (
	"os"

	"github.com/markbates/goth"
	"github.com/markbates/goth/providers/facebook"
	"github.com/markbates/goth/providers/google"
)

func Goth_init() {

	goth.UseProviders(
		google.New(
			os.Getenv("GOOGLE_CLIENT_ID"),
			os.Getenv("GOOGLE_CLIENT_SECRET"),
			"http://localhost:8000/v1/api/auth/google/callback",
			"email", "profile", // add scopes
		),
		facebook.New(
			os.Getenv("facebook_ID"),
			os.Getenv("facebook_secret"),
			"http://localhost:8000/v1/api/auth/facebook/callback",
			"public_profile", "email",
		),
	)
}
