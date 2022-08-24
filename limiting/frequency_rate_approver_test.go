package limiting

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestFrequencyRateApprover(t *testing.T) {
	ra, err := NewFrequencyRateApprover(Options{Range: 3 * time.Second, Limit: 6})
	require.NoError(t, err)

	require.True(t, ra.Approve())
	time.Sleep(50 * time.Millisecond)
	require.False(t, ra.Approve())

	time.Sleep(500 * time.Millisecond)

	require.True(t, ra.Approve())
	time.Sleep(50 * time.Millisecond)
	require.False(t, ra.Approve())
}
