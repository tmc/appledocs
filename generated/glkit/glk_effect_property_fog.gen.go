// Code generated from Apple documentation for GLKit. DO NOT EDIT.

package glkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

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

// An interface definition for the [GLKEffectPropertyFog] class.
type IGLKEffectPropertyFog interface {
	IGLKEffectProperty
}

// Fog drawing information for use in GLKit rendering effects.
//
// These properties are specifically designed to mimic the fog calculations provided by OpenGL ES 1.1. When fog is enabled, the fog component is calculated and clamped to a range from to . Then, the fog value is used as a blending factor between the computed fragment color and the fog color.
//
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

// Alloc allocates a new instance without initialization.
func (gc _GLKEffectPropertyFogClass) Alloc() GLKEffectPropertyFog {
	rv := objc.Send[GLKEffectPropertyFog](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


// The color of the fog at maximum density.
//
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKEffectPropertyFog/color
func (g_ GLKEffectPropertyFog) Color() GLKVector4 {
	rv := objc.Send[GLKVector4](g_.ID, objc.Sel("color"))
	return rv
}


// SetColor sets the value of the color property.
// The color of the fog at maximum density.

//
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKEffectPropertyFog/color
func (g_ GLKEffectPropertyFog) SetColor(value IGLKVector4) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setColor:"), value)
}

// The rate at which the fog exponent increases.
//
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKEffectPropertyFog/density
func (g_ GLKEffectPropertyFog) Density() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("density"))
	return rv
}


// SetDensity sets the value of the density property.
// The rate at which the fog exponent increases.

//
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKEffectPropertyFog/density
func (g_ GLKEffectPropertyFog) SetDensity(value unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setDensity:"), value)
}

// A Boolean value that indicates whether fog is applied to the fragment color.
//
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKEffectPropertyFog/enabled
func (g_ GLKEffectPropertyFog) Enabled() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("enabled"))
	return rv
}


// SetEnabled sets the value of the enabled property.
// A Boolean value that indicates whether fog is applied to the fragment color.

//
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKEffectPropertyFog/enabled
func (g_ GLKEffectPropertyFog) SetEnabled(value unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setEnabled:"), value)
}

// The distance in eye coordinates where fog completely covers the color fragment.
//
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKEffectPropertyFog/end
func (g_ GLKEffectPropertyFog) End() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("end"))
	return rv
}


// SetEnd sets the value of the end property.
// The distance in eye coordinates where fog completely covers the color fragment.

//
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKEffectPropertyFog/end
func (g_ GLKEffectPropertyFog) SetEnd(value unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setEnd:"), value)
}

// The algorithm used to compute the density of the fog applied to the fragment color.
//
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKEffectPropertyFog/mode
func (g_ GLKEffectPropertyFog) Mode() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("mode"))
	return rv
}


// SetMode sets the value of the mode property.
// The algorithm used to compute the density of the fog applied to the fragment color.

//
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKEffectPropertyFog/mode
func (g_ GLKEffectPropertyFog) SetMode(value unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setMode:"), value)
}

// The minimum distance in eye coordinates before fog is applied to the fragment color.
//
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKEffectPropertyFog/start
func (g_ GLKEffectPropertyFog) Start() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("start"))
	return rv
}


// SetStart sets the value of the start property.
// The minimum distance in eye coordinates before fog is applied to the fragment color.

//
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKEffectPropertyFog/start
func (g_ GLKEffectPropertyFog) SetStart(value unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setStart:"), value)
}



