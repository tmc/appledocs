
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [nameFieldLabel] class.
var nameFieldLabelClass _nameFieldLabelClass

func init() {
	nameFieldLabelClass = _nameFieldLabelClass{objc.GetClass("nameFieldLabel")}
}

type _nameFieldLabelClass struct {
	objc.Class
}

// An interface definition for the [nameFieldLabel] class.
type InameFieldLabel interface {
	ID() objc.ID
}

type nameFieldLabel struct {
	id objc.ID
}

func nameFieldLabelFrom(ptr unsafe.Pointer) nameFieldLabel {
	return nameFieldLabel{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (n_ nameFieldLabel) ID() objc.ID {
	return n_.id
}

// Alloc allocates a new instance without initialization.
func (nc _nameFieldLabelClass) Alloc() nameFieldLabel {
	rv := objc.Send[nameFieldLabel](objc.ID(nc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (nc _nameFieldLabelClass) New() nameFieldLabel {
	rv := objc.Send[nameFieldLabel](objc.ID(nc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewnameFieldLabel creates and returns a new initialized instance.
func NewnameFieldLabel() nameFieldLabel {
	return nameFieldLabelClass.New()
}

// Init initializes the instance.
func (n_ nameFieldLabel) Init() nameFieldLabel {
	rv := objc.Send[nameFieldLabel](n_.ID(), selInit)
	return rv
}
