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

// Information about an observed iBeacon device and its relative distance to a person’s device.
//
// The class represents a beacon that was observed during beacon ranging. You do not create instances of this class directly. The location manager ( ) object reports observed beacons to its associated delegate object. The identity of a beacon is defined by its , , and properties. These values are coded into the beacon itself. For a more thorough description of the meaning of those values, see .
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


// The accuracy of the proximity value, measured in meters from the beacon.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLBeacon/accuracy
func (b_ Beacon) Accuracy() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("accuracy"))
	return rv
}

// The major value that the observed beacon transmitted.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLBeacon/major
func (b_ Beacon) Major() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("major"))
	return rv
}

// The minor value that the observed beacon transmitted.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLBeacon/minor
func (b_ Beacon) Minor() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("minor"))
	return rv
}

// The relative distance to the beacon.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLBeacon/proximity
func (b_ Beacon) Proximity() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("proximity"))
	return rv
}

// The proximity ID of the beacon.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLBeacon/proximityUUID
func (b_ Beacon) ProximityUUID() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("proximityUUID"))
	return rv
}

// The received signal strength of the beacon, measured in decibels.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLBeacon/rssi
func (b_ Beacon) Rssi() int {
	rv := objc.Send[int](b_.ID, objc.Sel("rssi"))
	return rv
}

// A timestamp representing when the beacon was observed.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLBeacon/timestamp
func (b_ Beacon) Timestamp() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("timestamp"))
	return rv
}

// The UUID that the observed beacon transmitted.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLBeacon/uuid
func (b_ Beacon) UUID() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("UUID"))
	return rv
}



