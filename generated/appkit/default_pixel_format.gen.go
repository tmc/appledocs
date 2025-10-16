
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [defaultPixelFormat] class.
var defaultPixelFormatClass _defaultPixelFormatClass

func init() {
	defaultPixelFormatClass = _defaultPixelFormatClass{objc.GetClass("defaultPixelFormat")}
}

type _defaultPixelFormatClass struct {
	objc.Class
}

// An interface definition for the [defaultPixelFormat] class.
type IdefaultPixelFormat interface {
	ID() objc.ID
}

type defaultPixelFormat struct {
	id objc.ID
}

func defaultPixelFormatFrom(ptr unsafe.Pointer) defaultPixelFormat {
	return defaultPixelFormat{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (d_ defaultPixelFormat) ID() objc.ID {
	return d_.id
}

// Alloc allocates a new instance without initialization.
func (dc _defaultPixelFormatClass) Alloc() defaultPixelFormat {
	rv := objc.Send[defaultPixelFormat](objc.ID(dc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (dc _defaultPixelFormatClass) New() defaultPixelFormat {
	rv := objc.Send[defaultPixelFormat](objc.ID(dc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewdefaultPixelFormat creates and returns a new initialized instance.
func NewdefaultPixelFormat() defaultPixelFormat {
	return defaultPixelFormatClass.New()
}

// Init initializes the instance.
func (d_ defaultPixelFormat) Init() defaultPixelFormat {
	rv := objc.Send[defaultPixelFormat](d_.ID(), selInit)
	return rv
}
