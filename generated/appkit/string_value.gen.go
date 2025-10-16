
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [stringValue] class.
var stringValueClass _stringValueClass

func init() {
	stringValueClass = _stringValueClass{objc.GetClass("stringValue")}
}

type _stringValueClass struct {
	objc.Class
}

// An interface definition for the [stringValue] class.
type IstringValue interface {
	ID() objc.ID
}

type stringValue struct {
	id objc.ID
}

func stringValueFrom(ptr unsafe.Pointer) stringValue {
	return stringValue{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (s_ stringValue) ID() objc.ID {
	return s_.id
}

// Alloc allocates a new instance without initialization.
func (sc _stringValueClass) Alloc() stringValue {
	rv := objc.Send[stringValue](objc.ID(sc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (sc _stringValueClass) New() stringValue {
	rv := objc.Send[stringValue](objc.ID(sc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewstringValue creates and returns a new initialized instance.
func NewstringValue() stringValue {
	return stringValueClass.New()
}

// Init initializes the instance.
func (s_ stringValue) Init() stringValue {
	rv := objc.Send[stringValue](s_.ID(), selInit)
	return rv
}
