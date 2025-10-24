// Code generated from Apple documentation for PHASE. DO NOT EDIT.

package phase

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class PHASECardioidDirectivityModelSubbandParameters */


/* debug [class_header]: Header for PHASECardioidDirectivityModelSubbandParameters */
// The class instance for the [PHASECardioidDirectivityModelSubbandParameters] class.
var (
	PHASECardioidDirectivityModelSubbandParametersClass     _PHASECardioidDirectivityModelSubbandParametersClass
	PHASECardioidDirectivityModelSubbandParametersClassOnce sync.Once
)

func getPHASECardioidDirectivityModelSubbandParametersClass() _PHASECardioidDirectivityModelSubbandParametersClass {
	PHASECardioidDirectivityModelSubbandParametersClassOnce.Do(func() {
		PHASECardioidDirectivityModelSubbandParametersClass = _PHASECardioidDirectivityModelSubbandParametersClass{objc.GetClass("PHASECardioidDirectivityModelSubbandParameters")}
	})
	return PHASECardioidDirectivityModelSubbandParametersClass
}

type _PHASECardioidDirectivityModelSubbandParametersClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for PHASECardioidDirectivityModelSubbandParameters */
// An interface definition for the [PHASECardioidDirectivityModelSubbandParameters] class.
type IPHASECardioidDirectivityModelSubbandParameters interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for PHASECardioidDirectivityModelSubbandParameters */
	// properties:
	Frequency() float64
	SetFrequency(value float64)
	Pattern() float64
	SetPattern(value float64)
	Sharpness() float64
	SetSharpness(value float64)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for PHASECardioidDirectivityModelSubbandParameters */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for PHASECardioidDirectivityModelSubbandParameters */
// Alloc allocates a new instance without initialization.
func (pc _PHASECardioidDirectivityModelSubbandParametersClass) Alloc() PHASECardioidDirectivityModelSubbandParameters {
	rv := objc.Send[PHASECardioidDirectivityModelSubbandParameters](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (pc _PHASECardioidDirectivityModelSubbandParametersClass) New() PHASECardioidDirectivityModelSubbandParameters {
	rv := objc.Send[PHASECardioidDirectivityModelSubbandParameters](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PHASECardioidDirectivityModelSubbandParameters) Init() PHASECardioidDirectivityModelSubbandParameters {
	rv := objc.Send[PHASECardioidDirectivityModelSubbandParameters](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PHASECardioidDirectivityModelSubbandParameters) Autorelease() PHASECardioidDirectivityModelSubbandParameters {
	rv := objc.Send[PHASECardioidDirectivityModelSubbandParameters](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPHASECardioidDirectivityModelSubbandParameters creates a new PHASECardioidDirectivityModelSubbandParameters instance.
func NewPHASECardioidDirectivityModelSubbandParameters() PHASECardioidDirectivityModelSubbandParameters {
	return getPHASECardioidDirectivityModelSubbandParametersClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for PHASECardioidDirectivityModelSubbandParameters */
// A data set that projects sound of a certain frequency outward in the shape of a heart.
//
// This class defines one subband in the class’s . Depending on the specific shape you define with and , you can attenuate sound focused at to the sides of the listener, while leaving the sound in front of or behind the listener unchanged.


// A data set that projects sound of a certain frequency outward in the shape of a heart.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASECardioidDirectivityModelSubbandParameters
type PHASECardioidDirectivityModelSubbandParameters struct {
	objectivec.Object
}

// PHASECardioidDirectivityModelSubbandParametersFrom constructs a [PHASECardioidDirectivityModelSubbandParameters] from an unsafe.Pointer.
//
// A data set that projects sound of a certain frequency outward in the shape of a heart.
func PHASECardioidDirectivityModelSubbandParametersFrom(ptr unsafe.Pointer) PHASECardioidDirectivityModelSubbandParameters {
	return PHASECardioidDirectivityModelSubbandParameters{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for PHASECardioidDirectivityModelSubbandParameters */
/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for PHASECardioidDirectivityModelSubbandParameters */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for PHASECardioidDirectivityModelSubbandParameters */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for PHASECardioidDirectivityModelSubbandParameters */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for PHASECardioidDirectivityModelSubbandParameters */

// A frequency in the audio spectrum where the pattern and sharpness resonate most.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASECardioidDirectivityModelSubbandParameters/frequency
func (p_ PHASECardioidDirectivityModelSubbandParameters) Frequency() float64 {
	rv := objc.Send[float64](p_.ID, objc.Sel("frequency"))
	return rv
}/* debug [instance_properties/getter]: frequency */


// A frequency in the audio spectrum where the pattern and sharpness resonate most.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASECardioidDirectivityModelSubbandParameters/frequency
func (p_ PHASECardioidDirectivityModelSubbandParameters) SetFrequency(value float64) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setFrequency:"), value)
}/* debug [instance_properties/setter]: frequency */


// A shape that determines the direction of sound.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASECardioidDirectivityModelSubbandParameters/pattern
func (p_ PHASECardioidDirectivityModelSubbandParameters) Pattern() float64 {
	rv := objc.Send[float64](p_.ID, objc.Sel("pattern"))
	return rv
}/* debug [instance_properties/getter]: pattern */


// A shape that determines the direction of sound.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASECardioidDirectivityModelSubbandParameters/pattern
func (p_ PHASECardioidDirectivityModelSubbandParameters) SetPattern(value float64) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setPattern:"), value)
}/* debug [instance_properties/setter]: pattern */


// The amount that the shape overlaps with bordering subbands.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASECardioidDirectivityModelSubbandParameters/sharpness
func (p_ PHASECardioidDirectivityModelSubbandParameters) Sharpness() float64 {
	rv := objc.Send[float64](p_.ID, objc.Sel("sharpness"))
	return rv
}/* debug [instance_properties/getter]: sharpness */


// The amount that the shape overlaps with bordering subbands.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASECardioidDirectivityModelSubbandParameters/sharpness
func (p_ PHASECardioidDirectivityModelSubbandParameters) SetSharpness(value float64) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setSharpness:"), value)
}/* debug [instance_properties/setter]: sharpness */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class PHASECardioidDirectivityModelSubbandParameters */


