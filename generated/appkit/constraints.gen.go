
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [constraints] class.
var constraintsClass _constraintsClass

func init() {
	constraintsClass = _constraintsClass{objc.GetClass("constraints")}
}

type _constraintsClass struct {
	objc.Class
}

// An interface definition for the [constraints] class.
type Iconstraints interface {
	ID() objc.ID
}

type constraints struct {
	id objc.ID
}

func constraintsFrom(ptr unsafe.Pointer) constraints {
	return constraints{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (c_ constraints) ID() objc.ID {
	return c_.id
}

// Alloc allocates a new instance without initialization.
func (cc _constraintsClass) Alloc() constraints {
	rv := objc.Send[constraints](objc.ID(cc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (cc _constraintsClass) New() constraints {
	rv := objc.Send[constraints](objc.ID(cc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// Newconstraints creates and returns a new initialized instance.
func Newconstraints() constraints {
	return constraintsClass.New()
}

// Init initializes the instance.
func (c_ constraints) Init() constraints {
	rv := objc.Send[constraints](c_.ID(), selInit)
	return rv
}
