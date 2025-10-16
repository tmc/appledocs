
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [hasCloseBox] class.
var hasCloseBoxClass _hasCloseBoxClass

func init() {
	hasCloseBoxClass = _hasCloseBoxClass{objc.GetClass("hasCloseBox")}
}

type _hasCloseBoxClass struct {
	objc.Class
}

// An interface definition for the [hasCloseBox] class.
type IhasCloseBox interface {
	ID() objc.ID
}

type hasCloseBox struct {
	id objc.ID
}

func hasCloseBoxFrom(ptr unsafe.Pointer) hasCloseBox {
	return hasCloseBox{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (h_ hasCloseBox) ID() objc.ID {
	return h_.id
}

// Alloc allocates a new instance without initialization.
func (hc _hasCloseBoxClass) Alloc() hasCloseBox {
	rv := objc.Send[hasCloseBox](objc.ID(hc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (hc _hasCloseBoxClass) New() hasCloseBox {
	rv := objc.Send[hasCloseBox](objc.ID(hc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewhasCloseBox creates and returns a new initialized instance.
func NewhasCloseBox() hasCloseBox {
	return hasCloseBoxClass.New()
}

// Init initializes the instance.
func (h_ hasCloseBox) Init() hasCloseBox {
	rv := objc.Send[hasCloseBox](h_.ID(), selInit)
	return rv
}
