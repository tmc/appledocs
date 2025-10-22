// Code generated from Apple documentation for Accessibility. DO NOT EDIT.

package accessibility

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [AXMathExpressionFenced] class.
var (
	AXMathExpressionFencedClass     _AXMathExpressionFencedClass
	AXMathExpressionFencedClassOnce sync.Once
)

func getAXMathExpressionFencedClass() _AXMathExpressionFencedClass {
	AXMathExpressionFencedClassOnce.Do(func() {
		AXMathExpressionFencedClass = _AXMathExpressionFencedClass{objc.GetClass("AXMathExpressionFenced")}
	})
	return AXMathExpressionFencedClass
}

type _AXMathExpressionFencedClass struct {
	class objc.Class
}

// An interface definition for the [AXMathExpressionFenced] class.
type IAXMathExpressionFenced interface {
	IAXMathExpression
	CloseString() string
	Expressions() []AXMathExpression
	OpenString() string
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXMathExpressionFenced

type AXMathExpressionFenced struct {
	AXMathExpression
}

// AXMathExpressionFencedFrom constructs a [AXMathExpressionFenced] from an unsafe.Pointer.
func AXMathExpressionFencedFrom(ptr unsafe.Pointer) AXMathExpressionFenced {
	return AXMathExpressionFenced{
		AXMathExpression: AXMathExpressionFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ac _AXMathExpressionFencedClass) Alloc() AXMathExpressionFenced {
	rv := objc.Send[AXMathExpressionFenced](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _AXMathExpressionFencedClass) New() AXMathExpressionFenced {
	rv := objc.Send[AXMathExpressionFenced](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AXMathExpressionFenced) Init() AXMathExpressionFenced {
	rv := objc.Send[AXMathExpressionFenced](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AXMathExpressionFenced) Autorelease() AXMathExpressionFenced {
	rv := objc.Send[AXMathExpressionFenced](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAXMathExpressionFenced creates a new AXMathExpressionFenced instance.
func NewAXMathExpressionFenced() AXMathExpressionFenced {
	return getAXMathExpressionFencedClass().New()
}




// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXMathExpressionFenced/init(expressions:open:close:)

func NewAXMathExpressionFencedWithExpressionsOpenStringCloseString(expressions []AXMathExpression, openString string, closeString string) AXMathExpressionFenced {
	instance := getAXMathExpressionFencedClass().Alloc()
	rv := objc.Send[AXMathExpressionFenced](instance.ID, objc.Sel("initWithExpressions:openString:closeString:"), expressions, objc.String(openString), objc.String(closeString))
	rv.Autorelease()
	return rv
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXMathExpressionFenced/closeString

func (a_ AXMathExpressionFenced) CloseString() string {
	rv := objc.Send[string](a_.ID, objc.Sel("closeString"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXMathExpressionFenced/expressions

func (a_ AXMathExpressionFenced) Expressions() []AXMathExpression {
	rv := objc.Send[[]AXMathExpression](a_.ID, objc.Sel("expressions"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXMathExpressionFenced/openString

func (a_ AXMathExpressionFenced) OpenString() string {
	rv := objc.Send[string](a_.ID, objc.Sel("openString"))
	return rv
}


