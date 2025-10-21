// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [reserved] class.
var (
	ReservedClass     _reservedClass
	ReservedClassOnce sync.Once
)

func getreservedClass() _reservedClass {
	ReservedClassOnce.Do(func() {
		ReservedClass = _reservedClass{objc.GetClass("reserved")}
	})
	return ReservedClass
}

type _reservedClass struct {
	class objc.Class
}

// An interface definition for the [reserved] class.
type Ireserved interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPredicate/reserved
type reserved struct {
	objectivec.Object
}

// reservedFrom constructs a [reserved] from an unsafe.Pointer.
func reservedFrom(ptr unsafe.Pointer) reserved {
	return reserved{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (rc _reservedClass) Alloc() reserved {
	rv := objc.Send[reserved](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (rc _reservedClass) New() reserved {
	rv := objc.Send[reserved](objc.ID(rc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (r_ reserved) Init() reserved {
	rv := objc.Send[reserved](r_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r_ reserved) Autorelease() reserved {
	rv := objc.Send[reserved](r_.ID, objc.Sel("autorelease"))
	return rv
}

// Newreserved creates a new reserved instance.
func Newreserved() reserved {
	return getreservedClass().New()
}




