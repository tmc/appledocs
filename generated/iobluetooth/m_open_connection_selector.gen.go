// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [mOpenConnectionSelector] class.
var (
	MOpenConnectionSelectorClass     _mOpenConnectionSelectorClass
	MOpenConnectionSelectorClassOnce sync.Once
)

func getmOpenConnectionSelectorClass() _mOpenConnectionSelectorClass {
	MOpenConnectionSelectorClassOnce.Do(func() {
		MOpenConnectionSelectorClass = _mOpenConnectionSelectorClass{objc.GetClass("mOpenConnectionSelector")}
	})
	return MOpenConnectionSelectorClass
}

type _mOpenConnectionSelectorClass struct {
	class objc.Class
}

// An interface definition for the [mOpenConnectionSelector] class.
type ImOpenConnectionSelector interface {
	objectivec.IObject
	// properties:
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothOBEXSession/mOpenConnectionSelector
type mOpenConnectionSelector struct {
	objectivec.Object
}

// mOpenConnectionSelectorFrom constructs a [mOpenConnectionSelector] from an unsafe.Pointer.
func mOpenConnectionSelectorFrom(ptr unsafe.Pointer) mOpenConnectionSelector {
	return mOpenConnectionSelector{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _mOpenConnectionSelectorClass) Alloc() mOpenConnectionSelector {
	rv := objc.Send[mOpenConnectionSelector](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _mOpenConnectionSelectorClass) New() mOpenConnectionSelector {
	rv := objc.Send[mOpenConnectionSelector](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ mOpenConnectionSelector) Init() mOpenConnectionSelector {
	rv := objc.Send[mOpenConnectionSelector](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ mOpenConnectionSelector) Autorelease() mOpenConnectionSelector {
	rv := objc.Send[mOpenConnectionSelector](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewmOpenConnectionSelector creates a new mOpenConnectionSelector instance.
func NewmOpenConnectionSelector() mOpenConnectionSelector {
	return getmOpenConnectionSelectorClass().New()
}




