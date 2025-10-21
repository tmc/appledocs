// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [mEventDataListener] class.
var (
	MEventDataListenerClass     _mEventDataListenerClass
	MEventDataListenerClassOnce sync.Once
)

func getmEventDataListenerClass() _mEventDataListenerClass {
	MEventDataListenerClassOnce.Do(func() {
		MEventDataListenerClass = _mEventDataListenerClass{objc.GetClass("mEventDataListener")}
	})
	return MEventDataListenerClass
}

type _mEventDataListenerClass struct {
	class objc.Class
}

// An interface definition for the [mEventDataListener] class.
type ImEventDataListener interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothRFCOMMChannel/mEventDataListener
type mEventDataListener struct {
	objectivec.Object
}

// mEventDataListenerFrom constructs a [mEventDataListener] from an unsafe.Pointer.
func mEventDataListenerFrom(ptr unsafe.Pointer) mEventDataListener {
	return mEventDataListener{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _mEventDataListenerClass) Alloc() mEventDataListener {
	rv := objc.Send[mEventDataListener](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _mEventDataListenerClass) New() mEventDataListener {
	rv := objc.Send[mEventDataListener](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ mEventDataListener) Init() mEventDataListener {
	rv := objc.Send[mEventDataListener](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ mEventDataListener) Autorelease() mEventDataListener {
	rv := objc.Send[mEventDataListener](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewmEventDataListener creates a new mEventDataListener instance.
func NewmEventDataListener() mEventDataListener {
	return getmEventDataListenerClass().New()
}




