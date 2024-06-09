package ninjarmm

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/stellaraf/go-ninjarmm/internal/util"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_Timestamp(t *testing.T) {

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

func Test_InstallDate(t *testing.T) {
	type S struct {
		InstallDate InstallDate `json:"installDate"`
	}
	var date InstallDate
	asJSON := []byte(`{"installDate":"2024-06-08"}`)
	t.Run("unmarshal", func(t *testing.T) {
		var s *S
		err := json.Unmarshal(asJSON, &s)
		require.NoError(t, err)
		date = s.InstallDate
		assert.IsType(t, InstallDate{}, s.InstallDate)
		assert.Equal(t, s.InstallDate.Year(), 2024)
		assert.Equal(t, s.InstallDate.Month(), time.June)
		assert.Equal(t, s.InstallDate.Day(), 8)
	})
	t.Run("marshal", func(t *testing.T) {
		s := S{InstallDate: date}
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
		expected := fmt.Sprintf(`{"start":%.9f,"end":%.9f,"disabledFeatures":["ONE","TWO","THREE"]}`, util.TimeToFractional(start), util.TimeToFractional(end))
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
