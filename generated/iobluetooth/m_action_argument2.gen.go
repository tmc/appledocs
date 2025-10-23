// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [mActionArgument2] class.
var (
	MActionArgument2Class     _mActionArgument2Class
	MActionArgument2ClassOnce sync.Once
)

func getmActionArgument2Class() _mActionArgument2Class {
	MActionArgument2ClassOnce.Do(func() {
		MActionArgument2Class = _mActionArgument2Class{objc.GetClass("mActionArgument2")}
	})
	return MActionArgument2Class
}

type _mActionArgument2Class struct {
	class objc.Class
}

// An interface definition for the [mActionArgument2] class.
type ImActionArgument2 interface {
	objectivec.IObject
	// properties:
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/OBEXFileTransferServices/mActionArgument2
type mActionArgument2 struct {
	objectivec.Object
}

// mActionArgument2From constructs a [mActionArgument2] from an unsafe.Pointer.
func mActionArgument2From(ptr unsafe.Pointer) mActionArgument2 {
	return mActionArgument2{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _mActionArgument2Class) Alloc() mActionArgument2 {
	rv := objc.Send[mActionArgument2](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _mActionArgument2Class) New() mActionArgument2 {
	rv := objc.Send[mActionArgument2](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ mActionArgument2) Init() mActionArgument2 {
	rv := objc.Send[mActionArgument2](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ mActionArgument2) Autorelease() mActionArgument2 {
	rv := objc.Send[mActionArgument2](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewmActionArgument2 creates a new mActionArgument2 instance.
func NewmActionArgument2() mActionArgument2 {
	return getmActionArgument2Class().New()
}




