
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [OutlineView] class.
var OutlineViewClass _OutlineViewClass

func init() {
	OutlineViewClass = _OutlineViewClass{objc.GetClass("NSOutlineView")}
}

type _OutlineViewClass struct {
	objc.Class
}

// An interface definition for the [OutlineView] class.
type IOutlineView interface {
	ID() objc.ID
}

type OutlineView struct {
	id objc.ID
}

func OutlineViewFrom(ptr unsafe.Pointer) OutlineView {
	return OutlineView{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (o_ OutlineView) ID() objc.ID {
	return o_.id
}

// Alloc allocates a new instance without initialization.
func (oc _OutlineViewClass) Alloc() OutlineView {
	rv := objc.Send[OutlineView](objc.ID(oc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (oc _OutlineViewClass) New() OutlineView {
	rv := objc.Send[OutlineView](objc.ID(oc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewOutlineView creates and returns a new initialized instance.
func NewOutlineView() OutlineView {
	return OutlineViewClass.New()
}

// Init initializes the instance.
func (o_ OutlineView) Init() OutlineView {
	rv := objc.Send[OutlineView](o_.ID(), selInit)
	return rv
}
