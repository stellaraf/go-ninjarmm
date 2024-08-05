package ninjarmm_test

import (
	"encoding/json"
	"fmt"
	"regexp"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.stellar.af/go-ninjarmm"
	"go.stellar.af/go-ninjarmm/internal/test"
	"go.stellar.af/go-ninjarmm/internal/util"
)

func Test_Timestamp(t *testing.T) {

	type S struct {
		Timestamp ninjarmm.Timestamp `json:"timestamp"`
	}
	var ts ninjarmm.Timestamp
	asJSON := []byte(`{"timestamp":1689605132.309}`)

	t.Run("unmarshal", func(t *testing.T) {
		var s *S
		err := json.Unmarshal(asJSON, &s)
		ts = s.Timestamp
		require.NoError(t, err)
		assert.IsType(t, ninjarmm.Timestamp{}, s.Timestamp)
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
		InstallDate ninjarmm.InstallDate `json:"installDate"`
	}
	var date ninjarmm.InstallDate
	asJSON := []byte(`{"installDate":"2024-06-08"}`)
	t.Run("unmarshal", func(t *testing.T) {
		var s *S
		err := json.Unmarshal(asJSON, &s)
		require.NoError(t, err)
		date = s.InstallDate
		assert.IsType(t, ninjarmm.InstallDate{}, s.InstallDate)
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
		mr := &ninjarmm.MaintenanceRequest{
			Start:            start,
			End:              end,
			DisabledFeatures: disabledFeatures,
		}
		result, err := json.Marshal(&mr)
		require.NoError(t, err)
		assert.JSONEq(t, expected, string(result))
	})
}

func TestDevices_Filter(t *testing.T) {
	t.Parallel()
	td, err := test.LoadTestData()
	require.NoError(t, err)
	client, err := initClient()
	require.NoError(t, err)
	df := ninjarmm.NewDeviceFilter().Org(ninjarmm.EQ, td.OrgID)
	devices, err := client.Devices(df)
	require.NoError(t, err)
	filtered := devices.Filter(func(d ninjarmm.Device) bool {
		return d.ID == td.DeviceID
	})
	require.Len(t, filtered, 1)
	assert.Equal(t, td.DeviceID, filtered[0].ID)
}

func TestDevices_MatchName(t *testing.T) {
	t.Parallel()
	td, err := test.LoadTestData()
	require.NoError(t, err)
	client, err := initClient()
	require.NoError(t, err)
	df := ninjarmm.NewDeviceFilter().Org(ninjarmm.EQ, td.OrgID)
	devices, err := client.Devices(df)
	require.NoError(t, err)
	filtered := devices.MatchName(regexp.MustCompile(".+"))
	assert.Equal(t, len(devices), len(filtered))
}
