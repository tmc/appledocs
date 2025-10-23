// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [mIONotification] class.
var (
	MIONotificationClass     _mIONotificationClass
	MIONotificationClassOnce sync.Once
)

func getmIONotificationClass() _mIONotificationClass {
	MIONotificationClassOnce.Do(func() {
		MIONotificationClass = _mIONotificationClass{objc.GetClass("mIONotification")}
	})
	return MIONotificationClass
}

type _mIONotificationClass struct {
	class objc.Class
}

// An interface definition for the [mIONotification] class.
type ImIONotification interface {
	objectivec.IObject
	// properties:
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothObject/mIONotification
type mIONotification struct {
	objectivec.Object
}

// mIONotificationFrom constructs a [mIONotification] from an unsafe.Pointer.
func mIONotificationFrom(ptr unsafe.Pointer) mIONotification {
	return mIONotification{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _mIONotificationClass) Alloc() mIONotification {
	rv := objc.Send[mIONotification](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _mIONotificationClass) New() mIONotification {
	rv := objc.Send[mIONotification](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ mIONotification) Init() mIONotification {
	rv := objc.Send[mIONotification](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ mIONotification) Autorelease() mIONotification {
	rv := objc.Send[mIONotification](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewmIONotification creates a new mIONotification instance.
func NewmIONotification() mIONotification {
	return getmIONotificationClass().New()
}




