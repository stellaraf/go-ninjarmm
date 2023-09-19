package ninjarmm

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_ActivityTime(t *testing.T) {

	type S struct {
		Timestamp Timestamp `json:"timestamp"`
	}
	var ts Timestamp
	asJSON := []byte(`{"timestamp":1689605132.309}`)

	t.Run("unmarshal", func(t *testing.T) {
		var s *S
		err := json.Unmarshal(asJSON, &s)
		ts = s.Timestamp
		require.NoError(t, err)
		assert.IsType(t, Timestamp{}, s.Timestamp)
		assert.Equal(t, s.Timestamp.Month(), time.July)
	})
	t.Run("marshal", func(t *testing.T) {
		s := S{Timestamp: ts}
		result, err := json.Marshal(&s)
		require.NoError(t, err)
		assert.Equal(t, asJSON, result)
	})
}

func Test_MaintenanceRequest(t *testing.T) {
	t.Run("maintenance request json marshaller", func(t *testing.T) {
		start := time.Now()
		end := start.Add(time.Hour)
		disabledFeatures := []string{
			"ONE",
			"two",
			"THREE ",
		}
		expected := fmt.Sprintf(`{"start":%.9f,"end":%.9f,"disabledFeatures":["ONE","TWO","THREE"]}`, timeToFractional(start), timeToFractional(end))
		mr := &MaintenanceRequest{
			Start:            start,
			End:              end,
			DisabledFeatures: disabledFeatures,
		}
		result, err := json.Marshal(&mr)
		require.NoError(t, err)
		assert.JSONEq(t, expected, string(result))
	})
}
