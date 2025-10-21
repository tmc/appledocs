// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [mHasTargetHeader] class.
var (
	MHasTargetHeaderClass     _mHasTargetHeaderClass
	MHasTargetHeaderClassOnce sync.Once
)

func getmHasTargetHeaderClass() _mHasTargetHeaderClass {
	MHasTargetHeaderClassOnce.Do(func() {
		MHasTargetHeaderClass = _mHasTargetHeaderClass{objc.GetClass("mHasTargetHeader")}
	})
	return MHasTargetHeaderClass
}

type _mHasTargetHeaderClass struct {
	class objc.Class
}

// An interface definition for the [mHasTargetHeader] class.
type ImHasTargetHeader interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/OBEXFileTransferServices/mHasTargetHeader
type mHasTargetHeader struct {
	objectivec.Object
}

// mHasTargetHeaderFrom constructs a [mHasTargetHeader] from an unsafe.Pointer.
func mHasTargetHeaderFrom(ptr unsafe.Pointer) mHasTargetHeader {
	return mHasTargetHeader{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _mHasTargetHeaderClass) Alloc() mHasTargetHeader {
	rv := objc.Send[mHasTargetHeader](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _mHasTargetHeaderClass) New() mHasTargetHeader {
	rv := objc.Send[mHasTargetHeader](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ mHasTargetHeader) Init() mHasTargetHeader {
	rv := objc.Send[mHasTargetHeader](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ mHasTargetHeader) Autorelease() mHasTargetHeader {
	rv := objc.Send[mHasTargetHeader](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewmHasTargetHeader creates a new mHasTargetHeader instance.
func NewmHasTargetHeader() mHasTargetHeader {
	return getmHasTargetHeaderClass().New()
}




