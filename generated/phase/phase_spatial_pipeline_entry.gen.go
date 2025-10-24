// Code generated from Apple documentation for PHASE. DO NOT EDIT.

package phase

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class PHASESpatialPipelineEntry */


/* debug [class_header]: Header for PHASESpatialPipelineEntry */
// The class instance for the [PHASESpatialPipelineEntry] class.
var (
	PHASESpatialPipelineEntryClass     _PHASESpatialPipelineEntryClass
	PHASESpatialPipelineEntryClassOnce sync.Once
)

func getPHASESpatialPipelineEntryClass() _PHASESpatialPipelineEntryClass {
	PHASESpatialPipelineEntryClassOnce.Do(func() {
		PHASESpatialPipelineEntryClass = _PHASESpatialPipelineEntryClass{objc.GetClass("PHASESpatialPipelineEntry")}
	})
	return PHASESpatialPipelineEntryClass
}

type _PHASESpatialPipelineEntryClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for PHASESpatialPipelineEntry */
// An interface definition for the [PHASESpatialPipelineEntry] class.
type IPHASESpatialPipelineEntry interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for PHASESpatialPipelineEntry */
	// properties:
	SendLevel() float64
	SetSendLevel(value float64)
	SendLevelMetaParameterDefinition() IPHASENumberMetaParameterDefinition
	SetSendLevelMetaParameterDefinition(value IPHASENumberMetaParameterDefinition)
	SpatialPipeline() IPHASESpatialPipeline
	SetSpatialPipeline(value IPHASESpatialPipeline)
	Entries() IPHASESpatialPipelineEntry
	SetEntries(value IPHASESpatialPipelineEntry)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for PHASESpatialPipelineEntry */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for PHASESpatialPipelineEntry */
// Alloc allocates a new instance without initialization.
func (pc _PHASESpatialPipelineEntryClass) Alloc() PHASESpatialPipelineEntry {
	rv := objc.Send[PHASESpatialPipelineEntry](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (pc _PHASESpatialPipelineEntryClass) New() PHASESpatialPipelineEntry {
	rv := objc.Send[PHASESpatialPipelineEntry](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PHASESpatialPipelineEntry) Init() PHASESpatialPipelineEntry {
	rv := objc.Send[PHASESpatialPipelineEntry](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PHASESpatialPipelineEntry) Autorelease() PHASESpatialPipelineEntry {
	rv := objc.Send[PHASESpatialPipelineEntry](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPHASESpatialPipelineEntry creates a new PHASESpatialPipelineEntry instance.
func NewPHASESpatialPipelineEntry() PHASESpatialPipelineEntry {
	return getPHASESpatialPipelineEntryClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for PHASESpatialPipelineEntry */
// An audio layer with an adjustable volume for a spatial mixer’s output.
//
// This property adjusts the amount of audio that passes through a spatial mixer’s pipeline ( ) to the output. The pipeline’s contains an instance of this class for each type of audio layer that defines. Depending on the layer’s type, the audio may sound like spatial relections, environmental reverb, or the unfiltered signal. An app adjusts the layer’s presence in the mixer’s output by: Defining an initial volume using Adjusting the audio’s volume dynamically, for example, by fading it over a duration using


// An audio layer with an adjustable volume for a spatial mixer’s output.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASESpatialPipelineEntry
type PHASESpatialPipelineEntry struct {
	objectivec.Object
}

// PHASESpatialPipelineEntryFrom constructs a [PHASESpatialPipelineEntry] from an unsafe.Pointer.
//
// An audio layer with an adjustable volume for a spatial mixer’s output.
func PHASESpatialPipelineEntryFrom(ptr unsafe.Pointer) PHASESpatialPipelineEntry {
	return PHASESpatialPipelineEntry{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for PHASESpatialPipelineEntry *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for PHASESpatialPipelineEntry */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for PHASESpatialPipelineEntry */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for PHASESpatialPipelineEntry */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for PHASESpatialPipelineEntry */

// The amount of audio signal to add to the output.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASESpatialPipelineEntry/sendLevel
func (p_ PHASESpatialPipelineEntry) SendLevel() float64 {
	rv := objc.Send[float64](p_.ID, objc.Sel("sendLevel"))
	return rv
}/* debug [instance_properties/getter]: sendLevel */


// The amount of audio signal to add to the output.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASESpatialPipelineEntry/sendLevel
func (p_ PHASESpatialPipelineEntry) SetSendLevel(value float64) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setSendLevel:"), value)
}/* debug [instance_properties/setter]: sendLevel */


// A parameter that gradually updates the amount of audio signal that passes through to the output.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASESpatialPipelineEntry/sendLevelMetaParameterDefinition
func (p_ PHASESpatialPipelineEntry) SendLevelMetaParameterDefinition() IPHASENumberMetaParameterDefinition {
	rv := objc.Send[PHASENumberMetaParameterDefinition](p_.ID, objc.Sel("sendLevelMetaParameterDefinition"))
	return rv
}/* debug [instance_properties/getter]: sendLevelMetaParameterDefinition */


// A parameter that gradually updates the amount of audio signal that passes through to the output.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASESpatialPipelineEntry/sendLevelMetaParameterDefinition
func (p_ PHASESpatialPipelineEntry) SetSendLevelMetaParameterDefinition(value IPHASENumberMetaParameterDefinition) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setSendLevelMetaParameterDefinition:"), value)
}/* debug [instance_properties/setter]: sendLevelMetaParameterDefinition */


// An object that adds sound layers for environmental effects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/phase/phasespatialmixerdefinition/spatialpipeline
func (p_ PHASESpatialPipelineEntry) SpatialPipeline() IPHASESpatialPipeline {
	rv := objc.Send[PHASESpatialPipeline](p_.ID, objc.Sel("spatialPipeline"))
	return rv
}/* debug [instance_properties/getter]: spatialPipeline */


// An object that adds sound layers for environmental effects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/phase/phasespatialmixerdefinition/spatialpipeline
func (p_ PHASESpatialPipelineEntry) SetSpatialPipeline(value IPHASESpatialPipeline) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setSpatialPipeline:"), value)
}/* debug [instance_properties/setter]: spatialPipeline */


// Audio layers for environmental effects to add to the output.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/phase/phasespatialpipeline/entries
func (p_ PHASESpatialPipelineEntry) Entries() IPHASESpatialPipelineEntry {
	rv := objc.Send[PHASESpatialPipelineEntry](p_.ID, objc.Sel("entries"))
	return rv
}/* debug [instance_properties/getter]: entries */


// Audio layers for environmental effects to add to the output.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/phase/phasespatialpipeline/entries
func (p_ PHASESpatialPipelineEntry) SetEntries(value IPHASESpatialPipelineEntry) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setEntries:"), value)
}/* debug [instance_properties/setter]: entries */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class PHASESpatialPipelineEntry */



