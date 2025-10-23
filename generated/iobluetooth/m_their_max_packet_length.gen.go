// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [mTheirMaxPacketLength] class.
var (
	MTheirMaxPacketLengthClass     _mTheirMaxPacketLengthClass
	MTheirMaxPacketLengthClassOnce sync.Once
)

func getmTheirMaxPacketLengthClass() _mTheirMaxPacketLengthClass {
	MTheirMaxPacketLengthClassOnce.Do(func() {
		MTheirMaxPacketLengthClass = _mTheirMaxPacketLengthClass{objc.GetClass("mTheirMaxPacketLength")}
	})
	return MTheirMaxPacketLengthClass
}

type _mTheirMaxPacketLengthClass struct {
	class objc.Class
}

// An interface definition for the [mTheirMaxPacketLength] class.
type ImTheirMaxPacketLength interface {
	objectivec.IObject
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/OBEXSession/mTheirMaxPacketLength
type mTheirMaxPacketLength struct {
	objectivec.Object
}

// mTheirMaxPacketLengthFrom constructs a [mTheirMaxPacketLength] from an unsafe.Pointer.
func mTheirMaxPacketLengthFrom(ptr unsafe.Pointer) mTheirMaxPacketLength {
	return mTheirMaxPacketLength{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _mTheirMaxPacketLengthClass) Alloc() mTheirMaxPacketLength {
	rv := objc.Send[mTheirMaxPacketLength](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _mTheirMaxPacketLengthClass) New() mTheirMaxPacketLength {
	rv := objc.Send[mTheirMaxPacketLength](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ mTheirMaxPacketLength) Init() mTheirMaxPacketLength {
	rv := objc.Send[mTheirMaxPacketLength](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ mTheirMaxPacketLength) Autorelease() mTheirMaxPacketLength {
	rv := objc.Send[mTheirMaxPacketLength](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewmTheirMaxPacketLength creates a new mTheirMaxPacketLength instance.
func NewmTheirMaxPacketLength() mTheirMaxPacketLength {
	return getmTheirMaxPacketLengthClass().New()
}





