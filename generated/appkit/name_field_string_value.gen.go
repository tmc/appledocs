
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [nameFieldStringValue] class.
var nameFieldStringValueClass _nameFieldStringValueClass

func init() {
	nameFieldStringValueClass = _nameFieldStringValueClass{objc.GetClass("nameFieldStringValue")}
}

type _nameFieldStringValueClass struct {
	objc.Class
}

// An interface definition for the [nameFieldStringValue] class.
type InameFieldStringValue interface {
	ID() objc.ID
}

type nameFieldStringValue struct {
	id objc.ID
}

func nameFieldStringValueFrom(ptr unsafe.Pointer) nameFieldStringValue {
	return nameFieldStringValue{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (n_ nameFieldStringValue) ID() objc.ID {
	return n_.id
}

// Alloc allocates a new instance without initialization.
func (nc _nameFieldStringValueClass) Alloc() nameFieldStringValue {
	rv := objc.Send[nameFieldStringValue](objc.ID(nc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (nc _nameFieldStringValueClass) New() nameFieldStringValue {
	rv := objc.Send[nameFieldStringValue](objc.ID(nc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewnameFieldStringValue creates and returns a new initialized instance.
func NewnameFieldStringValue() nameFieldStringValue {
	return nameFieldStringValueClass.New()
}

// Init initializes the instance.
func (n_ nameFieldStringValue) Init() nameFieldStringValue {
	rv := objc.Send[nameFieldStringValue](n_.ID(), selInit)
	return rv
}
