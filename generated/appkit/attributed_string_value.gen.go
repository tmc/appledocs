
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [attributedStringValue] class.
var attributedStringValueClass _attributedStringValueClass

func init() {
	attributedStringValueClass = _attributedStringValueClass{objc.GetClass("attributedStringValue")}
}

type _attributedStringValueClass struct {
	objc.Class
}

// An interface definition for the [attributedStringValue] class.
type IattributedStringValue interface {
	ID() objc.ID
}

type attributedStringValue struct {
	id objc.ID
}

func attributedStringValueFrom(ptr unsafe.Pointer) attributedStringValue {
	return attributedStringValue{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (a_ attributedStringValue) ID() objc.ID {
	return a_.id
}

// Alloc allocates a new instance without initialization.
func (ac _attributedStringValueClass) Alloc() attributedStringValue {
	rv := objc.Send[attributedStringValue](objc.ID(ac.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (ac _attributedStringValueClass) New() attributedStringValue {
	rv := objc.Send[attributedStringValue](objc.ID(ac.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewattributedStringValue creates and returns a new initialized instance.
func NewattributedStringValue() attributedStringValue {
	return attributedStringValueClass.New()
}

// Init initializes the instance.
func (a_ attributedStringValue) Init() attributedStringValue {
	rv := objc.Send[attributedStringValue](a_.ID(), selInit)
	return rv
}
