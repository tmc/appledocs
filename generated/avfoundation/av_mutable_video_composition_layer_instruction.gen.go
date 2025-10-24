// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
	"github.com/tmc/appledocs/generated/corevideo"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class AVMutableVideoCompositionLayerInstruction */


/* debug [class_header]: Header for AVMutableVideoCompositionLayerInstruction */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MutableVideoCompositionLayerInstruction */
// An interface definition for the [MutableVideoCompositionLayerInstruction] class.
type IMutableVideoCompositionLayerInstruction interface {
	IVideoCompositionLayerInstruction
	
/* debug [class_interface_properties]: Properties for MutableVideoCompositionLayerInstruction */
	// properties:
	TrackID() PersistentTrackID /* not a class type */
	SetTrackID(value PersistentTrackID /* not a class type */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MutableVideoCompositionLayerInstruction */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MutableVideoCompositionLayerInstruction */
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MutableVideoCompositionLayerInstruction */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MutableVideoCompositionLayerInstruction */

// Creates a new mutable video composition layer instruction for the given track.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableVideoCompositionLayerInstruction/init(assetTrack:)
func NewMutableVideoCompositionLayerInstructionWithAssetTrack(track IAVAssetTrack) MutableVideoCompositionLayerInstruction {
	rv := objc.Send[MutableVideoCompositionLayerInstruction](objc.ID(getMutableVideoCompositionLayerInstructionClass().class), objc.Sel("videoCompositionLayerInstructionWithAssetTrack:"), track)
	return rv
}/* debug [class_init_methods/constructor]: NewMutableVideoCompositionLayerInstructionWithAssetTrack */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MutableVideoCompositionLayerInstruction */

// Creates a new mutable video composition layer instruction for the given track.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableVideoCompositionLayerInstruction/init(assetTrack:)
func (mc _MutableVideoCompositionLayerInstructionClass) VideoCompositionLayerInstructionWithAssetTrack(track IAVAssetTrack) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(mc.class), objc.Sel("videoCompositionLayerInstructionWithAssetTrack:"), track)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=VideoCompositionLayerInstructionWithAssetTrack) */


// Returns a new mutable video composition layer instruction.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableVideoCompositionLayerInstruction/videoCompositionLayerInstruction
func (mc _MutableVideoCompositionLayerInstructionClass) VideoCompositionLayerInstruction() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(mc.class), objc.Sel("videoCompositionLayerInstruction"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=VideoCompositionLayerInstruction) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MutableVideoCompositionLayerInstruction */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MutableVideoCompositionLayerInstruction */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MutableVideoCompositionLayerInstruction */

// The track identifier of the source track to which the compositor applies the instruction.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableVideoCompositionLayerInstruction/trackID
func (m_ MutableVideoCompositionLayerInstruction) TrackID() PersistentTrackID /* not a class type */ {
	rv := objc.Send[PersistentTrackID](m_.ID, objc.Sel("trackID"))
	return rv
}/* debug [instance_properties/getter]: trackID */


// The track identifier of the source track to which the compositor applies the instruction.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableVideoCompositionLayerInstruction/trackID
func (m_ MutableVideoCompositionLayerInstruction) SetTrackID(value PersistentTrackID /* not a class type */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTrackID:"), value)
}/* debug [instance_properties/setter]: trackID */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVMutableVideoCompositionLayerInstruction */


