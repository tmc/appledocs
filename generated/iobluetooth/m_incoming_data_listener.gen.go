// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [mIncomingDataListener] class.
var (
	MIncomingDataListenerClass     _mIncomingDataListenerClass
	MIncomingDataListenerClassOnce sync.Once
)

func getmIncomingDataListenerClass() _mIncomingDataListenerClass {
	MIncomingDataListenerClassOnce.Do(func() {
		MIncomingDataListenerClass = _mIncomingDataListenerClass{objc.GetClass("mIncomingDataListener")}
	})
	return MIncomingDataListenerClass
}

type _mIncomingDataListenerClass struct {
	class objc.Class
}

// An interface definition for the [mIncomingDataListener] class.
type ImIncomingDataListener interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothRFCOMMChannel/mIncomingDataListener
type mIncomingDataListener struct {
	objectivec.Object
}

// mIncomingDataListenerFrom constructs a [mIncomingDataListener] from an unsafe.Pointer.
func mIncomingDataListenerFrom(ptr unsafe.Pointer) mIncomingDataListener {
	return mIncomingDataListener{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _mIncomingDataListenerClass) Alloc() mIncomingDataListener {
	rv := objc.Send[mIncomingDataListener](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _mIncomingDataListenerClass) New() mIncomingDataListener {
	rv := objc.Send[mIncomingDataListener](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ mIncomingDataListener) Init() mIncomingDataListener {
	rv := objc.Send[mIncomingDataListener](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ mIncomingDataListener) Autorelease() mIncomingDataListener {
	rv := objc.Send[mIncomingDataListener](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewmIncomingDataListener creates a new mIncomingDataListener instance.
func NewmIncomingDataListener() mIncomingDataListener {
	return getmIncomingDataListenerClass().New()
}




