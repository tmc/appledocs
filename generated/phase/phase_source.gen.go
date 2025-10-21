// Code generated from Apple documentation for PHASE. DO NOT EDIT.

package phase

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [PHASESource] class.
var (
	PHASESourceClass     _PHASESourceClass
	PHASESourceClassOnce sync.Once
)

func getPHASESourceClass() _PHASESourceClass {
	PHASESourceClassOnce.Do(func() {
		PHASESourceClass = _PHASESourceClass{objc.GetClass("PHASESource")}
	})
	return PHASESourceClass
}

type _PHASESourceClass struct {
	class objc.Class
}

// An interface definition for the [PHASESource] class.
type IPHASESource interface {
	IPHASEObject
}

// An object that plays audio from a 3D location and orientation in a scene.
//
// This class represents a sound-emitting point or area in a virtual environment, positioned and oriented by a 3D . A spatial mixer, , adds environmental effects to sound sources. To tie a mixer to a sound source, create a object and pass it into the argument of a sound event’s initializer. For an example that demonstrates sound sources, see .
//
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASESource
type PHASESource struct {
	PHASEObject
}

// PHASESourceFrom constructs a [PHASESource] from an unsafe.Pointer.
//
// An object that plays audio from a 3D location and orientation in a scene.
func PHASESourceFrom(ptr unsafe.Pointer) PHASESource {
	return PHASESource{
		PHASEObject: PHASEObjectFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (pc _PHASESourceClass) Alloc() PHASESource {
	rv := objc.Send[PHASESource](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PHASESourceClass) New() PHASESource {
	rv := objc.Send[PHASESource](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PHASESource) Init() PHASESource {
	rv := objc.Send[PHASESource](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PHASESource) Autorelease() PHASESource {
	rv := objc.Send[PHASESource](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPHASESource creates a new PHASESource instance.
func NewPHASESource() PHASESource {
	return getPHASESourceClass().New()
}




// Creates a single point in the environment from which sound emanates.
//
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASESource/init(engine:)
func NewPHASESourceWithEngine(engine unsafe.Pointer) PHASESource {
	instance := getPHASESourceClass().Alloc()
	rv := objc.Send[PHASESource](instance.ID, objc.Sel("initWithEngine:"), engine)
	rv.Autorelease()
	return rv
}



// Creates a voluminous area in the environment from which sound emanates.
//
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASESource/init(engine:shapes:)
func NewPHASESourceWithEngineShapes(engine unsafe.Pointer, shapes unsafe.Pointer) PHASESource {
	instance := getPHASESourceClass().Alloc()
	rv := objc.Send[PHASESource](instance.ID, objc.Sel("initWithEngine:shapes:"), engine, shapes)
	rv.Autorelease()
	return rv
}


// A matrix, in local coordinates, that determines the object’s pose in the scene.
//
// [Full Topic]: https://developer.apple.com/documentation/phase/phaseobject/transform
func (p_ PHASESource) Transform() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("transform"))
	return rv
}


// SetTransform sets the value of the transform property.
// A matrix, in local coordinates, that determines the object’s pose in the scene.

//
// [Full Topic]: https://developer.apple.com/documentation/phase/phaseobject/transform
func (p_ PHASESource) SetTransform(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setTransform:"), value)
}

// The amount of sound the source emanates.
//
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASESource/gain
func (p_ PHASESource) Gain() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("gain"))
	return rv
}


// SetGain sets the value of the gain property.
// The amount of sound the source emanates.

//
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASESource/gain
func (p_ PHASESource) SetGain(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setGain:"), value)
}

// An array of shapes that collectively define the audio-emitting surface area of a volumetric source.
//
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASESource/shapes
func (p_ PHASESource) Shapes() []PHASEShape {
	rv := objc.Send[[]PHASEShape](p_.ID, objc.Sel("shapes"))
	return rv
}


