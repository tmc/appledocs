// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [addresses] class.
var (
	AddressesClass     _addressesClass
	AddressesClassOnce sync.Once
)

func getaddressesClass() _addressesClass {
	AddressesClassOnce.Do(func() {
		AddressesClass = _addressesClass{objc.GetClass("addresses")}
	})
	return AddressesClass
}

type _addressesClass struct {
	class objc.Class
}

// An interface definition for the [addresses] class.
type Iaddresses interface {
	objectivec.IObject
	// properties:
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSHost/addresses-c.ivar
type addresses struct {
	objectivec.Object
}

// addressesFrom constructs a [addresses] from an unsafe.Pointer.
func addressesFrom(ptr unsafe.Pointer) addresses {
	return addresses{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ac _addressesClass) Alloc() addresses {
	rv := objc.Send[addresses](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _addressesClass) New() addresses {
	rv := objc.Send[addresses](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ addresses) Init() addresses {
	rv := objc.Send[addresses](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ addresses) Autorelease() addresses {
	rv := objc.Send[addresses](a_.ID, objc.Sel("autorelease"))
	return rv
}

// Newaddresses creates a new addresses instance.
func Newaddresses() addresses {
	return getaddressesClass().New()
}




