
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [detachedViews] class.
var detachedViewsClass _detachedViewsClass

func init() {
	detachedViewsClass = _detachedViewsClass{objc.GetClass("detachedViews")}
}

type _detachedViewsClass struct {
	objc.Class
}

// An interface definition for the [detachedViews] class.
type IdetachedViews interface {
	ID() objc.ID
}

type detachedViews struct {
	id objc.ID
}

func detachedViewsFrom(ptr unsafe.Pointer) detachedViews {
	return detachedViews{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (d_ detachedViews) ID() objc.ID {
	return d_.id
}

// Alloc allocates a new instance without initialization.
func (dc _detachedViewsClass) Alloc() detachedViews {
	rv := objc.Send[detachedViews](objc.ID(dc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (dc _detachedViewsClass) New() detachedViews {
	rv := objc.Send[detachedViews](objc.ID(dc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewdetachedViews creates and returns a new initialized instance.
func NewdetachedViews() detachedViews {
	return detachedViewsClass.New()
}

// Init initializes the instance.
func (d_ detachedViews) Init() detachedViews {
	rv := objc.Send[detachedViews](d_.ID(), selInit)
	return rv
}
