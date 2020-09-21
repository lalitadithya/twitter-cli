package twitter

import (
	"github.com/dghubble/oauth1"
	"github.com/lalitadithya/twitter-cli/twitter/util"
)

var TwitterBaseURL string = "https://api.twitter.com/"
var TwitterAPIVersion string = "1.1"

var TwitterAuthenticateEndpoint = oauth1.Endpoint{
	RequestTokenURL: "https://api.twitter.com/oauth/request_token",
	AuthorizeURL:    "https://api.twitter.com/oauth/authenticate",
	AccessTokenURL:  "https://api.twitter.com/oauth/access_token",
}

type TwitterClient struct {
	userSecrets   *util.UserSecrets
	config        oauth1.Config
	token         *oauth1.Token
	requestToken  string
	requestSecret string
}
