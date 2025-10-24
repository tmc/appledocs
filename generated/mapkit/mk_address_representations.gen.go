// Code generated from Apple documentation for MapKit. DO NOT EDIT.

package mapkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corelocation"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MKAddressRepresentations */


/* debug [class_header]: Header for MKAddressRepresentations */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MKAddressRepresentations */
// An interface definition for the [MKAddressRepresentations] class.
type IMKAddressRepresentations interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MKAddressRepresentations */
	// properties:
	CityName() objc.IObject /* cross-framework: NSString */
	CityWithContext() objc.IObject /* cross-framework: NSString */
	RegionCode() objc.IObject /* cross-framework: NSString */
	RegionName() objc.IObject /* cross-framework: NSString */
	Region() corelocation.Region
	SetRegion(value corelocation.Region)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MKAddressRepresentations */
	// methods:
	CityWithContextUsingStyle(style MKAddressRepresentationsContextStyle) foundation.String
	FullAddressIncludingRegionSingleLine(includingRegion bool, singleLine bool) foundation.String
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MKAddressRepresentations */
// Alloc allocates a new instance without initialization.
func (mc _MKAddressRepresentationsClass) Alloc() MKAddressRepresentations {
	rv := objc.Send[MKAddressRepresentations](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MKAddressRepresentations */
// A class that provides formatted address strings.
//
// Use this class to obtain formatted address strings for a place’s full address, city, or region.


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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MKAddressRepresentations *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MKAddressRepresentations */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MKAddressRepresentations */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MKAddressRepresentations */

// The city name and, optionally and if applicable, state and region to provide additional disambiguating context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKAddressRepresentations/cityWithContext(_:)
func (m_ MKAddressRepresentations) CityWithContextUsingStyle(style MKAddressRepresentationsContextStyle) foundation.String {
	rv := objc.Send[foundation.String](m_.ID, objc.Sel("cityWithContextUsingStyle:"), style)
	return rv
}/* debug [instance_methods/method]: CityWithContextUsingStyle */


// Returns the the location’s full address, optionally including the country or on a single link without line breaks.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKAddressRepresentations/fullAddress(includingRegion:singleLine:)
func (m_ MKAddressRepresentations) FullAddressIncludingRegionSingleLine(includingRegion bool, singleLine bool) foundation.String {
	rv := objc.Send[foundation.String](m_.ID, objc.Sel("fullAddressIncludingRegion:singleLine:"), includingRegion, singleLine)
	return rv
}/* debug [instance_methods/method]: FullAddressIncludingRegionSingleLine */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MKAddressRepresentations */

// The name of the city.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKAddressRepresentations/cityName
func (m_ MKAddressRepresentations) CityName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("cityName"))
	return rv
}/* debug [instance_properties/getter]: cityName */


// The city name along with the country name, to provide additional disambiguating context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKAddressRepresentations/cityWithContext
func (m_ MKAddressRepresentations) CityWithContext() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("cityWithContext"))
	return rv
}/* debug [instance_properties/getter]: cityWithContext */


// The region’s ISO 3166-2 region code, such as “US”.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKAddressRepresentations/regionCode
func (m_ MKAddressRepresentations) RegionCode() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("regionCode"))
	return rv
}/* debug [instance_properties/getter]: regionCode */


// The region name, such as “United States”.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKAddressRepresentations/regionName
func (m_ MKAddressRepresentations) RegionName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("regionName"))
	return rv
}/* debug [instance_properties/getter]: regionName */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkaddressrepresentations/region
func (m_ MKAddressRepresentations) Region() corelocation.Region {
	rv := objc.Send[corelocation.Region](m_.ID, objc.Sel("region"))
	return rv
}/* debug [instance_properties/getter]: region */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkaddressrepresentations/region
func (m_ MKAddressRepresentations) SetRegion(value corelocation.Region) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRegion:"), value)
}/* debug [instance_properties/setter]: region */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MKAddressRepresentations */



