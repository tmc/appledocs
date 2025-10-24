// Code generated from Apple documentation for GameplayKit. DO NOT EDIT.

package gameplaykit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class GKCheckerboardNoiseSource */


/* debug [class_header]: Header for GKCheckerboardNoiseSource */
// The class instance for the [CheckerboardNoiseSource] class.
var (
	CheckerboardNoiseSourceClass     _CheckerboardNoiseSourceClass
	CheckerboardNoiseSourceClassOnce sync.Once
)

func getCheckerboardNoiseSourceClass() _CheckerboardNoiseSourceClass {
	CheckerboardNoiseSourceClassOnce.Do(func() {
		CheckerboardNoiseSourceClass = _CheckerboardNoiseSourceClass{objc.GetClass("GKCheckerboardNoiseSource")}
	})
	return CheckerboardNoiseSourceClass
}

type _CheckerboardNoiseSourceClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CheckerboardNoiseSource */
// An interface definition for the [CheckerboardNoiseSource] class.
type ICheckerboardNoiseSource interface {
	INoiseSource
	
/* debug [class_interface_properties]: Properties for CheckerboardNoiseSource */
	// properties:
	SquareSize() float64
	SetSquareSize(value float64)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CheckerboardNoiseSource */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CheckerboardNoiseSource */
// Alloc allocates a new instance without initialization.
func (cc _CheckerboardNoiseSourceClass) Alloc() CheckerboardNoiseSource {
	rv := objc.Send[CheckerboardNoiseSource](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CheckerboardNoiseSourceClass) New() CheckerboardNoiseSource {
	rv := objc.Send[CheckerboardNoiseSource](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CheckerboardNoiseSource) Init() CheckerboardNoiseSource {
	rv := objc.Send[CheckerboardNoiseSource](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CheckerboardNoiseSource) Autorelease() CheckerboardNoiseSource {
	rv := objc.Send[CheckerboardNoiseSource](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCheckerboardNoiseSource creates a new CheckerboardNoiseSource instance.
func NewCheckerboardNoiseSource() CheckerboardNoiseSource {
	return getCheckerboardNoiseSourceClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CheckerboardNoiseSource */
// A procedural noise generator whose output is an alternating square pattern.
//
// Checkerboard noise can be useful as an input to methods that create noise by combining other noise objects through various operations. Like all subclasses, a checkerboard noise source represents a noise generation algorithm and its parameters. To make use of a noise source, first create object from it (and optionally apply operations to that noise object or combine it with other noise objects). Then create a object from your noise object, generating a concrete field of values that you can sample from directly or visualize using the or class.


// A procedural noise generator whose output is an alternating square pattern.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKCheckerboardNoiseSource
type CheckerboardNoiseSource struct {
	NoiseSource
}

// CheckerboardNoiseSourceFrom constructs a [CheckerboardNoiseSource] from an unsafe.Pointer.
//
// A procedural noise generator whose output is an alternating square pattern.
func CheckerboardNoiseSourceFrom(ptr unsafe.Pointer) CheckerboardNoiseSource {
	return CheckerboardNoiseSource{
		NoiseSource: NoiseSourceFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CheckerboardNoiseSource */

// Initializes a checkerboard noise source with the specified square size.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKCheckerboardNoiseSource/init(squareSize:)
func NewCheckerboardNoiseSourceWithSquareSize(squareSize float64) CheckerboardNoiseSource {
	instance := getCheckerboardNoiseSourceClass().Alloc()
	rv := objc.Send[CheckerboardNoiseSource](instance.ID, objc.Sel("initWithSquareSize:"), squareSize)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCheckerboardNoiseSourceWithSquareSize */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CheckerboardNoiseSource */

// Creates a checkerboard noise source with the specified square size.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKCheckerboardNoiseSource/checkerboardNoise(withSquareSize:)
func (cc _CheckerboardNoiseSourceClass) CheckerboardNoiseWithSquareSize(squareSize float64) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(cc.class), objc.Sel("checkerboardNoiseWithSquareSize:"), squareSize)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=CheckerboardNoiseWithSquareSize) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CheckerboardNoiseSource */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CheckerboardNoiseSource */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CheckerboardNoiseSource */

// The size (both width and height) of squares in the generated checkerboard pattern.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKCheckerboardNoiseSource/squareSize
func (c_ CheckerboardNoiseSource) SquareSize() float64 {
	rv := objc.Send[float64](c_.ID, objc.Sel("squareSize"))
	return rv
}/* debug [instance_properties/getter]: squareSize */


// The size (both width and height) of squares in the generated checkerboard pattern.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKCheckerboardNoiseSource/squareSize
func (c_ CheckerboardNoiseSource) SetSquareSize(value float64) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSquareSize:"), value)
}/* debug [instance_properties/setter]: squareSize */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class GKCheckerboardNoiseSource */


