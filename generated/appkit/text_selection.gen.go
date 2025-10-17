
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [TextSelection] class.
var TextSelectionClass _TextSelectionClass

func init() {
	TextSelectionClass = _TextSelectionClass{objc.GetClass("NSTextSelection")}
}

type _TextSelectionClass struct {
	objc.Class
}

// An interface definition for the [TextSelection] class.
type ITextSelection interface {
	ID() objc.ID
}

type TextSelection struct {
	id objc.ID
}

func TextSelectionFrom(ptr unsafe.Pointer) TextSelection {
	return TextSelection{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (t_ TextSelection) ID() objc.ID {
	return t_.id
}

// Alloc allocates a new instance without initialization.
func (tc _TextSelectionClass) Alloc() TextSelection {
	rv := objc.Send[TextSelection](objc.ID(tc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (tc _TextSelectionClass) New() TextSelection {
	rv := objc.Send[TextSelection](objc.ID(tc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewTextSelection creates and returns a new initialized instance.
func NewTextSelection() TextSelection {
	return TextSelectionClass.New()
}

// Init initializes the instance.
func (t_ TextSelection) Init() TextSelection {
	rv := objc.Send[TextSelection](t_.ID(), selInit)
	return rv
}
