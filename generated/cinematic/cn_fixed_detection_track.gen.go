// Code generated from Apple documentation for Cinematic. DO NOT EDIT.

package cinematic

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class CNFixedDetectionTrack */


/* debug [class_header]: Header for CNFixedDetectionTrack */
// The class instance for the [CNFixedDetectionTrack] class.
var (
	CNFixedDetectionTrackClass     _CNFixedDetectionTrackClass
	CNFixedDetectionTrackClassOnce sync.Once
)

func getCNFixedDetectionTrackClass() _CNFixedDetectionTrackClass {
	CNFixedDetectionTrackClassOnce.Do(func() {
		CNFixedDetectionTrackClass = _CNFixedDetectionTrackClass{objc.GetClass("CNFixedDetectionTrack")}
	})
	return CNFixedDetectionTrackClass
}

type _CNFixedDetectionTrackClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CNFixedDetectionTrack */
// An interface definition for the [CNFixedDetectionTrack] class.
type ICNFixedDetectionTrack interface {
	ICNDetectionTrack
	
/* debug [class_interface_properties]: Properties for CNFixedDetectionTrack */
	// properties:
	FocusDisparity() float32
	OriginalDetection() ICNDetection
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CNFixedDetectionTrack */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CNFixedDetectionTrack */
// Alloc allocates a new instance without initialization.
func (cc _CNFixedDetectionTrackClass) Alloc() CNFixedDetectionTrack {
	rv := objc.Send[CNFixedDetectionTrack](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CNFixedDetectionTrackClass) New() CNFixedDetectionTrack {
	rv := objc.Send[CNFixedDetectionTrack](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNFixedDetectionTrack) Init() CNFixedDetectionTrack {
	rv := objc.Send[CNFixedDetectionTrack](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNFixedDetectionTrack) Autorelease() CNFixedDetectionTrack {
	rv := objc.Send[CNFixedDetectionTrack](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNFixedDetectionTrack creates a new CNFixedDetectionTrack instance.
func NewCNFixedDetectionTrack() CNFixedDetectionTrack {
	return getCNFixedDetectionTrackClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CNFixedDetectionTrack */
// An object representing the fixed detection track.


// An object representing the fixed detection track.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNFixedDetectionTrack-5aei2
type CNFixedDetectionTrack struct {
	CNDetectionTrack
}

// CNFixedDetectionTrackFrom constructs a [CNFixedDetectionTrack] from an unsafe.Pointer.
//
// An object representing the fixed detection track.
func CNFixedDetectionTrackFrom(ptr unsafe.Pointer) CNFixedDetectionTrack {
	return CNFixedDetectionTrack{
		CNDetectionTrack: CNDetectionTrackFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CNFixedDetectionTrack */

// Creates a detection track with fixed focus at the given disparity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNFixedDetectionTrack-5aei2/initWithFocusDisparity:
func NewCNFixedDetectionTrackWithFocusDisparity(focusDisparity float32) CNFixedDetectionTrack {
	instance := getCNFixedDetectionTrackClass().Alloc()
	rv := objc.Send[CNFixedDetectionTrack](instance.ID, objc.Sel("initWithFocusDisparity:"), focusDisparity)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCNFixedDetectionTrackWithFocusDisparity */


// Creates a detection track with fixed focus at the disparity of an existing detection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNFixedDetectionTrack-5aei2/initWithOriginalDetection:
func NewCNFixedDetectionTrackWithOriginalDetection(originalDetection ICNDetection) CNFixedDetectionTrack {
	instance := getCNFixedDetectionTrackClass().Alloc()
	rv := objc.Send[CNFixedDetectionTrack](instance.ID, objc.Sel("initWithOriginalDetection:"), originalDetection)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCNFixedDetectionTrackWithOriginalDetection */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CNFixedDetectionTrack */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CNFixedDetectionTrack */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CNFixedDetectionTrack */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CNFixedDetectionTrack */

// The disparity to use in order to focus on the object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNFixedDetectionTrack-5aei2/focusDisparity
func (c_ CNFixedDetectionTrack) FocusDisparity() float32 {
	rv := objc.Send[float32](c_.ID, objc.Sel("focusDisparity"))
	return rv
}/* debug [instance_properties/getter]: focusDisparity */


// The original detection based on the fixed detection track.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNFixedDetectionTrack-5aei2/originalDetection
func (c_ CNFixedDetectionTrack) OriginalDetection() ICNDetection {
	rv := objc.Send[CNDetection](c_.ID, objc.Sel("originalDetection"))
	return rv
}/* debug [instance_properties/getter]: originalDetection */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CNFixedDetectionTrack */


