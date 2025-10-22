// Code generated from Apple documentation for Accessibility. DO NOT EDIT.

package accessibility

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [AXMathExpressionSubSuperscript] class.
var (
	AXMathExpressionSubSuperscriptClass     _AXMathExpressionSubSuperscriptClass
	AXMathExpressionSubSuperscriptClassOnce sync.Once
)

func getAXMathExpressionSubSuperscriptClass() _AXMathExpressionSubSuperscriptClass {
	AXMathExpressionSubSuperscriptClassOnce.Do(func() {
		AXMathExpressionSubSuperscriptClass = _AXMathExpressionSubSuperscriptClass{objc.GetClass("AXMathExpressionSubSuperscript")}
	})
	return AXMathExpressionSubSuperscriptClass
}

type _AXMathExpressionSubSuperscriptClass struct {
	class objc.Class
}

// An interface definition for the [AXMathExpressionSubSuperscript] class.
type IAXMathExpressionSubSuperscript interface {
	IAXMathExpression
	BaseExpression() AXMathExpression
	SuperscriptExpressions() []AXMathExpression
	SubscriptExpressions() AXMathExpression
	SetSubscriptExpressions(value IAXMathExpression)
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXMathExpressionSubSuperscript

type AXMathExpressionSubSuperscript struct {
	AXMathExpression
}

// AXMathExpressionSubSuperscriptFrom constructs a [AXMathExpressionSubSuperscript] from an unsafe.Pointer.
func AXMathExpressionSubSuperscriptFrom(ptr unsafe.Pointer) AXMathExpressionSubSuperscript {
	return AXMathExpressionSubSuperscript{
		AXMathExpression: AXMathExpressionFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ac _AXMathExpressionSubSuperscriptClass) Alloc() AXMathExpressionSubSuperscript {
	rv := objc.Send[AXMathExpressionSubSuperscript](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _AXMathExpressionSubSuperscriptClass) New() AXMathExpressionSubSuperscript {
	rv := objc.Send[AXMathExpressionSubSuperscript](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AXMathExpressionSubSuperscript) Init() AXMathExpressionSubSuperscript {
	rv := objc.Send[AXMathExpressionSubSuperscript](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AXMathExpressionSubSuperscript) Autorelease() AXMathExpressionSubSuperscript {
	rv := objc.Send[AXMathExpressionSubSuperscript](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAXMathExpressionSubSuperscript creates a new AXMathExpressionSubSuperscript instance.
func NewAXMathExpressionSubSuperscript() AXMathExpressionSubSuperscript {
	return getAXMathExpressionSubSuperscriptClass().New()
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXMathExpressionSubSuperscript/baseExpression

func (a_ AXMathExpressionSubSuperscript) BaseExpression() AXMathExpression {
	rv := objc.Send[AXMathExpression](a_.ID, objc.Sel("baseExpression"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXMathExpressionSubSuperscript/superscriptExpressions

func (a_ AXMathExpressionSubSuperscript) SuperscriptExpressions() []AXMathExpression {
	rv := objc.Send[[]AXMathExpression](a_.ID, objc.Sel("superscriptExpressions"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/accessibility/axmathexpressionsubsuperscript/subscriptexpressions

func (a_ AXMathExpressionSubSuperscript) SubscriptExpressions() AXMathExpression {
	rv := objc.Send[AXMathExpression](a_.ID, objc.Sel("subscriptExpressions"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/accessibility/axmathexpressionsubsuperscript/subscriptexpressions

func (a_ AXMathExpressionSubSuperscript) SetSubscriptExpressions(value IAXMathExpression) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setSubscriptExpressions:"), value)
}



