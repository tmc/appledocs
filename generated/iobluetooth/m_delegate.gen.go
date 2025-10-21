// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [mDelegate] class.
var (
	MDelegateClass     _mDelegateClass
	MDelegateClassOnce sync.Once
)

func getmDelegateClass() _mDelegateClass {
	MDelegateClassOnce.Do(func() {
		MDelegateClass = _mDelegateClass{objc.GetClass("mDelegate")}
	})
	return MDelegateClass
}

type _mDelegateClass struct {
	class objc.Class
}

// An interface definition for the [mDelegate] class.
type ImDelegate interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/OBEXFileTransferServices/mDelegate
type mDelegate struct {
	objectivec.Object
}

// mDelegateFrom constructs a [mDelegate] from an unsafe.Pointer.
func mDelegateFrom(ptr unsafe.Pointer) mDelegate {
	return mDelegate{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _mDelegateClass) Alloc() mDelegate {
	rv := objc.Send[mDelegate](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _mDelegateClass) New() mDelegate {
	rv := objc.Send[mDelegate](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ mDelegate) Init() mDelegate {
	rv := objc.Send[mDelegate](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ mDelegate) Autorelease() mDelegate {
	rv := objc.Send[mDelegate](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewmDelegate creates a new mDelegate instance.
func NewmDelegate() mDelegate {
	return getmDelegateClass().New()
}




