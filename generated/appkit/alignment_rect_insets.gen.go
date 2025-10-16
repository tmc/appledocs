
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [alignmentRectInsets] class.
var alignmentRectInsetsClass _alignmentRectInsetsClass

func init() {
	alignmentRectInsetsClass = _alignmentRectInsetsClass{objc.GetClass("alignmentRectInsets")}
}

type _alignmentRectInsetsClass struct {
	objc.Class
}

// An interface definition for the [alignmentRectInsets] class.
type IalignmentRectInsets interface {
	ID() objc.ID
}

type alignmentRectInsets struct {
	id objc.ID
}

func alignmentRectInsetsFrom(ptr unsafe.Pointer) alignmentRectInsets {
	return alignmentRectInsets{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (a_ alignmentRectInsets) ID() objc.ID {
	return a_.id
}

// Alloc allocates a new instance without initialization.
func (ac _alignmentRectInsetsClass) Alloc() alignmentRectInsets {
	rv := objc.Send[alignmentRectInsets](objc.ID(ac.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (ac _alignmentRectInsetsClass) New() alignmentRectInsets {
	rv := objc.Send[alignmentRectInsets](objc.ID(ac.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewalignmentRectInsets creates and returns a new initialized instance.
func NewalignmentRectInsets() alignmentRectInsets {
	return alignmentRectInsetsClass.New()
}

// Init initializes the instance.
func (a_ alignmentRectInsets) Init() alignmentRectInsets {
	rv := objc.Send[alignmentRectInsets](a_.ID(), selInit)
	return rv
}
