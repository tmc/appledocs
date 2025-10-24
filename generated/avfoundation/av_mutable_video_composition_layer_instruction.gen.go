// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [MutableVideoCompositionLayerInstruction] class.
var (
	MutableVideoCompositionLayerInstructionClass     _MutableVideoCompositionLayerInstructionClass
	MutableVideoCompositionLayerInstructionClassOnce sync.Once
)

func getMutableVideoCompositionLayerInstructionClass() _MutableVideoCompositionLayerInstructionClass {
	MutableVideoCompositionLayerInstructionClassOnce.Do(func() {
		MutableVideoCompositionLayerInstructionClass = _MutableVideoCompositionLayerInstructionClass{objc.GetClass("AVMutableVideoCompositionLayerInstruction")}
	})
	return MutableVideoCompositionLayerInstructionClass
}

type _MutableVideoCompositionLayerInstructionClass struct {
	class objc.Class
}





// An interface definition for the [MutableVideoCompositionLayerInstruction] class.
type IMutableVideoCompositionLayerInstruction interface {
	IVideoCompositionLayerInstruction
	

	// properties:
	TrackID() PersistentTrackID /* not a class type */
	SetTrackID(value PersistentTrackID /* not a class type */)


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (mc _MutableVideoCompositionLayerInstructionClass) Alloc() MutableVideoCompositionLayerInstruction {
	rv := objc.Send[MutableVideoCompositionLayerInstruction](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MutableVideoCompositionLayerInstructionClass) New() MutableVideoCompositionLayerInstruction {
	rv := objc.Send[MutableVideoCompositionLayerInstruction](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MutableVideoCompositionLayerInstruction) Init() MutableVideoCompositionLayerInstruction {
	rv := objc.Send[MutableVideoCompositionLayerInstruction](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MutableVideoCompositionLayerInstruction) Autorelease() MutableVideoCompositionLayerInstruction {
	rv := objc.Send[MutableVideoCompositionLayerInstruction](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMutableVideoCompositionLayerInstruction creates a new MutableVideoCompositionLayerInstruction instance.
func NewMutableVideoCompositionLayerInstruction() MutableVideoCompositionLayerInstruction {
	return getMutableVideoCompositionLayerInstructionClass().New()
}





// An object used to modify the transform, cropping, and opacity ramps applied to a given track in a mutable composition.


// An object used to modify the transform, cropping, and opacity ramps applied to a given track in a mutable composition.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableVideoCompositionLayerInstruction
type MutableVideoCompositionLayerInstruction struct {
	VideoCompositionLayerInstruction
}

// MutableVideoCompositionLayerInstructionFrom constructs a [MutableVideoCompositionLayerInstruction] from an unsafe.Pointer.
//
// An object used to modify the transform, cropping, and opacity ramps applied to a given track in a mutable composition.
func MutableVideoCompositionLayerInstructionFrom(ptr unsafe.Pointer) MutableVideoCompositionLayerInstruction {
	return MutableVideoCompositionLayerInstruction{
		VideoCompositionLayerInstruction: VideoCompositionLayerInstructionFrom(ptr),
	}
}






// Creates a new mutable video composition layer instruction for the given track.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableVideoCompositionLayerInstruction/init(assetTrack:)
func NewMutableVideoCompositionLayerInstructionWithAssetTrack(track IAVAssetTrack) MutableVideoCompositionLayerInstruction {
	rv := objc.Send[MutableVideoCompositionLayerInstruction](objc.ID(getMutableVideoCompositionLayerInstructionClass().class), objc.Sel("videoCompositionLayerInstructionWithAssetTrack:"), track)
	return rv
}







// Creates a new mutable video composition layer instruction for the given track.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableVideoCompositionLayerInstruction/init(assetTrack:)
func (mc _MutableVideoCompositionLayerInstructionClass) VideoCompositionLayerInstructionWithAssetTrack(track IAVAssetTrack) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(mc.class), objc.Sel("videoCompositionLayerInstructionWithAssetTrack:"), track)
	return rv
}


// Returns a new mutable video composition layer instruction.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableVideoCompositionLayerInstruction/videoCompositionLayerInstruction
func (mc _MutableVideoCompositionLayerInstructionClass) VideoCompositionLayerInstruction() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(mc.class), objc.Sel("videoCompositionLayerInstruction"))
	return rv
}

















// The track identifier of the source track to which the compositor applies the instruction.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableVideoCompositionLayerInstruction/trackID
func (m_ MutableVideoCompositionLayerInstruction) TrackID() PersistentTrackID /* not a class type */ {
	rv := objc.Send[PersistentTrackID](m_.ID, objc.Sel("trackID"))
	return rv
}


// The track identifier of the source track to which the compositor applies the instruction.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableVideoCompositionLayerInstruction/trackID
func (m_ MutableVideoCompositionLayerInstruction) SetTrackID(value PersistentTrackID /* not a class type */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTrackID:"), value)
}







