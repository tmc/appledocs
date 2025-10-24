// Code generated from Apple documentation for MapKit. DO NOT EDIT.

package mapkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MKLocalSearchResponse */


/* debug [class_header]: Header for MKLocalSearchResponse */
// The class instance for the [MKLocalSearchResponse] class.
var (
	MKLocalSearchResponseClass     _MKLocalSearchResponseClass
	MKLocalSearchResponseClassOnce sync.Once
)

func getMKLocalSearchResponseClass() _MKLocalSearchResponseClass {
	MKLocalSearchResponseClassOnce.Do(func() {
		MKLocalSearchResponseClass = _MKLocalSearchResponseClass{objc.GetClass("MKLocalSearchResponse")}
	})
	return MKLocalSearchResponseClass
}

type _MKLocalSearchResponseClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MKLocalSearchResponse */
// An interface definition for the [MKLocalSearchResponse] class.
type IMKLocalSearchResponse interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MKLocalSearchResponse */
	// properties:
	BoundingRegion() objc.IObject /* cross-framework: MKCoordinateRegion */
	MapItems() []MKMapItem
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MKLocalSearchResponse */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MKLocalSearchResponse */
// Alloc allocates a new instance without initialization.
func (mc _MKLocalSearchResponseClass) Alloc() MKLocalSearchResponse {
	rv := objc.Send[MKLocalSearchResponse](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MKLocalSearchResponseClass) New() MKLocalSearchResponse {
	rv := objc.Send[MKLocalSearchResponse](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MKLocalSearchResponse) Init() MKLocalSearchResponse {
	rv := objc.Send[MKLocalSearchResponse](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MKLocalSearchResponse) Autorelease() MKLocalSearchResponse {
	rv := objc.Send[MKLocalSearchResponse](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMKLocalSearchResponse creates a new MKLocalSearchResponse instance.
func NewMKLocalSearchResponse() MKLocalSearchResponse {
	return getMKLocalSearchResponseClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MKLocalSearchResponse */
// The results from a map-based search.
//
// You don’t create instances of this class directly. After initiating a map search using an object, MapKit passes an instance of this class to your completion handler.


// The results from a map-based search.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKLocalSearch/Response
type MKLocalSearchResponse struct {
	objectivec.Object
}

// MKLocalSearchResponseFrom constructs a [MKLocalSearchResponse] from an unsafe.Pointer.
//
// The results from a map-based search.
func MKLocalSearchResponseFrom(ptr unsafe.Pointer) MKLocalSearchResponse {
	return MKLocalSearchResponse{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MKLocalSearchResponse *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MKLocalSearchResponse */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MKLocalSearchResponse */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MKLocalSearchResponse */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MKLocalSearchResponse */

// The map region that encloses the returned search results.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKLocalSearch/Response/boundingRegion
func (m_ MKLocalSearchResponse) BoundingRegion() objc.IObject /* cross-framework: MKCoordinateRegion */ {
	rv := objc.Send[objc.ID](m_.ID, objc.Sel("boundingRegion"))
	return rv
}/* debug [instance_properties/getter]: boundingRegion */


// An array of map items representing the search results.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKLocalSearch/Response/mapItems
func (m_ MKLocalSearchResponse) MapItems() []MKMapItem {
	rv := objc.Send[[]MKMapItem](m_.ID, objc.Sel("mapItems"))
	return rv
}/* debug [instance_properties/getter]: mapItems */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MKLocalSearchResponse */



