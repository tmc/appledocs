// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [mInactivityTimer] class.
var (
	MInactivityTimerClass     _mInactivityTimerClass
	MInactivityTimerClassOnce sync.Once
)

func getmInactivityTimerClass() _mInactivityTimerClass {
	MInactivityTimerClassOnce.Do(func() {
		MInactivityTimerClass = _mInactivityTimerClass{objc.GetClass("mInactivityTimer")}
	})
	return MInactivityTimerClass
}

type _mInactivityTimerClass struct {
	class objc.Class
}

// An interface definition for the [mInactivityTimer] class.
type ImInactivityTimer interface {
	objectivec.IObject
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/OBEXFileTransferServices/mInactivityTimer
type mInactivityTimer struct {
	objectivec.Object
}

// mInactivityTimerFrom constructs a [mInactivityTimer] from an unsafe.Pointer.
func mInactivityTimerFrom(ptr unsafe.Pointer) mInactivityTimer {
	return mInactivityTimer{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _mInactivityTimerClass) Alloc() mInactivityTimer {
	rv := objc.Send[mInactivityTimer](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _mInactivityTimerClass) New() mInactivityTimer {
	rv := objc.Send[mInactivityTimer](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ mInactivityTimer) Init() mInactivityTimer {
	rv := objc.Send[mInactivityTimer](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ mInactivityTimer) Autorelease() mInactivityTimer {
	rv := objc.Send[mInactivityTimer](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewmInactivityTimer creates a new mInactivityTimer instance.
func NewmInactivityTimer() mInactivityTimer {
	return getmInactivityTimerClass().New()
}




