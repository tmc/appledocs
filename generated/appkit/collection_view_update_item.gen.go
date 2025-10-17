
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [CollectionViewUpdateItem] class.
var CollectionViewUpdateItemClass _CollectionViewUpdateItemClass

func init() {
	CollectionViewUpdateItemClass = _CollectionViewUpdateItemClass{objc.GetClass("NSCollectionViewUpdateItem")}
}

type _CollectionViewUpdateItemClass struct {
	objc.Class
}

// An interface definition for the [CollectionViewUpdateItem] class.
type ICollectionViewUpdateItem interface {
	ID() objc.ID
}

type CollectionViewUpdateItem struct {
	id objc.ID
}

func CollectionViewUpdateItemFrom(ptr unsafe.Pointer) CollectionViewUpdateItem {
	return CollectionViewUpdateItem{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (c_ CollectionViewUpdateItem) ID() objc.ID {
	return c_.id
}

// Alloc allocates a new instance without initialization.
func (cc _CollectionViewUpdateItemClass) Alloc() CollectionViewUpdateItem {
	rv := objc.Send[CollectionViewUpdateItem](objc.ID(cc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (cc _CollectionViewUpdateItemClass) New() CollectionViewUpdateItem {
	rv := objc.Send[CollectionViewUpdateItem](objc.ID(cc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewCollectionViewUpdateItem creates and returns a new initialized instance.
func NewCollectionViewUpdateItem() CollectionViewUpdateItem {
	return CollectionViewUpdateItemClass.New()
}

// Init initializes the instance.
func (c_ CollectionViewUpdateItem) Init() CollectionViewUpdateItem {
	rv := objc.Send[CollectionViewUpdateItem](c_.ID(), selInit)
	return rv
}
