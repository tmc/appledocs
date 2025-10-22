// Code generated from Apple documentation for GLKit. DO NOT EDIT.

package glkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [GLKEffectPropertyLight] class.
var (
	GLKEffectPropertyLightClass     _GLKEffectPropertyLightClass
	GLKEffectPropertyLightClassOnce sync.Once
)

func getGLKEffectPropertyLightClass() _GLKEffectPropertyLightClass {
	GLKEffectPropertyLightClassOnce.Do(func() {
		GLKEffectPropertyLightClass = _GLKEffectPropertyLightClass{objc.GetClass("GLKEffectPropertyLight")}
	})
	return GLKEffectPropertyLightClass
}

type _GLKEffectPropertyLightClass struct {
	class objc.Class
}

// An interface definition for the [GLKEffectPropertyLight] class.
type IGLKEffectPropertyLight interface {
	IGLKEffectProperty
	AmbientColor() GLKVector4
	SetAmbientColor(value IGLKVector4)
	ConstantAttenuation() unsafe.Pointer
	SetConstantAttenuation(value unsafe.Pointer)
	DiffuseColor() GLKVector4
	SetDiffuseColor(value IGLKVector4)
	Enabled() unsafe.Pointer
	SetEnabled(value unsafe.Pointer)
	LinearAttenuation() unsafe.Pointer
	SetLinearAttenuation(value unsafe.Pointer)
	Position() GLKVector4
	SetPosition(value IGLKVector4)
	QuadraticAttenuation() unsafe.Pointer
	SetQuadraticAttenuation(value unsafe.Pointer)
	SpecularColor() GLKVector4
	SetSpecularColor(value IGLKVector4)
	SpotCutoff() unsafe.Pointer
	SetSpotCutoff(value unsafe.Pointer)
	SpotDirection() GLKVector3
	SetSpotDirection(value IGLKVector3)
	SpotExponent() unsafe.Pointer
	SetSpotExponent(value unsafe.Pointer)
	Transform() GLKEffectPropertyTransform
	SetTransform(value IGLKEffectPropertyTransform)
}

// Lighting information for use in GLKit rendering effects.
//
// The lighting model implemented by is identical to the lighting model implemented in OpenGL ES 1.1; each light interacts with any material properties on the effect to determine the intensity and color that particular light contributes to the scene at a fragment. There are three basic kinds of lights: directional, point and spotlights. A directional light is considered to be infinitely far away, and always directs light in the same direction. To create a directional light, set the property to a vector whose , , and components specify the direction to the light (that is, the negation of the direction the light is travelling in), and whose component is set to . A point light is placed at a position within the scene, and emits light in all directions. To create a directional light, set the property to a vector whose , , and components specify the homogenous coordinates for the position of the light in the scene. The component is typically set to and must not be set to . The intensity of a point light is adjusted using an distance attenuation function. This function is controlled by adjusting the , or properties. The default values for these properties create a light whose intensity is constant over distance. A spotlight is placed at a position within the scene, and emits light in a specific direction in a cone. To create a spotlight, set the property to a vector whose , , and components specify the homogenous coordinates for the position of the light in the scene. The component must not be set to . Then, set the property to a vector that specifies the direction of the light and the to a value less than . Like a point light, a spotlight’s intensity can be adjusted using the distance attenuation properties. A spotlight’s intensity can also be changed as a function of the spotlight angle by setting the property. The default value for a spotlight creates a spotlight whose intensity is not affected by the angle. That is, the spotlight radiates the same amount of light at the center and at the edge of the cone. Lighting calculations are performed in eye-space coordinates.  The eye-space coordinates for the position and the spot direction are calculated at the precise moment that new position values are specified and may be affected by other properties of the effect. For more information, see .
//
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKEffectPropertyLight
type GLKEffectPropertyLight struct {
	GLKEffectProperty
}

// GLKEffectPropertyLightFrom constructs a [GLKEffectPropertyLight] from an unsafe.Pointer.
//
// Lighting information for use in GLKit rendering effects.
func GLKEffectPropertyLightFrom(ptr unsafe.Pointer) GLKEffectPropertyLight {
	return GLKEffectPropertyLight{
		GLKEffectProperty: GLKEffectPropertyFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (gc _GLKEffectPropertyLightClass) Alloc() GLKEffectPropertyLight {
	rv := objc.Send[GLKEffectPropertyLight](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (gc _GLKEffectPropertyLightClass) New() GLKEffectPropertyLight {
	rv := objc.Send[GLKEffectPropertyLight](objc.ID(gc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (g_ GLKEffectPropertyLight) Init() GLKEffectPropertyLight {
	rv := objc.Send[GLKEffectPropertyLight](g_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g_ GLKEffectPropertyLight) Autorelease() GLKEffectPropertyLight {
	rv := objc.Send[GLKEffectPropertyLight](g_.ID, objc.Sel("autorelease"))
	return rv
}

// NewGLKEffectPropertyLight creates a new GLKEffectPropertyLight instance.
func NewGLKEffectPropertyLight() GLKEffectPropertyLight {
	return getGLKEffectPropertyLightClass().New()
}


// The ambient portion of the light.
//
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKEffectPropertyLight/ambientColor
func (g_ GLKEffectPropertyLight) AmbientColor() GLKVector4 {
	rv := objc.Send[GLKVector4](g_.ID, objc.Sel("ambientColor"))
	return rv
}


// SetAmbientColor sets the value of the ambientColor property.
// The ambient portion of the light.

//
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKEffectPropertyLight/ambientColor
func (g_ GLKEffectPropertyLight) SetAmbientColor(value IGLKVector4) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setAmbientColor:"), value)
}

// A constant factor applied to the attenuation of a point light or spotlight.
//
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKEffectPropertyLight/constantAttenuation
func (g_ GLKEffectPropertyLight) ConstantAttenuation() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("constantAttenuation"))
	return rv
}


// SetConstantAttenuation sets the value of the constantAttenuation property.
// A constant factor applied to the attenuation of a point light or spotlight.

//
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKEffectPropertyLight/constantAttenuation
func (g_ GLKEffectPropertyLight) SetConstantAttenuation(value unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setConstantAttenuation:"), value)
}

// The diffuse portion of the light.
//
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKEffectPropertyLight/diffuseColor
func (g_ GLKEffectPropertyLight) DiffuseColor() GLKVector4 {
	rv := objc.Send[GLKVector4](g_.ID, objc.Sel("diffuseColor"))
	return rv
}


// SetDiffuseColor sets the value of the diffuseColor property.
// The diffuse portion of the light.

//
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKEffectPropertyLight/diffuseColor
func (g_ GLKEffectPropertyLight) SetDiffuseColor(value IGLKVector4) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setDiffuseColor:"), value)
}

// A Boolean value that indicates whether calculations should be performed on this light.
//
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKEffectPropertyLight/enabled
func (g_ GLKEffectPropertyLight) Enabled() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("enabled"))
	return rv
}


// SetEnabled sets the value of the enabled property.
// A Boolean value that indicates whether calculations should be performed on this light.

//
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKEffectPropertyLight/enabled
func (g_ GLKEffectPropertyLight) SetEnabled(value unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setEnabled:"), value)
}

// A linear factor applied to the attenuation of a point light or spotlight.
//
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKEffectPropertyLight/linearAttenuation
func (g_ GLKEffectPropertyLight) LinearAttenuation() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("linearAttenuation"))
	return rv
}


// SetLinearAttenuation sets the value of the linearAttenuation property.
// A linear factor applied to the attenuation of a point light or spotlight.

//
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKEffectPropertyLight/linearAttenuation
func (g_ GLKEffectPropertyLight) SetLinearAttenuation(value unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setLinearAttenuation:"), value)
}

// The position of the light in world coordinates.
//
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKEffectPropertyLight/position
func (g_ GLKEffectPropertyLight) Position() GLKVector4 {
	rv := objc.Send[GLKVector4](g_.ID, objc.Sel("position"))
	return rv
}


// SetPosition sets the value of the position property.
// The position of the light in world coordinates.

//
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKEffectPropertyLight/position
func (g_ GLKEffectPropertyLight) SetPosition(value IGLKVector4) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setPosition:"), value)
}

// A quadratic factor applied to the attenuation of a point light or spotlight.
//
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKEffectPropertyLight/quadraticAttenuation
func (g_ GLKEffectPropertyLight) QuadraticAttenuation() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("quadraticAttenuation"))
	return rv
}


// SetQuadraticAttenuation sets the value of the quadraticAttenuation property.
// A quadratic factor applied to the attenuation of a point light or spotlight.

//
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKEffectPropertyLight/quadraticAttenuation
func (g_ GLKEffectPropertyLight) SetQuadraticAttenuation(value unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setQuadraticAttenuation:"), value)
}

// The specular portion of the light.
//
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKEffectPropertyLight/specularColor
func (g_ GLKEffectPropertyLight) SpecularColor() GLKVector4 {
	rv := objc.Send[GLKVector4](g_.ID, objc.Sel("specularColor"))
	return rv
}


// SetSpecularColor sets the value of the specularColor property.
// The specular portion of the light.

//
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKEffectPropertyLight/specularColor
func (g_ GLKEffectPropertyLight) SetSpecularColor(value IGLKVector4) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setSpecularColor:"), value)
}

// The angle in degrees where the spotlight is cut off.
//
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKEffectPropertyLight/spotCutoff
func (g_ GLKEffectPropertyLight) SpotCutoff() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("spotCutoff"))
	return rv
}


// SetSpotCutoff sets the value of the spotCutoff property.
// The angle in degrees where the spotlight is cut off.

//
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKEffectPropertyLight/spotCutoff
func (g_ GLKEffectPropertyLight) SetSpotCutoff(value unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setSpotCutoff:"), value)
}

// A vector indicating the direction the spotlight is projecting.
//
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKEffectPropertyLight/spotDirection
func (g_ GLKEffectPropertyLight) SpotDirection() GLKVector3 {
	rv := objc.Send[GLKVector3](g_.ID, objc.Sel("spotDirection"))
	return rv
}


// SetSpotDirection sets the value of the spotDirection property.
// A vector indicating the direction the spotlight is projecting.

//
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKEffectPropertyLight/spotDirection
func (g_ GLKEffectPropertyLight) SetSpotDirection(value IGLKVector3) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setSpotDirection:"), value)
}

// A value indicating how focused the spotlight is.
//
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKEffectPropertyLight/spotExponent
func (g_ GLKEffectPropertyLight) SpotExponent() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("spotExponent"))
	return rv
}


// SetSpotExponent sets the value of the spotExponent property.
// A value indicating how focused the spotlight is.

//
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKEffectPropertyLight/spotExponent
func (g_ GLKEffectPropertyLight) SetSpotExponent(value unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setSpotExponent:"), value)
}

// A transform applied to the light’s position and direction before calculating the contribution of the light.
//
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKEffectPropertyLight/transform
func (g_ GLKEffectPropertyLight) Transform() GLKEffectPropertyTransform {
	rv := objc.Send[GLKEffectPropertyTransform](g_.ID, objc.Sel("transform"))
	return rv
}


// SetTransform sets the value of the transform property.
// A transform applied to the light’s position and direction before calculating the contribution of the light.

//
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKEffectPropertyLight/transform
func (g_ GLKEffectPropertyLight) SetTransform(value IGLKEffectPropertyTransform) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setTransform:"), value)
}



