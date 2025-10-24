//go:build darwin && ios

// Code generated from Apple documentation for EventKit. DO NOT EDIT.

package eventkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// iOS-only methods for EKEventStore


// iOS-only properties

// The calendars associated with the event store.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKEventStore/calendars
func (e_ EKEventStore) Calendars() []IEKCalendar {
	rv := objc.Send[[]EKCalendar](e_.ID, objc.Sel("calendars"))
	return rv
}




