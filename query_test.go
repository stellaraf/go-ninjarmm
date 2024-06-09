package ninjarmm_test

import (
	"fmt"
	"testing"

	"github.com/stellaraf/go-ninjarmm"
	"github.com/stellaraf/go-ninjarmm/internal/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func run(t *testing.T, size int) {
	client, err := initClient()
	require.NoError(t, err)
	testData, err := test.LoadTestData()
	require.NoError(t, err)
	qc := ninjarmm.NewQueryClient[ninjarmm.SoftwareInventoryResult](client, 100)
	df := ninjarmm.NewDeviceFilter().ID(ninjarmm.EQ, testData.DeviceID)
	q := map[string]string{"df": df.String(), "pageSize": fmt.Sprint(size)}
	results, err := qc.Do("/api/v2/queries/software", q)
	require.NoError(t, err)
	assert.IsType(t, []ninjarmm.SoftwareInventoryResult{}, results)
	assert.True(t, len(results) > 2)
	swList := make([]string, 0, len(results))
	for _, sw := range results {
		swList = append(swList, sw.Name)
	}
	assert.Contains(t, swList, testData.SoftwareName)
}

func Test_Query(t *testing.T) {
	t.Run("with page", func(t *testing.T) {
		t.Parallel()
		run(t, 5)
	})
	t.Run("full", func(t *testing.T) {
		t.Parallel()
		run(t, 500)
	})
}
