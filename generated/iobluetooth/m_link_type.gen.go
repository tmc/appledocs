// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [mLinkType] class.
var (
	MLinkTypeClass     _mLinkTypeClass
	MLinkTypeClassOnce sync.Once
)

func getmLinkTypeClass() _mLinkTypeClass {
	MLinkTypeClassOnce.Do(func() {
		MLinkTypeClass = _mLinkTypeClass{objc.GetClass("mLinkType")}
	})
	return MLinkTypeClass
}

type _mLinkTypeClass struct {
	class objc.Class
}

// An interface definition for the [mLinkType] class.
type ImLinkType interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/mLinkType
type mLinkType struct {
	objectivec.Object
}

// mLinkTypeFrom constructs a [mLinkType] from an unsafe.Pointer.
func mLinkTypeFrom(ptr unsafe.Pointer) mLinkType {
	return mLinkType{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _mLinkTypeClass) Alloc() mLinkType {
	rv := objc.Send[mLinkType](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _mLinkTypeClass) New() mLinkType {
	rv := objc.Send[mLinkType](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ mLinkType) Init() mLinkType {
	rv := objc.Send[mLinkType](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ mLinkType) Autorelease() mLinkType {
	rv := objc.Send[mLinkType](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewmLinkType creates a new mLinkType instance.
func NewmLinkType() mLinkType {
	return getmLinkTypeClass().New()
}




