// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [mActionArgument1] class.
var (
	MActionArgument1Class     _mActionArgument1Class
	MActionArgument1ClassOnce sync.Once
)

func getmActionArgument1Class() _mActionArgument1Class {
	MActionArgument1ClassOnce.Do(func() {
		MActionArgument1Class = _mActionArgument1Class{objc.GetClass("mActionArgument1")}
	})
	return MActionArgument1Class
}

type _mActionArgument1Class struct {
	class objc.Class
}

// An interface definition for the [mActionArgument1] class.
type ImActionArgument1 interface {
	objectivec.IObject
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/OBEXFileTransferServices/mActionArgument1
type mActionArgument1 struct {
	objectivec.Object
}

// mActionArgument1From constructs a [mActionArgument1] from an unsafe.Pointer.
func mActionArgument1From(ptr unsafe.Pointer) mActionArgument1 {
	return mActionArgument1{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _mActionArgument1Class) Alloc() mActionArgument1 {
	rv := objc.Send[mActionArgument1](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _mActionArgument1Class) New() mActionArgument1 {
	rv := objc.Send[mActionArgument1](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ mActionArgument1) Init() mActionArgument1 {
	rv := objc.Send[mActionArgument1](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ mActionArgument1) Autorelease() mActionArgument1 {
	rv := objc.Send[mActionArgument1](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewmActionArgument1 creates a new mActionArgument1 instance.
func NewmActionArgument1() mActionArgument1 {
	return getmActionArgument1Class().New()
}




