// Code generated from Apple documentation for PHASE. DO NOT EDIT.

package phase

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class PHASEDistanceModelFadeOutParameters */


/* debug [class_header]: Header for PHASEDistanceModelFadeOutParameters */
// The class instance for the [PHASEDistanceModelFadeOutParameters] class.
var (
	PHASEDistanceModelFadeOutParametersClass     _PHASEDistanceModelFadeOutParametersClass
	PHASEDistanceModelFadeOutParametersClassOnce sync.Once
)

func getPHASEDistanceModelFadeOutParametersClass() _PHASEDistanceModelFadeOutParametersClass {
	PHASEDistanceModelFadeOutParametersClassOnce.Do(func() {
		PHASEDistanceModelFadeOutParametersClass = _PHASEDistanceModelFadeOutParametersClass{objc.GetClass("PHASEDistanceModelFadeOutParameters")}
	})
	return PHASEDistanceModelFadeOutParametersClass
}

type _PHASEDistanceModelFadeOutParametersClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for PHASEDistanceModelFadeOutParameters */
// An interface definition for the [PHASEDistanceModelFadeOutParameters] class.
type IPHASEDistanceModelFadeOutParameters interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for PHASEDistanceModelFadeOutParameters */
	// properties:
	CullDistance() float64
	FadeOutParameters() IPHASEDistanceModelFadeOutParameters
	SetFadeOutParameters(value IPHASEDistanceModelFadeOutParameters)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for PHASEDistanceModelFadeOutParameters */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for PHASEDistanceModelFadeOutParameters */
// Alloc allocates a new instance without initialization.
func (pc _PHASEDistanceModelFadeOutParametersClass) Alloc() PHASEDistanceModelFadeOutParameters {
	rv := objc.Send[PHASEDistanceModelFadeOutParameters](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (pc _PHASEDistanceModelFadeOutParametersClass) New() PHASEDistanceModelFadeOutParameters {
	rv := objc.Send[PHASEDistanceModelFadeOutParameters](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PHASEDistanceModelFadeOutParameters) Init() PHASEDistanceModelFadeOutParameters {
	rv := objc.Send[PHASEDistanceModelFadeOutParameters](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PHASEDistanceModelFadeOutParameters) Autorelease() PHASEDistanceModelFadeOutParameters {
	rv := objc.Send[PHASEDistanceModelFadeOutParameters](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPHASEDistanceModelFadeOutParameters creates a new PHASEDistanceModelFadeOutParameters instance.
func NewPHASEDistanceModelFadeOutParameters() PHASEDistanceModelFadeOutParameters {
	return getPHASEDistanceModelFadeOutParametersClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for PHASEDistanceModelFadeOutParameters */
// A distance over which the framework fades out sound.
//
// For spatial sound output, the framework stops playing a sound when its distance from the listener surpases . The framework gradually fades out the sound’s volume as the distance between the source and listener approaches . Likewise, the framework gradually fades in the sound as the distance between the source and listener approaches . A object provides an instance of this class to a spatial mixer; for more information, see .


// A distance over which the framework fades out sound.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEDistanceModelFadeOutParameters
type PHASEDistanceModelFadeOutParameters struct {
	objectivec.Object
}

// PHASEDistanceModelFadeOutParametersFrom constructs a [PHASEDistanceModelFadeOutParameters] from an unsafe.Pointer.
//
// A distance over which the framework fades out sound.
func PHASEDistanceModelFadeOutParametersFrom(ptr unsafe.Pointer) PHASEDistanceModelFadeOutParameters {
	return PHASEDistanceModelFadeOutParameters{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for PHASEDistanceModelFadeOutParameters */

// Creates a distance beyond which sound sources stop playing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEDistanceModelFadeOutParameters/init(cullDistance:)
func NewPHASEDistanceModelFadeOutParametersWithCullDistance(cullDistance float64) PHASEDistanceModelFadeOutParameters {
	instance := getPHASEDistanceModelFadeOutParametersClass().Alloc()
	rv := objc.Send[PHASEDistanceModelFadeOutParameters](instance.ID, objc.Sel("initWithCullDistance:"), cullDistance)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewPHASEDistanceModelFadeOutParametersWithCullDistance */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for PHASEDistanceModelFadeOutParameters */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for PHASEDistanceModelFadeOutParameters */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for PHASEDistanceModelFadeOutParameters */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for PHASEDistanceModelFadeOutParameters */

// The distance beyond which the framework doesn’t process the sound.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEDistanceModelFadeOutParameters/cullDistance
func (p_ PHASEDistanceModelFadeOutParameters) CullDistance() float64 {
	rv := objc.Send[float64](p_.ID, objc.Sel("cullDistance"))
	return rv
}/* debug [instance_properties/getter]: cullDistance */


// A distance over which the framework fades out the mixer’s sound.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/phase/phasedistancemodelparameters/fadeoutparameters
func (p_ PHASEDistanceModelFadeOutParameters) FadeOutParameters() IPHASEDistanceModelFadeOutParameters {
	rv := objc.Send[PHASEDistanceModelFadeOutParameters](p_.ID, objc.Sel("fadeOutParameters"))
	return rv
}/* debug [instance_properties/getter]: fadeOutParameters */


// A distance over which the framework fades out the mixer’s sound.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/phase/phasedistancemodelparameters/fadeoutparameters
func (p_ PHASEDistanceModelFadeOutParameters) SetFadeOutParameters(value IPHASEDistanceModelFadeOutParameters) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setFadeOutParameters:"), value)
}/* debug [instance_properties/setter]: fadeOutParameters */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class PHASEDistanceModelFadeOutParameters */


