// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [mOBEXFileGet] class.
var (
	MOBEXFileGetClass     _mOBEXFileGetClass
	MOBEXFileGetClassOnce sync.Once
)

func getmOBEXFileGetClass() _mOBEXFileGetClass {
	MOBEXFileGetClassOnce.Do(func() {
		MOBEXFileGetClass = _mOBEXFileGetClass{objc.GetClass("mOBEXFileGet")}
	})
	return MOBEXFileGetClass
}

type _mOBEXFileGetClass struct {
	class objc.Class
}

// An interface definition for the [mOBEXFileGet] class.
type ImOBEXFileGet interface {
	objectivec.IObject
	// properties:
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/OBEXFileTransferServices/mOBEXFileGet
type mOBEXFileGet struct {
	objectivec.Object
}

// mOBEXFileGetFrom constructs a [mOBEXFileGet] from an unsafe.Pointer.
func mOBEXFileGetFrom(ptr unsafe.Pointer) mOBEXFileGet {
	return mOBEXFileGet{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _mOBEXFileGetClass) Alloc() mOBEXFileGet {
	rv := objc.Send[mOBEXFileGet](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _mOBEXFileGetClass) New() mOBEXFileGet {
	rv := objc.Send[mOBEXFileGet](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ mOBEXFileGet) Init() mOBEXFileGet {
	rv := objc.Send[mOBEXFileGet](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ mOBEXFileGet) Autorelease() mOBEXFileGet {
	rv := objc.Send[mOBEXFileGet](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewmOBEXFileGet creates a new mOBEXFileGet instance.
func NewmOBEXFileGet() mOBEXFileGet {
	return getmOBEXFileGetClass().New()
}




