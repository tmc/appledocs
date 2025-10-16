
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [objectValue] class.
var objectValueClass _objectValueClass

func init() {
	objectValueClass = _objectValueClass{objc.GetClass("objectValue")}
}

type _objectValueClass struct {
	objc.Class
}

// An interface definition for the [objectValue] class.
type IobjectValue interface {
	ID() objc.ID
}

type objectValue struct {
	id objc.ID
}

func objectValueFrom(ptr unsafe.Pointer) objectValue {
	return objectValue{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (o_ objectValue) ID() objc.ID {
	return o_.id
}

// Alloc allocates a new instance without initialization.
func (oc _objectValueClass) Alloc() objectValue {
	rv := objc.Send[objectValue](objc.ID(oc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (oc _objectValueClass) New() objectValue {
	rv := objc.Send[objectValue](objc.ID(oc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewobjectValue creates and returns a new initialized instance.
func NewobjectValue() objectValue {
	return objectValueClass.New()
}

// Init initializes the instance.
func (o_ objectValue) Init() objectValue {
	rv := objc.Send[objectValue](o_.ID(), selInit)
	return rv
}
