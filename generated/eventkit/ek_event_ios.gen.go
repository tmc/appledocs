//go:build darwin && ios

// Code generated from Apple documentation for EventKit. DO NOT EDIT.

package eventkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

// iOS-only methods for EKEvent


// iOS-only properties

// The Address Book framework record identifier of the person for this birthday event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKEvent/birthdayPersonID
func (e_ EKEvent) BirthdayPersonID() int {
	rv := objc.Send[int](e_.ID, objc.Sel("birthdayPersonID"))
	return rv
}




