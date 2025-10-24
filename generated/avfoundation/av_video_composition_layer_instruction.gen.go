// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
	"github.com/tmc/appledocs/generated/objectivec"
)





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





// An interface definition for the [VideoCompositionLayerInstruction] class.
type IVideoCompositionLayerInstruction interface {
	objectivec.IObject
	

	// properties:
	TrackID() PersistentTrackID /* not a class type */


	

	// methods:
	GetCropRectangleRampForTimeStartCropRectangleEndCropRectangleTimeRange(time objc.IObject /* cross-framework: Time */, startCropRectangle corefoundation.CGRect, endCropRectangle corefoundation.CGRect, timeRange TimeRange /* not a class type */) bool
	GetOpacityRampForTimeStartOpacityEndOpacityTimeRange(time objc.IObject /* cross-framework: Time */, startOpacity objectivec.IObject, endOpacity objectivec.IObject, timeRange TimeRange /* not a class type */) bool
	GetTransformRampForTimeStartTransformEndTransformTimeRange(time objc.IObject /* cross-framework: Time */, startTransform corefoundation.CGAffineTransform, endTransform corefoundation.CGAffineTransform, timeRange TimeRange /* not a class type */) bool


}





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










// Pass-through initializer, for internal use in AVFoundation only
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVVideoCompositionLayerInstruction/videoCompositionLayerInstructionWithLayerInstruction:
func (vc _VideoCompositionLayerInstructionClass) VideoCompositionLayerInstructionWithLayerInstruction(instruction IAVVideoCompositionLayerInstruction) IVideoCompositionLayerInstruction {
	rv := objc.Send[VideoCompositionLayerInstruction](objc.ID(vc.class), objc.Sel("videoCompositionLayerInstructionWithLayerInstruction:"), instruction)
	return rv
}












// Obtains the crop rectangle ramp that includes the specified time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVVideoCompositionLayerInstruction/getCropRectangleRamp(for:startCropRectangle:endCropRectangle:timeRange:)
func (v_ VideoCompositionLayerInstruction) GetCropRectangleRampForTimeStartCropRectangleEndCropRectangleTimeRange(time objc.IObject /* cross-framework: Time */, startCropRectangle corefoundation.CGRect, endCropRectangle corefoundation.CGRect, timeRange TimeRange /* not a class type */) bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("getCropRectangleRampForTime:startCropRectangle:endCropRectangle:timeRange:"), time, startCropRectangle, endCropRectangle, timeRange)
	return rv
}


// Obtains the opacity ramp that includes a specified time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVVideoCompositionLayerInstruction/getOpacityRamp(for:startOpacity:endOpacity:timeRange:)
func (v_ VideoCompositionLayerInstruction) GetOpacityRampForTimeStartOpacityEndOpacityTimeRange(time objc.IObject /* cross-framework: Time */, startOpacity objectivec.IObject, endOpacity objectivec.IObject, timeRange TimeRange /* not a class type */) bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("getOpacityRampForTime:startOpacity:endOpacity:timeRange:"), time, startOpacity, endOpacity, timeRange)
	return rv
}


// Obtains the transform ramp that includes a specified time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVVideoCompositionLayerInstruction/getTransformRamp(for:start:end:timeRange:)
func (v_ VideoCompositionLayerInstruction) GetTransformRampForTimeStartTransformEndTransformTimeRange(time objc.IObject /* cross-framework: Time */, startTransform corefoundation.CGAffineTransform, endTransform corefoundation.CGAffineTransform, timeRange TimeRange /* not a class type */) bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("getTransformRampForTime:startTransform:endTransform:timeRange:"), time, startTransform, endTransform, timeRange)
	return rv
}







// The track identifier of the source track to which the compositor will apply the instruction.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVVideoCompositionLayerInstruction/trackID
func (v_ VideoCompositionLayerInstruction) TrackID() PersistentTrackID /* not a class type */ {
	rv := objc.Send[PersistentTrackID](v_.ID, objc.Sel("trackID"))
	return rv
}








