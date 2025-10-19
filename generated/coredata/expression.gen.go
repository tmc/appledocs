// Code generated from Apple documentation for CoreData. DO NOT EDIT.

package coredata

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [Expression] class.
var expressionClass = _ExpressionClass{objc.GetClass("NSExpression")}

type _ExpressionClass struct {
	class objc.Class
}

// An interface definition for the [Expression] class.
type IExpression interface {
	objectivec.IObject
}

// A parent class referenced by other CoreData classes. [Full Topic]

type Expression struct {
	objectivec.Object
}

// ExpressionFrom constructs a [Expression] from an unsafe.Pointer.
//
// A parent class referenced by other CoreData classes.
func ExpressionFrom(ptr unsafe.Pointer) Expression {
	return Expression{objectivec.Object{objc.ID(ptr)}}
}
// Alloc allocates a new instance without initialization.
func (ec _ExpressionClass) Alloc() Expression {
	rv := objc.Send[Expression](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
func (ec _ExpressionClass) New() Expression {
	rv := objc.Send[Expression](objc.ID(ec.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (e_ Expression) Init() Expression {
	rv := objc.Send[Expression](e_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e_ Expression) Autorelease() Expression {
	rv := objc.Send[Expression](e_.ID, objc.Sel("autorelease"))
	return rv
}

// NewExpression creates a new Expression instance.
func NewExpression() Expression {
	return expressionClass.New()
}




