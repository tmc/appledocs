// Code generated from Apple documentation for CoreLocation. DO NOT EDIT.

package corelocation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [BeaconRegion] class.
var (
	BeaconRegionClass     _BeaconRegionClass
	BeaconRegionClassOnce sync.Once
)

func getBeaconRegionClass() _BeaconRegionClass {
	BeaconRegionClassOnce.Do(func() {
		BeaconRegionClass = _BeaconRegionClass{objc.GetClass("CLBeaconRegion")}
	})
	return BeaconRegionClass
}

type _BeaconRegionClass struct {
	class objc.Class
}

// An interface definition for the [BeaconRegion] class.
type IBeaconRegion interface {
	IRegion
	PeripheralDataWithMeasuredPower(measuredPower foundation.INumber) unsafe.Pointer
	BeaconIdentityConstraint() CLBeaconIdentityConstraint
	Major() foundation.Number
	Minor() foundation.Number
	NotifyEntryStateOnDisplay() bool
	SetNotifyEntryStateOnDisplay(value bool)
	ProximityUUID() foundation.UUID
	UUID() foundation.UUID
}

// A region for detecting the presence of iBeacon devices.
//
// A object defines a region that you use to detect Bluetooth beacons conforming to the iBeacon specification. In contrast to a that centers on a geographic location, a focuses on an iBeacon with specific identifying characteristics, which you provide. When a matching device comes in range, Core Location notifies your app. You monitor beacon regions in two ways. To detect when a beacon is in range, use the method of your location manager object. After detecting a beacon, call the method to determine the relative distance to that beacon. When detecting an iBeacon, you need to specify the , , and values that you programmed into the beacon hardware. You use the values to identify your beacons uniquely, and you can specify a subset of values to detect multiple beacons. The property is typically the same for all of the beacons in your installation. Use the and values to distinguish among different beacons in your installation. If you want to configure the current iOS device as a Bluetooth beacon, create a beacon region with the appropriate identifying information. You can then call the method of the region to get a dictionary that you can use to advertise the device with the Core Bluetooth framework. For more information about using that framework to advertise the device as a beacon, see . For information about how to detect beacons, see .
//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLBeaconRegion
type BeaconRegion struct {
	Region
}

// BeaconRegionFrom constructs a [BeaconRegion] from an unsafe.Pointer.
//
// A region for detecting the presence of iBeacon devices.
func BeaconRegionFrom(ptr unsafe.Pointer) BeaconRegion {
	return BeaconRegion{
		Region: RegionFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (bc _BeaconRegionClass) Alloc() BeaconRegion {
	rv := objc.Send[BeaconRegion](objc.ID(bc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (bc _BeaconRegionClass) New() BeaconRegion {
	rv := objc.Send[BeaconRegion](objc.ID(bc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (b_ BeaconRegion) Init() BeaconRegion {
	rv := objc.Send[BeaconRegion](b_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (b_ BeaconRegion) Autorelease() BeaconRegion {
	rv := objc.Send[BeaconRegion](b_.ID, objc.Sel("autorelease"))
	return rv
}

// NewBeaconRegion creates a new BeaconRegion instance.
func NewBeaconRegion() BeaconRegion {
	return getBeaconRegionClass().New()
}




// Creates and returns a region object that targets beacons that satisfy the specified beacon identity constraints.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLBeaconRegion/init(beaconIdentityConstraint:identifier:)
func NewBeaconRegionWithBeaconIdentityConstraintIdentifier(beaconIdentityConstraint ICLBeaconIdentityConstraint, identifier string) BeaconRegion {
	instance := getBeaconRegionClass().Alloc()
	rv := objc.Send[BeaconRegion](instance.ID, objc.Sel("initWithBeaconIdentityConstraint:identifier:"), beaconIdentityConstraint, objc.String(identifier))
	rv.Autorelease()
	return rv
}



// Creates and returns a region object that targets a beacon with the specified UUID.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLBeaconRegion/init(proximityUUID:identifier:)
func NewBeaconRegionWithProximityUUIDIdentifier(proximityUUID foundation.IUUID, identifier string) BeaconRegion {
	instance := getBeaconRegionClass().Alloc()
	rv := objc.Send[BeaconRegion](instance.ID, objc.Sel("initWithProximityUUID:identifier:"), proximityUUID, objc.String(identifier))
	rv.Autorelease()
	return rv
}



// Creates and returns a region object that targets a beacon with the specified proximity ID and major value.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLBeaconRegion/init(proximityUUID:major:identifier:)
func NewBeaconRegionWithProximityUUIDMajorIdentifier(proximityUUID foundation.IUUID, major IBeaconMajorValue, identifier string) BeaconRegion {
	instance := getBeaconRegionClass().Alloc()
	rv := objc.Send[BeaconRegion](instance.ID, objc.Sel("initWithProximityUUID:major:identifier:"), proximityUUID, major, objc.String(identifier))
	rv.Autorelease()
	return rv
}



// Creates and returns a region object that targets a beacon with the specified proximity ID, major value, and minor value.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLBeaconRegion/init(proximityUUID:major:minor:identifier:)
func NewBeaconRegionWithProximityUUIDMajorMinorIdentifier(proximityUUID foundation.IUUID, major IBeaconMajorValue, minor IBeaconMinorValue, identifier string) BeaconRegion {
	instance := getBeaconRegionClass().Alloc()
	rv := objc.Send[BeaconRegion](instance.ID, objc.Sel("initWithProximityUUID:major:minor:identifier:"), proximityUUID, major, minor, objc.String(identifier))
	rv.Autorelease()
	return rv
}



// Creates and returns a region object that targets beacons with the specified UUID.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLBeaconRegion/init(uuid:identifier:)
func NewBeaconRegionWithUUIDIdentifier(uuid foundation.IUUID, identifier string) BeaconRegion {
	instance := getBeaconRegionClass().Alloc()
	rv := objc.Send[BeaconRegion](instance.ID, objc.Sel("initWithUUID:identifier:"), uuid, objc.String(identifier))
	rv.Autorelease()
	return rv
}



// Creates and returns a region object that targets beacons with the specified UUID and major value.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLBeaconRegion/init(uuid:major:identifier:)
func NewBeaconRegionWithUUIDMajorIdentifier(uuid foundation.IUUID, major IBeaconMajorValue, identifier string) BeaconRegion {
	instance := getBeaconRegionClass().Alloc()
	rv := objc.Send[BeaconRegion](instance.ID, objc.Sel("initWithUUID:major:identifier:"), uuid, major, objc.String(identifier))
	rv.Autorelease()
	return rv
}



// Creates and returns a region object that targets beacons with the specified UUID, and major and minor values.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLBeaconRegion/init(uuid:major:minor:identifier:)
func NewBeaconRegionWithUUIDMajorMinorIdentifier(uuid foundation.IUUID, major IBeaconMajorValue, minor IBeaconMinorValue, identifier string) BeaconRegion {
	instance := getBeaconRegionClass().Alloc()
	rv := objc.Send[BeaconRegion](instance.ID, objc.Sel("initWithUUID:major:minor:identifier:"), uuid, major, minor, objc.String(identifier))
	rv.Autorelease()
	return rv
}


// Retrieves data that you can use to advertise the current device as a beacon.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLBeaconRegion/peripheralData(withMeasuredPower:)
func (b_ BeaconRegion) PeripheralDataWithMeasuredPower(measuredPower foundation.INumber) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("peripheralDataWithMeasuredPower:"), measuredPower)
	return rv
}

// The beacon identity constraint that defines the beacon region.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLBeaconRegion/beaconIdentityConstraint
func (b_ BeaconRegion) BeaconIdentityConstraint() CLBeaconIdentityConstraint {
	rv := objc.Send[CLBeaconIdentityConstraint](b_.ID, objc.Sel("beaconIdentityConstraint"))
	return rv
}

// The major value from the beacon identity constraint that defines the beacon region.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLBeaconRegion/major
func (b_ BeaconRegion) Major() foundation.Number {
	rv := objc.Send[foundation.Number](b_.ID, objc.Sel("major"))
	return rv
}

// The minor value from the beacon identity constraint that defines the beacon region.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLBeaconRegion/minor
func (b_ BeaconRegion) Minor() foundation.Number {
	rv := objc.Send[foundation.Number](b_.ID, objc.Sel("minor"))
	return rv
}

// A Boolean value that indicates whether Core Location sends beacon notifications when the device’s display is on.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLBeaconRegion/notifyEntryStateOnDisplay
func (b_ BeaconRegion) NotifyEntryStateOnDisplay() bool {
	rv := objc.Send[bool](b_.ID, objc.Sel("notifyEntryStateOnDisplay"))
	return rv
}


// SetNotifyEntryStateOnDisplay sets the value of the notifyEntryStateOnDisplay property.
// A Boolean value that indicates whether Core Location sends beacon notifications when the device’s display is on.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLBeaconRegion/notifyEntryStateOnDisplay
func (b_ BeaconRegion) SetNotifyEntryStateOnDisplay(value bool) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setNotifyEntryStateOnDisplay:"), value)
}

// The unique ID of the beacons you’re targeting.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLBeaconRegion/proximityUUID
func (b_ BeaconRegion) ProximityUUID() foundation.UUID {
	rv := objc.Send[foundation.UUID](b_.ID, objc.Sel("proximityUUID"))
	return rv
}

// The UUID value from the beacon identity constraint that defines the beacon region.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLBeaconRegion/uuid
func (b_ BeaconRegion) UUID() foundation.UUID {
	rv := objc.Send[foundation.UUID](b_.ID, objc.Sel("UUID"))
	return rv
}


