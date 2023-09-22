package ninjarmm

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func initClient() (client *Client, err error) {
	env, err := loadEnv()
	if err != nil {
		return
	}
	getAccessToken, setAccessToken, getRefreshToken, setRefreshToken, err := setup()
	if err != nil {
		return
	}
	client, err = New(
		env.BaseURL,
		env.ClientID,
		env.ClientSecret,
		nil,
		getAccessToken,
		setAccessToken,
		getRefreshToken,
		setRefreshToken,
	)
	return
}

func Test_NinjaRMMClient(t *testing.T) {
	testData, err := loadTestData()
	require.NoError(t, err)
	client, err := initClient()
	require.NoError(t, err)
	assert.NotNil(t, client)

	t.Run("organizations", func(t *testing.T) {
		t.Parallel()
		orgs, err := client.Organizations()
		require.NoError(t, err)
		assert.IsType(t, []OrganizationSummary{}, orgs)
	})
	t.Run("device", func(t *testing.T) {
		t.Parallel()
		data, err := client.Device(testData.DeviceID)
		require.NoError(t, err)
		assert.IsType(t, DeviceDetails{}, data)
	})
	t.Run("device custom fields", func(t *testing.T) {
		t.Parallel()
		data, err := client.DeviceCustomFields(testData.DeviceID)
		require.NoError(t, err)
		assert.IsType(t, map[string]any{}, data)
	})
	t.Run("organization", func(t *testing.T) {
		t.Parallel()
		data, err := client.Organization(testData.OrgID)
		require.NoError(t, err)
		assert.IsType(t, Organization{}, data)
	})

	t.Run("os patches", func(t *testing.T) {
		t.Parallel()
		data, err := client.OSPatches(testData.OrgID)
		require.NoError(t, err)
		assert.IsType(t, OSPatchReportQuery{}, data)
	})
	t.Run("os patch report", func(t *testing.T) {
		t.Parallel()
		data, err := client.OSPatchReport(testData.OrgID)
		require.NoError(t, err)
		assert.IsType(t, []OSPatchReportDetail{}, data)
	})
	t.Run("org locations", func(t *testing.T) {
		t.Parallel()
		data, err := client.OrganizationLocations(testData.OrgID)
		require.NoError(t, err)
		assert.IsType(t, []Location{}, data)
		assert.True(t, len(data) > 0)
	})
	t.Run("location", func(t *testing.T) {
		t.Parallel()
		data, err := client.Location(testData.OrgID, testData.LocationID)
		require.NoError(t, err)
		assert.IsType(t, &Location{}, data)
		assert.Equal(t, testData.LocationID, data.ID)
	})
	t.Run("maintenance", func(t *testing.T) {
		t.Parallel()
		start := time.Now().Add(time.Hour)
		end := start.Add(time.Hour)
		disabledFeatures := []string{"ALERTS"}
		err := client.ScheduleMaintenance(testData.DeviceID, start, end, disabledFeatures)
		require.NoError(t, err)
		err = client.CancelMaintenance(testData.DeviceID)
		require.NoError(t, err)
	})
	t.Run("maintenance error", func(t *testing.T) {
		t.Parallel()
		start := time.Now().Add(-(time.Hour * 49))
		end := time.Now().Add(-(time.Hour * 48))
		disabledFeatures := []string{"ALERTS"}
		err := client.ScheduleMaintenance(testData.DeviceID, start, end, disabledFeatures)
		assert.Error(t, err)
		assert.ErrorContains(t, err, fmt.Sprint(testData.DeviceID))
	})
}
