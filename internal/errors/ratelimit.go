package errors

import (
	"fmt"
	"time"
)

type RateLimitError struct {
	resetTime time.Time
}

func NewRateLimitError(resetTime time.Time) error {
	return &RateLimitError{resetTime: resetTime}
}
func (e *RateLimitError) Error() string {
	return fmt.Sprintf("API rate limit exceeded. Reset Time: %s", e.resetTime)
}
