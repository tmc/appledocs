
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [leadingAnchor] class.
var leadingAnchorClass _leadingAnchorClass

func init() {
	leadingAnchorClass = _leadingAnchorClass{objc.GetClass("leadingAnchor")}
}

type _leadingAnchorClass struct {
	objc.Class
}

// An interface definition for the [leadingAnchor] class.
type IleadingAnchor interface {
	ID() objc.ID
}

type leadingAnchor struct {
	id objc.ID
}

func leadingAnchorFrom(ptr unsafe.Pointer) leadingAnchor {
	return leadingAnchor{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (l_ leadingAnchor) ID() objc.ID {
	return l_.id
}

// Alloc allocates a new instance without initialization.
func (lc _leadingAnchorClass) Alloc() leadingAnchor {
	rv := objc.Send[leadingAnchor](objc.ID(lc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (lc _leadingAnchorClass) New() leadingAnchor {
	rv := objc.Send[leadingAnchor](objc.ID(lc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewleadingAnchor creates and returns a new initialized instance.
func NewleadingAnchor() leadingAnchor {
	return leadingAnchorClass.New()
}

// Init initializes the instance.
func (l_ leadingAnchor) Init() leadingAnchor {
	rv := objc.Send[leadingAnchor](l_.ID(), selInit)
	return rv
}
