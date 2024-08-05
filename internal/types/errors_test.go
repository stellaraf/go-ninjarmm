package types_test

import (
	"encoding/json"

	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.stellar.af/go-ninjarmm/internal/types"
)

func Test_Error(t *testing.T) {
	t.Run("base", func(t *testing.T) {
		t.Parallel()
		e := []byte(`{"error":"an error"}`)
		var ninjaErr types.Error
		err := json.Unmarshal(e, &ninjaErr)
		require.NoError(t, err)
		assert.Equal(t, "an error", ninjaErr.Message)
	})
	t.Run("api", func(t *testing.T) {
		t.Parallel()
		e := []byte(`{"error":"an error","error_description":"description of error","error_code":100}`)
		var ninjaErr types.Error
		err := json.Unmarshal(e, &ninjaErr)
		require.NoError(t, err)
		assert.Equal(t, "an error", ninjaErr.Message)
		assert.Equal(t, "description of error", ninjaErr.Description)
		assert.Equal(t, "100", ninjaErr.Code)
	})
	t.Run("request", func(t *testing.T) {
		t.Parallel()
		e := []byte(`{"errorMessage":"an error","resultCode":"a_result_code","incidentId":"12345"}`)
		var ninjaErr types.Error
		err := json.Unmarshal(e, &ninjaErr)
		require.NoError(t, err)
		assert.Equal(t, "an error", ninjaErr.Message)
		assert.Equal(t, "a_result_code", ninjaErr.Code)
		assert.Equal(t, "12345", ninjaErr.Details["incidentId"])
	})
	t.Run("put", func(t *testing.T) {
		t.Parallel()
		e := []byte(`{"errorMessage":{"code":"some_error_code","params":{"key":"value"}}}`)
		var ninjaErr types.Error
		err := json.Unmarshal(e, &ninjaErr)
		require.NoError(t, err)
		assert.Equal(t, "some_error_code", ninjaErr.Message)
		assert.Empty(t, ninjaErr.Code)
		assert.Equal(t, "value", ninjaErr.Details["key"])
	})
	t.Run("string", func(t *testing.T) {
		t.Parallel()
		e := []byte(`"a string error"`)
		var ninjaErr types.Error
		err := json.Unmarshal(e, &ninjaErr)
		require.NoError(t, err)
		assert.Equal(t, `"a string error"`, ninjaErr.Message)
	})
}
