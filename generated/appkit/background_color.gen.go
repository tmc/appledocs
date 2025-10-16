
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [backgroundColor] class.
var backgroundColorClass _backgroundColorClass

func init() {
	backgroundColorClass = _backgroundColorClass{objc.GetClass("backgroundColor")}
}

type _backgroundColorClass struct {
	objc.Class
}

// An interface definition for the [backgroundColor] class.
type IbackgroundColor interface {
	ID() objc.ID
}

type backgroundColor struct {
	id objc.ID
}

func backgroundColorFrom(ptr unsafe.Pointer) backgroundColor {
	return backgroundColor{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (b_ backgroundColor) ID() objc.ID {
	return b_.id
}

// Alloc allocates a new instance without initialization.
func (bc _backgroundColorClass) Alloc() backgroundColor {
	rv := objc.Send[backgroundColor](objc.ID(bc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (bc _backgroundColorClass) New() backgroundColor {
	rv := objc.Send[backgroundColor](objc.ID(bc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewbackgroundColor creates and returns a new initialized instance.
func NewbackgroundColor() backgroundColor {
	return backgroundColorClass.New()
}

// Init initializes the instance.
func (b_ backgroundColor) Init() backgroundColor {
	rv := objc.Send[backgroundColor](b_.ID(), selInit)
	return rv
}
