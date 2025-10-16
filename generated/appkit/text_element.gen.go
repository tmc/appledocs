
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [TextElement] class.
var TextElementClass _TextElementClass

func init() {
	TextElementClass = _TextElementClass{objc.GetClass("NSTextElement")}
}

type _TextElementClass struct {
	objc.Class
}

// An interface definition for the [TextElement] class.
type ITextElement interface {
	ID() objc.ID
}

type TextElement struct {
	id objc.ID
}

func TextElementFrom(ptr unsafe.Pointer) TextElement {
	return TextElement{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (t_ TextElement) ID() objc.ID {
	return t_.id
}

// Alloc allocates a new instance without initialization.
func (tc _TextElementClass) Alloc() TextElement {
	rv := objc.Send[TextElement](objc.ID(tc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (tc _TextElementClass) New() TextElement {
	rv := objc.Send[TextElement](objc.ID(tc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewTextElement creates and returns a new initialized instance.
func NewTextElement() TextElement {
	return TextElementClass.New()
}

// Init initializes the instance.
func (t_ TextElement) Init() TextElement {
	rv := objc.Send[TextElement](t_.ID(), selInit)
	return rv
}
