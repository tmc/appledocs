
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [CollectionViewLayout] class.
var CollectionViewLayoutClass _CollectionViewLayoutClass

func init() {
	CollectionViewLayoutClass = _CollectionViewLayoutClass{objc.GetClass("NSCollectionViewLayout")}
}

type _CollectionViewLayoutClass struct {
	objc.Class
}

// An interface definition for the [CollectionViewLayout] class.
type ICollectionViewLayout interface {
	ID() objc.ID
}

type CollectionViewLayout struct {
	id objc.ID
}

func CollectionViewLayoutFrom(ptr unsafe.Pointer) CollectionViewLayout {
	return CollectionViewLayout{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (c_ CollectionViewLayout) ID() objc.ID {
	return c_.id
}

// Alloc allocates a new instance without initialization.
func (cc _CollectionViewLayoutClass) Alloc() CollectionViewLayout {
	rv := objc.Send[CollectionViewLayout](objc.ID(cc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (cc _CollectionViewLayoutClass) New() CollectionViewLayout {
	rv := objc.Send[CollectionViewLayout](objc.ID(cc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewCollectionViewLayout creates and returns a new initialized instance.
func NewCollectionViewLayout() CollectionViewLayout {
	return CollectionViewLayoutClass.New()
}

// Init initializes the instance.
func (c_ CollectionViewLayout) Init() CollectionViewLayout {
	rv := objc.Send[CollectionViewLayout](c_.ID(), selInit)
	return rv
}
