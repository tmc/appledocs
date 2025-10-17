
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [TextLayoutManager] class.
var TextLayoutManagerClass _TextLayoutManagerClass

func init() {
	TextLayoutManagerClass = _TextLayoutManagerClass{objc.GetClass("NSTextLayoutManager")}
}

type _TextLayoutManagerClass struct {
	objc.Class
}

// An interface definition for the [TextLayoutManager] class.
type ITextLayoutManager interface {
	ID() objc.ID
}

type TextLayoutManager struct {
	id objc.ID
}

func TextLayoutManagerFrom(ptr unsafe.Pointer) TextLayoutManager {
	return TextLayoutManager{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (t_ TextLayoutManager) ID() objc.ID {
	return t_.id
}

// Alloc allocates a new instance without initialization.
func (tc _TextLayoutManagerClass) Alloc() TextLayoutManager {
	rv := objc.Send[TextLayoutManager](objc.ID(tc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (tc _TextLayoutManagerClass) New() TextLayoutManager {
	rv := objc.Send[TextLayoutManager](objc.ID(tc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewTextLayoutManager creates and returns a new initialized instance.
func NewTextLayoutManager() TextLayoutManager {
	return TextLayoutManagerClass.New()
}

// Init initializes the instance.
func (t_ TextLayoutManager) Init() TextLayoutManager {
	rv := objc.Send[TextLayoutManager](t_.ID(), selInit)
	return rv
}
