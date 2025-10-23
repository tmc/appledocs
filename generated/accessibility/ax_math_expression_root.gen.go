// Code generated from Apple documentation for Accessibility. DO NOT EDIT.

package accessibility

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [AXMathExpressionRoot] class.
var (
	AXMathExpressionRootClass     _AXMathExpressionRootClass
	AXMathExpressionRootClassOnce sync.Once
)

func getAXMathExpressionRootClass() _AXMathExpressionRootClass {
	AXMathExpressionRootClassOnce.Do(func() {
		AXMathExpressionRootClass = _AXMathExpressionRootClass{objc.GetClass("AXMathExpressionRoot")}
	})
	return AXMathExpressionRootClass
}

type _AXMathExpressionRootClass struct {
	class objc.Class
}

// An interface definition for the [AXMathExpressionRoot] class.
type IAXMathExpressionRoot interface {
	IAXMathExpression
	RadicandExpressions() []AXMathExpression
	RootIndexExpression() AXMathExpression
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXMathExpressionRoot
type AXMathExpressionRoot struct {
	AXMathExpression
}

// AXMathExpressionRootFrom constructs a [AXMathExpressionRoot] from an unsafe.Pointer.
func AXMathExpressionRootFrom(ptr unsafe.Pointer) AXMathExpressionRoot {
	return AXMathExpressionRoot{
		AXMathExpression: AXMathExpressionFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ac _AXMathExpressionRootClass) Alloc() AXMathExpressionRoot {
	rv := objc.Send[AXMathExpressionRoot](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _AXMathExpressionRootClass) New() AXMathExpressionRoot {
	rv := objc.Send[AXMathExpressionRoot](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AXMathExpressionRoot) Init() AXMathExpressionRoot {
	rv := objc.Send[AXMathExpressionRoot](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AXMathExpressionRoot) Autorelease() AXMathExpressionRoot {
	rv := objc.Send[AXMathExpressionRoot](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAXMathExpressionRoot creates a new AXMathExpressionRoot instance.
func NewAXMathExpressionRoot() AXMathExpressionRoot {
	return getAXMathExpressionRootClass().New()
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXMathExpressionRoot/init(radicandExpressions:rootIndexExpression:)
func NewAXMathExpressionRootWithRadicandExpressionsRootIndexExpression(radicandExpressions []AXMathExpression, rootIndexExpression IAXMathExpression) AXMathExpressionRoot {
	instance := getAXMathExpressionRootClass().Alloc()
	rv := objc.Send[AXMathExpressionRoot](instance.ID, objc.Sel("initWithRadicandExpressions:rootIndexExpression:"), radicandExpressions, rootIndexExpression)
	rv.Autorelease()
	return rv
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXMathExpressionRoot/radicandExpressions
func (a_ AXMathExpressionRoot) RadicandExpressions() []AXMathExpression {
	rv := objc.Send[[]AXMathExpression](a_.ID, objc.Sel("radicandExpressions"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXMathExpressionRoot/rootIndexExpression
func (a_ AXMathExpressionRoot) RootIndexExpression() AXMathExpression {
	rv := objc.Send[AXMathExpression](a_.ID, objc.Sel("rootIndexExpression"))
	return rv
}


