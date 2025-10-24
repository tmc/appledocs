// Code generated from Apple documentation for PHASE. DO NOT EDIT.

package phase

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class PHASESpatialMixerDefinition */


/* debug [class_header]: Header for PHASESpatialMixerDefinition */
// The class instance for the [PHASESpatialMixerDefinition] class.
var (
	PHASESpatialMixerDefinitionClass     _PHASESpatialMixerDefinitionClass
	PHASESpatialMixerDefinitionClassOnce sync.Once
)

func getPHASESpatialMixerDefinitionClass() _PHASESpatialMixerDefinitionClass {
	PHASESpatialMixerDefinitionClassOnce.Do(func() {
		PHASESpatialMixerDefinitionClass = _PHASESpatialMixerDefinitionClass{objc.GetClass("PHASESpatialMixerDefinition")}
	})
	return PHASESpatialMixerDefinitionClass
}

type _PHASESpatialMixerDefinitionClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for PHASESpatialMixerDefinition */
// An interface definition for the [PHASESpatialMixerDefinition] class.
type IPHASESpatialMixerDefinition interface {
	IPHASEMixerDefinition
	
/* debug [class_interface_properties]: Properties for PHASESpatialMixerDefinition */
	// properties:
	DistanceModelParameters() IPHASEDistanceModelParameters
	SetDistanceModelParameters(value IPHASEDistanceModelParameters)
	ListenerDirectivityModelParameters() IPHASEDirectivityModelParameters
	SetListenerDirectivityModelParameters(value IPHASEDirectivityModelParameters)
	SourceDirectivityModelParameters() IPHASEDirectivityModelParameters
	SetSourceDirectivityModelParameters(value IPHASEDirectivityModelParameters)
	SpatialPipeline() IPHASESpatialPipeline
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for PHASESpatialMixerDefinition */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for PHASESpatialMixerDefinition */
// Alloc allocates a new instance without initialization.
func (pc _PHASESpatialMixerDefinitionClass) Alloc() PHASESpatialMixerDefinition {
	rv := objc.Send[PHASESpatialMixerDefinition](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (pc _PHASESpatialMixerDefinitionClass) New() PHASESpatialMixerDefinition {
	rv := objc.Send[PHASESpatialMixerDefinition](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PHASESpatialMixerDefinition) Init() PHASESpatialMixerDefinition {
	rv := objc.Send[PHASESpatialMixerDefinition](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PHASESpatialMixerDefinition) Autorelease() PHASESpatialMixerDefinition {
	rv := objc.Send[PHASESpatialMixerDefinition](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPHASESpatialMixerDefinition creates a new PHASESpatialMixerDefinition instance.
func NewPHASESpatialMixerDefinition() PHASESpatialMixerDefinition {
	return getPHASESpatialMixerDefinitionClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for PHASESpatialMixerDefinition */
// An audio-layering object that produces environmental effects and plays sound with a 3D position and orientation.
//
// This class enables the app to define a relationship between a source and listener in six degrees of freedom: orientation (roll, pitch, yaw) and a 3D position (x, y, z). The framework plays back an audio source with (see ), direct path transmission effects and any combination of environmental effects, such as reverb (see ), and directivity (see ).  The result enables an app to implement directive point or omnidirectional sound sources — with or without direction, respectively — and volumetric sources with a defined shape. For a walkthrough of spatial mixing, see .


// An audio-layering object that produces environmental effects and plays sound with a 3D position and orientation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASESpatialMixerDefinition
type PHASESpatialMixerDefinition struct {
	PHASEMixerDefinition
}

// PHASESpatialMixerDefinitionFrom constructs a [PHASESpatialMixerDefinition] from an unsafe.Pointer.
//
// An audio-layering object that produces environmental effects and plays sound with a 3D position and orientation.
func PHASESpatialMixerDefinitionFrom(ptr unsafe.Pointer) PHASESpatialMixerDefinition {
	return PHASESpatialMixerDefinition{
		PHASEMixerDefinition: PHASEMixerDefinitionFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for PHASESpatialMixerDefinition */

// Creates a mixer with the designated spatial pipeline.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASESpatialMixerDefinition/init(spatialPipeline:)
func NewPHASESpatialMixerDefinitionWithSpatialPipeline(spatialPipeline IPHASESpatialPipeline) PHASESpatialMixerDefinition {
	instance := getPHASESpatialMixerDefinitionClass().Alloc()
	rv := objc.Send[PHASESpatialMixerDefinition](instance.ID, objc.Sel("initWithSpatialPipeline:"), spatialPipeline)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewPHASESpatialMixerDefinitionWithSpatialPipeline */


// Creates a named mixer with the designated spatial pipeline.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASESpatialMixerDefinition/init(spatialPipeline:identifier:)
func NewPHASESpatialMixerDefinitionWithSpatialPipelineIdentifier(spatialPipeline IPHASESpatialPipeline, identifier objc.IObject /* cross-framework: NSString */) PHASESpatialMixerDefinition {
	instance := getPHASESpatialMixerDefinitionClass().Alloc()
	rv := objc.Send[PHASESpatialMixerDefinition](instance.ID, objc.Sel("initWithSpatialPipeline:identifier:"), spatialPipeline, identifier)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewPHASESpatialMixerDefinitionWithSpatialPipelineIdentifier */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for PHASESpatialMixerDefinition */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for PHASESpatialMixerDefinition */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for PHASESpatialMixerDefinition */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for PHASESpatialMixerDefinition */

// An effect that changes sound as it carries over a distance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASESpatialMixerDefinition/distanceModelParameters
func (p_ PHASESpatialMixerDefinition) DistanceModelParameters() IPHASEDistanceModelParameters {
	rv := objc.Send[PHASEDistanceModelParameters](p_.ID, objc.Sel("distanceModelParameters"))
	return rv
}/* debug [instance_properties/getter]: distanceModelParameters */


// An effect that changes sound as it carries over a distance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASESpatialMixerDefinition/distanceModelParameters
func (p_ PHASESpatialMixerDefinition) SetDistanceModelParameters(value IPHASEDistanceModelParameters) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setDistanceModelParameters:"), value)
}/* debug [instance_properties/setter]: distanceModelParameters */


// A data set that determines how well the listener hears depending on its direction relative to a sound source.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASESpatialMixerDefinition/listenerDirectivityModelParameters
func (p_ PHASESpatialMixerDefinition) ListenerDirectivityModelParameters() IPHASEDirectivityModelParameters {
	rv := objc.Send[PHASEDirectivityModelParameters](p_.ID, objc.Sel("listenerDirectivityModelParameters"))
	return rv
}/* debug [instance_properties/getter]: listenerDirectivityModelParameters */


// A data set that determines how well the listener hears depending on its direction relative to a sound source.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASESpatialMixerDefinition/listenerDirectivityModelParameters
func (p_ PHASESpatialMixerDefinition) SetListenerDirectivityModelParameters(value IPHASEDirectivityModelParameters) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setListenerDirectivityModelParameters:"), value)
}/* debug [instance_properties/setter]: listenerDirectivityModelParameters */


// A data set that directs sound such that it’s louder when directed at the listener.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASESpatialMixerDefinition/sourceDirectivityModelParameters
func (p_ PHASESpatialMixerDefinition) SourceDirectivityModelParameters() IPHASEDirectivityModelParameters {
	rv := objc.Send[PHASEDirectivityModelParameters](p_.ID, objc.Sel("sourceDirectivityModelParameters"))
	return rv
}/* debug [instance_properties/getter]: sourceDirectivityModelParameters */


// A data set that directs sound such that it’s louder when directed at the listener.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASESpatialMixerDefinition/sourceDirectivityModelParameters
func (p_ PHASESpatialMixerDefinition) SetSourceDirectivityModelParameters(value IPHASEDirectivityModelParameters) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setSourceDirectivityModelParameters:"), value)
}/* debug [instance_properties/setter]: sourceDirectivityModelParameters */


// An object that adds sound layers for environmental effects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASESpatialMixerDefinition/spatialPipeline
func (p_ PHASESpatialMixerDefinition) SpatialPipeline() IPHASESpatialPipeline {
	rv := objc.Send[PHASESpatialPipeline](p_.ID, objc.Sel("spatialPipeline"))
	return rv
}/* debug [instance_properties/getter]: spatialPipeline */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class PHASESpatialMixerDefinition */


