package ninjarmm

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_isGenericError(t *testing.T) {
	t.Run("returns true", func(t *testing.T) {
		fixture := ninjaRMMBaseError{Error: "test error"}
		result := isGenericError(fixture)
		assert.True(t, result)
	})
	t.Run("returns false", func(t *testing.T) {
		fixture := struct{ SomethingElse string }{SomethingElse: "test error"}
		result := isGenericError(fixture)
		assert.False(t, result)
	})
}

func Test_isRequestError(t *testing.T) {
	t.Run("returns true", func(t *testing.T) {
		fixture := ninjaRMMRequestError{ErrorMessage: "test error", ResultCode: "test error", IncidentID: "1"}
		result := isRequestError(fixture)
		assert.True(t, result)
	})
	t.Run("returns false", func(t *testing.T) {
		fixture := struct{ SomethingElse string }{SomethingElse: "test error"}
		result := isRequestError(fixture)
		assert.False(t, result)
	})
}

func Test_isApiError(t *testing.T) {
	t.Run("returns true", func(t *testing.T) {
		fixture := ninaRMMAPIError{Error: "test error", ErrorDescription: "test error", ErrorCode: 400}
		result := isApiError(fixture)
		assert.True(t, result)
	})
	t.Run("returns false", func(t *testing.T) {
		fixture := struct{ SomethingElse string }{SomethingElse: "test error"}
		result := isApiError(fixture)
		assert.False(t, result)
	})
}

func Test_getNinjaRMMError(t *testing.T) {
	t.Run("error from generic error", func(t *testing.T) {
		message := "test error"
		fixture := ninjaRMMBaseError{Error: message}
		result := getNinjaRMMError(fixture)
		assert.Equal(t, message, result)
	})
	t.Run("error from request error", func(t *testing.T) {
		message := "test error"
		fixture := ninjaRMMRequestError{ErrorMessage: message, ResultCode: "test error", IncidentID: "1"}
		result := getNinjaRMMError(fixture)
		assert.Equal(t, message, result)
	})
	t.Run("error from api error", func(t *testing.T) {
		message := "test error"
		fixture := ninaRMMAPIError{Error: "base error", ErrorDescription: message, ErrorCode: 400}
		result := getNinjaRMMError(fixture)
		assert.Equal(t, message, result)
	})
}
