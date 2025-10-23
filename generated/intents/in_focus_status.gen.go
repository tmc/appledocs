// Code generated from Apple documentation for Intents. DO NOT EDIT.

package intents

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [INFocusStatus] class.
var (
	INFocusStatusClass     _INFocusStatusClass
	INFocusStatusClassOnce sync.Once
)

func getINFocusStatusClass() _INFocusStatusClass {
	INFocusStatusClassOnce.Do(func() {
		INFocusStatusClass = _INFocusStatusClass{objc.GetClass("INFocusStatus")}
	})
	return INFocusStatusClass
}

type _INFocusStatusClass struct {
	class objc.Class
}

// An interface definition for the [INFocusStatus] class.
type IINFocusStatus interface {
	objectivec.IObject
	IsFocused() foundation.Number
}

// The user’s preference for receiving notifications.
//
// When a user wants to focus, they can choose to prevent or delay notifications from most apps. Use this information to display that the user is in a focus to other people in your communication service.


// The user’s preference for receiving notifications.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Intents/INFocusStatus
type INFocusStatus struct {
	objectivec.Object
}

// INFocusStatusFrom constructs a [INFocusStatus] from an unsafe.Pointer.
//
// The user’s preference for receiving notifications.
func INFocusStatusFrom(ptr unsafe.Pointer) INFocusStatus {
	return INFocusStatus{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ic _INFocusStatusClass) Alloc() INFocusStatus {
	rv := objc.Send[INFocusStatus](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ic _INFocusStatusClass) New() INFocusStatus {
	rv := objc.Send[INFocusStatus](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ INFocusStatus) Init() INFocusStatus {
	rv := objc.Send[INFocusStatus](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ INFocusStatus) Autorelease() INFocusStatus {
	rv := objc.Send[INFocusStatus](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewINFocusStatus creates a new INFocusStatus instance.
func NewINFocusStatus() INFocusStatus {
	return getINFocusStatusClass().New()
}



// Creates an object that indicates the user’s ability to receive communication notifications.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Intents/INFocusStatus/initWithIsFocused:
func NewINFocusStatusWithIsFocused(isFocused foundation.INumber) INFocusStatus {
	instance := getINFocusStatusClass().Alloc()
	rv := objc.Send[INFocusStatus](instance.ID, objc.Sel("initWithIsFocused:"), isFocused)
	rv.Autorelease()
	return rv
}



// The user’s preference for receiving communication notifications.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Intents/INFocusStatus/isFocused-78wbx
func (i_ INFocusStatus) IsFocused() foundation.Number {
	rv := objc.Send[foundation.Number](i_.ID, objc.Sel("isFocused"))
	return rv
}


