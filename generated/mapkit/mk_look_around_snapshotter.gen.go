// Code generated from Apple documentation for MapKit. DO NOT EDIT.

package mapkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MKLookAroundSnapshotter */


/* debug [class_header]: Header for MKLookAroundSnapshotter */
// The class instance for the [MKLookAroundSnapshotter] class.
var (
	MKLookAroundSnapshotterClass     _MKLookAroundSnapshotterClass
	MKLookAroundSnapshotterClassOnce sync.Once
)

func getMKLookAroundSnapshotterClass() _MKLookAroundSnapshotterClass {
	MKLookAroundSnapshotterClassOnce.Do(func() {
		MKLookAroundSnapshotterClass = _MKLookAroundSnapshotterClass{objc.GetClass("MKLookAroundSnapshotter")}
	})
	return MKLookAroundSnapshotterClass
}

type _MKLookAroundSnapshotterClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MKLookAroundSnapshotter */
// An interface definition for the [MKLookAroundSnapshotter] class.
type IMKLookAroundSnapshotter interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MKLookAroundSnapshotter */
	// properties:
	Loading() bool
	IsLoading() bool
	SetIsLoading(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MKLookAroundSnapshotter */
	// methods:
	Cancel()
	GetSnapshotWithCompletionHandler(completionHandler func(unsafe.Pointer, unsafe.Pointer))
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MKLookAroundSnapshotter */
// Alloc allocates a new instance without initialization.
func (mc _MKLookAroundSnapshotterClass) Alloc() MKLookAroundSnapshotter {
	rv := objc.Send[MKLookAroundSnapshotter](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MKLookAroundSnapshotterClass) New() MKLookAroundSnapshotter {
	rv := objc.Send[MKLookAroundSnapshotter](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MKLookAroundSnapshotter) Init() MKLookAroundSnapshotter {
	rv := objc.Send[MKLookAroundSnapshotter](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MKLookAroundSnapshotter) Autorelease() MKLookAroundSnapshotter {
	rv := objc.Send[MKLookAroundSnapshotter](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMKLookAroundSnapshotter creates a new MKLookAroundSnapshotter instance.
func NewMKLookAroundSnapshotter() MKLookAroundSnapshotter {
	return getMKLookAroundSnapshotterClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MKLookAroundSnapshotter */
// A utility class that you use to create a static image from a LookAround scene.


// A utility class that you use to create a static image from a LookAround scene.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKLookAroundSnapshotter
type MKLookAroundSnapshotter struct {
	objectivec.Object
}

// MKLookAroundSnapshotterFrom constructs a [MKLookAroundSnapshotter] from an unsafe.Pointer.
//
// A utility class that you use to create a static image from a LookAround scene.
func MKLookAroundSnapshotterFrom(ptr unsafe.Pointer) MKLookAroundSnapshotter {
	return MKLookAroundSnapshotter{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MKLookAroundSnapshotter */

// Create a new snapshotter object with the scene and options you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKLookAroundSnapshotter/init(scene:options:)
func NewMKLookAroundSnapshotterWithSceneOptions(scene IMKLookAroundScene, options IMKLookAroundSnapshotOptions) MKLookAroundSnapshotter {
	instance := getMKLookAroundSnapshotterClass().Alloc()
	rv := objc.Send[MKLookAroundSnapshotter](instance.ID, objc.Sel("initWithScene:options:"), scene, options)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMKLookAroundSnapshotterWithSceneOptions */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MKLookAroundSnapshotter */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MKLookAroundSnapshotter */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MKLookAroundSnapshotter */

// Cancels an in-progress snapshot request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKLookAroundSnapshotter/cancel()
func (m_ MKLookAroundSnapshotter) Cancel() {
	objc.Send[objc.ID](m_.ID, objc.Sel("cancel"))
}/* debug [instance_methods/method]: Cancel */


// Requests a new snapshot and calls the completion handler you provide.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKLookAroundSnapshotter/getSnapshotWithCompletionHandler(_:)
func (m_ MKLookAroundSnapshotter) GetSnapshotWithCompletionHandler(completionHandler func(unsafe.Pointer, unsafe.Pointer)) {
	objc.Send[objc.ID](m_.ID, objc.Sel("getSnapshotWithCompletionHandler:"), completionHandler)
}/* debug [instance_methods/method]: GetSnapshotWithCompletionHandler */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MKLookAroundSnapshotter */

// A Boolean value that indicates whether the snapshot request is loading.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKLookAroundSnapshotter/isLoading
func (m_ MKLookAroundSnapshotter) Loading() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("loading"))
	return rv
}/* debug [instance_properties/getter]: loading */


// A Boolean value that indicates whether the snapshot request is loading.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mklookaroundsnapshotter/isloading
func (m_ MKLookAroundSnapshotter) IsLoading() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("isLoading"))
	return rv
}/* debug [instance_properties/getter]: isLoading */


// A Boolean value that indicates whether the snapshot request is loading.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mklookaroundsnapshotter/isloading
func (m_ MKLookAroundSnapshotter) SetIsLoading(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIsLoading:"), value)
}/* debug [instance_properties/setter]: isLoading */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MKLookAroundSnapshotter */


