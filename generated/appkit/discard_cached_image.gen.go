
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [discardCachedImage] class.
var discardCachedImageClass _discardCachedImageClass

func init() {
	discardCachedImageClass = _discardCachedImageClass{objc.GetClass("discardCachedImage")}
}

type _discardCachedImageClass struct {
	objc.Class
}

// An interface definition for the [discardCachedImage] class.
type IdiscardCachedImage interface {
	ID() objc.ID
}

type discardCachedImage struct {
	id objc.ID
}

func discardCachedImageFrom(ptr unsafe.Pointer) discardCachedImage {
	return discardCachedImage{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (d_ discardCachedImage) ID() objc.ID {
	return d_.id
}

// Alloc allocates a new instance without initialization.
func (dc _discardCachedImageClass) Alloc() discardCachedImage {
	rv := objc.Send[discardCachedImage](objc.ID(dc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (dc _discardCachedImageClass) New() discardCachedImage {
	rv := objc.Send[discardCachedImage](objc.ID(dc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewdiscardCachedImage creates and returns a new initialized instance.
func NewdiscardCachedImage() discardCachedImage {
	return discardCachedImageClass.New()
}

// Init initializes the instance.
func (d_ discardCachedImage) Init() discardCachedImage {
	rv := objc.Send[discardCachedImage](d_.ID(), selInit)
	return rv
}
