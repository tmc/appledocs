// Code generated from Apple documentation for Accessibility. DO NOT EDIT.

package accessibility

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [AXMathExpressionNumber] class.
var (
	AXMathExpressionNumberClass     _AXMathExpressionNumberClass
	AXMathExpressionNumberClassOnce sync.Once
)

func getAXMathExpressionNumberClass() _AXMathExpressionNumberClass {
	AXMathExpressionNumberClassOnce.Do(func() {
		AXMathExpressionNumberClass = _AXMathExpressionNumberClass{objc.GetClass("AXMathExpressionNumber")}
	})
	return AXMathExpressionNumberClass
}

type _AXMathExpressionNumberClass struct {
	class objc.Class
}

// An interface definition for the [AXMathExpressionNumber] class.
type IAXMathExpressionNumber interface {
	IAXMathExpression
	// properties:
	Content() string /* primitive/slice/pointer. */
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXMathExpressionNumber
type AXMathExpressionNumber struct {
	AXMathExpression
}

// AXMathExpressionNumberFrom constructs a [AXMathExpressionNumber] from an unsafe.Pointer.
func AXMathExpressionNumberFrom(ptr unsafe.Pointer) AXMathExpressionNumber {
	return AXMathExpressionNumber{
		AXMathExpression: AXMathExpressionFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ac _AXMathExpressionNumberClass) Alloc() AXMathExpressionNumber {
	rv := objc.Send[AXMathExpressionNumber](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _AXMathExpressionNumberClass) New() AXMathExpressionNumber {
	rv := objc.Send[AXMathExpressionNumber](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AXMathExpressionNumber) Init() AXMathExpressionNumber {
	rv := objc.Send[AXMathExpressionNumber](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AXMathExpressionNumber) Autorelease() AXMathExpressionNumber {
	rv := objc.Send[AXMathExpressionNumber](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAXMathExpressionNumber creates a new AXMathExpressionNumber instance.
func NewAXMathExpressionNumber() AXMathExpressionNumber {
	return getAXMathExpressionNumberClass().New()
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXMathExpressionNumber/init(content:)
func NewAXMathExpressionNumberWithContent(content string /* primitive/slice/pointer. */) AXMathExpressionNumber {
	instance := getAXMathExpressionNumberClass().Alloc()
	rv := objc.Send[AXMathExpressionNumber](instance.ID, objc.Sel("initWithContent:"), objc.String(content))
	rv.Autorelease()
	return rv
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXMathExpressionNumber/content
func (a_ AXMathExpressionNumber) Content() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](a_.ID, objc.Sel("content"))
	return rv
}


