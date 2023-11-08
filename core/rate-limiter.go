package core

type RateLimiter interface {
	AllowRequest(ipAddr string) bool
}

const (
	TokenBucketLimiterType = iota
)

func NewRateLimiter(rateLimiterType int) RateLimiter {
	switch rateLimiterType {
	case TokenBucketLimiterType:
		return NewTokenBucketLimiter()
	default:
		return NewTokenBucketLimiter()
	}
}
