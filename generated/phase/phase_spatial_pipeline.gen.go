// Code generated from Apple documentation for PHASE. DO NOT EDIT.

package phase

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class PHASESpatialPipeline */


/* debug [class_header]: Header for PHASESpatialPipeline */
// The class instance for the [PHASESpatialPipeline] class.
var (
	PHASESpatialPipelineClass     _PHASESpatialPipelineClass
	PHASESpatialPipelineClassOnce sync.Once
)

func getPHASESpatialPipelineClass() _PHASESpatialPipelineClass {
	PHASESpatialPipelineClassOnce.Do(func() {
		PHASESpatialPipelineClass = _PHASESpatialPipelineClass{objc.GetClass("PHASESpatialPipeline")}
	})
	return PHASESpatialPipelineClass
}

type _PHASESpatialPipelineClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for PHASESpatialPipeline */
// An interface definition for the [PHASESpatialPipeline] class.
type IPHASESpatialPipeline interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for PHASESpatialPipeline */
	// properties:
	Entries() foundation.IDictionary
	Flags() PHASESpatialPipelineFlags
	SpatialPipeline() IPHASESpatialPipeline
	SetSpatialPipeline(value IPHASESpatialPipeline)
	SendLevel() float64
	SetSendLevel(value float64)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for PHASESpatialPipeline */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for PHASESpatialPipeline */
// Alloc allocates a new instance without initialization.
func (pc _PHASESpatialPipelineClass) Alloc() PHASESpatialPipeline {
	rv := objc.Send[PHASESpatialPipeline](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (pc _PHASESpatialPipelineClass) New() PHASESpatialPipeline {
	rv := objc.Send[PHASESpatialPipeline](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PHASESpatialPipeline) Init() PHASESpatialPipeline {
	rv := objc.Send[PHASESpatialPipeline](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PHASESpatialPipeline) Autorelease() PHASESpatialPipeline {
	rv := objc.Send[PHASESpatialPipeline](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPHASESpatialPipeline creates a new PHASESpatialPipeline instance.
func NewPHASESpatialPipeline() PHASESpatialPipeline {
	return getPHASESpatialPipelineClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for PHASESpatialPipeline */
// An object that specifies the volume of optional environmental effects.
//
// The class contains an instance of this class, , to add optional sound layers to the output. On top of the original audio signal designated by , this class optionally includes audio layers for environmental effects, such as or , in the output. To control the amount of volume that either audio layer possesses in the mixer’s output, adjust the for the layer’s respective member in the dictionary.


// An object that specifies the volume of optional environmental effects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASESpatialPipeline
type PHASESpatialPipeline struct {
	objectivec.Object
}

// PHASESpatialPipelineFrom constructs a [PHASESpatialPipeline] from an unsafe.Pointer.
//
// An object that specifies the volume of optional environmental effects.
func PHASESpatialPipelineFrom(ptr unsafe.Pointer) PHASESpatialPipeline {
	return PHASESpatialPipeline{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for PHASESpatialPipeline */

// Creates a spatial pipeline with the specified flags.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASESpatialPipeline/init(flags:)
func NewPHASESpatialPipelineWithFlags(flags PHASESpatialPipelineFlags) PHASESpatialPipeline {
	instance := getPHASESpatialPipelineClass().Alloc()
	rv := objc.Send[PHASESpatialPipeline](instance.ID, objc.Sel("initWithFlags:"), flags)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewPHASESpatialPipelineWithFlags */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for PHASESpatialPipeline */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for PHASESpatialPipeline */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for PHASESpatialPipeline */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for PHASESpatialPipeline */

// Audio layers for environmental effects to add to the output.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASESpatialPipeline/entries
func (p_ PHASESpatialPipeline) Entries() foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](p_.ID, objc.Sel("entries"))
	return rv
}/* debug [instance_properties/getter]: entries */


// A collection of environmental effects to include in the output.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASESpatialPipeline/flags-swift.property
func (p_ PHASESpatialPipeline) Flags() PHASESpatialPipelineFlags {
	rv := objc.Send[PHASESpatialPipelineFlags](p_.ID, objc.Sel("flags"))
	return rv
}/* debug [instance_properties/getter]: flags */


// An object that adds sound layers for environmental effects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/phase/phasespatialmixerdefinition/spatialpipeline
func (p_ PHASESpatialPipeline) SpatialPipeline() IPHASESpatialPipeline {
	rv := objc.Send[PHASESpatialPipeline](p_.ID, objc.Sel("spatialPipeline"))
	return rv
}/* debug [instance_properties/getter]: spatialPipeline */


// An object that adds sound layers for environmental effects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/phase/phasespatialmixerdefinition/spatialpipeline
func (p_ PHASESpatialPipeline) SetSpatialPipeline(value IPHASESpatialPipeline) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setSpatialPipeline:"), value)
}/* debug [instance_properties/setter]: spatialPipeline */


// The amount of audio signal to add to the output.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/phase/phasespatialpipelineentry/sendlevel
func (p_ PHASESpatialPipeline) SendLevel() float64 {
	rv := objc.Send[float64](p_.ID, objc.Sel("sendLevel"))
	return rv
}/* debug [instance_properties/getter]: sendLevel */


// The amount of audio signal to add to the output.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/phase/phasespatialpipelineentry/sendlevel
func (p_ PHASESpatialPipeline) SetSendLevel(value float64) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setSendLevel:"), value)
}/* debug [instance_properties/setter]: sendLevel */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class PHASESpatialPipeline */


