
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [appearanceSource] class.
var appearanceSourceClass _appearanceSourceClass

func init() {
	appearanceSourceClass = _appearanceSourceClass{objc.GetClass("appearanceSource")}
}

type _appearanceSourceClass struct {
	objc.Class
}

// An interface definition for the [appearanceSource] class.
type IappearanceSource interface {
	ID() objc.ID
}

type appearanceSource struct {
	id objc.ID
}

func appearanceSourceFrom(ptr unsafe.Pointer) appearanceSource {
	return appearanceSource{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (a_ appearanceSource) ID() objc.ID {
	return a_.id
}

// Alloc allocates a new instance without initialization.
func (ac _appearanceSourceClass) Alloc() appearanceSource {
	rv := objc.Send[appearanceSource](objc.ID(ac.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (ac _appearanceSourceClass) New() appearanceSource {
	rv := objc.Send[appearanceSource](objc.ID(ac.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewappearanceSource creates and returns a new initialized instance.
func NewappearanceSource() appearanceSource {
	return appearanceSourceClass.New()
}

// Init initializes the instance.
func (a_ appearanceSource) Init() appearanceSource {
	rv := objc.Send[appearanceSource](a_.ID(), selInit)
	return rv
}
