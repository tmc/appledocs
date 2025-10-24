// Code generated from Apple documentation for GameplayKit. DO NOT EDIT.

package gameplaykit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class GKCoherentNoiseSource */


/* debug [class_header]: Header for GKCoherentNoiseSource */
// The class instance for the [CoherentNoiseSource] class.
var (
	CoherentNoiseSourceClass     _CoherentNoiseSourceClass
	CoherentNoiseSourceClassOnce sync.Once
)

func getCoherentNoiseSourceClass() _CoherentNoiseSourceClass {
	CoherentNoiseSourceClassOnce.Do(func() {
		CoherentNoiseSourceClass = _CoherentNoiseSourceClass{objc.GetClass("GKCoherentNoiseSource")}
	})
	return CoherentNoiseSourceClass
}

type _CoherentNoiseSourceClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CoherentNoiseSource */
// An interface definition for the [CoherentNoiseSource] class.
type ICoherentNoiseSource interface {
	INoiseSource
	
/* debug [class_interface_properties]: Properties for CoherentNoiseSource */
	// properties:
	Frequency() float64
	SetFrequency(value float64)
	Lacunarity() float64
	SetLacunarity(value float64)
	OctaveCount() int
	SetOctaveCount(value int)
	Seed() int32 /* not a class type */
	SetSeed(value int32 /* not a class type */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CoherentNoiseSource */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CoherentNoiseSource */
// Alloc allocates a new instance without initialization.
func (cc _CoherentNoiseSourceClass) Alloc() CoherentNoiseSource {
	rv := objc.Send[CoherentNoiseSource](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CoherentNoiseSourceClass) New() CoherentNoiseSource {
	rv := objc.Send[CoherentNoiseSource](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CoherentNoiseSource) Init() CoherentNoiseSource {
	rv := objc.Send[CoherentNoiseSource](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CoherentNoiseSource) Autorelease() CoherentNoiseSource {
	rv := objc.Send[CoherentNoiseSource](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCoherentNoiseSource creates a new CoherentNoiseSource instance.
func NewCoherentNoiseSource() CoherentNoiseSource {
	return getCoherentNoiseSourceClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CoherentNoiseSource */
// The abstract superclass for procedural noise generators that create coherent noise.
//
// In general, is randomness across a (one-, two-, three-, or many-dimensional) domain—for example, you can create noise by filling an image with values from a random number generator. Unlike such truly random noise, is consistent and smooth: you can always generate the same output from a specific seed value, and small variations across the domain create only small variations in noise values. You don’t instantiate or work directly with this class. Instead, the concrete subclasses of each provide a different style of coherent noise.


// The abstract superclass for procedural noise generators that create coherent noise.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKCoherentNoiseSource
type CoherentNoiseSource struct {
	NoiseSource
}

// CoherentNoiseSourceFrom constructs a [CoherentNoiseSource] from an unsafe.Pointer.
//
// The abstract superclass for procedural noise generators that create coherent noise.
func CoherentNoiseSourceFrom(ptr unsafe.Pointer) CoherentNoiseSource {
	return CoherentNoiseSource{
		NoiseSource: NoiseSourceFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CoherentNoiseSource *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CoherentNoiseSource */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CoherentNoiseSource */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CoherentNoiseSource */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CoherentNoiseSource */

// A value that determines the size and spacing of features in generated noise.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKCoherentNoiseSource/frequency
func (c_ CoherentNoiseSource) Frequency() float64 {
	rv := objc.Send[float64](c_.ID, objc.Sel("frequency"))
	return rv
}/* debug [instance_properties/getter]: frequency */


// A value that determines the size and spacing of features in generated noise.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKCoherentNoiseSource/frequency
func (c_ CoherentNoiseSource) SetFrequency(value float64) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setFrequency:"), value)
}/* debug [instance_properties/setter]: frequency */


// The rate at which successive octaves of the noise function increase in frequency.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKCoherentNoiseSource/lacunarity
func (c_ CoherentNoiseSource) Lacunarity() float64 {
	rv := objc.Send[float64](c_.ID, objc.Sel("lacunarity"))
	return rv
}/* debug [instance_properties/getter]: lacunarity */


// The rate at which successive octaves of the noise function increase in frequency.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKCoherentNoiseSource/lacunarity
func (c_ CoherentNoiseSource) SetLacunarity(value float64) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setLacunarity:"), value)
}/* debug [instance_properties/setter]: lacunarity */


// The number of octaves of the underlying noise function to use for generating noise.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKCoherentNoiseSource/octaveCount
func (c_ CoherentNoiseSource) OctaveCount() int {
	rv := objc.Send[int](c_.ID, objc.Sel("octaveCount"))
	return rv
}/* debug [instance_properties/getter]: octaveCount */


// The number of octaves of the underlying noise function to use for generating noise.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKCoherentNoiseSource/octaveCount
func (c_ CoherentNoiseSource) SetOctaveCount(value int) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setOctaveCount:"), value)
}/* debug [instance_properties/setter]: octaveCount */


// The value that determines the specific configuration of noise produced by the noise source.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKCoherentNoiseSource/seed
func (c_ CoherentNoiseSource) Seed() int32 /* not a class type */ {
	rv := objc.Send[int32](c_.ID, objc.Sel("seed"))
	return rv
}/* debug [instance_properties/getter]: seed */


// The value that determines the specific configuration of noise produced by the noise source.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKCoherentNoiseSource/seed
func (c_ CoherentNoiseSource) SetSeed(value int32 /* not a class type */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSeed:"), value)
}/* debug [instance_properties/setter]: seed */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class GKCoherentNoiseSource */



