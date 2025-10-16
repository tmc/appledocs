
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [bottomAnchor] class.
var bottomAnchorClass _bottomAnchorClass

func init() {
	bottomAnchorClass = _bottomAnchorClass{objc.GetClass("bottomAnchor")}
}

type _bottomAnchorClass struct {
	objc.Class
}

// An interface definition for the [bottomAnchor] class.
type IbottomAnchor interface {
	ID() objc.ID
}

type bottomAnchor struct {
	id objc.ID
}

func bottomAnchorFrom(ptr unsafe.Pointer) bottomAnchor {
	return bottomAnchor{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (b_ bottomAnchor) ID() objc.ID {
	return b_.id
}

// Alloc allocates a new instance without initialization.
func (bc _bottomAnchorClass) Alloc() bottomAnchor {
	rv := objc.Send[bottomAnchor](objc.ID(bc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (bc _bottomAnchorClass) New() bottomAnchor {
	rv := objc.Send[bottomAnchor](objc.ID(bc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewbottomAnchor creates and returns a new initialized instance.
func NewbottomAnchor() bottomAnchor {
	return bottomAnchorClass.New()
}

// Init initializes the instance.
func (b_ bottomAnchor) Init() bottomAnchor {
	rv := objc.Send[bottomAnchor](b_.ID(), selInit)
	return rv
}
