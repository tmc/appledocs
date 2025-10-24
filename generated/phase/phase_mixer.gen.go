// Code generated from Apple documentation for PHASE. DO NOT EDIT.

package phase

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class PHASEMixer */


/* debug [class_header]: Header for PHASEMixer */
// The class instance for the [PHASEMixer] class.
var (
	PHASEMixerClass     _PHASEMixerClass
	PHASEMixerClassOnce sync.Once
)

func getPHASEMixerClass() _PHASEMixerClass {
	PHASEMixerClassOnce.Do(func() {
		PHASEMixerClass = _PHASEMixerClass{objc.GetClass("PHASEMixer")}
	})
	return PHASEMixerClass
}

type _PHASEMixerClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for PHASEMixer */
// An interface definition for the [PHASEMixer] class.
type IPHASEMixer interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for PHASEMixer */
	// properties:
	Gain() float64
	GainMetaParameter() IPHASEMetaParameter
	Identifier() objc.IObject /* cross-framework: NSString */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for PHASEMixer */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for PHASEMixer */
// Alloc allocates a new instance without initialization.
func (pc _PHASEMixerClass) Alloc() PHASEMixer {
	rv := objc.Send[PHASEMixer](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (pc _PHASEMixerClass) New() PHASEMixer {
	rv := objc.Send[PHASEMixer](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PHASEMixer) Init() PHASEMixer {
	rv := objc.Send[PHASEMixer](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PHASEMixer) Autorelease() PHASEMixer {
	rv := objc.Send[PHASEMixer](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPHASEMixer creates a new PHASEMixer instance.
func NewPHASEMixer() PHASEMixer {
	return getPHASEMixerClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for PHASEMixer */
// An object that combines multiple audio signals into a single signal.
//
// Mixers provide a single point of control over the multiple audio signals they combine. To create a mixer, you provide the framework with a mixer definition; see . Subclasses of this class define unique properties the app sets to control specific features. For example, the spatial mixer ( ) adds environmental effects into the output audio signal.


// An object that combines multiple audio signals into a single signal.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEMixer
type PHASEMixer struct {
	objectivec.Object
}

// PHASEMixerFrom constructs a [PHASEMixer] from an unsafe.Pointer.
//
// An object that combines multiple audio signals into a single signal.
func PHASEMixerFrom(ptr unsafe.Pointer) PHASEMixer {
	return PHASEMixer{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for PHASEMixer *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for PHASEMixer */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for PHASEMixer */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for PHASEMixer */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for PHASEMixer */

// The mixer’s volume.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEMixer/gain
func (p_ PHASEMixer) Gain() float64 {
	rv := objc.Send[float64](p_.ID, objc.Sel("gain"))
	return rv
}/* debug [instance_properties/getter]: gain */


// A parameter that changes the mixer’s volume gradually over a period of time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEMixer/gainMetaParameter
func (p_ PHASEMixer) GainMetaParameter() IPHASEMetaParameter {
	rv := objc.Send[PHASEMetaParameter](p_.ID, objc.Sel("gainMetaParameter"))
	return rv
}/* debug [instance_properties/getter]: gainMetaParameter */


// A unique name for the mixer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEMixer/identifier
func (p_ PHASEMixer) Identifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](p_.ID, objc.Sel("identifier"))
	return rv
}/* debug [instance_properties/getter]: identifier */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class PHASEMixer */



