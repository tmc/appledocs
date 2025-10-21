// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [mOBEXFilePut] class.
var (
	MOBEXFilePutClass     _mOBEXFilePutClass
	MOBEXFilePutClassOnce sync.Once
)

func getmOBEXFilePutClass() _mOBEXFilePutClass {
	MOBEXFilePutClassOnce.Do(func() {
		MOBEXFilePutClass = _mOBEXFilePutClass{objc.GetClass("mOBEXFilePut")}
	})
	return MOBEXFilePutClass
}

type _mOBEXFilePutClass struct {
	class objc.Class
}

// An interface definition for the [mOBEXFilePut] class.
type ImOBEXFilePut interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/OBEXFileTransferServices/mOBEXFilePut
type mOBEXFilePut struct {
	objectivec.Object
}

// mOBEXFilePutFrom constructs a [mOBEXFilePut] from an unsafe.Pointer.
func mOBEXFilePutFrom(ptr unsafe.Pointer) mOBEXFilePut {
	return mOBEXFilePut{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _mOBEXFilePutClass) Alloc() mOBEXFilePut {
	rv := objc.Send[mOBEXFilePut](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _mOBEXFilePutClass) New() mOBEXFilePut {
	rv := objc.Send[mOBEXFilePut](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ mOBEXFilePut) Init() mOBEXFilePut {
	rv := objc.Send[mOBEXFilePut](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ mOBEXFilePut) Autorelease() mOBEXFilePut {
	rv := objc.Send[mOBEXFilePut](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewmOBEXFilePut creates a new mOBEXFilePut instance.
func NewmOBEXFilePut() mOBEXFilePut {
	return getmOBEXFilePutClass().New()
}




