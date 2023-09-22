package ninjarmm

import (
	"encoding/json"
	"fmt"
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
	getAccessToken := func() (string, error) {
		res, err := cache.Value("access-token")
		if err != nil {
			return "", nil
		}
		token := res.Data().(string)
		return token, nil
	}

	setAccessToken := func(token string, expiresIn time.Duration) error {
		cache.Add("access-token", expiresIn, token)
		return nil
	}
	getRefreshToken := func() (string, error) {
		res, err := cache.Value("refresh-token")
		if err != nil {
			return "", nil
		}
		token := res.Data().(string)
		return token, nil
	}
	setRefreshToken := func(token string, expiresIn time.Duration) error {
		cache.Add("refresh-token", expiresIn, token)
		return nil
	}
	return getAccessToken, setAccessToken, getRefreshToken, setRefreshToken, nil
}

func loadTestData() (data testDataT, err error) {
	env, err := loadEnv()
	if err != nil {
		return
	}
	err = json.Unmarshal([]byte(env.TestData), &data)
	required := map[string]int{"deviceId": data.DeviceID, "orgId": data.OrgID, "locationId": data.LocationID}

	for k, v := range required {
		if v == 0 {
			err = fmt.Errorf("'%s' missing from test data", k)
			return
		}
	}
	return
}
