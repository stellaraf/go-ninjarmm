package ninjarmm_test

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stellaraf/go-ninjarmm"
	"github.com/stretchr/testify/assert"
)

func Test_ActivityTime(t *testing.T) {

	type S struct {
		Timestamp ninjarmm.Timestamp `json:"timestamp"`
	}
	t.Run("parse timestamp", func(t *testing.T) {
		var s *S
		// value := float64(1689605132.309)
		j := []byte(`{"timestamp":1689605132.309}`)
		err := json.Unmarshal(j, &s)
		assert.NoError(t, err)
		assert.IsType(t, ninjarmm.Timestamp{}, s.Timestamp)
		assert.Equal(t, s.Timestamp.Month(), time.July)
	})
}
