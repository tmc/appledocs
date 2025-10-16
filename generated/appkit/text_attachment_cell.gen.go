
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [TextAttachmentCell] class.
var TextAttachmentCellClass _TextAttachmentCellClass

func init() {
	TextAttachmentCellClass = _TextAttachmentCellClass{objc.GetClass("NSTextAttachmentCell")}
}

type _TextAttachmentCellClass struct {
	objc.Class
}

// An interface definition for the [TextAttachmentCell] class.
type ITextAttachmentCell interface {
	ID() objc.ID
}

type TextAttachmentCell struct {
	id objc.ID
}

func TextAttachmentCellFrom(ptr unsafe.Pointer) TextAttachmentCell {
	return TextAttachmentCell{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (t_ TextAttachmentCell) ID() objc.ID {
	return t_.id
}

// Alloc allocates a new instance without initialization.
func (tc _TextAttachmentCellClass) Alloc() TextAttachmentCell {
	rv := objc.Send[TextAttachmentCell](objc.ID(tc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (tc _TextAttachmentCellClass) New() TextAttachmentCell {
	rv := objc.Send[TextAttachmentCell](objc.ID(tc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewTextAttachmentCell creates and returns a new initialized instance.
func NewTextAttachmentCell() TextAttachmentCell {
	return TextAttachmentCellClass.New()
}

// Init initializes the instance.
func (t_ TextAttachmentCell) Init() TextAttachmentCell {
	rv := objc.Send[TextAttachmentCell](t_.ID(), selInit)
	return rv
}
