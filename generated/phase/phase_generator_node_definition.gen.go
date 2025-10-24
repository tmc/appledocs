// Code generated from Apple documentation for PHASE. DO NOT EDIT.

package phase

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class PHASEGeneratorNodeDefinition */


/* debug [class_header]: Header for PHASEGeneratorNodeDefinition */
// The class instance for the [PHASEGeneratorNodeDefinition] class.
var (
	PHASEGeneratorNodeDefinitionClass     _PHASEGeneratorNodeDefinitionClass
	PHASEGeneratorNodeDefinitionClassOnce sync.Once
)

func getPHASEGeneratorNodeDefinitionClass() _PHASEGeneratorNodeDefinitionClass {
	PHASEGeneratorNodeDefinitionClassOnce.Do(func() {
		PHASEGeneratorNodeDefinitionClass = _PHASEGeneratorNodeDefinitionClass{objc.GetClass("PHASEGeneratorNodeDefinition")}
	})
	return PHASEGeneratorNodeDefinitionClass
}

type _PHASEGeneratorNodeDefinitionClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for PHASEGeneratorNodeDefinition */
// An interface definition for the [PHASEGeneratorNodeDefinition] class.
type IPHASEGeneratorNodeDefinition interface {
	IPHASESoundEventNodeDefinition
	
/* debug [class_interface_properties]: Properties for PHASEGeneratorNodeDefinition */
	// properties:
	CalibrationMode() PHASECalibrationMode
	GainMetaParameterDefinition() IPHASENumberMetaParameterDefinition
	SetGainMetaParameterDefinition(value IPHASENumberMetaParameterDefinition)
	Group() IPHASEGroup
	SetGroup(value IPHASEGroup)
	Level() float64
	MixerDefinition() IPHASEMixerDefinition
	Rate() float64
	SetRate(value float64)
	RateMetaParameterDefinition() IPHASENumberMetaParameterDefinition
	SetRateMetaParameterDefinition(value IPHASENumberMetaParameterDefinition)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for PHASEGeneratorNodeDefinition */
	// methods:
	SetCalibrationModeLevel(calibrationMode PHASECalibrationMode, level float64)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for PHASEGeneratorNodeDefinition */
// Alloc allocates a new instance without initialization.
func (pc _PHASEGeneratorNodeDefinitionClass) Alloc() PHASEGeneratorNodeDefinition {
	rv := objc.Send[PHASEGeneratorNodeDefinition](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (pc _PHASEGeneratorNodeDefinitionClass) New() PHASEGeneratorNodeDefinition {
	rv := objc.Send[PHASEGeneratorNodeDefinition](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PHASEGeneratorNodeDefinition) Init() PHASEGeneratorNodeDefinition {
	rv := objc.Send[PHASEGeneratorNodeDefinition](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PHASEGeneratorNodeDefinition) Autorelease() PHASEGeneratorNodeDefinition {
	rv := objc.Send[PHASEGeneratorNodeDefinition](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPHASEGeneratorNodeDefinition creates a new PHASEGeneratorNodeDefinition instance.
func NewPHASEGeneratorNodeDefinition() PHASEGeneratorNodeDefinition {
	return getPHASEGeneratorNodeDefinitionClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for PHASEGeneratorNodeDefinition */
// A base class for nodes that provide audio data to generate sound.
//
// This class encapsulates shared logic for subclasses that provide audio data to a mixer for sound output, namely and .


// A base class for nodes that provide audio data to generate sound.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEGeneratorNodeDefinition
type PHASEGeneratorNodeDefinition struct {
	PHASESoundEventNodeDefinition
}

// PHASEGeneratorNodeDefinitionFrom constructs a [PHASEGeneratorNodeDefinition] from an unsafe.Pointer.
//
// A base class for nodes that provide audio data to generate sound.
func PHASEGeneratorNodeDefinitionFrom(ptr unsafe.Pointer) PHASEGeneratorNodeDefinition {
	return PHASEGeneratorNodeDefinition{
		PHASESoundEventNodeDefinition: PHASESoundEventNodeDefinitionFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for PHASEGeneratorNodeDefinition *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for PHASEGeneratorNodeDefinition */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for PHASEGeneratorNodeDefinition */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for PHASEGeneratorNodeDefinition */

// Selects a loudness correction strategy and reference level.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEGeneratorNodeDefinition/setCalibrationMode(calibrationMode:level:)
func (p_ PHASEGeneratorNodeDefinition) SetCalibrationModeLevel(calibrationMode PHASECalibrationMode, level float64) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setCalibrationMode:level:"), calibrationMode, level)
}/* debug [instance_methods/method]: SetCalibrationModeLevel */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for PHASEGeneratorNodeDefinition */

// A sound pressure level strategy for loudness correction.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEGeneratorNodeDefinition/calibrationMode
func (p_ PHASEGeneratorNodeDefinition) CalibrationMode() PHASECalibrationMode {
	rv := objc.Send[PHASECalibrationMode](p_.ID, objc.Sel("calibrationMode"))
	return rv
}/* debug [instance_properties/getter]: calibrationMode */


// A meta parameter that dynamically changes the audio’s loudness.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEGeneratorNodeDefinition/gainMetaParameterDefinition
func (p_ PHASEGeneratorNodeDefinition) GainMetaParameterDefinition() IPHASENumberMetaParameterDefinition {
	rv := objc.Send[PHASENumberMetaParameterDefinition](p_.ID, objc.Sel("gainMetaParameterDefinition"))
	return rv
}/* debug [instance_properties/getter]: gainMetaParameterDefinition */


// A meta parameter that dynamically changes the audio’s loudness.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEGeneratorNodeDefinition/gainMetaParameterDefinition
func (p_ PHASEGeneratorNodeDefinition) SetGainMetaParameterDefinition(value IPHASENumberMetaParameterDefinition) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setGainMetaParameterDefinition:"), value)
}/* debug [instance_properties/setter]: gainMetaParameterDefinition */


// A group this node conforms to for gain and rate control.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEGeneratorNodeDefinition/group
func (p_ PHASEGeneratorNodeDefinition) Group() IPHASEGroup {
	rv := objc.Send[PHASEGroup](p_.ID, objc.Sel("group"))
	return rv
}/* debug [instance_properties/getter]: group */


// A group this node conforms to for gain and rate control.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEGeneratorNodeDefinition/group
func (p_ PHASEGeneratorNodeDefinition) SetGroup(value IPHASEGroup) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setGroup:"), value)
}/* debug [instance_properties/setter]: group */


// The node’s loudness.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEGeneratorNodeDefinition/level
func (p_ PHASEGeneratorNodeDefinition) Level() float64 {
	rv := objc.Send[float64](p_.ID, objc.Sel("level"))
	return rv
}/* debug [instance_properties/getter]: level */


// An object that combines audio layers for the node’s output.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEGeneratorNodeDefinition/mixerDefinition
func (p_ PHASEGeneratorNodeDefinition) MixerDefinition() IPHASEMixerDefinition {
	rv := objc.Send[PHASEMixerDefinition](p_.ID, objc.Sel("mixerDefinition"))
	return rv
}/* debug [instance_properties/getter]: mixerDefinition */


// A playback speed for the node’s audio.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEGeneratorNodeDefinition/rate
func (p_ PHASEGeneratorNodeDefinition) Rate() float64 {
	rv := objc.Send[float64](p_.ID, objc.Sel("rate"))
	return rv
}/* debug [instance_properties/getter]: rate */


// A playback speed for the node’s audio.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEGeneratorNodeDefinition/rate
func (p_ PHASEGeneratorNodeDefinition) SetRate(value float64) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setRate:"), value)
}/* debug [instance_properties/setter]: rate */


// A meta parameter that dynamically changes the audio’s rate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEGeneratorNodeDefinition/rateMetaParameterDefinition
func (p_ PHASEGeneratorNodeDefinition) RateMetaParameterDefinition() IPHASENumberMetaParameterDefinition {
	rv := objc.Send[PHASENumberMetaParameterDefinition](p_.ID, objc.Sel("rateMetaParameterDefinition"))
	return rv
}/* debug [instance_properties/getter]: rateMetaParameterDefinition */


// A meta parameter that dynamically changes the audio’s rate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEGeneratorNodeDefinition/rateMetaParameterDefinition
func (p_ PHASEGeneratorNodeDefinition) SetRateMetaParameterDefinition(value IPHASENumberMetaParameterDefinition) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setRateMetaParameterDefinition:"), value)
}/* debug [instance_properties/setter]: rateMetaParameterDefinition */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class PHASEGeneratorNodeDefinition */



