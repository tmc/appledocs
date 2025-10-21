// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [mIncompletePacketResponseCode] class.
var (
	MIncompletePacketResponseCodeClass     _mIncompletePacketResponseCodeClass
	MIncompletePacketResponseCodeClassOnce sync.Once
)

func getmIncompletePacketResponseCodeClass() _mIncompletePacketResponseCodeClass {
	MIncompletePacketResponseCodeClassOnce.Do(func() {
		MIncompletePacketResponseCodeClass = _mIncompletePacketResponseCodeClass{objc.GetClass("mIncompletePacketResponseCode")}
	})
	return MIncompletePacketResponseCodeClass
}

type _mIncompletePacketResponseCodeClass struct {
	class objc.Class
}

// An interface definition for the [mIncompletePacketResponseCode] class.
type ImIncompletePacketResponseCode interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/OBEXSession/mIncompletePacketResponseCode
type mIncompletePacketResponseCode struct {
	objectivec.Object
}

// mIncompletePacketResponseCodeFrom constructs a [mIncompletePacketResponseCode] from an unsafe.Pointer.
func mIncompletePacketResponseCodeFrom(ptr unsafe.Pointer) mIncompletePacketResponseCode {
	return mIncompletePacketResponseCode{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _mIncompletePacketResponseCodeClass) Alloc() mIncompletePacketResponseCode {
	rv := objc.Send[mIncompletePacketResponseCode](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _mIncompletePacketResponseCodeClass) New() mIncompletePacketResponseCode {
	rv := objc.Send[mIncompletePacketResponseCode](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ mIncompletePacketResponseCode) Init() mIncompletePacketResponseCode {
	rv := objc.Send[mIncompletePacketResponseCode](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ mIncompletePacketResponseCode) Autorelease() mIncompletePacketResponseCode {
	rv := objc.Send[mIncompletePacketResponseCode](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewmIncompletePacketResponseCode creates a new mIncompletePacketResponseCode instance.
func NewmIncompletePacketResponseCode() mIncompletePacketResponseCode {
	return getmIncompletePacketResponseCodeClass().New()
}




