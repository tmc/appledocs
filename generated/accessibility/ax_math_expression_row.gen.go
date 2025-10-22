// Code generated from Apple documentation for Accessibility. DO NOT EDIT.

package accessibility

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [AXMathExpressionRow] class.
var (
	AXMathExpressionRowClass     _AXMathExpressionRowClass
	AXMathExpressionRowClassOnce sync.Once
)

func getAXMathExpressionRowClass() _AXMathExpressionRowClass {
	AXMathExpressionRowClassOnce.Do(func() {
		AXMathExpressionRowClass = _AXMathExpressionRowClass{objc.GetClass("AXMathExpressionRow")}
	})
	return AXMathExpressionRowClass
}

type _AXMathExpressionRowClass struct {
	class objc.Class
}

// An interface definition for the [AXMathExpressionRow] class.
type IAXMathExpressionRow interface {
	IAXMathExpression
	Expressions() []AXMathExpression
}

//
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXMathExpressionRow
type AXMathExpressionRow struct {
	AXMathExpression
}

// AXMathExpressionRowFrom constructs a [AXMathExpressionRow] from an unsafe.Pointer.
func AXMathExpressionRowFrom(ptr unsafe.Pointer) AXMathExpressionRow {
	return AXMathExpressionRow{
		AXMathExpression: AXMathExpressionFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ac _AXMathExpressionRowClass) Alloc() AXMathExpressionRow {
	rv := objc.Send[AXMathExpressionRow](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _AXMathExpressionRowClass) New() AXMathExpressionRow {
	rv := objc.Send[AXMathExpressionRow](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AXMathExpressionRow) Init() AXMathExpressionRow {
	rv := objc.Send[AXMathExpressionRow](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AXMathExpressionRow) Autorelease() AXMathExpressionRow {
	rv := objc.Send[AXMathExpressionRow](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAXMathExpressionRow creates a new AXMathExpressionRow instance.
func NewAXMathExpressionRow() AXMathExpressionRow {
	return getAXMathExpressionRowClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXMathExpressionRow/init(expressions:)
func NewAXMathExpressionRowWithExpressions(expressions []AXMathExpression) AXMathExpressionRow {
	instance := getAXMathExpressionRowClass().Alloc()
	rv := objc.Send[AXMathExpressionRow](instance.ID, objc.Sel("initWithExpressions:"), expressions)
	rv.Autorelease()
	return rv
}


//
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXMathExpressionRow/expressions
func (a_ AXMathExpressionRow) Expressions() []AXMathExpression {
	rv := objc.Send[[]AXMathExpression](a_.ID, objc.Sel("expressions"))
	return rv
}


