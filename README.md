#go-soundcloud-oauth


Go implementation of Soundcloud Oauth2 for Server-Side Web Applications.
Reference and Doc: [https://developers.soundcloud.com/docs/api/guide#authentication](https://developers.soundcloud.com/docs/api/guide#authentication).

-----------------------------------------------------------------------------------------------

##Example
```golang

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


##Documentation

###Method signature
```
//Creates a client object to be used for authorization, token exchange and get the current user.
func SoundcloudOauth(clientId, clientSecret, redirectUri, responseType, scope, grantType, failure, success string) (*OauthClient, error)
```

###Creating oauth object with the parameters
```

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
```

##Usage

Download `go get github.com/vitorsvvv/go-soundcloud-oauth`





---


Have a request, suggestion or question?

Drop me an email: vitorsvieira at yahoo.com

You can also find me at [@vitorsvvv](https://twitter.com/vitorsvvv)


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