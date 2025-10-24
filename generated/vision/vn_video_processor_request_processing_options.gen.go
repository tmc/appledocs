// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class VNVideoProcessorRequestProcessingOptions */


/* debug [class_header]: Header for VNVideoProcessorRequestProcessingOptions */
// The class instance for the [VideoProcessorRequestProcessingOptions] class.
var (
	VideoProcessorRequestProcessingOptionsClass     _VideoProcessorRequestProcessingOptionsClass
	VideoProcessorRequestProcessingOptionsClassOnce sync.Once
)

func getVideoProcessorRequestProcessingOptionsClass() _VideoProcessorRequestProcessingOptionsClass {
	VideoProcessorRequestProcessingOptionsClassOnce.Do(func() {
		VideoProcessorRequestProcessingOptionsClass = _VideoProcessorRequestProcessingOptionsClass{objc.GetClass("VNVideoProcessorRequestProcessingOptions")}
	})
	return VideoProcessorRequestProcessingOptionsClass
}

type _VideoProcessorRequestProcessingOptionsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for VideoProcessorRequestProcessingOptions */
// An interface definition for the [VideoProcessorRequestProcessingOptions] class.
type IVideoProcessorRequestProcessingOptions interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for VideoProcessorRequestProcessingOptions */
	// properties:
	Cadence() IVNVideoProcessorCadence
	SetCadence(value IVNVideoProcessorCadence)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for VideoProcessorRequestProcessingOptions */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for VideoProcessorRequestProcessingOptions */
// Alloc allocates a new instance without initialization.
func (vc _VideoProcessorRequestProcessingOptionsClass) Alloc() VideoProcessorRequestProcessingOptions {
	rv := objc.Send[VideoProcessorRequestProcessingOptions](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (vc _VideoProcessorRequestProcessingOptionsClass) New() VideoProcessorRequestProcessingOptions {
	rv := objc.Send[VideoProcessorRequestProcessingOptions](objc.ID(vc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (v_ VideoProcessorRequestProcessingOptions) Init() VideoProcessorRequestProcessingOptions {
	rv := objc.Send[VideoProcessorRequestProcessingOptions](v_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v_ VideoProcessorRequestProcessingOptions) Autorelease() VideoProcessorRequestProcessingOptions {
	rv := objc.Send[VideoProcessorRequestProcessingOptions](v_.ID, objc.Sel("autorelease"))
	return rv
}

// NewVideoProcessorRequestProcessingOptions creates a new VideoProcessorRequestProcessingOptions instance.
func NewVideoProcessorRequestProcessingOptions() VideoProcessorRequestProcessingOptions {
	return getVideoProcessorRequestProcessingOptionsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for VideoProcessorRequestProcessingOptions */
// An object that defines a video processor’s configuration options.


// An object that defines a video processor’s configuration options.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNVideoProcessor/RequestProcessingOptions
type VideoProcessorRequestProcessingOptions struct {
	objectivec.Object
}

// VideoProcessorRequestProcessingOptionsFrom constructs a [VideoProcessorRequestProcessingOptions] from an unsafe.Pointer.
//
// An object that defines a video processor’s configuration options.
func VideoProcessorRequestProcessingOptionsFrom(ptr unsafe.Pointer) VideoProcessorRequestProcessingOptions {
	return VideoProcessorRequestProcessingOptions{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for VideoProcessorRequestProcessingOptions *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for VideoProcessorRequestProcessingOptions */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for VideoProcessorRequestProcessingOptions */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for VideoProcessorRequestProcessingOptions */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for VideoProcessorRequestProcessingOptions */

// The cadence the video processor maintains to process the request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNVideoProcessor/RequestProcessingOptions/cadence
func (v_ VideoProcessorRequestProcessingOptions) Cadence() IVNVideoProcessorCadence {
	rv := objc.Send[VideoProcessorCadence](v_.ID, objc.Sel("cadence"))
	return rv
}/* debug [instance_properties/getter]: cadence */


// The cadence the video processor maintains to process the request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNVideoProcessor/RequestProcessingOptions/cadence
func (v_ VideoProcessorRequestProcessingOptions) SetCadence(value IVNVideoProcessorCadence) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setCadence:"), value)
}/* debug [instance_properties/setter]: cadence */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class VNVideoProcessorRequestProcessingOptions */



