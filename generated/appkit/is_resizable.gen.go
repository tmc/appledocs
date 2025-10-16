
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [isResizable] class.
var isResizableClass _isResizableClass

func init() {
	isResizableClass = _isResizableClass{objc.GetClass("isResizable")}
}

type _isResizableClass struct {
	objc.Class
}

// An interface definition for the [isResizable] class.
type IisResizable interface {
	ID() objc.ID
}

type isResizable struct {
	id objc.ID
}

func isResizableFrom(ptr unsafe.Pointer) isResizable {
	return isResizable{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (i_ isResizable) ID() objc.ID {
	return i_.id
}

// Alloc allocates a new instance without initialization.
func (ic _isResizableClass) Alloc() isResizable {
	rv := objc.Send[isResizable](objc.ID(ic.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (ic _isResizableClass) New() isResizable {
	rv := objc.Send[isResizable](objc.ID(ic.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewisResizable creates and returns a new initialized instance.
func NewisResizable() isResizable {
	return isResizableClass.New()
}

// Init initializes the instance.
func (i_ isResizable) Init() isResizable {
	rv := objc.Send[isResizable](i_.ID(), selInit)
	return rv
}
