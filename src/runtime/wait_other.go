//go:build !tinygo.riscv && !cortexm && !(linux && !baremetal && !tinygo.wasm) && !darwin

package runtime

func waitForEvents() {
	runtimePanic("deadlocked: no event source")
}
