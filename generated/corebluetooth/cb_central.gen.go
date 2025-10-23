// Code generated from Apple documentation for CoreBluetooth. DO NOT EDIT.

package corebluetooth

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [CBCentral] class.
var (
	CBCentralClass     _CBCentralClass
	CBCentralClassOnce sync.Once
)

func getCBCentralClass() _CBCentralClass {
	CBCentralClassOnce.Do(func() {
		CBCentralClass = _CBCentralClass{objc.GetClass("CBCentral")}
	})
	return CBCentralClass
}

type _CBCentralClass struct {
	class objc.Class
}

// An interface definition for the [CBCentral] class.
type ICBCentral interface {
	ICBPeer
	MaximumUpdateValueLength() uint
}

// A remote device connected to a local app, which is acting as a peripheral.
//
// The class represents remote central devices (or ) that have connected to an app implementing the peripheral role on a local device. Remote centrals use universally unique identifiers (UUIDs), represented by objects, to identify themselves.


// A remote device connected to a local app, which is acting as a peripheral.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBCentral
type CBCentral struct {
	CBPeer
}

// CBCentralFrom constructs a [CBCentral] from an unsafe.Pointer.
//
// A remote device connected to a local app, which is acting as a peripheral.
func CBCentralFrom(ptr unsafe.Pointer) CBCentral {
	return CBCentral{
		CBPeer: CBPeerFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (cc _CBCentralClass) Alloc() CBCentral {
	rv := objc.Send[CBCentral](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CBCentralClass) New() CBCentral {
	rv := objc.Send[CBCentral](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CBCentral) Init() CBCentral {
	rv := objc.Send[CBCentral](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CBCentral) Autorelease() CBCentral {
	rv := objc.Send[CBCentral](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCBCentral creates a new CBCentral instance.
func NewCBCentral() CBCentral {
	return getCBCentralClass().New()
}



// The maximum amount of data, in bytes, that the central can receive in a single notification or indication.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBCentral/maximumUpdateValueLength
func (c_ CBCentral) MaximumUpdateValueLength() uint {
	rv := objc.Send[uint](c_.ID, objc.Sel("maximumUpdateValueLength"))
	return rv
}



