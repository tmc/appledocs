// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [TokenField] class.
var (
	tokenFieldClass     _TokenFieldClass
	tokenFieldClassOnce sync.Once
)

func getTokenFieldClass() _TokenFieldClass {
	tokenFieldClassOnce.Do(func() {
		tokenFieldClass = _TokenFieldClass{objc.GetClass("NSTokenField")}
	})
	return tokenFieldClass
}

type _TokenFieldClass struct {
	class objc.Class
}

// An interface definition for the [TokenField] class.
type ITokenField interface {
	ITextField
}

// A text field that converts text into visually distinct tokens.
//
// Use a token field when you want typed text to be transformed into “tokens”, which are visually distinct elements in the text field interface. For example, you might use a token field in a mail app to display email addresses for individual users. The distinct appearance of tokens makes them easy for users to distinguish from surrounding text. uses an to implement much of the control’s functionality. provides cover methods for most methods of , which invoke the corresponding cell method.
//
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

// Alloc allocates a new instance without initialization.
func (tc _TokenFieldClass) Alloc() TokenField {
	rv := objc.Send[TokenField](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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




