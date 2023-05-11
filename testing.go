package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/muesli/cache2go"
)

func setup() (
	getAccessTokenCallback CachedTokenCallback,
	setAccessTokenCallback SetTokenCallback,
	getRefreshTokenCallback CachedTokenCallback,
	setRefreshTokenCallback SetTokenCallback,
	err error) {
	cache := cache2go.Cache("go-ninjarmm-test")
	getAccessToken := func() (*string, error) {
		res, err := cache.Value("access-token")
		if err != nil {
			return nil, nil
		}
		token := res.Data().(string)
		return &token, nil
	}

	setAccessToken := func(token string, expiresIn float64) error {
		cache.Add("access-token", time.Duration(expiresIn), token)
		return nil
	}
	getRefreshToken := func() (*string, error) {
		res, err := cache.Value("refresh-token")
		if err != nil {
			return nil, nil
		}
		token := res.Data().(string)
		return &token, nil
	}
	setRefreshToken := func(token string, expiresIn float64) error {
		cache.Add("refresh-token", time.Duration(expiresIn), token)
		return nil
	}
	return getAccessToken, setAccessToken, getRefreshToken, setRefreshToken, nil
}

func loadTestData() (data TestData, err error) {
	err = loadDotEnv()
	if err != nil {
		return
	}
	rawData := os.Getenv("NINJARMM_TEST_DATA")
	err = json.Unmarshal([]byte(rawData), &data)
	if data.DeviceID == 0 {
		err = fmt.Errorf("'deviceId' missing from test data")
		return
	}
	if data.OrgID == 0 {
		err = fmt.Errorf("'orgId' missing from test data")
		return
	}
	return
}
