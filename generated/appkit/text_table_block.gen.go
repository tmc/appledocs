
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [TextTableBlock] class.
var TextTableBlockClass _TextTableBlockClass

func init() {
	TextTableBlockClass = _TextTableBlockClass{objc.GetClass("NSTextTableBlock")}
}

type _TextTableBlockClass struct {
	objc.Class
}

// An interface definition for the [TextTableBlock] class.
type ITextTableBlock interface {
	ID() objc.ID
}

type TextTableBlock struct {
	id objc.ID
}

func TextTableBlockFrom(ptr unsafe.Pointer) TextTableBlock {
	return TextTableBlock{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (t_ TextTableBlock) ID() objc.ID {
	return t_.id
}

// Alloc allocates a new instance without initialization.
func (tc _TextTableBlockClass) Alloc() TextTableBlock {
	rv := objc.Send[TextTableBlock](objc.ID(tc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (tc _TextTableBlockClass) New() TextTableBlock {
	rv := objc.Send[TextTableBlock](objc.ID(tc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewTextTableBlock creates and returns a new initialized instance.
func NewTextTableBlock() TextTableBlock {
	return TextTableBlockClass.New()
}

// Init initializes the instance.
func (t_ TextTableBlock) Init() TextTableBlock {
	rv := objc.Send[TextTableBlock](t_.ID(), selInit)
	return rv
}
