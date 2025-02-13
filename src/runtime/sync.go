// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package runtime

import (
	"internal/task"
	"unsafe"
)

// This file contains stub implementations for internal/poll.
// The official golang implementation states:
//
// "That is, don't think of these as semaphores.
// Think of them as a way to implement sleep and wakeup
// such that every sleep is paired with a single wakeup,
// even if, due to races, the wakeup happens before the sleep."
//
// This is an experimental and probably incomplete implementation of the
// semaphore system, tailored to the network use case. That means, that it does not
// implement the modularity that the semacquire/semacquire1 implementation model
// offers, which in fact is emitted here entirely.
// This means we assume the following constant settings from the golang standard
// library: lifo=false,profile=semaBlock,skipframe=0,reason=waitReasonSemaquire

// TODO: is this per-process or per-thread?
// Are there potential collisions for the addresses?
var semAddrMap map[uint32]task.Semaphore

func init() {
	semAddrMap = make(map[uint32]task.Semaphore)
}

func printSemMap() {
	println("[dbg] semAddrMap:")
	for k, _ := range semAddrMap {
		println("[dbg]", k)
	}
}

// Go runtime expects this function to receive a pointer to a uint32
// Tinygo has it's own implementation of semaphores though, so we map
// the raw address to a semaphore object and call the Wait method.
//
//go:linkname semacquire internal/poll.runtime_Semacquire
func semacquire(sema *uint32) {
	skip := false
	if !skip {
		if sema == nil {
			panic("semacquire: sema==0")
		}

		printSemMap()
		// get absolute uint32 value of the semaphore for map index
		semAddru32 := uint32(uintptr(unsafe.Pointer(sema)))
		println("[dbg] semAddru32: ", semAddru32)

		// check if there already is a semaphore object for this address
		// and call the Wait method
		if semaphore, ok := semAddrMap[semAddru32]; ok {
			println("[dbg] acquire semaphore: ", sema)
			semaphore.Wait()
		} else {
			// if there is no semaphore object, create one and add it to the map
			semaphore := *task.NewSemaphore(semAddru32)
			semAddrMap[semAddru32] = semaphore
			println("[dbg] add semaphore: ", sema)
			printSemMap()
			semaphore.Wait()
		}
	}
}

//go:linkname semrelease internal/poll.runtime_Semrelease
func semrelease(sema *uint32) {
	skip := false
	if !skip {
		if sema == nil {
			panic("semrelease: sema==0")
		}
		printSemMap()
		semAddru32 := uint32(uintptr(unsafe.Pointer(sema)))
		println("[dbg] semAddru32: ", semAddru32)

		if semaphore, ok := semAddrMap[semAddru32]; ok {
			println("[dbg] release semaphore: ", sema)
			semaphore.Post()
		} else {
			println("[dbg] semaphore release (unknown): ", sema)
			// semacquire(sema)
			// semrelease(sema)
		}
	}
}
