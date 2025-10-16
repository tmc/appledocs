
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [TextTab] class.
var TextTabClass _TextTabClass

func init() {
	TextTabClass = _TextTabClass{objc.GetClass("NSTextTab")}
}

type _TextTabClass struct {
	objc.Class
}

// An interface definition for the [TextTab] class.
type ITextTab interface {
	ID() objc.ID
}

type TextTab struct {
	id objc.ID
}

func TextTabFrom(ptr unsafe.Pointer) TextTab {
	return TextTab{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (t_ TextTab) ID() objc.ID {
	return t_.id
}

// Alloc allocates a new instance without initialization.
func (tc _TextTabClass) Alloc() TextTab {
	rv := objc.Send[TextTab](objc.ID(tc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (tc _TextTabClass) New() TextTab {
	rv := objc.Send[TextTab](objc.ID(tc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewTextTab creates and returns a new initialized instance.
func NewTextTab() TextTab {
	return TextTabClass.New()
}

// Init initializes the instance.
func (t_ TextTab) Init() TextTab {
	rv := objc.Send[TextTab](t_.ID(), selInit)
	return rv
}
