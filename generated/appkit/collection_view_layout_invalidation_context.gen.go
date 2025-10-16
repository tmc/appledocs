
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [CollectionViewLayoutInvalidationContext] class.
var CollectionViewLayoutInvalidationContextClass _CollectionViewLayoutInvalidationContextClass

func init() {
	CollectionViewLayoutInvalidationContextClass = _CollectionViewLayoutInvalidationContextClass{objc.GetClass("NSCollectionViewLayoutInvalidationContext")}
}

type _CollectionViewLayoutInvalidationContextClass struct {
	objc.Class
}

// An interface definition for the [CollectionViewLayoutInvalidationContext] class.
type ICollectionViewLayoutInvalidationContext interface {
	ID() objc.ID
}

type CollectionViewLayoutInvalidationContext struct {
	id objc.ID
}

func CollectionViewLayoutInvalidationContextFrom(ptr unsafe.Pointer) CollectionViewLayoutInvalidationContext {
	return CollectionViewLayoutInvalidationContext{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (c_ CollectionViewLayoutInvalidationContext) ID() objc.ID {
	return c_.id
}

// Alloc allocates a new instance without initialization.
func (cc _CollectionViewLayoutInvalidationContextClass) Alloc() CollectionViewLayoutInvalidationContext {
	rv := objc.Send[CollectionViewLayoutInvalidationContext](objc.ID(cc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (cc _CollectionViewLayoutInvalidationContextClass) New() CollectionViewLayoutInvalidationContext {
	rv := objc.Send[CollectionViewLayoutInvalidationContext](objc.ID(cc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewCollectionViewLayoutInvalidationContext creates and returns a new initialized instance.
func NewCollectionViewLayoutInvalidationContext() CollectionViewLayoutInvalidationContext {
	return CollectionViewLayoutInvalidationContextClass.New()
}

// Init initializes the instance.
func (c_ CollectionViewLayoutInvalidationContext) Init() CollectionViewLayoutInvalidationContext {
	rv := objc.Send[CollectionViewLayoutInvalidationContext](c_.ID(), selInit)
	return rv
}
