// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [mOpenConnectionSelectorTarget] class.
var (
	MOpenConnectionSelectorTargetClass     _mOpenConnectionSelectorTargetClass
	MOpenConnectionSelectorTargetClassOnce sync.Once
)

func getmOpenConnectionSelectorTargetClass() _mOpenConnectionSelectorTargetClass {
	MOpenConnectionSelectorTargetClassOnce.Do(func() {
		MOpenConnectionSelectorTargetClass = _mOpenConnectionSelectorTargetClass{objc.GetClass("mOpenConnectionSelectorTarget")}
	})
	return MOpenConnectionSelectorTargetClass
}

type _mOpenConnectionSelectorTargetClass struct {
	class objc.Class
}

// An interface definition for the [mOpenConnectionSelectorTarget] class.
type ImOpenConnectionSelectorTarget interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothOBEXSession/mOpenConnectionSelectorTarget
type mOpenConnectionSelectorTarget struct {
	objectivec.Object
}

// mOpenConnectionSelectorTargetFrom constructs a [mOpenConnectionSelectorTarget] from an unsafe.Pointer.
func mOpenConnectionSelectorTargetFrom(ptr unsafe.Pointer) mOpenConnectionSelectorTarget {
	return mOpenConnectionSelectorTarget{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _mOpenConnectionSelectorTargetClass) Alloc() mOpenConnectionSelectorTarget {
	rv := objc.Send[mOpenConnectionSelectorTarget](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _mOpenConnectionSelectorTargetClass) New() mOpenConnectionSelectorTarget {
	rv := objc.Send[mOpenConnectionSelectorTarget](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ mOpenConnectionSelectorTarget) Init() mOpenConnectionSelectorTarget {
	rv := objc.Send[mOpenConnectionSelectorTarget](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ mOpenConnectionSelectorTarget) Autorelease() mOpenConnectionSelectorTarget {
	rv := objc.Send[mOpenConnectionSelectorTarget](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewmOpenConnectionSelectorTarget creates a new mOpenConnectionSelectorTarget instance.
func NewmOpenConnectionSelectorTarget() mOpenConnectionSelectorTarget {
	return getmOpenConnectionSelectorTargetClass().New()
}




