// Code generated from Apple documentation for GameplayKit. DO NOT EDIT.

package gameplaykit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

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

// An interface definition for the [CoherentNoiseSource] class.
type ICoherentNoiseSource interface {
	INoiseSource
	Frequency() float64
	SetFrequency(value float64)
	Lacunarity() float64
	SetLacunarity(value float64)
	OctaveCount() int
	SetOctaveCount(value int)
	Seed() unsafe.Pointer
	SetSeed(value unsafe.Pointer)
}

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

// Alloc allocates a new instance without initialization.
func (cc _CoherentNoiseSourceClass) Alloc() CoherentNoiseSource {
	rv := objc.Send[CoherentNoiseSource](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// A value that determines the size and spacing of features in generated noise.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKCoherentNoiseSource/frequency
func (c_ CoherentNoiseSource) Frequency() float64 {
	rv := objc.Send[float64](c_.ID, objc.Sel("frequency"))
	return rv
}


// A value that determines the size and spacing of features in generated noise.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKCoherentNoiseSource/frequency
func (c_ CoherentNoiseSource) SetFrequency(value float64) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setFrequency:"), value)
}


// The rate at which successive octaves of the noise function increase in frequency.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKCoherentNoiseSource/lacunarity
func (c_ CoherentNoiseSource) Lacunarity() float64 {
	rv := objc.Send[float64](c_.ID, objc.Sel("lacunarity"))
	return rv
}


// The rate at which successive octaves of the noise function increase in frequency.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKCoherentNoiseSource/lacunarity
func (c_ CoherentNoiseSource) SetLacunarity(value float64) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setLacunarity:"), value)
}


// The number of octaves of the underlying noise function to use for generating noise.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKCoherentNoiseSource/octaveCount
func (c_ CoherentNoiseSource) OctaveCount() int {
	rv := objc.Send[int](c_.ID, objc.Sel("octaveCount"))
	return rv
}


// The number of octaves of the underlying noise function to use for generating noise.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKCoherentNoiseSource/octaveCount
func (c_ CoherentNoiseSource) SetOctaveCount(value int) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setOctaveCount:"), value)
}


// The value that determines the specific configuration of noise produced by the noise source.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKCoherentNoiseSource/seed
func (c_ CoherentNoiseSource) Seed() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("seed"))
	return rv
}


// The value that determines the specific configuration of noise produced by the noise source.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKCoherentNoiseSource/seed
func (c_ CoherentNoiseSource) SetSeed(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSeed:"), value)
}



