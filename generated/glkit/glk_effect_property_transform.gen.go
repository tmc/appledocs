// Code generated from Apple documentation for GLKit. DO NOT EDIT.

package glkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

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

// An interface definition for the [GLKEffectPropertyTransform] class.
type IGLKEffectPropertyTransform interface {
	IGLKEffectProperty
	// properties:
	ModelviewMatrix() GLKMatrix4
	SetModelviewMatrix(value GLKMatrix4)
	NormalMatrix() GLKMatrix3
	ProjectionMatrix() GLKMatrix4
	SetProjectionMatrix(value GLKMatrix4)
	// methods:
}

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

// Alloc allocates a new instance without initialization.
func (gc _GLKEffectPropertyTransformClass) Alloc() GLKEffectPropertyTransform {
	rv := objc.Send[GLKEffectPropertyTransform](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// The matrix used to transform position coordinates from world space to eye space.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKEffectPropertyTransform/modelviewMatrix
func (g_ GLKEffectPropertyTransform) ModelviewMatrix() GLKMatrix4 {
	rv := objc.Send[GLKMatrix4](g_.ID, objc.Sel("modelviewMatrix"))
	return rv
}


// The matrix used to transform position coordinates from world space to eye space.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKEffectPropertyTransform/modelviewMatrix
func (g_ GLKEffectPropertyTransform) SetModelviewMatrix(value GLKMatrix4) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setModelviewMatrix:"), value)
}


// The matrix used to transform normal coordinates from world space to eye space.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKEffectPropertyTransform/normalMatrix
func (g_ GLKEffectPropertyTransform) NormalMatrix() GLKMatrix3 {
	rv := objc.Send[GLKMatrix3](g_.ID, objc.Sel("normalMatrix"))
	return rv
}


// The matrix used to transform position coordinates from eye space to projection space.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKEffectPropertyTransform/projectionMatrix
func (g_ GLKEffectPropertyTransform) ProjectionMatrix() GLKMatrix4 {
	rv := objc.Send[GLKMatrix4](g_.ID, objc.Sel("projectionMatrix"))
	return rv
}


// The matrix used to transform position coordinates from eye space to projection space.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKEffectPropertyTransform/projectionMatrix
func (g_ GLKEffectPropertyTransform) SetProjectionMatrix(value GLKMatrix4) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setProjectionMatrix:"), value)
}



