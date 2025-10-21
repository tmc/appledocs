// Code generated from Apple documentation for Accessibility. DO NOT EDIT.

package accessibility

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [AXMathExpressionFraction] class.
var (
	AXMathExpressionFractionClass     _AXMathExpressionFractionClass
	AXMathExpressionFractionClassOnce sync.Once
)

func getAXMathExpressionFractionClass() _AXMathExpressionFractionClass {
	AXMathExpressionFractionClassOnce.Do(func() {
		AXMathExpressionFractionClass = _AXMathExpressionFractionClass{objc.GetClass("AXMathExpressionFraction")}
	})
	return AXMathExpressionFractionClass
}

type _AXMathExpressionFractionClass struct {
	class objc.Class
}

// An interface definition for the [AXMathExpressionFraction] class.
type IAXMathExpressionFraction interface {
	IAXMathExpression
}

//
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXMathExpressionFraction
type AXMathExpressionFraction struct {
	AXMathExpression
}

// AXMathExpressionFractionFrom constructs a [AXMathExpressionFraction] from an unsafe.Pointer.
func AXMathExpressionFractionFrom(ptr unsafe.Pointer) AXMathExpressionFraction {
	return AXMathExpressionFraction{
		AXMathExpression: AXMathExpressionFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ac _AXMathExpressionFractionClass) Alloc() AXMathExpressionFraction {
	rv := objc.Send[AXMathExpressionFraction](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _AXMathExpressionFractionClass) New() AXMathExpressionFraction {
	rv := objc.Send[AXMathExpressionFraction](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AXMathExpressionFraction) Init() AXMathExpressionFraction {
	rv := objc.Send[AXMathExpressionFraction](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AXMathExpressionFraction) Autorelease() AXMathExpressionFraction {
	rv := objc.Send[AXMathExpressionFraction](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAXMathExpressionFraction creates a new AXMathExpressionFraction instance.
func NewAXMathExpressionFraction() AXMathExpressionFraction {
	return getAXMathExpressionFractionClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXMathExpressionFraction/denimonatorExpression
func (a_ AXMathExpressionFraction) DenimonatorExpression() AXMathExpression {
	rv := objc.Send[AXMathExpression](a_.ID, objc.Sel("denimonatorExpression"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/accessibility/axmathexpressionfraction/numeratorexpression
func (a_ AXMathExpressionFraction) NumeratorExpression() AXMathExpression {
	rv := objc.Send[AXMathExpression](a_.ID, objc.Sel("numeratorExpression"))
	return rv
}


// SetNumeratorExpression sets the value of the numeratorExpression property.
//
// [Full Topic]: https://developer.apple.com/documentation/accessibility/axmathexpressionfraction/numeratorexpression
func (a_ AXMathExpressionFraction) SetNumeratorExpression(value IAXMathExpression) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setNumeratorExpression:"), value)
}



