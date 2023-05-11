package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func initClient() (client *NinjaRMMClient, err error) {
	env, err := LoadEnv()
	if err != nil {
		return
	}
	getAccessToken, setAccessToken, getRefreshToken, setRefreshToken, err := setup()
	if err != nil {
		return
	}
	client, err = CreateNinjaRMMClient(
		env.BaseURL,
		env.ClientID,
		env.ClientSecret,
		&env.EncryptionPassphrase,
		getAccessToken,
		setAccessToken,
		getRefreshToken,
		setRefreshToken,
	)
	return
}

func Test_NinjaRMMClient(t *testing.T) {
	testData, err := loadTestData()
	assert.NoError(t, err)

	t.Run("organizations", func(t *testing.T) {
		client, err := initClient()
		assert.NoError(t, err)
		assert.NotNil(t, client)
		orgs, err := client.Organizations()
		assert.NoError(t, err)
		assert.IsType(t, []OrganizationSummary{}, orgs)
	})
	t.Run("device", func(t *testing.T) {
		client, err := initClient()
		assert.NoError(t, err)
		assert.NotNil(t, client)
		data, err := client.Device(testData.DeviceID)
		assert.NoError(t, err)
		assert.IsType(t, DeviceDetails{}, data)
	})
	t.Run("organization", func(t *testing.T) {
		client, err := initClient()
		assert.NoError(t, err)
		assert.NotNil(t, client)
		data, err := client.Organization(testData.OrgID)
		assert.NoError(t, err)
		assert.IsType(t, Organization{}, data)
	})

	t.Run("os patches", func(t *testing.T) {
		client, err := initClient()
		assert.NoError(t, err)
		assert.NotNil(t, client)
		data, err := client.OSPatches(testData.OrgID)
		assert.NoError(t, err)
		assert.IsType(t, OSPatchReportQuery{}, data)
	})
	t.Run("os patch report", func(t *testing.T) {
		client, err := initClient()
		assert.NoError(t, err)
		assert.NotNil(t, client)
		data, err := client.OSPatchReport(testData.OrgID)
		assert.NoError(t, err)
		assert.IsType(t, []OSPatchReportDetail{}, data)
	})
}
