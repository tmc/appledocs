
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [TextAttachmentViewProvider] class.
var TextAttachmentViewProviderClass _TextAttachmentViewProviderClass

func init() {
	TextAttachmentViewProviderClass = _TextAttachmentViewProviderClass{objc.GetClass("NSTextAttachmentViewProvider")}
}

type _TextAttachmentViewProviderClass struct {
	objc.Class
}

// An interface definition for the [TextAttachmentViewProvider] class.
type ITextAttachmentViewProvider interface {
	ID() objc.ID
}

type TextAttachmentViewProvider struct {
	id objc.ID
}

func TextAttachmentViewProviderFrom(ptr unsafe.Pointer) TextAttachmentViewProvider {
	return TextAttachmentViewProvider{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (t_ TextAttachmentViewProvider) ID() objc.ID {
	return t_.id
}

// Alloc allocates a new instance without initialization.
func (tc _TextAttachmentViewProviderClass) Alloc() TextAttachmentViewProvider {
	rv := objc.Send[TextAttachmentViewProvider](objc.ID(tc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (tc _TextAttachmentViewProviderClass) New() TextAttachmentViewProvider {
	rv := objc.Send[TextAttachmentViewProvider](objc.ID(tc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewTextAttachmentViewProvider creates and returns a new initialized instance.
func NewTextAttachmentViewProvider() TextAttachmentViewProvider {
	return TextAttachmentViewProviderClass.New()
}

// Init initializes the instance.
func (t_ TextAttachmentViewProvider) Init() TextAttachmentViewProvider {
	rv := objc.Send[TextAttachmentViewProvider](t_.ID(), selInit)
	return rv
}
