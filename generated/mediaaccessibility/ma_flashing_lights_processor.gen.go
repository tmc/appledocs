// Code generated from Apple documentation for MediaAccessibility. DO NOT EDIT.

package mediaaccessibility

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MAFlashingLightsProcessor */


/* debug [class_header]: Header for MAFlashingLightsProcessor */
// The class instance for the [MAFlashingLightsProcessor] class.
var (
	MAFlashingLightsProcessorClass     _MAFlashingLightsProcessorClass
	MAFlashingLightsProcessorClassOnce sync.Once
)

func getMAFlashingLightsProcessorClass() _MAFlashingLightsProcessorClass {
	MAFlashingLightsProcessorClassOnce.Do(func() {
		MAFlashingLightsProcessorClass = _MAFlashingLightsProcessorClass{objc.GetClass("MAFlashingLightsProcessor")}
	})
	return MAFlashingLightsProcessorClass
}

type _MAFlashingLightsProcessorClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MAFlashingLightsProcessor */
// An interface definition for the [MAFlashingLightsProcessor] class.
type IMAFlashingLightsProcessor interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MAFlashingLightsProcessor */
	// properties:
	KMADimFlashingLightsChangedNotification() foundation.String
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MAFlashingLightsProcessor */
	// methods:
	CanProcessSurface(surface SurfaceRef /* not a class type */) bool
	ProcessSurfaceOutSurfaceTimestampOptions(inSurface SurfaceRef /* not a class type */, outSurface SurfaceRef /* not a class type */, timestamp AbsoluteTime /* not a class type */, options foundation.IDictionary) IMAFlashingLightsProcessorResult
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MAFlashingLightsProcessor */
// Alloc allocates a new instance without initialization.
func (mc _MAFlashingLightsProcessorClass) Alloc() MAFlashingLightsProcessor {
	rv := objc.Send[MAFlashingLightsProcessor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MAFlashingLightsProcessorClass) New() MAFlashingLightsProcessor {
	rv := objc.Send[MAFlashingLightsProcessor](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MAFlashingLightsProcessor) Init() MAFlashingLightsProcessor {
	rv := objc.Send[MAFlashingLightsProcessor](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MAFlashingLightsProcessor) Autorelease() MAFlashingLightsProcessor {
	rv := objc.Send[MAFlashingLightsProcessor](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMAFlashingLightsProcessor creates a new MAFlashingLightsProcessor instance.
func NewMAFlashingLightsProcessor() MAFlashingLightsProcessor {
	return getMAFlashingLightsProcessorClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MAFlashingLightsProcessor */
// A class that processes a framebuffer object to detect and dim sequences of flashing lights.
//
// A device with the Dim Flashing Lights setting on automatically dims the brightness of flashing effect sequences when it detects them in video content. If your app performs custom video drawing instead of using APIs, you can use the class to detect and mitigate sequences of flashing effects in your video content. The following example shows how you might incorporate into code that uses APIs. For more information, see .


// A class that processes a framebuffer object to detect and dim sequences of flashing lights.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaAccessibility/MAFlashingLightsProcessor
type MAFlashingLightsProcessor struct {
	objectivec.Object
}

// MAFlashingLightsProcessorFrom constructs a [MAFlashingLightsProcessor] from an unsafe.Pointer.
//
// A class that processes a framebuffer object to detect and dim sequences of flashing lights.
func MAFlashingLightsProcessorFrom(ptr unsafe.Pointer) MAFlashingLightsProcessor {
	return MAFlashingLightsProcessor{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MAFlashingLightsProcessor *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MAFlashingLightsProcessor */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MAFlashingLightsProcessor */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MAFlashingLightsProcessor */

// Returns a Boolean value that indicates whether the flashing lights processor can process the content in the surface for sequences of flashing lights.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaAccessibility/MAFlashingLightsProcessor/canProcessSurface(_:)
func (m_ MAFlashingLightsProcessor) CanProcessSurface(surface SurfaceRef /* not a class type */) bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("canProcessSurface:"), surface)
	return rv
}/* debug [instance_methods/method]: CanProcessSurface */


// Processes a surface by analyzing pixels for sequences of flashing lights and mitigates them by dimming the content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaAccessibility/MAFlashingLightsProcessor/processSurface:outSurface:timestamp:options:
func (m_ MAFlashingLightsProcessor) ProcessSurfaceOutSurfaceTimestampOptions(inSurface SurfaceRef /* not a class type */, outSurface SurfaceRef /* not a class type */, timestamp AbsoluteTime /* not a class type */, options foundation.IDictionary) IMAFlashingLightsProcessorResult {
	rv := objc.Send[MAFlashingLightsProcessorResult](m_.ID, objc.Sel("processSurface:outSurface:timestamp:options:"), inSurface, outSurface, timestamp, options)
	return rv
}/* debug [instance_methods/method]: ProcessSurfaceOutSurfaceTimestampOptions */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MAFlashingLightsProcessor */

// A notification that posts when a person changes the flashing lights setting on the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaaccessibility/kmadimflashinglightschangednotification
func (m_ MAFlashingLightsProcessor) KMADimFlashingLightsChangedNotification() foundation.String {
	rv := objc.Send[foundation.String](m_.ID, objc.Sel("kMADimFlashingLightsChangedNotification"))
	return rv
}/* debug [instance_properties/getter]: kMADimFlashingLightsChangedNotification */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MAFlashingLightsProcessor */



