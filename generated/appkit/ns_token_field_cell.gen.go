// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)





// The class instance for the [TokenFieldCell] class.
var (
	TokenFieldCellClass     _TokenFieldCellClass
	TokenFieldCellClassOnce sync.Once
)

func getTokenFieldCellClass() _TokenFieldCellClass {
	TokenFieldCellClassOnce.Do(func() {
		TokenFieldCellClass = _TokenFieldCellClass{objc.GetClass("NSTokenFieldCell")}
	})
	return TokenFieldCellClass
}

type _TokenFieldCellClass struct {
	class objc.Class
}





// An interface definition for the [TokenFieldCell] class.
type ITokenFieldCell interface {
	ITextFieldCell
	

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
func (tc _TokenFieldCellClass) Alloc() TokenFieldCell {
	rv := objc.Send[TokenFieldCell](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (tc _TokenFieldCellClass) New() TokenFieldCell {
	rv := objc.Send[TokenFieldCell](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ TokenFieldCell) Init() TokenFieldCell {
	rv := objc.Send[TokenFieldCell](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ TokenFieldCell) Autorelease() TokenFieldCell {
	rv := objc.Send[TokenFieldCell](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewTokenFieldCell creates a new TokenFieldCell instance.
func NewTokenFieldCell() TokenFieldCell {
	return getTokenFieldCellClass().New()
}





// A text field cell subclass that enables tokenized editing of an array of objects.
//
// is a subclass of that provides tokenized editing of an array of objects similar to the address field in the Mail app. The objects may be strings or objects that can be represented as strings. A single token field cell can be presented in an control.


// A text field cell subclass that enables tokenized editing of an array of objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTokenFieldCell
type TokenFieldCell struct {
	TextFieldCell
}

// TokenFieldCellFrom constructs a [TokenFieldCell] from an unsafe.Pointer.
//
// A text field cell subclass that enables tokenized editing of an array of objects.
func TokenFieldCellFrom(ptr unsafe.Pointer) TokenFieldCell {
	return TokenFieldCell{
		TextFieldCell: TextFieldCellFrom(ptr),
	}
}















// Returns the default completion delay.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTokenFieldCell/defaultCompletionDelay
func (tc _TokenFieldCellClass) DefaultCompletionDelay() float64 {
	rv := objc.Send[float64](objc.ID(tc.class), objc.Sel("defaultCompletionDelay"))
	return rv
}

// Returns the default tokenizing character set.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTokenFieldCell/defaultTokenizingCharacterSet
func (tc _TokenFieldCellClass) DefaultTokenizingCharacterSet() foundation.CharacterSet {
	rv := objc.Send[foundation.CharacterSet](objc.ID(tc.class), objc.Sel("defaultTokenizingCharacterSet"))
	return rv
}











// The receiver’s completion delay to a given delay.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTokenFieldCell/completionDelay
func (t_ TokenFieldCell) CompletionDelay() float64 {
	rv := objc.Send[float64](t_.ID, objc.Sel("completionDelay"))
	return rv
}


// The receiver’s completion delay to a given delay.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTokenFieldCell/completionDelay
func (t_ TokenFieldCell) SetCompletionDelay(value float64) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setCompletionDelay:"), value)
}


// Returns the default completion delay.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTokenFieldCell/defaultCompletionDelay
func (t_ TokenFieldCell) DefaultCompletionDelay() float64 {
	rv := objc.Send[float64](t_.ID, objc.Sel("defaultCompletionDelay"))
	return rv
}


// Returns the default tokenizing character set.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTokenFieldCell/defaultTokenizingCharacterSet
func (t_ TokenFieldCell) DefaultTokenizingCharacterSet() foundation.CharacterSet {
	rv := objc.Send[foundation.CharacterSet](t_.ID, objc.Sel("defaultTokenizingCharacterSet"))
	return rv
}


// The token style of the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTokenFieldCell/tokenStyle
func (t_ TokenFieldCell) TokenStyle() TokenStyle {
	rv := objc.Send[TokenStyle](t_.ID, objc.Sel("tokenStyle"))
	return rv
}


// The token style of the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTokenFieldCell/tokenStyle
func (t_ TokenFieldCell) SetTokenStyle(value TokenStyle) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setTokenStyle:"), value)
}


// The receiver’s tokenizing character set to a given character set.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTokenFieldCell/tokenizingCharacterSet
func (t_ TokenFieldCell) TokenizingCharacterSet() foundation.CharacterSet {
	rv := objc.Send[foundation.CharacterSet](t_.ID, objc.Sel("tokenizingCharacterSet"))
	return rv
}


// The receiver’s tokenizing character set to a given character set.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTokenFieldCell/tokenizingCharacterSet
func (t_ TokenFieldCell) SetTokenizingCharacterSet(value foundation.CharacterSet) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setTokenizingCharacterSet:"), value)
}








