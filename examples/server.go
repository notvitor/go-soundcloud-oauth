package main

import (
	"fmt"
	"github.com/vitorsvvv/gosoundcloudoauth"
	"log"
	"net/http"
)

func main() {

	// OFFICIAL SOUNDCLOUD AUTHORIZATION DOCUMENTATION
	// https://developers.soundcloud.com/docs/api/guide#authentication

	// If you don't have an account and an app registered
	// Go https://soundcloud.com/you/apps/
	id := "{SOUNDCLOUD-APP-CLIENT-ID}"
	secret := "{SOUNDCLOUD-APP-CLIENT-SECRET}"
	redirectUri := "{SOUNDCLOUD-APP-CLIENT-ID}"

	// If any error occur during the authorization process the user will be redirected to the home page
	failureUrl := "/"
	// If everything goes ok during the authorization process the user will be redirected to the profile page
	successUrl := "/profile"

	oauth, err := gosoundcloudoauth.SoundcloudOauth(id, secret, redirectUri, "", "", "", failureUrl, successUrl)
	if err != nil {
		log.Fatal(err)
	}

	// Home Page Route
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(res, "<h1>Home</h1><br/><ul><li><a href='/login'>Login</a></li><li><a href='/profile'>Profile</a></li></ul>")
	})

	// Login Route
	http.HandleFunc("/login", func(res http.ResponseWriter, req *http.Request) {
		oauth.AuthorizeUrl(res, req)
	})

	// Route registered as a Callback/Redirect URI in your app profile.
	// Note that this same route must be provided as a parameter when instantiating SoundcloudOauth.
	http.HandleFunc("/login/callback", func(res http.ResponseWriter, req *http.Request) {
		oauth.ExchangeToken(res, req)
	})

	// Profile Route
	http.HandleFunc("/profile", func(res http.ResponseWriter, req *http.Request) {

		// You should now store the access token in a database.
		// Associate it with the user it belongs to and use it from now on
		// instead of sending the user through the authorization flow.
		user, token, err := oauth.GetCurrentUser()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Fprintf(res, "<h1>Welcome!</h1><br/><a href='/'>Home</a><br/><h3>Token: %s</h3><br/><h3>ID: %d</h3><br/><h3>UserName: %s</h3><br/><h3>City: %s</h3><br/><h3>Country: %s</h3><br/><h3>Description: %s</h3><br/><br/>",
			token, user.Id, user.Username, user.City, user.Country, user.Description)
	})

	err = http.ListenAndServe(":3000", nil)
	if err != nil {
		fmt.Println(err)
	}
}
