// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class AVVideoCompositionLayerInstruction */


/* debug [class_header]: Header for AVVideoCompositionLayerInstruction */
// The class instance for the [VideoCompositionLayerInstruction] class.
var (
	VideoCompositionLayerInstructionClass     _VideoCompositionLayerInstructionClass
	VideoCompositionLayerInstructionClassOnce sync.Once
)

func getVideoCompositionLayerInstructionClass() _VideoCompositionLayerInstructionClass {
	VideoCompositionLayerInstructionClassOnce.Do(func() {
		VideoCompositionLayerInstructionClass = _VideoCompositionLayerInstructionClass{objc.GetClass("AVVideoCompositionLayerInstruction")}
	})
	return VideoCompositionLayerInstructionClass
}

type _VideoCompositionLayerInstructionClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for VideoCompositionLayerInstruction */
// An interface definition for the [VideoCompositionLayerInstruction] class.
type IVideoCompositionLayerInstruction interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for VideoCompositionLayerInstruction */
	// properties:
	TrackID() PersistentTrackID /* not a class type */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for VideoCompositionLayerInstruction */
	// methods:
	GetCropRectangleRampForTimeStartCropRectangleEndCropRectangleTimeRange(time objc.IObject /* cross-framework: Time */, startCropRectangle corefoundation.CGRect, endCropRectangle corefoundation.CGRect, timeRange TimeRange /* not a class type */) bool
	GetOpacityRampForTimeStartOpacityEndOpacityTimeRange(time objc.IObject /* cross-framework: Time */, startOpacity objectivec.IObject, endOpacity objectivec.IObject, timeRange TimeRange /* not a class type */) bool
	GetTransformRampForTimeStartTransformEndTransformTimeRange(time objc.IObject /* cross-framework: Time */, startTransform corefoundation.CGAffineTransform, endTransform corefoundation.CGAffineTransform, timeRange TimeRange /* not a class type */) bool
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for VideoCompositionLayerInstruction */
// Alloc allocates a new instance without initialization.
func (vc _VideoCompositionLayerInstructionClass) Alloc() VideoCompositionLayerInstruction {
	rv := objc.Send[VideoCompositionLayerInstruction](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (vc _VideoCompositionLayerInstructionClass) New() VideoCompositionLayerInstruction {
	rv := objc.Send[VideoCompositionLayerInstruction](objc.ID(vc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (v_ VideoCompositionLayerInstruction) Init() VideoCompositionLayerInstruction {
	rv := objc.Send[VideoCompositionLayerInstruction](v_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v_ VideoCompositionLayerInstruction) Autorelease() VideoCompositionLayerInstruction {
	rv := objc.Send[VideoCompositionLayerInstruction](v_.ID, objc.Sel("autorelease"))
	return rv
}

// NewVideoCompositionLayerInstruction creates a new VideoCompositionLayerInstruction instance.
func NewVideoCompositionLayerInstruction() VideoCompositionLayerInstruction {
	return getVideoCompositionLayerInstructionClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for VideoCompositionLayerInstruction */
// An object used to modify the transform, cropping, and opacity ramps applied to a given track in a composition.


// An object used to modify the transform, cropping, and opacity ramps applied to a given track in a composition.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVVideoCompositionLayerInstruction
type VideoCompositionLayerInstruction struct {
	objectivec.Object
}

// VideoCompositionLayerInstructionFrom constructs a [VideoCompositionLayerInstruction] from an unsafe.Pointer.
//
// An object used to modify the transform, cropping, and opacity ramps applied to a given track in a composition.
func VideoCompositionLayerInstructionFrom(ptr unsafe.Pointer) VideoCompositionLayerInstruction {
	return VideoCompositionLayerInstruction{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for VideoCompositionLayerInstruction *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for VideoCompositionLayerInstruction */

// Pass-through initializer, for internal use in AVFoundation only
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVVideoCompositionLayerInstruction/videoCompositionLayerInstructionWithLayerInstruction:
func (vc _VideoCompositionLayerInstructionClass) VideoCompositionLayerInstructionWithLayerInstruction(instruction IAVVideoCompositionLayerInstruction) IVideoCompositionLayerInstruction {
	rv := objc.Send[VideoCompositionLayerInstruction](objc.ID(vc.class), objc.Sel("videoCompositionLayerInstructionWithLayerInstruction:"), instruction)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=VideoCompositionLayerInstructionWithLayerInstruction) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for VideoCompositionLayerInstruction */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for VideoCompositionLayerInstruction */

// Obtains the crop rectangle ramp that includes the specified time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVVideoCompositionLayerInstruction/getCropRectangleRamp(for:startCropRectangle:endCropRectangle:timeRange:)
func (v_ VideoCompositionLayerInstruction) GetCropRectangleRampForTimeStartCropRectangleEndCropRectangleTimeRange(time objc.IObject /* cross-framework: Time */, startCropRectangle corefoundation.CGRect, endCropRectangle corefoundation.CGRect, timeRange TimeRange /* not a class type */) bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("getCropRectangleRampForTime:startCropRectangle:endCropRectangle:timeRange:"), time, startCropRectangle, endCropRectangle, timeRange)
	return rv
}/* debug [instance_methods/method]: GetCropRectangleRampForTimeStartCropRectangleEndCropRectangleTimeRange */


// Obtains the opacity ramp that includes a specified time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVVideoCompositionLayerInstruction/getOpacityRamp(for:startOpacity:endOpacity:timeRange:)
func (v_ VideoCompositionLayerInstruction) GetOpacityRampForTimeStartOpacityEndOpacityTimeRange(time objc.IObject /* cross-framework: Time */, startOpacity objectivec.IObject, endOpacity objectivec.IObject, timeRange TimeRange /* not a class type */) bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("getOpacityRampForTime:startOpacity:endOpacity:timeRange:"), time, startOpacity, endOpacity, timeRange)
	return rv
}/* debug [instance_methods/method]: GetOpacityRampForTimeStartOpacityEndOpacityTimeRange */


// Obtains the transform ramp that includes a specified time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVVideoCompositionLayerInstruction/getTransformRamp(for:start:end:timeRange:)
func (v_ VideoCompositionLayerInstruction) GetTransformRampForTimeStartTransformEndTransformTimeRange(time objc.IObject /* cross-framework: Time */, startTransform corefoundation.CGAffineTransform, endTransform corefoundation.CGAffineTransform, timeRange TimeRange /* not a class type */) bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("getTransformRampForTime:startTransform:endTransform:timeRange:"), time, startTransform, endTransform, timeRange)
	return rv
}/* debug [instance_methods/method]: GetTransformRampForTimeStartTransformEndTransformTimeRange */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for VideoCompositionLayerInstruction */

// The track identifier of the source track to which the compositor will apply the instruction.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVVideoCompositionLayerInstruction/trackID
func (v_ VideoCompositionLayerInstruction) TrackID() PersistentTrackID /* not a class type */ {
	rv := objc.Send[PersistentTrackID](v_.ID, objc.Sel("trackID"))
	return rv
}/* debug [instance_properties/getter]: trackID */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVVideoCompositionLayerInstruction */



