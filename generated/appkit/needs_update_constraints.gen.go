
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [needsUpdateConstraints] class.
var needsUpdateConstraintsClass _needsUpdateConstraintsClass

func init() {
	needsUpdateConstraintsClass = _needsUpdateConstraintsClass{objc.GetClass("needsUpdateConstraints")}
}

type _needsUpdateConstraintsClass struct {
	objc.Class
}

// An interface definition for the [needsUpdateConstraints] class.
type IneedsUpdateConstraints interface {
	ID() objc.ID
}

type needsUpdateConstraints struct {
	id objc.ID
}

func needsUpdateConstraintsFrom(ptr unsafe.Pointer) needsUpdateConstraints {
	return needsUpdateConstraints{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (n_ needsUpdateConstraints) ID() objc.ID {
	return n_.id
}

// Alloc allocates a new instance without initialization.
func (nc _needsUpdateConstraintsClass) Alloc() needsUpdateConstraints {
	rv := objc.Send[needsUpdateConstraints](objc.ID(nc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (nc _needsUpdateConstraintsClass) New() needsUpdateConstraints {
	rv := objc.Send[needsUpdateConstraints](objc.ID(nc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewneedsUpdateConstraints creates and returns a new initialized instance.
func NewneedsUpdateConstraints() needsUpdateConstraints {
	return needsUpdateConstraintsClass.New()
}

// Init initializes the instance.
func (n_ needsUpdateConstraints) Init() needsUpdateConstraints {
	rv := objc.Send[needsUpdateConstraints](n_.ID(), selInit)
	return rv
}
