// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [mAbortedByOwner] class.
var (
	MAbortedByOwnerClass     _mAbortedByOwnerClass
	MAbortedByOwnerClassOnce sync.Once
)

func getmAbortedByOwnerClass() _mAbortedByOwnerClass {
	MAbortedByOwnerClassOnce.Do(func() {
		MAbortedByOwnerClass = _mAbortedByOwnerClass{objc.GetClass("mAbortedByOwner")}
	})
	return MAbortedByOwnerClass
}

type _mAbortedByOwnerClass struct {
	class objc.Class
}

// An interface definition for the [mAbortedByOwner] class.
type ImAbortedByOwner interface {
	objectivec.IObject
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/OBEXFileTransferServices/mAbortedByOwner
type mAbortedByOwner struct {
	objectivec.Object
}

// mAbortedByOwnerFrom constructs a [mAbortedByOwner] from an unsafe.Pointer.
func mAbortedByOwnerFrom(ptr unsafe.Pointer) mAbortedByOwner {
	return mAbortedByOwner{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _mAbortedByOwnerClass) Alloc() mAbortedByOwner {
	rv := objc.Send[mAbortedByOwner](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _mAbortedByOwnerClass) New() mAbortedByOwner {
	rv := objc.Send[mAbortedByOwner](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ mAbortedByOwner) Init() mAbortedByOwner {
	rv := objc.Send[mAbortedByOwner](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ mAbortedByOwner) Autorelease() mAbortedByOwner {
	rv := objc.Send[mAbortedByOwner](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewmAbortedByOwner creates a new mAbortedByOwner instance.
func NewmAbortedByOwner() mAbortedByOwner {
	return getmAbortedByOwnerClass().New()
}




