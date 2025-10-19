// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [TokenFieldCell] class.
var (
	tokenFieldCellClass     _TokenFieldCellClass
	tokenFieldCellClassOnce sync.Once
)

func getTokenFieldCellClass() _TokenFieldCellClass {
	tokenFieldCellClassOnce.Do(func() {
		tokenFieldCellClass = _TokenFieldCellClass{objc.GetClass("NSTokenFieldCell")}
	})
	return tokenFieldCellClass
}

type _TokenFieldCellClass struct {
	class objc.Class
}

// An interface definition for the [TokenFieldCell] class.
type ITokenFieldCell interface {
	ITextFieldCell
}

// A text field cell subclass that enables tokenized editing of an array of objects. [Full Topic]
//
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

// Alloc allocates a new instance without initialization.
func (tc _TokenFieldCellClass) Alloc() TokenFieldCell {
	rv := objc.Send[TokenFieldCell](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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




