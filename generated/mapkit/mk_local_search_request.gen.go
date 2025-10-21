// Code generated from Apple documentation for MapKit. DO NOT EDIT.

package mapkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [MKLocalSearchRequest] class.
type IMKLocalSearchRequest interface {
	objectivec.IObject
}

// The parameters to use when searching for points of interest on the map.
//
// You create an object when you want to search for map locations based on a natural language string. For example, if your interface allows the user to type in addresses, place the typed text in this object and pass it to an object to begin the search process. When specifying your search strings, include a map region to narrow the search results to the specified geographical area. When creating an MKLocalSearch.Request object yourself, set the property to an appropriate search string, as in the following example: If your app uses an object to implement autocomplete support for user-supplied search strings, initialize your search request using the search completion that the user selects. In that case, use the method instead of the method to initialize your search request object. The completion object automatically provides the value for the property.
//
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

// Alloc allocates a new instance without initialization.
func (mc _MKLocalSearchRequestClass) Alloc() MKLocalSearchRequest {
	rv := objc.Send[MKLocalSearchRequest](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


// Creates and returns a search request based on the specified search completion data.
//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKLocalSearch/Request/init(completion:)
func NewMKLocalSearchRequestWithCompletion(completion unsafe.Pointer) MKLocalSearchRequest {
	instance := getMKLocalSearchRequestClass().Alloc()
	rv := objc.Send[MKLocalSearchRequest](instance.ID, objc.Sel("initWithCompletion:"), completion)
	rv.Autorelease()
	return rv
}

// Initializes and returns a local search request based on the provided string.
//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKLocalSearch/Request/init(naturalLanguageQuery:)
func NewMKLocalSearchRequestWithNaturalLanguageQuery(naturalLanguageQuery string) MKLocalSearchRequest {
	instance := getMKLocalSearchRequestClass().Alloc()
	rv := objc.Send[MKLocalSearchRequest](instance.ID, objc.Sel("initWithNaturalLanguageQuery:"), objc.String(naturalLanguageQuery))
	rv.Autorelease()
	return rv
}

// Initializes and returns a local search request based on the provided string and region.
//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKLocalSearch/Request/init(naturalLanguageQuery:region:)
func NewMKLocalSearchRequestWithNaturalLanguageQueryRegion(naturalLanguageQuery string, region unsafe.Pointer) MKLocalSearchRequest {
	instance := getMKLocalSearchRequestClass().Alloc()
	rv := objc.Send[MKLocalSearchRequest](instance.ID, objc.Sel("initWithNaturalLanguageQuery:region:"), objc.String(naturalLanguageQuery), region)
	rv.Autorelease()
	return rv
}


// A filter that lists which address options to include or exclude in search results.
//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKLocalSearch/Request/addressFilter
func (m_ MKLocalSearchRequest) AddressFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("addressFilter"))
	return rv
}


// SetAddressFilter sets the value of the addressFilter property.
// A filter that lists which address options to include or exclude in search results.

//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKLocalSearch/Request/addressFilter
func (m_ MKLocalSearchRequest) SetAddressFilter(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAddressFilter:"), value)
}
// A string containing the desired search item.
//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKLocalSearch/Request/naturalLanguageQuery
func (m_ MKLocalSearchRequest) NaturalLanguageQuery() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("naturalLanguageQuery"))
	return rv
}


// SetNaturalLanguageQuery sets the value of the naturalLanguageQuery property.
// A string containing the desired search item.

//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKLocalSearch/Request/naturalLanguageQuery
func (m_ MKLocalSearchRequest) SetNaturalLanguageQuery(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNaturalLanguageQuery:"), value)
}
// A filter that lists point-of-interest categories to include or exclude in search results.
//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKLocalSearch/Request/pointOfInterestFilter
func (m_ MKLocalSearchRequest) PointOfInterestFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("pointOfInterestFilter"))
	return rv
}


// SetPointOfInterestFilter sets the value of the pointOfInterestFilter property.
// A filter that lists point-of-interest categories to include or exclude in search results.

//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKLocalSearch/Request/pointOfInterestFilter
func (m_ MKLocalSearchRequest) SetPointOfInterestFilter(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPointOfInterestFilter:"), value)
}
// A map region that provides a hint as to where to search.
//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKLocalSearch/Request/region
func (m_ MKLocalSearchRequest) Region() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("region"))
	return rv
}


// SetRegion sets the value of the region property.
// A map region that provides a hint as to where to search.

//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKLocalSearch/Request/region
func (m_ MKLocalSearchRequest) SetRegion(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRegion:"), value)
}
// A value that indicates the importance of the configured region.
//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKLocalSearch/Request/regionPriority
func (m_ MKLocalSearchRequest) RegionPriority() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("regionPriority"))
	return rv
}


// SetRegionPriority sets the value of the regionPriority property.
// A value that indicates the importance of the configured region.

//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKLocalSearch/Request/regionPriority
func (m_ MKLocalSearchRequest) SetRegionPriority(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRegionPriority:"), value)
}
// The types of items to include in the search results.
//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKLocalSearch/Request/resultTypes
func (m_ MKLocalSearchRequest) ResultTypes() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("resultTypes"))
	return rv
}


// SetResultTypes sets the value of the resultTypes property.
// The types of items to include in the search results.

//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKLocalSearch/Request/resultTypes
func (m_ MKLocalSearchRequest) SetResultTypes(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setResultTypes:"), value)
}

