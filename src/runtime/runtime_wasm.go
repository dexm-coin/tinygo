// +build wasm,!arm,!avr

package runtime

type timeUnit int64	
const tickMicros = 1
var timestamp timeUnit

func putchar(c byte) {
}

func sleepTicks(d timeUnit) {
	timestamp += d
}

func ticks() timeUnit {
	return timestamp
}

// Align on word boundary.
func align(ptr uintptr) uintptr {
	return (ptr + 3) &^ 3
}

// Abort executes the wasm 'unreachable' instruction.
func abort() {
	trap()
}
