// Code generated from Apple documentation for Accessibility. DO NOT EDIT.

package accessibility

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [AXMathExpressionTableCell] class.
var (
	AXMathExpressionTableCellClass     _AXMathExpressionTableCellClass
	AXMathExpressionTableCellClassOnce sync.Once
)

func getAXMathExpressionTableCellClass() _AXMathExpressionTableCellClass {
	AXMathExpressionTableCellClassOnce.Do(func() {
		AXMathExpressionTableCellClass = _AXMathExpressionTableCellClass{objc.GetClass("AXMathExpressionTableCell")}
	})
	return AXMathExpressionTableCellClass
}

type _AXMathExpressionTableCellClass struct {
	class objc.Class
}

// An interface definition for the [AXMathExpressionTableCell] class.
type IAXMathExpressionTableCell interface {
	IAXMathExpression
	// properties:
	Expressions() []AXMathExpression /* primitive/slice/pointer. */
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXMathExpressionTableCell
type AXMathExpressionTableCell struct {
	AXMathExpression
}

// AXMathExpressionTableCellFrom constructs a [AXMathExpressionTableCell] from an unsafe.Pointer.
func AXMathExpressionTableCellFrom(ptr unsafe.Pointer) AXMathExpressionTableCell {
	return AXMathExpressionTableCell{
		AXMathExpression: AXMathExpressionFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ac _AXMathExpressionTableCellClass) Alloc() AXMathExpressionTableCell {
	rv := objc.Send[AXMathExpressionTableCell](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _AXMathExpressionTableCellClass) New() AXMathExpressionTableCell {
	rv := objc.Send[AXMathExpressionTableCell](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AXMathExpressionTableCell) Init() AXMathExpressionTableCell {
	rv := objc.Send[AXMathExpressionTableCell](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AXMathExpressionTableCell) Autorelease() AXMathExpressionTableCell {
	rv := objc.Send[AXMathExpressionTableCell](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAXMathExpressionTableCell creates a new AXMathExpressionTableCell instance.
func NewAXMathExpressionTableCell() AXMathExpressionTableCell {
	return getAXMathExpressionTableCellClass().New()
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXMathExpressionTableCell/init(expressions:)
func NewAXMathExpressionTableCellWithExpressions(expressions []AXMathExpression /* primitive/slice/pointer. */) AXMathExpressionTableCell {
	instance := getAXMathExpressionTableCellClass().Alloc()
	rv := objc.Send[AXMathExpressionTableCell](instance.ID, objc.Sel("initWithExpressions:"), expressions)
	rv.Autorelease()
	return rv
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXMathExpressionTableCell/expressions
func (a_ AXMathExpressionTableCell) Expressions() []AXMathExpression /* primitive/slice/pointer. */ {
	rv := objc.Send[[]AXMathExpression](a_.ID, objc.Sel("expressions"))
	return rv
}


