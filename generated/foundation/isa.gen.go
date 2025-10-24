// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [ISA] class.
var (
	IsaClass     _ISAClass
	IsaClassOnce sync.Once
)

func getISAClass() _ISAClass {
	IsaClassOnce.Do(func() {
		IsaClass = _ISAClass{objc.GetClass("isa")}
	})
	return IsaClass
}

type _ISAClass struct {
	class objc.Class
}

// An interface definition for the [ISA] class.
type IISA interface {
	objectivec.IObject
	// properties:
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSProxy/isa
type ISA struct {
	objectivec.Object
}

// ISAFrom constructs a [ISA] from an unsafe.Pointer.
func ISAFrom(ptr unsafe.Pointer) ISA {
	return ISA{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ic _ISAClass) Alloc() ISA {
	rv := objc.Send[ISA](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ic _ISAClass) New() ISA {
	rv := objc.Send[ISA](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ ISA) Init() ISA {
	rv := objc.Send[ISA](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ ISA) Autorelease() ISA {
	rv := objc.Send[ISA](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewISA creates a new ISA instance.
func NewISA() ISA {
	return getISAClass().New()
}




