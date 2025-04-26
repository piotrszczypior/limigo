package limiter

import (
	"time"
)

func CreateTokenBucket(capacity int, refillRate time.Duration) Limiter {
	return &tokenBucket{
		maxCapacity: capacity,
		tokens:      capacity,
		refillRate:  refillRate,
	}
}

func (bucket *tokenBucket) Allow() bool {
	return shouldLimit(bucket)
}

type tokenBucket struct {
	maxCapacity int
	tokens      int
	refillRate  time.Duration
	lastRefill  time.Time
}

func shouldLimit(bucket *tokenBucket) bool {
	elapsedTime := time.Since(bucket.lastRefill)

	tokensToAdd := int(elapsedTime / bucket.refillRate)

	if tokensToAdd > 0 {
		bucketContent := bucket.tokens + tokensToAdd
		bucket.tokens = min(bucketContent, bucket.maxCapacity)
		bucket.lastRefill = time.Now()
	}

	if bucket.tokens > 0 {
		bucket.tokens--
		return true
	}
	return false
}
