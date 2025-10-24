// Code generated from Apple documentation for MapKit. DO NOT EDIT.

package mapkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MKMapCameraBoundary */


/* debug [class_header]: Header for MKMapCameraBoundary */
// The class instance for the [MKMapCameraBoundary] class.
var (
	MKMapCameraBoundaryClass     _MKMapCameraBoundaryClass
	MKMapCameraBoundaryClassOnce sync.Once
)

func getMKMapCameraBoundaryClass() _MKMapCameraBoundaryClass {
	MKMapCameraBoundaryClassOnce.Do(func() {
		MKMapCameraBoundaryClass = _MKMapCameraBoundaryClass{objc.GetClass("MKMapCameraBoundary")}
	})
	return MKMapCameraBoundaryClass
}

type _MKMapCameraBoundaryClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MKMapCameraBoundary */
// An interface definition for the [MKMapCameraBoundary] class.
type IMKMapCameraBoundary interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MKMapCameraBoundary */
	// properties:
	MapRect() objc.IObject /* cross-framework: MKMapRect */
	Region() objc.IObject /* cross-framework: MKCoordinateRegion */
	CameraBoundary() IMKMapCameraBoundary
	SetCameraBoundary(value IMKMapCameraBoundary)
	CameraZoomRange() IMKMapCameraZoomRange
	SetCameraZoomRange(value IMKMapCameraZoomRange)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MKMapCameraBoundary */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MKMapCameraBoundary */
// Alloc allocates a new instance without initialization.
func (mc _MKMapCameraBoundaryClass) Alloc() MKMapCameraBoundary {
	rv := objc.Send[MKMapCameraBoundary](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MKMapCameraBoundaryClass) New() MKMapCameraBoundary {
	rv := objc.Send[MKMapCameraBoundary](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MKMapCameraBoundary) Init() MKMapCameraBoundary {
	rv := objc.Send[MKMapCameraBoundary](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MKMapCameraBoundary) Autorelease() MKMapCameraBoundary {
	rv := objc.Send[MKMapCameraBoundary](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMKMapCameraBoundary creates a new MKMapCameraBoundary instance.
func NewMKMapCameraBoundary() MKMapCameraBoundary {
	return getMKMapCameraBoundaryClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MKMapCameraBoundary */
// A boundary of an area within which the map’s center needs to remain.
//
// The constraints of the camera boundary restrict the center point of your map.


// A boundary of an area within which the map’s center needs to remain.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapView/CameraBoundary-swift.class
type MKMapCameraBoundary struct {
	objectivec.Object
}

// MKMapCameraBoundaryFrom constructs a [MKMapCameraBoundary] from an unsafe.Pointer.
//
// A boundary of an area within which the map’s center needs to remain.
func MKMapCameraBoundaryFrom(ptr unsafe.Pointer) MKMapCameraBoundary {
	return MKMapCameraBoundary{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MKMapCameraBoundary */

// Creates a camera boundary using the provided coder.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapView/CameraBoundary-swift.class/init(coder:)
func NewMKMapCameraBoundaryWithCoder(coder foundation.Coder) MKMapCameraBoundary {
	instance := getMKMapCameraBoundaryClass().Alloc()
	rv := objc.Send[MKMapCameraBoundary](instance.ID, objc.Sel("initWithCoder:"), coder)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMKMapCameraBoundaryWithCoder */


// Creates a camera boundary using the provided coordinate region.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapView/CameraBoundary-swift.class/init(coordinateRegion:)
func NewMKMapCameraBoundaryWithCoordinateRegion(region objc.IObject /* cross-framework: MKCoordinateRegion */) MKMapCameraBoundary {
	instance := getMKMapCameraBoundaryClass().Alloc()
	rv := objc.Send[MKMapCameraBoundary](instance.ID, objc.Sel("initWithCoordinateRegion:"), region)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMKMapCameraBoundaryWithCoordinateRegion */


// Creates a camera boundary using the provided map rectangle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapView/CameraBoundary-swift.class/init(mapRect:)
func NewMKMapCameraBoundaryWithMapRect(mapRect objc.IObject /* cross-framework: MKMapRect */) MKMapCameraBoundary {
	instance := getMKMapCameraBoundaryClass().Alloc()
	rv := objc.Send[MKMapCameraBoundary](instance.ID, objc.Sel("initWithMapRect:"), mapRect)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMKMapCameraBoundaryWithMapRect */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MKMapCameraBoundary */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MKMapCameraBoundary */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MKMapCameraBoundary */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MKMapCameraBoundary */

// The map rectangle that describes the camera boundary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapView/CameraBoundary-swift.class/mapRect
func (m_ MKMapCameraBoundary) MapRect() objc.IObject /* cross-framework: MKMapRect */ {
	rv := objc.Send[objc.ID](m_.ID, objc.Sel("mapRect"))
	return rv
}/* debug [instance_properties/getter]: mapRect */


// The coordinate region that describes the camera boundary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapView/CameraBoundary-swift.class/region
func (m_ MKMapCameraBoundary) Region() objc.IObject /* cross-framework: MKCoordinateRegion */ {
	rv := objc.Send[objc.ID](m_.ID, objc.Sel("region"))
	return rv
}/* debug [instance_properties/getter]: region */


// The boundary of the area within which the map view’s center needs to remain.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapview/cameraboundary-swift.property
func (m_ MKMapCameraBoundary) CameraBoundary() IMKMapCameraBoundary {
	rv := objc.Send[MKMapCameraBoundary](m_.ID, objc.Sel("cameraBoundary"))
	return rv
}/* debug [instance_properties/getter]: cameraBoundary */


// The boundary of the area within which the map view’s center needs to remain.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapview/cameraboundary-swift.property
func (m_ MKMapCameraBoundary) SetCameraBoundary(value IMKMapCameraBoundary) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCameraBoundary:"), value)
}/* debug [instance_properties/setter]: cameraBoundary */


// The zoom range to apply to the map view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapview/camerazoomrange-swift.property
func (m_ MKMapCameraBoundary) CameraZoomRange() IMKMapCameraZoomRange {
	rv := objc.Send[MKMapCameraZoomRange](m_.ID, objc.Sel("cameraZoomRange"))
	return rv
}/* debug [instance_properties/getter]: cameraZoomRange */


// The zoom range to apply to the map view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapview/camerazoomrange-swift.property
func (m_ MKMapCameraBoundary) SetCameraZoomRange(value IMKMapCameraZoomRange) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCameraZoomRange:"), value)
}/* debug [instance_properties/setter]: cameraZoomRange */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MKMapCameraBoundary */


