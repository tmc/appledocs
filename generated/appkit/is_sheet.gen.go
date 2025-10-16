
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [isSheet] class.
var isSheetClass _isSheetClass

func init() {
	isSheetClass = _isSheetClass{objc.GetClass("isSheet")}
}

type _isSheetClass struct {
	objc.Class
}

// An interface definition for the [isSheet] class.
type IisSheet interface {
	ID() objc.ID
}

type isSheet struct {
	id objc.ID
}

func isSheetFrom(ptr unsafe.Pointer) isSheet {
	return isSheet{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (i_ isSheet) ID() objc.ID {
	return i_.id
}

// Alloc allocates a new instance without initialization.
func (ic _isSheetClass) Alloc() isSheet {
	rv := objc.Send[isSheet](objc.ID(ic.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (ic _isSheetClass) New() isSheet {
	rv := objc.Send[isSheet](objc.ID(ic.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewisSheet creates and returns a new initialized instance.
func NewisSheet() isSheet {
	return isSheetClass.New()
}

// Init initializes the instance.
func (i_ isSheet) Init() isSheet {
	rv := objc.Send[isSheet](i_.ID(), selInit)
	return rv
}
