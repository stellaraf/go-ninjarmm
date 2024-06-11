package ninjarmm_test

import (
	"testing"

	"github.com/stellaraf/go-ninjarmm"
	"github.com/stellaraf/go-ninjarmm/internal/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_Query(t *testing.T) {
	t.Run("full", func(t *testing.T) {
		client, err := initClient()
		require.NoError(t, err)
		testData, err := test.LoadTestData()
		require.NoError(t, err)
		qc := ninjarmm.NewQueryClient[ninjarmm.SoftwareInventoryResult](client, 100)
		df := ninjarmm.NewDeviceFilter().ID(ninjarmm.EQ, testData.DeviceID)
		q := map[string]string{"df": df.String()}
		results, err := qc.Do("/api/v2/queries/software", q)
		require.NoError(t, err)
		assert.IsType(t, []ninjarmm.SoftwareInventoryResult{}, results.Results)
		assert.True(t, len(results.Results) > 2)
		swList := make([]string, 0, len(results.Results))
		for _, sw := range results.Results {
			swList = append(swList, sw.Name)
		}
		assert.Contains(t, swList, testData.SoftwareName)
	})

}
