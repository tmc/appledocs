
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [runModal] class.
var runModalClass _runModalClass

func init() {
	runModalClass = _runModalClass{objc.GetClass("runModal")}
}

type _runModalClass struct {
	objc.Class
}

// An interface definition for the [runModal] class.
type IrunModal interface {
	ID() objc.ID
}

type runModal struct {
	id objc.ID
}

func runModalFrom(ptr unsafe.Pointer) runModal {
	return runModal{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (r_ runModal) ID() objc.ID {
	return r_.id
}

// Alloc allocates a new instance without initialization.
func (rc _runModalClass) Alloc() runModal {
	rv := objc.Send[runModal](objc.ID(rc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (rc _runModalClass) New() runModal {
	rv := objc.Send[runModal](objc.ID(rc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewrunModal creates and returns a new initialized instance.
func NewrunModal() runModal {
	return runModalClass.New()
}

// Init initializes the instance.
func (r_ runModal) Init() runModal {
	rv := objc.Send[runModal](r_.ID(), selInit)
	return rv
}
