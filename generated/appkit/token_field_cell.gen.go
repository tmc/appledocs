
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [TokenFieldCell] class.
var TokenFieldCellClass _TokenFieldCellClass

func init() {
	TokenFieldCellClass = _TokenFieldCellClass{objc.GetClass("NSTokenFieldCell")}
}

type _TokenFieldCellClass struct {
	objc.Class
}

// An interface definition for the [TokenFieldCell] class.
type ITokenFieldCell interface {
	ID() objc.ID
}

type TokenFieldCell struct {
	id objc.ID
}

func TokenFieldCellFrom(ptr unsafe.Pointer) TokenFieldCell {
	return TokenFieldCell{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (t_ TokenFieldCell) ID() objc.ID {
	return t_.id
}

// Alloc allocates a new instance without initialization.
func (tc _TokenFieldCellClass) Alloc() TokenFieldCell {
	rv := objc.Send[TokenFieldCell](objc.ID(tc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (tc _TokenFieldCellClass) New() TokenFieldCell {
	rv := objc.Send[TokenFieldCell](objc.ID(tc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewTokenFieldCell creates and returns a new initialized instance.
func NewTokenFieldCell() TokenFieldCell {
	return TokenFieldCellClass.New()
}

// Init initializes the instance.
func (t_ TokenFieldCell) Init() TokenFieldCell {
	rv := objc.Send[TokenFieldCell](t_.ID(), selInit)
	return rv
}
