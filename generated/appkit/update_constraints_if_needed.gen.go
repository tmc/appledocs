
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [updateConstraintsIfNeeded] class.
var updateConstraintsIfNeededClass _updateConstraintsIfNeededClass

func init() {
	updateConstraintsIfNeededClass = _updateConstraintsIfNeededClass{objc.GetClass("updateConstraintsIfNeeded")}
}

type _updateConstraintsIfNeededClass struct {
	objc.Class
}

// An interface definition for the [updateConstraintsIfNeeded] class.
type IupdateConstraintsIfNeeded interface {
	ID() objc.ID
}

type updateConstraintsIfNeeded struct {
	id objc.ID
}

func updateConstraintsIfNeededFrom(ptr unsafe.Pointer) updateConstraintsIfNeeded {
	return updateConstraintsIfNeeded{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (u_ updateConstraintsIfNeeded) ID() objc.ID {
	return u_.id
}

// Alloc allocates a new instance without initialization.
func (uc _updateConstraintsIfNeededClass) Alloc() updateConstraintsIfNeeded {
	rv := objc.Send[updateConstraintsIfNeeded](objc.ID(uc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (uc _updateConstraintsIfNeededClass) New() updateConstraintsIfNeeded {
	rv := objc.Send[updateConstraintsIfNeeded](objc.ID(uc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewupdateConstraintsIfNeeded creates and returns a new initialized instance.
func NewupdateConstraintsIfNeeded() updateConstraintsIfNeeded {
	return updateConstraintsIfNeededClass.New()
}

// Init initializes the instance.
func (u_ updateConstraintsIfNeeded) Init() updateConstraintsIfNeeded {
	rv := objc.Send[updateConstraintsIfNeeded](u_.ID(), selInit)
	return rv
}
