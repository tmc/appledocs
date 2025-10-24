// Code generated from Apple documentation for MapKit. DO NOT EDIT.

package mapkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MKLocalSearchRequest */


/* debug [class_header]: Header for MKLocalSearchRequest */
// The class instance for the [MKLocalSearchRequest] class.
var (
	MKLocalSearchRequestClass     _MKLocalSearchRequestClass
	MKLocalSearchRequestClassOnce sync.Once
)

func getMKLocalSearchRequestClass() _MKLocalSearchRequestClass {
	MKLocalSearchRequestClassOnce.Do(func() {
		MKLocalSearchRequestClass = _MKLocalSearchRequestClass{objc.GetClass("MKLocalSearchRequest")}
	})
	return MKLocalSearchRequestClass
}

type _MKLocalSearchRequestClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MKLocalSearchRequest */
// An interface definition for the [MKLocalSearchRequest] class.
type IMKLocalSearchRequest interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MKLocalSearchRequest */
	// properties:
	AddressFilter() IMKAddressFilter
	SetAddressFilter(value IMKAddressFilter)
	NaturalLanguageQuery() objc.IObject /* cross-framework: NSString */
	SetNaturalLanguageQuery(value objc.IObject /* cross-framework: NSString */)
	PointOfInterestFilter() IMKPointOfInterestFilter
	SetPointOfInterestFilter(value IMKPointOfInterestFilter)
	Region() objc.IObject /* cross-framework: MKCoordinateRegion */
	SetRegion(value objc.IObject /* cross-framework: MKCoordinateRegion */)
	RegionPriority() MKLocalSearchRegionPriority
	SetRegionPriority(value MKLocalSearchRegionPriority)
	ResultTypes() MKLocalSearchResultType
	SetResultTypes(value MKLocalSearchResultType)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MKLocalSearchRequest */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MKLocalSearchRequest */
// Alloc allocates a new instance without initialization.
func (mc _MKLocalSearchRequestClass) Alloc() MKLocalSearchRequest {
	rv := objc.Send[MKLocalSearchRequest](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MKLocalSearchRequestClass) New() MKLocalSearchRequest {
	rv := objc.Send[MKLocalSearchRequest](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MKLocalSearchRequest) Init() MKLocalSearchRequest {
	rv := objc.Send[MKLocalSearchRequest](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MKLocalSearchRequest) Autorelease() MKLocalSearchRequest {
	rv := objc.Send[MKLocalSearchRequest](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMKLocalSearchRequest creates a new MKLocalSearchRequest instance.
func NewMKLocalSearchRequest() MKLocalSearchRequest {
	return getMKLocalSearchRequestClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MKLocalSearchRequest */
// The parameters to use when searching for points of interest on the map.
//
// You create an object when you want to search for map locations based on a natural language string. For example, if your interface allows the user to type in addresses, place the typed text in this object and pass it to an object to begin the search process. When specifying your search strings, include a map region to narrow the search results to the specified geographical area. When creating an MKLocalSearch.Request object yourself, set the property to an appropriate search string, as in the following example: If your app uses an object to implement autocomplete support for user-supplied search strings, initialize your search request using the search completion that the user selects. In that case, use the method instead of the method to initialize your search request object. The completion object automatically provides the value for the property.


// The parameters to use when searching for points of interest on the map.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKLocalSearch/Request
type MKLocalSearchRequest struct {
	objectivec.Object
}

// MKLocalSearchRequestFrom constructs a [MKLocalSearchRequest] from an unsafe.Pointer.
//
// The parameters to use when searching for points of interest on the map.
func MKLocalSearchRequestFrom(ptr unsafe.Pointer) MKLocalSearchRequest {
	return MKLocalSearchRequest{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MKLocalSearchRequest */

// Creates and returns a search request based on the specified search completion data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKLocalSearch/Request/init(completion:)
func NewMKLocalSearchRequestWithCompletion(completion IMKLocalSearchCompletion) MKLocalSearchRequest {
	instance := getMKLocalSearchRequestClass().Alloc()
	rv := objc.Send[MKLocalSearchRequest](instance.ID, objc.Sel("initWithCompletion:"), completion)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMKLocalSearchRequestWithCompletion */


// Initializes and returns a local search request based on the provided string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKLocalSearch/Request/init(naturalLanguageQuery:)
func NewMKLocalSearchRequestWithNaturalLanguageQuery(naturalLanguageQuery objc.IObject /* cross-framework: NSString */) MKLocalSearchRequest {
	instance := getMKLocalSearchRequestClass().Alloc()
	rv := objc.Send[MKLocalSearchRequest](instance.ID, objc.Sel("initWithNaturalLanguageQuery:"), naturalLanguageQuery)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMKLocalSearchRequestWithNaturalLanguageQuery */


// Initializes and returns a local search request based on the provided string and region.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKLocalSearch/Request/init(naturalLanguageQuery:region:)
func NewMKLocalSearchRequestWithNaturalLanguageQueryRegion(naturalLanguageQuery objc.IObject /* cross-framework: NSString */, region objc.IObject /* cross-framework: MKCoordinateRegion */) MKLocalSearchRequest {
	instance := getMKLocalSearchRequestClass().Alloc()
	rv := objc.Send[MKLocalSearchRequest](instance.ID, objc.Sel("initWithNaturalLanguageQuery:region:"), naturalLanguageQuery, region)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMKLocalSearchRequestWithNaturalLanguageQueryRegion */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MKLocalSearchRequest */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MKLocalSearchRequest */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MKLocalSearchRequest */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MKLocalSearchRequest */

// A filter that lists which address options to include or exclude in search results.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKLocalSearch/Request/addressFilter
func (m_ MKLocalSearchRequest) AddressFilter() IMKAddressFilter {
	rv := objc.Send[MKAddressFilter](m_.ID, objc.Sel("addressFilter"))
	return rv
}/* debug [instance_properties/getter]: addressFilter */


// A filter that lists which address options to include or exclude in search results.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKLocalSearch/Request/addressFilter
func (m_ MKLocalSearchRequest) SetAddressFilter(value IMKAddressFilter) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAddressFilter:"), value)
}/* debug [instance_properties/setter]: addressFilter */


// A string containing the desired search item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKLocalSearch/Request/naturalLanguageQuery
func (m_ MKLocalSearchRequest) NaturalLanguageQuery() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("naturalLanguageQuery"))
	return rv
}/* debug [instance_properties/getter]: naturalLanguageQuery */


// A string containing the desired search item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKLocalSearch/Request/naturalLanguageQuery
func (m_ MKLocalSearchRequest) SetNaturalLanguageQuery(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNaturalLanguageQuery:"), value)
}/* debug [instance_properties/setter]: naturalLanguageQuery */


// A filter that lists point-of-interest categories to include or exclude in search results.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKLocalSearch/Request/pointOfInterestFilter
func (m_ MKLocalSearchRequest) PointOfInterestFilter() IMKPointOfInterestFilter {
	rv := objc.Send[MKPointOfInterestFilter](m_.ID, objc.Sel("pointOfInterestFilter"))
	return rv
}/* debug [instance_properties/getter]: pointOfInterestFilter */


// A filter that lists point-of-interest categories to include or exclude in search results.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKLocalSearch/Request/pointOfInterestFilter
func (m_ MKLocalSearchRequest) SetPointOfInterestFilter(value IMKPointOfInterestFilter) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPointOfInterestFilter:"), value)
}/* debug [instance_properties/setter]: pointOfInterestFilter */


// A map region that provides a hint as to where to search.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKLocalSearch/Request/region
func (m_ MKLocalSearchRequest) Region() objc.IObject /* cross-framework: MKCoordinateRegion */ {
	rv := objc.Send[objc.ID](m_.ID, objc.Sel("region"))
	return rv
}/* debug [instance_properties/getter]: region */


// A map region that provides a hint as to where to search.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKLocalSearch/Request/region
func (m_ MKLocalSearchRequest) SetRegion(value objc.IObject /* cross-framework: MKCoordinateRegion */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRegion:"), value)
}/* debug [instance_properties/setter]: region */


// A value that indicates the importance of the configured region.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKLocalSearch/Request/regionPriority
func (m_ MKLocalSearchRequest) RegionPriority() MKLocalSearchRegionPriority {
	rv := objc.Send[MKLocalSearchRegionPriority](m_.ID, objc.Sel("regionPriority"))
	return rv
}/* debug [instance_properties/getter]: regionPriority */


// A value that indicates the importance of the configured region.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKLocalSearch/Request/regionPriority
func (m_ MKLocalSearchRequest) SetRegionPriority(value MKLocalSearchRegionPriority) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRegionPriority:"), value)
}/* debug [instance_properties/setter]: regionPriority */


// The types of items to include in the search results.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKLocalSearch/Request/resultTypes
func (m_ MKLocalSearchRequest) ResultTypes() MKLocalSearchResultType {
	rv := objc.Send[MKLocalSearchResultType](m_.ID, objc.Sel("resultTypes"))
	return rv
}/* debug [instance_properties/getter]: resultTypes */


// The types of items to include in the search results.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKLocalSearch/Request/resultTypes
func (m_ MKLocalSearchRequest) SetResultTypes(value MKLocalSearchResultType) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setResultTypes:"), value)
}/* debug [instance_properties/setter]: resultTypes */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MKLocalSearchRequest */


