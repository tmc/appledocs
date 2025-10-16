
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [TextParagraph] class.
var TextParagraphClass _TextParagraphClass

func init() {
	TextParagraphClass = _TextParagraphClass{objc.GetClass("NSTextParagraph")}
}

type _TextParagraphClass struct {
	objc.Class
}

// An interface definition for the [TextParagraph] class.
type ITextParagraph interface {
	ID() objc.ID
}

type TextParagraph struct {
	id objc.ID
}

func TextParagraphFrom(ptr unsafe.Pointer) TextParagraph {
	return TextParagraph{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (t_ TextParagraph) ID() objc.ID {
	return t_.id
}

// Alloc allocates a new instance without initialization.
func (tc _TextParagraphClass) Alloc() TextParagraph {
	rv := objc.Send[TextParagraph](objc.ID(tc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (tc _TextParagraphClass) New() TextParagraph {
	rv := objc.Send[TextParagraph](objc.ID(tc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewTextParagraph creates and returns a new initialized instance.
func NewTextParagraph() TextParagraph {
	return TextParagraphClass.New()
}

// Init initializes the instance.
func (t_ TextParagraph) Init() TextParagraph {
	rv := objc.Send[TextParagraph](t_.ID(), selInit)
	return rv
}
