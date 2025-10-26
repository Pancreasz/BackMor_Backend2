package config

import (
	"os"

	"github.com/markbates/goth"
	"github.com/markbates/goth/providers/facebook"
	"github.com/markbates/goth/providers/google"

	"fmt"
)

func Goth_init() {

	fmt.Println(os.Getenv("REDIRECT_URL_GOOGLE"), "alalalalalalalalal")
	fmt.Println(os.Getenv("REDIRECT_URL_FACEBOOK"), "alalalalalalalalal")

	goth.UseProviders(
		google.New(
			os.Getenv("GOOGLE_CLIENT_ID"),
			os.Getenv("GOOGLE_CLIENT_SECRET"),
			os.Getenv("REDIRECT_URL_GOOGLE"),
			"email", "profile", // add scopes
		),
		facebook.New(
			os.Getenv("facebook_ID"),
			os.Getenv("facebook_secret"),
			os.Getenv("REDIRECT_URL_FACEBOOK"),
			"public_profile", "email",
		),
	)
}
