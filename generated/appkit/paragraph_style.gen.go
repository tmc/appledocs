
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [ParagraphStyle] class.
var ParagraphStyleClass _ParagraphStyleClass

func init() {
	ParagraphStyleClass = _ParagraphStyleClass{objc.GetClass("NSParagraphStyle")}
}

type _ParagraphStyleClass struct {
	objc.Class
}

// An interface definition for the [ParagraphStyle] class.
type IParagraphStyle interface {
	ID() objc.ID
}

type ParagraphStyle struct {
	id objc.ID
}

func ParagraphStyleFrom(ptr unsafe.Pointer) ParagraphStyle {
	return ParagraphStyle{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (p_ ParagraphStyle) ID() objc.ID {
	return p_.id
}

// Alloc allocates a new instance without initialization.
func (pc _ParagraphStyleClass) Alloc() ParagraphStyle {
	rv := objc.Send[ParagraphStyle](objc.ID(pc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (pc _ParagraphStyleClass) New() ParagraphStyle {
	rv := objc.Send[ParagraphStyle](objc.ID(pc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewParagraphStyle creates and returns a new initialized instance.
func NewParagraphStyle() ParagraphStyle {
	return ParagraphStyleClass.New()
}

// Init initializes the instance.
func (p_ ParagraphStyle) Init() ParagraphStyle {
	rv := objc.Send[ParagraphStyle](p_.ID(), selInit)
	return rv
}
