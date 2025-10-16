
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [isHorizontallyResizable] class.
var isHorizontallyResizableClass _isHorizontallyResizableClass

func init() {
	isHorizontallyResizableClass = _isHorizontallyResizableClass{objc.GetClass("isHorizontallyResizable")}
}

type _isHorizontallyResizableClass struct {
	objc.Class
}

// An interface definition for the [isHorizontallyResizable] class.
type IisHorizontallyResizable interface {
	ID() objc.ID
}

type isHorizontallyResizable struct {
	id objc.ID
}

func isHorizontallyResizableFrom(ptr unsafe.Pointer) isHorizontallyResizable {
	return isHorizontallyResizable{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (i_ isHorizontallyResizable) ID() objc.ID {
	return i_.id
}

// Alloc allocates a new instance without initialization.
func (ic _isHorizontallyResizableClass) Alloc() isHorizontallyResizable {
	rv := objc.Send[isHorizontallyResizable](objc.ID(ic.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (ic _isHorizontallyResizableClass) New() isHorizontallyResizable {
	rv := objc.Send[isHorizontallyResizable](objc.ID(ic.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewisHorizontallyResizable creates and returns a new initialized instance.
func NewisHorizontallyResizable() isHorizontallyResizable {
	return isHorizontallyResizableClass.New()
}

// Init initializes the instance.
func (i_ isHorizontallyResizable) Init() isHorizontallyResizable {
	rv := objc.Send[isHorizontallyResizable](i_.ID(), selInit)
	return rv
}
