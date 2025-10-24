// Code generated from Apple documentation for GameplayKit. DO NOT EDIT.

package gameplaykit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class GKPerlinNoiseSource */


/* debug [class_header]: Header for GKPerlinNoiseSource */
// The class instance for the [PerlinNoiseSource] class.
var (
	PerlinNoiseSourceClass     _PerlinNoiseSourceClass
	PerlinNoiseSourceClassOnce sync.Once
)

func getPerlinNoiseSourceClass() _PerlinNoiseSourceClass {
	PerlinNoiseSourceClassOnce.Do(func() {
		PerlinNoiseSourceClass = _PerlinNoiseSourceClass{objc.GetClass("GKPerlinNoiseSource")}
	})
	return PerlinNoiseSourceClass
}

type _PerlinNoiseSourceClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for PerlinNoiseSource */
// An interface definition for the [PerlinNoiseSource] class.
type IPerlinNoiseSource interface {
	ICoherentNoiseSource
	
/* debug [class_interface_properties]: Properties for PerlinNoiseSource */
	// properties:
	Persistence() float64
	SetPersistence(value float64)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for PerlinNoiseSource */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for PerlinNoiseSource */
// Alloc allocates a new instance without initialization.
func (pc _PerlinNoiseSourceClass) Alloc() PerlinNoiseSource {
	rv := objc.Send[PerlinNoiseSource](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (pc _PerlinNoiseSourceClass) New() PerlinNoiseSource {
	rv := objc.Send[PerlinNoiseSource](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PerlinNoiseSource) Init() PerlinNoiseSource {
	rv := objc.Send[PerlinNoiseSource](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PerlinNoiseSource) Autorelease() PerlinNoiseSource {
	rv := objc.Send[PerlinNoiseSource](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPerlinNoiseSource creates a new PerlinNoiseSource instance.
func NewPerlinNoiseSource() PerlinNoiseSource {
	return getPerlinNoiseSourceClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for PerlinNoiseSource */
// A procedural noise generator whose output is a type of fractal coherent noise resembling natural phenomena such as clouds and terrain.
//
// Like all subclasses, a Perlin noise source represents a noise generation algorithm and its parameters. To make use of a noise source, first create object from it (and optionally apply operations to that noise object or combine it with other noise objects). Then create a object from your noise object, generating a concrete field of values that you can sample from directly or visualize using the or class.


// A procedural noise generator whose output is a type of fractal coherent noise resembling natural phenomena such as clouds and terrain.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKPerlinNoiseSource
type PerlinNoiseSource struct {
	CoherentNoiseSource
}

// PerlinNoiseSourceFrom constructs a [PerlinNoiseSource] from an unsafe.Pointer.
//
// A procedural noise generator whose output is a type of fractal coherent noise resembling natural phenomena such as clouds and terrain.
func PerlinNoiseSourceFrom(ptr unsafe.Pointer) PerlinNoiseSource {
	return PerlinNoiseSource{
		CoherentNoiseSource: CoherentNoiseSourceFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for PerlinNoiseSource */

// Initializes a Perlin noise source with the specified parameters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKPerlinNoiseSource/init(frequency:octaveCount:persistence:lacunarity:seed:)
func NewPerlinNoiseSourceWithFrequencyOctaveCountPersistenceLacunaritySeed(frequency float64, octaveCount int, persistence float64, lacunarity float64, seed int32 /* not a class type */) PerlinNoiseSource {
	instance := getPerlinNoiseSourceClass().Alloc()
	rv := objc.Send[PerlinNoiseSource](instance.ID, objc.Sel("initWithFrequency:octaveCount:persistence:lacunarity:seed:"), frequency, octaveCount, persistence, lacunarity, seed)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewPerlinNoiseSourceWithFrequencyOctaveCountPersistenceLacunaritySeed */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for PerlinNoiseSource */

// Creates a Perlin noise source with the specified parameters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKPerlinNoiseSource/perlinNoiseSourceWithFrequency:octaveCount:persistence:lacunarity:seed:
func (pc _PerlinNoiseSourceClass) PerlinNoiseSourceWithFrequencyOctaveCountPersistenceLacunaritySeed(frequency float64, octaveCount int, persistence float64, lacunarity float64, seed int32 /* not a class type */) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(pc.class), objc.Sel("perlinNoiseSourceWithFrequency:octaveCount:persistence:lacunarity:seed:"), frequency, octaveCount, persistence, lacunarity, seed)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=PerlinNoiseSourceWithFrequencyOctaveCountPersistenceLacunaritySeed) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for PerlinNoiseSource */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for PerlinNoiseSource */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for PerlinNoiseSource */

// The rate at which successive octaves of the noise function decrease in amplitude.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKPerlinNoiseSource/persistence
func (p_ PerlinNoiseSource) Persistence() float64 {
	rv := objc.Send[float64](p_.ID, objc.Sel("persistence"))
	return rv
}/* debug [instance_properties/getter]: persistence */


// The rate at which successive octaves of the noise function decrease in amplitude.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKPerlinNoiseSource/persistence
func (p_ PerlinNoiseSource) SetPersistence(value float64) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setPersistence:"), value)
}/* debug [instance_properties/setter]: persistence */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class GKPerlinNoiseSource */


