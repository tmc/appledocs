// Code generated from Apple documentation for GLKit. DO NOT EDIT.

package glkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class GLKEffectPropertyTransform */


/* debug [class_header]: Header for GLKEffectPropertyTransform */
// The class instance for the [GLKEffectPropertyTransform] class.
var (
	GLKEffectPropertyTransformClass     _GLKEffectPropertyTransformClass
	GLKEffectPropertyTransformClassOnce sync.Once
)

func getGLKEffectPropertyTransformClass() _GLKEffectPropertyTransformClass {
	GLKEffectPropertyTransformClassOnce.Do(func() {
		GLKEffectPropertyTransformClass = _GLKEffectPropertyTransformClass{objc.GetClass("GLKEffectPropertyTransform")}
	})
	return GLKEffectPropertyTransformClass
}

type _GLKEffectPropertyTransformClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for GLKEffectPropertyTransform */
// An interface definition for the [GLKEffectPropertyTransform] class.
type IGLKEffectPropertyTransform interface {
	IGLKEffectProperty
	
/* debug [class_interface_properties]: Properties for GLKEffectPropertyTransform */
	// properties:
	ModelviewMatrix() GLKMatrix4 /* typedef */
	SetModelviewMatrix(value GLKMatrix4 /* typedef */)
	NormalMatrix() GLKMatrix3 /* typedef */
	ProjectionMatrix() GLKMatrix4 /* typedef */
	SetProjectionMatrix(value GLKMatrix4 /* typedef */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for GLKEffectPropertyTransform */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for GLKEffectPropertyTransform */
// Alloc allocates a new instance without initialization.
func (gc _GLKEffectPropertyTransformClass) Alloc() GLKEffectPropertyTransform {
	rv := objc.Send[GLKEffectPropertyTransform](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (gc _GLKEffectPropertyTransformClass) New() GLKEffectPropertyTransform {
	rv := objc.Send[GLKEffectPropertyTransform](objc.ID(gc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (g_ GLKEffectPropertyTransform) Init() GLKEffectPropertyTransform {
	rv := objc.Send[GLKEffectPropertyTransform](g_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g_ GLKEffectPropertyTransform) Autorelease() GLKEffectPropertyTransform {
	rv := objc.Send[GLKEffectPropertyTransform](g_.ID, objc.Sel("autorelease"))
	return rv
}

// NewGLKEffectPropertyTransform creates a new GLKEffectPropertyTransform instance.
func NewGLKEffectPropertyTransform() GLKEffectPropertyTransform {
	return getGLKEffectPropertyTransformClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for GLKEffectPropertyTransform */
// Coordinate transform information for use in GLKit rendering effects.
//
// The class defines properties that provide the coordinate transformations to be performed when rendering the effect.


// Coordinate transform information for use in GLKit rendering effects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKEffectPropertyTransform
type GLKEffectPropertyTransform struct {
	GLKEffectProperty
}

// GLKEffectPropertyTransformFrom constructs a [GLKEffectPropertyTransform] from an unsafe.Pointer.
//
// Coordinate transform information for use in GLKit rendering effects.
func GLKEffectPropertyTransformFrom(ptr unsafe.Pointer) GLKEffectPropertyTransform {
	return GLKEffectPropertyTransform{
		GLKEffectProperty: GLKEffectPropertyFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for GLKEffectPropertyTransform *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for GLKEffectPropertyTransform */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for GLKEffectPropertyTransform */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for GLKEffectPropertyTransform */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for GLKEffectPropertyTransform */

// The matrix used to transform position coordinates from world space to eye space.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKEffectPropertyTransform/modelviewMatrix
func (g_ GLKEffectPropertyTransform) ModelviewMatrix() GLKMatrix4 /* typedef */ {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("modelviewMatrix"))
	return rv
}/* debug [instance_properties/getter]: modelviewMatrix */


// The matrix used to transform position coordinates from world space to eye space.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKEffectPropertyTransform/modelviewMatrix
func (g_ GLKEffectPropertyTransform) SetModelviewMatrix(value GLKMatrix4 /* typedef */) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setModelviewMatrix:"), value)
}/* debug [instance_properties/setter]: modelviewMatrix */


// The matrix used to transform normal coordinates from world space to eye space.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKEffectPropertyTransform/normalMatrix
func (g_ GLKEffectPropertyTransform) NormalMatrix() GLKMatrix3 /* typedef */ {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("normalMatrix"))
	return rv
}/* debug [instance_properties/getter]: normalMatrix */


// The matrix used to transform position coordinates from eye space to projection space.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKEffectPropertyTransform/projectionMatrix
func (g_ GLKEffectPropertyTransform) ProjectionMatrix() GLKMatrix4 /* typedef */ {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("projectionMatrix"))
	return rv
}/* debug [instance_properties/getter]: projectionMatrix */


// The matrix used to transform position coordinates from eye space to projection space.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKEffectPropertyTransform/projectionMatrix
func (g_ GLKEffectPropertyTransform) SetProjectionMatrix(value GLKMatrix4 /* typedef */) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setProjectionMatrix:"), value)
}/* debug [instance_properties/setter]: projectionMatrix */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class GLKEffectPropertyTransform */



