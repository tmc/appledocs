// Code generated from Apple documentation for PHASE. DO NOT EDIT.

package phase

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class PHASESource */


/* debug [class_header]: Header for PHASESource */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for PHASESource */
// An interface definition for the [PHASESource] class.
type IPHASESource interface {
	IPHASEObject
	
/* debug [class_interface_properties]: Properties for PHASESource */
	// properties:
	Gain() float64
	SetGain(value float64)
	Shapes() []PHASEShape
	Transform() unsafe.Pointer
	SetTransform(value unsafe.Pointer)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for PHASESource */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for PHASESource */
// Alloc allocates a new instance without initialization.
func (pc _PHASESourceClass) Alloc() PHASESource {
	rv := objc.Send[PHASESource](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for PHASESource */
// An object that plays audio from a 3D location and orientation in a scene.
//
// This class represents a sound-emitting point or area in a virtual environment, positioned and oriented by a 3D . A spatial mixer, , adds environmental effects to sound sources. To tie a mixer to a sound source, create a object and pass it into the argument of a sound event’s initializer. For an example that demonstrates sound sources, see .


// An object that plays audio from a 3D location and orientation in a scene.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for PHASESource */

// Creates a single point in the environment from which sound emanates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASESource/init(engine:)
func NewPHASESourceWithEngine(engine IPHASEEngine) PHASESource {
	instance := getPHASESourceClass().Alloc()
	rv := objc.Send[PHASESource](instance.ID, objc.Sel("initWithEngine:"), engine)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewPHASESourceWithEngine */


// Creates a voluminous area in the environment from which sound emanates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASESource/init(engine:shapes:)
func NewPHASESourceWithEngineShapes(engine IPHASEEngine, shapes []PHASEShape) PHASESource {
	instance := getPHASESourceClass().Alloc()
	rv := objc.Send[PHASESource](instance.ID, objc.Sel("initWithEngine:shapes:"), engine, shapes)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewPHASESourceWithEngineShapes */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for PHASESource */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for PHASESource */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for PHASESource */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for PHASESource */

// The amount of sound the source emanates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASESource/gain
func (p_ PHASESource) Gain() float64 {
	rv := objc.Send[float64](p_.ID, objc.Sel("gain"))
	return rv
}/* debug [instance_properties/getter]: gain */


// The amount of sound the source emanates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASESource/gain
func (p_ PHASESource) SetGain(value float64) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setGain:"), value)
}/* debug [instance_properties/setter]: gain */


// An array of shapes that collectively define the audio-emitting surface area of a volumetric source.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASESource/shapes
func (p_ PHASESource) Shapes() []PHASEShape {
	rv := objc.Send[[]PHASEShape](p_.ID, objc.Sel("shapes"))
	return rv
}/* debug [instance_properties/getter]: shapes */


// A matrix, in local coordinates, that determines the object’s pose in the scene.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/phase/phaseobject/transform
func (p_ PHASESource) Transform() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("transform"))
	return rv
}/* debug [instance_properties/getter]: transform */


// A matrix, in local coordinates, that determines the object’s pose in the scene.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/phase/phaseobject/transform
func (p_ PHASESource) SetTransform(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setTransform:"), value)
}/* debug [instance_properties/setter]: transform */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class PHASESource */


