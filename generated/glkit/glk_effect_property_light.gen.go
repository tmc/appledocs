// Code generated from Apple documentation for GLKit. DO NOT EDIT.

package glkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class GLKEffectPropertyLight */


/* debug [class_header]: Header for GLKEffectPropertyLight */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for GLKEffectPropertyLight */
// An interface definition for the [GLKEffectPropertyLight] class.
type IGLKEffectPropertyLight interface {
	IGLKEffectProperty
	
/* debug [class_interface_properties]: Properties for GLKEffectPropertyLight */
	// properties:
	AmbientColor() GLKVector4 /* typedef */
	SetAmbientColor(value GLKVector4 /* typedef */)
	ConstantAttenuation() unsafe.Pointer
	SetConstantAttenuation(value unsafe.Pointer)
	DiffuseColor() GLKVector4 /* typedef */
	SetDiffuseColor(value GLKVector4 /* typedef */)
	Enabled() unsafe.Pointer
	SetEnabled(value unsafe.Pointer)
	LinearAttenuation() unsafe.Pointer
	SetLinearAttenuation(value unsafe.Pointer)
	Position() GLKVector4 /* typedef */
	SetPosition(value GLKVector4 /* typedef */)
	QuadraticAttenuation() unsafe.Pointer
	SetQuadraticAttenuation(value unsafe.Pointer)
	SpecularColor() GLKVector4 /* typedef */
	SetSpecularColor(value GLKVector4 /* typedef */)
	SpotCutoff() unsafe.Pointer
	SetSpotCutoff(value unsafe.Pointer)
	SpotDirection() GLKVector3 /* typedef */
	SetSpotDirection(value GLKVector3 /* typedef */)
	SpotExponent() unsafe.Pointer
	SetSpotExponent(value unsafe.Pointer)
	Transform() IGLKEffectPropertyTransform
	SetTransform(value IGLKEffectPropertyTransform)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for GLKEffectPropertyLight */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for GLKEffectPropertyLight */
// Alloc allocates a new instance without initialization.
func (gc _GLKEffectPropertyLightClass) Alloc() GLKEffectPropertyLight {
	rv := objc.Send[GLKEffectPropertyLight](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for GLKEffectPropertyLight */
// Lighting information for use in GLKit rendering effects.
//
// The lighting model implemented by is identical to the lighting model implemented in OpenGL ES 1.1; each light interacts with any material properties on the effect to determine the intensity and color that particular light contributes to the scene at a fragment. There are three basic kinds of lights: directional, point and spotlights. A directional light is considered to be infinitely far away, and always directs light in the same direction. To create a directional light, set the property to a vector whose , , and components specify the direction to the light (that is, the negation of the direction the light is travelling in), and whose component is set to . A point light is placed at a position within the scene, and emits light in all directions. To create a directional light, set the property to a vector whose , , and components specify the homogenous coordinates for the position of the light in the scene. The component is typically set to and must not be set to . The intensity of a point light is adjusted using an distance attenuation function. This function is controlled by adjusting the , or properties. The default values for these properties create a light whose intensity is constant over distance. A spotlight is placed at a position within the scene, and emits light in a specific direction in a cone. To create a spotlight, set the property to a vector whose , , and components specify the homogenous coordinates for the position of the light in the scene. The component must not be set to . Then, set the property to a vector that specifies the direction of the light and the to a value less than . Like a point light, a spotlight’s intensity can be adjusted using the distance attenuation properties. A spotlight’s intensity can also be changed as a function of the spotlight angle by setting the property. The default value for a spotlight creates a spotlight whose intensity is not affected by the angle. That is, the spotlight radiates the same amount of light at the center and at the edge of the cone. Lighting calculations are performed in eye-space coordinates.  The eye-space coordinates for the position and the spot direction are calculated at the precise moment that new position values are specified and may be affected by other properties of the effect. For more information, see .


// Lighting information for use in GLKit rendering effects.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for GLKEffectPropertyLight *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for GLKEffectPropertyLight */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for GLKEffectPropertyLight */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for GLKEffectPropertyLight */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for GLKEffectPropertyLight */

// The ambient portion of the light.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKEffectPropertyLight/ambientColor
func (g_ GLKEffectPropertyLight) AmbientColor() GLKVector4 /* typedef */ {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("ambientColor"))
	return rv
}/* debug [instance_properties/getter]: ambientColor */


// The ambient portion of the light.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKEffectPropertyLight/ambientColor
func (g_ GLKEffectPropertyLight) SetAmbientColor(value GLKVector4 /* typedef */) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setAmbientColor:"), value)
}/* debug [instance_properties/setter]: ambientColor */


// A constant factor applied to the attenuation of a point light or spotlight.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKEffectPropertyLight/constantAttenuation
func (g_ GLKEffectPropertyLight) ConstantAttenuation() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("constantAttenuation"))
	return rv
}/* debug [instance_properties/getter]: constantAttenuation */


// A constant factor applied to the attenuation of a point light or spotlight.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKEffectPropertyLight/constantAttenuation
func (g_ GLKEffectPropertyLight) SetConstantAttenuation(value unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setConstantAttenuation:"), value)
}/* debug [instance_properties/setter]: constantAttenuation */


// The diffuse portion of the light.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKEffectPropertyLight/diffuseColor
func (g_ GLKEffectPropertyLight) DiffuseColor() GLKVector4 /* typedef */ {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("diffuseColor"))
	return rv
}/* debug [instance_properties/getter]: diffuseColor */


// The diffuse portion of the light.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKEffectPropertyLight/diffuseColor
func (g_ GLKEffectPropertyLight) SetDiffuseColor(value GLKVector4 /* typedef */) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setDiffuseColor:"), value)
}/* debug [instance_properties/setter]: diffuseColor */


// A Boolean value that indicates whether calculations should be performed on this light.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKEffectPropertyLight/enabled
func (g_ GLKEffectPropertyLight) Enabled() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("enabled"))
	return rv
}/* debug [instance_properties/getter]: enabled */


// A Boolean value that indicates whether calculations should be performed on this light.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKEffectPropertyLight/enabled
func (g_ GLKEffectPropertyLight) SetEnabled(value unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setEnabled:"), value)
}/* debug [instance_properties/setter]: enabled */


// A linear factor applied to the attenuation of a point light or spotlight.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKEffectPropertyLight/linearAttenuation
func (g_ GLKEffectPropertyLight) LinearAttenuation() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("linearAttenuation"))
	return rv
}/* debug [instance_properties/getter]: linearAttenuation */


// A linear factor applied to the attenuation of a point light or spotlight.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKEffectPropertyLight/linearAttenuation
func (g_ GLKEffectPropertyLight) SetLinearAttenuation(value unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setLinearAttenuation:"), value)
}/* debug [instance_properties/setter]: linearAttenuation */


// The position of the light in world coordinates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKEffectPropertyLight/position
func (g_ GLKEffectPropertyLight) Position() GLKVector4 /* typedef */ {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("position"))
	return rv
}/* debug [instance_properties/getter]: position */


// The position of the light in world coordinates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKEffectPropertyLight/position
func (g_ GLKEffectPropertyLight) SetPosition(value GLKVector4 /* typedef */) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setPosition:"), value)
}/* debug [instance_properties/setter]: position */


// A quadratic factor applied to the attenuation of a point light or spotlight.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKEffectPropertyLight/quadraticAttenuation
func (g_ GLKEffectPropertyLight) QuadraticAttenuation() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("quadraticAttenuation"))
	return rv
}/* debug [instance_properties/getter]: quadraticAttenuation */


// A quadratic factor applied to the attenuation of a point light or spotlight.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKEffectPropertyLight/quadraticAttenuation
func (g_ GLKEffectPropertyLight) SetQuadraticAttenuation(value unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setQuadraticAttenuation:"), value)
}/* debug [instance_properties/setter]: quadraticAttenuation */


// The specular portion of the light.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKEffectPropertyLight/specularColor
func (g_ GLKEffectPropertyLight) SpecularColor() GLKVector4 /* typedef */ {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("specularColor"))
	return rv
}/* debug [instance_properties/getter]: specularColor */


// The specular portion of the light.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKEffectPropertyLight/specularColor
func (g_ GLKEffectPropertyLight) SetSpecularColor(value GLKVector4 /* typedef */) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setSpecularColor:"), value)
}/* debug [instance_properties/setter]: specularColor */


// The angle in degrees where the spotlight is cut off.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKEffectPropertyLight/spotCutoff
func (g_ GLKEffectPropertyLight) SpotCutoff() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("spotCutoff"))
	return rv
}/* debug [instance_properties/getter]: spotCutoff */


// The angle in degrees where the spotlight is cut off.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKEffectPropertyLight/spotCutoff
func (g_ GLKEffectPropertyLight) SetSpotCutoff(value unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setSpotCutoff:"), value)
}/* debug [instance_properties/setter]: spotCutoff */


// A vector indicating the direction the spotlight is projecting.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKEffectPropertyLight/spotDirection
func (g_ GLKEffectPropertyLight) SpotDirection() GLKVector3 /* typedef */ {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("spotDirection"))
	return rv
}/* debug [instance_properties/getter]: spotDirection */


// A vector indicating the direction the spotlight is projecting.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKEffectPropertyLight/spotDirection
func (g_ GLKEffectPropertyLight) SetSpotDirection(value GLKVector3 /* typedef */) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setSpotDirection:"), value)
}/* debug [instance_properties/setter]: spotDirection */


// A value indicating how focused the spotlight is.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKEffectPropertyLight/spotExponent
func (g_ GLKEffectPropertyLight) SpotExponent() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("spotExponent"))
	return rv
}/* debug [instance_properties/getter]: spotExponent */


// A value indicating how focused the spotlight is.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKEffectPropertyLight/spotExponent
func (g_ GLKEffectPropertyLight) SetSpotExponent(value unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setSpotExponent:"), value)
}/* debug [instance_properties/setter]: spotExponent */


// A transform applied to the light’s position and direction before calculating the contribution of the light.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKEffectPropertyLight/transform
func (g_ GLKEffectPropertyLight) Transform() IGLKEffectPropertyTransform {
	rv := objc.Send[GLKEffectPropertyTransform](g_.ID, objc.Sel("transform"))
	return rv
}/* debug [instance_properties/getter]: transform */


// A transform applied to the light’s position and direction before calculating the contribution of the light.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKEffectPropertyLight/transform
func (g_ GLKEffectPropertyLight) SetTransform(value IGLKEffectPropertyTransform) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setTransform:"), value)
}/* debug [instance_properties/setter]: transform */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class GLKEffectPropertyLight */



