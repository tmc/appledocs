
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [backgroundFilters] class.
var backgroundFiltersClass _backgroundFiltersClass

func init() {
	backgroundFiltersClass = _backgroundFiltersClass{objc.GetClass("backgroundFilters")}
}

type _backgroundFiltersClass struct {
	objc.Class
}

// An interface definition for the [backgroundFilters] class.
type IbackgroundFilters interface {
	ID() objc.ID
}

type backgroundFilters struct {
	id objc.ID
}

func backgroundFiltersFrom(ptr unsafe.Pointer) backgroundFilters {
	return backgroundFilters{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (b_ backgroundFilters) ID() objc.ID {
	return b_.id
}

// Alloc allocates a new instance without initialization.
func (bc _backgroundFiltersClass) Alloc() backgroundFilters {
	rv := objc.Send[backgroundFilters](objc.ID(bc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (bc _backgroundFiltersClass) New() backgroundFilters {
	rv := objc.Send[backgroundFilters](objc.ID(bc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewbackgroundFilters creates and returns a new initialized instance.
func NewbackgroundFilters() backgroundFilters {
	return backgroundFiltersClass.New()
}

// Init initializes the instance.
func (b_ backgroundFilters) Init() backgroundFilters {
	rv := objc.Send[backgroundFilters](b_.ID(), selInit)
	return rv
}
