
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [isRichText] class.
var isRichTextClass _isRichTextClass

func init() {
	isRichTextClass = _isRichTextClass{objc.GetClass("isRichText")}
}

type _isRichTextClass struct {
	objc.Class
}

// An interface definition for the [isRichText] class.
type IisRichText interface {
	ID() objc.ID
}

type isRichText struct {
	id objc.ID
}

func isRichTextFrom(ptr unsafe.Pointer) isRichText {
	return isRichText{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (i_ isRichText) ID() objc.ID {
	return i_.id
}

// Alloc allocates a new instance without initialization.
func (ic _isRichTextClass) Alloc() isRichText {
	rv := objc.Send[isRichText](objc.ID(ic.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (ic _isRichTextClass) New() isRichText {
	rv := objc.Send[isRichText](objc.ID(ic.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewisRichText creates and returns a new initialized instance.
func NewisRichText() isRichText {
	return isRichTextClass.New()
}

// Init initializes the instance.
func (i_ isRichText) Init() isRichText {
	rv := objc.Send[isRichText](i_.ID(), selInit)
	return rv
}
