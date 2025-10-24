// Code generated from Apple documentation for Cinematic. DO NOT EDIT.

package cinematic

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class CNCustomDetectionTrack */


/* debug [class_header]: Header for CNCustomDetectionTrack */
// The class instance for the [CNCustomDetectionTrack] class.
var (
	CNCustomDetectionTrackClass     _CNCustomDetectionTrackClass
	CNCustomDetectionTrackClassOnce sync.Once
)

func getCNCustomDetectionTrackClass() _CNCustomDetectionTrackClass {
	CNCustomDetectionTrackClassOnce.Do(func() {
		CNCustomDetectionTrackClass = _CNCustomDetectionTrackClass{objc.GetClass("CNCustomDetectionTrack")}
	})
	return CNCustomDetectionTrackClass
}

type _CNCustomDetectionTrackClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CNCustomDetectionTrack */
// An interface definition for the [CNCustomDetectionTrack] class.
type ICNCustomDetectionTrack interface {
	ICNDetectionTrack
	
/* debug [class_interface_properties]: Properties for CNCustomDetectionTrack */
	// properties:
	AllDetections() []CNDetection
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CNCustomDetectionTrack */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CNCustomDetectionTrack */
// Alloc allocates a new instance without initialization.
func (cc _CNCustomDetectionTrackClass) Alloc() CNCustomDetectionTrack {
	rv := objc.Send[CNCustomDetectionTrack](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CNCustomDetectionTrackClass) New() CNCustomDetectionTrack {
	rv := objc.Send[CNCustomDetectionTrack](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNCustomDetectionTrack) Init() CNCustomDetectionTrack {
	rv := objc.Send[CNCustomDetectionTrack](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNCustomDetectionTrack) Autorelease() CNCustomDetectionTrack {
	rv := objc.Send[CNCustomDetectionTrack](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNCustomDetectionTrack creates a new CNCustomDetectionTrack instance.
func NewCNCustomDetectionTrack() CNCustomDetectionTrack {
	return getCNCustomDetectionTrackClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CNCustomDetectionTrack */
// An object representing a discrete detection track composed of individual detections.


// An object representing a discrete detection track composed of individual detections.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNCustomDetectionTrack-891hc
type CNCustomDetectionTrack struct {
	CNDetectionTrack
}

// CNCustomDetectionTrackFrom constructs a [CNCustomDetectionTrack] from an unsafe.Pointer.
//
// An object representing a discrete detection track composed of individual detections.
func CNCustomDetectionTrackFrom(ptr unsafe.Pointer) CNCustomDetectionTrack {
	return CNCustomDetectionTrack{
		CNDetectionTrack: CNDetectionTrackFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CNCustomDetectionTrack */

// Initializes a custom detection track with an array of detections, optionally applying smoothing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNCustomDetectionTrack-891hc/initWithDetections:smooth:
func NewCNCustomDetectionTrackWithDetectionsSmooth(detections []CNDetection, applySmoothing bool) CNCustomDetectionTrack {
	instance := getCNCustomDetectionTrackClass().Alloc()
	rv := objc.Send[CNCustomDetectionTrack](instance.ID, objc.Sel("initWithDetections:smooth:"), detections, applySmoothing)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCNCustomDetectionTrackWithDetectionsSmooth */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CNCustomDetectionTrack */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CNCustomDetectionTrack */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CNCustomDetectionTrack */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CNCustomDetectionTrack */

// All detected objects in the track.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNCustomDetectionTrack-891hc/allDetections
func (c_ CNCustomDetectionTrack) AllDetections() []CNDetection {
	rv := objc.Send[[]CNDetection](c_.ID, objc.Sel("allDetections"))
	return rv
}/* debug [instance_properties/getter]: allDetections */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CNCustomDetectionTrack */


