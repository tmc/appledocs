// Code generated from Apple documentation for MapKit. DO NOT EDIT.

package mapkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MKMapSnapshotter */


/* debug [class_header]: Header for MKMapSnapshotter */
// The class instance for the [MKMapSnapshotter] class.
var (
	MKMapSnapshotterClass     _MKMapSnapshotterClass
	MKMapSnapshotterClassOnce sync.Once
)

func getMKMapSnapshotterClass() _MKMapSnapshotterClass {
	MKMapSnapshotterClassOnce.Do(func() {
		MKMapSnapshotterClass = _MKMapSnapshotterClass{objc.GetClass("MKMapSnapshotter")}
	})
	return MKMapSnapshotterClass
}

type _MKMapSnapshotterClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MKMapSnapshotter */
// An interface definition for the [MKMapSnapshotter] class.
type IMKMapSnapshotter interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MKMapSnapshotter */
	// properties:
	Loading() bool
	IsLoading() bool
	SetIsLoading(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MKMapSnapshotter */
	// methods:
	Cancel()
	StartWithCompletionHandler(completionHandler objectivec.IObject)
	StartWithQueueCompletionHandler(queue objectivec.IObject, completionHandler objectivec.IObject)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MKMapSnapshotter */
// Alloc allocates a new instance without initialization.
func (mc _MKMapSnapshotterClass) Alloc() MKMapSnapshotter {
	rv := objc.Send[MKMapSnapshotter](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MKMapSnapshotterClass) New() MKMapSnapshotter {
	rv := objc.Send[MKMapSnapshotter](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MKMapSnapshotter) Init() MKMapSnapshotter {
	rv := objc.Send[MKMapSnapshotter](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MKMapSnapshotter) Autorelease() MKMapSnapshotter {
	rv := objc.Send[MKMapSnapshotter](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMKMapSnapshotter creates a new MKMapSnapshotter instance.
func NewMKMapSnapshotter() MKMapSnapshotter {
	return getMKMapSnapshotterClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MKMapSnapshotter */
// A utility class for capturing a map and its content into an image.
//
// Use an object when you want to capture the system-provided map content, including the map tiles and imagery. The snapshotter object captures the best image possible by loading all of the available map tiles before capturing the image. Configure a snapshotter object using an object. The snapshot options specify the appearance of the map, including which portion of the map the snapshotter captures.


// A utility class for capturing a map and its content into an image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapSnapshotter
type MKMapSnapshotter struct {
	objectivec.Object
}

// MKMapSnapshotterFrom constructs a [MKMapSnapshotter] from an unsafe.Pointer.
//
// A utility class for capturing a map and its content into an image.
func MKMapSnapshotterFrom(ptr unsafe.Pointer) MKMapSnapshotter {
	return MKMapSnapshotter{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MKMapSnapshotter */

// Creates and returns a snapshotter object based on the specified options.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapSnapshotter/init(options:)
func NewMKMapSnapshotterWithOptions(options IMKMapSnapshotOptions) MKMapSnapshotter {
	instance := getMKMapSnapshotterClass().Alloc()
	rv := objc.Send[MKMapSnapshotter](instance.ID, objc.Sel("initWithOptions:"), options)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMKMapSnapshotterWithOptions */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MKMapSnapshotter */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MKMapSnapshotter */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MKMapSnapshotter */

// Cancels the request to create a snapshot.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapSnapshotter/cancel()
func (m_ MKMapSnapshotter) Cancel() {
	objc.Send[objc.ID](m_.ID, objc.Sel("cancel"))
}/* debug [instance_methods/method]: Cancel */


// Submits the request to create a snapshot and delivers the results to the specified block.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapSnapshotter/start(completionHandler:)
func (m_ MKMapSnapshotter) StartWithCompletionHandler(completionHandler objectivec.IObject) {
	objc.Send[objc.ID](m_.ID, objc.Sel("startWithCompletionHandler:"), completionHandler)
}/* debug [instance_methods/method]: StartWithCompletionHandler */


// Submits the request to create a snapshot and executes the resulting block on the specified queue.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapSnapshotter/start(with:completionHandler:)
func (m_ MKMapSnapshotter) StartWithQueueCompletionHandler(queue objectivec.IObject, completionHandler objectivec.IObject) {
	objc.Send[objc.ID](m_.ID, objc.Sel("startWithQueue:completionHandler:"), queue, completionHandler)
}/* debug [instance_methods/method]: StartWithQueueCompletionHandler */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MKMapSnapshotter */

// A Boolean value that indicates whether the snapshotter is generating an image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapSnapshotter/isLoading
func (m_ MKMapSnapshotter) Loading() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("loading"))
	return rv
}/* debug [instance_properties/getter]: loading */


// A Boolean value that indicates whether the snapshotter is generating an image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapsnapshotter/isloading
func (m_ MKMapSnapshotter) IsLoading() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("isLoading"))
	return rv
}/* debug [instance_properties/getter]: isLoading */


// A Boolean value that indicates whether the snapshotter is generating an image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapsnapshotter/isloading
func (m_ MKMapSnapshotter) SetIsLoading(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIsLoading:"), value)
}/* debug [instance_properties/setter]: isLoading */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MKMapSnapshotter */


