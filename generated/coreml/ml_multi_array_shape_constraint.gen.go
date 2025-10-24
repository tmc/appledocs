// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MLMultiArrayShapeConstraint */


/* debug [class_header]: Header for MLMultiArrayShapeConstraint */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MultiArrayShapeConstraint */
// An interface definition for the [MultiArrayShapeConstraint] class.
type IMultiArrayShapeConstraint interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MultiArrayShapeConstraint */
	// properties:
	EnumeratedShapes() []foundation.Array
	SizeRangeForDimension() []foundation.Value
	Type() MultiArrayShapeConstraintType
	DataType() MultiArrayDataType
	SetDataType(value MultiArrayDataType)
	Shape() objc.IObject /* cross-framework: NSNumber */
	SetShape(value objc.IObject /* cross-framework: NSNumber */)
	ShapeConstraint() IMLMultiArrayShapeConstraint
	SetShapeConstraint(value IMLMultiArrayShapeConstraint)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MultiArrayShapeConstraint */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MultiArrayShapeConstraint */
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MultiArrayShapeConstraint */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MultiArrayShapeConstraint *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MultiArrayShapeConstraint */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MultiArrayShapeConstraint */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MultiArrayShapeConstraint */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MultiArrayShapeConstraint */

// Array of allowed shapes for a multiarray feature.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLMultiArrayShapeConstraint/enumeratedShapes
func (m_ MultiArrayShapeConstraint) EnumeratedShapes() []foundation.Array {
	rv := objc.Send[[]foundation.Array](m_.ID, objc.Sel("enumeratedShapes"))
	return rv
}/* debug [instance_properties/getter]: enumeratedShapes */


// The allowable range for a dimention of the multiarray.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLMultiArrayShapeConstraint/sizeRangeForDimension
func (m_ MultiArrayShapeConstraint) SizeRangeForDimension() []foundation.Value {
	rv := objc.Send[[]foundation.Value](m_.ID, objc.Sel("sizeRangeForDimension"))
	return rv
}/* debug [instance_properties/getter]: sizeRangeForDimension */


// The type of the shape constraint.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLMultiArrayShapeConstraint/type
func (m_ MultiArrayShapeConstraint) Type() MultiArrayShapeConstraintType {
	rv := objc.Send[MultiArrayShapeConstraintType](m_.ID, objc.Sel("type"))
	return rv
}/* debug [instance_properties/getter]: type */


// The type for the multi array.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlmultiarrayconstraint/datatype
func (m_ MultiArrayShapeConstraint) DataType() MultiArrayDataType {
	rv := objc.Send[MultiArrayDataType](m_.ID, objc.Sel("dataType"))
	return rv
}/* debug [instance_properties/getter]: dataType */


// The type for the multi array.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlmultiarrayconstraint/datatype
func (m_ MultiArrayShapeConstraint) SetDataType(value MultiArrayDataType) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDataType:"), value)
}/* debug [instance_properties/setter]: dataType */


// The shape of the multi array.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlmultiarrayconstraint/shape
func (m_ MultiArrayShapeConstraint) Shape() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("shape"))
	return rv
}/* debug [instance_properties/getter]: shape */


// The shape of the multi array.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlmultiarrayconstraint/shape
func (m_ MultiArrayShapeConstraint) SetShape(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setShape:"), value)
}/* debug [instance_properties/setter]: shape */


// The constraint on the shape of the multiarray.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlmultiarrayconstraint/shapeconstraint
func (m_ MultiArrayShapeConstraint) ShapeConstraint() IMLMultiArrayShapeConstraint {
	rv := objc.Send[MultiArrayShapeConstraint](m_.ID, objc.Sel("shapeConstraint"))
	return rv
}/* debug [instance_properties/getter]: shapeConstraint */


// The constraint on the shape of the multiarray.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlmultiarrayconstraint/shapeconstraint
func (m_ MultiArrayShapeConstraint) SetShapeConstraint(value IMLMultiArrayShapeConstraint) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setShapeConstraint:"), value)
}/* debug [instance_properties/setter]: shapeConstraint */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MLMultiArrayShapeConstraint */



