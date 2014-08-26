//Go implementation of Soundcloud Oauth2 for Server-Side Web Applications
//Reference and Doc: https://developers.soundcloud.com/docs/api/guide#authentication
package gosoundcloudoauth

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

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
	//AvatarData           []byte `json:"AvatarData"`				//binary data of user avatar. Example: (only for uploading)
}

type Oauth2Token struct {
	AccessToken string `json:"access_token"`
	Scope       string `json:"scope"`
}

type OauthClient struct {
	ID             string //The client id belonging to your application.
	Secret         string //The client secret belonging to your application.
	RedirectUri    string //The redirect uri you have configured for your application.
	ResponseType   string //(code, token_and_code). By default this package uses 'code'
	Scope          string //Must be '*' or 'non-expiring'. By default this package uses 'non-expiring'.
	//Display        string //Can specify a value of 'popup' for mobile optimized screen.
	//State          string //Any value included here will be appended to the redirect URI
	GrantType      string //Enumeration (authorization_code, refresh_token, password, client_credentials, oauth1_token).By default this package uses 'authorization_code'.
	code           string //The authorization code obtained when user is sent to redirect_uri.
	AccessToken    string //Access Tokens are provided once a user has authorized your application.
	connectUrl     string //The OAuth2 authorization endpoint. Your app redirects a user to this endpoint, allowing them to delegate access to their account.
	oauth2TokenUrl string //The OAuth2 token endpoint. This endpoint accepts POST requests and is used to provision access tokens once a user has authorized your application.
	FailureUrl     string //In case of error a Redirect will be made to this URL.
	SuccessUrl     string //In case of success a Redirect will be made to this URL.
	authorizeUrl   string //Authorize Endpoint /connect with the parameters.
}

//Creates a client object to be used for authorization, token exchange and get the current user.
func SoundcloudOauth(clientId, clientSecret, redirectUri, responseType, scope, grantType, failure, success string) (*OauthClient, error) {
	if clientId == "" || clientSecret == "" || redirectUri == "" {
		return nil, throw("SoundcloudOauth():Client ID, Client Secret and Redirect URI are required fields. If you opt to use the default values for ResponseType, Scope, GrantType, just pass blank.")
	}

	client := new(OauthClient)
	if clientId != ""{
		client.ID = clientId
	}
	if clientSecret != ""{
		client.Secret = clientSecret
	}
	if redirectUri != ""{
		client.RedirectUri = redirectUri
	}
	if responseType != ""{
		client.ResponseType = responseType
	}else{
		client.ResponseType = "code"
	}
	if scope != ""{
		client.Scope = scope
	}else{
		client.Scope = "non-expiring"
	}
	if grantType != ""{
		client.GrantType = grantType
	}else{
		client.GrantType = "authorization_code"
	}
	if failure != ""{
		client.FailureUrl = failure
	}else{
		client.FailureUrl = "/"
	}
	if success != ""{
		client.SuccessUrl = success
	}else{
		client.SuccessUrl = "/profile"
	}
	client.connectUrl = "https://soundcloud.com/connect"
	client.oauth2TokenUrl = "https://api.soundcloud.com/oauth2/token"
	client.authorizeUrl = client.connectUrl + "?client_id=" + client.ID + "&scope=" + client.Scope + "&response_type=" + client.ResponseType + "&redirect_uri=" + client.RedirectUri

	return client, nil
}

//This method redirects a user to /connect endpoint, allowing them to delegate access to their account.
//https://developers.soundcloud.com/docs/api/reference#connect
func (client *OauthClient) AuthorizeUrl(res http.ResponseWriter, req *http.Request) {
	http.Redirect(res, req, client.authorizeUrl, http.StatusMovedPermanently)
}

//This method sends a POST request to /oauth2/token endpoint and is used to provision access tokens once a user has authorized your application.
//https://developers.soundcloud.com/docs/api/reference#token
func (client *OauthClient) ExchangeToken(res http.ResponseWriter, req *http.Request) {
	code := req.FormValue("code")
	if code == "" {
		http.Redirect(res, req, client.FailureUrl, http.StatusMovedPermanently )
		return
	}
	client.code = code

	response, err := http.PostForm(client.oauth2TokenUrl,
		url.Values{"client_id": {client.ID}, "client_secret": {client.Secret}, "grant_type": {client.GrantType}, "redirect_uri": {client.RedirectUri}, "code": {client.code}})
	if err != nil {
		http.Redirect(res, req, client.FailureUrl, http.StatusMovedPermanently )
		return
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		http.Redirect(res, req, client.FailureUrl, http.StatusMovedPermanently )
		return
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		http.Redirect(res, req, client.FailureUrl, http.StatusMovedPermanently )
		return
	}

	token, err := parseOauth2Token(body)
	if err != nil {
		http.Redirect(res, req, client.FailureUrl, http.StatusMovedPermanently )
		return
	}

	client.AccessToken = token.AccessToken
	client.Scope = token.Scope

	http.Redirect(res, req, client.SuccessUrl, http.StatusFound )
}

//This method returns the current logged-in user after receiving the token
func (client *OauthClient) GetCurrentUser() (user *User, oauthToken string, err error) {
	if client.AccessToken == "" {
		return nil, "", throw("GetCurrentUser():Error: Please provide a valid Access Token.")
	}

	url := "https://api.soundcloud.com/me.json?oauth_token=" + client.AccessToken
	result, err := doRequest(url)
	if err != nil {
		return nil, "", throw("%s", err)
	}

	user, err = parseUser(result)
	if err != nil {
		return nil, "", throw("%s", err)
	}
	return user, client.AccessToken, err
}

func parseOauth2Token(body []byte) (*Oauth2Token, error) {
	token := new(Oauth2Token)
	err := json.Unmarshal(body, &token)
	if err != nil {
		return nil, throw("parseOauth2Token():json.Unmarshal:Oauth2Token: %s", err)
	}
	return token, nil
}

func parseUser(body []byte) (*User, error) {
	user := new(User)
	err := json.Unmarshal(body, &user)
	if err != nil {
		return nil, throw("parseUser():json.Unmarshal:User: %s", err)
	}
	return user, err
}

func doRequest(url string) ([]byte, error) {

	httpClient := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, throw("http.NewRequest('GET'):Error: error creating request object.")
	}
	res, err := httpClient.Do(req)
	if err != nil {
		return nil, throw("httpClient.Do(req):Error: error getting response object.")
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return nil, throw("doRequest(url):Error:res.StatusCode != 200: %s, URL: %s", res.Status, url)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, throw("doRequest(url):ioutil.ReadAllError:(res.Body): %s", err)
	}

	return []byte(body), err
}

func throw(message string, params ...interface{}) error {
	return fmt.Errorf(message, params...)
}


