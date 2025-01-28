package ninjarmm_test

import (
	"fmt"
	"regexp"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.stellar.af/go-ninjarmm"
	"go.stellar.af/go-ninjarmm/internal/test"
)

func initClient() (client *ninjarmm.Client, err error) {
	env, err := test.LoadEnv()
	if err != nil {
		return
	}
	tc := test.NewTokenCache()
	client, err = ninjarmm.New(
		env.BaseURL,
		env.ClientID,
		env.ClientSecret,
		nil,
		tc,
	)
	return
}

func Test_NinjaRMMClient(t *testing.T) {
	testData, err := test.LoadTestData()
	require.NoError(t, err)
	client, err := initClient()
	require.NoError(t, err)
	assert.NotNil(t, client)

	t.Run("organizations", func(t *testing.T) {
		t.Parallel()
		orgs, err := client.Organizations()
		require.NoError(t, err)
		assert.IsType(t, []ninjarmm.OrganizationSummary{}, orgs)
	})
	t.Run("device", func(t *testing.T) {
		t.Parallel()
		data, err := client.Device(testData.DeviceID)
		require.NoError(t, err)
		assert.IsType(t, &ninjarmm.DeviceDetails{}, data)
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
		assert.IsType(t, &ninjarmm.Organization{}, data)
	})
	t.Run("organization custom fields", func(t *testing.T) {
		t.Parallel()
		data, err := client.OrganizationCustomFields(testData.OrgID)
		require.NoError(t, err)
		assert.IsType(t, map[string]any{}, data)
		assert.True(t, len(data) > 0, "custom fields are empty")
	})

	t.Run("os patches", func(t *testing.T) {
		t.Parallel()
		df := ninjarmm.NewDeviceFilter().Org(ninjarmm.EQ, testData.OrgID)
		data, err := client.OSPatches(df)
		require.NoError(t, err)
		assert.IsType(t, &ninjarmm.OSPatchReportQuery{}, data)
	})
	t.Run("os patch report", func(t *testing.T) {
		t.Parallel()
		df := ninjarmm.NewDeviceFilter().Org(ninjarmm.EQ, testData.OrgID)
		data, err := client.OSPatchReport(df)
		require.NoError(t, err)
		assert.IsType(t, []ninjarmm.OSPatchReportDetail{}, data)
	})
	t.Run("org devices", func(t *testing.T) {
		t.Parallel()
		data, err := client.OrganizationDevices(testData.OrgID)
		require.NoError(t, err)
		assert.IsType(t, []ninjarmm.Device{}, data)
		assert.True(t, len(data) > 0)
	})
	t.Run("org locations", func(t *testing.T) {
		t.Parallel()
		data, err := client.OrganizationLocations(testData.OrgID)
		require.NoError(t, err)
		assert.IsType(t, []ninjarmm.Location{}, data)
		assert.True(t, len(data) > 0)
	})
	t.Run("location", func(t *testing.T) {
		t.Parallel()
		data, err := client.Location(testData.OrgID, testData.LocationID)
		require.NoError(t, err)
		assert.IsType(t, &ninjarmm.Location{}, data)
		assert.Equal(t, testData.LocationID, data.ID)
	})
	t.Run("location custom fields", func(t *testing.T) {
		t.Parallel()
		data, err := client.LocationCustomFields(testData.OrgID, testData.LocationID)
		require.NoError(t, err)
		assert.IsType(t, map[string]any{}, data)
		assert.True(t, len(data) > 0)
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
		assert.ErrorIs(t, &ninjarmm.Error{Message: "node_maintenance_date_in_past"}, err)
	})
	t.Run("roles all", func(t *testing.T) {
		t.Parallel()
		data, err := client.Roles()
		require.NoError(t, err)
		assert.IsType(t, []ninjarmm.Role{}, data)
		assert.True(t, len(data) > 0)
	})
	t.Run("roles filter", func(t *testing.T) {
		t.Parallel()
		data, err := client.Roles("Linux Server")
		require.NoError(t, err)
		assert.IsType(t, []ninjarmm.Role{}, data)
		assert.Len(t, data, 1)
		assert.Equal(t, "LINUX_SERVER", data[0].Name)
	})
	t.Run("set role", func(t *testing.T) {
		t.Parallel()
		err := client.SetDeviceRole(testData.DeviceID, testData.RoleID)
		require.NoError(t, err)
	})
	t.Run("get role", func(t *testing.T) {
		t.Parallel()
		role, err := client.Role(testData.RoleID)
		require.NoError(t, err)
		assert.IsType(t, &ninjarmm.Role{}, role)
		assert.Equal(t, testData.RoleID, role.ID)
	})
	t.Run("get devices", func(t *testing.T) {
		t.Parallel()
		devices, err := client.Devices(nil)
		require.NoError(t, err)
		assert.True(t, len(devices) > 1_000)
	})
}

func TestClient_SoftwareInventory(t *testing.T) {
	t.Parallel()
	td, err := test.LoadTestData()
	require.NoError(t, err)
	df := ninjarmm.NewDeviceFilter().ID(ninjarmm.EQ, td.DeviceID)
	client, err := initClient()
	require.NoError(t, err)
	results, err := client.SoftwareInventory(df)
	require.NoError(t, err)
	assert.IsType(t, []ninjarmm.SoftwareInventoryResult{}, results)
	assert.True(t, len(results) > 2)
	swList := make([]string, 0, len(results))
	for _, sw := range results {
		swList = append(swList, sw.Name)
	}
	assert.Contains(t, swList, td.SoftwareName)
}

func TestClient_DevicesWithSoftware(t *testing.T) {
	td, err := test.LoadTestData()
	require.NoError(t, err)
	client, err := initClient()
	require.NoError(t, err)
	ninjarmm.DefaultQueryBatchSize = 10
	t.Run("base", func(t *testing.T) {
		t.Parallel()
		df := ninjarmm.NewDeviceFilter().Org(ninjarmm.EQ, td.OrgID).Class(ninjarmm.EQ, ninjarmm.NodeClass_WINDOWS_SERVER)
		devices, err := client.Devices(df)
		require.NoError(t, err)
		results, err := client.DevicesWithSoftware(devices, regexp.MustCompile(td.SoftwareName))
		require.NoError(t, err)
		assert.True(t, len(results) > 5, fmt.Sprintf("result=%d != expected=>%d", len(results), 5))
	})
	t.Run("single", func(t *testing.T) {
		df := ninjarmm.NewDeviceFilter().ID(ninjarmm.IN, td.DeviceID)
		devices, err := client.Devices(df)
		require.NoError(t, err)
		results, err := client.DevicesWithSoftware(devices, regexp.MustCompile(td.SoftwareName))
		require.NoError(t, err)
		assert.Len(t, results, 1)
	})

}

func TestClient_SearchDevices(t *testing.T) {
	t.Parallel()
	td, err := test.LoadTestData()
	require.NoError(t, err)
	client, err := initClient()
	require.NoError(t, err)
	df := ninjarmm.NewDeviceFilter().ID(ninjarmm.EQ, td.DeviceID)
	devices, err := client.SearchDevices(regexp.MustCompile(".+"), df)
	require.NoError(t, err)
	require.Len(t, devices, 1)
	assert.Equal(t, devices[0].ID, td.DeviceID)
}
