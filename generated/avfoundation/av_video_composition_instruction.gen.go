// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [VideoCompositionInstruction] class.
var (
	VideoCompositionInstructionClass     _VideoCompositionInstructionClass
	VideoCompositionInstructionClassOnce sync.Once
)

func getVideoCompositionInstructionClass() _VideoCompositionInstructionClass {
	VideoCompositionInstructionClassOnce.Do(func() {
		VideoCompositionInstructionClass = _VideoCompositionInstructionClass{objc.GetClass("AVVideoCompositionInstruction")}
	})
	return VideoCompositionInstructionClass
}

type _VideoCompositionInstructionClass struct {
	class objc.Class
}





// An interface definition for the [VideoCompositionInstruction] class.
type IVideoCompositionInstruction interface {
	objectivec.IObject
	

	// properties:
	BackgroundColor() ColorRef /* not a class type */
	EnablePostProcessing() bool
	LayerInstructions() []VideoCompositionLayerInstruction
	PassthroughTrackID() PersistentTrackID /* not a class type */
	RequiredSourceSampleDataTrackIDs() []foundation.Number
	RequiredSourceTrackIDs() []foundation.Value
	TimeRange() TimeRange /* not a class type */
	Instructions() VideoCompositionInstructionProtocol /* not a class type */
	SetInstructions(value VideoCompositionInstructionProtocol /* not a class type */)


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (vc _VideoCompositionInstructionClass) Alloc() VideoCompositionInstruction {
	rv := objc.Send[VideoCompositionInstruction](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (vc _VideoCompositionInstructionClass) New() VideoCompositionInstruction {
	rv := objc.Send[VideoCompositionInstruction](objc.ID(vc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (v_ VideoCompositionInstruction) Init() VideoCompositionInstruction {
	rv := objc.Send[VideoCompositionInstruction](v_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v_ VideoCompositionInstruction) Autorelease() VideoCompositionInstruction {
	rv := objc.Send[VideoCompositionInstruction](v_.ID, objc.Sel("autorelease"))
	return rv
}

// NewVideoCompositionInstruction creates a new VideoCompositionInstruction instance.
func NewVideoCompositionInstruction() VideoCompositionInstruction {
	return getVideoCompositionInstructionClass().New()
}





// An operation that a compositor performs.
//
// An object maintains an array of to perform its composition.


// An operation that a compositor performs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVVideoCompositionInstruction-swift.class
type VideoCompositionInstruction struct {
	objectivec.Object
}

// VideoCompositionInstructionFrom constructs a [VideoCompositionInstruction] from an unsafe.Pointer.
//
// An operation that a compositor performs.
func VideoCompositionInstructionFrom(ptr unsafe.Pointer) VideoCompositionInstruction {
	return VideoCompositionInstruction{objectivec.Object{objc.ID(ptr)}}
}










// Pass-through initializer, for internal use in AVFoundation only
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVVideoCompositionInstruction-swift.class/videoCompositionInstructionWithInstruction:
func (vc _VideoCompositionInstructionClass) VideoCompositionInstructionWithInstruction(instruction IAVVideoCompositionInstruction) IVideoCompositionInstruction {
	rv := objc.Send[VideoCompositionInstruction](objc.ID(vc.class), objc.Sel("videoCompositionInstructionWithInstruction:"), instruction)
	return rv
}

















// The background color of the composition.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVVideoCompositionInstruction-swift.class/backgroundColor
func (v_ VideoCompositionInstruction) BackgroundColor() ColorRef /* not a class type */ {
	rv := objc.Send[ColorRef](v_.ID, objc.Sel("backgroundColor"))
	return rv
}


// A Boolean value that indicates whether the instruction requires post processing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVVideoCompositionInstruction-swift.class/enablePostProcessing
func (v_ VideoCompositionInstruction) EnablePostProcessing() bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("enablePostProcessing"))
	return rv
}


// Instructions that specify how to layer and compose video frames from source tracks.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVVideoCompositionInstruction-swift.class/layerInstructions
func (v_ VideoCompositionInstruction) LayerInstructions() []VideoCompositionLayerInstruction {
	rv := objc.Send[[]VideoCompositionLayerInstruction](v_.ID, objc.Sel("layerInstructions"))
	return rv
}


// The track identifier from an instruction source frame.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVVideoCompositionInstruction-swift.class/passthroughTrackID
func (v_ VideoCompositionInstruction) PassthroughTrackID() PersistentTrackID /* not a class type */ {
	rv := objc.Send[PersistentTrackID](v_.ID, objc.Sel("passthroughTrackID"))
	return rv
}


// The identifiers of source sample data tracks that the compositor requires to compose frames for the instruction.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVVideoCompositionInstruction-swift.class/requiredSourceSampleDataTrackIDs
func (v_ VideoCompositionInstruction) RequiredSourceSampleDataTrackIDs() []foundation.Number {
	rv := objc.Send[[]foundation.Number](v_.ID, objc.Sel("requiredSourceSampleDataTrackIDs"))
	return rv
}


// The identifiers of source video tracks that the compositor requires to compose frames for the instruction.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVVideoCompositionInstruction-swift.class/requiredSourceTrackIDs
func (v_ VideoCompositionInstruction) RequiredSourceTrackIDs() []foundation.Value {
	rv := objc.Send[[]foundation.Value](v_.ID, objc.Sel("requiredSourceTrackIDs"))
	return rv
}


// The time range to which the instruction applies.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVVideoCompositionInstruction-swift.class/timeRange
func (v_ VideoCompositionInstruction) TimeRange() TimeRange /* not a class type */ {
	rv := objc.Send[TimeRange](v_.ID, objc.Sel("timeRange"))
	return rv
}


// The video composition instructions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avvideocomposition/instructions
func (v_ VideoCompositionInstruction) Instructions() VideoCompositionInstructionProtocol /* not a class type */ {
	rv := objc.Send[VideoCompositionInstructionProtocol](v_.ID, objc.Sel("instructions"))
	return rv
}


// The video composition instructions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avvideocomposition/instructions
func (v_ VideoCompositionInstruction) SetInstructions(value VideoCompositionInstructionProtocol /* not a class type */) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setInstructions:"), value)
}








