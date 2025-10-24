// Code generated from Apple documentation for GLKit. DO NOT EDIT.

package glkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class GLKEffectPropertyMaterial */


/* debug [class_header]: Header for GLKEffectPropertyMaterial */
// The class instance for the [GLKEffectPropertyMaterial] class.
var (
	GLKEffectPropertyMaterialClass     _GLKEffectPropertyMaterialClass
	GLKEffectPropertyMaterialClassOnce sync.Once
)

func getGLKEffectPropertyMaterialClass() _GLKEffectPropertyMaterialClass {
	GLKEffectPropertyMaterialClassOnce.Do(func() {
		GLKEffectPropertyMaterialClass = _GLKEffectPropertyMaterialClass{objc.GetClass("GLKEffectPropertyMaterial")}
	})
	return GLKEffectPropertyMaterialClass
}

type _GLKEffectPropertyMaterialClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for GLKEffectPropertyMaterial */
// An interface definition for the [GLKEffectPropertyMaterial] class.
type IGLKEffectPropertyMaterial interface {
	IGLKEffectProperty
	
/* debug [class_interface_properties]: Properties for GLKEffectPropertyMaterial */
	// properties:
	AmbientColor() GLKVector4 /* typedef */
	SetAmbientColor(value GLKVector4 /* typedef */)
	DiffuseColor() GLKVector4 /* typedef */
	SetDiffuseColor(value GLKVector4 /* typedef */)
	EmissiveColor() GLKVector4 /* typedef */
	SetEmissiveColor(value GLKVector4 /* typedef */)
	Shininess() unsafe.Pointer
	SetShininess(value unsafe.Pointer)
	SpecularColor() GLKVector4 /* typedef */
	SetSpecularColor(value GLKVector4 /* typedef */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for GLKEffectPropertyMaterial */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for GLKEffectPropertyMaterial */
// Alloc allocates a new instance without initialization.
func (gc _GLKEffectPropertyMaterialClass) Alloc() GLKEffectPropertyMaterial {
	rv := objc.Send[GLKEffectPropertyMaterial](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (gc _GLKEffectPropertyMaterialClass) New() GLKEffectPropertyMaterial {
	rv := objc.Send[GLKEffectPropertyMaterial](objc.ID(gc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (g_ GLKEffectPropertyMaterial) Init() GLKEffectPropertyMaterial {
	rv := objc.Send[GLKEffectPropertyMaterial](g_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g_ GLKEffectPropertyMaterial) Autorelease() GLKEffectPropertyMaterial {
	rv := objc.Send[GLKEffectPropertyMaterial](g_.ID, objc.Sel("autorelease"))
	return rv
}

// NewGLKEffectPropertyMaterial creates a new GLKEffectPropertyMaterial instance.
func NewGLKEffectPropertyMaterial() GLKEffectPropertyMaterial {
	return getGLKEffectPropertyMaterialClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for GLKEffectPropertyMaterial */
// Surface appearance properties for use in GLKit rendering effects.
//
// The class defines properties used to configure the characteristics of the surface being lit. The material properties for an effect interact with light properties on the same effect to determine how that surface is lit within the scene. The behavior of this class matches the material properties and lighting calculations defined in the OpenGL ES 1.1 specification.


// Surface appearance properties for use in GLKit rendering effects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKEffectPropertyMaterial
type GLKEffectPropertyMaterial struct {
	GLKEffectProperty
}

// GLKEffectPropertyMaterialFrom constructs a [GLKEffectPropertyMaterial] from an unsafe.Pointer.
//
// Surface appearance properties for use in GLKit rendering effects.
func GLKEffectPropertyMaterialFrom(ptr unsafe.Pointer) GLKEffectPropertyMaterial {
	return GLKEffectPropertyMaterial{
		GLKEffectProperty: GLKEffectPropertyFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for GLKEffectPropertyMaterial *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for GLKEffectPropertyMaterial */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for GLKEffectPropertyMaterial */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for GLKEffectPropertyMaterial */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for GLKEffectPropertyMaterial */

// The ambient color of the material.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKEffectPropertyMaterial/ambientColor
func (g_ GLKEffectPropertyMaterial) AmbientColor() GLKVector4 /* typedef */ {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("ambientColor"))
	return rv
}/* debug [instance_properties/getter]: ambientColor */


// The ambient color of the material.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKEffectPropertyMaterial/ambientColor
func (g_ GLKEffectPropertyMaterial) SetAmbientColor(value GLKVector4 /* typedef */) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setAmbientColor:"), value)
}/* debug [instance_properties/setter]: ambientColor */


// The diffuse color of the material.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKEffectPropertyMaterial/diffuseColor
func (g_ GLKEffectPropertyMaterial) DiffuseColor() GLKVector4 /* typedef */ {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("diffuseColor"))
	return rv
}/* debug [instance_properties/getter]: diffuseColor */


// The diffuse color of the material.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKEffectPropertyMaterial/diffuseColor
func (g_ GLKEffectPropertyMaterial) SetDiffuseColor(value GLKVector4 /* typedef */) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setDiffuseColor:"), value)
}/* debug [instance_properties/setter]: diffuseColor */


// The emissive color of the material.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKEffectPropertyMaterial/emissiveColor
func (g_ GLKEffectPropertyMaterial) EmissiveColor() GLKVector4 /* typedef */ {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("emissiveColor"))
	return rv
}/* debug [instance_properties/getter]: emissiveColor */


// The emissive color of the material.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKEffectPropertyMaterial/emissiveColor
func (g_ GLKEffectPropertyMaterial) SetEmissiveColor(value GLKVector4 /* typedef */) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setEmissiveColor:"), value)
}/* debug [instance_properties/setter]: emissiveColor */


// The shininess of the material, used when calculating specular lighting effects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKEffectPropertyMaterial/shininess
func (g_ GLKEffectPropertyMaterial) Shininess() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("shininess"))
	return rv
}/* debug [instance_properties/getter]: shininess */


// The shininess of the material, used when calculating specular lighting effects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKEffectPropertyMaterial/shininess
func (g_ GLKEffectPropertyMaterial) SetShininess(value unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setShininess:"), value)
}/* debug [instance_properties/setter]: shininess */


// The specular color of the material.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKEffectPropertyMaterial/specularColor
func (g_ GLKEffectPropertyMaterial) SpecularColor() GLKVector4 /* typedef */ {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("specularColor"))
	return rv
}/* debug [instance_properties/getter]: specularColor */


// The specular color of the material.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKEffectPropertyMaterial/specularColor
func (g_ GLKEffectPropertyMaterial) SetSpecularColor(value GLKVector4 /* typedef */) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setSpecularColor:"), value)
}/* debug [instance_properties/setter]: specularColor */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class GLKEffectPropertyMaterial */



