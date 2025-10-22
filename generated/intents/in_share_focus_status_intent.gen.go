// Code generated from Apple documentation for Intents. DO NOT EDIT.

package intents

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [INShareFocusStatusIntent] class.
var (
	INShareFocusStatusIntentClass     _INShareFocusStatusIntentClass
	INShareFocusStatusIntentClassOnce sync.Once
)

func getINShareFocusStatusIntentClass() _INShareFocusStatusIntentClass {
	INShareFocusStatusIntentClassOnce.Do(func() {
		INShareFocusStatusIntentClass = _INShareFocusStatusIntentClass{objc.GetClass("INShareFocusStatusIntent")}
	})
	return INShareFocusStatusIntentClass
}

type _INShareFocusStatusIntentClass struct {
	class objc.Class
}

// An interface definition for the [INShareFocusStatusIntent] class.
type IINShareFocusStatusIntent interface {
	IINIntent
	FocusStatus() INFocusStatus
}

// An object that indicates the user’s focus status is changing.
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INShareFocusStatusIntent
type INShareFocusStatusIntent struct {
	INIntent
}

// INShareFocusStatusIntentFrom constructs a [INShareFocusStatusIntent] from an unsafe.Pointer.
//
// An object that indicates the user’s focus status is changing.
func INShareFocusStatusIntentFrom(ptr unsafe.Pointer) INShareFocusStatusIntent {
	return INShareFocusStatusIntent{
		INIntent: INIntentFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ic _INShareFocusStatusIntentClass) Alloc() INShareFocusStatusIntent {
	rv := objc.Send[INShareFocusStatusIntent](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ic _INShareFocusStatusIntentClass) New() INShareFocusStatusIntent {
	rv := objc.Send[INShareFocusStatusIntent](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ INShareFocusStatusIntent) Init() INShareFocusStatusIntent {
	rv := objc.Send[INShareFocusStatusIntent](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ INShareFocusStatusIntent) Autorelease() INShareFocusStatusIntent {
	rv := objc.Send[INShareFocusStatusIntent](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewINShareFocusStatusIntent creates a new INShareFocusStatusIntent instance.
func NewINShareFocusStatusIntent() INShareFocusStatusIntent {
	return getINShareFocusStatusIntentClass().New()
}




// Creates an intent with the specified focus status.
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INShareFocusStatusIntent/init(focusStatus:)
func NewINShareFocusStatusIntentWithFocusStatus(focusStatus INFocusStatus) INShareFocusStatusIntent {
	instance := getINShareFocusStatusIntentClass().Alloc()
	rv := objc.Send[INShareFocusStatusIntent](instance.ID, objc.Sel("initWithFocusStatus:"), focusStatus)
	rv.Autorelease()
	return rv
}


// The user’s preference for receiving communication notifications.
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INShareFocusStatusIntent/focusStatus
func (i_ INShareFocusStatusIntent) FocusStatus() INFocusStatus {
	rv := objc.Send[INFocusStatus](i_.ID, objc.Sel("focusStatus"))
	return rv
}


