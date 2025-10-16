
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [detachesHiddenViews] class.
var detachesHiddenViewsClass _detachesHiddenViewsClass

func init() {
	detachesHiddenViewsClass = _detachesHiddenViewsClass{objc.GetClass("detachesHiddenViews")}
}

type _detachesHiddenViewsClass struct {
	objc.Class
}

// An interface definition for the [detachesHiddenViews] class.
type IdetachesHiddenViews interface {
	ID() objc.ID
}

type detachesHiddenViews struct {
	id objc.ID
}

func detachesHiddenViewsFrom(ptr unsafe.Pointer) detachesHiddenViews {
	return detachesHiddenViews{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (d_ detachesHiddenViews) ID() objc.ID {
	return d_.id
}

// Alloc allocates a new instance without initialization.
func (dc _detachesHiddenViewsClass) Alloc() detachesHiddenViews {
	rv := objc.Send[detachesHiddenViews](objc.ID(dc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (dc _detachesHiddenViewsClass) New() detachesHiddenViews {
	rv := objc.Send[detachesHiddenViews](objc.ID(dc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewdetachesHiddenViews creates and returns a new initialized instance.
func NewdetachesHiddenViews() detachesHiddenViews {
	return detachesHiddenViewsClass.New()
}

// Init initializes the instance.
func (d_ detachesHiddenViews) Init() detachesHiddenViews {
	rv := objc.Send[detachesHiddenViews](d_.ID(), selInit)
	return rv
}
