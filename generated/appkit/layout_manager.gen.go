
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [LayoutManager] class.
var LayoutManagerClass _LayoutManagerClass

func init() {
	LayoutManagerClass = _LayoutManagerClass{objc.GetClass("NSLayoutManager")}
}

type _LayoutManagerClass struct {
	objc.Class
}

// An interface definition for the [LayoutManager] class.
type ILayoutManager interface {
	ID() objc.ID
}

type LayoutManager struct {
	id objc.ID
}

func LayoutManagerFrom(ptr unsafe.Pointer) LayoutManager {
	return LayoutManager{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (l_ LayoutManager) ID() objc.ID {
	return l_.id
}

// Alloc allocates a new instance without initialization.
func (lc _LayoutManagerClass) Alloc() LayoutManager {
	rv := objc.Send[LayoutManager](objc.ID(lc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (lc _LayoutManagerClass) New() LayoutManager {
	rv := objc.Send[LayoutManager](objc.ID(lc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewLayoutManager creates and returns a new initialized instance.
func NewLayoutManager() LayoutManager {
	return LayoutManagerClass.New()
}

// Init initializes the instance.
func (l_ LayoutManager) Init() LayoutManager {
	rv := objc.Send[LayoutManager](l_.ID(), selInit)
	return rv
}
