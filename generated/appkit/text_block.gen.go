
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [TextBlock] class.
var TextBlockClass _TextBlockClass

func init() {
	TextBlockClass = _TextBlockClass{objc.GetClass("NSTextBlock")}
}

type _TextBlockClass struct {
	objc.Class
}

// An interface definition for the [TextBlock] class.
type ITextBlock interface {
	ID() objc.ID
}

type TextBlock struct {
	id objc.ID
}

func TextBlockFrom(ptr unsafe.Pointer) TextBlock {
	return TextBlock{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (t_ TextBlock) ID() objc.ID {
	return t_.id
}

// Alloc allocates a new instance without initialization.
func (tc _TextBlockClass) Alloc() TextBlock {
	rv := objc.Send[TextBlock](objc.ID(tc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (tc _TextBlockClass) New() TextBlock {
	rv := objc.Send[TextBlock](objc.ID(tc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewTextBlock creates and returns a new initialized instance.
func NewTextBlock() TextBlock {
	return TextBlockClass.New()
}

// Init initializes the instance.
func (t_ TextBlock) Init() TextBlock {
	rv := objc.Send[TextBlock](t_.ID(), selInit)
	return rv
}
