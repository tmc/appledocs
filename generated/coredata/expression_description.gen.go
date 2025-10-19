// Code generated from Apple documentation for CoreData. DO NOT EDIT.

package coredata

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [ExpressionDescription] class.
var expressionDescriptionClass = _ExpressionDescriptionClass{objc.GetClass("NSExpressionDescription")}

type _ExpressionDescriptionClass struct {
	class objc.Class
}

// An interface definition for the [ExpressionDescription] class.
type IExpressionDescription interface {
	IPropertyDescription
}

// An object that describes an expression to include with a fetch request. [Full Topic]
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

// New creates and returns a new instance with a +1 retain count.
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
	return expressionDescriptionClass.New()
}




