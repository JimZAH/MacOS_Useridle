package macos_idle

// #cgo LDFLAGS: -framework CoreGraphics
// #include <CoreGraphics/CoreGraphics.h>
import "C"

import (
	"math"
)

func Check() uint {
	return uint(C.CGEventSourceSecondsSinceLastEventType(1, math.MaxInt32))
}
