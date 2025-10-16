
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [timestamp] class.
var timestampClass _timestampClass

func init() {
	timestampClass = _timestampClass{objc.GetClass("timestamp")}
}

type _timestampClass struct {
	objc.Class
}

// An interface definition for the [timestamp] class.
type Itimestamp interface {
	ID() objc.ID
}

type timestamp struct {
	id objc.ID
}

func timestampFrom(ptr unsafe.Pointer) timestamp {
	return timestamp{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (t_ timestamp) ID() objc.ID {
	return t_.id
}

// Alloc allocates a new instance without initialization.
func (tc _timestampClass) Alloc() timestamp {
	rv := objc.Send[timestamp](objc.ID(tc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (tc _timestampClass) New() timestamp {
	rv := objc.Send[timestamp](objc.ID(tc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// Newtimestamp creates and returns a new initialized instance.
func Newtimestamp() timestamp {
	return timestampClass.New()
}

// Init initializes the instance.
func (t_ timestamp) Init() timestamp {
	rv := objc.Send[timestamp](t_.ID(), selInit)
	return rv
}
