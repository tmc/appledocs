// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





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





// An interface definition for the [NumericConstraint] class.
type INumericConstraint interface {
	objectivec.IObject
	

	// properties:
	EnumeratedNumbers() unsafe.Pointer
	MaxNumber() foundation.foundation.INSNumber
	MinNumber() foundation.foundation.INSNumber
	NumericConstraint() IMLNumericConstraint
	SetNumericConstraint(value IMLNumericConstraint)


	

	// methods:


}





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

























// A set of the numbers allowed in this constraint.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLNumericConstraint/enumeratedNumbers
func (n_ NumericConstraint) EnumeratedNumbers() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("enumeratedNumbers"))
	return rv
}


// The largest numerical value allowed by this constraint.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLNumericConstraint/maxNumber
func (n_ NumericConstraint) MaxNumber() foundation.foundation.INSNumber {
	rv := objc.Send[foundation.NSNumber](n_.ID, objc.Sel("maxNumber"))
	return rv
}


// The smallest numerical value allowed by this constraint.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLNumericConstraint/minNumber
func (n_ NumericConstraint) MinNumber() foundation.foundation.INSNumber {
	rv := objc.Send[foundation.NSNumber](n_.ID, objc.Sel("minNumber"))
	return rv
}


// The constraints of this paramter description value, if and only if the value is numerical.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlparameterdescription/numericconstraint
func (n_ NumericConstraint) NumericConstraint() IMLNumericConstraint {
	rv := objc.Send[NumericConstraint](n_.ID, objc.Sel("numericConstraint"))
	return rv
}


// The constraints of this paramter description value, if and only if the value is numerical.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlparameterdescription/numericconstraint
func (n_ NumericConstraint) SetNumericConstraint(value IMLNumericConstraint) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setNumericConstraint:"), value)
}








