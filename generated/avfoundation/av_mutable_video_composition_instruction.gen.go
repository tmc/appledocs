// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [MutableVideoCompositionInstruction] class.
var (
	MutableVideoCompositionInstructionClass     _MutableVideoCompositionInstructionClass
	MutableVideoCompositionInstructionClassOnce sync.Once
)

func getMutableVideoCompositionInstructionClass() _MutableVideoCompositionInstructionClass {
	MutableVideoCompositionInstructionClassOnce.Do(func() {
		MutableVideoCompositionInstructionClass = _MutableVideoCompositionInstructionClass{objc.GetClass("AVMutableVideoCompositionInstruction")}
	})
	return MutableVideoCompositionInstructionClass
}

type _MutableVideoCompositionInstructionClass struct {
	class objc.Class
}





// An interface definition for the [MutableVideoCompositionInstruction] class.
type IMutableVideoCompositionInstruction interface {
	IVideoCompositionInstruction
	

	// properties:
	BackgroundColor() ColorRef /* not a class type */
	SetBackgroundColor(value ColorRef /* not a class type */)
	EnablePostProcessing() bool
	SetEnablePostProcessing(value bool)
	LayerInstructions() []VideoCompositionLayerInstruction
	SetLayerInstructions(value []VideoCompositionLayerInstruction)
	RequiredSourceSampleDataTrackIDs() []foundation.Number
	SetRequiredSourceSampleDataTrackIDs(value []foundation.Number)
	TimeRange() objectivec.IObject
	SetTimeRange(value objectivec.IObject)
	Instructions() VideoCompositionInstructionProtocol /* not a class type */
	SetInstructions(value VideoCompositionInstructionProtocol /* not a class type */)


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (mc _MutableVideoCompositionInstructionClass) Alloc() MutableVideoCompositionInstruction {
	rv := objc.Send[MutableVideoCompositionInstruction](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MutableVideoCompositionInstructionClass) New() MutableVideoCompositionInstruction {
	rv := objc.Send[MutableVideoCompositionInstruction](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MutableVideoCompositionInstruction) Init() MutableVideoCompositionInstruction {
	rv := objc.Send[MutableVideoCompositionInstruction](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MutableVideoCompositionInstruction) Autorelease() MutableVideoCompositionInstruction {
	rv := objc.Send[MutableVideoCompositionInstruction](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMutableVideoCompositionInstruction creates a new MutableVideoCompositionInstruction instance.
func NewMutableVideoCompositionInstruction() MutableVideoCompositionInstruction {
	return getMutableVideoCompositionInstructionClass().New()
}





// A mutable video composition instruction subclass.
//
// An object maintains an array of to perform its composition.


// A mutable video composition instruction subclass.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableVideoCompositionInstruction
type MutableVideoCompositionInstruction struct {
	VideoCompositionInstruction
}

// MutableVideoCompositionInstructionFrom constructs a [MutableVideoCompositionInstruction] from an unsafe.Pointer.
//
// A mutable video composition instruction subclass.
func MutableVideoCompositionInstructionFrom(ptr unsafe.Pointer) MutableVideoCompositionInstruction {
	return MutableVideoCompositionInstruction{
		VideoCompositionInstruction: VideoCompositionInstructionFrom(ptr),
	}
}










// Returns a new mutable video composition instruction.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableVideoCompositionInstruction/videoCompositionInstruction
func (mc _MutableVideoCompositionInstructionClass) VideoCompositionInstruction() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(mc.class), objc.Sel("videoCompositionInstruction"))
	return rv
}

















// The background color of the composition.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableVideoCompositionInstruction/backgroundColor
func (m_ MutableVideoCompositionInstruction) BackgroundColor() ColorRef /* not a class type */ {
	rv := objc.Send[ColorRef](m_.ID, objc.Sel("backgroundColor"))
	return rv
}


// The background color of the composition.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableVideoCompositionInstruction/backgroundColor
func (m_ MutableVideoCompositionInstruction) SetBackgroundColor(value ColorRef /* not a class type */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setBackgroundColor:"), value)
}


// A Boolean value that indicates whether the instruction requires post processing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableVideoCompositionInstruction/enablePostProcessing
func (m_ MutableVideoCompositionInstruction) EnablePostProcessing() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("enablePostProcessing"))
	return rv
}


// A Boolean value that indicates whether the instruction requires post processing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableVideoCompositionInstruction/enablePostProcessing
func (m_ MutableVideoCompositionInstruction) SetEnablePostProcessing(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setEnablePostProcessing:"), value)
}


// Instructions that specify how to layer and compose video frames from source tracks.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableVideoCompositionInstruction/layerInstructions
func (m_ MutableVideoCompositionInstruction) LayerInstructions() []VideoCompositionLayerInstruction {
	rv := objc.Send[[]VideoCompositionLayerInstruction](m_.ID, objc.Sel("layerInstructions"))
	return rv
}


// Instructions that specify how to layer and compose video frames from source tracks.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableVideoCompositionInstruction/layerInstructions
func (m_ MutableVideoCompositionInstruction) SetLayerInstructions(value []VideoCompositionLayerInstruction) {
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](m_.ID, objc.Sel("setLayerInstructions:"), nsArray)
}


// The track identifiers of source sample data that the compositor requires to compose frames for the instruction.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableVideoCompositionInstruction/requiredSourceSampleDataTrackIDs
func (m_ MutableVideoCompositionInstruction) RequiredSourceSampleDataTrackIDs() []foundation.Number {
	rv := objc.Send[[]foundation.Number](m_.ID, objc.Sel("requiredSourceSampleDataTrackIDs"))
	return rv
}


// The track identifiers of source sample data that the compositor requires to compose frames for the instruction.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableVideoCompositionInstruction/requiredSourceSampleDataTrackIDs
func (m_ MutableVideoCompositionInstruction) SetRequiredSourceSampleDataTrackIDs(value []foundation.Number) {
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](m_.ID, objc.Sel("setRequiredSourceSampleDataTrackIDs:"), nsArray)
}


// The time range to which the instruction applies.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableVideoCompositionInstruction/timeRange
func (m_ MutableVideoCompositionInstruction) TimeRange() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](m_.ID, objc.Sel("timeRange"))
	return rv
}


// The time range to which the instruction applies.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableVideoCompositionInstruction/timeRange
func (m_ MutableVideoCompositionInstruction) SetTimeRange(value objectivec.IObject) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimeRange:"), value)
}


// The video composition instructions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avvideocomposition/instructions
func (m_ MutableVideoCompositionInstruction) Instructions() VideoCompositionInstructionProtocol /* not a class type */ {
	rv := objc.Send[VideoCompositionInstructionProtocol](m_.ID, objc.Sel("instructions"))
	return rv
}


// The video composition instructions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avvideocomposition/instructions
func (m_ MutableVideoCompositionInstruction) SetInstructions(value VideoCompositionInstructionProtocol /* not a class type */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setInstructions:"), value)
}








