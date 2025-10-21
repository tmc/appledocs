// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MultiArrayShapeConstraint] class.
var (
	MultiArrayShapeConstraintClass     _MultiArrayShapeConstraintClass
	MultiArrayShapeConstraintClassOnce sync.Once
)

func getMultiArrayShapeConstraintClass() _MultiArrayShapeConstraintClass {
	MultiArrayShapeConstraintClassOnce.Do(func() {
		MultiArrayShapeConstraintClass = _MultiArrayShapeConstraintClass{objc.GetClass("MLMultiArrayShapeConstraint")}
	})
	return MultiArrayShapeConstraintClass
}

type _MultiArrayShapeConstraintClass struct {
	class objc.Class
}

// An interface definition for the [MultiArrayShapeConstraint] class.
type IMultiArrayShapeConstraint interface {
	objectivec.IObject
}

// The lists of shapes or ranges of shapes that constrain a multiarray feature.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLMultiArrayShapeConstraint
type MultiArrayShapeConstraint struct {
	objectivec.Object
}

// MultiArrayShapeConstraintFrom constructs a [MultiArrayShapeConstraint] from an unsafe.Pointer.
//
// The lists of shapes or ranges of shapes that constrain a multiarray feature.
func MultiArrayShapeConstraintFrom(ptr unsafe.Pointer) MultiArrayShapeConstraint {
	return MultiArrayShapeConstraint{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MultiArrayShapeConstraintClass) Alloc() MultiArrayShapeConstraint {
	rv := objc.Send[MultiArrayShapeConstraint](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MultiArrayShapeConstraintClass) New() MultiArrayShapeConstraint {
	rv := objc.Send[MultiArrayShapeConstraint](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MultiArrayShapeConstraint) Init() MultiArrayShapeConstraint {
	rv := objc.Send[MultiArrayShapeConstraint](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MultiArrayShapeConstraint) Autorelease() MultiArrayShapeConstraint {
	rv := objc.Send[MultiArrayShapeConstraint](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMultiArrayShapeConstraint creates a new MultiArrayShapeConstraint instance.
func NewMultiArrayShapeConstraint() MultiArrayShapeConstraint {
	return getMultiArrayShapeConstraintClass().New()
}


// The type for the multi array.
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlmultiarrayconstraint/datatype
func (m_ MultiArrayShapeConstraint) DataType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("dataType"))
	return rv
}


// SetDataType sets the value of the dataType property.
// The type for the multi array.

//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlmultiarrayconstraint/datatype
func (m_ MultiArrayShapeConstraint) SetDataType(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDataType:"), value)
}

// The shape of the multi array.
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlmultiarrayconstraint/shape
func (m_ MultiArrayShapeConstraint) Shape() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("shape"))
	return rv
}


// SetShape sets the value of the shape property.
// The shape of the multi array.

//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlmultiarrayconstraint/shape
func (m_ MultiArrayShapeConstraint) SetShape(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setShape:"), value)
}

// The constraint on the shape of the multiarray.
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlmultiarrayconstraint/shapeconstraint
func (m_ MultiArrayShapeConstraint) ShapeConstraint() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("shapeConstraint"))
	return rv
}


// SetShapeConstraint sets the value of the shapeConstraint property.
// The constraint on the shape of the multiarray.

//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlmultiarrayconstraint/shapeconstraint
func (m_ MultiArrayShapeConstraint) SetShapeConstraint(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setShapeConstraint:"), value)
}

// Array of allowed shapes for a multiarray feature.
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlmultiarrayshapeconstraint/enumeratedshapes
func (m_ MultiArrayShapeConstraint) EnumeratedShapes() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("enumeratedShapes"))
	return rv
}


// SetEnumeratedShapes sets the value of the enumeratedShapes property.
// Array of allowed shapes for a multiarray feature.

//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlmultiarrayshapeconstraint/enumeratedshapes
func (m_ MultiArrayShapeConstraint) SetEnumeratedShapes(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setEnumeratedShapes:"), value)
}

// The allowable range for a dimention of the multiarray.
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlmultiarrayshapeconstraint/sizerangefordimension
func (m_ MultiArrayShapeConstraint) SizeRangeForDimension() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("sizeRangeForDimension"))
	return rv
}


// SetSizeRangeForDimension sets the value of the sizeRangeForDimension property.
// The allowable range for a dimention of the multiarray.

//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlmultiarrayshapeconstraint/sizerangefordimension
func (m_ MultiArrayShapeConstraint) SetSizeRangeForDimension(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSizeRangeForDimension:"), value)
}

// The type of the shape constraint.
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlmultiarrayshapeconstraint/type
func (m_ MultiArrayShapeConstraint) Type() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("type"))
	return rv
}


// SetType sets the value of the type property.
// The type of the shape constraint.

//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlmultiarrayshapeconstraint/type
func (m_ MultiArrayShapeConstraint) SetType(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setType:"), value)
}



