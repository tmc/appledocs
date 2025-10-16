
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [updateConstraintsForSubtreeIfNeeded] class.
var updateConstraintsForSubtreeIfNeededClass _updateConstraintsForSubtreeIfNeededClass

func init() {
	updateConstraintsForSubtreeIfNeededClass = _updateConstraintsForSubtreeIfNeededClass{objc.GetClass("updateConstraintsForSubtreeIfNeeded")}
}

type _updateConstraintsForSubtreeIfNeededClass struct {
	objc.Class
}

// An interface definition for the [updateConstraintsForSubtreeIfNeeded] class.
type IupdateConstraintsForSubtreeIfNeeded interface {
	ID() objc.ID
}

type updateConstraintsForSubtreeIfNeeded struct {
	id objc.ID
}

func updateConstraintsForSubtreeIfNeededFrom(ptr unsafe.Pointer) updateConstraintsForSubtreeIfNeeded {
	return updateConstraintsForSubtreeIfNeeded{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (u_ updateConstraintsForSubtreeIfNeeded) ID() objc.ID {
	return u_.id
}

// Alloc allocates a new instance without initialization.
func (uc _updateConstraintsForSubtreeIfNeededClass) Alloc() updateConstraintsForSubtreeIfNeeded {
	rv := objc.Send[updateConstraintsForSubtreeIfNeeded](objc.ID(uc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (uc _updateConstraintsForSubtreeIfNeededClass) New() updateConstraintsForSubtreeIfNeeded {
	rv := objc.Send[updateConstraintsForSubtreeIfNeeded](objc.ID(uc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewupdateConstraintsForSubtreeIfNeeded creates and returns a new initialized instance.
func NewupdateConstraintsForSubtreeIfNeeded() updateConstraintsForSubtreeIfNeeded {
	return updateConstraintsForSubtreeIfNeededClass.New()
}

// Init initializes the instance.
func (u_ updateConstraintsForSubtreeIfNeeded) Init() updateConstraintsForSubtreeIfNeeded {
	rv := objc.Send[updateConstraintsForSubtreeIfNeeded](u_.ID(), selInit)
	return rv
}
