// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [mMaxPacketLength] class.
var (
	MMaxPacketLengthClass     _mMaxPacketLengthClass
	MMaxPacketLengthClassOnce sync.Once
)

func getmMaxPacketLengthClass() _mMaxPacketLengthClass {
	MMaxPacketLengthClassOnce.Do(func() {
		MMaxPacketLengthClass = _mMaxPacketLengthClass{objc.GetClass("mMaxPacketLength")}
	})
	return MMaxPacketLengthClass
}

type _mMaxPacketLengthClass struct {
	class objc.Class
}

// An interface definition for the [mMaxPacketLength] class.
type ImMaxPacketLength interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/OBEXSession/mMaxPacketLength
type mMaxPacketLength struct {
	objectivec.Object
}

// mMaxPacketLengthFrom constructs a [mMaxPacketLength] from an unsafe.Pointer.
func mMaxPacketLengthFrom(ptr unsafe.Pointer) mMaxPacketLength {
	return mMaxPacketLength{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _mMaxPacketLengthClass) Alloc() mMaxPacketLength {
	rv := objc.Send[mMaxPacketLength](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _mMaxPacketLengthClass) New() mMaxPacketLength {
	rv := objc.Send[mMaxPacketLength](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ mMaxPacketLength) Init() mMaxPacketLength {
	rv := objc.Send[mMaxPacketLength](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ mMaxPacketLength) Autorelease() mMaxPacketLength {
	rv := objc.Send[mMaxPacketLength](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewmMaxPacketLength creates a new mMaxPacketLength instance.
func NewmMaxPacketLength() mMaxPacketLength {
	return getmMaxPacketLengthClass().New()
}




