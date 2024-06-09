package auth_test

import (
	"testing"

	"github.com/stellaraf/go-ninjarmm/internal/auth"
	"github.com/stellaraf/go-ninjarmm/internal/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func initAuth() (*auth.Auth, error) {
	env, err := test.LoadEnv()
	if err != nil {
		return nil, err
	}
	tc := test.NewTokenCache()
	return auth.New(
		env.BaseURL,
		env.ClientID,
		env.ClientSecret,
		nil,
		tc,
	)
}

func Test_Auth(t *testing.T) {

	t.Run("get an access token", func(t *testing.T) {
		auth, err := initAuth()
		require.NoError(t, err)
		token, err := auth.GetAccessToken()
		require.NoError(t, err)
		assert.IsType(t, "", token)
		t.Logf("access token: %s", token)
	})

	t.Run("get refresh token", func(t *testing.T) {
		var tokenType string
		auth, err := initAuth()
		require.NoError(t, err)
		assert.NotNil(t, auth)
		token, err := auth.GetRefreshToken()
		require.NoError(t, err)
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
