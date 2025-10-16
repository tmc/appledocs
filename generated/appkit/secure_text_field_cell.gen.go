
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [SecureTextFieldCell] class.
var SecureTextFieldCellClass _SecureTextFieldCellClass

func init() {
	SecureTextFieldCellClass = _SecureTextFieldCellClass{objc.GetClass("NSSecureTextFieldCell")}
}

type _SecureTextFieldCellClass struct {
	objc.Class
}

// An interface definition for the [SecureTextFieldCell] class.
type ISecureTextFieldCell interface {
	ID() objc.ID
}

type SecureTextFieldCell struct {
	id objc.ID
}

func SecureTextFieldCellFrom(ptr unsafe.Pointer) SecureTextFieldCell {
	return SecureTextFieldCell{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (s_ SecureTextFieldCell) ID() objc.ID {
	return s_.id
}

// Alloc allocates a new instance without initialization.
func (sc _SecureTextFieldCellClass) Alloc() SecureTextFieldCell {
	rv := objc.Send[SecureTextFieldCell](objc.ID(sc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (sc _SecureTextFieldCellClass) New() SecureTextFieldCell {
	rv := objc.Send[SecureTextFieldCell](objc.ID(sc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewSecureTextFieldCell creates and returns a new initialized instance.
func NewSecureTextFieldCell() SecureTextFieldCell {
	return SecureTextFieldCellClass.New()
}

// Init initializes the instance.
func (s_ SecureTextFieldCell) Init() SecureTextFieldCell {
	rv := objc.Send[SecureTextFieldCell](s_.ID(), selInit)
	return rv
}
