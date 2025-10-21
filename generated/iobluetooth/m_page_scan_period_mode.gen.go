// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [mPageScanPeriodMode] class.
var (
	MPageScanPeriodModeClass     _mPageScanPeriodModeClass
	MPageScanPeriodModeClassOnce sync.Once
)

func getmPageScanPeriodModeClass() _mPageScanPeriodModeClass {
	MPageScanPeriodModeClassOnce.Do(func() {
		MPageScanPeriodModeClass = _mPageScanPeriodModeClass{objc.GetClass("mPageScanPeriodMode")}
	})
	return MPageScanPeriodModeClass
}

type _mPageScanPeriodModeClass struct {
	class objc.Class
}

// An interface definition for the [mPageScanPeriodMode] class.
type ImPageScanPeriodMode interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/mPageScanPeriodMode
type mPageScanPeriodMode struct {
	objectivec.Object
}

// mPageScanPeriodModeFrom constructs a [mPageScanPeriodMode] from an unsafe.Pointer.
func mPageScanPeriodModeFrom(ptr unsafe.Pointer) mPageScanPeriodMode {
	return mPageScanPeriodMode{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _mPageScanPeriodModeClass) Alloc() mPageScanPeriodMode {
	rv := objc.Send[mPageScanPeriodMode](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _mPageScanPeriodModeClass) New() mPageScanPeriodMode {
	rv := objc.Send[mPageScanPeriodMode](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ mPageScanPeriodMode) Init() mPageScanPeriodMode {
	rv := objc.Send[mPageScanPeriodMode](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ mPageScanPeriodMode) Autorelease() mPageScanPeriodMode {
	rv := objc.Send[mPageScanPeriodMode](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewmPageScanPeriodMode creates a new mPageScanPeriodMode instance.
func NewmPageScanPeriodMode() mPageScanPeriodMode {
	return getmPageScanPeriodModeClass().New()
}




