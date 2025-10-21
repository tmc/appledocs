// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [mPageScanRepetitionMode] class.
var (
	MPageScanRepetitionModeClass     _mPageScanRepetitionModeClass
	MPageScanRepetitionModeClassOnce sync.Once
)

func getmPageScanRepetitionModeClass() _mPageScanRepetitionModeClass {
	MPageScanRepetitionModeClassOnce.Do(func() {
		MPageScanRepetitionModeClass = _mPageScanRepetitionModeClass{objc.GetClass("mPageScanRepetitionMode")}
	})
	return MPageScanRepetitionModeClass
}

type _mPageScanRepetitionModeClass struct {
	class objc.Class
}

// An interface definition for the [mPageScanRepetitionMode] class.
type ImPageScanRepetitionMode interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/mPageScanRepetitionMode
type mPageScanRepetitionMode struct {
	objectivec.Object
}

// mPageScanRepetitionModeFrom constructs a [mPageScanRepetitionMode] from an unsafe.Pointer.
func mPageScanRepetitionModeFrom(ptr unsafe.Pointer) mPageScanRepetitionMode {
	return mPageScanRepetitionMode{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _mPageScanRepetitionModeClass) Alloc() mPageScanRepetitionMode {
	rv := objc.Send[mPageScanRepetitionMode](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _mPageScanRepetitionModeClass) New() mPageScanRepetitionMode {
	rv := objc.Send[mPageScanRepetitionMode](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ mPageScanRepetitionMode) Init() mPageScanRepetitionMode {
	rv := objc.Send[mPageScanRepetitionMode](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ mPageScanRepetitionMode) Autorelease() mPageScanRepetitionMode {
	rv := objc.Send[mPageScanRepetitionMode](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewmPageScanRepetitionMode creates a new mPageScanRepetitionMode instance.
func NewmPageScanRepetitionMode() mPageScanRepetitionMode {
	return getmPageScanRepetitionModeClass().New()
}




