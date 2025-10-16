
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [directoryURL] class.
var directoryURLClass _directoryURLClass

func init() {
	directoryURLClass = _directoryURLClass{objc.GetClass("directoryURL")}
}

type _directoryURLClass struct {
	objc.Class
}

// An interface definition for the [directoryURL] class.
type IdirectoryURL interface {
	ID() objc.ID
}

type directoryURL struct {
	id objc.ID
}

func directoryURLFrom(ptr unsafe.Pointer) directoryURL {
	return directoryURL{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (d_ directoryURL) ID() objc.ID {
	return d_.id
}

// Alloc allocates a new instance without initialization.
func (dc _directoryURLClass) Alloc() directoryURL {
	rv := objc.Send[directoryURL](objc.ID(dc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (dc _directoryURLClass) New() directoryURL {
	rv := objc.Send[directoryURL](objc.ID(dc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewdirectoryURL creates and returns a new initialized instance.
func NewdirectoryURL() directoryURL {
	return directoryURLClass.New()
}

// Init initializes the instance.
func (d_ directoryURL) Init() directoryURL {
	rv := objc.Send[directoryURL](d_.ID(), selInit)
	return rv
}
