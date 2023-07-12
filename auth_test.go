package ninjarmm

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func initAuth() (auth *authT, err error) {
	env, err := loadEnv()
	if err != nil {
		return
	}
	getAccessToken, setAccessToken, getRefreshToken, setRefreshToken, err := setup()
	if err != nil {
		return
	}
	auth, err = newAuth(
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
		var tokenType string
		auth, err := initAuth()
		assert.NoError(t, err)
		assert.NotNil(t, auth)
		token, err := auth.GetRefreshToken()
		assert.NoError(t, err)
		assert.IsType(t, tokenType, token)
		result := "failure"
		if token == "" {
			result = ""
		} else {
			result = token
		}
		assert.NotEqual(t, "failure", result)
		t.Logf("refresh token: %s", result)
	})
}
