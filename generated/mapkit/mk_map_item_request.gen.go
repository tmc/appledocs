// Code generated from Apple documentation for MapKit. DO NOT EDIT.

package mapkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MKMapItemRequest */


/* debug [class_header]: Header for MKMapItemRequest */
// The class instance for the [MKMapItemRequest] class.
var (
	MKMapItemRequestClass     _MKMapItemRequestClass
	MKMapItemRequestClassOnce sync.Once
)

func getMKMapItemRequestClass() _MKMapItemRequestClass {
	MKMapItemRequestClassOnce.Do(func() {
		MKMapItemRequestClass = _MKMapItemRequestClass{objc.GetClass("MKMapItemRequest")}
	})
	return MKMapItemRequestClass
}

type _MKMapItemRequestClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MKMapItemRequest */
// An interface definition for the [MKMapItemRequest] class.
type IMKMapItemRequest interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MKMapItemRequest */
	// properties:
	Cancelled() bool
	Loading() bool
	MapItemIdentifier() IMKMapItemIdentifier
	Feature() objectivec.IObject
	SetFeature(value objectivec.IObject)
	IsCancelled() bool
	SetIsCancelled(value bool)
	IsLoading() bool
	SetIsLoading(value bool)
	MapFeature() objectivec.IObject
	SetMapFeature(value objectivec.IObject)
	PlaceDescriptor() objectivec.IObject
	SetPlaceDescriptor(value objectivec.IObject)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MKMapItemRequest */
	// methods:
	Cancel()
	GetMapItemWithCompletionHandler(completionHandler func(unsafe.Pointer, unsafe.Pointer))
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MKMapItemRequest */
// Alloc allocates a new instance without initialization.
func (mc _MKMapItemRequestClass) Alloc() MKMapItemRequest {
	rv := objc.Send[MKMapItemRequest](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MKMapItemRequestClass) New() MKMapItemRequest {
	rv := objc.Send[MKMapItemRequest](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MKMapItemRequest) Init() MKMapItemRequest {
	rv := objc.Send[MKMapItemRequest](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MKMapItemRequest) Autorelease() MKMapItemRequest {
	rv := objc.Send[MKMapItemRequest](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMKMapItemRequest creates a new MKMapItemRequest instance.
func NewMKMapItemRequest() MKMapItemRequest {
	return getMKMapItemRequestClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MKMapItemRequest */
// A utility class you use to request additional information about a map feature.


// A utility class you use to request additional information about a map feature.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapItemRequest
type MKMapItemRequest struct {
	objectivec.Object
}

// MKMapItemRequestFrom constructs a [MKMapItemRequest] from an unsafe.Pointer.
//
// A utility class you use to request additional information about a map feature.
func MKMapItemRequestFrom(ptr unsafe.Pointer) MKMapItemRequest {
	return MKMapItemRequest{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MKMapItemRequest */

// Creates a new map item request with the specified feature annotation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapItemRequest/init(mapFeatureAnnotation:)
func NewMKMapItemRequestWithMapFeatureAnnotation(mapFeatureAnnotation IMKMapFeatureAnnotation) MKMapItemRequest {
	instance := getMKMapItemRequestClass().Alloc()
	rv := objc.Send[MKMapItemRequest](instance.ID, objc.Sel("initWithMapFeatureAnnotation:"), mapFeatureAnnotation)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMKMapItemRequestWithMapFeatureAnnotation */


// Create a request with a map item identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapItemRequest/init(mapItemIdentifier:)
func NewMKMapItemRequestWithMapItemIdentifier(identifier IMKMapItemIdentifier) MKMapItemRequest {
	instance := getMKMapItemRequestClass().Alloc()
	rv := objc.Send[MKMapItemRequest](instance.ID, objc.Sel("initWithMapItemIdentifier:"), identifier)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMKMapItemRequestWithMapItemIdentifier */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MKMapItemRequest */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MKMapItemRequest */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MKMapItemRequest */

// Cancels an in-progress map item request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapItemRequest/cancel()
func (m_ MKMapItemRequest) Cancel() {
	objc.Send[objc.ID](m_.ID, objc.Sel("cancel"))
}/* debug [instance_methods/method]: Cancel */


// Requests a map item and calls the provided completion handler.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapItemRequest/getMapItem(completionHandler:)
func (m_ MKMapItemRequest) GetMapItemWithCompletionHandler(completionHandler func(unsafe.Pointer, unsafe.Pointer)) {
	objc.Send[objc.ID](m_.ID, objc.Sel("getMapItemWithCompletionHandler:"), completionHandler)
}/* debug [instance_methods/method]: GetMapItemWithCompletionHandler */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MKMapItemRequest */

// A Boolean value that indicates if the cancellation of the request was successful.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapItemRequest/isCancelled
func (m_ MKMapItemRequest) Cancelled() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("cancelled"))
	return rv
}/* debug [instance_properties/getter]: cancelled */


// A Boolean value that indicates if the request is loading.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapItemRequest/isLoading
func (m_ MKMapItemRequest) Loading() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("loading"))
	return rv
}/* debug [instance_properties/getter]: loading */


// The map item identifer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapItemRequest/mapItemIdentifier
func (m_ MKMapItemRequest) MapItemIdentifier() IMKMapItemIdentifier {
	rv := objc.Send[MKMapItemIdentifier](m_.ID, objc.Sel("mapItemIdentifier"))
	return rv
}/* debug [instance_properties/getter]: mapItemIdentifier */


// The map feature.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapitemrequest/feature
func (m_ MKMapItemRequest) Feature() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](m_.ID, objc.Sel("feature"))
	return rv
}/* debug [instance_properties/getter]: feature */


// The map feature.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapitemrequest/feature
func (m_ MKMapItemRequest) SetFeature(value objectivec.IObject) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setFeature:"), value)
}/* debug [instance_properties/setter]: feature */


// A Boolean value that indicates if the cancellation of the request was successful.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapitemrequest/iscancelled
func (m_ MKMapItemRequest) IsCancelled() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("isCancelled"))
	return rv
}/* debug [instance_properties/getter]: isCancelled */


// A Boolean value that indicates if the cancellation of the request was successful.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapitemrequest/iscancelled
func (m_ MKMapItemRequest) SetIsCancelled(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIsCancelled:"), value)
}/* debug [instance_properties/setter]: isCancelled */


// A Boolean value that indicates if the request is loading.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapitemrequest/isloading
func (m_ MKMapItemRequest) IsLoading() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("isLoading"))
	return rv
}/* debug [instance_properties/getter]: isLoading */


// A Boolean value that indicates if the request is loading.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapitemrequest/isloading
func (m_ MKMapItemRequest) SetIsLoading(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIsLoading:"), value)
}/* debug [instance_properties/setter]: isLoading */


// The map feature.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapitemrequest/mapfeature
func (m_ MKMapItemRequest) MapFeature() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](m_.ID, objc.Sel("mapFeature"))
	return rv
}/* debug [instance_properties/getter]: mapFeature */


// The map feature.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapitemrequest/mapfeature
func (m_ MKMapItemRequest) SetMapFeature(value objectivec.IObject) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMapFeature:"), value)
}/* debug [instance_properties/setter]: mapFeature */


// The place descriptor that contains information that’s helpful in uniquely identifying this place.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapitemrequest/placedescriptor
func (m_ MKMapItemRequest) PlaceDescriptor() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](m_.ID, objc.Sel("placeDescriptor"))
	return rv
}/* debug [instance_properties/getter]: placeDescriptor */


// The place descriptor that contains information that’s helpful in uniquely identifying this place.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapitemrequest/placedescriptor
func (m_ MKMapItemRequest) SetPlaceDescriptor(value objectivec.IObject) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPlaceDescriptor:"), value)
}/* debug [instance_properties/setter]: placeDescriptor */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MKMapItemRequest */


