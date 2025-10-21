// Code generated from Apple documentation for GLKit. DO NOT EDIT.

package glkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [GLKBaseEffect] class.
var (
	GLKBaseEffectClass     _GLKBaseEffectClass
	GLKBaseEffectClassOnce sync.Once
)

func getGLKBaseEffectClass() _GLKBaseEffectClass {
	GLKBaseEffectClassOnce.Do(func() {
		GLKBaseEffectClass = _GLKBaseEffectClass{objc.GetClass("GLKBaseEffect")}
	})
	return GLKBaseEffectClass
}

type _GLKBaseEffectClass struct {
	class objc.Class
}

// An interface definition for the [GLKBaseEffect] class.
type IGLKBaseEffect interface {
	objectivec.IObject
	PrepareToDraw()
}

// A simple lighting and shading system for use in shader-based OpenGL rendering.
//
// The class provides shaders that mimic many of the behaviors provided by the OpenGL ES 1.1 lighting and shading model, including materials, lighting and texturing. The base effect allows up to three lights and two textures to be applied to a scene. At initialization time, your application first creates a compatible OpenGL or OpenGL ES context and makes it current. Then, it allocates and initializes a new effect object, configures its properties, and calls its method. Binding an effect causes a shader to be compiled and bound to the current context. The base effect also requires vertex data to be supplied by your application. To supply vertex data, create one or more vertex array objects. For each attribute required by the shader, the vertex array object should enable the attribute and point to data stored in a vertex buffer object. At rendering time, your application calls the effect’s method to prepare the effect. Then, it binds a vertex array object and submits one or more drawing commands. Lighting calculations for the base effect are done in eye-space coordinates.  The , and properties hold the position and spot direction of the base effect’s lights. The property contains the model view matrix assigned to the scene. When a light is assigned a new position or spot direction, those values are immediately modified by the current model view matrix. Thus, it is important to sequence changes to the model view matrix and changes to the lights to achieve the desired light positioning.  Light positions that need to be transformed in a manner similar to scene geometry should be set after the model view matrix is updated.
//
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKBaseEffect
type GLKBaseEffect struct {
	objectivec.Object
}

// GLKBaseEffectFrom constructs a [GLKBaseEffect] from an unsafe.Pointer.
//
// A simple lighting and shading system for use in shader-based OpenGL rendering.
func GLKBaseEffectFrom(ptr unsafe.Pointer) GLKBaseEffect {
	return GLKBaseEffect{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (gc _GLKBaseEffectClass) Alloc() GLKBaseEffect {
	rv := objc.Send[GLKBaseEffect](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (gc _GLKBaseEffectClass) New() GLKBaseEffect {
	rv := objc.Send[GLKBaseEffect](objc.ID(gc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (g_ GLKBaseEffect) Init() GLKBaseEffect {
	rv := objc.Send[GLKBaseEffect](g_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g_ GLKBaseEffect) Autorelease() GLKBaseEffect {
	rv := objc.Send[GLKBaseEffect](g_.ID, objc.Sel("autorelease"))
	return rv
}

// NewGLKBaseEffect creates a new GLKBaseEffect instance.
func NewGLKBaseEffect() GLKBaseEffect {
	return getGLKBaseEffectClass().New()
}


// Prepares an effect for rendering.
//
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKBaseEffect/prepareToDraw()
func (g_ GLKBaseEffect) PrepareToDraw() {
	objc.Send[objc.ID](g_.ID, objc.Sel("prepareToDraw"))
}

// A Boolean value that indicates whether or not to use the color vertex attribute when calculating the light’s interaction with the material.
//
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKBaseEffect/colorMaterialEnabled
func (g_ GLKBaseEffect) ColorMaterialEnabled() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("colorMaterialEnabled"))
	return rv
}


// SetColorMaterialEnabled sets the value of the colorMaterialEnabled property.
// A Boolean value that indicates whether or not to use the color vertex attribute when calculating the light’s interaction with the material.

//
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKBaseEffect/colorMaterialEnabled
func (g_ GLKBaseEffect) SetColorMaterialEnabled(value unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setColorMaterialEnabled:"), value)
}

// A constant color, used when per-vertex color data is not provided.
//
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKBaseEffect/constantColor
func (g_ GLKBaseEffect) ConstantColor() GLKVector4 {
	rv := objc.Send[GLKVector4](g_.ID, objc.Sel("constantColor"))
	return rv
}


// SetConstantColor sets the value of the constantColor property.
// A constant color, used when per-vertex color data is not provided.

//
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKBaseEffect/constantColor
func (g_ GLKBaseEffect) SetConstantColor(value IGLKVector4) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setConstantColor:"), value)
}

// The fog properties to apply to the scene.
//
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKBaseEffect/fog
func (g_ GLKBaseEffect) Fog() GLKEffectPropertyFog {
	rv := objc.Send[GLKEffectPropertyFog](g_.ID, objc.Sel("fog"))
	return rv
}

// A string used to name your effect.
//
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKBaseEffect/label
func (g_ GLKBaseEffect) Label() appkit.string {
	rv := objc.Send[appkit.string](g_.ID, objc.Sel("label"))
	return rv
}


// SetLabel sets the value of the label property.
// A string used to name your effect.

//
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKBaseEffect/label
func (g_ GLKBaseEffect) SetLabel(value appkit.string) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setLabel:"), value)
}

// The lighting properties for the first light in the scene.
//
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKBaseEffect/light0
func (g_ GLKBaseEffect) Light0() GLKEffectPropertyLight {
	rv := objc.Send[GLKEffectPropertyLight](g_.ID, objc.Sel("light0"))
	return rv
}

// The lighting properties for the second light in the scene.
//
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKBaseEffect/light1
func (g_ GLKBaseEffect) Light1() GLKEffectPropertyLight {
	rv := objc.Send[GLKEffectPropertyLight](g_.ID, objc.Sel("light1"))
	return rv
}

// The lighting properties for the third light in the scene.
//
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKBaseEffect/light2
func (g_ GLKBaseEffect) Light2() GLKEffectPropertyLight {
	rv := objc.Send[GLKEffectPropertyLight](g_.ID, objc.Sel("light2"))
	return rv
}

// The ambient color applied to all primitives rendered by the effect.
//
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKBaseEffect/lightModelAmbientColor
func (g_ GLKBaseEffect) LightModelAmbientColor() GLKVector4 {
	rv := objc.Send[GLKVector4](g_.ID, objc.Sel("lightModelAmbientColor"))
	return rv
}


// SetLightModelAmbientColor sets the value of the lightModelAmbientColor property.
// The ambient color applied to all primitives rendered by the effect.

//
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKBaseEffect/lightModelAmbientColor
func (g_ GLKBaseEffect) SetLightModelAmbientColor(value IGLKVector4) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setLightModelAmbientColor:"), value)
}

// A Boolean value that indicates whether lighting is calculated for both sides of a primitive.
//
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKBaseEffect/lightModelTwoSided
func (g_ GLKBaseEffect) LightModelTwoSided() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("lightModelTwoSided"))
	return rv
}


// SetLightModelTwoSided sets the value of the lightModelTwoSided property.
// A Boolean value that indicates whether lighting is calculated for both sides of a primitive.

//
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKBaseEffect/lightModelTwoSided
func (g_ GLKBaseEffect) SetLightModelTwoSided(value unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setLightModelTwoSided:"), value)
}

// The strategy the effect uses to calculate light values at each fragment. See .
//
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKBaseEffect/lightingType
func (g_ GLKBaseEffect) LightingType() GLKLightingType {
	rv := objc.Send[GLKLightingType](g_.ID, objc.Sel("lightingType"))
	return rv
}


// SetLightingType sets the value of the lightingType property.
// The strategy the effect uses to calculate light values at each fragment. See .

//
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKBaseEffect/lightingType
func (g_ GLKBaseEffect) SetLightingType(value GLKLightingType) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setLightingType:"), value)
}

// The material properties used when calculating the light values for a rendered primitive.
//
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKBaseEffect/material
func (g_ GLKBaseEffect) Material() GLKEffectPropertyMaterial {
	rv := objc.Send[GLKEffectPropertyMaterial](g_.ID, objc.Sel("material"))
	return rv
}

// The properties for the first texture.
//
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKBaseEffect/texture2d0
func (g_ GLKBaseEffect) Texture2d0() GLKEffectPropertyTexture {
	rv := objc.Send[GLKEffectPropertyTexture](g_.ID, objc.Sel("texture2d0"))
	return rv
}

// The properties for the second texture.
//
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKBaseEffect/texture2d1
func (g_ GLKBaseEffect) Texture2d1() GLKEffectPropertyTexture {
	rv := objc.Send[GLKEffectPropertyTexture](g_.ID, objc.Sel("texture2d1"))
	return rv
}

// The order in which textures are applied to rendered primitives.
//
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKBaseEffect/textureOrder
func (g_ GLKBaseEffect) TextureOrder() []GLKEffectPropertyTexture {
	rv := objc.Send[[]GLKEffectPropertyTexture](g_.ID, objc.Sel("textureOrder"))
	return rv
}


// SetTextureOrder sets the value of the textureOrder property.
// The order in which textures are applied to rendered primitives.

//
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKBaseEffect/textureOrder
func (g_ GLKBaseEffect) SetTextureOrder(value []GLKEffectPropertyTexture) {
	// Convert Go slice to NSArray
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](g_.ID, objc.Sel("setTextureOrder:"), nsArray)
}

// The modelview, projection and texture transformations applied to the vertex data when the effect is bound.
//
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKBaseEffect/transform
func (g_ GLKBaseEffect) Transform() GLKEffectPropertyTransform {
	rv := objc.Send[GLKEffectPropertyTransform](g_.ID, objc.Sel("transform"))
	return rv
}

// A Boolean value that indicates whether or not to use the constant color.
//
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKBaseEffect/useConstantColor
func (g_ GLKBaseEffect) UseConstantColor() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("useConstantColor"))
	return rv
}


// SetUseConstantColor sets the value of the useConstantColor property.
// A Boolean value that indicates whether or not to use the constant color.

//
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKBaseEffect/useConstantColor
func (g_ GLKBaseEffect) SetUseConstantColor(value unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setUseConstantColor:"), value)
}



