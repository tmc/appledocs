
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [TextRange] class.
var TextRangeClass _TextRangeClass

func init() {
	TextRangeClass = _TextRangeClass{objc.GetClass("NSTextRange")}
}

type _TextRangeClass struct {
	objc.Class
}

// An interface definition for the [TextRange] class.
type ITextRange interface {
	ID() objc.ID
}

type TextRange struct {
	id objc.ID
}

func TextRangeFrom(ptr unsafe.Pointer) TextRange {
	return TextRange{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (t_ TextRange) ID() objc.ID {
	return t_.id
}

// Alloc allocates a new instance without initialization.
func (tc _TextRangeClass) Alloc() TextRange {
	rv := objc.Send[TextRange](objc.ID(tc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (tc _TextRangeClass) New() TextRange {
	rv := objc.Send[TextRange](objc.ID(tc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewTextRange creates and returns a new initialized instance.
func NewTextRange() TextRange {
	return TextRangeClass.New()
}

// Init initializes the instance.
func (t_ TextRange) Init() TextRange {
	rv := objc.Send[TextRange](t_.ID(), selInit)
	return rv
}
