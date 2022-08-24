package limiting

import (
	"fmt"
	"sync"
)

type RateApproverVec[K comparable] struct {
	constructor RateApproverConstructor
	options     Options
	mutex       sync.Mutex
	approvers   map[K]RateApprover
}

func NewRateApproverVec[K comparable](
	constructor RateApproverConstructor, options Options) *RateApproverVec[K] {

	return &RateApproverVec[K]{
		constructor: constructor,
		options:     options,
		approvers:   make(map[K]RateApprover),
	}
}

func (rav *RateApproverVec[K]) With(key K) (RateApprover, error) {
	rav.mutex.Lock()
	defer rav.mutex.Unlock()

	if approver, found := rav.approvers[key]; found {
		return approver, nil
	}

	approver, err := rav.constructor(rav.options)
	if err != nil {
		return nil, fmt.Errorf("key %+v: %w", key, err)
	}

	rav.approvers[key] = approver
	return approver, nil
}

func (rav *RateApproverVec[K]) Setup(options Options, overrides map[K]Options) error {
	rav.mutex.Lock()
	defer rav.mutex.Unlock()

	rav.options = options

	for key, approver := range rav.approvers {
		value := options

		override, found := overrides[key]
		if found {
			value = override
		}

		if err := approver.Setup(value); err != nil {
			return fmt.Errorf("key %+v: %w", key, err)
		}
	}

	return nil
}
