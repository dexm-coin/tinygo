// +build wasm,!arm,!avr

package runtime

func putchar(c byte) {
}

func sleepTicks(d timeUnit) {
}

func ticks() timeUnit {}

// Align on word boundary.
func align(ptr uintptr) uintptr {
	return (ptr + 3) &^ 3
}

// Abort executes the wasm 'unreachable' instruction.
func abort() {
	trap()
}
