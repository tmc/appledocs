// Code generated from Apple documentation for Accessibility. DO NOT EDIT.

package accessibility

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [AXMathExpressionMultiscript] class.
var (
	AXMathExpressionMultiscriptClass     _AXMathExpressionMultiscriptClass
	AXMathExpressionMultiscriptClassOnce sync.Once
)

func getAXMathExpressionMultiscriptClass() _AXMathExpressionMultiscriptClass {
	AXMathExpressionMultiscriptClassOnce.Do(func() {
		AXMathExpressionMultiscriptClass = _AXMathExpressionMultiscriptClass{objc.GetClass("AXMathExpressionMultiscript")}
	})
	return AXMathExpressionMultiscriptClass
}

type _AXMathExpressionMultiscriptClass struct {
	class objc.Class
}

// An interface definition for the [AXMathExpressionMultiscript] class.
type IAXMathExpressionMultiscript interface {
	IAXMathExpression
	// properties:
	BaseExpression() IAXMathExpression
	PostscriptExpressions() []IAXMathExpressionSubSuperscript
	PrescriptExpressions() []IAXMathExpressionSubSuperscript
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXMathExpressionMultiscript
type AXMathExpressionMultiscript struct {
	AXMathExpression
}

// AXMathExpressionMultiscriptFrom constructs a [AXMathExpressionMultiscript] from an unsafe.Pointer.
func AXMathExpressionMultiscriptFrom(ptr unsafe.Pointer) AXMathExpressionMultiscript {
	return AXMathExpressionMultiscript{
		AXMathExpression: AXMathExpressionFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ac _AXMathExpressionMultiscriptClass) Alloc() AXMathExpressionMultiscript {
	rv := objc.Send[AXMathExpressionMultiscript](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _AXMathExpressionMultiscriptClass) New() AXMathExpressionMultiscript {
	rv := objc.Send[AXMathExpressionMultiscript](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AXMathExpressionMultiscript) Init() AXMathExpressionMultiscript {
	rv := objc.Send[AXMathExpressionMultiscript](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AXMathExpressionMultiscript) Autorelease() AXMathExpressionMultiscript {
	rv := objc.Send[AXMathExpressionMultiscript](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAXMathExpressionMultiscript creates a new AXMathExpressionMultiscript instance.
func NewAXMathExpressionMultiscript() AXMathExpressionMultiscript {
	return getAXMathExpressionMultiscriptClass().New()
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXMathExpressionMultiscript/init(baseExpression:prescriptExpressions:postscriptExpressions:)
func NewAXMathExpressionMultiscriptWithBaseExpressionPrescriptExpressionsPostscriptExpressions(baseExpression IAXMathExpression, prescriptExpressions []IAXMathExpressionSubSuperscript, postscriptExpressions []IAXMathExpressionSubSuperscript) AXMathExpressionMultiscript {
	instance := getAXMathExpressionMultiscriptClass().Alloc()
	rv := objc.Send[AXMathExpressionMultiscript](instance.ID, objc.Sel("initWithBaseExpression:prescriptExpressions:postscriptExpressions:"), baseExpression, prescriptExpressions, postscriptExpressions)
	rv.Autorelease()
	return rv
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXMathExpressionMultiscript/baseExpression
func (a_ AXMathExpressionMultiscript) BaseExpression() IAXMathExpression {
	rv := objc.Send[AXMathExpression](a_.ID, objc.Sel("baseExpression"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXMathExpressionMultiscript/postscriptExpressions
func (a_ AXMathExpressionMultiscript) PostscriptExpressions() []IAXMathExpressionSubSuperscript {
	rv := objc.Send[[]AXMathExpressionSubSuperscript](a_.ID, objc.Sel("postscriptExpressions"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXMathExpressionMultiscript/prescriptExpressions
func (a_ AXMathExpressionMultiscript) PrescriptExpressions() []IAXMathExpressionSubSuperscript {
	rv := objc.Send[[]AXMathExpressionSubSuperscript](a_.ID, objc.Sel("prescriptExpressions"))
	return rv
}


