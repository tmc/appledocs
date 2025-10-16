
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [SecureTextField] class.
var SecureTextFieldClass _SecureTextFieldClass

func init() {
	SecureTextFieldClass = _SecureTextFieldClass{objc.GetClass("NSSecureTextField")}
}

type _SecureTextFieldClass struct {
	objc.Class
}

// An interface definition for the [SecureTextField] class.
type ISecureTextField interface {
	ID() objc.ID
}

type SecureTextField struct {
	id objc.ID
}

func SecureTextFieldFrom(ptr unsafe.Pointer) SecureTextField {
	return SecureTextField{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (s_ SecureTextField) ID() objc.ID {
	return s_.id
}

// Alloc allocates a new instance without initialization.
func (sc _SecureTextFieldClass) Alloc() SecureTextField {
	rv := objc.Send[SecureTextField](objc.ID(sc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (sc _SecureTextFieldClass) New() SecureTextField {
	rv := objc.Send[SecureTextField](objc.ID(sc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewSecureTextField creates and returns a new initialized instance.
func NewSecureTextField() SecureTextField {
	return SecureTextFieldClass.New()
}

// Init initializes the instance.
func (s_ SecureTextField) Init() SecureTextField {
	rv := objc.Send[SecureTextField](s_.ID(), selInit)
	return rv
}
