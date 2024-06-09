package util_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/stellaraf/go-ninjarmm/internal/util"
	"github.com/stretchr/testify/assert"
)

func Test_TimeToFractional(t *testing.T) {
	t.Run("time to fractional", func(t *testing.T) {
		ts := time.Date(
			2023,      // year
			7,         // month
			18,        // day
			1,         // hour
			23,        // minute
			45,        // second
			670000076, // millisecond
			time.UTC,  // timezone
		)
		expected := 1689643425.670000076
		result := util.TimeToFractional(ts)
		fs := "%.9f"
		assert.Equal(t, fmt.Sprintf(fs, expected), fmt.Sprintf(fs, result))
	})
}

func Test_MatchWithUpper(t *testing.T) {
	t.Run("exact", func(t *testing.T) {
		t.Parallel()
		assert.True(t, util.MatchWithUpper("LINUX_SERVER", "LINUX_SERVER"))
	})
	t.Run("partial", func(t *testing.T) {
		t.Parallel()
		assert.True(t, util.MatchWithUpper("LINUX_SERVER", "Linux Server"))
		assert.False(t, util.MatchWithUpper("WINDOWS_SERVER", "Linux Server"))
	})
}
