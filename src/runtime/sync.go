package runtime

import (
	"internal/task"
	"unsafe"
)

// This file contains stub implementations for internal/poll.

// The golang runtime exposes the function semacquire to acquire a semaphore and semrelease to release a semaphore.
// These functions are used by the runtime to implement sync.Mutex and sync.Cond. The runtime uses a futex-based
// implementation of these functions. The futex-based implementation is not available in the TinyGo runtime. This

// The semaphore is uniquely identified by the address of the uint32 variable that is passed to semacquire.
// Hence we have to implement a global table (similar to golang) that maps the address of the uint32 variable to
// and instance of a semaphore implemented in internal/task.
// It is assumed that the address of the uint32 variable is unique and does not change during the lifetime of the
// program.

var semMap = make(map[uint32]task.Semaphore)

//go:linkname semacquire internal/poll.runtime_Semacquire
func semacquire(sema *uint32) {
	if sem, ok := semMap[uint32(uintptr(unsafe.Pointer(sema)))]; ok {
		sem.Wait()
	}
	// TODO: set error
}

//go:linkname semrelease internal/poll.runtime_Semrelease
func semrelease(sema *uint32) {
	if sem, ok := semMap[uint32(uintptr(unsafe.Pointer(sema)))]; ok {
		sem.Post()
	}
	// TODO: set error
}
