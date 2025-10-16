
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [heightAnchor] class.
var heightAnchorClass _heightAnchorClass

func init() {
	heightAnchorClass = _heightAnchorClass{objc.GetClass("heightAnchor")}
}

type _heightAnchorClass struct {
	objc.Class
}

// An interface definition for the [heightAnchor] class.
type IheightAnchor interface {
	ID() objc.ID
}

type heightAnchor struct {
	id objc.ID
}

func heightAnchorFrom(ptr unsafe.Pointer) heightAnchor {
	return heightAnchor{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (h_ heightAnchor) ID() objc.ID {
	return h_.id
}

// Alloc allocates a new instance without initialization.
func (hc _heightAnchorClass) Alloc() heightAnchor {
	rv := objc.Send[heightAnchor](objc.ID(hc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (hc _heightAnchorClass) New() heightAnchor {
	rv := objc.Send[heightAnchor](objc.ID(hc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewheightAnchor creates and returns a new initialized instance.
func NewheightAnchor() heightAnchor {
	return heightAnchorClass.New()
}

// Init initializes the instance.
func (h_ heightAnchor) Init() heightAnchor {
	rv := objc.Send[heightAnchor](h_.ID(), selInit)
	return rv
}
