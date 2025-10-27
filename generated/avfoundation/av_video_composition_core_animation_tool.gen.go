// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [VideoCompositionCoreAnimationTool] class.
var (
	VideoCompositionCoreAnimationToolClass     _VideoCompositionCoreAnimationToolClass
	VideoCompositionCoreAnimationToolClassOnce sync.Once
)

func getVideoCompositionCoreAnimationToolClass() _VideoCompositionCoreAnimationToolClass {
	VideoCompositionCoreAnimationToolClassOnce.Do(func() {
		VideoCompositionCoreAnimationToolClass = _VideoCompositionCoreAnimationToolClass{objc.GetClass("AVVideoCompositionCoreAnimationTool")}
	})
	return VideoCompositionCoreAnimationToolClass
}

type _VideoCompositionCoreAnimationToolClass struct {
	class objc.Class
}





// An interface definition for the [VideoCompositionCoreAnimationTool] class.
type IVideoCompositionCoreAnimationTool interface {
	objectivec.IObject
	

	// properties:
	AVCoreAnimationBeginTimeAtZero() float64
	IsRemovedOnCompletion() bool
	SetIsRemovedOnCompletion(value bool)
	BeginTime() float64
	SetBeginTime(value float64)


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (vc _VideoCompositionCoreAnimationToolClass) Alloc() VideoCompositionCoreAnimationTool {
	rv := objc.Send[VideoCompositionCoreAnimationTool](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (vc _VideoCompositionCoreAnimationToolClass) New() VideoCompositionCoreAnimationTool {
	rv := objc.Send[VideoCompositionCoreAnimationTool](objc.ID(vc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (v_ VideoCompositionCoreAnimationTool) Init() VideoCompositionCoreAnimationTool {
	rv := objc.Send[VideoCompositionCoreAnimationTool](v_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v_ VideoCompositionCoreAnimationTool) Autorelease() VideoCompositionCoreAnimationTool {
	rv := objc.Send[VideoCompositionCoreAnimationTool](v_.ID, objc.Sel("autorelease"))
	return rv
}

// NewVideoCompositionCoreAnimationTool creates a new VideoCompositionCoreAnimationTool instance.
func NewVideoCompositionCoreAnimationTool() VideoCompositionCoreAnimationTool {
	return getVideoCompositionCoreAnimationToolClass().New()
}





// An object used to incorporate Core Animation into a video composition.
//
// Any animations will be interpreted on the video’s timeline, not real-time, so you should: Set animations’ property to rather than (which CoreAnimation replaces with ); Set to on animations so they are not automatically removed; Avoid using layers that are associated with objects.


// An object used to incorporate Core Animation into a video composition.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVVideoCompositionCoreAnimationTool
type VideoCompositionCoreAnimationTool struct {
	objectivec.Object
}

// VideoCompositionCoreAnimationToolFrom constructs a [VideoCompositionCoreAnimationTool] from an unsafe.Pointer.
//
// An object used to incorporate Core Animation into a video composition.
func VideoCompositionCoreAnimationToolFrom(ptr unsafe.Pointer) VideoCompositionCoreAnimationTool {
	return VideoCompositionCoreAnimationTool{objectivec.Object{objc.ID(ptr)}}
}






// Adds a Core Animation layer to the video composition.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVVideoCompositionCoreAnimationTool/init(additionalLayer:asTrackID:)
func NewVideoCompositionCoreAnimationToolWithAdditionalLayerAsTrackID(layer Layer, trackID PersistentTrackID /* not a class type */) VideoCompositionCoreAnimationTool {
	rv := objc.Send[VideoCompositionCoreAnimationTool](objc.ID(getVideoCompositionCoreAnimationToolClass().class), objc.Sel("videoCompositionCoreAnimationToolWithAdditionalLayer:asTrackID:"), layer, trackID)
	return rv
}


// Composes the composited video frame with a Core Animation layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVVideoCompositionCoreAnimationTool/init(postProcessingAsVideoLayer:in:)
func NewVideoCompositionCoreAnimationToolWithPostProcessingAsVideoLayerInLayer(videoLayer Layer, animationLayer Layer) VideoCompositionCoreAnimationTool {
	rv := objc.Send[VideoCompositionCoreAnimationTool](objc.ID(getVideoCompositionCoreAnimationToolClass().class), objc.Sel("videoCompositionCoreAnimationToolWithPostProcessingAsVideoLayer:inLayer:"), videoLayer, animationLayer)
	return rv
}


// Composes the composited video frames with the Core Animation layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVVideoCompositionCoreAnimationTool/init(postProcessingAsVideoLayers:in:)
func NewVideoCompositionCoreAnimationToolWithPostProcessingAsVideoLayersInLayer(videoLayers []Layer, animationLayer Layer) VideoCompositionCoreAnimationTool {
	rv := objc.Send[VideoCompositionCoreAnimationTool](objc.ID(getVideoCompositionCoreAnimationToolClass().class), objc.Sel("videoCompositionCoreAnimationToolWithPostProcessingAsVideoLayers:inLayer:"), videoLayers, animationLayer)
	return rv
}







// Adds a Core Animation layer to the video composition.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVVideoCompositionCoreAnimationTool/init(additionalLayer:asTrackID:)
func (vc _VideoCompositionCoreAnimationToolClass) VideoCompositionCoreAnimationToolWithAdditionalLayerAsTrackID(layer Layer, trackID PersistentTrackID /* not a class type */) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(vc.class), objc.Sel("videoCompositionCoreAnimationToolWithAdditionalLayer:asTrackID:"), layer, trackID)
	return rv
}


// Composes the composited video frame with a Core Animation layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVVideoCompositionCoreAnimationTool/init(postProcessingAsVideoLayer:in:)
func (vc _VideoCompositionCoreAnimationToolClass) VideoCompositionCoreAnimationToolWithPostProcessingAsVideoLayerInLayer(videoLayer Layer, animationLayer Layer) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(vc.class), objc.Sel("videoCompositionCoreAnimationToolWithPostProcessingAsVideoLayer:inLayer:"), videoLayer, animationLayer)
	return rv
}


// Composes the composited video frames with the Core Animation layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVVideoCompositionCoreAnimationTool/init(postProcessingAsVideoLayers:in:)
func (vc _VideoCompositionCoreAnimationToolClass) VideoCompositionCoreAnimationToolWithPostProcessingAsVideoLayersInLayer(videoLayers []Layer, animationLayer Layer) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(vc.class), objc.Sel("videoCompositionCoreAnimationToolWithPostProcessingAsVideoLayers:inLayer:"), videoLayers, animationLayer)
	return rv
}

















// A value that sets an animation begin time to
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcoreanimationbegintimeatzero
func (v_ VideoCompositionCoreAnimationTool) AVCoreAnimationBeginTimeAtZero() float64 {
	rv := objc.Send[objc.ID](v_.ID, objc.Sel("AVCoreAnimationBeginTimeAtZero"))
	return rv
}


// Determines if the animation is removed from the target layer’s animations upon completion.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAAnimation/isRemovedOnCompletion
func (v_ VideoCompositionCoreAnimationTool) IsRemovedOnCompletion() bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("isRemovedOnCompletion"))
	return rv
}


// Determines if the animation is removed from the target layer’s animations upon completion.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAAnimation/isRemovedOnCompletion
func (v_ VideoCompositionCoreAnimationTool) SetIsRemovedOnCompletion(value bool) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setIsRemovedOnCompletion:"), value)
}


// Specifies the begin time of the receiver in relation to its parent object, if applicable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAMediaTiming/beginTime
func (v_ VideoCompositionCoreAnimationTool) BeginTime() float64 {
	rv := objc.Send[objc.ID](v_.ID, objc.Sel("beginTime"))
	return rv
}


// Specifies the begin time of the receiver in relation to its parent object, if applicable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAMediaTiming/beginTime
func (v_ VideoCompositionCoreAnimationTool) SetBeginTime(value float64) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setBeginTime:"), value)
}







