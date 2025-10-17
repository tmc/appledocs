
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [TextList] class.
var TextListClass _TextListClass

func init() {
	TextListClass = _TextListClass{objc.GetClass("NSTextList")}
}

type _TextListClass struct {
	objc.Class
}

// An interface definition for the [TextList] class.
type ITextList interface {
	ID() objc.ID
}

type TextList struct {
	id objc.ID
}

func TextListFrom(ptr unsafe.Pointer) TextList {
	return TextList{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (t_ TextList) ID() objc.ID {
	return t_.id
}

// Alloc allocates a new instance without initialization.
func (tc _TextListClass) Alloc() TextList {
	rv := objc.Send[TextList](objc.ID(tc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (tc _TextListClass) New() TextList {
	rv := objc.Send[TextList](objc.ID(tc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewTextList creates and returns a new initialized instance.
func NewTextList() TextList {
	return TextListClass.New()
}

// Init initializes the instance.
func (t_ TextList) Init() TextList {
	rv := objc.Send[TextList](t_.ID(), selInit)
	return rv
}
