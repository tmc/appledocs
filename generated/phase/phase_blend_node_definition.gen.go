// Code generated from Apple documentation for PHASE. DO NOT EDIT.

package phase

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class PHASEBlendNodeDefinition */


/* debug [class_header]: Header for PHASEBlendNodeDefinition */
// The class instance for the [PHASEBlendNodeDefinition] class.
var (
	PHASEBlendNodeDefinitionClass     _PHASEBlendNodeDefinitionClass
	PHASEBlendNodeDefinitionClassOnce sync.Once
)

func getPHASEBlendNodeDefinitionClass() _PHASEBlendNodeDefinitionClass {
	PHASEBlendNodeDefinitionClassOnce.Do(func() {
		PHASEBlendNodeDefinitionClass = _PHASEBlendNodeDefinitionClass{objc.GetClass("PHASEBlendNodeDefinition")}
	})
	return PHASEBlendNodeDefinitionClass
}

type _PHASEBlendNodeDefinitionClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for PHASEBlendNodeDefinition */
// An interface definition for the [PHASEBlendNodeDefinition] class.
type IPHASEBlendNodeDefinition interface {
	IPHASESoundEventNodeDefinition
	
/* debug [class_interface_properties]: Properties for PHASEBlendNodeDefinition */
	// properties:
	BlendParameterDefinition() IPHASENumberMetaParameterDefinition
	SpatialMixerDefinitionForDistance() IPHASESpatialMixerDefinition
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for PHASEBlendNodeDefinition */
	// methods:
	AddRangeWithEnvelopeSubtree(envelope IPHASEEnvelope, subtree IPHASESoundEventNodeDefinition)
	AddRangeForInputValuesAboveFullGainAtValueFadeCurveTypeSubtree(value float64, fullGainAtValue float64, fadeCurveType PHASECurveType, subtree IPHASESoundEventNodeDefinition)
	AddRangeForInputValuesBelowFullGainAtValueFadeCurveTypeSubtree(value float64, fullGainAtValue float64, fadeCurveType PHASECurveType, subtree IPHASESoundEventNodeDefinition)
	AddRangeForInputValuesBetweenHighValueFullGainAtLowValueFullGainAtHighValueLowFadeCurveTypeHighFadeCurveTypeSubtree(lowValue float64, highValue float64, fullGainAtLowValue float64, fullGainAtHighValue float64, lowFadeCurveType PHASECurveType, highFadeCurveType PHASECurveType, subtree IPHASESoundEventNodeDefinition)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for PHASEBlendNodeDefinition */
// Alloc allocates a new instance without initialization.
func (pc _PHASEBlendNodeDefinitionClass) Alloc() PHASEBlendNodeDefinition {
	rv := objc.Send[PHASEBlendNodeDefinition](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (pc _PHASEBlendNodeDefinitionClass) New() PHASEBlendNodeDefinition {
	rv := objc.Send[PHASEBlendNodeDefinition](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PHASEBlendNodeDefinition) Init() PHASEBlendNodeDefinition {
	rv := objc.Send[PHASEBlendNodeDefinition](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PHASEBlendNodeDefinition) Autorelease() PHASEBlendNodeDefinition {
	rv := objc.Send[PHASEBlendNodeDefinition](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPHASEBlendNodeDefinition creates a new PHASEBlendNodeDefinition instance.
func NewPHASEBlendNodeDefinition() PHASEBlendNodeDefinition {
	return getPHASEBlendNodeDefinitionClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for PHASEBlendNodeDefinition */
// A node that smoothly fades between the audio of its child nodes.
//
// This class defines a threshold and a numeric parameter the app increases and decreases to fade between child nodes. Each child node defines a range within the threshold in which the child node plays audio. As the app moves the blend parameter value between and the threshold, the blend node plays the audio of its child nodes whose range and fade curve overlap at the current value.


// A node that smoothly fades between the audio of its child nodes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEBlendNodeDefinition
type PHASEBlendNodeDefinition struct {
	PHASESoundEventNodeDefinition
}

// PHASEBlendNodeDefinitionFrom constructs a [PHASEBlendNodeDefinition] from an unsafe.Pointer.
//
// A node that smoothly fades between the audio of its child nodes.
func PHASEBlendNodeDefinitionFrom(ptr unsafe.Pointer) PHASEBlendNodeDefinition {
	return PHASEBlendNodeDefinition{
		PHASESoundEventNodeDefinition: PHASESoundEventNodeDefinitionFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for PHASEBlendNodeDefinition */

// Creates a blend node for spatial audio output.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEBlendNodeDefinition/init(spatialMixerDefinition:)
func NewPHASEBlendNodeDefinitionDistanceBlendWithSpatialMixerDefinition(spatialMixerDefinition IPHASESpatialMixerDefinition) PHASEBlendNodeDefinition {
	instance := getPHASEBlendNodeDefinitionClass().Alloc()
	rv := objc.Send[PHASEBlendNodeDefinition](instance.ID, objc.Sel("initDistanceBlendWithSpatialMixerDefinition:"), spatialMixerDefinition)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewPHASEBlendNodeDefinitionDistanceBlendWithSpatialMixerDefinition */


// Creates a named blend node for spatial audio output.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEBlendNodeDefinition/init(spatialMixerDefinition:identifier:)
func NewPHASEBlendNodeDefinitionDistanceBlendWithSpatialMixerDefinitionIdentifier(spatialMixerDefinition IPHASESpatialMixerDefinition, identifier objc.IObject /* cross-framework: NSString */) PHASEBlendNodeDefinition {
	instance := getPHASEBlendNodeDefinitionClass().Alloc()
	rv := objc.Send[PHASEBlendNodeDefinition](instance.ID, objc.Sel("initDistanceBlendWithSpatialMixerDefinition:identifier:"), spatialMixerDefinition, identifier)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewPHASEBlendNodeDefinitionDistanceBlendWithSpatialMixerDefinitionIdentifier */


// Creates a blend node with a maxiumum blend range value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEBlendNodeDefinition/init(blendMetaParameterDefinition:)
func NewPHASEBlendNodeDefinitionWithBlendMetaParameterDefinition(blendMetaParameterDefinition IPHASENumberMetaParameterDefinition) PHASEBlendNodeDefinition {
	instance := getPHASEBlendNodeDefinitionClass().Alloc()
	rv := objc.Send[PHASEBlendNodeDefinition](instance.ID, objc.Sel("initWithBlendMetaParameterDefinition:"), blendMetaParameterDefinition)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewPHASEBlendNodeDefinitionWithBlendMetaParameterDefinition */


// Creates a named blend node with a maxiumum blend range value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEBlendNodeDefinition/init(blendMetaParameterDefinition:identifier:)
func NewPHASEBlendNodeDefinitionWithBlendMetaParameterDefinitionIdentifier(blendMetaParameterDefinition IPHASENumberMetaParameterDefinition, identifier objc.IObject /* cross-framework: NSString */) PHASEBlendNodeDefinition {
	instance := getPHASEBlendNodeDefinitionClass().Alloc()
	rv := objc.Send[PHASEBlendNodeDefinition](instance.ID, objc.Sel("initWithBlendMetaParameterDefinition:identifier:"), blendMetaParameterDefinition, identifier)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewPHASEBlendNodeDefinitionWithBlendMetaParameterDefinitionIdentifier */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for PHASEBlendNodeDefinition */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for PHASEBlendNodeDefinition */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for PHASEBlendNodeDefinition */

// Adds a child node with an envelope.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEBlendNodeDefinition/addRange(envelope:subtree:)
func (p_ PHASEBlendNodeDefinition) AddRangeWithEnvelopeSubtree(envelope IPHASEEnvelope, subtree IPHASESoundEventNodeDefinition) {
	objc.Send[objc.ID](p_.ID, objc.Sel("addRangeWithEnvelope:subtree:"), envelope, subtree)
}/* debug [instance_methods/method]: AddRangeWithEnvelopeSubtree */


// Adds a child node that blends above a given value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEBlendNodeDefinition/addRangeForInputValuesAbove(value:fullGainAtValue:fadeCurveType:subtree:)
func (p_ PHASEBlendNodeDefinition) AddRangeForInputValuesAboveFullGainAtValueFadeCurveTypeSubtree(value float64, fullGainAtValue float64, fadeCurveType PHASECurveType, subtree IPHASESoundEventNodeDefinition) {
	objc.Send[objc.ID](p_.ID, objc.Sel("addRangeForInputValuesAbove:fullGainAtValue:fadeCurveType:subtree:"), value, fullGainAtValue, fadeCurveType, subtree)
}/* debug [instance_methods/method]: AddRangeForInputValuesAboveFullGainAtValueFadeCurveTypeSubtree */


// Adds a child node that blends below a given value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEBlendNodeDefinition/addRangeForInputValuesBelow(value:fullGainAtValue:fadeCurveType:subtree:)
func (p_ PHASEBlendNodeDefinition) AddRangeForInputValuesBelowFullGainAtValueFadeCurveTypeSubtree(value float64, fullGainAtValue float64, fadeCurveType PHASECurveType, subtree IPHASESoundEventNodeDefinition) {
	objc.Send[objc.ID](p_.ID, objc.Sel("addRangeForInputValuesBelow:fullGainAtValue:fadeCurveType:subtree:"), value, fullGainAtValue, fadeCurveType, subtree)
}/* debug [instance_methods/method]: AddRangeForInputValuesBelowFullGainAtValueFadeCurveTypeSubtree */


// Adds a child node that blends between a given high and low value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEBlendNodeDefinition/addRangeForInputValuesBetween(lowValue:highValue:fullGainAtLowValue:fullGainAtHighValue:lowFadeCurveType:highFadeCurveType:subtree:)
func (p_ PHASEBlendNodeDefinition) AddRangeForInputValuesBetweenHighValueFullGainAtLowValueFullGainAtHighValueLowFadeCurveTypeHighFadeCurveTypeSubtree(lowValue float64, highValue float64, fullGainAtLowValue float64, fullGainAtHighValue float64, lowFadeCurveType PHASECurveType, highFadeCurveType PHASECurveType, subtree IPHASESoundEventNodeDefinition) {
	objc.Send[objc.ID](p_.ID, objc.Sel("addRangeForInputValuesBetween:highValue:fullGainAtLowValue:fullGainAtHighValue:lowFadeCurveType:highFadeCurveType:subtree:"), lowValue, highValue, fullGainAtLowValue, fullGainAtHighValue, lowFadeCurveType, highFadeCurveType, subtree)
}/* debug [instance_methods/method]: AddRangeForInputValuesBetweenHighValueFullGainAtLowValueFullGainAtHighValueLowFadeCurveTypeHighFadeCurveTypeSubtree */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for PHASEBlendNodeDefinition */

// The meta parameter definition that caps the blend range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEBlendNodeDefinition/blendParameterDefinition
func (p_ PHASEBlendNodeDefinition) BlendParameterDefinition() IPHASENumberMetaParameterDefinition {
	rv := objc.Send[PHASENumberMetaParameterDefinition](p_.ID, objc.Sel("blendParameterDefinition"))
	return rv
}/* debug [instance_properties/getter]: blendParameterDefinition */


// An object that combines spatial audio layers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEBlendNodeDefinition/spatialMixerDefinitionForDistance
func (p_ PHASEBlendNodeDefinition) SpatialMixerDefinitionForDistance() IPHASESpatialMixerDefinition {
	rv := objc.Send[PHASESpatialMixerDefinition](p_.ID, objc.Sel("spatialMixerDefinitionForDistance"))
	return rv
}/* debug [instance_properties/getter]: spatialMixerDefinitionForDistance */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class PHASEBlendNodeDefinition */


