
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [orderedIndex] class.
var orderedIndexClass _orderedIndexClass

func init() {
	orderedIndexClass = _orderedIndexClass{objc.GetClass("orderedIndex")}
}

type _orderedIndexClass struct {
	objc.Class
}

// An interface definition for the [orderedIndex] class.
type IorderedIndex interface {
	ID() objc.ID
}

type orderedIndex struct {
	id objc.ID
}

func orderedIndexFrom(ptr unsafe.Pointer) orderedIndex {
	return orderedIndex{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (o_ orderedIndex) ID() objc.ID {
	return o_.id
}

// Alloc allocates a new instance without initialization.
func (oc _orderedIndexClass) Alloc() orderedIndex {
	rv := objc.Send[orderedIndex](objc.ID(oc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (oc _orderedIndexClass) New() orderedIndex {
	rv := objc.Send[orderedIndex](objc.ID(oc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NeworderedIndex creates and returns a new initialized instance.
func NeworderedIndex() orderedIndex {
	return orderedIndexClass.New()
}

// Init initializes the instance.
func (o_ orderedIndex) Init() orderedIndex {
	rv := objc.Send[orderedIndex](o_.ID(), selInit)
	return rv
}
