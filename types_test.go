package ninjarmm

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_ActivityTime(t *testing.T) {

	type S struct {
		Timestamp Timestamp `json:"timestamp"`
	}
	t.Run("parse timestamp", func(t *testing.T) {
		var s *S
		j := []byte(`{"timestamp":1689605132.309}`)
		err := json.Unmarshal(j, &s)
		assert.NoError(t, err)
		assert.IsType(t, Timestamp{}, s.Timestamp)
		assert.Equal(t, s.Timestamp.Month(), time.July)
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
		assert.NoError(t, err)
		assert.JSONEq(t, expected, string(result))
	})
}
