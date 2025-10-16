
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [CollectionViewFlowLayout] class.
var CollectionViewFlowLayoutClass _CollectionViewFlowLayoutClass

func init() {
	CollectionViewFlowLayoutClass = _CollectionViewFlowLayoutClass{objc.GetClass("NSCollectionViewFlowLayout")}
}

type _CollectionViewFlowLayoutClass struct {
	objc.Class
}

// An interface definition for the [CollectionViewFlowLayout] class.
type ICollectionViewFlowLayout interface {
	ID() objc.ID
}

type CollectionViewFlowLayout struct {
	id objc.ID
}

func CollectionViewFlowLayoutFrom(ptr unsafe.Pointer) CollectionViewFlowLayout {
	return CollectionViewFlowLayout{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (c_ CollectionViewFlowLayout) ID() objc.ID {
	return c_.id
}

// Alloc allocates a new instance without initialization.
func (cc _CollectionViewFlowLayoutClass) Alloc() CollectionViewFlowLayout {
	rv := objc.Send[CollectionViewFlowLayout](objc.ID(cc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (cc _CollectionViewFlowLayoutClass) New() CollectionViewFlowLayout {
	rv := objc.Send[CollectionViewFlowLayout](objc.ID(cc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewCollectionViewFlowLayout creates and returns a new initialized instance.
func NewCollectionViewFlowLayout() CollectionViewFlowLayout {
	return CollectionViewFlowLayoutClass.New()
}

// Init initializes the instance.
func (c_ CollectionViewFlowLayout) Init() CollectionViewFlowLayout {
	rv := objc.Send[CollectionViewFlowLayout](c_.ID(), selInit)
	return rv
}
