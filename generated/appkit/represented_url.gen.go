
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [representedURL] class.
var representedURLClass _representedURLClass

func init() {
	representedURLClass = _representedURLClass{objc.GetClass("representedURL")}
}

type _representedURLClass struct {
	objc.Class
}

// An interface definition for the [representedURL] class.
type IrepresentedURL interface {
	ID() objc.ID
}

type representedURL struct {
	id objc.ID
}

func representedURLFrom(ptr unsafe.Pointer) representedURL {
	return representedURL{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (r_ representedURL) ID() objc.ID {
	return r_.id
}

// Alloc allocates a new instance without initialization.
func (rc _representedURLClass) Alloc() representedURL {
	rv := objc.Send[representedURL](objc.ID(rc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (rc _representedURLClass) New() representedURL {
	rv := objc.Send[representedURL](objc.ID(rc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewrepresentedURL creates and returns a new initialized instance.
func NewrepresentedURL() representedURL {
	return representedURLClass.New()
}

// Init initializes the instance.
func (r_ representedURL) Init() representedURL {
	rv := objc.Send[representedURL](r_.ID(), selInit)
	return rv
}
