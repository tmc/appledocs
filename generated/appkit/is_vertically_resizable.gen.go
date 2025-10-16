
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [isVerticallyResizable] class.
var isVerticallyResizableClass _isVerticallyResizableClass

func init() {
	isVerticallyResizableClass = _isVerticallyResizableClass{objc.GetClass("isVerticallyResizable")}
}

type _isVerticallyResizableClass struct {
	objc.Class
}

// An interface definition for the [isVerticallyResizable] class.
type IisVerticallyResizable interface {
	ID() objc.ID
}

type isVerticallyResizable struct {
	id objc.ID
}

func isVerticallyResizableFrom(ptr unsafe.Pointer) isVerticallyResizable {
	return isVerticallyResizable{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (i_ isVerticallyResizable) ID() objc.ID {
	return i_.id
}

// Alloc allocates a new instance without initialization.
func (ic _isVerticallyResizableClass) Alloc() isVerticallyResizable {
	rv := objc.Send[isVerticallyResizable](objc.ID(ic.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (ic _isVerticallyResizableClass) New() isVerticallyResizable {
	rv := objc.Send[isVerticallyResizable](objc.ID(ic.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewisVerticallyResizable creates and returns a new initialized instance.
func NewisVerticallyResizable() isVerticallyResizable {
	return isVerticallyResizableClass.New()
}

// Init initializes the instance.
func (i_ isVerticallyResizable) Init() isVerticallyResizable {
	rv := objc.Send[isVerticallyResizable](i_.ID(), selInit)
	return rv
}
