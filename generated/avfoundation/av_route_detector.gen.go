// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class AVRouteDetector */


/* debug [class_header]: Header for AVRouteDetector */
// The class instance for the [RouteDetector] class.
var (
	RouteDetectorClass     _RouteDetectorClass
	RouteDetectorClassOnce sync.Once
)

func getRouteDetectorClass() _RouteDetectorClass {
	RouteDetectorClassOnce.Do(func() {
		RouteDetectorClass = _RouteDetectorClass{objc.GetClass("AVRouteDetector")}
	})
	return RouteDetectorClass
}

type _RouteDetectorClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for RouteDetector */
// An interface definition for the [RouteDetector] class.
type IRouteDetector interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for RouteDetector */
	// properties:
	RouteDetectionEnabled() bool
	SetRouteDetectionEnabled(value bool)
	MultipleRoutesDetected() bool
	IsRouteDetectionEnabled() bool
	SetIsRouteDetectionEnabled(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for RouteDetector */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for RouteDetector */
// Alloc allocates a new instance without initialization.
func (rc _RouteDetectorClass) Alloc() RouteDetector {
	rv := objc.Send[RouteDetector](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (rc _RouteDetectorClass) New() RouteDetector {
	rv := objc.Send[RouteDetector](objc.ID(rc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (r_ RouteDetector) Init() RouteDetector {
	rv := objc.Send[RouteDetector](r_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r_ RouteDetector) Autorelease() RouteDetector {
	rv := objc.Send[RouteDetector](r_.ID, objc.Sel("autorelease"))
	return rv
}

// NewRouteDetector creates a new RouteDetector instance.
func NewRouteDetector() RouteDetector {
	return getRouteDetectorClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for RouteDetector */
// An object that detects available media playback routes.
//
// If you enable route detection, the object reports whether it detects multiple playback routes. If it does, use to present the UI for the user to select an appropriate route.


// An object that detects available media playback routes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVRouteDetector
type RouteDetector struct {
	objectivec.Object
}

// RouteDetectorFrom constructs a [RouteDetector] from an unsafe.Pointer.
//
// An object that detects available media playback routes.
func RouteDetectorFrom(ptr unsafe.Pointer) RouteDetector {
	return RouteDetector{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for RouteDetector *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for RouteDetector */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for RouteDetector */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for RouteDetector */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for RouteDetector */

// A Boolean value that indicates whether route detection is in an enabled state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVRouteDetector/isRouteDetectionEnabled
func (r_ RouteDetector) RouteDetectionEnabled() bool {
	rv := objc.Send[bool](r_.ID, objc.Sel("routeDetectionEnabled"))
	return rv
}/* debug [instance_properties/getter]: routeDetectionEnabled */


// A Boolean value that indicates whether route detection is in an enabled state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVRouteDetector/isRouteDetectionEnabled
func (r_ RouteDetector) SetRouteDetectionEnabled(value bool) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setRouteDetectionEnabled:"), value)
}/* debug [instance_properties/setter]: routeDetectionEnabled */


// A Boolean value that indicates whether the object detects more than one playback route.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVRouteDetector/multipleRoutesDetected
func (r_ RouteDetector) MultipleRoutesDetected() bool {
	rv := objc.Send[bool](r_.ID, objc.Sel("multipleRoutesDetected"))
	return rv
}/* debug [instance_properties/getter]: multipleRoutesDetected */


// A Boolean value that indicates whether route detection is in an enabled state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avroutedetector/isroutedetectionenabled
func (r_ RouteDetector) IsRouteDetectionEnabled() bool {
	rv := objc.Send[bool](r_.ID, objc.Sel("isRouteDetectionEnabled"))
	return rv
}/* debug [instance_properties/getter]: isRouteDetectionEnabled */


// A Boolean value that indicates whether route detection is in an enabled state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avroutedetector/isroutedetectionenabled
func (r_ RouteDetector) SetIsRouteDetectionEnabled(value bool) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setIsRouteDetectionEnabled:"), value)
}/* debug [instance_properties/setter]: isRouteDetectionEnabled */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVRouteDetector */


