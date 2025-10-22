// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [SecureTextField] class.
var (
	SecureTextFieldClass     _SecureTextFieldClass
	SecureTextFieldClassOnce sync.Once
)

func getSecureTextFieldClass() _SecureTextFieldClass {
	SecureTextFieldClassOnce.Do(func() {
		SecureTextFieldClass = _SecureTextFieldClass{objc.GetClass("NSSecureTextField")}
	})
	return SecureTextFieldClass
}

type _SecureTextFieldClass struct {
	class objc.Class
}

// An interface definition for the [SecureTextField] class.
type ISecureTextField interface {
	ITextField
}

// A text field that hides the typed text.
//
// A secure text field is suitable for use as a password-entry object or for any item in which the text value must be kept secret. uses to implement its user interface.


// A text field that hides the typed text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSecureTextField

type SecureTextField struct {
	TextField
}

// SecureTextFieldFrom constructs a [SecureTextField] from an unsafe.Pointer.
//
// A text field that hides the typed text.
func SecureTextFieldFrom(ptr unsafe.Pointer) SecureTextField {
	return SecureTextField{
		TextField: TextFieldFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (sc _SecureTextFieldClass) Alloc() SecureTextField {
	rv := objc.Send[SecureTextField](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (sc _SecureTextFieldClass) New() SecureTextField {
	rv := objc.Send[SecureTextField](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SecureTextField) Init() SecureTextField {
	rv := objc.Send[SecureTextField](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SecureTextField) Autorelease() SecureTextField {
	rv := objc.Send[SecureTextField](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSecureTextField creates a new SecureTextField instance.
func NewSecureTextField() SecureTextField {
	return getSecureTextFieldClass().New()
}




