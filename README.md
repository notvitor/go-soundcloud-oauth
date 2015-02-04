#go-soundcloud-oauth


Go implementation of Soundcloud Oauth2 for Server-Side Web Applications.

Reference and Doc: [https://developers.soundcloud.com/docs/api/guide#authentication](https://developers.soundcloud.com/docs/api/guide#authentication).

-----------------------------------------------------------------------------------------------


##Example

`go get github.com/vitorsvvv/go-soundcloud-oauth`

```

    package main

    import (
        "fmt"
        "github.com/vitorsvvv/gosoundcloudoauth"
        "log"
        "net/http"
    )

    func main() {
        id := "{SOUNDCLOUD-APP-CLIENT-ID}"
        secret := "{SOUNDCLOUD-APP-CLIENT-SECRET}"
        redirectUri := "{SOUNDCLOUD-APP-CLIENT-ID}"
        failureUrl := "/"
        successUrl := "/profile"

        oauth, err := gosoundcloudoauth.SoundcloudOauth(id, secret, redirectUri, "", "", "", failureUrl, successUrl)
        if err != nil {
            log.Fatal(err)
        }

        http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
            fmt.Fprintln(res, "<h1>Home</h1><br/><ul><li><a href='/login'>Login</a></li><li><a href='/profile'>Profile</a></li></ul>")
        })

        http.HandleFunc("/login", func(res http.ResponseWriter, req *http.Request) {
            oauth.AuthorizeUrl(res, req)
        })

        http.HandleFunc("/login/callback", func(res http.ResponseWriter, req *http.Request) {
            oauth.ExchangeToken(res, req)
        })

        http.HandleFunc("/profile", func(res http.ResponseWriter, req *http.Request) {
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
```

##Structs

###OauthClient
Reference:
[https://developers.soundcloud.com/docs/api/reference#connect](https://developers.soundcloud.com/docs/api/reference#connect)
[https://developers.soundcloud.com/docs/api/reference#token](https://developers.soundcloud.com/docs/api/reference#token)
```
type OauthClient struct {
	ID             string //The client id belonging to your application.
	Secret         string //The client secret belonging to your application.
	RedirectUri    string //The redirect uri you have configured for your application.
	ResponseType   string //(code, token_and_code). By default this package uses 'code'
	Scope          string //Must be '*' or 'non-expiring'. By default this package uses 'non-expiring'.
	GrantType      string //Enumeration (authorization_code, refresh_token, password, client_credentials, oauth1_token).By default this package uses 'authorization_code'.
	code           string //The authorization code obtained when user is sent to redirect_uri.
	AccessToken    string //Access Tokens are provided once a user has authorized your application.
	connectUrl     string //The OAuth2 authorization endpoint. Your app redirects a user to this endpoint, allowing them to delegate access to their account.
	oauth2TokenUrl string //The OAuth2 token endpoint. This endpoint accepts POST requests and is used to provision access tokens once a user has authorized your application.
	FailureUrl     string //In case of error a Redirect will be made to this URL.
	SuccessUrl     string //In case of success a Redirect will be made to this URL.
	authorizeUrl   string //Authorize Endpoint /connect with the parameters.
}
```
###User
Reference: [https://developers.soundcloud.com/docs/api/reference#me](https://developers.soundcloud.com/docs/api/reference#me)
```
type User struct {
	Id                    int    `json:"id"`                      //integer ID. Example: 123
	Permalink             string `json:"permalink"`               //permalink of the resource. Example: "sbahn-sounds"
	Username              string `json:"username"`                //username. Example: "Doctor Wilson"
	Uri                   string `json:"uri"`                     //API resource URL.Example: http://api.soundcloud.com/comments/32562
	PermalinkUrl          string `json:"permalink_url"`           //URL to the SoundCloud.com page. Example: "http://soundcloud.com/bryan/sbahn-sounds"
	AvatarUrl             string `json:"avatar_url"`              //URL to a JPEG image.	Example: "http://i1.sndcdn.com/avatars-000011353294-n0axp1-large.jpg"
	Country               string `json:"country"`                 //country.	Example: "Germany"
	FullName              string `json:"full_name"`               //first and last name.	Example: "Tom Wilson"
	City                  string `json:"city"`                    //city. Example: "Berlin"
	Description           string `json:"description"`             //description.	Example: "Buskers playing in the S-Bahn station in Berlin"
	DiscogsName           string `json:"discogs_name"`            //Discogs name. Example: "myrandomband"
	MyspaceName           string `json:"myspace_name"`            //MySpace name. Example: "myrandomband"
	Website               string `json:"website"`                 //a URL to the website. Example: "http://facebook.com/myrandomband"
	WebsiteTitle          string `json:"website_title"`           //a custom title for the website. Example: "myrandomband on Facebook"
	Online                bool   `json:"online"`                  //online status (boolean).	Example: true
	TrackCount            int    `json:"track_count"`             //number of public tracks.	Example: 4
	PlaylistCount         int    `json:"playlist_count"`          //number of public playlists. Example: 5
	FollowersCount        int    `json:"followers_count"`         //number of followers.	Example: 54
	FollowingsCount       int    `json:"followings_count"`        //number of followed users. Example: 75
	PublicFavoritesCount  int    `json:"public_favorites_count"`  //number of favorited public tracks. Example:	7
	Plan                  string `json:"plan"`                    //subscription plan of the user. Example: Pro Plus
	PrivateTracksCount    int    `json:"private_tracks_count"`    //number of private tracks. Example: 34
	PrivatePlaylistsCount int    `json:"private_playlists_count"` //number of private playlists. Example: 6
	PrimaryEmailConfirmed bool   `json:"primary_email_confirmed"` //boolean if email is confirmed. Example true
}
```




Check the full documented example [here.](https://github.com/vitorsvvv/go-soundcloud-oauth/blob/master/examples/server.go)





---


Have a request, suggestion or question?

Drop me an email: vitorsvieira at yahoo.com

You can also find me at [@notvitor](https://twitter.com/notvitor)


---


### License

The MIT License (MIT)

Copyright (c) 2014 Vitor Vieira

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
