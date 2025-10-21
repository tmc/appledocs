// Code generated from Apple documentation for PHASE. DO NOT EDIT.

package phase

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [PHASESpatialPipelineEntry] class.
type IPHASESpatialPipelineEntry interface {
	objectivec.IObject
}

// An audio layer with an adjustable volume for a spatial mixer’s output.
//
// This property adjusts the amount of audio that passes through a spatial mixer’s pipeline ( ) to the output. The pipeline’s contains an instance of this class for each type of audio layer that defines. Depending on the layer’s type, the audio may sound like spatial relections, environmental reverb, or the unfiltered signal. An app adjusts the layer’s presence in the mixer’s output by: Defining an initial volume using Adjusting the audio’s volume dynamically, for example, by fading it over a duration using
//
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

// Alloc allocates a new instance without initialization.
func (pc _PHASESpatialPipelineEntryClass) Alloc() PHASESpatialPipelineEntry {
	rv := objc.Send[PHASESpatialPipelineEntry](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


// A parameter that gradually updates the amount of audio signal that passes through to the output.
//
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASESpatialPipelineEntry/sendLevelMetaParameterDefinition
func (p_ PHASESpatialPipelineEntry) SendLevelMetaParameterDefinition() PHASENumberMetaParameterDefinition {
	rv := objc.Send[PHASENumberMetaParameterDefinition](p_.ID, objc.Sel("sendLevelMetaParameterDefinition"))
	return rv
}


// SetSendLevelMetaParameterDefinition sets the value of the sendLevelMetaParameterDefinition property.
// A parameter that gradually updates the amount of audio signal that passes through to the output.

//
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASESpatialPipelineEntry/sendLevelMetaParameterDefinition
func (p_ PHASESpatialPipelineEntry) SetSendLevelMetaParameterDefinition(value IPHASENumberMetaParameterDefinition) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setSendLevelMetaParameterDefinition:"), value)
}

// An object that adds sound layers for environmental effects.
//
// [Full Topic]: https://developer.apple.com/documentation/phase/phasespatialmixerdefinition/spatialpipeline
func (p_ PHASESpatialPipelineEntry) SpatialPipeline() PHASESpatialPipeline {
	rv := objc.Send[PHASESpatialPipeline](p_.ID, objc.Sel("spatialPipeline"))
	return rv
}


// SetSpatialPipeline sets the value of the spatialPipeline property.
// An object that adds sound layers for environmental effects.

//
// [Full Topic]: https://developer.apple.com/documentation/phase/phasespatialmixerdefinition/spatialpipeline
func (p_ PHASESpatialPipelineEntry) SetSpatialPipeline(value IPHASESpatialPipeline) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setSpatialPipeline:"), value)
}

// Audio layers for environmental effects to add to the output.
//
// [Full Topic]: https://developer.apple.com/documentation/phase/phasespatialpipeline/entries
func (p_ PHASESpatialPipelineEntry) Entries() PHASESpatialPipelineEntry {
	rv := objc.Send[PHASESpatialPipelineEntry](p_.ID, objc.Sel("entries"))
	return rv
}


// SetEntries sets the value of the entries property.
// Audio layers for environmental effects to add to the output.

//
// [Full Topic]: https://developer.apple.com/documentation/phase/phasespatialpipeline/entries
func (p_ PHASESpatialPipelineEntry) SetEntries(value IPHASESpatialPipelineEntry) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setEntries:"), value)
}

// The amount of audio signal to add to the output.
//
// [Full Topic]: https://developer.apple.com/documentation/phase/phasespatialpipelineentry/sendlevel
func (p_ PHASESpatialPipelineEntry) SendLevel() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("sendLevel"))
	return rv
}


// SetSendLevel sets the value of the sendLevel property.
// The amount of audio signal to add to the output.

//
// [Full Topic]: https://developer.apple.com/documentation/phase/phasespatialpipelineentry/sendlevel
func (p_ PHASESpatialPipelineEntry) SetSendLevel(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setSendLevel:"), value)
}



