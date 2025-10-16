
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [updateConstraints] class.
var updateConstraintsClass _updateConstraintsClass

func init() {
	updateConstraintsClass = _updateConstraintsClass{objc.GetClass("updateConstraints")}
}

type _updateConstraintsClass struct {
	objc.Class
}

// An interface definition for the [updateConstraints] class.
type IupdateConstraints interface {
	ID() objc.ID
}

type updateConstraints struct {
	id objc.ID
}

func updateConstraintsFrom(ptr unsafe.Pointer) updateConstraints {
	return updateConstraints{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (u_ updateConstraints) ID() objc.ID {
	return u_.id
}

// Alloc allocates a new instance without initialization.
func (uc _updateConstraintsClass) Alloc() updateConstraints {
	rv := objc.Send[updateConstraints](objc.ID(uc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (uc _updateConstraintsClass) New() updateConstraints {
	rv := objc.Send[updateConstraints](objc.ID(uc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewupdateConstraints creates and returns a new initialized instance.
func NewupdateConstraints() updateConstraints {
	return updateConstraintsClass.New()
}

// Init initializes the instance.
func (u_ updateConstraints) Init() updateConstraints {
	rv := objc.Send[updateConstraints](u_.ID(), selInit)
	return rv
}
