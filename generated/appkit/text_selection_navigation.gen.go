
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [TextSelectionNavigation] class.
var TextSelectionNavigationClass _TextSelectionNavigationClass

func init() {
	TextSelectionNavigationClass = _TextSelectionNavigationClass{objc.GetClass("NSTextSelectionNavigation")}
}

type _TextSelectionNavigationClass struct {
	objc.Class
}

// An interface definition for the [TextSelectionNavigation] class.
type ITextSelectionNavigation interface {
	ID() objc.ID
}

type TextSelectionNavigation struct {
	id objc.ID
}

func TextSelectionNavigationFrom(ptr unsafe.Pointer) TextSelectionNavigation {
	return TextSelectionNavigation{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (t_ TextSelectionNavigation) ID() objc.ID {
	return t_.id
}

// Alloc allocates a new instance without initialization.
func (tc _TextSelectionNavigationClass) Alloc() TextSelectionNavigation {
	rv := objc.Send[TextSelectionNavigation](objc.ID(tc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (tc _TextSelectionNavigationClass) New() TextSelectionNavigation {
	rv := objc.Send[TextSelectionNavigation](objc.ID(tc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewTextSelectionNavigation creates and returns a new initialized instance.
func NewTextSelectionNavigation() TextSelectionNavigation {
	return TextSelectionNavigationClass.New()
}

// Init initializes the instance.
func (t_ TextSelectionNavigation) Init() TextSelectionNavigation {
	rv := objc.Send[TextSelectionNavigation](t_.ID(), selInit)
	return rv
}
