// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
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
	TrackID() unsafe.Pointer
	SetTrackID(value unsafe.Pointer)
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

// Alloc allocates a new instance without initialization.
func (vc _VideoCompositionLayerInstructionClass) Alloc() VideoCompositionLayerInstruction {
	rv := objc.Send[VideoCompositionLayerInstruction](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// The track identifier of the source track to which the compositor will apply the instruction.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avvideocompositionlayerinstruction/trackid
func (v_ VideoCompositionLayerInstruction) TrackID() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID, objc.Sel("trackID"))
	return rv
}


// The track identifier of the source track to which the compositor will apply the instruction.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avvideocompositionlayerinstruction/trackid
func (v_ VideoCompositionLayerInstruction) SetTrackID(value unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setTrackID:"), value)
}




