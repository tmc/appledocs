// Code generated from Apple documentation for PHASE. DO NOT EDIT.

package phase

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class PHASEDistanceModelParameters */


/* debug [class_header]: Header for PHASEDistanceModelParameters */
// The class instance for the [PHASEDistanceModelParameters] class.
var (
	PHASEDistanceModelParametersClass     _PHASEDistanceModelParametersClass
	PHASEDistanceModelParametersClassOnce sync.Once
)

func getPHASEDistanceModelParametersClass() _PHASEDistanceModelParametersClass {
	PHASEDistanceModelParametersClassOnce.Do(func() {
		PHASEDistanceModelParametersClass = _PHASEDistanceModelParametersClass{objc.GetClass("PHASEDistanceModelParameters")}
	})
	return PHASEDistanceModelParametersClass
}

type _PHASEDistanceModelParametersClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for PHASEDistanceModelParameters */
// An interface definition for the [PHASEDistanceModelParameters] class.
type IPHASEDistanceModelParameters interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for PHASEDistanceModelParameters */
	// properties:
	FadeOutParameters() IPHASEDistanceModelFadeOutParameters
	SetFadeOutParameters(value IPHASEDistanceModelFadeOutParameters)
	DistanceModelParameters() IPHASEDistanceModelParameters
	SetDistanceModelParameters(value IPHASEDistanceModelParameters)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for PHASEDistanceModelParameters */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for PHASEDistanceModelParameters */
// Alloc allocates a new instance without initialization.
func (pc _PHASEDistanceModelParametersClass) Alloc() PHASEDistanceModelParameters {
	rv := objc.Send[PHASEDistanceModelParameters](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (pc _PHASEDistanceModelParametersClass) New() PHASEDistanceModelParameters {
	rv := objc.Send[PHASEDistanceModelParameters](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PHASEDistanceModelParameters) Init() PHASEDistanceModelParameters {
	rv := objc.Send[PHASEDistanceModelParameters](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PHASEDistanceModelParameters) Autorelease() PHASEDistanceModelParameters {
	rv := objc.Send[PHASEDistanceModelParameters](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPHASEDistanceModelParameters creates a new PHASEDistanceModelParameters instance.
func NewPHASEDistanceModelParameters() PHASEDistanceModelParameters {
	return getPHASEDistanceModelParametersClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for PHASEDistanceModelParameters */
// A base class for a sound’s rate of change over distance.
//
// When your app outputs sound with a 3D position and orientation, designate a subclass of this class to indicate the manner in which PHASE changes sound with distance. Assign an instance of either or , depending on your app’s needs, to the class’s property.


// A base class for a sound’s rate of change over distance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEDistanceModelParameters
type PHASEDistanceModelParameters struct {
	objectivec.Object
}

// PHASEDistanceModelParametersFrom constructs a [PHASEDistanceModelParameters] from an unsafe.Pointer.
//
// A base class for a sound’s rate of change over distance.
func PHASEDistanceModelParametersFrom(ptr unsafe.Pointer) PHASEDistanceModelParameters {
	return PHASEDistanceModelParameters{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for PHASEDistanceModelParameters *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for PHASEDistanceModelParameters */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for PHASEDistanceModelParameters */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for PHASEDistanceModelParameters */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for PHASEDistanceModelParameters */

// A distance over which the framework fades out the mixer’s sound.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEDistanceModelParameters/fadeOutParameters
func (p_ PHASEDistanceModelParameters) FadeOutParameters() IPHASEDistanceModelFadeOutParameters {
	rv := objc.Send[PHASEDistanceModelFadeOutParameters](p_.ID, objc.Sel("fadeOutParameters"))
	return rv
}/* debug [instance_properties/getter]: fadeOutParameters */


// A distance over which the framework fades out the mixer’s sound.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEDistanceModelParameters/fadeOutParameters
func (p_ PHASEDistanceModelParameters) SetFadeOutParameters(value IPHASEDistanceModelFadeOutParameters) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setFadeOutParameters:"), value)
}/* debug [instance_properties/setter]: fadeOutParameters */


// An effect that changes sound as it carries over a distance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/phase/phasespatialmixerdefinition/distancemodelparameters
func (p_ PHASEDistanceModelParameters) DistanceModelParameters() IPHASEDistanceModelParameters {
	rv := objc.Send[PHASEDistanceModelParameters](p_.ID, objc.Sel("distanceModelParameters"))
	return rv
}/* debug [instance_properties/getter]: distanceModelParameters */


// An effect that changes sound as it carries over a distance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/phase/phasespatialmixerdefinition/distancemodelparameters
func (p_ PHASEDistanceModelParameters) SetDistanceModelParameters(value IPHASEDistanceModelParameters) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setDistanceModelParameters:"), value)
}/* debug [instance_properties/setter]: distanceModelParameters */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class PHASEDistanceModelParameters */



