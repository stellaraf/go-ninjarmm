package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func initAuth() (auth *NinjaRMMAuth, err error) {
	env, err := LoadEnv()
	if err != nil {
		return
	}
	getAccessToken, setAccessToken, getRefreshToken, setRefreshToken, err := setup()
	if err != nil {
		return
	}
	auth, err = createNinjaRMMAuth(
		env.BaseURL,
		env.ClientID,
		env.ClientSecret,
		&env.EncryptionPassphrase,
		getAccessToken,
		setAccessToken,
		getRefreshToken,
		setRefreshToken,
	)
	return
}

func Test_Auth(t *testing.T) {

	t.Run("get an access token", func(t *testing.T) {
		auth, err := initAuth()
		assert.NoError(t, err)
		token, err := auth.GetAccessToken()
		assert.NoError(t, err)
		assert.IsType(t, "", token)
		t.Logf("access token: %s", token)

	})

	t.Run("get refresh token", func(t *testing.T) {
		var tokenType *string
		auth, err := initAuth()
		assert.NoError(t, err)
		assert.NotNil(t, auth)
		token, err := auth.GetRefreshToken()
		assert.NoError(t, err)
		assert.IsType(t, tokenType, token)
		result := "failure"
		if token == nil {
			result = "<nil>"
		} else if *token == "" {
			result = "none"
		} else {
			result = *token
		}
		assert.NotEqual(t, "failure", result)
		t.Logf("refresh token: %s", result)
	})
}
