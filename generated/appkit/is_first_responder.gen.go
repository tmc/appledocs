
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [isFirstResponder] class.
var isFirstResponderClass _isFirstResponderClass

func init() {
	isFirstResponderClass = _isFirstResponderClass{objc.GetClass("isFirstResponder")}
}

type _isFirstResponderClass struct {
	objc.Class
}

// An interface definition for the [isFirstResponder] class.
type IisFirstResponder interface {
	ID() objc.ID
}

type isFirstResponder struct {
	id objc.ID
}

func isFirstResponderFrom(ptr unsafe.Pointer) isFirstResponder {
	return isFirstResponder{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (i_ isFirstResponder) ID() objc.ID {
	return i_.id
}

// Alloc allocates a new instance without initialization.
func (ic _isFirstResponderClass) Alloc() isFirstResponder {
	rv := objc.Send[isFirstResponder](objc.ID(ic.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (ic _isFirstResponderClass) New() isFirstResponder {
	rv := objc.Send[isFirstResponder](objc.ID(ic.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewisFirstResponder creates and returns a new initialized instance.
func NewisFirstResponder() isFirstResponder {
	return isFirstResponderClass.New()
}

// Init initializes the instance.
func (i_ isFirstResponder) Init() isFirstResponder {
	rv := objc.Send[isFirstResponder](i_.ID(), selInit)
	return rv
}
