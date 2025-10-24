// Code generated from Apple documentation for MapKit. DO NOT EDIT.

package mapkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/corefoundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MKMapSnapshot */


/* debug [class_header]: Header for MKMapSnapshot */
// The class instance for the [MKMapSnapshot] class.
var (
	MKMapSnapshotClass     _MKMapSnapshotClass
	MKMapSnapshotClassOnce sync.Once
)

func getMKMapSnapshotClass() _MKMapSnapshotClass {
	MKMapSnapshotClassOnce.Do(func() {
		MKMapSnapshotClass = _MKMapSnapshotClass{objc.GetClass("MKMapSnapshot")}
	})
	return MKMapSnapshotClass
}

type _MKMapSnapshotClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MKMapSnapshot */
// An interface definition for the [MKMapSnapshot] class.
type IMKMapSnapshot interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MKMapSnapshot */
	// properties:
	Appearance() appkit.Appearance
	Image() appkit.Image
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MKMapSnapshot */
	// methods:
	PointForCoordinate(coordinate LocationCoordinate2D /* not a class type */) corefoundation.CGPoint
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MKMapSnapshot */
// Alloc allocates a new instance without initialization.
func (mc _MKMapSnapshotClass) Alloc() MKMapSnapshot {
	rv := objc.Send[MKMapSnapshot](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MKMapSnapshotClass) New() MKMapSnapshot {
	rv := objc.Send[MKMapSnapshot](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MKMapSnapshot) Init() MKMapSnapshot {
	rv := objc.Send[MKMapSnapshot](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MKMapSnapshot) Autorelease() MKMapSnapshot {
	rv := objc.Send[MKMapSnapshot](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMKMapSnapshot creates a new MKMapSnapshot instance.
func NewMKMapSnapshot() MKMapSnapshot {
	return getMKMapSnapshotClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MKMapSnapshot */
// An image that a snapshotter object generates.
//
// You don’t create instances of this class directly. Instead, you use an object to capture the map contents asynchronously. An object contains the image that the snapshotter generates from the map contents. Snapshot images don’t include any custom overlays or annotations that your app adds to the map view. If you want your annotations and overlays to appear on the final image, you need to draw them yourself. To position those items correctly on the image, use the method of this class to translate the overlay or annotation coordinate value to an appropriate location inside the image’s coordinate space.


// An image that a snapshotter object generates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapSnapshotter/Snapshot
type MKMapSnapshot struct {
	objectivec.Object
}

// MKMapSnapshotFrom constructs a [MKMapSnapshot] from an unsafe.Pointer.
//
// An image that a snapshotter object generates.
func MKMapSnapshotFrom(ptr unsafe.Pointer) MKMapSnapshot {
	return MKMapSnapshot{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MKMapSnapshot *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MKMapSnapshot */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MKMapSnapshot */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MKMapSnapshot */

// Converts the specified map coordinate to a point in the coordinate space of the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapSnapshotter/Snapshot/point(for:)
func (m_ MKMapSnapshot) PointForCoordinate(coordinate LocationCoordinate2D /* not a class type */) corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](m_.ID, objc.Sel("pointForCoordinate:"), coordinate)
	return rv
}/* debug [instance_methods/method]: PointForCoordinate */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MKMapSnapshot */

// The visual style that MapKit uses when rendering the snapshot.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapSnapshotter/Snapshot/appearance
func (m_ MKMapSnapshot) Appearance() appkit.Appearance {
	rv := objc.Send[appkit.Appearance](m_.ID, objc.Sel("appearance"))
	return rv
}/* debug [instance_properties/getter]: appearance */


// The image of the map’s content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapSnapshotter/Snapshot/image
func (m_ MKMapSnapshot) Image() appkit.Image {
	rv := objc.Send[appkit.Image](m_.ID, objc.Sel("image"))
	return rv
}/* debug [instance_properties/getter]: image */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MKMapSnapshot */


