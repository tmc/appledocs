// Code generated from Apple documentation for GameplayKit. DO NOT EDIT.

package gameplaykit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class GKVoronoiNoiseSource */


/* debug [class_header]: Header for GKVoronoiNoiseSource */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for VoronoiNoiseSource */
// An interface definition for the [VoronoiNoiseSource] class.
type IVoronoiNoiseSource interface {
	INoiseSource
	
/* debug [class_interface_properties]: Properties for VoronoiNoiseSource */
	// properties:
	Displacement() float64
	SetDisplacement(value float64)
	Frequency() float64
	SetFrequency(value float64)
	DistanceEnabled() bool
	SetDistanceEnabled(value bool)
	Seed() int32 /* not a class type */
	SetSeed(value int32 /* not a class type */)
	IsDistanceEnabled() bool
	SetIsDistanceEnabled(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for VoronoiNoiseSource */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for VoronoiNoiseSource */
// Alloc allocates a new instance without initialization.
func (vc _VoronoiNoiseSourceClass) Alloc() VoronoiNoiseSource {
	rv := objc.Send[VoronoiNoiseSource](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for VoronoiNoiseSource */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for VoronoiNoiseSource */

// Initializes a Voronoi noise source with the specified parameters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKVoronoiNoiseSource/init(frequency:displacement:distanceEnabled:seed:)
func NewVoronoiNoiseSourceWithFrequencyDisplacementDistanceEnabledSeed(frequency float64, displacement float64, distanceEnabled bool, seed int32 /* not a class type */) VoronoiNoiseSource {
	instance := getVoronoiNoiseSourceClass().Alloc()
	rv := objc.Send[VoronoiNoiseSource](instance.ID, objc.Sel("initWithFrequency:displacement:distanceEnabled:seed:"), frequency, displacement, distanceEnabled, seed)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewVoronoiNoiseSourceWithFrequencyDisplacementDistanceEnabledSeed */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for VoronoiNoiseSource */

// Creates a Voronoi noise source with the specified parameters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKVoronoiNoiseSource/voronoiNoise(withFrequency:displacement:distanceEnabled:seed:)
func (vc _VoronoiNoiseSourceClass) VoronoiNoiseWithFrequencyDisplacementDistanceEnabledSeed(frequency float64, displacement float64, distanceEnabled bool, seed int32 /* not a class type */) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(vc.class), objc.Sel("voronoiNoiseWithFrequency:displacement:distanceEnabled:seed:"), frequency, displacement, distanceEnabled, seed)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=VoronoiNoiseWithFrequencyDisplacementDistanceEnabledSeed) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for VoronoiNoiseSource */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for VoronoiNoiseSource */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for VoronoiNoiseSource */

// The range of random values to assign to each cell in generated noise.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKVoronoiNoiseSource/displacement
func (v_ VoronoiNoiseSource) Displacement() float64 {
	rv := objc.Send[float64](v_.ID, objc.Sel("displacement"))
	return rv
}/* debug [instance_properties/getter]: displacement */


// The range of random values to assign to each cell in generated noise.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKVoronoiNoiseSource/displacement
func (v_ VoronoiNoiseSource) SetDisplacement(value float64) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setDisplacement:"), value)
}/* debug [instance_properties/setter]: displacement */


// A value that determines the number and size of cells in generated noise.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKVoronoiNoiseSource/frequency
func (v_ VoronoiNoiseSource) Frequency() float64 {
	rv := objc.Send[float64](v_.ID, objc.Sel("frequency"))
	return rv
}/* debug [instance_properties/getter]: frequency */


// A value that determines the number and size of cells in generated noise.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKVoronoiNoiseSource/frequency
func (v_ VoronoiNoiseSource) SetFrequency(value float64) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setFrequency:"), value)
}/* debug [instance_properties/setter]: frequency */


// A Boolean value that specifies whether generated noise values incorporate the distance from each point to the nearest seed point.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKVoronoiNoiseSource/isDistanceEnabled
func (v_ VoronoiNoiseSource) DistanceEnabled() bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("distanceEnabled"))
	return rv
}/* debug [instance_properties/getter]: distanceEnabled */


// A Boolean value that specifies whether generated noise values incorporate the distance from each point to the nearest seed point.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKVoronoiNoiseSource/isDistanceEnabled
func (v_ VoronoiNoiseSource) SetDistanceEnabled(value bool) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setDistanceEnabled:"), value)
}/* debug [instance_properties/setter]: distanceEnabled */


// The value that determines the specific configuration of noise produced by the noise source.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKVoronoiNoiseSource/seed
func (v_ VoronoiNoiseSource) Seed() int32 /* not a class type */ {
	rv := objc.Send[int32](v_.ID, objc.Sel("seed"))
	return rv
}/* debug [instance_properties/getter]: seed */


// The value that determines the specific configuration of noise produced by the noise source.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKVoronoiNoiseSource/seed
func (v_ VoronoiNoiseSource) SetSeed(value int32 /* not a class type */) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setSeed:"), value)
}/* debug [instance_properties/setter]: seed */


// A Boolean value that specifies whether generated noise values incorporate the distance from each point to the nearest seed point.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gameplaykit/gkvoronoinoisesource/isdistanceenabled
func (v_ VoronoiNoiseSource) IsDistanceEnabled() bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("isDistanceEnabled"))
	return rv
}/* debug [instance_properties/getter]: isDistanceEnabled */


// A Boolean value that specifies whether generated noise values incorporate the distance from each point to the nearest seed point.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gameplaykit/gkvoronoinoisesource/isdistanceenabled
func (v_ VoronoiNoiseSource) SetIsDistanceEnabled(value bool) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setIsDistanceEnabled:"), value)
}/* debug [instance_properties/setter]: isDistanceEnabled */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class GKVoronoiNoiseSource */


