
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [layoutGuides] class.
var layoutGuidesClass _layoutGuidesClass

func init() {
	layoutGuidesClass = _layoutGuidesClass{objc.GetClass("layoutGuides")}
}

type _layoutGuidesClass struct {
	objc.Class
}

// An interface definition for the [layoutGuides] class.
type IlayoutGuides interface {
	ID() objc.ID
}

type layoutGuides struct {
	id objc.ID
}

func layoutGuidesFrom(ptr unsafe.Pointer) layoutGuides {
	return layoutGuides{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (l_ layoutGuides) ID() objc.ID {
	return l_.id
}

// Alloc allocates a new instance without initialization.
func (lc _layoutGuidesClass) Alloc() layoutGuides {
	rv := objc.Send[layoutGuides](objc.ID(lc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (lc _layoutGuidesClass) New() layoutGuides {
	rv := objc.Send[layoutGuides](objc.ID(lc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewlayoutGuides creates and returns a new initialized instance.
func NewlayoutGuides() layoutGuides {
	return layoutGuidesClass.New()
}

// Init initializes the instance.
func (l_ layoutGuides) Init() layoutGuides {
	rv := objc.Send[layoutGuides](l_.ID(), selInit)
	return rv
}
