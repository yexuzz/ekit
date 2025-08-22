package retry

import (
	"math"
	"sync/atomic"
	"time"
)

var _ Strategy = (*ExponentialBackoffRetryStrategy)(nil)

// ExponentialBackoffRetryStrategy 指数退避重试
type ExponentialBackoffRetryStrategy struct {
	initialInterval    time.Duration //初始重试间隔
	maxInterval        time.Duration //最大重试间隔
	maxRetries         int32         //最大重试次数
	retries            int32         //当前重试次数
	maxIntervalReached atomic.Value  //是否已经达到最大重试间隔
}

func NewExponentialBackoffRetryStrategy(initialInterval, maxInterval time.Duration, maxRetires int32) (*ExponentialBackoffRetryStrategy, error) {
	if initialInterval <= 0 {
		return nil, NewErrInvalidIntervalValue(initialInterval)
	}
	if initialInterval > maxInterval {
		return nil, NewErrInvalidMaxIntervalValue(maxInterval, initialInterval)
	}

	return &ExponentialBackoffRetryStrategy{
		initialInterval: initialInterval,
		maxInterval:     maxInterval,
		maxRetries:      maxRetires,
	}, nil
}

func (s *ExponentialBackoffRetryStrategy) Report(err error) Strategy {
	return s
}
func (s *ExponentialBackoffRetryStrategy) Next() (time.Duration, bool) {
	retries := atomic.AddInt32(&s.retries, 1)
	if s.maxRetries <= 0 || retries < s.maxRetries {
		if reached, ok := s.maxIntervalReached.Load().(bool); ok && reached {
			return s.maxInterval, true
		}
		interval := s.initialInterval * time.Duration(math.Pow(2, float64(retries-1)))
		//溢出或当前重试间隔大于最大重试间隔
		if interval <= 0 || interval > s.maxInterval {
			s.maxIntervalReached.Store(true)
			return s.maxInterval, true
		}
		return interval, true
	}
	return 0, false
}
