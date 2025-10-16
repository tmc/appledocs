
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [selectionIndexPaths] class.
var selectionIndexPathsClass _selectionIndexPathsClass

func init() {
	selectionIndexPathsClass = _selectionIndexPathsClass{objc.GetClass("selectionIndexPaths")}
}

type _selectionIndexPathsClass struct {
	objc.Class
}

// An interface definition for the [selectionIndexPaths] class.
type IselectionIndexPaths interface {
	ID() objc.ID
}

type selectionIndexPaths struct {
	id objc.ID
}

func selectionIndexPathsFrom(ptr unsafe.Pointer) selectionIndexPaths {
	return selectionIndexPaths{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (s_ selectionIndexPaths) ID() objc.ID {
	return s_.id
}

// Alloc allocates a new instance without initialization.
func (sc _selectionIndexPathsClass) Alloc() selectionIndexPaths {
	rv := objc.Send[selectionIndexPaths](objc.ID(sc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (sc _selectionIndexPathsClass) New() selectionIndexPaths {
	rv := objc.Send[selectionIndexPaths](objc.ID(sc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewselectionIndexPaths creates and returns a new initialized instance.
func NewselectionIndexPaths() selectionIndexPaths {
	return selectionIndexPathsClass.New()
}

// Init initializes the instance.
func (s_ selectionIndexPaths) Init() selectionIndexPaths {
	rv := objc.Send[selectionIndexPaths](s_.ID(), selInit)
	return rv
}
