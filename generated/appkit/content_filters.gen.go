
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [contentFilters] class.
var contentFiltersClass _contentFiltersClass

func init() {
	contentFiltersClass = _contentFiltersClass{objc.GetClass("contentFilters")}
}

type _contentFiltersClass struct {
	objc.Class
}

// An interface definition for the [contentFilters] class.
type IcontentFilters interface {
	ID() objc.ID
}

type contentFilters struct {
	id objc.ID
}

func contentFiltersFrom(ptr unsafe.Pointer) contentFilters {
	return contentFilters{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (c_ contentFilters) ID() objc.ID {
	return c_.id
}

// Alloc allocates a new instance without initialization.
func (cc _contentFiltersClass) Alloc() contentFilters {
	rv := objc.Send[contentFilters](objc.ID(cc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (cc _contentFiltersClass) New() contentFilters {
	rv := objc.Send[contentFilters](objc.ID(cc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewcontentFilters creates and returns a new initialized instance.
func NewcontentFilters() contentFilters {
	return contentFiltersClass.New()
}

// Init initializes the instance.
func (c_ contentFilters) Init() contentFilters {
	rv := objc.Send[contentFilters](c_.ID(), selInit)
	return rv
}
