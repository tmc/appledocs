//go:build darwin && ios

// Code generated from Apple documentation for EventKit. DO NOT EDIT.

package eventkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

// iOS-only methods for EKSource


// iOS-only properties

// The calendars that belong to this source object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKSource/calendars
func (e_ EKSource) Calendars() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](e_.ID, objc.Sel("calendars"))
	return rv
}





