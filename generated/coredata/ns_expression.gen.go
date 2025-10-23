// Code generated from Apple documentation for CoreData. DO NOT EDIT.

package coredata

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [Expression] class.
var (
	ExpressionClass     _ExpressionClass
	ExpressionClassOnce sync.Once
)

func getExpressionClass() _ExpressionClass {
	ExpressionClassOnce.Do(func() {
		ExpressionClass = _ExpressionClass{objc.GetClass("NSExpression")}
	})
	return ExpressionClass
}

type _ExpressionClass struct {
	class objc.Class
}

// An interface definition for the [Expression] class.
type IExpression interface {
	objectivec.IObject
	// properties:
	// methods:
}

// A parent class referenced by other CoreData classes.


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

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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
	return getExpressionClass().New()
}




