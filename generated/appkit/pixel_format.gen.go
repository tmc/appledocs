
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [pixelFormat] class.
var pixelFormatClass _pixelFormatClass

func init() {
	pixelFormatClass = _pixelFormatClass{objc.GetClass("pixelFormat")}
}

type _pixelFormatClass struct {
	objc.Class
}

// An interface definition for the [pixelFormat] class.
type IpixelFormat interface {
	ID() objc.ID
}

type pixelFormat struct {
	id objc.ID
}

func pixelFormatFrom(ptr unsafe.Pointer) pixelFormat {
	return pixelFormat{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (p_ pixelFormat) ID() objc.ID {
	return p_.id
}

// Alloc allocates a new instance without initialization.
func (pc _pixelFormatClass) Alloc() pixelFormat {
	rv := objc.Send[pixelFormat](objc.ID(pc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (pc _pixelFormatClass) New() pixelFormat {
	rv := objc.Send[pixelFormat](objc.ID(pc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewpixelFormat creates and returns a new initialized instance.
func NewpixelFormat() pixelFormat {
	return pixelFormatClass.New()
}

// Init initializes the instance.
func (p_ pixelFormat) Init() pixelFormat {
	rv := objc.Send[pixelFormat](p_.ID(), selInit)
	return rv
}
