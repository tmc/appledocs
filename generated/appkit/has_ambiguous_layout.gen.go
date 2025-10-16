
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [hasAmbiguousLayout] class.
var hasAmbiguousLayoutClass _hasAmbiguousLayoutClass

func init() {
	hasAmbiguousLayoutClass = _hasAmbiguousLayoutClass{objc.GetClass("hasAmbiguousLayout")}
}

type _hasAmbiguousLayoutClass struct {
	objc.Class
}

// An interface definition for the [hasAmbiguousLayout] class.
type IhasAmbiguousLayout interface {
	ID() objc.ID
}

type hasAmbiguousLayout struct {
	id objc.ID
}

func hasAmbiguousLayoutFrom(ptr unsafe.Pointer) hasAmbiguousLayout {
	return hasAmbiguousLayout{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (h_ hasAmbiguousLayout) ID() objc.ID {
	return h_.id
}

// Alloc allocates a new instance without initialization.
func (hc _hasAmbiguousLayoutClass) Alloc() hasAmbiguousLayout {
	rv := objc.Send[hasAmbiguousLayout](objc.ID(hc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (hc _hasAmbiguousLayoutClass) New() hasAmbiguousLayout {
	rv := objc.Send[hasAmbiguousLayout](objc.ID(hc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewhasAmbiguousLayout creates and returns a new initialized instance.
func NewhasAmbiguousLayout() hasAmbiguousLayout {
	return hasAmbiguousLayoutClass.New()
}

// Init initializes the instance.
func (h_ hasAmbiguousLayout) Init() hasAmbiguousLayout {
	rv := objc.Send[hasAmbiguousLayout](h_.ID(), selInit)
	return rv
}
