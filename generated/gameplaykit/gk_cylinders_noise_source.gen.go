// Code generated from Apple documentation for GameplayKit. DO NOT EDIT.

package gameplaykit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [CylindersNoiseSource] class.
var (
	CylindersNoiseSourceClass     _CylindersNoiseSourceClass
	CylindersNoiseSourceClassOnce sync.Once
)

func getCylindersNoiseSourceClass() _CylindersNoiseSourceClass {
	CylindersNoiseSourceClassOnce.Do(func() {
		CylindersNoiseSourceClass = _CylindersNoiseSourceClass{objc.GetClass("GKCylindersNoiseSource")}
	})
	return CylindersNoiseSourceClass
}

type _CylindersNoiseSourceClass struct {
	class objc.Class
}

// An interface definition for the [CylindersNoiseSource] class.
type ICylindersNoiseSource interface {
	INoiseSource
}

// A procedural noise generator whose output is a 3D field of concentric cylindrical shells.
//
// All noise sources generate infinite 3D fields of noise values, but this fact is especially relevant to cylinder noise: by rotating a noise object in 3D, you can sample the noise in ways that “slice” across or along the cylinders. Use this technique (combined with other noise sources and noise processing operations) to create effects such as wood-grain textures. Like all subclasses, a cylinder noise source represents a noise generation algorithm and its parameters. To make use of a noise source, first create object from it (and optionally apply operations to that noise object or combine it with other noise objects). Then create a object from your noise object, generating a concrete field of values that you can sample from directly or visualize using the or class.
//
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKCylindersNoiseSource
type CylindersNoiseSource struct {
	NoiseSource
}

// CylindersNoiseSourceFrom constructs a [CylindersNoiseSource] from an unsafe.Pointer.
//
// A procedural noise generator whose output is a 3D field of concentric cylindrical shells.
func CylindersNoiseSourceFrom(ptr unsafe.Pointer) CylindersNoiseSource {
	return CylindersNoiseSource{
		NoiseSource: NoiseSourceFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (cc _CylindersNoiseSourceClass) Alloc() CylindersNoiseSource {
	rv := objc.Send[CylindersNoiseSource](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CylindersNoiseSourceClass) New() CylindersNoiseSource {
	rv := objc.Send[CylindersNoiseSource](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CylindersNoiseSource) Init() CylindersNoiseSource {
	rv := objc.Send[CylindersNoiseSource](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CylindersNoiseSource) Autorelease() CylindersNoiseSource {
	rv := objc.Send[CylindersNoiseSource](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCylindersNoiseSource creates a new CylindersNoiseSource instance.
func NewCylindersNoiseSource() CylindersNoiseSource {
	return getCylindersNoiseSourceClass().New()
}


// Initializes a cylinder noise source with the specified frequency.
//
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKCylindersNoiseSource/init(frequency:)
func NewCylindersNoiseSourceWithFrequency(frequency unsafe.Pointer) CylindersNoiseSource {
	instance := getCylindersNoiseSourceClass().Alloc()
	rv := objc.Send[CylindersNoiseSource](instance.ID, objc.Sel("initWithFrequency:"), frequency)
	rv.Autorelease()
	return rv
}


// Creates a cylinder noise source with the specified frequency.
//
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKCylindersNoiseSource/cylindersNoise(withFrequency:)
func (cc _CylindersNoiseSourceClass) CylindersNoiseWithFrequency(frequency unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("cylindersNoiseWithFrequency:"), frequency)
	return rv
}

// A value that determines the size and spacing of concentric cylinders.
//
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKCylindersNoiseSource/frequency
func (c_ CylindersNoiseSource) Frequency() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("frequency"))
	return rv
}


// SetFrequency sets the value of the frequency property.
// A value that determines the size and spacing of concentric cylinders.

//
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKCylindersNoiseSource/frequency
func (c_ CylindersNoiseSource) SetFrequency(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setFrequency:"), value)
}

