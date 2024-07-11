// Based on code: https://github.com/dghubble/oauth1/blob/main/examples/twitter-login.go

package twitter

import (
	"fmt"
	"log"

	"github.com/dghubble/oauth1"
	oauth_twitter "github.com/dghubble/oauth1/twitter"
	"github.com/ignytis/deletex/system/config"
)

func login(config oauth1.Config) (requestToken string, err error) {
	requestToken, _, err = config.RequestToken()
	if err != nil {
		return "", err
	}
	authorizationURL, err := config.AuthorizationURL(requestToken)
	if err != nil {
		return "", err
	}
	fmt.Printf("Open this URL in your browser:\n%s\n", authorizationURL.String())
	return requestToken, err
}

func receivePIN(config oauth1.Config, requestToken string) (*oauth1.Token, error) {
	fmt.Printf("Paste your PIN here: ")
	var verifier string
	_, err := fmt.Scanf("%s", &verifier)
	if err != nil {
		return nil, err
	}

	accessToken, accessSecret, err := config.AccessToken(requestToken, "secret does not matter", verifier)
	if err != nil {
		return nil, err
	}
	return oauth1.NewToken(accessToken, accessSecret), err
}

func mustGetTokens(cfg *config.Config) (string, string) {
	config := oauth1.Config{
		ConsumerKey:    cfg.Auth.ConsumerKey,
		ConsumerSecret: cfg.Auth.ConsumerKeySecret,
		CallbackURL:    "oob",
		Endpoint:       oauth_twitter.AuthorizeEndpoint,
	}

	requestToken, err := login(config)
	if err != nil {
		log.Fatalf("Failed to get Request Token: %s", err.Error())
	}
	accessToken, err := receivePIN(config, requestToken)
	if err != nil {
		log.Fatalf("Failed to get Access Token: %s", err.Error())
	}

	return accessToken.Token, accessToken.TokenSecret
}
