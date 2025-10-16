
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [preventsApplicationTerminationWhenModal] class.
var preventsApplicationTerminationWhenModalClass _preventsApplicationTerminationWhenModalClass

func init() {
	preventsApplicationTerminationWhenModalClass = _preventsApplicationTerminationWhenModalClass{objc.GetClass("preventsApplicationTerminationWhenModal")}
}

type _preventsApplicationTerminationWhenModalClass struct {
	objc.Class
}

// An interface definition for the [preventsApplicationTerminationWhenModal] class.
type IpreventsApplicationTerminationWhenModal interface {
	ID() objc.ID
}

type preventsApplicationTerminationWhenModal struct {
	id objc.ID
}

func preventsApplicationTerminationWhenModalFrom(ptr unsafe.Pointer) preventsApplicationTerminationWhenModal {
	return preventsApplicationTerminationWhenModal{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (p_ preventsApplicationTerminationWhenModal) ID() objc.ID {
	return p_.id
}

// Alloc allocates a new instance without initialization.
func (pc _preventsApplicationTerminationWhenModalClass) Alloc() preventsApplicationTerminationWhenModal {
	rv := objc.Send[preventsApplicationTerminationWhenModal](objc.ID(pc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (pc _preventsApplicationTerminationWhenModalClass) New() preventsApplicationTerminationWhenModal {
	rv := objc.Send[preventsApplicationTerminationWhenModal](objc.ID(pc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewpreventsApplicationTerminationWhenModal creates and returns a new initialized instance.
func NewpreventsApplicationTerminationWhenModal() preventsApplicationTerminationWhenModal {
	return preventsApplicationTerminationWhenModalClass.New()
}

// Init initializes the instance.
func (p_ preventsApplicationTerminationWhenModal) Init() preventsApplicationTerminationWhenModal {
	rv := objc.Send[preventsApplicationTerminationWhenModal](p_.ID(), selInit)
	return rv
}
