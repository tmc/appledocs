// Code generated from Apple documentation for Accessibility. DO NOT EDIT.

package accessibility

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [AXMathExpressionText] class.
var (
	AXMathExpressionTextClass     _AXMathExpressionTextClass
	AXMathExpressionTextClassOnce sync.Once
)

func getAXMathExpressionTextClass() _AXMathExpressionTextClass {
	AXMathExpressionTextClassOnce.Do(func() {
		AXMathExpressionTextClass = _AXMathExpressionTextClass{objc.GetClass("AXMathExpressionText")}
	})
	return AXMathExpressionTextClass
}

type _AXMathExpressionTextClass struct {
	class objc.Class
}

// An interface definition for the [AXMathExpressionText] class.
type IAXMathExpressionText interface {
	IAXMathExpression
	Content() string
}

//
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXMathExpressionText
type AXMathExpressionText struct {
	AXMathExpression
}

// AXMathExpressionTextFrom constructs a [AXMathExpressionText] from an unsafe.Pointer.
func AXMathExpressionTextFrom(ptr unsafe.Pointer) AXMathExpressionText {
	return AXMathExpressionText{
		AXMathExpression: AXMathExpressionFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ac _AXMathExpressionTextClass) Alloc() AXMathExpressionText {
	rv := objc.Send[AXMathExpressionText](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _AXMathExpressionTextClass) New() AXMathExpressionText {
	rv := objc.Send[AXMathExpressionText](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AXMathExpressionText) Init() AXMathExpressionText {
	rv := objc.Send[AXMathExpressionText](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AXMathExpressionText) Autorelease() AXMathExpressionText {
	rv := objc.Send[AXMathExpressionText](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAXMathExpressionText creates a new AXMathExpressionText instance.
func NewAXMathExpressionText() AXMathExpressionText {
	return getAXMathExpressionTextClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXMathExpressionText/init(content:)
func NewAXMathExpressionTextWithContent(content string) AXMathExpressionText {
	instance := getAXMathExpressionTextClass().Alloc()
	rv := objc.Send[AXMathExpressionText](instance.ID, objc.Sel("initWithContent:"), objc.String(content))
	rv.Autorelease()
	return rv
}


//
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXMathExpressionText/content
func (a_ AXMathExpressionText) Content() string {
	rv := objc.Send[string](a_.ID, objc.Sel("content"))
	return rv
}


