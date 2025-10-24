// Code generated from Apple documentation for GameplayKit. DO NOT EDIT.

package gameplaykit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class GKBillowNoiseSource */


/* debug [class_header]: Header for GKBillowNoiseSource */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for BillowNoiseSource */
// An interface definition for the [BillowNoiseSource] class.
type IBillowNoiseSource interface {
	ICoherentNoiseSource
	
/* debug [class_interface_properties]: Properties for BillowNoiseSource */
	// properties:
	Persistence() float64
	SetPersistence(value float64)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for BillowNoiseSource */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for BillowNoiseSource */
// Alloc allocates a new instance without initialization.
func (bc _BillowNoiseSourceClass) Alloc() BillowNoiseSource {
	rv := objc.Send[BillowNoiseSource](objc.ID(bc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for BillowNoiseSource */
// A procedural noise generator whose output is a type of fractal coherent noise with smooth features.
//
// Billow noise is similar to Perlin noise (see the class), but with more rounded features resembling natural phenomena such as treetops and hills. Like all subclasses, a billow noise source represents a noise generation algorithm and its parameters. To make use of a noise source, first create object from it (and optionally apply operations to that noise object or combine it with other noise objects). Then create a object from your noise object, generating a concrete field of values that you can sample from directly or visualize using the or class.


// A procedural noise generator whose output is a type of fractal coherent noise with smooth features.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for BillowNoiseSource */

// Creates a billow noise source with the specified parameters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKBillowNoiseSource/init(frequency:octaveCount:persistence:lacunarity:seed:)
func NewBillowNoiseSourceWithFrequencyOctaveCountPersistenceLacunaritySeed(frequency float64, octaveCount int, persistence float64, lacunarity float64, seed int32 /* not a class type */) BillowNoiseSource {
	instance := getBillowNoiseSourceClass().Alloc()
	rv := objc.Send[BillowNoiseSource](instance.ID, objc.Sel("initWithFrequency:octaveCount:persistence:lacunarity:seed:"), frequency, octaveCount, persistence, lacunarity, seed)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewBillowNoiseSourceWithFrequencyOctaveCountPersistenceLacunaritySeed */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for BillowNoiseSource */

// Initializes a billow noise source with the specified parameters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKBillowNoiseSource/billowNoiseSourceWithFrequency:octaveCount:persistence:lacunarity:seed:
func (bc _BillowNoiseSourceClass) BillowNoiseSourceWithFrequencyOctaveCountPersistenceLacunaritySeed(frequency float64, octaveCount int, persistence float64, lacunarity float64, seed int32 /* not a class type */) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(bc.class), objc.Sel("billowNoiseSourceWithFrequency:octaveCount:persistence:lacunarity:seed:"), frequency, octaveCount, persistence, lacunarity, seed)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=BillowNoiseSourceWithFrequencyOctaveCountPersistenceLacunaritySeed) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for BillowNoiseSource */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for BillowNoiseSource */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for BillowNoiseSource */

// The rate at which successive octaves of the noise function decrease in amplitude.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKBillowNoiseSource/persistence
func (b_ BillowNoiseSource) Persistence() float64 {
	rv := objc.Send[float64](b_.ID, objc.Sel("persistence"))
	return rv
}/* debug [instance_properties/getter]: persistence */


// The rate at which successive octaves of the noise function decrease in amplitude.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKBillowNoiseSource/persistence
func (b_ BillowNoiseSource) SetPersistence(value float64) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setPersistence:"), value)
}/* debug [instance_properties/setter]: persistence */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class GKBillowNoiseSource */


