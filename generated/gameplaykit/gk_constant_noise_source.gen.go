// Code generated from Apple documentation for GameplayKit. DO NOT EDIT.

package gameplaykit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [ConstantNoiseSource] class.
var (
	ConstantNoiseSourceClass     _ConstantNoiseSourceClass
	ConstantNoiseSourceClassOnce sync.Once
)

func getConstantNoiseSourceClass() _ConstantNoiseSourceClass {
	ConstantNoiseSourceClassOnce.Do(func() {
		ConstantNoiseSourceClass = _ConstantNoiseSourceClass{objc.GetClass("GKConstantNoiseSource")}
	})
	return ConstantNoiseSourceClass
}

type _ConstantNoiseSourceClass struct {
	class objc.Class
}

// An interface definition for the [ConstantNoiseSource] class.
type IConstantNoiseSource interface {
	INoiseSource
	Value() float64
	SetValue(value float64)
}

// A procedural noise generator that outputs a field of a single constant value.
//
// Constant noise can be useful as an input to methods that create noise by combining other noise objects through various operations. For example, when using the method you can pass constant noise for some parameters and non-constant noise for other parameters, resulting in a variable displacement along one axis but constant or no displacement along the others. Like all subclasses, a constant noise source represents a noise generation algorithm and its parameters. To make use of a noise source, first create object from it (and optionally apply operations to that noise object or combine it with other noise objects). Then create a object from your noise object, generating a concrete field of values that you can sample from directly or visualize using the or class.
//
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKConstantNoiseSource
type ConstantNoiseSource struct {
	NoiseSource
}

// ConstantNoiseSourceFrom constructs a [ConstantNoiseSource] from an unsafe.Pointer.
//
// A procedural noise generator that outputs a field of a single constant value.
func ConstantNoiseSourceFrom(ptr unsafe.Pointer) ConstantNoiseSource {
	return ConstantNoiseSource{
		NoiseSource: NoiseSourceFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (cc _ConstantNoiseSourceClass) Alloc() ConstantNoiseSource {
	rv := objc.Send[ConstantNoiseSource](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _ConstantNoiseSourceClass) New() ConstantNoiseSource {
	rv := objc.Send[ConstantNoiseSource](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ ConstantNoiseSource) Init() ConstantNoiseSource {
	rv := objc.Send[ConstantNoiseSource](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ ConstantNoiseSource) Autorelease() ConstantNoiseSource {
	rv := objc.Send[ConstantNoiseSource](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewConstantNoiseSource creates a new ConstantNoiseSource instance.
func NewConstantNoiseSource() ConstantNoiseSource {
	return getConstantNoiseSourceClass().New()
}




// Initializes a noise source with the specified constant value.
//
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKConstantNoiseSource/init(value:)
func NewConstantNoiseSourceWithValue(value float64) ConstantNoiseSource {
	instance := getConstantNoiseSourceClass().Alloc()
	rv := objc.Send[ConstantNoiseSource](instance.ID, objc.Sel("initWithValue:"), value)
	rv.Autorelease()
	return rv
}


// Creates a noise source with the specified constant value.
//
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKConstantNoiseSource/constantNoise(withValue:)
func (cc _ConstantNoiseSourceClass) ConstantNoiseWithValue(value float64) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("constantNoiseWithValue:"), value)
	return rv
}

// The constant value for the generated noise.
//
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKConstantNoiseSource/value
func (c_ ConstantNoiseSource) Value() float64 {
	rv := objc.Send[float64](c_.ID, objc.Sel("value"))
	return rv
}


// SetValue sets the value of the value property.
// The constant value for the generated noise.

//
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKConstantNoiseSource/value
func (c_ ConstantNoiseSource) SetValue(value float64) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setValue:"), value)
}


