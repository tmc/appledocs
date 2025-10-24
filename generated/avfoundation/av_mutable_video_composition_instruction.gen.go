// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class AVMutableVideoCompositionInstruction */


/* debug [class_header]: Header for AVMutableVideoCompositionInstruction */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MutableVideoCompositionInstruction */
// An interface definition for the [MutableVideoCompositionInstruction] class.
type IMutableVideoCompositionInstruction interface {
	IVideoCompositionInstruction
	
/* debug [class_interface_properties]: Properties for MutableVideoCompositionInstruction */
	// properties:
	BackgroundColor() ColorRef /* not a class type */
	SetBackgroundColor(value ColorRef /* not a class type */)
	EnablePostProcessing() bool
	SetEnablePostProcessing(value bool)
	LayerInstructions() []VideoCompositionLayerInstruction
	SetLayerInstructions(value []VideoCompositionLayerInstruction)
	RequiredSourceSampleDataTrackIDs() []foundation.Number
	SetRequiredSourceSampleDataTrackIDs(value []foundation.Number)
	TimeRange() TimeRange /* not a class type */
	SetTimeRange(value TimeRange /* not a class type */)
	Instructions() VideoCompositionInstructionProtocol /* not a class type */
	SetInstructions(value VideoCompositionInstructionProtocol /* not a class type */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MutableVideoCompositionInstruction */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MutableVideoCompositionInstruction */
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MutableVideoCompositionInstruction */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MutableVideoCompositionInstruction *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MutableVideoCompositionInstruction */

// Returns a new mutable video composition instruction.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableVideoCompositionInstruction/videoCompositionInstruction
func (mc _MutableVideoCompositionInstructionClass) VideoCompositionInstruction() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(mc.class), objc.Sel("videoCompositionInstruction"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=VideoCompositionInstruction) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MutableVideoCompositionInstruction */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MutableVideoCompositionInstruction */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MutableVideoCompositionInstruction */

// The background color of the composition.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableVideoCompositionInstruction/backgroundColor
func (m_ MutableVideoCompositionInstruction) BackgroundColor() ColorRef /* not a class type */ {
	rv := objc.Send[ColorRef](m_.ID, objc.Sel("backgroundColor"))
	return rv
}/* debug [instance_properties/getter]: backgroundColor */


// The background color of the composition.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableVideoCompositionInstruction/backgroundColor
func (m_ MutableVideoCompositionInstruction) SetBackgroundColor(value ColorRef /* not a class type */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setBackgroundColor:"), value)
}/* debug [instance_properties/setter]: backgroundColor */


// A Boolean value that indicates whether the instruction requires post processing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableVideoCompositionInstruction/enablePostProcessing
func (m_ MutableVideoCompositionInstruction) EnablePostProcessing() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("enablePostProcessing"))
	return rv
}/* debug [instance_properties/getter]: enablePostProcessing */


// A Boolean value that indicates whether the instruction requires post processing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableVideoCompositionInstruction/enablePostProcessing
func (m_ MutableVideoCompositionInstruction) SetEnablePostProcessing(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setEnablePostProcessing:"), value)
}/* debug [instance_properties/setter]: enablePostProcessing */


// Instructions that specify how to layer and compose video frames from source tracks.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableVideoCompositionInstruction/layerInstructions
func (m_ MutableVideoCompositionInstruction) LayerInstructions() []VideoCompositionLayerInstruction {
	rv := objc.Send[[]VideoCompositionLayerInstruction](m_.ID, objc.Sel("layerInstructions"))
	return rv
}/* debug [instance_properties/getter]: layerInstructions */


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
}/* debug [instance_properties/setter]: layerInstructions */


// The track identifiers of source sample data that the compositor requires to compose frames for the instruction.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableVideoCompositionInstruction/requiredSourceSampleDataTrackIDs
func (m_ MutableVideoCompositionInstruction) RequiredSourceSampleDataTrackIDs() []foundation.Number {
	rv := objc.Send[[]foundation.Number](m_.ID, objc.Sel("requiredSourceSampleDataTrackIDs"))
	return rv
}/* debug [instance_properties/getter]: requiredSourceSampleDataTrackIDs */


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
}/* debug [instance_properties/setter]: requiredSourceSampleDataTrackIDs */


// The time range to which the instruction applies.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableVideoCompositionInstruction/timeRange
func (m_ MutableVideoCompositionInstruction) TimeRange() TimeRange /* not a class type */ {
	rv := objc.Send[TimeRange](m_.ID, objc.Sel("timeRange"))
	return rv
}/* debug [instance_properties/getter]: timeRange */


// The time range to which the instruction applies.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableVideoCompositionInstruction/timeRange
func (m_ MutableVideoCompositionInstruction) SetTimeRange(value TimeRange /* not a class type */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimeRange:"), value)
}/* debug [instance_properties/setter]: timeRange */


// The video composition instructions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avvideocomposition/instructions
func (m_ MutableVideoCompositionInstruction) Instructions() VideoCompositionInstructionProtocol /* not a class type */ {
	rv := objc.Send[VideoCompositionInstructionProtocol](m_.ID, objc.Sel("instructions"))
	return rv
}/* debug [instance_properties/getter]: instructions */


// The video composition instructions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avvideocomposition/instructions
func (m_ MutableVideoCompositionInstruction) SetInstructions(value VideoCompositionInstructionProtocol /* not a class type */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setInstructions:"), value)
}/* debug [instance_properties/setter]: instructions */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVMutableVideoCompositionInstruction */



