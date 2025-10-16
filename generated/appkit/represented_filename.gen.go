
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [representedFilename] class.
var representedFilenameClass _representedFilenameClass

func init() {
	representedFilenameClass = _representedFilenameClass{objc.GetClass("representedFilename")}
}

type _representedFilenameClass struct {
	objc.Class
}

// An interface definition for the [representedFilename] class.
type IrepresentedFilename interface {
	ID() objc.ID
}

type representedFilename struct {
	id objc.ID
}

func representedFilenameFrom(ptr unsafe.Pointer) representedFilename {
	return representedFilename{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (r_ representedFilename) ID() objc.ID {
	return r_.id
}

// Alloc allocates a new instance without initialization.
func (rc _representedFilenameClass) Alloc() representedFilename {
	rv := objc.Send[representedFilename](objc.ID(rc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (rc _representedFilenameClass) New() representedFilename {
	rv := objc.Send[representedFilename](objc.ID(rc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewrepresentedFilename creates and returns a new initialized instance.
func NewrepresentedFilename() representedFilename {
	return representedFilenameClass.New()
}

// Init initializes the instance.
func (r_ representedFilename) Init() representedFilename {
	rv := objc.Send[representedFilename](r_.ID(), selInit)
	return rv
}
