
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [CollectionViewItem] class.
var CollectionViewItemClass _CollectionViewItemClass

func init() {
	CollectionViewItemClass = _CollectionViewItemClass{objc.GetClass("NSCollectionViewItem")}
}

type _CollectionViewItemClass struct {
	objc.Class
}

// An interface definition for the [CollectionViewItem] class.
type ICollectionViewItem interface {
	ID() objc.ID
}

type CollectionViewItem struct {
	id objc.ID
}

func CollectionViewItemFrom(ptr unsafe.Pointer) CollectionViewItem {
	return CollectionViewItem{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (c_ CollectionViewItem) ID() objc.ID {
	return c_.id
}

// Alloc allocates a new instance without initialization.
func (cc _CollectionViewItemClass) Alloc() CollectionViewItem {
	rv := objc.Send[CollectionViewItem](objc.ID(cc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (cc _CollectionViewItemClass) New() CollectionViewItem {
	rv := objc.Send[CollectionViewItem](objc.ID(cc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewCollectionViewItem creates and returns a new initialized instance.
func NewCollectionViewItem() CollectionViewItem {
	return CollectionViewItemClass.New()
}

// Init initializes the instance.
func (c_ CollectionViewItem) Init() CollectionViewItem {
	rv := objc.Send[CollectionViewItem](c_.ID(), selInit)
	return rv
}
