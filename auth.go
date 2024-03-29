package ninjarmm

import (
	"encoding/json"
	"fmt"
	"net/url"
	"time"

	"github.com/go-resty/resty/v2"
)

type CachedTokenCallback func() (string, error)

type SetTokenCallback func(token string, expiresIn time.Duration) error

type ninjaRMMAccessToken struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    int    `json:"expires_in"`
	TokenType    string `json:"token_type"`
	RefreshToken string `json:"refresh_token"`
	Scope        string `json:"scope"`
}

func (token *ninjaRMMAccessToken) Expiry() time.Duration {
	return time.Duration(token.ExpiresIn) * time.Second
}

type authT struct {
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

func (authT *authT) RefreshTokenExpiry() time.Duration {
	return time.Duration(REFRESH_TOKEN_EXPIRY_DAYS*24) * time.Hour
}

func (auth *authT) GetRefreshToken() (string, error) {
	rawToken, err := auth.getRefreshTokenCallback()
	if err != nil {
		return "", err
	}
	if rawToken != "" && auth.encryption {
		decrypted := decrypt(auth.encryptionPassphrase, rawToken)
		return decrypted, nil
	}
	return "", nil
}

func (auth *authT) GetNewToken() (*ninjaRMMAccessToken, error) {
	refreshToken, err := auth.GetRefreshToken()
	if err != nil {
		return nil, err
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
		return nil, err
	}
	err = checkForError(res)
	if err != nil {
		return nil, err
	}

	bodyBytes := res.Body()
	if res.StatusCode() >= 400 {
		errorDetail := string(bodyBytes)
		err = fmt.Errorf("failed to request new NinjaRMM access token due to %d error: '%s'", res.StatusCode(), errorDetail)
		return nil, err
	}
	var token *ninjaRMMAccessToken
	err = json.Unmarshal(bodyBytes, &token)
	if err != nil {
		return nil, err
	}
	if token == nil {
		err = fmt.Errorf("failed to get new NinjaRMM access token")
		return nil, err
	}
	if isGenericError(token) || isRequestError(token) {
		errorDetail := getNinjaRMMError(token)
		err = fmt.Errorf("failed to get new NinjaRMM access token due to error: '%s'", errorDetail)
		return nil, err
	}
	return token, nil
}

func (auth *authT) GetAccessToken() (string, error) {
	cachedToken, err := auth.getAccessTokenCallback()
	if err != nil {
		return "", err
	}
	if cachedToken == "" {
		newToken, err := auth.GetNewToken()
		if err != nil {
			return "", err
		}
		if newToken == nil {
			err = fmt.Errorf("failed to retrieve new access token")
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

func (auth *authT) CacheNewToken(token *ninjaRMMAccessToken) error {
	err := auth.SetAccessToken(token)
	if err != nil {
		return err
	}
	err = auth.SetRefreshToken(token.RefreshToken)
	if err != nil {
		return err
	}
	return nil
}

func (auth *authT) SetRefreshToken(value string) error {
	if auth.encryption {
		encryptedToken := encrypt(value, auth.encryptionPassphrase)
		err := auth.setRefreshTokenCallback(encryptedToken, auth.RefreshTokenExpiry())
		if err != nil {
			return err
		}
		return nil
	}
	err := auth.setRefreshTokenCallback(value, auth.RefreshTokenExpiry())
	if err != nil {
		return err
	}
	return nil
}

func (auth *authT) SetAccessToken(token *ninjaRMMAccessToken) error {
	if auth.encryption {
		encrypted := encrypt(token.AccessToken, auth.encryptionPassphrase)
		err := auth.setAccessTokenCallback(encrypted, token.Expiry())
		if err != nil {
			return err
		}
		return nil
	}
	err := auth.setAccessTokenCallback(token.AccessToken, token.Expiry())
	if err != nil {
		return err
	}
	return nil
}

func newAuth(
	baseURL, clientID, clientSecret string,
	encryption *string,
	getAccessTokenCallback CachedTokenCallback,
	setAccessTokenCallback SetTokenCallback,
	getRefreshTokenCallback CachedTokenCallback,
	setRefreshTokenCallback SetTokenCallback) (*authT, error) {
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
	auth := &authT{
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
