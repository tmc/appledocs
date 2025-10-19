// Code generated from Apple documentation for CoreLocation. DO NOT EDIT.

package corelocation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [Placemark] class.
var (
	placemarkClass     _PlacemarkClass
	placemarkClassOnce sync.Once
)

func getPlacemarkClass() _PlacemarkClass {
	placemarkClassOnce.Do(func() {
		placemarkClass = _PlacemarkClass{objc.GetClass("CLPlacemark")}
	})
	return placemarkClass
}

type _PlacemarkClass struct {
	class objc.Class
}

// An interface definition for the [Placemark] class.
type IPlacemark interface {
	objectivec.IObject
}

// A user-friendly description of a geographic coordinate, often containing the name of the place, its address, and other relevant information. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLPlacemark
type Placemark struct {
	objectivec.Object
}

// PlacemarkFrom constructs a [Placemark] from an unsafe.Pointer.
//
// A user-friendly description of a geographic coordinate, often containing the name of the place, its address, and other relevant information.
func PlacemarkFrom(ptr unsafe.Pointer) Placemark {
	return Placemark{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (pc _PlacemarkClass) Alloc() Placemark {
	rv := objc.Send[Placemark](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PlacemarkClass) New() Placemark {
	rv := objc.Send[Placemark](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ Placemark) Init() Placemark {
	rv := objc.Send[Placemark](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ Placemark) Autorelease() Placemark {
	rv := objc.Send[Placemark](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPlacemark creates a new Placemark instance.
func NewPlacemark() Placemark {
	return getPlacemarkClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLPlacemark/init(location:name:postalAddress:)
func NewPlacemarkWithLocationNamePostalAddress(location unsafe.Pointer, name string, postalAddress unsafe.Pointer) Placemark {
	rv := objc.Send[Placemark](objc.ID(getPlacemarkClass().class), objc.Sel("placemarkWithLocation:name:postalAddress:"), location, objc.String(name), postalAddress)
	return rv
}
// Initializes and returns a placemark object from another placemark object. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLPlacemark/init(placemark:)
func NewPlacemarkWithPlacemark(placemark unsafe.Pointer) Placemark {
	instance := getPlacemarkClass().Alloc()
	rv := objc.Send[Placemark](instance.ID, objc.Sel("initWithPlacemark:"), placemark)
	rv.Autorelease()
	return rv
}


//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLPlacemark/init(location:name:postalAddress:)
func (pc _PlacemarkClass) PlacemarkWithLocationNamePostalAddress(location unsafe.Pointer, name string, postalAddress unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(pc.class), objc.Sel("placemarkWithLocation:name:postalAddress:"), location, objc.String(name), postalAddress)
	return rv
}

