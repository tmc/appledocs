// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [mSavedCommand] class.
var (
	MSavedCommandClass     _mSavedCommandClass
	MSavedCommandClassOnce sync.Once
)

func getmSavedCommandClass() _mSavedCommandClass {
	MSavedCommandClassOnce.Do(func() {
		MSavedCommandClass = _mSavedCommandClass{objc.GetClass("mSavedCommand")}
	})
	return MSavedCommandClass
}

type _mSavedCommandClass struct {
	class objc.Class
}

// An interface definition for the [mSavedCommand] class.
type ImSavedCommand interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/OBEXFileTransferServices/mSavedCommand
type mSavedCommand struct {
	objectivec.Object
}

// mSavedCommandFrom constructs a [mSavedCommand] from an unsafe.Pointer.
func mSavedCommandFrom(ptr unsafe.Pointer) mSavedCommand {
	return mSavedCommand{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _mSavedCommandClass) Alloc() mSavedCommand {
	rv := objc.Send[mSavedCommand](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _mSavedCommandClass) New() mSavedCommand {
	rv := objc.Send[mSavedCommand](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ mSavedCommand) Init() mSavedCommand {
	rv := objc.Send[mSavedCommand](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ mSavedCommand) Autorelease() mSavedCommand {
	rv := objc.Send[mSavedCommand](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewmSavedCommand creates a new mSavedCommand instance.
func NewmSavedCommand() mSavedCommand {
	return getmSavedCommandClass().New()
}




