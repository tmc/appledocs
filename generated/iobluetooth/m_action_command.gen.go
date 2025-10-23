// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [mActionCommand] class.
var (
	MActionCommandClass     _mActionCommandClass
	MActionCommandClassOnce sync.Once
)

func getmActionCommandClass() _mActionCommandClass {
	MActionCommandClassOnce.Do(func() {
		MActionCommandClass = _mActionCommandClass{objc.GetClass("mActionCommand")}
	})
	return MActionCommandClass
}

type _mActionCommandClass struct {
	class objc.Class
}

// An interface definition for the [mActionCommand] class.
type ImActionCommand interface {
	objectivec.IObject
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/OBEXFileTransferServices/mActionCommand
type mActionCommand struct {
	objectivec.Object
}

// mActionCommandFrom constructs a [mActionCommand] from an unsafe.Pointer.
func mActionCommandFrom(ptr unsafe.Pointer) mActionCommand {
	return mActionCommand{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _mActionCommandClass) Alloc() mActionCommand {
	rv := objc.Send[mActionCommand](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _mActionCommandClass) New() mActionCommand {
	rv := objc.Send[mActionCommand](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ mActionCommand) Init() mActionCommand {
	rv := objc.Send[mActionCommand](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ mActionCommand) Autorelease() mActionCommand {
	rv := objc.Send[mActionCommand](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewmActionCommand creates a new mActionCommand instance.
func NewmActionCommand() mActionCommand {
	return getmActionCommandClass().New()
}




