// Code generated from Apple documentation for Accessibility. DO NOT EDIT.

package accessibility

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [AXMathExpressionIdentifier] class.
var (
	AXMathExpressionIdentifierClass     _AXMathExpressionIdentifierClass
	AXMathExpressionIdentifierClassOnce sync.Once
)

func getAXMathExpressionIdentifierClass() _AXMathExpressionIdentifierClass {
	AXMathExpressionIdentifierClassOnce.Do(func() {
		AXMathExpressionIdentifierClass = _AXMathExpressionIdentifierClass{objc.GetClass("AXMathExpressionIdentifier")}
	})
	return AXMathExpressionIdentifierClass
}

type _AXMathExpressionIdentifierClass struct {
	class objc.Class
}

// An interface definition for the [AXMathExpressionIdentifier] class.
type IAXMathExpressionIdentifier interface {
	IAXMathExpression
	Content() string
	SetContent(value string)
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXMathExpressionIdentifier

type AXMathExpressionIdentifier struct {
	AXMathExpression
}

// AXMathExpressionIdentifierFrom constructs a [AXMathExpressionIdentifier] from an unsafe.Pointer.
func AXMathExpressionIdentifierFrom(ptr unsafe.Pointer) AXMathExpressionIdentifier {
	return AXMathExpressionIdentifier{
		AXMathExpression: AXMathExpressionFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ac _AXMathExpressionIdentifierClass) Alloc() AXMathExpressionIdentifier {
	rv := objc.Send[AXMathExpressionIdentifier](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _AXMathExpressionIdentifierClass) New() AXMathExpressionIdentifier {
	rv := objc.Send[AXMathExpressionIdentifier](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AXMathExpressionIdentifier) Init() AXMathExpressionIdentifier {
	rv := objc.Send[AXMathExpressionIdentifier](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AXMathExpressionIdentifier) Autorelease() AXMathExpressionIdentifier {
	rv := objc.Send[AXMathExpressionIdentifier](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAXMathExpressionIdentifier creates a new AXMathExpressionIdentifier instance.
func NewAXMathExpressionIdentifier() AXMathExpressionIdentifier {
	return getAXMathExpressionIdentifierClass().New()
}




// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXMathExpressionIdentifier/init(content:)

func NewAXMathExpressionIdentifierWithContent(content string) AXMathExpressionIdentifier {
	instance := getAXMathExpressionIdentifierClass().Alloc()
	rv := objc.Send[AXMathExpressionIdentifier](instance.ID, objc.Sel("initWithContent:"), objc.String(content))
	rv.Autorelease()
	return rv
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/accessibility/axmathexpressionidentifier/content

func (a_ AXMathExpressionIdentifier) Content() string {
	rv := objc.Send[string](a_.ID, objc.Sel("content"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/accessibility/axmathexpressionidentifier/content

func (a_ AXMathExpressionIdentifier) SetContent(value string) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setContent:"), objc.String(value))
}


