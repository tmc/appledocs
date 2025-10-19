// Code generated from Apple documentation for CoreLocation. DO NOT EDIT.

package corelocation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [Beacon] class.
var (
	beaconClass     _BeaconClass
	beaconClassOnce sync.Once
)

func getBeaconClass() _BeaconClass {
	beaconClassOnce.Do(func() {
		beaconClass = _BeaconClass{objc.GetClass("CLBeacon")}
	})
	return beaconClass
}

type _BeaconClass struct {
	class objc.Class
}

// An interface definition for the [Beacon] class.
type IBeacon interface {
	objectivec.IObject
}

// Information about an observed iBeacon device and its relative distance to a person’s device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLBeacon
type Beacon struct {
	objectivec.Object
}

// BeaconFrom constructs a [Beacon] from an unsafe.Pointer.
//
// Information about an observed iBeacon device and its relative distance to a person’s device.
func BeaconFrom(ptr unsafe.Pointer) Beacon {
	return Beacon{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (bc _BeaconClass) Alloc() Beacon {
	rv := objc.Send[Beacon](objc.ID(bc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (bc _BeaconClass) New() Beacon {
	rv := objc.Send[Beacon](objc.ID(bc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (b_ Beacon) Init() Beacon {
	rv := objc.Send[Beacon](b_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (b_ Beacon) Autorelease() Beacon {
	rv := objc.Send[Beacon](b_.ID, objc.Sel("autorelease"))
	return rv
}

// NewBeacon creates a new Beacon instance.
func NewBeacon() Beacon {
	return getBeaconClass().New()
}




