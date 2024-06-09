package check_test

import (
	"testing"

	"github.com/stellaraf/go-ninjarmm/internal/check"
	"github.com/stellaraf/go-ninjarmm/internal/types"
	"github.com/stretchr/testify/assert"
)

func Test_IsGenericError(t *testing.T) {
	t.Run("returns true", func(t *testing.T) {
		fixture := types.NinjaRMMBaseError{Error: "test error"}
		result := check.IsGenericError(fixture)
		assert.True(t, result)
	})
	t.Run("returns false", func(t *testing.T) {
		fixture := struct{ SomethingElse string }{SomethingElse: "test error"}
		result := check.IsGenericError(fixture)
		assert.False(t, result)
	})
}

func Test_isRequestError(t *testing.T) {
	t.Run("returns true", func(t *testing.T) {
		fixture := types.NinjaRMMRequestError{ErrorMessage: "test error", ResultCode: "test error", IncidentID: "1"}
		result := check.IsRequestError(fixture)
		assert.True(t, result)
	})
	t.Run("returns false", func(t *testing.T) {
		fixture := struct{ SomethingElse string }{SomethingElse: "test error"}
		result := check.IsRequestError(fixture)
		assert.False(t, result)
	})
}

func Test_isApiError(t *testing.T) {
	t.Run("returns true", func(t *testing.T) {
		fixture := types.NinjaRMMAPIError{Error: "test error", ErrorDescription: "test error", ErrorCode: 400}
		result := check.IsApiError(fixture)
		assert.True(t, result)
	})
	t.Run("returns false", func(t *testing.T) {
		fixture := struct{ SomethingElse string }{SomethingElse: "test error"}
		result := check.IsApiError(fixture)
		assert.False(t, result)
	})
}

func Test_GetNinjaRMMError(t *testing.T) {
	t.Run("error from generic error", func(t *testing.T) {
		message := "test error"
		fixture := types.NinjaRMMBaseError{Error: message}
		result := check.GetNinjaRMMError(fixture)
		assert.Equal(t, message, result)
	})
	t.Run("error from request error", func(t *testing.T) {
		message := "test error"
		fixture := types.NinjaRMMRequestError{ErrorMessage: message, ResultCode: "test error", IncidentID: "1"}
		result := check.GetNinjaRMMError(fixture)
		assert.Equal(t, message, result)
	})
	t.Run("error from api error", func(t *testing.T) {
		message := "test error"
		fixture := types.NinjaRMMAPIError{Error: "base error", ErrorDescription: message, ErrorCode: 400}
		result := check.GetNinjaRMMError(fixture)
		assert.Equal(t, message, result)
	})
}
