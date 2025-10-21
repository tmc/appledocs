// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [VideoComposition] class.
var (
	VideoCompositionClass     _VideoCompositionClass
	VideoCompositionClassOnce sync.Once
)

func getVideoCompositionClass() _VideoCompositionClass {
	VideoCompositionClassOnce.Do(func() {
		VideoCompositionClass = _VideoCompositionClass{objc.GetClass("AVVideoComposition")}
	})
	return VideoCompositionClass
}

type _VideoCompositionClass struct {
	class objc.Class
}

// An interface definition for the [VideoComposition] class.
type IVideoComposition interface {
	objectivec.IObject
	IsValidForAssetTimeRangeValidationDelegate(asset unsafe.Pointer, timeRange unsafe.Pointer, validationDelegate objc.ID) bool
}

// An object that describes how to compose video frames at particular points in time.
//
// If you use the built-in video compositor, the instructions a video composition contain can specify a spatial transformation, an opacity value, and a cropping rectangle for each video source. This values can vary over time by applying linear ramping functions. You can create a custom video compositor by implementing the protocol. The system provides the custom video compositor with pixel buffers for each of its video sources during playback, and can perform arbitrary graphical operations on them to produce visual output.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVVideoComposition
type VideoComposition struct {
	objectivec.Object
}

// VideoCompositionFrom constructs a [VideoComposition] from an unsafe.Pointer.
//
// An object that describes how to compose video frames at particular points in time.
func VideoCompositionFrom(ptr unsafe.Pointer) VideoComposition {
	return VideoComposition{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (vc _VideoCompositionClass) Alloc() VideoComposition {
	rv := objc.Send[VideoComposition](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (vc _VideoCompositionClass) New() VideoComposition {
	rv := objc.Send[VideoComposition](objc.ID(vc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (v_ VideoComposition) Init() VideoComposition {
	rv := objc.Send[VideoComposition](v_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v_ VideoComposition) Autorelease() VideoComposition {
	rv := objc.Send[VideoComposition](v_.ID, objc.Sel("autorelease"))
	return rv
}

// NewVideoComposition creates a new VideoComposition instance.
func NewVideoComposition() VideoComposition {
	return getVideoCompositionClass().New()
}


// Indicates whether the time ranges of the composition’s instructions conform to validation requirements.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVVideoComposition/isValid(for:timeRange:validationDelegate:)
func (v_ VideoComposition) IsValidForAssetTimeRangeValidationDelegate(asset unsafe.Pointer, timeRange unsafe.Pointer, validationDelegate objc.ID) bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("isValidForAsset:timeRange:validationDelegate:"), asset, timeRange, validationDelegate)
	return rv
}

// A custom compositor class to use.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVVideoComposition/customVideoCompositorClass
func (v_ VideoComposition) CustomVideoCompositorClass() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID, objc.Sel("customVideoCompositorClass"))
	return rv
}

// A time interval for which the video composition should render composed video frames.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVVideoComposition/frameDuration
func (v_ VideoComposition) FrameDuration() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID, objc.Sel("frameDuration"))
	return rv
}

// The video composition instructions.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVVideoComposition/instructions
func (v_ VideoComposition) Instructions() []objc.ID {
	rv := objc.Send[[]objc.ID](v_.ID, objc.Sel("instructions"))
	return rv
}

// The scale at which the video composition should render.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVVideoComposition/renderScale
func (v_ VideoComposition) RenderScale() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID, objc.Sel("renderScale"))
	return rv
}



