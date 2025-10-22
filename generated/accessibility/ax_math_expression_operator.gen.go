// Code generated from Apple documentation for Accessibility. DO NOT EDIT.

package accessibility

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [AXMathExpressionOperator] class.
var (
	AXMathExpressionOperatorClass     _AXMathExpressionOperatorClass
	AXMathExpressionOperatorClassOnce sync.Once
)

func getAXMathExpressionOperatorClass() _AXMathExpressionOperatorClass {
	AXMathExpressionOperatorClassOnce.Do(func() {
		AXMathExpressionOperatorClass = _AXMathExpressionOperatorClass{objc.GetClass("AXMathExpressionOperator")}
	})
	return AXMathExpressionOperatorClass
}

type _AXMathExpressionOperatorClass struct {
	class objc.Class
}

// An interface definition for the [AXMathExpressionOperator] class.
type IAXMathExpressionOperator interface {
	IAXMathExpression
	Content() string
}

//
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXMathExpressionOperator
type AXMathExpressionOperator struct {
	AXMathExpression
}

// AXMathExpressionOperatorFrom constructs a [AXMathExpressionOperator] from an unsafe.Pointer.
func AXMathExpressionOperatorFrom(ptr unsafe.Pointer) AXMathExpressionOperator {
	return AXMathExpressionOperator{
		AXMathExpression: AXMathExpressionFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ac _AXMathExpressionOperatorClass) Alloc() AXMathExpressionOperator {
	rv := objc.Send[AXMathExpressionOperator](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _AXMathExpressionOperatorClass) New() AXMathExpressionOperator {
	rv := objc.Send[AXMathExpressionOperator](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AXMathExpressionOperator) Init() AXMathExpressionOperator {
	rv := objc.Send[AXMathExpressionOperator](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AXMathExpressionOperator) Autorelease() AXMathExpressionOperator {
	rv := objc.Send[AXMathExpressionOperator](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAXMathExpressionOperator creates a new AXMathExpressionOperator instance.
func NewAXMathExpressionOperator() AXMathExpressionOperator {
	return getAXMathExpressionOperatorClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXMathExpressionOperator/init(content:)
func NewAXMathExpressionOperatorWithContent(content string) AXMathExpressionOperator {
	instance := getAXMathExpressionOperatorClass().Alloc()
	rv := objc.Send[AXMathExpressionOperator](instance.ID, objc.Sel("initWithContent:"), objc.String(content))
	rv.Autorelease()
	return rv
}


//
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXMathExpressionOperator/content
func (a_ AXMathExpressionOperator) Content() string {
	rv := objc.Send[string](a_.ID, objc.Sel("content"))
	return rv
}


