// Code generated from Apple documentation for Accessibility. DO NOT EDIT.

package accessibility

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [AXMathExpressionUnderOver] class.
var (
	AXMathExpressionUnderOverClass     _AXMathExpressionUnderOverClass
	AXMathExpressionUnderOverClassOnce sync.Once
)

func getAXMathExpressionUnderOverClass() _AXMathExpressionUnderOverClass {
	AXMathExpressionUnderOverClassOnce.Do(func() {
		AXMathExpressionUnderOverClass = _AXMathExpressionUnderOverClass{objc.GetClass("AXMathExpressionUnderOver")}
	})
	return AXMathExpressionUnderOverClass
}

type _AXMathExpressionUnderOverClass struct {
	class objc.Class
}

// An interface definition for the [AXMathExpressionUnderOver] class.
type IAXMathExpressionUnderOver interface {
	IAXMathExpression
}

//
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXMathExpressionUnderOver
type AXMathExpressionUnderOver struct {
	AXMathExpression
}

// AXMathExpressionUnderOverFrom constructs a [AXMathExpressionUnderOver] from an unsafe.Pointer.
func AXMathExpressionUnderOverFrom(ptr unsafe.Pointer) AXMathExpressionUnderOver {
	return AXMathExpressionUnderOver{
		AXMathExpression: AXMathExpressionFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ac _AXMathExpressionUnderOverClass) Alloc() AXMathExpressionUnderOver {
	rv := objc.Send[AXMathExpressionUnderOver](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _AXMathExpressionUnderOverClass) New() AXMathExpressionUnderOver {
	rv := objc.Send[AXMathExpressionUnderOver](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AXMathExpressionUnderOver) Init() AXMathExpressionUnderOver {
	rv := objc.Send[AXMathExpressionUnderOver](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AXMathExpressionUnderOver) Autorelease() AXMathExpressionUnderOver {
	rv := objc.Send[AXMathExpressionUnderOver](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAXMathExpressionUnderOver creates a new AXMathExpressionUnderOver instance.
func NewAXMathExpressionUnderOver() AXMathExpressionUnderOver {
	return getAXMathExpressionUnderOverClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXMathExpressionUnderOver/init(baseExpression:underExpression:overExpression:)
func NewAXMathExpressionUnderOverWithBaseExpressionUnderExpressionOverExpression(baseExpression IAXMathExpression, underExpression IAXMathExpression, overExpression IAXMathExpression) AXMathExpressionUnderOver {
	instance := getAXMathExpressionUnderOverClass().Alloc()
	rv := objc.Send[AXMathExpressionUnderOver](instance.ID, objc.Sel("initWithBaseExpression:underExpression:overExpression:"), baseExpression, underExpression, overExpression)
	rv.Autorelease()
	return rv
}


//
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXMathExpressionUnderOver/underExpression
func (a_ AXMathExpressionUnderOver) UnderExpression() AXMathExpression {
	rv := objc.Send[AXMathExpression](a_.ID, objc.Sel("underExpression"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/accessibility/axmathexpressionunderover/baseexpression
func (a_ AXMathExpressionUnderOver) BaseExpression() AXMathExpression {
	rv := objc.Send[AXMathExpression](a_.ID, objc.Sel("baseExpression"))
	return rv
}


// SetBaseExpression sets the value of the baseExpression property.
//
// [Full Topic]: https://developer.apple.com/documentation/accessibility/axmathexpressionunderover/baseexpression
func (a_ AXMathExpressionUnderOver) SetBaseExpression(value IAXMathExpression) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setBaseExpression:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/accessibility/axmathexpressionunderover/overexpression
func (a_ AXMathExpressionUnderOver) OverExpression() AXMathExpression {
	rv := objc.Send[AXMathExpression](a_.ID, objc.Sel("overExpression"))
	return rv
}


// SetOverExpression sets the value of the overExpression property.
//
// [Full Topic]: https://developer.apple.com/documentation/accessibility/axmathexpressionunderover/overexpression
func (a_ AXMathExpressionUnderOver) SetOverExpression(value IAXMathExpression) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setOverExpression:"), value)
}


