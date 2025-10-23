// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [mActionArgument3] class.
var (
	MActionArgument3Class     _mActionArgument3Class
	MActionArgument3ClassOnce sync.Once
)

func getmActionArgument3Class() _mActionArgument3Class {
	MActionArgument3ClassOnce.Do(func() {
		MActionArgument3Class = _mActionArgument3Class{objc.GetClass("mActionArgument3")}
	})
	return MActionArgument3Class
}

type _mActionArgument3Class struct {
	class objc.Class
}

// An interface definition for the [mActionArgument3] class.
type ImActionArgument3 interface {
	objectivec.IObject
	// properties:
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/OBEXFileTransferServices/mActionArgument3
type mActionArgument3 struct {
	objectivec.Object
}

// mActionArgument3From constructs a [mActionArgument3] from an unsafe.Pointer.
func mActionArgument3From(ptr unsafe.Pointer) mActionArgument3 {
	return mActionArgument3{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _mActionArgument3Class) Alloc() mActionArgument3 {
	rv := objc.Send[mActionArgument3](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _mActionArgument3Class) New() mActionArgument3 {
	rv := objc.Send[mActionArgument3](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ mActionArgument3) Init() mActionArgument3 {
	rv := objc.Send[mActionArgument3](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ mActionArgument3) Autorelease() mActionArgument3 {
	rv := objc.Send[mActionArgument3](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewmActionArgument3 creates a new mActionArgument3 instance.
func NewmActionArgument3() mActionArgument3 {
	return getmActionArgument3Class().New()
}




