// Code generated from Apple documentation for GLKit. DO NOT EDIT.

package glkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [GLKEffectPropertyTexture] class.
var (
	GLKEffectPropertyTextureClass     _GLKEffectPropertyTextureClass
	GLKEffectPropertyTextureClassOnce sync.Once
)

func getGLKEffectPropertyTextureClass() _GLKEffectPropertyTextureClass {
	GLKEffectPropertyTextureClassOnce.Do(func() {
		GLKEffectPropertyTextureClass = _GLKEffectPropertyTextureClass{objc.GetClass("GLKEffectPropertyTexture")}
	})
	return GLKEffectPropertyTextureClass
}

type _GLKEffectPropertyTextureClass struct {
	class objc.Class
}

// An interface definition for the [GLKEffectPropertyTexture] class.
type IGLKEffectPropertyTexture interface {
	IGLKEffectProperty
	Enabled() unsafe.Pointer
	SetEnabled(value unsafe.Pointer)
	EnvMode() GLKTextureEnvMode
	SetEnvMode(value GLKTextureEnvMode)
	Name() unsafe.Pointer
	SetName(value unsafe.Pointer)
	Target() GLKTextureTarget
	SetTarget(value IGLKTextureTarget)
}

// Texture drawing parameters for use in GLKit rendering effects.
//
// The class defines properties that are used to configure an OpenGL texturing operation. The texturing operation combines an input color and a color sampled from the texture and outputs a new color to the next stage of calculations. The property determines the function used to calculate the output color from the two input colors. If an effect only includes a single texture property, then the input color is the lighting color calculated by the lighting stage of the graphics pipeline. An effect can also include multiple objects. When an effect includes multiple properties, the first texture stage uses the lighting color as the first input color. Each texture stage after that uses the output of the previous stage as the input color.
//
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKEffectPropertyTexture
type GLKEffectPropertyTexture struct {
	GLKEffectProperty
}

// GLKEffectPropertyTextureFrom constructs a [GLKEffectPropertyTexture] from an unsafe.Pointer.
//
// Texture drawing parameters for use in GLKit rendering effects.
func GLKEffectPropertyTextureFrom(ptr unsafe.Pointer) GLKEffectPropertyTexture {
	return GLKEffectPropertyTexture{
		GLKEffectProperty: GLKEffectPropertyFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (gc _GLKEffectPropertyTextureClass) Alloc() GLKEffectPropertyTexture {
	rv := objc.Send[GLKEffectPropertyTexture](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (gc _GLKEffectPropertyTextureClass) New() GLKEffectPropertyTexture {
	rv := objc.Send[GLKEffectPropertyTexture](objc.ID(gc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (g_ GLKEffectPropertyTexture) Init() GLKEffectPropertyTexture {
	rv := objc.Send[GLKEffectPropertyTexture](g_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g_ GLKEffectPropertyTexture) Autorelease() GLKEffectPropertyTexture {
	rv := objc.Send[GLKEffectPropertyTexture](g_.ID, objc.Sel("autorelease"))
	return rv
}

// NewGLKEffectPropertyTexture creates a new GLKEffectPropertyTexture instance.
func NewGLKEffectPropertyTexture() GLKEffectPropertyTexture {
	return getGLKEffectPropertyTextureClass().New()
}


// A Boolean value that indicates whether this texture is used to texture drawn primitives.
//
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKEffectPropertyTexture/enabled
func (g_ GLKEffectPropertyTexture) Enabled() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("enabled"))
	return rv
}


// SetEnabled sets the value of the enabled property.
// A Boolean value that indicates whether this texture is used to texture drawn primitives.

//
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKEffectPropertyTexture/enabled
func (g_ GLKEffectPropertyTexture) SetEnabled(value unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setEnabled:"), value)
}

// The mode the texture uses to compute its output fragment color. See .
//
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKEffectPropertyTexture/envMode
func (g_ GLKEffectPropertyTexture) EnvMode() GLKTextureEnvMode {
	rv := objc.Send[GLKTextureEnvMode](g_.ID, objc.Sel("envMode"))
	return rv
}


// SetEnvMode sets the value of the envMode property.
// The mode the texture uses to compute its output fragment color. See .

//
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKEffectPropertyTexture/envMode
func (g_ GLKEffectPropertyTexture) SetEnvMode(value GLKTextureEnvMode) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setEnvMode:"), value)
}

// The OpenGL name for the texture being sampled by this texture stage.
//
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKEffectPropertyTexture/name
func (g_ GLKEffectPropertyTexture) Name() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("name"))
	return rv
}


// SetName sets the value of the name property.
// The OpenGL name for the texture being sampled by this texture stage.

//
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKEffectPropertyTexture/name
func (g_ GLKEffectPropertyTexture) SetName(value unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setName:"), value)
}

// The kind of texture pointed to by the texture stage. See .
//
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKEffectPropertyTexture/target
func (g_ GLKEffectPropertyTexture) Target() GLKTextureTarget {
	rv := objc.Send[GLKTextureTarget](g_.ID, objc.Sel("target"))
	return rv
}


// SetTarget sets the value of the target property.
// The kind of texture pointed to by the texture stage. See .

//
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKEffectPropertyTexture/target
func (g_ GLKEffectPropertyTexture) SetTarget(value IGLKTextureTarget) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setTarget:"), value)
}



