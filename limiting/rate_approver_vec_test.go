package limiting

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestRateApproverVec(t *testing.T) {
	options := Options{Range: 500 * time.Millisecond, Limit: 4}
	rav := NewRateApproverVec[string](NewFrequencyRateApprover, options)

	ra, err := rav.With("hello")
	require.NoError(t, err)

	ra2, err := rav.With("hello")
	require.NoError(t, err)

	ra3, err := rav.With("world")
	require.NoError(t, err)

	require.True(t, ra == ra2)
	require.False(t, ra == ra3)
	require.False(t, ra2 == ra3)
}
