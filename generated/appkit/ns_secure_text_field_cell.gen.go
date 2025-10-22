// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [SecureTextFieldCell] class.
var (
	SecureTextFieldCellClass     _SecureTextFieldCellClass
	SecureTextFieldCellClassOnce sync.Once
)

func getSecureTextFieldCellClass() _SecureTextFieldCellClass {
	SecureTextFieldCellClassOnce.Do(func() {
		SecureTextFieldCellClass = _SecureTextFieldCellClass{objc.GetClass("NSSecureTextFieldCell")}
	})
	return SecureTextFieldCellClass
}

type _SecureTextFieldCellClass struct {
	class objc.Class
}

// An interface definition for the [SecureTextFieldCell] class.
type ISecureTextFieldCell interface {
	ITextFieldCell
	EchosBullets() bool
	SetEchosBullets(value bool)
}

// A text field whose value is hidden from the user.
//
// works with and overrides the general cell use of the field editor to provide its own field editor, which doesn’t display text or allow the user to cut or copy its value.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSecureTextFieldCell
type SecureTextFieldCell struct {
	TextFieldCell
}

// SecureTextFieldCellFrom constructs a [SecureTextFieldCell] from an unsafe.Pointer.
//
// A text field whose value is hidden from the user.
func SecureTextFieldCellFrom(ptr unsafe.Pointer) SecureTextFieldCell {
	return SecureTextFieldCell{
		TextFieldCell: TextFieldCellFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (sc _SecureTextFieldCellClass) Alloc() SecureTextFieldCell {
	rv := objc.Send[SecureTextFieldCell](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (sc _SecureTextFieldCellClass) New() SecureTextFieldCell {
	rv := objc.Send[SecureTextFieldCell](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SecureTextFieldCell) Init() SecureTextFieldCell {
	rv := objc.Send[SecureTextFieldCell](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SecureTextFieldCell) Autorelease() SecureTextFieldCell {
	rv := objc.Send[SecureTextFieldCell](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSecureTextFieldCell creates a new SecureTextFieldCell instance.
func NewSecureTextFieldCell() SecureTextFieldCell {
	return getSecureTextFieldCellClass().New()
}


// A Boolean that indicates whether the receiver echoes a bullet character rather than each character typed.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssecuretextfieldcell/echosbullets
func (s_ SecureTextFieldCell) EchosBullets() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("echosBullets"))
	return rv
}


// SetEchosBullets sets the value of the echosBullets property.
// A Boolean that indicates whether the receiver echoes a bullet character rather than each character typed.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssecuretextfieldcell/echosbullets
func (s_ SecureTextFieldCell) SetEchosBullets(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setEchosBullets:"), value)
}



