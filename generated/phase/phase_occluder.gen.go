// Code generated from Apple documentation for PHASE. DO NOT EDIT.

package phase

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class PHASEOccluder */


/* debug [class_header]: Header for PHASEOccluder */
// The class instance for the [PHASEOccluder] class.
var (
	PHASEOccluderClass     _PHASEOccluderClass
	PHASEOccluderClassOnce sync.Once
)

func getPHASEOccluderClass() _PHASEOccluderClass {
	PHASEOccluderClassOnce.Do(func() {
		PHASEOccluderClass = _PHASEOccluderClass{objc.GetClass("PHASEOccluder")}
	})
	return PHASEOccluderClass
}

type _PHASEOccluderClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for PHASEOccluder */
// An interface definition for the [PHASEOccluder] class.
type IPHASEOccluder interface {
	IPHASEObject
	
/* debug [class_interface_properties]: Properties for PHASEOccluder */
	// properties:
	Shapes() []PHASEShape
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for PHASEOccluder */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for PHASEOccluder */
// Alloc allocates a new instance without initialization.
func (pc _PHASEOccluderClass) Alloc() PHASEOccluder {
	rv := objc.Send[PHASEOccluder](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (pc _PHASEOccluderClass) New() PHASEOccluder {
	rv := objc.Send[PHASEOccluder](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PHASEOccluder) Init() PHASEOccluder {
	rv := objc.Send[PHASEOccluder](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PHASEOccluder) Autorelease() PHASEOccluder {
	rv := objc.Send[PHASEOccluder](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPHASEOccluder creates a new PHASEOccluder instance.
func NewPHASEOccluder() PHASEOccluder {
	return getPHASEOccluderClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for PHASEOccluder */
// An object with a shape and position that blocks audio from reaching the listener.
//
// The framework lowers the volume of an audio signal when an instance of this class positions somewhere along the path between the sound source and the listener. For an example that demonstrates sound occlusion, see .


// An object with a shape and position that blocks audio from reaching the listener.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEOccluder
type PHASEOccluder struct {
	PHASEObject
}

// PHASEOccluderFrom constructs a [PHASEOccluder] from an unsafe.Pointer.
//
// An object with a shape and position that blocks audio from reaching the listener.
func PHASEOccluderFrom(ptr unsafe.Pointer) PHASEOccluder {
	return PHASEOccluder{
		PHASEObject: PHASEObjectFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for PHASEOccluder */

// Creates an occluder with the given engine and shapes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEOccluder/init(engine:shapes:)
func NewPHASEOccluderWithEngineShapes(engine IPHASEEngine, shapes []PHASEShape) PHASEOccluder {
	instance := getPHASEOccluderClass().Alloc()
	rv := objc.Send[PHASEOccluder](instance.ID, objc.Sel("initWithEngine:shapes:"), engine, shapes)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewPHASEOccluderWithEngineShapes */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for PHASEOccluder */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for PHASEOccluder */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for PHASEOccluder */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for PHASEOccluder */

// An array of shapes that collectively define the occluder’s audio-deflecting surface and texture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEOccluder/shapes
func (p_ PHASEOccluder) Shapes() []PHASEShape {
	rv := objc.Send[[]PHASEShape](p_.ID, objc.Sel("shapes"))
	return rv
}/* debug [instance_properties/getter]: shapes */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class PHASEOccluder */


