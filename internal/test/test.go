package test

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/muesli/cache2go"
	"github.com/stellaraf/go-utils/environment"
)

type TestData struct {
	OrgID        int    `json:"orgId"`
	DeviceID     int    `json:"deviceId"`
	LocationID   int    `json:"locationId"`
	SoftwareName string `json:"softwareName"`
	RoleID       int    `json:"roleId"`
}

type Environment struct {
	ClientID             string `env:"CLIENT_ID"`
	ClientSecret         string `env:"CLIENT_SECRET"`
	BaseURL              string `env:"BASE_URL"`
	EncryptionPassphrase string `env:"ENCRYPTION_PASSPHRASE"`
	TestData             string `env:"TEST_DATA"`
}

type TokenCache struct {
	cache *cache2go.CacheTable
}

func (tc *TokenCache) GetAccessToken() (string, error) {
	res, err := tc.cache.Value("access-token")
	if err != nil {
		return "", nil
	}
	token := res.Data().(string)
	return token, nil
}

func (tc *TokenCache) SetAccessToken(token string, expiresIn time.Duration) error {
	tc.cache.Add("access-token", expiresIn, token)
	return nil
}

func (tc *TokenCache) GetRefreshToken() (string, error) {
	res, err := tc.cache.Value("refresh-token")
	if err != nil {
		return "", nil
	}
	token := res.Data().(string)
	return token, nil
}

func (tc *TokenCache) SetRefreshToken(token string, expiresIn time.Duration) error {
	tc.cache.Add("refresh-token", expiresIn, token)
	return nil
}

func LoadEnv() (env Environment, err error) {
	isCI := os.Getenv("CI") == "true"
	err = environment.Load(&env, &environment.EnvironmentOptions{
		DotEnv: !isCI,
	})
	return
}

func NewTokenCache() *TokenCache {
	cache := cache2go.Cache("go-ninjarmm-test")
	return &TokenCache{cache: cache}
}

func LoadTestData() (data TestData, err error) {
	env, err := LoadEnv()
	if err != nil {
		return
	}
	err = json.Unmarshal([]byte(env.TestData), &data)
	required := map[string]any{
		"deviceId":     data.DeviceID,
		"orgId":        data.OrgID,
		"locationId":   data.LocationID,
		"softwareName": data.SoftwareName,
	}

	for k, v := range required {
		if v == 0 {
			err = fmt.Errorf("'%s' missing from test data", k)
			return
		}
	}
	return
}
