// Code generated from Apple documentation for GLKit. DO NOT EDIT.

package glkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class GLKReflectionMapEffect */


/* debug [class_header]: Header for GLKReflectionMapEffect */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for GLKReflectionMapEffect */
// An interface definition for the [GLKReflectionMapEffect] class.
type IGLKReflectionMapEffect interface {
	IGLKBaseEffect
	
/* debug [class_interface_properties]: Properties for GLKReflectionMapEffect */
	// properties:
	Matrix() GLKMatrix3 /* typedef */
	SetMatrix(value GLKMatrix3 /* typedef */)
	TextureCubeMap() IGLKEffectPropertyTexture
	TextureOrder() IGLKEffectPropertyTexture
	SetTextureOrder(value IGLKEffectPropertyTexture)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for GLKReflectionMapEffect */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for GLKReflectionMapEffect */
// Alloc allocates a new instance without initialization.
func (gc _GLKReflectionMapEffectClass) Alloc() GLKReflectionMapEffect {
	rv := objc.Send[GLKReflectionMapEffect](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for GLKReflectionMapEffect */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for GLKReflectionMapEffect *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for GLKReflectionMapEffect */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for GLKReflectionMapEffect */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for GLKReflectionMapEffect */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for GLKReflectionMapEffect */

// The reflection matrix to apply to the normals of the submitted vertices.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKReflectionMapEffect/matrix
func (g_ GLKReflectionMapEffect) Matrix() GLKMatrix3 /* typedef */ {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("matrix"))
	return rv
}/* debug [instance_properties/getter]: matrix */


// The reflection matrix to apply to the normals of the submitted vertices.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKReflectionMapEffect/matrix
func (g_ GLKReflectionMapEffect) SetMatrix(value GLKMatrix3 /* typedef */) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setMatrix:"), value)
}/* debug [instance_properties/setter]: matrix */


// The texture map to apply in the reflection stage.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKReflectionMapEffect/textureCubeMap
func (g_ GLKReflectionMapEffect) TextureCubeMap() IGLKEffectPropertyTexture {
	rv := objc.Send[GLKEffectPropertyTexture](g_.ID, objc.Sel("textureCubeMap"))
	return rv
}/* debug [instance_properties/getter]: textureCubeMap */


// The order in which textures are applied to rendered primitives.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/glkit/glkbaseeffect/textureorder
func (g_ GLKReflectionMapEffect) TextureOrder() IGLKEffectPropertyTexture {
	rv := objc.Send[GLKEffectPropertyTexture](g_.ID, objc.Sel("textureOrder"))
	return rv
}/* debug [instance_properties/getter]: textureOrder */


// The order in which textures are applied to rendered primitives.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/glkit/glkbaseeffect/textureorder
func (g_ GLKReflectionMapEffect) SetTextureOrder(value IGLKEffectPropertyTexture) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setTextureOrder:"), value)
}/* debug [instance_properties/setter]: textureOrder */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class GLKReflectionMapEffect */



