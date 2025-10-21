// Code generated from Apple documentation for CoreTelephony. DO NOT EDIT.

package coretelephony

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [SubscriberInfo] class.
var (
	SubscriberInfoClass     _SubscriberInfoClass
	SubscriberInfoClassOnce sync.Once
)

func getSubscriberInfoClass() _SubscriberInfoClass {
	SubscriberInfoClassOnce.Do(func() {
		SubscriberInfoClass = _SubscriberInfoClass{objc.GetClass("CTSubscriberInfo")}
	})
	return SubscriberInfoClass
}

type _SubscriberInfoClass struct {
	class objc.Class
}

// An interface definition for the [SubscriberInfo] class.
type ISubscriberInfo interface {
	objectivec.IObject
}

// An object that provides an array of cellular network subscribers.
//
// Use the instances provided by this class to identify individual subscribers by their or properties.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreTelephony/CTSubscriberInfo
type SubscriberInfo struct {
	objectivec.Object
}

// SubscriberInfoFrom constructs a [SubscriberInfo] from an unsafe.Pointer.
//
// An object that provides an array of cellular network subscribers.
func SubscriberInfoFrom(ptr unsafe.Pointer) SubscriberInfo {
	return SubscriberInfo{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (sc _SubscriberInfoClass) Alloc() SubscriberInfo {
	rv := objc.Send[SubscriberInfo](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (sc _SubscriberInfoClass) New() SubscriberInfo {
	rv := objc.Send[SubscriberInfo](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SubscriberInfo) Init() SubscriberInfo {
	rv := objc.Send[SubscriberInfo](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SubscriberInfo) Autorelease() SubscriberInfo {
	rv := objc.Send[SubscriberInfo](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSubscriberInfo creates a new SubscriberInfo instance.
func NewSubscriberInfo() SubscriberInfo {
	return getSubscriberInfoClass().New()
}


// Returns the cellular network subscribers.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreTelephony/CTSubscriberInfo/subscriber()
func (sc _SubscriberInfoClass) Subscriber() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(sc.class), objc.Sel("subscriber"))
	return rv
}



