// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [mOpenConnectionCallback] class.
var (
	MOpenConnectionCallbackClass     _mOpenConnectionCallbackClass
	MOpenConnectionCallbackClassOnce sync.Once
)

func getmOpenConnectionCallbackClass() _mOpenConnectionCallbackClass {
	MOpenConnectionCallbackClassOnce.Do(func() {
		MOpenConnectionCallbackClass = _mOpenConnectionCallbackClass{objc.GetClass("mOpenConnectionCallback")}
	})
	return MOpenConnectionCallbackClass
}

type _mOpenConnectionCallbackClass struct {
	class objc.Class
}

// An interface definition for the [mOpenConnectionCallback] class.
type ImOpenConnectionCallback interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothOBEXSession/mOpenConnectionCallback
type mOpenConnectionCallback struct {
	objectivec.Object
}

// mOpenConnectionCallbackFrom constructs a [mOpenConnectionCallback] from an unsafe.Pointer.
func mOpenConnectionCallbackFrom(ptr unsafe.Pointer) mOpenConnectionCallback {
	return mOpenConnectionCallback{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _mOpenConnectionCallbackClass) Alloc() mOpenConnectionCallback {
	rv := objc.Send[mOpenConnectionCallback](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _mOpenConnectionCallbackClass) New() mOpenConnectionCallback {
	rv := objc.Send[mOpenConnectionCallback](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ mOpenConnectionCallback) Init() mOpenConnectionCallback {
	rv := objc.Send[mOpenConnectionCallback](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ mOpenConnectionCallback) Autorelease() mOpenConnectionCallback {
	rv := objc.Send[mOpenConnectionCallback](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewmOpenConnectionCallback creates a new mOpenConnectionCallback instance.
func NewmOpenConnectionCallback() mOpenConnectionCallback {
	return getmOpenConnectionCallbackClass().New()
}




