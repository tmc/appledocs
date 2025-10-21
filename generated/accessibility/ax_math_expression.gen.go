// Code generated from Apple documentation for Accessibility. DO NOT EDIT.

package accessibility

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [AXMathExpression] class.
var (
	AXMathExpressionClass     _AXMathExpressionClass
	AXMathExpressionClassOnce sync.Once
)

func getAXMathExpressionClass() _AXMathExpressionClass {
	AXMathExpressionClassOnce.Do(func() {
		AXMathExpressionClass = _AXMathExpressionClass{objc.GetClass("AXMathExpression")}
	})
	return AXMathExpressionClass
}

type _AXMathExpressionClass struct {
	class objc.Class
}

// An interface definition for the [AXMathExpression] class.
type IAXMathExpression interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXMathExpression
type AXMathExpression struct {
	objectivec.Object
}

// AXMathExpressionFrom constructs a [AXMathExpression] from an unsafe.Pointer.
func AXMathExpressionFrom(ptr unsafe.Pointer) AXMathExpression {
	return AXMathExpression{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ac _AXMathExpressionClass) Alloc() AXMathExpression {
	rv := objc.Send[AXMathExpression](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _AXMathExpressionClass) New() AXMathExpression {
	rv := objc.Send[AXMathExpression](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AXMathExpression) Init() AXMathExpression {
	rv := objc.Send[AXMathExpression](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AXMathExpression) Autorelease() AXMathExpression {
	rv := objc.Send[AXMathExpression](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAXMathExpression creates a new AXMathExpression instance.
func NewAXMathExpression() AXMathExpression {
	return getAXMathExpressionClass().New()
}




