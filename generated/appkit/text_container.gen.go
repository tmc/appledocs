
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [TextContainer] class.
var TextContainerClass _TextContainerClass

func init() {
	TextContainerClass = _TextContainerClass{objc.GetClass("NSTextContainer")}
}

type _TextContainerClass struct {
	objc.Class
}

// An interface definition for the [TextContainer] class.
type ITextContainer interface {
	ID() objc.ID
}

type TextContainer struct {
	id objc.ID
}

func TextContainerFrom(ptr unsafe.Pointer) TextContainer {
	return TextContainer{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (t_ TextContainer) ID() objc.ID {
	return t_.id
}

// Alloc allocates a new instance without initialization.
func (tc _TextContainerClass) Alloc() TextContainer {
	rv := objc.Send[TextContainer](objc.ID(tc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (tc _TextContainerClass) New() TextContainer {
	rv := objc.Send[TextContainer](objc.ID(tc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewTextContainer creates and returns a new initialized instance.
func NewTextContainer() TextContainer {
	return TextContainerClass.New()
}

// Init initializes the instance.
func (t_ TextContainer) Init() TextContainer {
	rv := objc.Send[TextContainer](t_.ID(), selInit)
	return rv
}
