// Code generated from Apple documentation for MapKit. DO NOT EDIT.

package mapkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MKLocalSearch */


/* debug [class_header]: Header for MKLocalSearch */
// The class instance for the [MKLocalSearch] class.
var (
	MKLocalSearchClass     _MKLocalSearchClass
	MKLocalSearchClassOnce sync.Once
)

func getMKLocalSearchClass() _MKLocalSearchClass {
	MKLocalSearchClassOnce.Do(func() {
		MKLocalSearchClass = _MKLocalSearchClass{objc.GetClass("MKLocalSearch")}
	})
	return MKLocalSearchClass
}

type _MKLocalSearchClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MKLocalSearch */
// An interface definition for the [MKLocalSearch] class.
type IMKLocalSearch interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MKLocalSearch */
	// properties:
	Searching() bool
	IsSearching() bool
	SetIsSearching(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MKLocalSearch */
	// methods:
	Cancel()
	StartWithCompletionHandler(completionHandler objectivec.IObject)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MKLocalSearch */
// Alloc allocates a new instance without initialization.
func (mc _MKLocalSearchClass) Alloc() MKLocalSearch {
	rv := objc.Send[MKLocalSearch](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MKLocalSearchClass) New() MKLocalSearch {
	rv := objc.Send[MKLocalSearch](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MKLocalSearch) Init() MKLocalSearch {
	rv := objc.Send[MKLocalSearch](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MKLocalSearch) Autorelease() MKLocalSearch {
	rv := objc.Send[MKLocalSearch](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMKLocalSearch creates a new MKLocalSearch instance.
func NewMKLocalSearch() MKLocalSearch {
	return getMKLocalSearchClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MKLocalSearch */
// A utility object for initiating map-based searches and processing the results.
//
// Use an object to execute a single search request. You might use this class to search for addresses or points of interest on the map. Upon completion of the request, the object delivers the results to the completion handler that you provide.


// A utility object for initiating map-based searches and processing the results.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKLocalSearch
type MKLocalSearch struct {
	objectivec.Object
}

// MKLocalSearchFrom constructs a [MKLocalSearch] from an unsafe.Pointer.
//
// A utility object for initiating map-based searches and processing the results.
func MKLocalSearchFrom(ptr unsafe.Pointer) MKLocalSearch {
	return MKLocalSearch{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MKLocalSearch */

// Creates and returns a search object for fetching points of interest.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKLocalSearch/init(request:)-9x8kn
func NewMKLocalSearchWithPointsOfInterestRequest(request IMKLocalPointsOfInterestRequest) MKLocalSearch {
	instance := getMKLocalSearchClass().Alloc()
	rv := objc.Send[MKLocalSearch](instance.ID, objc.Sel("initWithPointsOfInterestRequest:"), request)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMKLocalSearchWithPointsOfInterestRequest */


// Creates and returns a search object with the specified parameters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKLocalSearch/init(request:)-12tf0
func NewMKLocalSearchWithRequest(request IMKLocalSearchRequest) MKLocalSearch {
	instance := getMKLocalSearchClass().Alloc()
	rv := objc.Send[MKLocalSearch](instance.ID, objc.Sel("initWithRequest:"), request)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMKLocalSearchWithRequest */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MKLocalSearch */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MKLocalSearch */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MKLocalSearch */

// Cancels an in-progress search operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKLocalSearch/cancel()
func (m_ MKLocalSearch) Cancel() {
	objc.Send[objc.ID](m_.ID, objc.Sel("cancel"))
}/* debug [instance_methods/method]: Cancel */


// Starts the search and delivers the results to the specified completion handler.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKLocalSearch/start(completionHandler:)
func (m_ MKLocalSearch) StartWithCompletionHandler(completionHandler objectivec.IObject) {
	objc.Send[objc.ID](m_.ID, objc.Sel("startWithCompletionHandler:"), completionHandler)
}/* debug [instance_methods/method]: StartWithCompletionHandler */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MKLocalSearch */

// A Boolean value that indicates whether the search is in progress.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKLocalSearch/isSearching
func (m_ MKLocalSearch) Searching() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("searching"))
	return rv
}/* debug [instance_properties/getter]: searching */


// A Boolean value that indicates whether the search is in progress.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mklocalsearch/issearching
func (m_ MKLocalSearch) IsSearching() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("isSearching"))
	return rv
}/* debug [instance_properties/getter]: isSearching */


// A Boolean value that indicates whether the search is in progress.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mklocalsearch/issearching
func (m_ MKLocalSearch) SetIsSearching(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIsSearching:"), value)
}/* debug [instance_properties/setter]: isSearching */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MKLocalSearch */


