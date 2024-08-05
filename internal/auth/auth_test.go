package auth_test

import (
	"testing"

	"github.com/stellaraf/go-ninjarmm/internal/auth"
	"github.com/stellaraf/go-ninjarmm/internal/test"
	"github.com/stellaraf/go-ninjarmm/internal/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var Auth *auth.Auth

func init() {
	env, err := test.LoadEnv()
	if err != nil {
		panic(err)
	}
	tc := test.NewTokenCache()
	a, err := auth.New(
		env.BaseURL,
		env.ClientID,
		env.ClientSecret,
		nil,
		tc,
	)
	if err != nil {
		panic(err)
	}
	Auth = a
}

func Test_Auth(t *testing.T) {
	t.Run("get an access token", func(t *testing.T) {
		token, err := Auth.GetAccessToken()
		require.NoError(t, err)
		assert.NotEmpty(t, token)
	})

	t.Run("get refresh token", func(t *testing.T) {
		token, err := Auth.GetRefreshToken()
		require.NoError(t, err)
		assert.NotEmpty(t, token)
	})
	t.Run("auth error", func(t *testing.T) {
		t.Parallel()
		env, err := test.LoadEnv()
		require.NoError(t, err)
		tc := test.NewTokenCache()
		a, err := auth.New(env.BaseURL, "invalid", "invalid", nil, tc)
		require.NoError(t, err)
		_, err = a.GetAccessToken()
		assert.IsType(t, &types.Error{}, err)
	})
}
