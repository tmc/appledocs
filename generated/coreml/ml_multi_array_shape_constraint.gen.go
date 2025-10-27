// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
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
	

	// properties:
	EnumeratedShapes() []foundation.Array
	SizeRangeForDimension() []foundation.Value
	Type() MultiArrayShapeConstraintType
	DataType() MultiArrayDataType
	SetDataType(value MultiArrayDataType)
	Shape() foundation.foundation.INSNumber
	SetShape(value foundation.foundation.INSNumber)
	ShapeConstraint() IMLMultiArrayShapeConstraint
	SetShapeConstraint(value IMLMultiArrayShapeConstraint)


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (mc _MultiArrayShapeConstraintClass) Alloc() MultiArrayShapeConstraint {
	rv := objc.Send[MultiArrayShapeConstraint](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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





// The lists of shapes or ranges of shapes that constrain a multiarray feature.


// The lists of shapes or ranges of shapes that constrain a multiarray feature.
//
// [Full Topic]
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

























// Array of allowed shapes for a multiarray feature.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLMultiArrayShapeConstraint/enumeratedShapes
func (m_ MultiArrayShapeConstraint) EnumeratedShapes() []foundation.Array {
	rv := objc.Send[[]foundation.Array](m_.ID, objc.Sel("enumeratedShapes"))
	return rv
}


// The allowable range for a dimention of the multiarray.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLMultiArrayShapeConstraint/sizeRangeForDimension
func (m_ MultiArrayShapeConstraint) SizeRangeForDimension() []foundation.Value {
	rv := objc.Send[[]foundation.Value](m_.ID, objc.Sel("sizeRangeForDimension"))
	return rv
}


// The type of the shape constraint.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLMultiArrayShapeConstraint/type
func (m_ MultiArrayShapeConstraint) Type() MultiArrayShapeConstraintType {
	rv := objc.Send[MultiArrayShapeConstraintType](m_.ID, objc.Sel("type"))
	return rv
}


// The type for the multi array.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlmultiarrayconstraint/datatype
func (m_ MultiArrayShapeConstraint) DataType() MultiArrayDataType {
	rv := objc.Send[MultiArrayDataType](m_.ID, objc.Sel("dataType"))
	return rv
}


// The type for the multi array.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlmultiarrayconstraint/datatype
func (m_ MultiArrayShapeConstraint) SetDataType(value MultiArrayDataType) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDataType:"), value)
}


// The shape of the multi array.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlmultiarrayconstraint/shape
func (m_ MultiArrayShapeConstraint) Shape() foundation.foundation.INSNumber {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("shape"))
	return rv
}


// The shape of the multi array.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlmultiarrayconstraint/shape
func (m_ MultiArrayShapeConstraint) SetShape(value foundation.foundation.INSNumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setShape:"), value)
}


// The constraint on the shape of the multiarray.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlmultiarrayconstraint/shapeconstraint
func (m_ MultiArrayShapeConstraint) ShapeConstraint() IMLMultiArrayShapeConstraint {
	rv := objc.Send[MultiArrayShapeConstraint](m_.ID, objc.Sel("shapeConstraint"))
	return rv
}


// The constraint on the shape of the multiarray.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlmultiarrayconstraint/shapeconstraint
func (m_ MultiArrayShapeConstraint) SetShapeConstraint(value IMLMultiArrayShapeConstraint) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setShapeConstraint:"), value)
}








