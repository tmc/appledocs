
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [endPage] class.
var endPageClass _endPageClass

func init() {
	endPageClass = _endPageClass{objc.GetClass("endPage")}
}

type _endPageClass struct {
	objc.Class
}

// An interface definition for the [endPage] class.
type IendPage interface {
	ID() objc.ID
}

type endPage struct {
	id objc.ID
}

func endPageFrom(ptr unsafe.Pointer) endPage {
	return endPage{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (e_ endPage) ID() objc.ID {
	return e_.id
}

// Alloc allocates a new instance without initialization.
func (ec _endPageClass) Alloc() endPage {
	rv := objc.Send[endPage](objc.ID(ec.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (ec _endPageClass) New() endPage {
	rv := objc.Send[endPage](objc.ID(ec.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewendPage creates and returns a new initialized instance.
func NewendPage() endPage {
	return endPageClass.New()
}

// Init initializes the instance.
func (e_ endPage) Init() endPage {
	rv := objc.Send[endPage](e_.ID(), selInit)
	return rv
}
