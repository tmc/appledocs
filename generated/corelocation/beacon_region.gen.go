// Code generated from Apple documentation for CoreLocation. DO NOT EDIT.

package corelocation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [BeaconRegion] class.
var beaconRegionClass = _BeaconRegionClass{objc.GetClass("CLBeaconRegion")}

type _BeaconRegionClass struct {
	class objc.Class
}

// A region for detecting the presence of iBeacon devices. [Full Topic]
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

// Retrieves data that you can use to advertise the current device as a beacon. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLBeaconRegion/peripheralData(withMeasuredPower:)
func (b_ BeaconRegion) PeripheralDataWithMeasuredPower(measuredPower unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("peripheralDataWithMeasuredPower:"), measuredPower)
	return rv
}


