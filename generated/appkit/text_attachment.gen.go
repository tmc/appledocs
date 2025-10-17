
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [TextAttachment] class.
var TextAttachmentClass _TextAttachmentClass

func init() {
	TextAttachmentClass = _TextAttachmentClass{objc.GetClass("NSTextAttachment")}
}

type _TextAttachmentClass struct {
	objc.Class
}

// An interface definition for the [TextAttachment] class.
type ITextAttachment interface {
	ID() objc.ID
}

type TextAttachment struct {
	id objc.ID
}

func TextAttachmentFrom(ptr unsafe.Pointer) TextAttachment {
	return TextAttachment{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (t_ TextAttachment) ID() objc.ID {
	return t_.id
}

// Alloc allocates a new instance without initialization.
func (tc _TextAttachmentClass) Alloc() TextAttachment {
	rv := objc.Send[TextAttachment](objc.ID(tc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (tc _TextAttachmentClass) New() TextAttachment {
	rv := objc.Send[TextAttachment](objc.ID(tc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewTextAttachment creates and returns a new initialized instance.
func NewTextAttachment() TextAttachment {
	return TextAttachmentClass.New()
}

// Init initializes the instance.
func (t_ TextAttachment) Init() TextAttachment {
	rv := objc.Send[TextAttachment](t_.ID(), selInit)
	return rv
}
