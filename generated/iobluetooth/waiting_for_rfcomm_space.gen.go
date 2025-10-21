// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [waitingForRfcommSpace] class.
var (
	WaitingForRfcommSpaceClass     _waitingForRfcommSpaceClass
	WaitingForRfcommSpaceClassOnce sync.Once
)

func getwaitingForRfcommSpaceClass() _waitingForRfcommSpaceClass {
	WaitingForRfcommSpaceClassOnce.Do(func() {
		WaitingForRfcommSpaceClass = _waitingForRfcommSpaceClass{objc.GetClass("waitingForRfcommSpace")}
	})
	return WaitingForRfcommSpaceClass
}

type _waitingForRfcommSpaceClass struct {
	class objc.Class
}

// An interface definition for the [waitingForRfcommSpace] class.
type IwaitingForRfcommSpace interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothOBEXSession/waitingForRfcommSpace
type waitingForRfcommSpace struct {
	objectivec.Object
}

// waitingForRfcommSpaceFrom constructs a [waitingForRfcommSpace] from an unsafe.Pointer.
func waitingForRfcommSpaceFrom(ptr unsafe.Pointer) waitingForRfcommSpace {
	return waitingForRfcommSpace{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (wc _waitingForRfcommSpaceClass) Alloc() waitingForRfcommSpace {
	rv := objc.Send[waitingForRfcommSpace](objc.ID(wc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (wc _waitingForRfcommSpaceClass) New() waitingForRfcommSpace {
	rv := objc.Send[waitingForRfcommSpace](objc.ID(wc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (w_ waitingForRfcommSpace) Init() waitingForRfcommSpace {
	rv := objc.Send[waitingForRfcommSpace](w_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (w_ waitingForRfcommSpace) Autorelease() waitingForRfcommSpace {
	rv := objc.Send[waitingForRfcommSpace](w_.ID, objc.Sel("autorelease"))
	return rv
}

// NewwaitingForRfcommSpace creates a new waitingForRfcommSpace instance.
func NewwaitingForRfcommSpace() waitingForRfcommSpace {
	return getwaitingForRfcommSpaceClass().New()
}




