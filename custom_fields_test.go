package ninjarmm_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.stellar.af/go-ninjarmm"
)

func Test_ParseCustomFields(t *testing.T) {
	t.Run("base", func(t *testing.T) {
		t.Parallel()
		type CF struct {
			FieldOne string `json:"fieldOne"`
			FieldTwo int    `json:"fieldTwo"`
		}
		m := map[string]any{
			"fieldOne": "value",
			"fieldTwo": 2,
		}
		cf, err := ninjarmm.ParseCustomFields[CF](m)
		require.NoError(t, err)
		require.NotNil(t, cf)
		assert.Equal(t, m["fieldOne"], cf.FieldOne)
		assert.Equal(t, m["fieldTwo"], cf.FieldTwo)
	})
}
