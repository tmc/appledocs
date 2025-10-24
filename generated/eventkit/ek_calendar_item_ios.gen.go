//go:build darwin && ios

// Code generated from Apple documentation for EventKit. DO NOT EDIT.

package eventkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

// iOS-only methods for EKCalendarItem


// iOS-only properties

// The calendar item’s unique identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKCalendarItem/uuid
func (e_ EKCalendarItem) UUID() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](e_.ID, objc.Sel("UUID"))
	return rv
}





