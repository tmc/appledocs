
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [alignment] class.
var alignmentClass _alignmentClass

func init() {
	alignmentClass = _alignmentClass{objc.GetClass("alignment")}
}

type _alignmentClass struct {
	objc.Class
}

// An interface definition for the [alignment] class.
type Ialignment interface {
	ID() objc.ID
}

type alignment struct {
	id objc.ID
}

func alignmentFrom(ptr unsafe.Pointer) alignment {
	return alignment{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (a_ alignment) ID() objc.ID {
	return a_.id
}

// Alloc allocates a new instance without initialization.
func (ac _alignmentClass) Alloc() alignment {
	rv := objc.Send[alignment](objc.ID(ac.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (ac _alignmentClass) New() alignment {
	rv := objc.Send[alignment](objc.ID(ac.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// Newalignment creates and returns a new initialized instance.
func Newalignment() alignment {
	return alignmentClass.New()
}

// Init initializes the instance.
func (a_ alignment) Init() alignment {
	rv := objc.Send[alignment](a_.ID(), selInit)
	return rv
}
