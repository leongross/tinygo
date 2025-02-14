//go:build scheduler.threads

package task

import "internal/futex"

type Futex = futex.Futex

func NewFutex(runtimeAddr *uint32) Futex {
	return futex.NewFutex(runtimeAddr)
}
