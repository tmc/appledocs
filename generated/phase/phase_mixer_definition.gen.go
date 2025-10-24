// Code generated from Apple documentation for PHASE. DO NOT EDIT.

package phase

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class PHASEMixerDefinition */


/* debug [class_header]: Header for PHASEMixerDefinition */
// The class instance for the [PHASEMixerDefinition] class.
var (
	PHASEMixerDefinitionClass     _PHASEMixerDefinitionClass
	PHASEMixerDefinitionClassOnce sync.Once
)

func getPHASEMixerDefinitionClass() _PHASEMixerDefinitionClass {
	PHASEMixerDefinitionClassOnce.Do(func() {
		PHASEMixerDefinitionClass = _PHASEMixerDefinitionClass{objc.GetClass("PHASEMixerDefinition")}
	})
	return PHASEMixerDefinitionClass
}

type _PHASEMixerDefinitionClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for PHASEMixerDefinition */
// An interface definition for the [PHASEMixerDefinition] class.
type IPHASEMixerDefinition interface {
	IPHASEDefinition
	
/* debug [class_interface_properties]: Properties for PHASEMixerDefinition */
	// properties:
	Gain() float64
	SetGain(value float64)
	GainMetaParameterDefinition() IPHASENumberMetaParameterDefinition
	SetGainMetaParameterDefinition(value IPHASENumberMetaParameterDefinition)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for PHASEMixerDefinition */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for PHASEMixerDefinition */
// Alloc allocates a new instance without initialization.
func (pc _PHASEMixerDefinitionClass) Alloc() PHASEMixerDefinition {
	rv := objc.Send[PHASEMixerDefinition](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (pc _PHASEMixerDefinitionClass) New() PHASEMixerDefinition {
	rv := objc.Send[PHASEMixerDefinition](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PHASEMixerDefinition) Init() PHASEMixerDefinition {
	rv := objc.Send[PHASEMixerDefinition](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PHASEMixerDefinition) Autorelease() PHASEMixerDefinition {
	rv := objc.Send[PHASEMixerDefinition](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPHASEMixerDefinition creates a new PHASEMixerDefinition instance.
func NewPHASEMixerDefinition() PHASEMixerDefinition {
	return getPHASEMixerDefinitionClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for PHASEMixerDefinition */
// An object to initialize a mixer with a given configuration.
//
// A mixer combines multiple layers of audio to a single signal for transmission to the output device. The framework creates a mixer when you provide a mixer definition. Instead of creating an instance of this class, instantiate one of the mixer definition subclasses instead:


// An object to initialize a mixer with a given configuration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEMixerDefinition
type PHASEMixerDefinition struct {
	PHASEDefinition
}

// PHASEMixerDefinitionFrom constructs a [PHASEMixerDefinition] from an unsafe.Pointer.
//
// An object to initialize a mixer with a given configuration.
func PHASEMixerDefinitionFrom(ptr unsafe.Pointer) PHASEMixerDefinition {
	return PHASEMixerDefinition{
		PHASEDefinition: PHASEDefinitionFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for PHASEMixerDefinition *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for PHASEMixerDefinition */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for PHASEMixerDefinition */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for PHASEMixerDefinition */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for PHASEMixerDefinition */

// The mixer’s volume.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEMixerDefinition/gain
func (p_ PHASEMixerDefinition) Gain() float64 {
	rv := objc.Send[float64](p_.ID, objc.Sel("gain"))
	return rv
}/* debug [instance_properties/getter]: gain */


// The mixer’s volume.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEMixerDefinition/gain
func (p_ PHASEMixerDefinition) SetGain(value float64) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setGain:"), value)
}/* debug [instance_properties/setter]: gain */


// A template for a parameter that changes the mixer’s volume gradually over a period of time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEMixerDefinition/gainMetaParameterDefinition
func (p_ PHASEMixerDefinition) GainMetaParameterDefinition() IPHASENumberMetaParameterDefinition {
	rv := objc.Send[PHASENumberMetaParameterDefinition](p_.ID, objc.Sel("gainMetaParameterDefinition"))
	return rv
}/* debug [instance_properties/getter]: gainMetaParameterDefinition */


// A template for a parameter that changes the mixer’s volume gradually over a period of time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEMixerDefinition/gainMetaParameterDefinition
func (p_ PHASEMixerDefinition) SetGainMetaParameterDefinition(value IPHASENumberMetaParameterDefinition) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setGainMetaParameterDefinition:"), value)
}/* debug [instance_properties/setter]: gainMetaParameterDefinition */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class PHASEMixerDefinition */



