
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [hasEqualSpacing] class.
var hasEqualSpacingClass _hasEqualSpacingClass

func init() {
	hasEqualSpacingClass = _hasEqualSpacingClass{objc.GetClass("hasEqualSpacing")}
}

type _hasEqualSpacingClass struct {
	objc.Class
}

// An interface definition for the [hasEqualSpacing] class.
type IhasEqualSpacing interface {
	ID() objc.ID
}

type hasEqualSpacing struct {
	id objc.ID
}

func hasEqualSpacingFrom(ptr unsafe.Pointer) hasEqualSpacing {
	return hasEqualSpacing{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (h_ hasEqualSpacing) ID() objc.ID {
	return h_.id
}

// Alloc allocates a new instance without initialization.
func (hc _hasEqualSpacingClass) Alloc() hasEqualSpacing {
	rv := objc.Send[hasEqualSpacing](objc.ID(hc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (hc _hasEqualSpacingClass) New() hasEqualSpacing {
	rv := objc.Send[hasEqualSpacing](objc.ID(hc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewhasEqualSpacing creates and returns a new initialized instance.
func NewhasEqualSpacing() hasEqualSpacing {
	return hasEqualSpacingClass.New()
}

// Init initializes the instance.
func (h_ hasEqualSpacing) Init() hasEqualSpacing {
	rv := objc.Send[hasEqualSpacing](h_.ID(), selInit)
	return rv
}
