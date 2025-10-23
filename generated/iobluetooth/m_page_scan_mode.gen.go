// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [mPageScanMode] class.
var (
	MPageScanModeClass     _mPageScanModeClass
	MPageScanModeClassOnce sync.Once
)

func getmPageScanModeClass() _mPageScanModeClass {
	MPageScanModeClassOnce.Do(func() {
		MPageScanModeClass = _mPageScanModeClass{objc.GetClass("mPageScanMode")}
	})
	return MPageScanModeClass
}

type _mPageScanModeClass struct {
	class objc.Class
}

// An interface definition for the [mPageScanMode] class.
type ImPageScanMode interface {
	objectivec.IObject
	// properties:
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/mPageScanMode
type mPageScanMode struct {
	objectivec.Object
}

// mPageScanModeFrom constructs a [mPageScanMode] from an unsafe.Pointer.
func mPageScanModeFrom(ptr unsafe.Pointer) mPageScanMode {
	return mPageScanMode{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _mPageScanModeClass) Alloc() mPageScanMode {
	rv := objc.Send[mPageScanMode](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _mPageScanModeClass) New() mPageScanMode {
	rv := objc.Send[mPageScanMode](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ mPageScanMode) Init() mPageScanMode {
	rv := objc.Send[mPageScanMode](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ mPageScanMode) Autorelease() mPageScanMode {
	rv := objc.Send[mPageScanMode](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewmPageScanMode creates a new mPageScanMode instance.
func NewmPageScanMode() mPageScanMode {
	return getmPageScanModeClass().New()
}




