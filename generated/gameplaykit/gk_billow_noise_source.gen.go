// Code generated from Apple documentation for GameplayKit. DO NOT EDIT.

package gameplaykit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [BillowNoiseSource] class.
var (
	BillowNoiseSourceClass     _BillowNoiseSourceClass
	BillowNoiseSourceClassOnce sync.Once
)

func getBillowNoiseSourceClass() _BillowNoiseSourceClass {
	BillowNoiseSourceClassOnce.Do(func() {
		BillowNoiseSourceClass = _BillowNoiseSourceClass{objc.GetClass("GKBillowNoiseSource")}
	})
	return BillowNoiseSourceClass
}

type _BillowNoiseSourceClass struct {
	class objc.Class
}

// An interface definition for the [BillowNoiseSource] class.
type IBillowNoiseSource interface {
	ICoherentNoiseSource
}

// A procedural noise generator whose output is a type of fractal coherent noise with smooth features.
//
// Billow noise is similar to Perlin noise (see the class), but with more rounded features resembling natural phenomena such as treetops and hills. Like all subclasses, a billow noise source represents a noise generation algorithm and its parameters. To make use of a noise source, first create object from it (and optionally apply operations to that noise object or combine it with other noise objects). Then create a object from your noise object, generating a concrete field of values that you can sample from directly or visualize using the or class.
//
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKBillowNoiseSource
type BillowNoiseSource struct {
	CoherentNoiseSource
}

// BillowNoiseSourceFrom constructs a [BillowNoiseSource] from an unsafe.Pointer.
//
// A procedural noise generator whose output is a type of fractal coherent noise with smooth features.
func BillowNoiseSourceFrom(ptr unsafe.Pointer) BillowNoiseSource {
	return BillowNoiseSource{
		CoherentNoiseSource: CoherentNoiseSourceFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (bc _BillowNoiseSourceClass) Alloc() BillowNoiseSource {
	rv := objc.Send[BillowNoiseSource](objc.ID(bc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (bc _BillowNoiseSourceClass) New() BillowNoiseSource {
	rv := objc.Send[BillowNoiseSource](objc.ID(bc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (b_ BillowNoiseSource) Init() BillowNoiseSource {
	rv := objc.Send[BillowNoiseSource](b_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (b_ BillowNoiseSource) Autorelease() BillowNoiseSource {
	rv := objc.Send[BillowNoiseSource](b_.ID, objc.Sel("autorelease"))
	return rv
}

// NewBillowNoiseSource creates a new BillowNoiseSource instance.
func NewBillowNoiseSource() BillowNoiseSource {
	return getBillowNoiseSourceClass().New()
}




// Creates a billow noise source with the specified parameters.
//
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKBillowNoiseSource/init(frequency:octaveCount:persistence:lacunarity:seed:)
func NewBillowNoiseSourceWithFrequencyOctaveCountPersistenceLacunaritySeed(frequency unsafe.Pointer, octaveCount int, persistence unsafe.Pointer, lacunarity unsafe.Pointer, seed unsafe.Pointer) BillowNoiseSource {
	instance := getBillowNoiseSourceClass().Alloc()
	rv := objc.Send[BillowNoiseSource](instance.ID, objc.Sel("initWithFrequency:octaveCount:persistence:lacunarity:seed:"), frequency, octaveCount, persistence, lacunarity, seed)
	rv.Autorelease()
	return rv
}


// Initializes a billow noise source with the specified parameters.
//
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKBillowNoiseSource/billowNoiseSourceWithFrequency:octaveCount:persistence:lacunarity:seed:
func (bc _BillowNoiseSourceClass) BillowNoiseSourceWithFrequencyOctaveCountPersistenceLacunaritySeed(frequency unsafe.Pointer, octaveCount int, persistence unsafe.Pointer, lacunarity unsafe.Pointer, seed unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(bc.class), objc.Sel("billowNoiseSourceWithFrequency:octaveCount:persistence:lacunarity:seed:"), frequency, octaveCount, persistence, lacunarity, seed)
	return rv
}

// The rate at which successive octaves of the noise function decrease in amplitude.
//
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKBillowNoiseSource/persistence
func (b_ BillowNoiseSource) Persistence() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("persistence"))
	return rv
}


// SetPersistence sets the value of the persistence property.
// The rate at which successive octaves of the noise function decrease in amplitude.

//
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKBillowNoiseSource/persistence
func (b_ BillowNoiseSource) SetPersistence(value unsafe.Pointer) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setPersistence:"), value)
}


