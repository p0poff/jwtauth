package params

import (
	"time"
)

type Init struct {
	Expire time.Duration
	TokenKey string
}

