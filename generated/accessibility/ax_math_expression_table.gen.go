// Code generated from Apple documentation for Accessibility. DO NOT EDIT.

package accessibility

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [AXMathExpressionTable] class.
var (
	AXMathExpressionTableClass     _AXMathExpressionTableClass
	AXMathExpressionTableClassOnce sync.Once
)

func getAXMathExpressionTableClass() _AXMathExpressionTableClass {
	AXMathExpressionTableClassOnce.Do(func() {
		AXMathExpressionTableClass = _AXMathExpressionTableClass{objc.GetClass("AXMathExpressionTable")}
	})
	return AXMathExpressionTableClass
}

type _AXMathExpressionTableClass struct {
	class objc.Class
}

// An interface definition for the [AXMathExpressionTable] class.
type IAXMathExpressionTable interface {
	IAXMathExpression
	Expressions() []AXMathExpression
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXMathExpressionTable

type AXMathExpressionTable struct {
	AXMathExpression
}

// AXMathExpressionTableFrom constructs a [AXMathExpressionTable] from an unsafe.Pointer.
func AXMathExpressionTableFrom(ptr unsafe.Pointer) AXMathExpressionTable {
	return AXMathExpressionTable{
		AXMathExpression: AXMathExpressionFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ac _AXMathExpressionTableClass) Alloc() AXMathExpressionTable {
	rv := objc.Send[AXMathExpressionTable](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _AXMathExpressionTableClass) New() AXMathExpressionTable {
	rv := objc.Send[AXMathExpressionTable](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AXMathExpressionTable) Init() AXMathExpressionTable {
	rv := objc.Send[AXMathExpressionTable](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AXMathExpressionTable) Autorelease() AXMathExpressionTable {
	rv := objc.Send[AXMathExpressionTable](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAXMathExpressionTable creates a new AXMathExpressionTable instance.
func NewAXMathExpressionTable() AXMathExpressionTable {
	return getAXMathExpressionTableClass().New()
}




// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXMathExpressionTable/init(expressions:)

func NewAXMathExpressionTableWithExpressions(expressions []AXMathExpression) AXMathExpressionTable {
	instance := getAXMathExpressionTableClass().Alloc()
	rv := objc.Send[AXMathExpressionTable](instance.ID, objc.Sel("initWithExpressions:"), expressions)
	rv.Autorelease()
	return rv
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXMathExpressionTable/expressions

func (a_ AXMathExpressionTable) Expressions() []AXMathExpression {
	rv := objc.Send[[]AXMathExpression](a_.ID, objc.Sel("expressions"))
	return rv
}


