// Code generated from Apple documentation for Accessibility. DO NOT EDIT.

package accessibility

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [AXMathExpressionTableRow] class.
var (
	AXMathExpressionTableRowClass     _AXMathExpressionTableRowClass
	AXMathExpressionTableRowClassOnce sync.Once
)

func getAXMathExpressionTableRowClass() _AXMathExpressionTableRowClass {
	AXMathExpressionTableRowClassOnce.Do(func() {
		AXMathExpressionTableRowClass = _AXMathExpressionTableRowClass{objc.GetClass("AXMathExpressionTableRow")}
	})
	return AXMathExpressionTableRowClass
}

type _AXMathExpressionTableRowClass struct {
	class objc.Class
}

// An interface definition for the [AXMathExpressionTableRow] class.
type IAXMathExpressionTableRow interface {
	IAXMathExpression
	// properties:
	Expressions() []IAXMathExpression
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXMathExpressionTableRow
type AXMathExpressionTableRow struct {
	AXMathExpression
}

// AXMathExpressionTableRowFrom constructs a [AXMathExpressionTableRow] from an unsafe.Pointer.
func AXMathExpressionTableRowFrom(ptr unsafe.Pointer) AXMathExpressionTableRow {
	return AXMathExpressionTableRow{
		AXMathExpression: AXMathExpressionFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ac _AXMathExpressionTableRowClass) Alloc() AXMathExpressionTableRow {
	rv := objc.Send[AXMathExpressionTableRow](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _AXMathExpressionTableRowClass) New() AXMathExpressionTableRow {
	rv := objc.Send[AXMathExpressionTableRow](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AXMathExpressionTableRow) Init() AXMathExpressionTableRow {
	rv := objc.Send[AXMathExpressionTableRow](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AXMathExpressionTableRow) Autorelease() AXMathExpressionTableRow {
	rv := objc.Send[AXMathExpressionTableRow](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAXMathExpressionTableRow creates a new AXMathExpressionTableRow instance.
func NewAXMathExpressionTableRow() AXMathExpressionTableRow {
	return getAXMathExpressionTableRowClass().New()
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXMathExpressionTableRow/init(expressions:)
func NewAXMathExpressionTableRowWithExpressions(expressions []IAXMathExpression) AXMathExpressionTableRow {
	instance := getAXMathExpressionTableRowClass().Alloc()
	rv := objc.Send[AXMathExpressionTableRow](instance.ID, objc.Sel("initWithExpressions:"), expressions)
	rv.Autorelease()
	return rv
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXMathExpressionTableRow/expressions
func (a_ AXMathExpressionTableRow) Expressions() []IAXMathExpression {
	rv := objc.Send[[]AXMathExpression](a_.ID, objc.Sel("expressions"))
	return rv
}


