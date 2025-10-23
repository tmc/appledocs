// Code generated from Apple documentation for CoreBluetooth. DO NOT EDIT.

package corebluetooth

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [CBManager] class.
var (
	CBManagerClass     _CBManagerClass
	CBManagerClassOnce sync.Once
)

func getCBManagerClass() _CBManagerClass {
	CBManagerClassOnce.Do(func() {
		CBManagerClass = _CBManagerClass{objc.GetClass("CBManager")}
	})
	return CBManagerClass
}

type _CBManagerClass struct {
	class objc.Class
}

// An interface definition for the [CBManager] class.
type ICBManager interface {
	objectivec.IObject
	// properties:
	Authorization() CBManagerAuthorization
	State() CBManagerState
	// methods:
}

// The abstract base class that manages central and peripheral objects.


// The abstract base class that manages central and peripheral objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBManager
type CBManager struct {
	objectivec.Object
}

// CBManagerFrom constructs a [CBManager] from an unsafe.Pointer.
//
// The abstract base class that manages central and peripheral objects.
func CBManagerFrom(ptr unsafe.Pointer) CBManager {
	return CBManager{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (cc _CBManagerClass) Alloc() CBManager {
	rv := objc.Send[CBManager](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CBManagerClass) New() CBManager {
	rv := objc.Send[CBManager](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CBManager) Init() CBManager {
	rv := objc.Send[CBManager](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CBManager) Autorelease() CBManager {
	rv := objc.Send[CBManager](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCBManager creates a new CBManager instance.
func NewCBManager() CBManager {
	return getCBManagerClass().New()
}



// The current authorization status for using Bluetooth.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBManager/authorization-swift.property
func (c_ CBManager) Authorization() CBManagerAuthorization {
	rv := objc.Send[CBManagerAuthorization](c_.ID, objc.Sel("authorization"))
	return rv
}


// The current state of the manager.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBManager/state
func (c_ CBManager) State() CBManagerState {
	rv := objc.Send[CBManagerState](c_.ID, objc.Sel("state"))
	return rv
}



