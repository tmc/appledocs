// Code generated from Apple documentation for Intents. DO NOT EDIT.

package intents

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [INFocusStatusCenter] class.
var (
	INFocusStatusCenterClass     _INFocusStatusCenterClass
	INFocusStatusCenterClassOnce sync.Once
)

func getINFocusStatusCenterClass() _INFocusStatusCenterClass {
	INFocusStatusCenterClassOnce.Do(func() {
		INFocusStatusCenterClass = _INFocusStatusCenterClass{objc.GetClass("INFocusStatusCenter")}
	})
	return INFocusStatusCenterClass
}

type _INFocusStatusCenterClass struct {
	class objc.Class
}

// An interface definition for the [INFocusStatusCenter] class.
type IINFocusStatusCenter interface {
	objectivec.IObject
	RequestAuthorizationWithCompletionHandler(completionHandler unsafe.Pointer)
	AuthorizationStatus() INFocusStatusAuthorizationStatus
	FocusStatus() INFocusStatus
}

// An object that maintains the user’s current focus status and your app’s ability to access it.


// An object that maintains the user’s current focus status and your app’s ability to access it.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Intents/INFocusStatusCenter
type INFocusStatusCenter struct {
	objectivec.Object
}

// INFocusStatusCenterFrom constructs a [INFocusStatusCenter] from an unsafe.Pointer.
//
// An object that maintains the user’s current focus status and your app’s ability to access it.
func INFocusStatusCenterFrom(ptr unsafe.Pointer) INFocusStatusCenter {
	return INFocusStatusCenter{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ic _INFocusStatusCenterClass) Alloc() INFocusStatusCenter {
	rv := objc.Send[INFocusStatusCenter](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ic _INFocusStatusCenterClass) New() INFocusStatusCenter {
	rv := objc.Send[INFocusStatusCenter](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ INFocusStatusCenter) Init() INFocusStatusCenter {
	rv := objc.Send[INFocusStatusCenter](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ INFocusStatusCenter) Autorelease() INFocusStatusCenter {
	rv := objc.Send[INFocusStatusCenter](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewINFocusStatusCenter creates a new INFocusStatusCenter instance.
func NewINFocusStatusCenter() INFocusStatusCenter {
	return getINFocusStatusCenterClass().New()
}



// The systemwide object that holds the user’s focus status.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Intents/INFocusStatusCenter/default
func (ic _INFocusStatusCenterClass) DefaultCenter() INFocusStatusCenter {
	rv := objc.Send[INFocusStatusCenter](objc.ID(ic.class), objc.Sel("defaultCenter"))
	return rv
}

// Asks the system for access to the user’s focus status.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Intents/INFocusStatusCenter/requestAuthorization(completionHandler:)
func (i_ INFocusStatusCenter) RequestAuthorizationWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("requestAuthorizationWithCompletionHandler:"), completionHandler)
}


// Returns your app’s current ability to access the user’s focus status.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Intents/INFocusStatusCenter/authorizationStatus
func (i_ INFocusStatusCenter) AuthorizationStatus() INFocusStatusAuthorizationStatus {
	rv := objc.Send[INFocusStatusAuthorizationStatus](i_.ID, objc.Sel("authorizationStatus"))
	return rv
}


// The systemwide object that holds the user’s focus status.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Intents/INFocusStatusCenter/default
func (i_ INFocusStatusCenter) DefaultCenter() INFocusStatusCenter {
	rv := objc.Send[INFocusStatusCenter](i_.ID, objc.Sel("defaultCenter"))
	return rv
}


// The user’s ability to receive notifications.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Intents/INFocusStatusCenter/focusStatus
func (i_ INFocusStatusCenter) FocusStatus() INFocusStatus {
	rv := objc.Send[INFocusStatus](i_.ID, objc.Sel("focusStatus"))
	return rv
}



