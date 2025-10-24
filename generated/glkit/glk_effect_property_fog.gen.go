// Code generated from Apple documentation for GLKit. DO NOT EDIT.

package glkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class GLKEffectPropertyFog */


/* debug [class_header]: Header for GLKEffectPropertyFog */
// The class instance for the [GLKEffectPropertyFog] class.
var (
	GLKEffectPropertyFogClass     _GLKEffectPropertyFogClass
	GLKEffectPropertyFogClassOnce sync.Once
)

func getGLKEffectPropertyFogClass() _GLKEffectPropertyFogClass {
	GLKEffectPropertyFogClassOnce.Do(func() {
		GLKEffectPropertyFogClass = _GLKEffectPropertyFogClass{objc.GetClass("GLKEffectPropertyFog")}
	})
	return GLKEffectPropertyFogClass
}

type _GLKEffectPropertyFogClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for GLKEffectPropertyFog */
// An interface definition for the [GLKEffectPropertyFog] class.
type IGLKEffectPropertyFog interface {
	IGLKEffectProperty
	
/* debug [class_interface_properties]: Properties for GLKEffectPropertyFog */
	// properties:
	Color() GLKVector4 /* typedef */
	SetColor(value GLKVector4 /* typedef */)
	Density() unsafe.Pointer
	SetDensity(value unsafe.Pointer)
	Enabled() unsafe.Pointer
	SetEnabled(value unsafe.Pointer)
	End() unsafe.Pointer
	SetEnd(value unsafe.Pointer)
	Mode() unsafe.Pointer
	SetMode(value unsafe.Pointer)
	Start() unsafe.Pointer
	SetStart(value unsafe.Pointer)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for GLKEffectPropertyFog */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for GLKEffectPropertyFog */
// Alloc allocates a new instance without initialization.
func (gc _GLKEffectPropertyFogClass) Alloc() GLKEffectPropertyFog {
	rv := objc.Send[GLKEffectPropertyFog](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (gc _GLKEffectPropertyFogClass) New() GLKEffectPropertyFog {
	rv := objc.Send[GLKEffectPropertyFog](objc.ID(gc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (g_ GLKEffectPropertyFog) Init() GLKEffectPropertyFog {
	rv := objc.Send[GLKEffectPropertyFog](g_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g_ GLKEffectPropertyFog) Autorelease() GLKEffectPropertyFog {
	rv := objc.Send[GLKEffectPropertyFog](g_.ID, objc.Sel("autorelease"))
	return rv
}

// NewGLKEffectPropertyFog creates a new GLKEffectPropertyFog instance.
func NewGLKEffectPropertyFog() GLKEffectPropertyFog {
	return getGLKEffectPropertyFogClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for GLKEffectPropertyFog */
// Fog drawing information for use in GLKit rendering effects.
//
// These properties are specifically designed to mimic the fog calculations provided by OpenGL ES 1.1. When fog is enabled, the fog component is calculated and clamped to a range from to . Then, the fog value is used as a blending factor between the computed fragment color and the fog color.


// Fog drawing information for use in GLKit rendering effects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKEffectPropertyFog
type GLKEffectPropertyFog struct {
	GLKEffectProperty
}

// GLKEffectPropertyFogFrom constructs a [GLKEffectPropertyFog] from an unsafe.Pointer.
//
// Fog drawing information for use in GLKit rendering effects.
func GLKEffectPropertyFogFrom(ptr unsafe.Pointer) GLKEffectPropertyFog {
	return GLKEffectPropertyFog{
		GLKEffectProperty: GLKEffectPropertyFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for GLKEffectPropertyFog *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for GLKEffectPropertyFog */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for GLKEffectPropertyFog */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for GLKEffectPropertyFog */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for GLKEffectPropertyFog */

// The color of the fog at maximum density.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKEffectPropertyFog/color
func (g_ GLKEffectPropertyFog) Color() GLKVector4 /* typedef */ {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("color"))
	return rv
}/* debug [instance_properties/getter]: color */


// The color of the fog at maximum density.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKEffectPropertyFog/color
func (g_ GLKEffectPropertyFog) SetColor(value GLKVector4 /* typedef */) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setColor:"), value)
}/* debug [instance_properties/setter]: color */


// The rate at which the fog exponent increases.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKEffectPropertyFog/density
func (g_ GLKEffectPropertyFog) Density() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("density"))
	return rv
}/* debug [instance_properties/getter]: density */


// The rate at which the fog exponent increases.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKEffectPropertyFog/density
func (g_ GLKEffectPropertyFog) SetDensity(value unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setDensity:"), value)
}/* debug [instance_properties/setter]: density */


// A Boolean value that indicates whether fog is applied to the fragment color.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKEffectPropertyFog/enabled
func (g_ GLKEffectPropertyFog) Enabled() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("enabled"))
	return rv
}/* debug [instance_properties/getter]: enabled */


// A Boolean value that indicates whether fog is applied to the fragment color.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKEffectPropertyFog/enabled
func (g_ GLKEffectPropertyFog) SetEnabled(value unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setEnabled:"), value)
}/* debug [instance_properties/setter]: enabled */


// The distance in eye coordinates where fog completely covers the color fragment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKEffectPropertyFog/end
func (g_ GLKEffectPropertyFog) End() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("end"))
	return rv
}/* debug [instance_properties/getter]: end */


// The distance in eye coordinates where fog completely covers the color fragment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKEffectPropertyFog/end
func (g_ GLKEffectPropertyFog) SetEnd(value unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setEnd:"), value)
}/* debug [instance_properties/setter]: end */


// The algorithm used to compute the density of the fog applied to the fragment color.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKEffectPropertyFog/mode
func (g_ GLKEffectPropertyFog) Mode() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("mode"))
	return rv
}/* debug [instance_properties/getter]: mode */


// The algorithm used to compute the density of the fog applied to the fragment color.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKEffectPropertyFog/mode
func (g_ GLKEffectPropertyFog) SetMode(value unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setMode:"), value)
}/* debug [instance_properties/setter]: mode */


// The minimum distance in eye coordinates before fog is applied to the fragment color.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKEffectPropertyFog/start
func (g_ GLKEffectPropertyFog) Start() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("start"))
	return rv
}/* debug [instance_properties/getter]: start */


// The minimum distance in eye coordinates before fog is applied to the fragment color.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKEffectPropertyFog/start
func (g_ GLKEffectPropertyFog) SetStart(value unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setStart:"), value)
}/* debug [instance_properties/setter]: start */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class GLKEffectPropertyFog */



