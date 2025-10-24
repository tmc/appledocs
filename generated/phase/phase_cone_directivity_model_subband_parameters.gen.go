// Code generated from Apple documentation for PHASE. DO NOT EDIT.

package phase

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class PHASEConeDirectivityModelSubbandParameters */


/* debug [class_header]: Header for PHASEConeDirectivityModelSubbandParameters */
// The class instance for the [PHASEConeDirectivityModelSubbandParameters] class.
var (
	PHASEConeDirectivityModelSubbandParametersClass     _PHASEConeDirectivityModelSubbandParametersClass
	PHASEConeDirectivityModelSubbandParametersClassOnce sync.Once
)

func getPHASEConeDirectivityModelSubbandParametersClass() _PHASEConeDirectivityModelSubbandParametersClass {
	PHASEConeDirectivityModelSubbandParametersClassOnce.Do(func() {
		PHASEConeDirectivityModelSubbandParametersClass = _PHASEConeDirectivityModelSubbandParametersClass{objc.GetClass("PHASEConeDirectivityModelSubbandParameters")}
	})
	return PHASEConeDirectivityModelSubbandParametersClass
}

type _PHASEConeDirectivityModelSubbandParametersClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for PHASEConeDirectivityModelSubbandParameters */
// An interface definition for the [PHASEConeDirectivityModelSubbandParameters] class.
type IPHASEConeDirectivityModelSubbandParameters interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for PHASEConeDirectivityModelSubbandParameters */
	// properties:
	Frequency() float64
	SetFrequency(value float64)
	InnerAngle() float64
	OuterAngle() float64
	OuterGain() float64
	SetOuterGain(value float64)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for PHASEConeDirectivityModelSubbandParameters */
	// methods:
	SetInnerAngleOuterAngle(innerAngle float64, outerAngle float64)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for PHASEConeDirectivityModelSubbandParameters */
// Alloc allocates a new instance without initialization.
func (pc _PHASEConeDirectivityModelSubbandParametersClass) Alloc() PHASEConeDirectivityModelSubbandParameters {
	rv := objc.Send[PHASEConeDirectivityModelSubbandParameters](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (pc _PHASEConeDirectivityModelSubbandParametersClass) New() PHASEConeDirectivityModelSubbandParameters {
	rv := objc.Send[PHASEConeDirectivityModelSubbandParameters](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PHASEConeDirectivityModelSubbandParameters) Init() PHASEConeDirectivityModelSubbandParameters {
	rv := objc.Send[PHASEConeDirectivityModelSubbandParameters](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PHASEConeDirectivityModelSubbandParameters) Autorelease() PHASEConeDirectivityModelSubbandParameters {
	rv := objc.Send[PHASEConeDirectivityModelSubbandParameters](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPHASEConeDirectivityModelSubbandParameters creates a new PHASEConeDirectivityModelSubbandParameters instance.
func NewPHASEConeDirectivityModelSubbandParameters() PHASEConeDirectivityModelSubbandParameters {
	return getPHASEConeDirectivityModelSubbandParametersClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for PHASEConeDirectivityModelSubbandParameters */
// A data set that projects sound of a certain frequency outward in the shape of a cone.
//
// This class defines one subband in the class’s . The inner and outer angles you define with describe a cone that directs sound of a given toward the listener. The cone’s point rests at the 3D position of the sound source. The framework adjusts the volume of the sound according to location of the listener in the 3D scene: If the listener positions in an area outside of the subband’s , the sound emanates from the source at the volume defined by . If the listener positions inside the area defined by , the sound emanates from the source at maximum volume. If the listener positions in between the outer and inner angles, the framework blends the volume to a value between and the maximum.


// A data set that projects sound of a certain frequency outward in the shape of a cone.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEConeDirectivityModelSubbandParameters
type PHASEConeDirectivityModelSubbandParameters struct {
	objectivec.Object
}

// PHASEConeDirectivityModelSubbandParametersFrom constructs a [PHASEConeDirectivityModelSubbandParameters] from an unsafe.Pointer.
//
// A data set that projects sound of a certain frequency outward in the shape of a cone.
func PHASEConeDirectivityModelSubbandParametersFrom(ptr unsafe.Pointer) PHASEConeDirectivityModelSubbandParameters {
	return PHASEConeDirectivityModelSubbandParameters{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for PHASEConeDirectivityModelSubbandParameters */
/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for PHASEConeDirectivityModelSubbandParameters */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for PHASEConeDirectivityModelSubbandParameters */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for PHASEConeDirectivityModelSubbandParameters */

// Configures a focus area for cone-based sound directivity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEConeDirectivityModelSubbandParameters/setAngles(innerAngle:outerAngle:)
func (p_ PHASEConeDirectivityModelSubbandParameters) SetInnerAngleOuterAngle(innerAngle float64, outerAngle float64) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setInnerAngle:outerAngle:"), innerAngle, outerAngle)
}/* debug [instance_methods/method]: SetInnerAngleOuterAngle */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for PHASEConeDirectivityModelSubbandParameters */

// A frequency in the audio spectrum where the subband resonates most.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEConeDirectivityModelSubbandParameters/frequency
func (p_ PHASEConeDirectivityModelSubbandParameters) Frequency() float64 {
	rv := objc.Send[float64](p_.ID, objc.Sel("frequency"))
	return rv
}/* debug [instance_properties/getter]: frequency */


// A frequency in the audio spectrum where the subband resonates most.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEConeDirectivityModelSubbandParameters/frequency
func (p_ PHASEConeDirectivityModelSubbandParameters) SetFrequency(value float64) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setFrequency:"), value)
}/* debug [instance_properties/setter]: frequency */


// An angle, in degrees, that determines the size of the audio emitting area inside the cone.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEConeDirectivityModelSubbandParameters/innerAngle
func (p_ PHASEConeDirectivityModelSubbandParameters) InnerAngle() float64 {
	rv := objc.Send[float64](p_.ID, objc.Sel("innerAngle"))
	return rv
}/* debug [instance_properties/getter]: innerAngle */


// An angle, in degrees, that determines the size of the audio emitting area outside the cone.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEConeDirectivityModelSubbandParameters/outerAngle
func (p_ PHASEConeDirectivityModelSubbandParameters) OuterAngle() float64 {
	rv := objc.Send[float64](p_.ID, objc.Sel("outerAngle"))
	return rv
}/* debug [instance_properties/getter]: outerAngle */


// The loudness of the audio the outside area of the cone emits.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEConeDirectivityModelSubbandParameters/outerGain
func (p_ PHASEConeDirectivityModelSubbandParameters) OuterGain() float64 {
	rv := objc.Send[float64](p_.ID, objc.Sel("outerGain"))
	return rv
}/* debug [instance_properties/getter]: outerGain */


// The loudness of the audio the outside area of the cone emits.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEConeDirectivityModelSubbandParameters/outerGain
func (p_ PHASEConeDirectivityModelSubbandParameters) SetOuterGain(value float64) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setOuterGain:"), value)
}/* debug [instance_properties/setter]: outerGain */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class PHASEConeDirectivityModelSubbandParameters */


