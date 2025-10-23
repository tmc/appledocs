// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [mConnectionID] class.
var (
	MConnectionIDClass     _mConnectionIDClass
	MConnectionIDClassOnce sync.Once
)

func getmConnectionIDClass() _mConnectionIDClass {
	MConnectionIDClassOnce.Do(func() {
		MConnectionIDClass = _mConnectionIDClass{objc.GetClass("mConnectionID")}
	})
	return MConnectionIDClass
}

type _mConnectionIDClass struct {
	class objc.Class
}

// An interface definition for the [mConnectionID] class.
type ImConnectionID interface {
	objectivec.IObject
	// properties:
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/OBEXFileTransferServices/mConnectionID
type mConnectionID struct {
	objectivec.Object
}

// mConnectionIDFrom constructs a [mConnectionID] from an unsafe.Pointer.
func mConnectionIDFrom(ptr unsafe.Pointer) mConnectionID {
	return mConnectionID{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _mConnectionIDClass) Alloc() mConnectionID {
	rv := objc.Send[mConnectionID](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _mConnectionIDClass) New() mConnectionID {
	rv := objc.Send[mConnectionID](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ mConnectionID) Init() mConnectionID {
	rv := objc.Send[mConnectionID](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ mConnectionID) Autorelease() mConnectionID {
	rv := objc.Send[mConnectionID](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewmConnectionID creates a new mConnectionID instance.
func NewmConnectionID() mConnectionID {
	return getmConnectionIDClass().New()
}




