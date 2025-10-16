
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [TextPreview] class.
var TextPreviewClass _TextPreviewClass

func init() {
	TextPreviewClass = _TextPreviewClass{objc.GetClass("NSTextPreview")}
}

type _TextPreviewClass struct {
	objc.Class
}

// An interface definition for the [TextPreview] class.
type ITextPreview interface {
	ID() objc.ID
}

type TextPreview struct {
	id objc.ID
}

func TextPreviewFrom(ptr unsafe.Pointer) TextPreview {
	return TextPreview{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (t_ TextPreview) ID() objc.ID {
	return t_.id
}

// Alloc allocates a new instance without initialization.
func (tc _TextPreviewClass) Alloc() TextPreview {
	rv := objc.Send[TextPreview](objc.ID(tc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (tc _TextPreviewClass) New() TextPreview {
	rv := objc.Send[TextPreview](objc.ID(tc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewTextPreview creates and returns a new initialized instance.
func NewTextPreview() TextPreview {
	return TextPreviewClass.New()
}

// Init initializes the instance.
func (t_ TextPreview) Init() TextPreview {
	rv := objc.Send[TextPreview](t_.ID(), selInit)
	return rv
}
