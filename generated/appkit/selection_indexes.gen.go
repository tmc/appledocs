
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [selectionIndexes] class.
var selectionIndexesClass _selectionIndexesClass

func init() {
	selectionIndexesClass = _selectionIndexesClass{objc.GetClass("selectionIndexes")}
}

type _selectionIndexesClass struct {
	objc.Class
}

// An interface definition for the [selectionIndexes] class.
type IselectionIndexes interface {
	ID() objc.ID
}

type selectionIndexes struct {
	id objc.ID
}

func selectionIndexesFrom(ptr unsafe.Pointer) selectionIndexes {
	return selectionIndexes{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (s_ selectionIndexes) ID() objc.ID {
	return s_.id
}

// Alloc allocates a new instance without initialization.
func (sc _selectionIndexesClass) Alloc() selectionIndexes {
	rv := objc.Send[selectionIndexes](objc.ID(sc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (sc _selectionIndexesClass) New() selectionIndexes {
	rv := objc.Send[selectionIndexes](objc.ID(sc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewselectionIndexes creates and returns a new initialized instance.
func NewselectionIndexes() selectionIndexes {
	return selectionIndexesClass.New()
}

// Init initializes the instance.
func (s_ selectionIndexes) Init() selectionIndexes {
	rv := objc.Send[selectionIndexes](s_.ID(), selInit)
	return rv
}
