package ninjarmm

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_isGenericError(t *testing.T) {
	t.Run("returns true", func(t *testing.T) {
		fixture := NinjaRMMBaseError{Error: "test error"}
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
		fixture := NinjaRMMRequestError{ErrorMessage: "test error", ResultCode: "test error", IncidentID: "1"}
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
		fixture := NinjaRMMAPIError{Error: "test error", ErrorDescription: "test error", ErrorCode: 400}
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
		fixture := NinjaRMMBaseError{Error: message}
		result := getNinjaRMMError(fixture)
		assert.Equal(t, message, result)
	})
	t.Run("error from request error", func(t *testing.T) {
		message := "test error"
		fixture := NinjaRMMRequestError{ErrorMessage: message, ResultCode: "test error", IncidentID: "1"}
		result := getNinjaRMMError(fixture)
		assert.Equal(t, message, result)
	})
	t.Run("error from api error", func(t *testing.T) {
		message := "test error"
		fixture := NinjaRMMAPIError{Error: "base error", ErrorDescription: message, ErrorCode: 400}
		result := getNinjaRMMError(fixture)
		assert.Equal(t, message, result)
	})
}

func Test_timeToFractional(t *testing.T) {
	t.Run("time to fractional", func(t *testing.T) {
		ts := time.Date(
			2023,      // year
			7,         // month
			18,        // day
			1,         // hour
			23,        // minute
			45,        // second
			670000076, // millisecond
			time.UTC,  // timezone
		)
		expected := 1689643425.670000076
		result := timeToFractional(ts)
		fs := "%.9f"
		assert.Equal(t, fmt.Sprintf(fs, expected), fmt.Sprintf(fs, result))
	})
}
