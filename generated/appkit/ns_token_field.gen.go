// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)





// The class instance for the [TokenField] class.
var (
	TokenFieldClass     _TokenFieldClass
	TokenFieldClassOnce sync.Once
)

func getTokenFieldClass() _TokenFieldClass {
	TokenFieldClassOnce.Do(func() {
		TokenFieldClass = _TokenFieldClass{objc.GetClass("NSTokenField")}
	})
	return TokenFieldClass
}

type _TokenFieldClass struct {
	class objc.Class
}





// An interface definition for the [TokenField] class.
type ITokenField interface {
	ITextField
	

	// properties:
	CompletionDelay() float64
	SetCompletionDelay(value float64)
	TokenStyle() TokenStyle
	SetTokenStyle(value TokenStyle)
	TokenizingCharacterSet() foundation.CharacterSet
	SetTokenizingCharacterSet(value foundation.CharacterSet)


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (tc _TokenFieldClass) Alloc() TokenField {
	rv := objc.Send[TokenField](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (tc _TokenFieldClass) New() TokenField {
	rv := objc.Send[TokenField](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ TokenField) Init() TokenField {
	rv := objc.Send[TokenField](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ TokenField) Autorelease() TokenField {
	rv := objc.Send[TokenField](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewTokenField creates a new TokenField instance.
func NewTokenField() TokenField {
	return getTokenFieldClass().New()
}





// A text field that converts text into visually distinct tokens.
//
// Use a token field when you want typed text to be transformed into “tokens”, which are visually distinct elements in the text field interface. For example, you might use a token field in a mail app to display email addresses for individual users. The distinct appearance of tokens makes them easy for users to distinguish from surrounding text. uses an to implement much of the control’s functionality. provides cover methods for most methods of , which invoke the corresponding cell method.


// A text field that converts text into visually distinct tokens.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTokenField
type TokenField struct {
	TextField
}

// TokenFieldFrom constructs a [TokenField] from an unsafe.Pointer.
//
// A text field that converts text into visually distinct tokens.
func TokenFieldFrom(ptr unsafe.Pointer) TokenField {
	return TokenField{
		TextField: TextFieldFrom(ptr),
	}
}















// Returns the default completion delay.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTokenField/defaultCompletionDelay
func (tc _TokenFieldClass) DefaultCompletionDelay() float64 {
	rv := objc.Send[float64](objc.ID(tc.class), objc.Sel("defaultCompletionDelay"))
	return rv
}

// Returns the default tokenizing character set.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTokenField/defaultTokenizingCharacterSet
func (tc _TokenFieldClass) DefaultTokenizingCharacterSet() foundation.CharacterSet {
	rv := objc.Send[foundation.CharacterSet](objc.ID(tc.class), objc.Sel("defaultTokenizingCharacterSet"))
	return rv
}











// The receiver’s completion delay.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTokenField/completionDelay
func (t_ TokenField) CompletionDelay() float64 {
	rv := objc.Send[float64](t_.ID, objc.Sel("completionDelay"))
	return rv
}


// The receiver’s completion delay.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTokenField/completionDelay
func (t_ TokenField) SetCompletionDelay(value float64) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setCompletionDelay:"), value)
}


// Returns the default completion delay.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTokenField/defaultCompletionDelay
func (t_ TokenField) DefaultCompletionDelay() float64 {
	rv := objc.Send[float64](t_.ID, objc.Sel("defaultCompletionDelay"))
	return rv
}


// Returns the default tokenizing character set.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTokenField/defaultTokenizingCharacterSet
func (t_ TokenField) DefaultTokenizingCharacterSet() foundation.CharacterSet {
	rv := objc.Send[foundation.CharacterSet](t_.ID, objc.Sel("defaultTokenizingCharacterSet"))
	return rv
}


// The token style of the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTokenField/tokenStyle-swift.property
func (t_ TokenField) TokenStyle() TokenStyle {
	rv := objc.Send[TokenStyle](t_.ID, objc.Sel("tokenStyle"))
	return rv
}


// The token style of the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTokenField/tokenStyle-swift.property
func (t_ TokenField) SetTokenStyle(value TokenStyle) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setTokenStyle:"), value)
}


// The recevier’s tokenizing character set to .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTokenField/tokenizingCharacterSet
func (t_ TokenField) TokenizingCharacterSet() foundation.CharacterSet {
	rv := objc.Send[foundation.CharacterSet](t_.ID, objc.Sel("tokenizingCharacterSet"))
	return rv
}


// The recevier’s tokenizing character set to .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTokenField/tokenizingCharacterSet
func (t_ TokenField) SetTokenizingCharacterSet(value foundation.CharacterSet) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setTokenizingCharacterSet:"), value)
}








