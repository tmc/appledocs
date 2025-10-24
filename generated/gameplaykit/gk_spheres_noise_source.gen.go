// Code generated from Apple documentation for GameplayKit. DO NOT EDIT.

package gameplaykit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class GKSpheresNoiseSource */


/* debug [class_header]: Header for GKSpheresNoiseSource */
// The class instance for the [SpheresNoiseSource] class.
var (
	SpheresNoiseSourceClass     _SpheresNoiseSourceClass
	SpheresNoiseSourceClassOnce sync.Once
)

func getSpheresNoiseSourceClass() _SpheresNoiseSourceClass {
	SpheresNoiseSourceClassOnce.Do(func() {
		SpheresNoiseSourceClass = _SpheresNoiseSourceClass{objc.GetClass("GKSpheresNoiseSource")}
	})
	return SpheresNoiseSourceClass
}

type _SpheresNoiseSourceClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for SpheresNoiseSource */
// An interface definition for the [SpheresNoiseSource] class.
type ISpheresNoiseSource interface {
	INoiseSource
	
/* debug [class_interface_properties]: Properties for SpheresNoiseSource */
	// properties:
	Frequency() float64
	SetFrequency(value float64)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for SpheresNoiseSource */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for SpheresNoiseSource */
// Alloc allocates a new instance without initialization.
func (sc _SpheresNoiseSourceClass) Alloc() SpheresNoiseSource {
	rv := objc.Send[SpheresNoiseSource](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (sc _SpheresNoiseSourceClass) New() SpheresNoiseSource {
	rv := objc.Send[SpheresNoiseSource](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SpheresNoiseSource) Init() SpheresNoiseSource {
	rv := objc.Send[SpheresNoiseSource](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SpheresNoiseSource) Autorelease() SpheresNoiseSource {
	rv := objc.Send[SpheresNoiseSource](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSpheresNoiseSource creates a new SpheresNoiseSource instance.
func NewSpheresNoiseSource() SpheresNoiseSource {
	return getSpheresNoiseSourceClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for SpheresNoiseSource */
// A procedural noise generator whose output is a 3D field of concentric spherical shells.
//
// All noise sources generate infinite 3D fields of noise values, but this fact is especially relevant to sphere noise: by transforming a noise object in 3D, you can sample the noise in ways that “slice” through the spheres. Use this technique (combined with other noise sources and noise processing operations) to create effects such as wood-grain textures. Like all subclasses, a sphere noise source represents a noise generation algorithm and its parameters. To make use of a noise source, first create object from it (and optionally apply operations to that noise object or combine it with other noise objects). Then create a object from your noise object, generating a concrete field of values that you can sample from directly or visualize using the or class.


// A procedural noise generator whose output is a 3D field of concentric spherical shells.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKSpheresNoiseSource
type SpheresNoiseSource struct {
	NoiseSource
}

// SpheresNoiseSourceFrom constructs a [SpheresNoiseSource] from an unsafe.Pointer.
//
// A procedural noise generator whose output is a 3D field of concentric spherical shells.
func SpheresNoiseSourceFrom(ptr unsafe.Pointer) SpheresNoiseSource {
	return SpheresNoiseSource{
		NoiseSource: NoiseSourceFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for SpheresNoiseSource */

// Initializes a sphere noise source with the specified frequency.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKSpheresNoiseSource/init(frequency:)
func NewSpheresNoiseSourceWithFrequency(frequency float64) SpheresNoiseSource {
	instance := getSpheresNoiseSourceClass().Alloc()
	rv := objc.Send[SpheresNoiseSource](instance.ID, objc.Sel("initWithFrequency:"), frequency)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewSpheresNoiseSourceWithFrequency */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for SpheresNoiseSource */

// Creates a sphere noise source with the specified frequency.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKSpheresNoiseSource/spheresNoise(withFrequency:)
func (sc _SpheresNoiseSourceClass) SpheresNoiseWithFrequency(frequency float64) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(sc.class), objc.Sel("spheresNoiseWithFrequency:"), frequency)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=SpheresNoiseWithFrequency) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for SpheresNoiseSource */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for SpheresNoiseSource */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for SpheresNoiseSource */

// A value that determines the size and spacing of concentric spheres.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKSpheresNoiseSource/frequency
func (s_ SpheresNoiseSource) Frequency() float64 {
	rv := objc.Send[float64](s_.ID, objc.Sel("frequency"))
	return rv
}/* debug [instance_properties/getter]: frequency */


// A value that determines the size and spacing of concentric spheres.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKSpheresNoiseSource/frequency
func (s_ SpheresNoiseSource) SetFrequency(value float64) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setFrequency:"), value)
}/* debug [instance_properties/setter]: frequency */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class GKSpheresNoiseSource */


