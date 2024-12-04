package auth

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"time"

	"github.com/go-resty/resty/v2"
	"go.stellar.af/go-ninjarmm/internal/types"
	"go.stellar.af/go-utils/encryption"
)

type TokenCache interface {
	GetAccessToken() (string, error)
	SetAccessToken(string, time.Duration) error
	GetRefreshToken() (string, error)
	SetRefreshToken(string, time.Duration) error
}

var (
	REFRESH_TOKEN_EXPIRY_DAYS uint = 29
)

type AccessToken struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    int    `json:"expires_in"`
	TokenType    string `json:"token_type"`
	RefreshToken string `json:"refresh_token"`
	Scope        string `json:"scope"`
}

func (token *AccessToken) Expiry() time.Duration {
	return time.Duration(token.ExpiresIn-60) * time.Second
}

type Auth struct {
	baseURL              string
	clientID             string
	clientSecret         string
	encryption           bool
	encryptionPassphrase string
	tokenCache           TokenCache
	httpClient           *resty.Client
}

func (authT *Auth) RefreshTokenExpiry() time.Duration {
	return time.Duration(REFRESH_TOKEN_EXPIRY_DAYS*24) * time.Hour
}

func (auth *Auth) GetRefreshToken() (string, error) {
	rawToken, err := auth.tokenCache.GetRefreshToken()
	if err != nil {
		return "", err
	}
	if rawToken != "" && auth.encryption {
		return encryption.Decrypt(auth.encryptionPassphrase, rawToken)
	}
	return rawToken, nil
}

func (auth *Auth) GetNewToken() (*AccessToken, error) {
	refreshToken, err := auth.GetRefreshToken()
	if err != nil && !errors.Is(err, types.ErrTokenCacheMiss) {
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
	req.
		SetHeader("content-type", "application/x-www-form-urlencoded").
		SetBody(q.Encode()).
		SetError(&types.Error{})
	res, err := req.Post("/ws/oauth/token")
	if err != nil {
		return nil, err
	}

	if res.IsError() {
		err := res.Error().(*types.Error)
		if refreshToken != "" && err.Message == "invalid_token" {
			auth.SetRefreshToken("")
			return auth.GetNewToken()
		}
		return nil, err
	}

	bodyBytes := res.Body()
	if res.StatusCode() >= 400 {
		errorDetail := string(bodyBytes)
		err = fmt.Errorf("failed to request new NinjaRMM access token due to %d error: '%s'", res.StatusCode(), errorDetail)
		return nil, err
	}
	var token *AccessToken
	err = json.Unmarshal(bodyBytes, &token)
	if err != nil {
		return nil, err
	}
	if token == nil {
		err = fmt.Errorf("failed to get new NinjaRMM access token")
		return nil, err
	}
	return token, nil
}

func (auth *Auth) GetAccessToken() (string, error) {
	cachedToken, err := auth.tokenCache.GetAccessToken()
	if err != nil && !errors.Is(err, types.ErrTokenCacheMiss) {
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
		return encryption.Decrypt(auth.encryptionPassphrase, cachedToken)
	}
	return cachedToken, nil
}

func (auth *Auth) CacheNewToken(token *AccessToken) error {
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

func (auth *Auth) SetRefreshToken(value string) error {
	if auth.encryption {
		encryptedToken, err := encryption.Encrypt(value, auth.encryptionPassphrase)
		if err != nil {
			return err
		}
		err = auth.tokenCache.SetRefreshToken(encryptedToken, auth.RefreshTokenExpiry())
		if err != nil {
			return err
		}
		return nil
	}
	err := auth.tokenCache.SetRefreshToken(value, auth.RefreshTokenExpiry())
	if err != nil {
		return err
	}
	return nil
}

func (auth *Auth) SetAccessToken(token *AccessToken) error {
	if auth.encryption {
		encrypted, err := encryption.Encrypt(token.AccessToken, auth.encryptionPassphrase)
		if err != nil {
			return err
		}
		err = auth.tokenCache.SetAccessToken(encrypted, token.Expiry())
		if err != nil {
			return err
		}
		return nil
	}
	err := auth.tokenCache.SetAccessToken(token.AccessToken, token.Expiry())
	if err != nil {
		return err
	}
	return nil
}

func New(
	baseURL, clientID, clientSecret string,
	encryption *string,
	tokenCache TokenCache,
) (*Auth, error) {
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
	auth := &Auth{
		baseURL:              baseURL,
		clientID:             clientID,
		clientSecret:         clientSecret,
		encryption:           doEncrypt,
		encryptionPassphrase: passphrase,
		tokenCache:           tokenCache,
		httpClient:           httpClient,
	}
	return auth, nil
}
