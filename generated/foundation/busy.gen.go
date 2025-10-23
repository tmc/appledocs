// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [busy] class.
var (
	BusyClass     _busyClass
	BusyClassOnce sync.Once
)

func getbusyClass() _busyClass {
	BusyClassOnce.Do(func() {
		BusyClass = _busyClass{objc.GetClass("busy")}
	})
	return BusyClass
}

type _busyClass struct {
	class objc.Class
}

// An interface definition for the [busy] class.
type Ibusy interface {
	objectivec.IObject
	// properties:
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSConnection/busy
type busy struct {
	objectivec.Object
}

// busyFrom constructs a [busy] from an unsafe.Pointer.
func busyFrom(ptr unsafe.Pointer) busy {
	return busy{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (bc _busyClass) Alloc() busy {
	rv := objc.Send[busy](objc.ID(bc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (bc _busyClass) New() busy {
	rv := objc.Send[busy](objc.ID(bc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (b_ busy) Init() busy {
	rv := objc.Send[busy](b_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (b_ busy) Autorelease() busy {
	rv := objc.Send[busy](b_.ID, objc.Sel("autorelease"))
	return rv
}

// Newbusy creates a new busy instance.
func Newbusy() busy {
	return getbusyClass().New()
}




