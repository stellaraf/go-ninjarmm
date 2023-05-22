package ninjarmm

import (
	"encoding/json"
	"fmt"
	"math"
	"net/url"
	"time"

	"github.com/go-resty/resty/v2"
)

type CachedTokenCallback func() (string, error)

type SetTokenCallback func(token string, expiresIn float64) error

type ninjaRMMAccessToken struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    int    `json:"expires_in"`
	TokenType    string `json:"token_type"`
	RefreshToken string `json:"refresh_token"`
	Scope        string `json:"scope"`
}

type NinjaRMMAuth struct {
	baseURL                 string
	clientID                string
	clientSecret            string
	encryption              bool
	encryptionPassphrase    string
	getAccessTokenCallback  CachedTokenCallback
	setAccessTokenCallback  SetTokenCallback
	getRefreshTokenCallback CachedTokenCallback
	setRefreshTokenCallback SetTokenCallback
	httpClient              *resty.Client
}

func (auth *NinjaRMMAuth) GetRefreshToken() (token string, err error) {
	rawToken, err := auth.getRefreshTokenCallback()
	if err != nil {
		return
	}
	if rawToken != "" && auth.encryption {
		decrypted := decrypt(auth.encryptionPassphrase, rawToken)
		token = decrypted
		return
	}
	return "", nil
}

func (auth *NinjaRMMAuth) GetNewToken() (token ninjaRMMAccessToken, err error) {
	refreshToken, err := auth.GetRefreshToken()
	if err != nil {
		return
	}
	q := url.Values{}
	q.Set("client_id", auth.clientID)
	q.Set("client_secret", auth.clientSecret)

	if refreshToken != "" {
		q.Set("grant_type", "refresh_token")
		q.Set("refresh_token", refreshToken)
	} else {
		q.Set("grant_type", "client_credentials")
		q.Set("scope", "monitoring management control offline_access")
	}
	req := auth.httpClient.R()
	req.SetHeader("content-type", "application/x-www-form-urlencoded")
	res, err := req.SetBody(q.Encode()).Post("/ws/oauth/token")

	if err != nil {
		return
	}
	err = checkForError(res)
	if err != nil {
		return
	}

	bodyBytes := res.Body()
	if res.StatusCode() >= 400 {
		errorDetail := string(bodyBytes)
		err = fmt.Errorf("failed to request new NinjaRMM access token due to %d error: '%s'", res.StatusCode(), errorDetail)
		return
	}

	err = json.Unmarshal(bodyBytes, &token)
	if err != nil {
		return
	}
	if isGenericError(token) || isRequestError(token) {
		errorDetail := getNinjaRMMError(token)
		err = fmt.Errorf("failed to get new NinjaRMM access token due to error: '%s'", errorDetail)
		return
	}
	return
}

func (auth *NinjaRMMAuth) GetAccessToken() (token string, err error) {
	cachedToken, err := auth.getAccessTokenCallback()
	if err != nil {
		return
	}
	if cachedToken == "" {
		newToken, err := auth.GetNewToken()
		if err != nil {
			return "", err
		}
		err = auth.CacheNewToken(newToken)
		if err != nil {
			return "", err
		}
		return newToken.AccessToken, nil
	}
	if auth.encryption {
		return decrypt(auth.encryptionPassphrase, cachedToken), nil
	}
	return cachedToken, nil
}

func (auth *NinjaRMMAuth) CacheNewToken(token ninjaRMMAccessToken) (err error) {
	err = auth.SetAccessToken(token)
	if err != nil {
		return
	}
	err = auth.SetRefreshToken(token.RefreshToken)
	return
}

func (auth *NinjaRMMAuth) SetRefreshToken(value string) (err error) {
	expiresAt := time.Now()
	expiresAt.AddDate(0, 0, 29)
	now := time.Now()
	expiresIn := math.Abs(float64(expiresAt.UnixMilli())-float64(now.UnixMilli())) / 1000
	if auth.encryption {
		encryptedToken := encrypt(value, auth.encryptionPassphrase)
		auth.setRefreshTokenCallback(encryptedToken, expiresIn)
		return
	}
	auth.setRefreshTokenCallback(value, expiresIn)
	return
}

func (auth *NinjaRMMAuth) SetAccessToken(token ninjaRMMAccessToken) (err error) {
	if auth.encryption {
		encrypted := encrypt(token.AccessToken, auth.encryptionPassphrase)
		auth.setAccessTokenCallback(encrypted, float64(token.ExpiresIn))
		return
	}
	auth.setAccessTokenCallback(token.AccessToken, float64(token.ExpiresIn))
	return
}

func createNinjaRMMAuth(
	baseURL, clientID, clientSecret string,
	encryption *string,
	getAccessTokenCallback CachedTokenCallback,
	setAccessTokenCallback SetTokenCallback,
	getRefreshTokenCallback CachedTokenCallback,
	setRefreshTokenCallback SetTokenCallback) (*NinjaRMMAuth, error) {
	var doEncrypt bool
	passphrase := ""
	if encryption == nil {
		doEncrypt = false
	} else {
		doEncrypt = true
		passphrase = *encryption
	}
	httpClient := resty.New()
	httpClient.SetBaseURL(baseURL)
	httpClient.SetHeader("user-agent", "go-ninjarmm")
	auth := &NinjaRMMAuth{
		baseURL:                 baseURL,
		clientID:                clientID,
		clientSecret:            clientSecret,
		encryption:              doEncrypt,
		encryptionPassphrase:    passphrase,
		getAccessTokenCallback:  getAccessTokenCallback,
		setAccessTokenCallback:  setAccessTokenCallback,
		getRefreshTokenCallback: getRefreshTokenCallback,
		setRefreshTokenCallback: setRefreshTokenCallback,
		httpClient:              httpClient,
	}
	return auth, nil
}
