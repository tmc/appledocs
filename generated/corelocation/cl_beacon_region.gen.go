// Code generated from Apple documentation for CoreLocation. DO NOT EDIT.

package corelocation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objc"
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
	// properties:
	BeaconIdentityConstraint() ICLBeaconIdentityConstraint
	SetBeaconIdentityConstraint(value ICLBeaconIdentityConstraint)
	Major() objc.IObject /* cross-framework: NSNumber */
	SetMajor(value objc.IObject /* cross-framework: NSNumber */)
	Minor() objc.IObject /* cross-framework: NSNumber */
	SetMinor(value objc.IObject /* cross-framework: NSNumber */)
	NotifyEntryStateOnDisplay() bool
	SetNotifyEntryStateOnDisplay(value bool)
	ProximityUUID() objc.IObject /* cross-framework: UUID */
	SetProximityUUID(value objc.IObject /* cross-framework: UUID */)
	Uuid() objc.IObject /* cross-framework: UUID */
	SetUuid(value objc.IObject /* cross-framework: UUID */)
	// methods:
}

// A region for detecting the presence of iBeacon devices.
//
// A object defines a region that you use to detect Bluetooth beacons conforming to the iBeacon specification. In contrast to a that centers on a geographic location, a focuses on an iBeacon with specific identifying characteristics, which you provide. When a matching device comes in range, Core Location notifies your app. You monitor beacon regions in two ways. To detect when a beacon is in range, use the method of your location manager object. After detecting a beacon, call the method to determine the relative distance to that beacon. When detecting an iBeacon, you need to specify the , , and values that you programmed into the beacon hardware. You use the values to identify your beacons uniquely, and you can specify a subset of values to detect multiple beacons. The property is typically the same for all of the beacons in your installation. Use the and values to distinguish among different beacons in your installation. If you want to configure the current iOS device as a Bluetooth beacon, create a beacon region with the appropriate identifying information. You can then call the method of the region to get a dictionary that you can use to advertise the device with the Core Bluetooth framework. For more information about using that framework to advertise the device as a beacon, see . For information about how to detect beacons, see .

// A region for detecting the presence of iBeacon devices.
//
// [Full Topic]
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

// The beacon identity constraint that defines the beacon region.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corelocation/clbeaconregion/beaconidentityconstraint
func (b_ BeaconRegion) BeaconIdentityConstraint() ICLBeaconIdentityConstraint {
	rv := objc.Send[BeaconIdentityConstraint](b_.ID, objc.Sel("beaconIdentityConstraint"))
	return rv
}

// The beacon identity constraint that defines the beacon region.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corelocation/clbeaconregion/beaconidentityconstraint
func (b_ BeaconRegion) SetBeaconIdentityConstraint(value ICLBeaconIdentityConstraint) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setBeaconIdentityConstraint:"), value)
}

// The major value from the beacon identity constraint that defines the beacon region.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corelocation/clbeaconregion/major
func (b_ BeaconRegion) Major() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](b_.ID, objc.Sel("major"))
	return rv
}

// The major value from the beacon identity constraint that defines the beacon region.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corelocation/clbeaconregion/major
func (b_ BeaconRegion) SetMajor(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setMajor:"), value)
}

// The minor value from the beacon identity constraint that defines the beacon region.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corelocation/clbeaconregion/minor
func (b_ BeaconRegion) Minor() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](b_.ID, objc.Sel("minor"))
	return rv
}

// The minor value from the beacon identity constraint that defines the beacon region.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corelocation/clbeaconregion/minor
func (b_ BeaconRegion) SetMinor(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setMinor:"), value)
}

// A Boolean value that indicates whether Core Location sends beacon notifications when the device’s display is on.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corelocation/clbeaconregion/notifyentrystateondisplay
func (b_ BeaconRegion) NotifyEntryStateOnDisplay() bool {
	rv := objc.Send[bool](b_.ID, objc.Sel("notifyEntryStateOnDisplay"))
	return rv
}

// A Boolean value that indicates whether Core Location sends beacon notifications when the device’s display is on.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corelocation/clbeaconregion/notifyentrystateondisplay
func (b_ BeaconRegion) SetNotifyEntryStateOnDisplay(value bool) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setNotifyEntryStateOnDisplay:"), value)
}

// The unique ID of the beacons you’re targeting.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corelocation/clbeaconregion/proximityuuid
func (b_ BeaconRegion) ProximityUUID() objc.IObject /* cross-framework: UUID */ {
	rv := objc.Send[foundation.UUID](b_.ID, objc.Sel("proximityUUID"))
	return rv
}

// The unique ID of the beacons you’re targeting.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corelocation/clbeaconregion/proximityuuid
func (b_ BeaconRegion) SetProximityUUID(value objc.IObject /* cross-framework: UUID */) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setProximityUUID:"), value)
}

// The UUID value from the beacon identity constraint that defines the beacon region.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corelocation/clbeaconregion/uuid
func (b_ BeaconRegion) Uuid() objc.IObject /* cross-framework: UUID */ {
	rv := objc.Send[foundation.UUID](b_.ID, objc.Sel("uuid"))
	return rv
}

// The UUID value from the beacon identity constraint that defines the beacon region.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corelocation/clbeaconregion/uuid
func (b_ BeaconRegion) SetUuid(value objc.IObject /* cross-framework: UUID */) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setUuid:"), value)
}
