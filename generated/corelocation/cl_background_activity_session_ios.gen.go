//go:build darwin && ios

// Code generated from Apple documentation for CoreLocation. DO NOT EDIT.

package corelocation

import (
	"github.com/tmc/appledocs/generated/objc"
)

// iOS-only methods for BackgroundActivitySession

// Invalidates the background activity session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLBackgroundActivitySession-4nl4y/invalidate
func (b_ BackgroundActivitySession) Invalidate() {
	objc.Send[objc.ID](b_.ID, objc.Sel("invalidate"))
}

// iOS-only properties
