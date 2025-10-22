// Code generated from Apple documentation for GameplayKit. DO NOT EDIT.

package gameplaykit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

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

// An interface definition for the [PerlinNoiseSource] class.
type IPerlinNoiseSource interface {
	ICoherentNoiseSource
	Persistence() float64
	SetPersistence(value float64)
}

// A procedural noise generator whose output is a type of fractal coherent noise resembling natural phenomena such as clouds and terrain.
//
// Like all subclasses, a Perlin noise source represents a noise generation algorithm and its parameters. To make use of a noise source, first create object from it (and optionally apply operations to that noise object or combine it with other noise objects). Then create a object from your noise object, generating a concrete field of values that you can sample from directly or visualize using the or class.
//
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

// Alloc allocates a new instance without initialization.
func (pc _PerlinNoiseSourceClass) Alloc() PerlinNoiseSource {
	rv := objc.Send[PerlinNoiseSource](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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




// Initializes a Perlin noise source with the specified parameters.
//
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKPerlinNoiseSource/init(frequency:octaveCount:persistence:lacunarity:seed:)
func NewPerlinNoiseSourceWithFrequencyOctaveCountPersistenceLacunaritySeed(frequency float64, octaveCount int, persistence float64, lacunarity float64, seed unsafe.Pointer) PerlinNoiseSource {
	instance := getPerlinNoiseSourceClass().Alloc()
	rv := objc.Send[PerlinNoiseSource](instance.ID, objc.Sel("initWithFrequency:octaveCount:persistence:lacunarity:seed:"), frequency, octaveCount, persistence, lacunarity, seed)
	rv.Autorelease()
	return rv
}


// Creates a Perlin noise source with the specified parameters.
//
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKPerlinNoiseSource/perlinNoiseSourceWithFrequency:octaveCount:persistence:lacunarity:seed:
func (pc _PerlinNoiseSourceClass) PerlinNoiseSourceWithFrequencyOctaveCountPersistenceLacunaritySeed(frequency float64, octaveCount int, persistence float64, lacunarity float64, seed unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(pc.class), objc.Sel("perlinNoiseSourceWithFrequency:octaveCount:persistence:lacunarity:seed:"), frequency, octaveCount, persistence, lacunarity, seed)
	return rv
}

// The rate at which successive octaves of the noise function decrease in amplitude.
//
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKPerlinNoiseSource/persistence
func (p_ PerlinNoiseSource) Persistence() float64 {
	rv := objc.Send[float64](p_.ID, objc.Sel("persistence"))
	return rv
}


// SetPersistence sets the value of the persistence property.
// The rate at which successive octaves of the noise function decrease in amplitude.

//
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKPerlinNoiseSource/persistence
func (p_ PerlinNoiseSource) SetPersistence(value float64) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setPersistence:"), value)
}


