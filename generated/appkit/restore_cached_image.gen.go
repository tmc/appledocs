
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [restoreCachedImage] class.
var restoreCachedImageClass _restoreCachedImageClass

func init() {
	restoreCachedImageClass = _restoreCachedImageClass{objc.GetClass("restoreCachedImage")}
}

type _restoreCachedImageClass struct {
	objc.Class
}

// An interface definition for the [restoreCachedImage] class.
type IrestoreCachedImage interface {
	ID() objc.ID
}

type restoreCachedImage struct {
	id objc.ID
}

func restoreCachedImageFrom(ptr unsafe.Pointer) restoreCachedImage {
	return restoreCachedImage{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (r_ restoreCachedImage) ID() objc.ID {
	return r_.id
}

// Alloc allocates a new instance without initialization.
func (rc _restoreCachedImageClass) Alloc() restoreCachedImage {
	rv := objc.Send[restoreCachedImage](objc.ID(rc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (rc _restoreCachedImageClass) New() restoreCachedImage {
	rv := objc.Send[restoreCachedImage](objc.ID(rc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewrestoreCachedImage creates and returns a new initialized instance.
func NewrestoreCachedImage() restoreCachedImage {
	return restoreCachedImageClass.New()
}

// Init initializes the instance.
func (r_ restoreCachedImage) Init() restoreCachedImage {
	rv := objc.Send[restoreCachedImage](r_.ID(), selInit)
	return rv
}
