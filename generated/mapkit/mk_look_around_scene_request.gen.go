// Code generated from Apple documentation for MapKit. DO NOT EDIT.

package mapkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MKLookAroundSceneRequest */


/* debug [class_header]: Header for MKLookAroundSceneRequest */
// The class instance for the [MKLookAroundSceneRequest] class.
var (
	MKLookAroundSceneRequestClass     _MKLookAroundSceneRequestClass
	MKLookAroundSceneRequestClassOnce sync.Once
)

func getMKLookAroundSceneRequestClass() _MKLookAroundSceneRequestClass {
	MKLookAroundSceneRequestClassOnce.Do(func() {
		MKLookAroundSceneRequestClass = _MKLookAroundSceneRequestClass{objc.GetClass("MKLookAroundSceneRequest")}
	})
	return MKLookAroundSceneRequestClass
}

type _MKLookAroundSceneRequestClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MKLookAroundSceneRequest */
// An interface definition for the [MKLookAroundSceneRequest] class.
type IMKLookAroundSceneRequest interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MKLookAroundSceneRequest */
	// properties:
	Coordinate() LocationCoordinate2D /* not a class type */
	Cancelled() bool
	Loading() bool
	MapItem() IMKMapItem
	IsCancelled() bool
	SetIsCancelled(value bool)
	IsLoading() bool
	SetIsLoading(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MKLookAroundSceneRequest */
	// methods:
	Cancel()
	GetSceneWithCompletionHandler(completionHandler func(unsafe.Pointer, unsafe.Pointer))
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MKLookAroundSceneRequest */
// Alloc allocates a new instance without initialization.
func (mc _MKLookAroundSceneRequestClass) Alloc() MKLookAroundSceneRequest {
	rv := objc.Send[MKLookAroundSceneRequest](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MKLookAroundSceneRequestClass) New() MKLookAroundSceneRequest {
	rv := objc.Send[MKLookAroundSceneRequest](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MKLookAroundSceneRequest) Init() MKLookAroundSceneRequest {
	rv := objc.Send[MKLookAroundSceneRequest](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MKLookAroundSceneRequest) Autorelease() MKLookAroundSceneRequest {
	rv := objc.Send[MKLookAroundSceneRequest](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMKLookAroundSceneRequest creates a new MKLookAroundSceneRequest instance.
func NewMKLookAroundSceneRequest() MKLookAroundSceneRequest {
	return getMKLookAroundSceneRequestClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MKLookAroundSceneRequest */
// A class you use to request a LookAround scene at the location you specify.


// A class you use to request a LookAround scene at the location you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKLookAroundSceneRequest
type MKLookAroundSceneRequest struct {
	objectivec.Object
}

// MKLookAroundSceneRequestFrom constructs a [MKLookAroundSceneRequest] from an unsafe.Pointer.
//
// A class you use to request a LookAround scene at the location you specify.
func MKLookAroundSceneRequestFrom(ptr unsafe.Pointer) MKLookAroundSceneRequest {
	return MKLookAroundSceneRequest{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MKLookAroundSceneRequest */

// Creates a LookAround scene at the specified coordinates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKLookAroundSceneRequest/init(coordinate:)
func NewMKLookAroundSceneRequestWithCoordinate(coordinate LocationCoordinate2D /* not a class type */) MKLookAroundSceneRequest {
	instance := getMKLookAroundSceneRequestClass().Alloc()
	rv := objc.Send[MKLookAroundSceneRequest](instance.ID, objc.Sel("initWithCoordinate:"), coordinate)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMKLookAroundSceneRequestWithCoordinate */


// Creates a LookAround scene with the location described by the specified map item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKLookAroundSceneRequest/init(mapItem:)
func NewMKLookAroundSceneRequestWithMapItem(mapItem IMKMapItem) MKLookAroundSceneRequest {
	instance := getMKLookAroundSceneRequestClass().Alloc()
	rv := objc.Send[MKLookAroundSceneRequest](instance.ID, objc.Sel("initWithMapItem:"), mapItem)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMKLookAroundSceneRequestWithMapItem */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MKLookAroundSceneRequest */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MKLookAroundSceneRequest */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MKLookAroundSceneRequest */

// Cancels the pending scene request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKLookAroundSceneRequest/cancel()
func (m_ MKLookAroundSceneRequest) Cancel() {
	objc.Send[objc.ID](m_.ID, objc.Sel("cancel"))
}/* debug [instance_methods/method]: Cancel */


// Requests a LookAround scene and calls the specified completion handler.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKLookAroundSceneRequest/getSceneWithCompletionHandler(_:)
func (m_ MKLookAroundSceneRequest) GetSceneWithCompletionHandler(completionHandler func(unsafe.Pointer, unsafe.Pointer)) {
	objc.Send[objc.ID](m_.ID, objc.Sel("getSceneWithCompletionHandler:"), completionHandler)
}/* debug [instance_methods/method]: GetSceneWithCompletionHandler */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MKLookAroundSceneRequest */

// A coordinate value that describes the location of the LookAround scene.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKLookAroundSceneRequest/coordinate
func (m_ MKLookAroundSceneRequest) Coordinate() LocationCoordinate2D /* not a class type */ {
	rv := objc.Send[LocationCoordinate2D](m_.ID, objc.Sel("coordinate"))
	return rv
}/* debug [instance_properties/getter]: coordinate */


// A Boolean value that indicates if the cancellation of a scene request was successful.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKLookAroundSceneRequest/isCancelled
func (m_ MKLookAroundSceneRequest) Cancelled() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("cancelled"))
	return rv
}/* debug [instance_properties/getter]: cancelled */


// A Boolean value that indicates whether a scene request is loading.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKLookAroundSceneRequest/isLoading
func (m_ MKLookAroundSceneRequest) Loading() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("loading"))
	return rv
}/* debug [instance_properties/getter]: loading */


// A map item that describes the location of the LookAround scene.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKLookAroundSceneRequest/mapItem
func (m_ MKLookAroundSceneRequest) MapItem() IMKMapItem {
	rv := objc.Send[MKMapItem](m_.ID, objc.Sel("mapItem"))
	return rv
}/* debug [instance_properties/getter]: mapItem */


// A Boolean value that indicates if the cancellation of a scene request was successful.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mklookaroundscenerequest/iscancelled
func (m_ MKLookAroundSceneRequest) IsCancelled() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("isCancelled"))
	return rv
}/* debug [instance_properties/getter]: isCancelled */


// A Boolean value that indicates if the cancellation of a scene request was successful.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mklookaroundscenerequest/iscancelled
func (m_ MKLookAroundSceneRequest) SetIsCancelled(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIsCancelled:"), value)
}/* debug [instance_properties/setter]: isCancelled */


// A Boolean value that indicates whether a scene request is loading.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mklookaroundscenerequest/isloading
func (m_ MKLookAroundSceneRequest) IsLoading() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("isLoading"))
	return rv
}/* debug [instance_properties/getter]: isLoading */


// A Boolean value that indicates whether a scene request is loading.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mklookaroundscenerequest/isloading
func (m_ MKLookAroundSceneRequest) SetIsLoading(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIsLoading:"), value)
}/* debug [instance_properties/setter]: isLoading */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MKLookAroundSceneRequest */


