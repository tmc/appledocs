// Code generated from Apple documentation for GameplayKit. DO NOT EDIT.

package gameplaykit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

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

// An interface definition for the [CheckerboardNoiseSource] class.
type ICheckerboardNoiseSource interface {
	INoiseSource
}

// A procedural noise generator whose output is an alternating square pattern.
//
// Checkerboard noise can be useful as an input to methods that create noise by combining other noise objects through various operations. Like all subclasses, a checkerboard noise source represents a noise generation algorithm and its parameters. To make use of a noise source, first create object from it (and optionally apply operations to that noise object or combine it with other noise objects). Then create a object from your noise object, generating a concrete field of values that you can sample from directly or visualize using the or class.
//
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

// Alloc allocates a new instance without initialization.
func (cc _CheckerboardNoiseSourceClass) Alloc() CheckerboardNoiseSource {
	rv := objc.Send[CheckerboardNoiseSource](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


// Initializes a checkerboard noise source with the specified square size.
//
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKCheckerboardNoiseSource/init(squareSize:)
func NewCheckerboardNoiseSourceWithSquareSize(squareSize unsafe.Pointer) CheckerboardNoiseSource {
	instance := getCheckerboardNoiseSourceClass().Alloc()
	rv := objc.Send[CheckerboardNoiseSource](instance.ID, objc.Sel("initWithSquareSize:"), squareSize)
	rv.Autorelease()
	return rv
}


// Creates a checkerboard noise source with the specified square size.
//
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKCheckerboardNoiseSource/checkerboardNoise(withSquareSize:)
func (cc _CheckerboardNoiseSourceClass) CheckerboardNoiseWithSquareSize(squareSize unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("checkerboardNoiseWithSquareSize:"), squareSize)
	return rv
}

// The size (both width and height) of squares in the generated checkerboard pattern.
//
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKCheckerboardNoiseSource/squareSize
func (c_ CheckerboardNoiseSource) SquareSize() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("squareSize"))
	return rv
}


// SetSquareSize sets the value of the squareSize property.
// The size (both width and height) of squares in the generated checkerboard pattern.

//
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKCheckerboardNoiseSource/squareSize
func (c_ CheckerboardNoiseSource) SetSquareSize(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSquareSize:"), value)
}

