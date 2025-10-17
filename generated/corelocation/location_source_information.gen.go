// Code generated from Apple documentation for CoreLocation. DO NOT EDIT.

package corelocation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [LocationSourceInformation] class.
var locationSourceInformationClass = _LocationSourceInformationClass{objc.GetClass("CLLocationSourceInformation")}

type _LocationSourceInformationClass struct {
	class objc.Class
}

// Information about the source that provides a location. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLLocationSourceInformation

type LocationSourceInformation struct {
	objectivec.Object
}

// LocationSourceInformationFrom constructs a [LocationSourceInformation] from an unsafe.Pointer.
//
// Information about the source that provides a location.
func LocationSourceInformationFrom(ptr unsafe.Pointer) LocationSourceInformation {
	return LocationSourceInformation{objectivec.Object{objc.ID(ptr)}}
}



