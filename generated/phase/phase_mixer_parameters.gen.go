// Code generated from Apple documentation for PHASE. DO NOT EDIT.

package phase

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class PHASEMixerParameters */


/* debug [class_header]: Header for PHASEMixerParameters */
// The class instance for the [PHASEMixerParameters] class.
var (
	PHASEMixerParametersClass     _PHASEMixerParametersClass
	PHASEMixerParametersClassOnce sync.Once
)

func getPHASEMixerParametersClass() _PHASEMixerParametersClass {
	PHASEMixerParametersClassOnce.Do(func() {
		PHASEMixerParametersClass = _PHASEMixerParametersClass{objc.GetClass("PHASEMixerParameters")}
	})
	return PHASEMixerParametersClass
}

type _PHASEMixerParametersClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for PHASEMixerParameters */
// An interface definition for the [PHASEMixerParameters] class.
type IPHASEMixerParameters interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for PHASEMixerParameters */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for PHASEMixerParameters */
	// methods:
	AddAmbientMixerParametersWithIdentifierListener(identifier objc.IObject /* cross-framework: NSString */, listener IPHASEListener)
	AddSpatialMixerParametersWithIdentifierSourceListener(identifier objc.IObject /* cross-framework: NSString */, source IPHASESource, listener IPHASEListener)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for PHASEMixerParameters */
// Alloc allocates a new instance without initialization.
func (pc _PHASEMixerParametersClass) Alloc() PHASEMixerParameters {
	rv := objc.Send[PHASEMixerParameters](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (pc _PHASEMixerParametersClass) New() PHASEMixerParameters {
	rv := objc.Send[PHASEMixerParameters](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PHASEMixerParameters) Init() PHASEMixerParameters {
	rv := objc.Send[PHASEMixerParameters](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PHASEMixerParameters) Autorelease() PHASEMixerParameters {
	rv := objc.Send[PHASEMixerParameters](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPHASEMixerParameters creates a new PHASEMixerParameters instance.
func NewPHASEMixerParameters() PHASEMixerParameters {
	return getPHASEMixerParametersClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for PHASEMixerParameters */
// An object that specifies a mixer for sound events and orients them in 3D space.
//
// This class orients a sound event in 3D space relative to a listener. When you configure an ambient mixer’s orientation and a listener’s orientation, PHASE lowers the volume of the sound event if the two orientations point away from each other, and plays the sound at full volume if they point at each other. To add an instance of this class to a sound event, use the argument of a sound event’s initializer. Alternatively, PHASE can adjust a sound event’s loudness based on its distance from the listener in 3D space. By calling this class’s function, you supply a sound source that defines the location. For more information, see . Ambient sound events define only a listener and play with a consistent loudness, regardless of the listener’s position in the scene. To define a listener and select a particular ambient mixer that outputs the sound, call this class’s function.


// An object that specifies a mixer for sound events and orients them in 3D space.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEMixerParameters
type PHASEMixerParameters struct {
	objectivec.Object
}

// PHASEMixerParametersFrom constructs a [PHASEMixerParameters] from an unsafe.Pointer.
//
// An object that specifies a mixer for sound events and orients them in 3D space.
func PHASEMixerParametersFrom(ptr unsafe.Pointer) PHASEMixerParameters {
	return PHASEMixerParameters{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for PHASEMixerParameters *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for PHASEMixerParameters */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for PHASEMixerParameters */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for PHASEMixerParameters */

// Adds runtime parameters for an ambient mixer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEMixerParameters/addAmbientMixerParameters(identifier:listener:)
func (p_ PHASEMixerParameters) AddAmbientMixerParametersWithIdentifierListener(identifier objc.IObject /* cross-framework: NSString */, listener IPHASEListener) {
	objc.Send[objc.ID](p_.ID, objc.Sel("addAmbientMixerParametersWithIdentifier:listener:"), identifier, listener)
}/* debug [instance_methods/method]: AddAmbientMixerParametersWithIdentifierListener */


// Adds runtime parameters for a spatial mixer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEMixerParameters/addSpatialMixerParameters(identifier:source:listener:)
func (p_ PHASEMixerParameters) AddSpatialMixerParametersWithIdentifierSourceListener(identifier objc.IObject /* cross-framework: NSString */, source IPHASESource, listener IPHASEListener) {
	objc.Send[objc.ID](p_.ID, objc.Sel("addSpatialMixerParametersWithIdentifier:source:listener:"), identifier, source, listener)
}/* debug [instance_methods/method]: AddSpatialMixerParametersWithIdentifierSourceListener */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for PHASEMixerParameters */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class PHASEMixerParameters */



