// Code generated from Apple documentation for GameplayKit. DO NOT EDIT.

package gameplaykit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [VoronoiNoiseSource] class.
var (
	VoronoiNoiseSourceClass     _VoronoiNoiseSourceClass
	VoronoiNoiseSourceClassOnce sync.Once
)

func getVoronoiNoiseSourceClass() _VoronoiNoiseSourceClass {
	VoronoiNoiseSourceClassOnce.Do(func() {
		VoronoiNoiseSourceClass = _VoronoiNoiseSourceClass{objc.GetClass("GKVoronoiNoiseSource")}
	})
	return VoronoiNoiseSourceClass
}

type _VoronoiNoiseSourceClass struct {
	class objc.Class
}

// An interface definition for the [VoronoiNoiseSource] class.
type IVoronoiNoiseSource interface {
	INoiseSource
	// properties:
	Displacement() float64 /* primitive/slice/pointer. */
	SetDisplacement(value float64 /* primitive/slice/pointer. */)
	Frequency() float64 /* primitive/slice/pointer. */
	SetFrequency(value float64 /* primitive/slice/pointer. */)
	DistanceEnabled() bool /* primitive/slice/pointer. */
	SetDistanceEnabled(value bool /* primitive/slice/pointer. */)
	Seed() unsafe.Pointer
	SetSeed(value unsafe.Pointer)
	IsDistanceEnabled() bool /* primitive/slice/pointer. */
	SetIsDistanceEnabled(value bool /* primitive/slice/pointer. */)
	// methods:
}

// A procedural noise generator whose output (also called Worley noise or cellular noise) divides space into discrete cells surrounding random seed points.
//
// Voronoi noise can generate textures resembling natural phenomena such as crystalline structures, cracked mud, or star fields. Like all subclasses, a Voronoi noise source represents a noise generation algorithm and its parameters. To make use of a noise source, first create object from it (and optionally apply operations to that noise object or combine it with other noise objects). Then create a object from your noise object, generating a concrete field of values that you can sample from directly or visualize using the or class.


// A procedural noise generator whose output (also called Worley noise or cellular noise) divides space into discrete cells surrounding random seed points.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKVoronoiNoiseSource
type VoronoiNoiseSource struct {
	NoiseSource
}

// VoronoiNoiseSourceFrom constructs a [VoronoiNoiseSource] from an unsafe.Pointer.
//
// A procedural noise generator whose output (also called Worley noise or cellular noise) divides space into discrete cells surrounding random seed points.
func VoronoiNoiseSourceFrom(ptr unsafe.Pointer) VoronoiNoiseSource {
	return VoronoiNoiseSource{
		NoiseSource: NoiseSourceFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (vc _VoronoiNoiseSourceClass) Alloc() VoronoiNoiseSource {
	rv := objc.Send[VoronoiNoiseSource](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (vc _VoronoiNoiseSourceClass) New() VoronoiNoiseSource {
	rv := objc.Send[VoronoiNoiseSource](objc.ID(vc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (v_ VoronoiNoiseSource) Init() VoronoiNoiseSource {
	rv := objc.Send[VoronoiNoiseSource](v_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v_ VoronoiNoiseSource) Autorelease() VoronoiNoiseSource {
	rv := objc.Send[VoronoiNoiseSource](v_.ID, objc.Sel("autorelease"))
	return rv
}

// NewVoronoiNoiseSource creates a new VoronoiNoiseSource instance.
func NewVoronoiNoiseSource() VoronoiNoiseSource {
	return getVoronoiNoiseSourceClass().New()
}



// Initializes a Voronoi noise source with the specified parameters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKVoronoiNoiseSource/init(frequency:displacement:distanceEnabled:seed:)
func NewVoronoiNoiseSourceWithFrequencyDisplacementDistanceEnabledSeed(frequency float64 /* primitive/slice/pointer. */, displacement float64 /* primitive/slice/pointer. */, distanceEnabled bool /* primitive/slice/pointer. */, seed unsafe.Pointer) VoronoiNoiseSource {
	instance := getVoronoiNoiseSourceClass().Alloc()
	rv := objc.Send[VoronoiNoiseSource](instance.ID, objc.Sel("initWithFrequency:displacement:distanceEnabled:seed:"), frequency, displacement, distanceEnabled, seed)
	rv.Autorelease()
	return rv
}



// Creates a Voronoi noise source with the specified parameters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKVoronoiNoiseSource/voronoiNoise(withFrequency:displacement:distanceEnabled:seed:)
func (vc _VoronoiNoiseSourceClass) VoronoiNoiseWithFrequencyDisplacementDistanceEnabledSeed(frequency float64 /* primitive/slice/pointer. */, displacement float64 /* primitive/slice/pointer. */, distanceEnabled bool /* primitive/slice/pointer. */, seed unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(vc.class), objc.Sel("voronoiNoiseWithFrequency:displacement:distanceEnabled:seed:"), frequency, displacement, distanceEnabled, seed)
	return rv
}


// The range of random values to assign to each cell in generated noise.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKVoronoiNoiseSource/displacement
func (v_ VoronoiNoiseSource) Displacement() float64 /* primitive/slice/pointer. */ {
	rv := objc.Send[float64](v_.ID, objc.Sel("displacement"))
	return rv
}


// The range of random values to assign to each cell in generated noise.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKVoronoiNoiseSource/displacement
func (v_ VoronoiNoiseSource) SetDisplacement(value float64 /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setDisplacement:"), value)
}


// A value that determines the number and size of cells in generated noise.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKVoronoiNoiseSource/frequency
func (v_ VoronoiNoiseSource) Frequency() float64 /* primitive/slice/pointer. */ {
	rv := objc.Send[float64](v_.ID, objc.Sel("frequency"))
	return rv
}


// A value that determines the number and size of cells in generated noise.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKVoronoiNoiseSource/frequency
func (v_ VoronoiNoiseSource) SetFrequency(value float64 /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setFrequency:"), value)
}


// A Boolean value that specifies whether generated noise values incorporate the distance from each point to the nearest seed point.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKVoronoiNoiseSource/isDistanceEnabled
func (v_ VoronoiNoiseSource) DistanceEnabled() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](v_.ID, objc.Sel("distanceEnabled"))
	return rv
}


// A Boolean value that specifies whether generated noise values incorporate the distance from each point to the nearest seed point.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKVoronoiNoiseSource/isDistanceEnabled
func (v_ VoronoiNoiseSource) SetDistanceEnabled(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setDistanceEnabled:"), value)
}


// The value that determines the specific configuration of noise produced by the noise source.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKVoronoiNoiseSource/seed
func (v_ VoronoiNoiseSource) Seed() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID, objc.Sel("seed"))
	return rv
}


// The value that determines the specific configuration of noise produced by the noise source.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKVoronoiNoiseSource/seed
func (v_ VoronoiNoiseSource) SetSeed(value unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setSeed:"), value)
}


// A Boolean value that specifies whether generated noise values incorporate the distance from each point to the nearest seed point.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gameplaykit/gkvoronoinoisesource/isdistanceenabled
func (v_ VoronoiNoiseSource) IsDistanceEnabled() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](v_.ID, objc.Sel("isDistanceEnabled"))
	return rv
}


// A Boolean value that specifies whether generated noise values incorporate the distance from each point to the nearest seed point.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gameplaykit/gkvoronoinoisesource/isdistanceenabled
func (v_ VoronoiNoiseSource) SetIsDistanceEnabled(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setIsDistanceEnabled:"), value)
}


