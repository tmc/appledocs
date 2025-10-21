// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [mOurMaxPacketLength] class.
var (
	MOurMaxPacketLengthClass     _mOurMaxPacketLengthClass
	MOurMaxPacketLengthClassOnce sync.Once
)

func getmOurMaxPacketLengthClass() _mOurMaxPacketLengthClass {
	MOurMaxPacketLengthClassOnce.Do(func() {
		MOurMaxPacketLengthClass = _mOurMaxPacketLengthClass{objc.GetClass("mOurMaxPacketLength")}
	})
	return MOurMaxPacketLengthClass
}

type _mOurMaxPacketLengthClass struct {
	class objc.Class
}

// An interface definition for the [mOurMaxPacketLength] class.
type ImOurMaxPacketLength interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/OBEXSession/mOurMaxPacketLength
type mOurMaxPacketLength struct {
	objectivec.Object
}

// mOurMaxPacketLengthFrom constructs a [mOurMaxPacketLength] from an unsafe.Pointer.
func mOurMaxPacketLengthFrom(ptr unsafe.Pointer) mOurMaxPacketLength {
	return mOurMaxPacketLength{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _mOurMaxPacketLengthClass) Alloc() mOurMaxPacketLength {
	rv := objc.Send[mOurMaxPacketLength](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _mOurMaxPacketLengthClass) New() mOurMaxPacketLength {
	rv := objc.Send[mOurMaxPacketLength](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ mOurMaxPacketLength) Init() mOurMaxPacketLength {
	rv := objc.Send[mOurMaxPacketLength](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ mOurMaxPacketLength) Autorelease() mOurMaxPacketLength {
	rv := objc.Send[mOurMaxPacketLength](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewmOurMaxPacketLength creates a new mOurMaxPacketLength instance.
func NewmOurMaxPacketLength() mOurMaxPacketLength {
	return getmOurMaxPacketLengthClass().New()
}




