
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [requiresConstraintBasedLayout] class.
var requiresConstraintBasedLayoutClass _requiresConstraintBasedLayoutClass

func init() {
	requiresConstraintBasedLayoutClass = _requiresConstraintBasedLayoutClass{objc.GetClass("requiresConstraintBasedLayout")}
}

type _requiresConstraintBasedLayoutClass struct {
	objc.Class
}

// An interface definition for the [requiresConstraintBasedLayout] class.
type IrequiresConstraintBasedLayout interface {
	ID() objc.ID
}

type requiresConstraintBasedLayout struct {
	id objc.ID
}

func requiresConstraintBasedLayoutFrom(ptr unsafe.Pointer) requiresConstraintBasedLayout {
	return requiresConstraintBasedLayout{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (r_ requiresConstraintBasedLayout) ID() objc.ID {
	return r_.id
}

// Alloc allocates a new instance without initialization.
func (rc _requiresConstraintBasedLayoutClass) Alloc() requiresConstraintBasedLayout {
	rv := objc.Send[requiresConstraintBasedLayout](objc.ID(rc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (rc _requiresConstraintBasedLayoutClass) New() requiresConstraintBasedLayout {
	rv := objc.Send[requiresConstraintBasedLayout](objc.ID(rc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewrequiresConstraintBasedLayout creates and returns a new initialized instance.
func NewrequiresConstraintBasedLayout() requiresConstraintBasedLayout {
	return requiresConstraintBasedLayoutClass.New()
}

// Init initializes the instance.
func (r_ requiresConstraintBasedLayout) Init() requiresConstraintBasedLayout {
	rv := objc.Send[requiresConstraintBasedLayout](r_.ID(), selInit)
	return rv
}
