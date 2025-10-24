// Code generated from Apple documentation for GameplayKit. DO NOT EDIT.

package gameplaykit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class GKCylindersNoiseSource */


/* debug [class_header]: Header for GKCylindersNoiseSource */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CylindersNoiseSource */
// An interface definition for the [CylindersNoiseSource] class.
type ICylindersNoiseSource interface {
	INoiseSource
	
/* debug [class_interface_properties]: Properties for CylindersNoiseSource */
	// properties:
	Frequency() float64
	SetFrequency(value float64)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CylindersNoiseSource */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CylindersNoiseSource */
// Alloc allocates a new instance without initialization.
func (cc _CylindersNoiseSourceClass) Alloc() CylindersNoiseSource {
	rv := objc.Send[CylindersNoiseSource](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CylindersNoiseSource */
// A procedural noise generator whose output is a 3D field of concentric cylindrical shells.
//
// All noise sources generate infinite 3D fields of noise values, but this fact is especially relevant to cylinder noise: by rotating a noise object in 3D, you can sample the noise in ways that “slice” across or along the cylinders. Use this technique (combined with other noise sources and noise processing operations) to create effects such as wood-grain textures. Like all subclasses, a cylinder noise source represents a noise generation algorithm and its parameters. To make use of a noise source, first create object from it (and optionally apply operations to that noise object or combine it with other noise objects). Then create a object from your noise object, generating a concrete field of values that you can sample from directly or visualize using the or class.


// A procedural noise generator whose output is a 3D field of concentric cylindrical shells.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CylindersNoiseSource */

// Initializes a cylinder noise source with the specified frequency.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKCylindersNoiseSource/init(frequency:)
func NewCylindersNoiseSourceWithFrequency(frequency float64) CylindersNoiseSource {
	instance := getCylindersNoiseSourceClass().Alloc()
	rv := objc.Send[CylindersNoiseSource](instance.ID, objc.Sel("initWithFrequency:"), frequency)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCylindersNoiseSourceWithFrequency */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CylindersNoiseSource */

// Creates a cylinder noise source with the specified frequency.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKCylindersNoiseSource/cylindersNoise(withFrequency:)
func (cc _CylindersNoiseSourceClass) CylindersNoiseWithFrequency(frequency float64) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(cc.class), objc.Sel("cylindersNoiseWithFrequency:"), frequency)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=CylindersNoiseWithFrequency) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CylindersNoiseSource */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CylindersNoiseSource */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CylindersNoiseSource */

// A value that determines the size and spacing of concentric cylinders.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKCylindersNoiseSource/frequency
func (c_ CylindersNoiseSource) Frequency() float64 {
	rv := objc.Send[float64](c_.ID, objc.Sel("frequency"))
	return rv
}/* debug [instance_properties/getter]: frequency */


// A value that determines the size and spacing of concentric cylinders.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKCylindersNoiseSource/frequency
func (c_ CylindersNoiseSource) SetFrequency(value float64) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setFrequency:"), value)
}/* debug [instance_properties/setter]: frequency */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class GKCylindersNoiseSource */


