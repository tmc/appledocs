
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [indexPathsForVisibleItems] class.
var indexPathsForVisibleItemsClass _indexPathsForVisibleItemsClass

func init() {
	indexPathsForVisibleItemsClass = _indexPathsForVisibleItemsClass{objc.GetClass("indexPathsForVisibleItems")}
}

type _indexPathsForVisibleItemsClass struct {
	objc.Class
}

// An interface definition for the [indexPathsForVisibleItems] class.
type IindexPathsForVisibleItems interface {
	ID() objc.ID
}

type indexPathsForVisibleItems struct {
	id objc.ID
}

func indexPathsForVisibleItemsFrom(ptr unsafe.Pointer) indexPathsForVisibleItems {
	return indexPathsForVisibleItems{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (i_ indexPathsForVisibleItems) ID() objc.ID {
	return i_.id
}

// Alloc allocates a new instance without initialization.
func (ic _indexPathsForVisibleItemsClass) Alloc() indexPathsForVisibleItems {
	rv := objc.Send[indexPathsForVisibleItems](objc.ID(ic.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (ic _indexPathsForVisibleItemsClass) New() indexPathsForVisibleItems {
	rv := objc.Send[indexPathsForVisibleItems](objc.ID(ic.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewindexPathsForVisibleItems creates and returns a new initialized instance.
func NewindexPathsForVisibleItems() indexPathsForVisibleItems {
	return indexPathsForVisibleItemsClass.New()
}

// Init initializes the instance.
func (i_ indexPathsForVisibleItems) Init() indexPathsForVisibleItems {
	rv := objc.Send[indexPathsForVisibleItems](i_.ID(), selInit)
	return rv
}
