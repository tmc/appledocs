// Code generated from Apple documentation for CoreData. DO NOT EDIT.

package coredata

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [ExpressionDescription] class.
var (
	expressionDescriptionClass     _ExpressionDescriptionClass
	expressionDescriptionClassOnce sync.Once
)

func getExpressionDescriptionClass() _ExpressionDescriptionClass {
	expressionDescriptionClassOnce.Do(func() {
		expressionDescriptionClass = _ExpressionDescriptionClass{objc.GetClass("NSExpressionDescription")}
	})
	return expressionDescriptionClass
}

type _ExpressionDescriptionClass struct {
	class objc.Class
}

// An interface definition for the [ExpressionDescription] class.
type IExpressionDescription interface {
	IPropertyDescription
}

// An object that describes an expression to include with a fetch request.
//
// An expression description describes a value that a fetch request returns, which doesn’t appear as an attribute or relationship on an entity. For example, expressions can aggregate data, or transform an attribute’s value. You add expression descriptions to a fetch request using the method.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSExpressionDescription
type ExpressionDescription struct {
	PropertyDescription
}

// ExpressionDescriptionFrom constructs a [ExpressionDescription] from an unsafe.Pointer.
//
// An object that describes an expression to include with a fetch request.
func ExpressionDescriptionFrom(ptr unsafe.Pointer) ExpressionDescription {
	return ExpressionDescription{
		PropertyDescription: PropertyDescriptionFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ec _ExpressionDescriptionClass) Alloc() ExpressionDescription {
	rv := objc.Send[ExpressionDescription](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ec _ExpressionDescriptionClass) New() ExpressionDescription {
	rv := objc.Send[ExpressionDescription](objc.ID(ec.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (e_ ExpressionDescription) Init() ExpressionDescription {
	rv := objc.Send[ExpressionDescription](e_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e_ ExpressionDescription) Autorelease() ExpressionDescription {
	rv := objc.Send[ExpressionDescription](e_.ID, objc.Sel("autorelease"))
	return rv
}

// NewExpressionDescription creates a new ExpressionDescription instance.
func NewExpressionDescription() ExpressionDescription {
	return getExpressionDescriptionClass().New()
}


// The expression to evaluate.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSExpressionDescription/expression
func (e_ ExpressionDescription) Expression() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](e_.ID, objc.Sel("expression"))
	return rv
}

// SetExpression sets the value of the expression property.
// The expression to evaluate.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSExpressionDescription/expression
func (e_ ExpressionDescription) SetExpression(value unsafe.Pointer) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setExpression:"), value)
}
// The attribute type of the expression’s result.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSExpressionDescription/expressionResultType
func (e_ ExpressionDescription) ExpressionResultType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](e_.ID, objc.Sel("expressionResultType"))
	return rv
}

// SetExpressionResultType sets the value of the expressionResultType property.
// The attribute type of the expression’s result.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSExpressionDescription/expressionResultType
func (e_ ExpressionDescription) SetExpressionResultType(value unsafe.Pointer) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setExpressionResultType:"), value)
}


