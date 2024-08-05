package auth_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.stellar.af/go-ninjarmm/internal/auth"
	"go.stellar.af/go-ninjarmm/internal/test"
	"go.stellar.af/go-ninjarmm/internal/types"
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
	t.Run("get tokens", func(t *testing.T) {
		t.Parallel()
		token, err := Auth.GetAccessToken()
		require.NoError(t, err)
		assert.NotEmpty(t, token)

		token, err = Auth.GetRefreshToken()
		require.NoError(t, err)
		assert.NotEmpty(t, token)
		t.Cleanup(func() {
			test.NewTokenCache().Clear()
		})
	})
	t.Run("auth error", func(t *testing.T) {
		t.Parallel()
		env, err := test.LoadEnv()
		require.NoError(t, err)
		tc := test.NewTokenCache()
		a, err := auth.New(env.BaseURL, "invalid", "invalid", nil, tc)
		require.NoError(t, err)
		_, err = a.GetAccessToken()
		require.Error(t, err)
		assert.IsType(t, &types.Error{}, err)
	})
}
