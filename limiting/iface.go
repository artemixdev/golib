package limiting

import "time"

type RateApprover interface {
	Approve() bool
	Setup(options Options) error
}

type RateApproverConstructor func(options Options) (RateApprover, error)

type Options struct {
	Range time.Duration
	Limit uint64
}
