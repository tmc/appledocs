// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [mCurrentRemoteDirectory] class.
var (
	MCurrentRemoteDirectoryClass     _mCurrentRemoteDirectoryClass
	MCurrentRemoteDirectoryClassOnce sync.Once
)

func getmCurrentRemoteDirectoryClass() _mCurrentRemoteDirectoryClass {
	MCurrentRemoteDirectoryClassOnce.Do(func() {
		MCurrentRemoteDirectoryClass = _mCurrentRemoteDirectoryClass{objc.GetClass("mCurrentRemoteDirectory")}
	})
	return MCurrentRemoteDirectoryClass
}

type _mCurrentRemoteDirectoryClass struct {
	class objc.Class
}

// An interface definition for the [mCurrentRemoteDirectory] class.
type ImCurrentRemoteDirectory interface {
	objectivec.IObject
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/OBEXFileTransferServices/mCurrentRemoteDirectory
type mCurrentRemoteDirectory struct {
	objectivec.Object
}

// mCurrentRemoteDirectoryFrom constructs a [mCurrentRemoteDirectory] from an unsafe.Pointer.
func mCurrentRemoteDirectoryFrom(ptr unsafe.Pointer) mCurrentRemoteDirectory {
	return mCurrentRemoteDirectory{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _mCurrentRemoteDirectoryClass) Alloc() mCurrentRemoteDirectory {
	rv := objc.Send[mCurrentRemoteDirectory](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _mCurrentRemoteDirectoryClass) New() mCurrentRemoteDirectory {
	rv := objc.Send[mCurrentRemoteDirectory](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ mCurrentRemoteDirectory) Init() mCurrentRemoteDirectory {
	rv := objc.Send[mCurrentRemoteDirectory](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ mCurrentRemoteDirectory) Autorelease() mCurrentRemoteDirectory {
	rv := objc.Send[mCurrentRemoteDirectory](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewmCurrentRemoteDirectory creates a new mCurrentRemoteDirectory instance.
func NewmCurrentRemoteDirectory() mCurrentRemoteDirectory {
	return getmCurrentRemoteDirectoryClass().New()
}




