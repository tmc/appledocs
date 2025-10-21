// Code generated from Apple documentation for GLKit. DO NOT EDIT.

package glkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

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

// An interface definition for the [GLKEffectPropertyMaterial] class.
type IGLKEffectPropertyMaterial interface {
	IGLKEffectProperty
}

// Surface appearance properties for use in GLKit rendering effects.
//
// The class defines properties used to configure the characteristics of the surface being lit. The material properties for an effect interact with light properties on the same effect to determine how that surface is lit within the scene. The behavior of this class matches the material properties and lighting calculations defined in the OpenGL ES 1.1 specification.
//
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

// Alloc allocates a new instance without initialization.
func (gc _GLKEffectPropertyMaterialClass) Alloc() GLKEffectPropertyMaterial {
	rv := objc.Send[GLKEffectPropertyMaterial](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


// The ambient color of the material.
//
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKEffectPropertyMaterial/ambientColor
func (g_ GLKEffectPropertyMaterial) AmbientColor() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("ambientColor"))
	return rv
}


// SetAmbientColor sets the value of the ambientColor property.
// The ambient color of the material.

//
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKEffectPropertyMaterial/ambientColor
func (g_ GLKEffectPropertyMaterial) SetAmbientColor(value unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setAmbientColor:"), value)
}
// The diffuse color of the material.
//
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKEffectPropertyMaterial/diffuseColor
func (g_ GLKEffectPropertyMaterial) DiffuseColor() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("diffuseColor"))
	return rv
}


// SetDiffuseColor sets the value of the diffuseColor property.
// The diffuse color of the material.

//
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKEffectPropertyMaterial/diffuseColor
func (g_ GLKEffectPropertyMaterial) SetDiffuseColor(value unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setDiffuseColor:"), value)
}
// The emissive color of the material.
//
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKEffectPropertyMaterial/emissiveColor
func (g_ GLKEffectPropertyMaterial) EmissiveColor() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("emissiveColor"))
	return rv
}


// SetEmissiveColor sets the value of the emissiveColor property.
// The emissive color of the material.

//
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKEffectPropertyMaterial/emissiveColor
func (g_ GLKEffectPropertyMaterial) SetEmissiveColor(value unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setEmissiveColor:"), value)
}
// The shininess of the material, used when calculating specular lighting effects.
//
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKEffectPropertyMaterial/shininess
func (g_ GLKEffectPropertyMaterial) Shininess() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("shininess"))
	return rv
}


// SetShininess sets the value of the shininess property.
// The shininess of the material, used when calculating specular lighting effects.

//
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKEffectPropertyMaterial/shininess
func (g_ GLKEffectPropertyMaterial) SetShininess(value unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setShininess:"), value)
}
// The specular color of the material.
//
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKEffectPropertyMaterial/specularColor
func (g_ GLKEffectPropertyMaterial) SpecularColor() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("specularColor"))
	return rv
}


// SetSpecularColor sets the value of the specularColor property.
// The specular color of the material.

//
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKEffectPropertyMaterial/specularColor
func (g_ GLKEffectPropertyMaterial) SetSpecularColor(value unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setSpecularColor:"), value)
}


