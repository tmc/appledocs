// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MLNumericConstraint */


/* debug [class_header]: Header for MLNumericConstraint */
// The class instance for the [NumericConstraint] class.
var (
	NumericConstraintClass     _NumericConstraintClass
	NumericConstraintClassOnce sync.Once
)

func getNumericConstraintClass() _NumericConstraintClass {
	NumericConstraintClassOnce.Do(func() {
		NumericConstraintClass = _NumericConstraintClass{objc.GetClass("MLNumericConstraint")}
	})
	return NumericConstraintClass
}

type _NumericConstraintClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for NumericConstraint */
// An interface definition for the [NumericConstraint] class.
type INumericConstraint interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for NumericConstraint */
	// properties:
	EnumeratedNumbers() unsafe.Pointer
	MaxNumber() objc.IObject /* cross-framework: NSNumber */
	MinNumber() objc.IObject /* cross-framework: NSNumber */
	NumericConstraint() IMLNumericConstraint
	SetNumericConstraint(value IMLNumericConstraint)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for NumericConstraint */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for NumericConstraint */
// Alloc allocates a new instance without initialization.
func (nc _NumericConstraintClass) Alloc() NumericConstraint {
	rv := objc.Send[NumericConstraint](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (nc _NumericConstraintClass) New() NumericConstraint {
	rv := objc.Send[NumericConstraint](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NumericConstraint) Init() NumericConstraint {
	rv := objc.Send[NumericConstraint](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NumericConstraint) Autorelease() NumericConstraint {
	rv := objc.Send[NumericConstraint](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNumericConstraint creates a new NumericConstraint instance.
func NewNumericConstraint() NumericConstraint {
	return getNumericConstraintClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for NumericConstraint */
// The value limitations of a number.


// The value limitations of a number.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLNumericConstraint
type NumericConstraint struct {
	objectivec.Object
}

// NumericConstraintFrom constructs a [NumericConstraint] from an unsafe.Pointer.
//
// The value limitations of a number.
func NumericConstraintFrom(ptr unsafe.Pointer) NumericConstraint {
	return NumericConstraint{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for NumericConstraint *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for NumericConstraint */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for NumericConstraint */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for NumericConstraint */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for NumericConstraint */

// A set of the numbers allowed in this constraint.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLNumericConstraint/enumeratedNumbers
func (n_ NumericConstraint) EnumeratedNumbers() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("enumeratedNumbers"))
	return rv
}/* debug [instance_properties/getter]: enumeratedNumbers */


// The largest numerical value allowed by this constraint.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLNumericConstraint/maxNumber
func (n_ NumericConstraint) MaxNumber() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](n_.ID, objc.Sel("maxNumber"))
	return rv
}/* debug [instance_properties/getter]: maxNumber */


// The smallest numerical value allowed by this constraint.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLNumericConstraint/minNumber
func (n_ NumericConstraint) MinNumber() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](n_.ID, objc.Sel("minNumber"))
	return rv
}/* debug [instance_properties/getter]: minNumber */


// The constraints of this paramter description value, if and only if the value is numerical.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlparameterdescription/numericconstraint
func (n_ NumericConstraint) NumericConstraint() IMLNumericConstraint {
	rv := objc.Send[NumericConstraint](n_.ID, objc.Sel("numericConstraint"))
	return rv
}/* debug [instance_properties/getter]: numericConstraint */


// The constraints of this paramter description value, if and only if the value is numerical.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlparameterdescription/numericconstraint
func (n_ NumericConstraint) SetNumericConstraint(value IMLNumericConstraint) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setNumericConstraint:"), value)
}/* debug [instance_properties/setter]: numericConstraint */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MLNumericConstraint */



