// Code generated from Apple documentation for GLKit. DO NOT EDIT.

package glkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [GLKReflectionMapEffect] class.
var (
	GLKReflectionMapEffectClass     _GLKReflectionMapEffectClass
	GLKReflectionMapEffectClassOnce sync.Once
)

func getGLKReflectionMapEffectClass() _GLKReflectionMapEffectClass {
	GLKReflectionMapEffectClassOnce.Do(func() {
		GLKReflectionMapEffectClass = _GLKReflectionMapEffectClass{objc.GetClass("GLKReflectionMapEffect")}
	})
	return GLKReflectionMapEffectClass
}

type _GLKReflectionMapEffectClass struct {
	class objc.Class
}

// An interface definition for the [GLKReflectionMapEffect] class.
type IGLKReflectionMapEffect interface {
	IGLKBaseEffect
	PrepareToDraw()
	Matrix() GLKMatrix3
	SetMatrix(value IGLKMatrix3)
	TextureCubeMap() GLKEffectPropertyTexture
	TextureOrder() GLKEffectPropertyTexture
	SetTextureOrder(value IGLKEffectPropertyTexture)
}

// A lighting and shading system that supports reflection mapping for use in shader-based OpenGL rendering.
//
// In addition to any of the properties provided by the class, your application must also configure the properties on the reflection map. The default value of the property provided by the base effect is modified to include the reflection map as a final texturing stage; your application can modify the value of that property to change the order in which texturing occurs. The reflection map effect is calculated in accordance to section 2.11.4 of the OpenGL 2.1 specification glTexGen() mode. It requires a cube map texture to define the enclosing envelope from which to reflection map the scene.


// A lighting and shading system that supports reflection mapping for use in shader-based OpenGL rendering.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKReflectionMapEffect

type GLKReflectionMapEffect struct {
	GLKBaseEffect
}

// GLKReflectionMapEffectFrom constructs a [GLKReflectionMapEffect] from an unsafe.Pointer.
//
// A lighting and shading system that supports reflection mapping for use in shader-based OpenGL rendering.
func GLKReflectionMapEffectFrom(ptr unsafe.Pointer) GLKReflectionMapEffect {
	return GLKReflectionMapEffect{
		GLKBaseEffect: GLKBaseEffectFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (gc _GLKReflectionMapEffectClass) Alloc() GLKReflectionMapEffect {
	rv := objc.Send[GLKReflectionMapEffect](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (gc _GLKReflectionMapEffectClass) New() GLKReflectionMapEffect {
	rv := objc.Send[GLKReflectionMapEffect](objc.ID(gc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (g_ GLKReflectionMapEffect) Init() GLKReflectionMapEffect {
	rv := objc.Send[GLKReflectionMapEffect](g_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g_ GLKReflectionMapEffect) Autorelease() GLKReflectionMapEffect {
	rv := objc.Send[GLKReflectionMapEffect](g_.ID, objc.Sel("autorelease"))
	return rv
}

// NewGLKReflectionMapEffect creates a new GLKReflectionMapEffect instance.
func NewGLKReflectionMapEffect() GLKReflectionMapEffect {
	return getGLKReflectionMapEffectClass().New()
}




// Prepares an effect for rendering.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKReflectionMapEffect/prepareToDraw()

func (g_ GLKReflectionMapEffect) PrepareToDraw() {
	objc.Send[objc.ID](g_.ID, objc.Sel("prepareToDraw"))
}


// The reflection matrix to apply to the normals of the submitted vertices.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKReflectionMapEffect/matrix

func (g_ GLKReflectionMapEffect) Matrix() GLKMatrix3 {
	rv := objc.Send[GLKMatrix3](g_.ID, objc.Sel("matrix"))
	return rv
}


// The reflection matrix to apply to the normals of the submitted vertices.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKReflectionMapEffect/matrix

func (g_ GLKReflectionMapEffect) SetMatrix(value IGLKMatrix3) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setMatrix:"), value)
}


// The texture map to apply in the reflection stage.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKReflectionMapEffect/textureCubeMap

func (g_ GLKReflectionMapEffect) TextureCubeMap() GLKEffectPropertyTexture {
	rv := objc.Send[GLKEffectPropertyTexture](g_.ID, objc.Sel("textureCubeMap"))
	return rv
}


// The order in which textures are applied to rendered primitives.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/glkit/glkbaseeffect/textureorder

func (g_ GLKReflectionMapEffect) TextureOrder() GLKEffectPropertyTexture {
	rv := objc.Send[GLKEffectPropertyTexture](g_.ID, objc.Sel("textureOrder"))
	return rv
}


// The order in which textures are applied to rendered primitives.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/glkit/glkbaseeffect/textureorder

func (g_ GLKReflectionMapEffect) SetTextureOrder(value IGLKEffectPropertyTexture) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setTextureOrder:"), value)
}



