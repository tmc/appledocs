// Code generated from Apple documentation for MapKit. DO NOT EDIT.

package mapkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corelocation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MKAddressRepresentations] class.
var (
	MKAddressRepresentationsClass     _MKAddressRepresentationsClass
	MKAddressRepresentationsClassOnce sync.Once
)

func getMKAddressRepresentationsClass() _MKAddressRepresentationsClass {
	MKAddressRepresentationsClassOnce.Do(func() {
		MKAddressRepresentationsClass = _MKAddressRepresentationsClass{objc.GetClass("MKAddressRepresentations")}
	})
	return MKAddressRepresentationsClass
}

type _MKAddressRepresentationsClass struct {
	class objc.Class
}

// An interface definition for the [MKAddressRepresentations] class.
type IMKAddressRepresentations interface {
	objectivec.IObject
	CityName() string
	SetCityName(value string)
	CityWithContext() string
	SetCityWithContext(value string)
	Region() corelocation.Region
	SetRegion(value corelocation.IRegion)
	RegionName() string
	SetRegionName(value string)
}

// A class that provides formatted address strings.


// A class that provides formatted address strings.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKAddressRepresentations
type MKAddressRepresentations struct {
	objectivec.Object
}

// MKAddressRepresentationsFrom constructs a [MKAddressRepresentations] from an unsafe.Pointer.
//
// A class that provides formatted address strings.
func MKAddressRepresentationsFrom(ptr unsafe.Pointer) MKAddressRepresentations {
	return MKAddressRepresentations{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MKAddressRepresentationsClass) Alloc() MKAddressRepresentations {
	rv := objc.Send[MKAddressRepresentations](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MKAddressRepresentationsClass) New() MKAddressRepresentations {
	rv := objc.Send[MKAddressRepresentations](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MKAddressRepresentations) Init() MKAddressRepresentations {
	rv := objc.Send[MKAddressRepresentations](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MKAddressRepresentations) Autorelease() MKAddressRepresentations {
	rv := objc.Send[MKAddressRepresentations](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMKAddressRepresentations creates a new MKAddressRepresentations instance.
func NewMKAddressRepresentations() MKAddressRepresentations {
	return getMKAddressRepresentationsClass().New()
}



// The name of the city.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkaddressrepresentations/cityname
func (m_ MKAddressRepresentations) CityName() string {
	rv := objc.Send[string](m_.ID, objc.Sel("cityName"))
	return rv
}


// The name of the city.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkaddressrepresentations/cityname
func (m_ MKAddressRepresentations) SetCityName(value string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCityName:"), objc.String(value))
}


// The city name along with the country name, to provide additional disambiguating context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkaddressrepresentations/citywithcontext
func (m_ MKAddressRepresentations) CityWithContext() string {
	rv := objc.Send[string](m_.ID, objc.Sel("cityWithContext"))
	return rv
}


// The city name along with the country name, to provide additional disambiguating context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkaddressrepresentations/citywithcontext
func (m_ MKAddressRepresentations) SetCityWithContext(value string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCityWithContext:"), objc.String(value))
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkaddressrepresentations/region
func (m_ MKAddressRepresentations) Region() corelocation.Region {
	rv := objc.Send[corelocation.Region](m_.ID, objc.Sel("region"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkaddressrepresentations/region
func (m_ MKAddressRepresentations) SetRegion(value corelocation.IRegion) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRegion:"), value)
}


// The region name, such as “United States”.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkaddressrepresentations/regionname
func (m_ MKAddressRepresentations) RegionName() string {
	rv := objc.Send[string](m_.ID, objc.Sel("regionName"))
	return rv
}


// The region name, such as “United States”.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkaddressrepresentations/regionname
func (m_ MKAddressRepresentations) SetRegionName(value string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRegionName:"), objc.String(value))
}



