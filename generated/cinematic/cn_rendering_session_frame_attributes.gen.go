// Code generated from Apple documentation for Cinematic. DO NOT EDIT.

package cinematic

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/avfoundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class CNRenderingSessionFrameAttributes */


/* debug [class_header]: Header for CNRenderingSessionFrameAttributes */
// The class instance for the [CNRenderingSessionFrameAttributes] class.
var (
	CNRenderingSessionFrameAttributesClass     _CNRenderingSessionFrameAttributesClass
	CNRenderingSessionFrameAttributesClassOnce sync.Once
)

func getCNRenderingSessionFrameAttributesClass() _CNRenderingSessionFrameAttributesClass {
	CNRenderingSessionFrameAttributesClassOnce.Do(func() {
		CNRenderingSessionFrameAttributesClass = _CNRenderingSessionFrameAttributesClass{objc.GetClass("CNRenderingSessionFrameAttributes")}
	})
	return CNRenderingSessionFrameAttributesClass
}

type _CNRenderingSessionFrameAttributesClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CNRenderingSessionFrameAttributes */
// An interface definition for the [CNRenderingSessionFrameAttributes] class.
type ICNRenderingSessionFrameAttributes interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for CNRenderingSessionFrameAttributes */
	// properties:
	FNumber() float32
	SetFNumber(value float32)
	FocusDisparity() float32
	SetFocusDisparity(value float32)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CNRenderingSessionFrameAttributes */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CNRenderingSessionFrameAttributes */
// Alloc allocates a new instance without initialization.
func (cc _CNRenderingSessionFrameAttributesClass) Alloc() CNRenderingSessionFrameAttributes {
	rv := objc.Send[CNRenderingSessionFrameAttributes](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CNRenderingSessionFrameAttributesClass) New() CNRenderingSessionFrameAttributes {
	rv := objc.Send[CNRenderingSessionFrameAttributes](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNRenderingSessionFrameAttributes) Init() CNRenderingSessionFrameAttributes {
	rv := objc.Send[CNRenderingSessionFrameAttributes](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNRenderingSessionFrameAttributes) Autorelease() CNRenderingSessionFrameAttributes {
	rv := objc.Send[CNRenderingSessionFrameAttributes](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNRenderingSessionFrameAttributes creates a new CNRenderingSessionFrameAttributes instance.
func NewCNRenderingSessionFrameAttributes() CNRenderingSessionFrameAttributes {
	return getCNRenderingSessionFrameAttributesClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CNRenderingSessionFrameAttributes */
// Creates an object with the per frame attributes that control the appearance of a single frame of the Cinematic movie.
//
// The object exposes properties such as focus disparity and f-stop. It initializes these to the values that the original recorded movie used for that frame. However, you can change them before rendering to adjust focus and aperture.


// Creates an object with the per frame attributes that control the appearance of a single frame of the Cinematic movie.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNRenderingSessionFrameAttributes
type CNRenderingSessionFrameAttributes struct {
	objectivec.Object
}

// CNRenderingSessionFrameAttributesFrom constructs a [CNRenderingSessionFrameAttributes] from an unsafe.Pointer.
//
// Creates an object with the per frame attributes that control the appearance of a single frame of the Cinematic movie.
func CNRenderingSessionFrameAttributesFrom(ptr unsafe.Pointer) CNRenderingSessionFrameAttributes {
	return CNRenderingSessionFrameAttributes{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CNRenderingSessionFrameAttributes */

// Initializes the rendering frame attributes from a sample buffer read from a Cinematic metadata track.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNRenderingSessionFrameAttributes/initWithSampleBuffer:sessionAttributes:
func NewCNRenderingSessionFrameAttributesWithSampleBufferSessionAttributes(sampleBuffer SampleBufferRef /* not a class type */, sessionAttributes ICNRenderingSessionAttributes) CNRenderingSessionFrameAttributes {
	instance := getCNRenderingSessionFrameAttributesClass().Alloc()
	rv := objc.Send[CNRenderingSessionFrameAttributes](instance.ID, objc.Sel("initWithSampleBuffer:sessionAttributes:"), sampleBuffer, sessionAttributes)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCNRenderingSessionFrameAttributesWithSampleBufferSessionAttributes */


// Initializes the rendering frame attributes from a timed metadata group read from a Cinematic metadata track.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNRenderingSessionFrameAttributes/initWithTimedMetadataGroup:sessionAttributes:
func NewCNRenderingSessionFrameAttributesWithTimedMetadataGroupSessionAttributes(metadataGroup avfoundation.TimedMetadataGroup, sessionAttributes ICNRenderingSessionAttributes) CNRenderingSessionFrameAttributes {
	instance := getCNRenderingSessionFrameAttributesClass().Alloc()
	rv := objc.Send[CNRenderingSessionFrameAttributes](instance.ID, objc.Sel("initWithTimedMetadataGroup:sessionAttributes:"), metadataGroup, sessionAttributes)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCNRenderingSessionFrameAttributesWithTimedMetadataGroupSessionAttributes */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CNRenderingSessionFrameAttributes */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CNRenderingSessionFrameAttributes */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CNRenderingSessionFrameAttributes */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CNRenderingSessionFrameAttributes */

// The f-stop value that inversely affects the aperture used to render the Cinematic image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNRenderingSessionFrameAttributes/fNumber
func (c_ CNRenderingSessionFrameAttributes) FNumber() float32 {
	rv := objc.Send[float32](c_.ID, objc.Sel("fNumber"))
	return rv
}/* debug [instance_properties/getter]: fNumber */


// The f-stop value that inversely affects the aperture used to render the Cinematic image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNRenderingSessionFrameAttributes/fNumber
func (c_ CNRenderingSessionFrameAttributes) SetFNumber(value float32) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setFNumber:"), value)
}/* debug [instance_properties/setter]: fNumber */


// Represents the focus plane at which the rendered image should be in focus.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNRenderingSessionFrameAttributes/focusDisparity
func (c_ CNRenderingSessionFrameAttributes) FocusDisparity() float32 {
	rv := objc.Send[float32](c_.ID, objc.Sel("focusDisparity"))
	return rv
}/* debug [instance_properties/getter]: focusDisparity */


// Represents the focus plane at which the rendered image should be in focus.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNRenderingSessionFrameAttributes/focusDisparity
func (c_ CNRenderingSessionFrameAttributes) SetFocusDisparity(value float32) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setFocusDisparity:"), value)
}/* debug [instance_properties/setter]: focusDisparity */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CNRenderingSessionFrameAttributes */


