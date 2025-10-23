// Code generated from Apple documentation for GameplayKit. DO NOT EDIT.

package gameplaykit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [RidgedNoiseSource] class.
var (
	RidgedNoiseSourceClass     _RidgedNoiseSourceClass
	RidgedNoiseSourceClassOnce sync.Once
)

func getRidgedNoiseSourceClass() _RidgedNoiseSourceClass {
	RidgedNoiseSourceClassOnce.Do(func() {
		RidgedNoiseSourceClass = _RidgedNoiseSourceClass{objc.GetClass("GKRidgedNoiseSource")}
	})
	return RidgedNoiseSourceClass
}

type _RidgedNoiseSourceClass struct {
	class objc.Class
}

// An interface definition for the [RidgedNoiseSource] class.
type IRidgedNoiseSource interface {
	ICoherentNoiseSource
	// properties:
	// methods:
}

// A procedural noise generator whose output is a type of multifractal coherent noise with sharply defined features.
//
// Ridged noise is similar to Perlin noise (see the class), but with thinner features resembling natural phenomena such as forked lightning and mountain peaks. Like all subclasses, a ridged noise source represents a noise generation algorithm and its parameters. To make use of a noise source, first create object from it (and optionally apply operations to that noise object or combine it with other noise objects). Then create a object from your noise object, generating a concrete field of values that you can sample from directly or visualize using the or class.


// A procedural noise generator whose output is a type of multifractal coherent noise with sharply defined features.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKRidgedNoiseSource
type RidgedNoiseSource struct {
	CoherentNoiseSource
}

// RidgedNoiseSourceFrom constructs a [RidgedNoiseSource] from an unsafe.Pointer.
//
// A procedural noise generator whose output is a type of multifractal coherent noise with sharply defined features.
func RidgedNoiseSourceFrom(ptr unsafe.Pointer) RidgedNoiseSource {
	return RidgedNoiseSource{
		CoherentNoiseSource: CoherentNoiseSourceFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (rc _RidgedNoiseSourceClass) Alloc() RidgedNoiseSource {
	rv := objc.Send[RidgedNoiseSource](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (rc _RidgedNoiseSourceClass) New() RidgedNoiseSource {
	rv := objc.Send[RidgedNoiseSource](objc.ID(rc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (r_ RidgedNoiseSource) Init() RidgedNoiseSource {
	rv := objc.Send[RidgedNoiseSource](r_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r_ RidgedNoiseSource) Autorelease() RidgedNoiseSource {
	rv := objc.Send[RidgedNoiseSource](r_.ID, objc.Sel("autorelease"))
	return rv
}

// NewRidgedNoiseSource creates a new RidgedNoiseSource instance.
func NewRidgedNoiseSource() RidgedNoiseSource {
	return getRidgedNoiseSourceClass().New()
}



// Initializes a ridged noise source with the specified parameters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKRidgedNoiseSource/init(frequency:octaveCount:lacunarity:seed:)
func NewRidgedNoiseSourceWithFrequencyOctaveCountLacunaritySeed(frequency float64 /* primitive/slice/pointer. */, octaveCount int /* primitive/slice/pointer. */, lacunarity float64 /* primitive/slice/pointer. */, seed unsafe.Pointer) RidgedNoiseSource {
	instance := getRidgedNoiseSourceClass().Alloc()
	rv := objc.Send[RidgedNoiseSource](instance.ID, objc.Sel("initWithFrequency:octaveCount:lacunarity:seed:"), frequency, octaveCount, lacunarity, seed)
	rv.Autorelease()
	return rv
}



// Creates a ridged noise source with the specified parameters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKRidgedNoiseSource/ridgedNoiseSourceWithFrequency:octaveCount:lacunarity:seed:
func (rc _RidgedNoiseSourceClass) RidgedNoiseSourceWithFrequencyOctaveCountLacunaritySeed(frequency float64 /* primitive/slice/pointer. */, octaveCount int /* primitive/slice/pointer. */, lacunarity float64 /* primitive/slice/pointer. */, seed unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(rc.class), objc.Sel("ridgedNoiseSourceWithFrequency:octaveCount:lacunarity:seed:"), frequency, octaveCount, lacunarity, seed)
	return rv
}


