
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [TextListElement] class.
var TextListElementClass _TextListElementClass

func init() {
	TextListElementClass = _TextListElementClass{objc.GetClass("NSTextListElement")}
}

type _TextListElementClass struct {
	objc.Class
}

// An interface definition for the [TextListElement] class.
type ITextListElement interface {
	ID() objc.ID
}

type TextListElement struct {
	id objc.ID
}

func TextListElementFrom(ptr unsafe.Pointer) TextListElement {
	return TextListElement{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (t_ TextListElement) ID() objc.ID {
	return t_.id
}

// Alloc allocates a new instance without initialization.
func (tc _TextListElementClass) Alloc() TextListElement {
	rv := objc.Send[TextListElement](objc.ID(tc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (tc _TextListElementClass) New() TextListElement {
	rv := objc.Send[TextListElement](objc.ID(tc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewTextListElement creates and returns a new initialized instance.
func NewTextListElement() TextListElement {
	return TextListElementClass.New()
}

// Init initializes the instance.
func (t_ TextListElement) Init() TextListElement {
	rv := objc.Send[TextListElement](t_.ID(), selInit)
	return rv
}
