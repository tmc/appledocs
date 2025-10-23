// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [receivePort] class.
var (
	ReceivePortClass     _receivePortClass
	ReceivePortClassOnce sync.Once
)

func getreceivePortClass() _receivePortClass {
	ReceivePortClassOnce.Do(func() {
		ReceivePortClass = _receivePortClass{objc.GetClass("receivePort")}
	})
	return ReceivePortClass
}

type _receivePortClass struct {
	class objc.Class
}

// An interface definition for the [receivePort] class.
type IreceivePort interface {
	objectivec.IObject
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSConnection/receivePort-c.ivar
type receivePort struct {
	objectivec.Object
}

// receivePortFrom constructs a [receivePort] from an unsafe.Pointer.
func receivePortFrom(ptr unsafe.Pointer) receivePort {
	return receivePort{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (rc _receivePortClass) Alloc() receivePort {
	rv := objc.Send[receivePort](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (rc _receivePortClass) New() receivePort {
	rv := objc.Send[receivePort](objc.ID(rc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (r_ receivePort) Init() receivePort {
	rv := objc.Send[receivePort](r_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r_ receivePort) Autorelease() receivePort {
	rv := objc.Send[receivePort](r_.ID, objc.Sel("autorelease"))
	return rv
}

// NewreceivePort creates a new receivePort instance.
func NewreceivePort() receivePort {
	return getreceivePortClass().New()
}




