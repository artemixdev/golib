package limiting

import (
	"fmt"
	"sync"
	"time"
)

type frequencyRateApprover struct {
	interval time.Duration
	mutex    sync.Mutex
	last     time.Time
}

func NewFrequencyRateApprover(options Options) (RateApprover, error) {
	ra := &frequencyRateApprover{}
	err := ra.Setup(options)
	return ra, err
}

func (ra *frequencyRateApprover) Approve() bool {
	now := time.Now()

	ra.mutex.Lock()
	defer ra.mutex.Unlock()

	if now.Sub(ra.last) < ra.interval {
		return false
	}

	ra.last = now
	return true
}

func (ra *frequencyRateApprover) Setup(options Options) error {
	if options.Range <= 0 {
		return fmt.Errorf("rate approver: range must be positive")
	}

	if options.Limit == 0 {
		return fmt.Errorf("rate approver: limit must be positive")
	}

	ra.interval = options.Range / time.Duration(options.Limit)
	return nil
}
