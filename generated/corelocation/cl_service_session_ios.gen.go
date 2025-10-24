//go:build darwin && ios

// Code generated from Apple documentation for CoreLocation. DO NOT EDIT.

package corelocation

import (
	"github.com/tmc/appledocs/generated/objc"
)

// iOS-only methods for ServiceSession

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLServiceSession-2ddhd/invalidate
func (s_ ServiceSession) Invalidate() {
	objc.Send[objc.ID](s_.ID, objc.Sel("invalidate"))
}

// iOS-only properties
