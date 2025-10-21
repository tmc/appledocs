// Code generated from Apple documentation for CryptoTokenKit. DO NOT EDIT.

package cryptotokenkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [TKSmartCardSlot] class.
var (
	TKSmartCardSlotClass     _TKSmartCardSlotClass
	TKSmartCardSlotClassOnce sync.Once
)

func getTKSmartCardSlotClass() _TKSmartCardSlotClass {
	TKSmartCardSlotClassOnce.Do(func() {
		TKSmartCardSlotClass = _TKSmartCardSlotClass{objc.GetClass("TKSmartCardSlot")}
	})
	return TKSmartCardSlotClass
}

type _TKSmartCardSlotClass struct {
	class objc.Class
}

// An interface definition for the [TKSmartCardSlot] class.
type ITKSmartCardSlot interface {
	objectivec.IObject
	MakeSmartCard() unsafe.Pointer
}

// A single smart card reader slot in the system.
//
// Use the class to manage all the smart card reader slots available to the system. You can retrieve the names of available smart card reader slots for a system using the property of a manager object, and access instances of using the method.
//
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardSlot
type TKSmartCardSlot struct {
	objectivec.Object
}

// TKSmartCardSlotFrom constructs a [TKSmartCardSlot] from an unsafe.Pointer.
//
// A single smart card reader slot in the system.
func TKSmartCardSlotFrom(ptr unsafe.Pointer) TKSmartCardSlot {
	return TKSmartCardSlot{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (tc _TKSmartCardSlotClass) Alloc() TKSmartCardSlot {
	rv := objc.Send[TKSmartCardSlot](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (tc _TKSmartCardSlotClass) New() TKSmartCardSlot {
	rv := objc.Send[TKSmartCardSlot](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ TKSmartCardSlot) Init() TKSmartCardSlot {
	rv := objc.Send[TKSmartCardSlot](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ TKSmartCardSlot) Autorelease() TKSmartCardSlot {
	rv := objc.Send[TKSmartCardSlot](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewTKSmartCardSlot creates a new TKSmartCardSlot instance.
func NewTKSmartCardSlot() TKSmartCardSlot {
	return getTKSmartCardSlotClass().New()
}


// Creates a new object representing the currently inserted Smart Card.
//
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardSlot/makeSmartCard()
func (t_ TKSmartCardSlot) MakeSmartCard() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("makeSmartCard"))
	return rv
}

// The name of the Smart Card reader slot.
//
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardSlot/name
func (t_ TKSmartCardSlot) Name() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("name"))
	return rv
}

// The current state of the Smart Card reader slot.
//
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardSlot/state-swift.property
func (t_ TKSmartCardSlot) State() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("state"))
	return rv
}



